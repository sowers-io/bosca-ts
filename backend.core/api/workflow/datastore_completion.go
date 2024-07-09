package workflow

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
)

type ExecutionNotificationType int

const (
	ExecutionCompletion ExecutionNotificationType = iota
	JobAvailable
	ExecutionFailed
)

type ExecutionNotification struct {
	Type        ExecutionNotificationType
	ExecutionId string
	Error       string
}

func (ds *DataStore) GetQueueChannel(queue string) chan string {
	ds.executionChannelMutex.Lock()
	defer ds.executionChannelMutex.Unlock()

	queueChannel := ds.queueChannels[queue]
	if queueChannel == nil {
		queueChannel = make(chan string)
		ds.queueChannels[queue] = queueChannel
	}

	return queueChannel
}

func (ds *DataStore) ListenForCompletions(ctx context.Context) error {
	conn, err := ds.db.Conn(ctx)
	if err != nil {
		return err
	}
	return conn.Raw(func(driverConn any) error {
		pgxConn := driverConn.(*pgxpool.Conn)
		_, err = ds.db.ExecContext(ctx, "select listen_all()")
		if err != nil {
			return err
		}
		errChan := make(chan error)
		go func() {
			ctx := context.Background()
			conn := pgxConn.Conn()
			for {
				notification, err := conn.WaitForNotification(ctx)
				if err != nil {
					errChan <- err
					return
				}
				executionNotification := &ExecutionNotification{}
				err = json.Unmarshal([]byte(notification.Payload), &executionNotification)
				if err != nil {
					errChan <- err
					return
				}
				switch executionNotification.Type {
				case ExecutionFailed:
					ds.OnExecutionFailed(ctx, executionNotification)
				case ExecutionCompletion:
					ds.OnExecutionCompletion(ctx, executionNotification)
				case JobAvailable:
					ds.OnJobAvailable(ctx, notification.Channel, executionNotification)
				}
			}
		}()
		err := <-errChan
		close(errChan)
		return err
	})
}

func (ds *DataStore) WaitForExecutionCompletion(executionId string) error {
	ds.executionChannelMutex.Lock()
	executionChannel := ds.executionChannels[executionId]
	if executionChannel == nil {
		executionChannel = make(chan error)
		ds.executionChannels[executionId] = executionChannel
	}
	ds.executionChannelMutex.Unlock()

	err := <-executionChannel

	ds.executionChannelMutex.Lock()
	delete(ds.executionChannels, executionId)
	close(executionChannel)
	ds.executionChannelMutex.Unlock()

	return err
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
	queueChannel := ds.GetQueueChannel(queue)
	queueChannel <- notification.ExecutionId
}

func (ds *DataStore) NotifyExecutionCompletion(ctx context.Context, queues []string, executionId string) error {
	notification := &ExecutionNotification{
		Type:        ExecutionCompletion,
		ExecutionId: executionId,
	}
	return ds.notify(ctx, queues, notification)
}

func (ds *DataStore) OnExecutionCompletion(ctx context.Context, notification *ExecutionNotification) {
	ds.executionChannelMutex.Lock()
	defer ds.executionChannelMutex.Unlock()
	executionChannel := ds.executionChannels[notification.ExecutionId]
	if executionChannel != nil {
		executionChannel <- nil
	}
}

func (ds *DataStore) NotifyExecutionFailed(ctx context.Context, queues []string, executionId string, err error) error {
	notification := &ExecutionNotification{
		Type:        ExecutionFailed,
		ExecutionId: executionId,
	}
	return ds.notify(ctx, queues, notification)
}

func (ds *DataStore) OnExecutionFailed(ctx context.Context, notification *ExecutionNotification) {
	ds.executionChannelMutex.Lock()
	defer ds.executionChannelMutex.Unlock()
	executionChannel := ds.executionChannels[notification.ExecutionId]
	if executionChannel != nil {
		executionChannel <- errors.New(notification.Error)
	}
}
