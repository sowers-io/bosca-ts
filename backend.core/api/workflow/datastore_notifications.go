package workflow

import (
	"context"
	"database/sql/driver"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5/stdlib"
	"io"
	"log/slog"
)

type ExecutionNotificationType int

const (
	ExecutionCompletion ExecutionNotificationType = iota
	JobAvailable
	ExecutionFailed
)

type ExecutionNotification struct {
	Type        ExecutionNotificationType
	Success     bool
	ExecutionId string
	Error       string
}

func (ds *DataStore) GetQueueChannel(queue string) chan *ExecutionNotification {
	ds.executionChannelMutex.Lock()
	defer ds.executionChannelMutex.Unlock()

	queueChannel := ds.queueChannels[queue]
	if queueChannel == nil {
		queueChannel = make(chan *ExecutionNotification)
		ds.queueChannels[queue] = queueChannel
	}

	return queueChannel
}

func (ds *DataStore) ListenForCompletions(ctx context.Context) error {
	dbConn, err := ds.db.Conn(ctx)
	if err != nil {
		return err
	}
	err = dbConn.Raw(func(driverConn any) error {
		pgxConn := driverConn.(*stdlib.Conn)

		rows, err := pgxConn.QueryContext(ctx, "select queue from workflows", nil)
		if err != nil {
			return err
		}

		queues := make([][]driver.Value, 0)
		for {
			queue := make([]driver.Value, 1)
			err = rows.Next(queue)
			if err != nil {
				if errors.Is(err, io.EOF) {
					break
				} else {
					rows.Close()
					return err
				}
			}
			queues = append(queues, queue)
		}
		rows.Close()

		for _, queue := range queues {
			_, err = pgxConn.ExecContext(ctx, "select listen($1)", []driver.NamedValue{
				{Name: "1", Ordinal: 1, Value: queue[0]},
			})
			if err != nil {
				return err
			}
		}

		conn := pgxConn.Conn()
		for {
			notification, err := conn.WaitForNotification(ctx)
			if err != nil {
				return err
			}
			executionNotification := &ExecutionNotification{}
			err = json.Unmarshal([]byte(notification.Payload), &executionNotification)
			if err != nil {
				return err
			}
			switch executionNotification.Type {
			case ExecutionFailed:
				ds.OnExecutionFailed(ctx, executionNotification)
			case ExecutionCompletion:
				ds.OnExecutionCompletion(ctx, executionNotification)
				ds.OnJobAvailable(ctx, notification.Channel, executionNotification)
			case JobAvailable:
				ds.OnJobAvailable(ctx, notification.Channel, executionNotification)
			}
		}
	})
	dbConn.Close()
	return err
}

func (ds *DataStore) WaitForExecutionCompletion(executionId string) *ExecutionNotification {
	ds.executionChannelMutex.Lock()
	executionChannel := ds.executionChannels[executionId]
	if executionChannel == nil {
		executionChannel = make(chan *ExecutionNotification)
		ds.executionChannels[executionId] = executionChannel
	}
	ds.executionChannelMutex.Unlock()

	notification := <-executionChannel

	ds.executionChannelMutex.Lock()
	delete(ds.executionChannels, executionId)
	close(executionChannel)
	ds.executionChannelMutex.Unlock()

	return notification
}

func (ds *DataStore) notify(ctx context.Context, queues []string, notification *ExecutionNotification) error {
	notificationJson, err := json.Marshal(notification)
	if err != nil {
		return err
	}
	for _, queue := range queues {
		_, err = ds.db.ExecContext(ctx, "select pg_notify($1, $2)", queue, notificationJson)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ds *DataStore) NotifyJobAvailable(ctx context.Context, queues []string) error {
	notification := &ExecutionNotification{
		Type: JobAvailable,
	}
	return ds.notify(ctx, queues, notification)
}

func (ds *DataStore) OnJobAvailable(ctx context.Context, queue string, notification *ExecutionNotification) {
	slog.InfoContext(ctx, "OnJobAvailable", slog.String("queue", queue), slog.Any("notification", notification))
	queueChannel := ds.GetQueueChannel(queue)
	queueChannel <- notification
}

func (ds *DataStore) NotifyExecutionCompletion(ctx context.Context, queues []string, executionId string, success bool) error {
	notification := &ExecutionNotification{
		Type:        ExecutionCompletion,
		ExecutionId: executionId,
		Success:     success,
	}
	return ds.notify(ctx, queues, notification)
}

func (ds *DataStore) OnExecutionCompletion(ctx context.Context, notification *ExecutionNotification) {
	slog.InfoContext(ctx, "OnExecutionCompletion", slog.Any("notification", notification))
	ds.executionChannelMutex.Lock()
	defer ds.executionChannelMutex.Unlock()
	executionChannel := ds.executionChannels[notification.ExecutionId]
	if executionChannel != nil {
		executionChannel <- notification
	}
}

func (ds *DataStore) NotifyExecutionFailed(ctx context.Context, queues []string, executionId string, err error) error {
	notification := &ExecutionNotification{
		Type:        ExecutionFailed,
		ExecutionId: executionId,
	}
	if err != nil {
		notification.Error = err.Error()
	}
	return ds.notify(ctx, queues, notification)
}

func (ds *DataStore) OnExecutionFailed(ctx context.Context, notification *ExecutionNotification) {
	slog.InfoContext(ctx, "OnExecutionFailed", slog.Any("notification", notification))
	ds.executionChannelMutex.Lock()
	defer ds.executionChannelMutex.Unlock()
	executionChannel := ds.executionChannels[notification.ExecutionId]
	if executionChannel != nil {
		executionChannel <- notification
	}
}
