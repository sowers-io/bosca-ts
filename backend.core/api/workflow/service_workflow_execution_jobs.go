package workflow

import (
	grpc "bosca.io/api/protobuf/bosca/workflow"
	"log/slog"
)

func (svc *service) claimNextJob(svr grpc.WorkflowService_GetWorkflowActivityJobsServer, workerId string, request *grpc.WorkflowActivityJobRequest) error {
	txn, err := svc.ds.NewTransaction(svr.Context())
	if err != nil {
		return err
	}
	job, err := svc.ds.ClaimNextJob(svr.Context(), txn, workerId, request.Queue, request.ActivityId)
	if err != nil {
		txn.Rollback()
		return err
	}
	err = txn.Commit()
	if err != nil {
		return err
	}
	if job != nil {
		err = svr.Send(job)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) GetWorkflowActivityJobs(request *grpc.WorkflowActivityJobRequest, svr grpc.WorkflowService_GetWorkflowActivityJobsServer) error {
	workerId, err := svc.ds.RegisterWorker(svr.Context(), request)
	if err != nil {
		return err
	}
	queueChannel := svc.ds.GetQueueChannel(request.Queue)
	if err = svc.claimNextJob(svr, workerId, request); err != nil {
		return err
	}
	for {
		select {
		case <-svr.Context().Done():
			return svc.ds.UnregisterWorker(svr.Context(), workerId)
		case notification := <-queueChannel:
			slog.InfoContext(svr.Context(), "checking for new jobs", slog.String("queue", request.Queue), slog.String("executionId", notification.ExecutionId))
			if err = svc.claimNextJob(svr, workerId, request); err != nil {
				return err
			}
		}
	}
}
