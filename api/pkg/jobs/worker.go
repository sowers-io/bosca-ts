package jobs

import (
	protojobs "bosca.io/api/protobuf/jobs"
	"bosca.io/pkg/clients"
	"bosca.io/pkg/configuration"
	"context"
	"google.golang.org/grpc"
)

type Worker struct {
	queue      string
	connection *grpc.ClientConn
	client     protojobs.JobsServiceClient
	pollClient protojobs.JobsService_PollClient
}

func NewWorker(ctx context.Context, cfg *configuration.WorkerConfiguration, queue string, timeout int64) (*Worker, error) {
	if timeout == 0 {
		timeout = 120
	}
	connection, err := clients.NewClientConnection(cfg.ClientEndPoints.JobsApiAddress)
	if err != nil {
		return nil, err
	}
	client := protojobs.NewJobsServiceClient(connection)
	pollClient, err := client.Poll(ctx, &protojobs.PollRequest{
		Queue:   queue,
		Timeout: timeout,
	})
	if err != nil {
		return nil, err
	}
	return &Worker{
		queue:      queue,
		connection: connection,
		client:     client,
		pollClient: pollClient,
	}, nil
}

func (w *Worker) Close() {
	w.connection.Close()
}

func (w *Worker) Work(ctx context.Context, worker func(job *protojobs.Job) error) error {
	for {
		job, err := w.pollClient.Recv()
		if err != nil {
			return err
		}
		err = worker(job)
		_, err = w.client.Finish(ctx, &protojobs.FinishRequest{
			Queue:   w.queue,
			Id:      job.Id,
			Success: err == nil,
		})
		if err != nil {
			return err
		}
	}
}
