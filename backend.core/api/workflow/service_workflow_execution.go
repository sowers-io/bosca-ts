package workflow

import (
	protobuf "bosca.io/api/protobuf/bosca"
	grpc "bosca.io/api/protobuf/bosca/workflow"
	"context"
	"log/slog"
)

func (svc *service) executeWorkflow(ctx context.Context, metadataId string, workflowId string, context map[string]string, waitForCompletion bool) error {
	executionContext, err := svc.getNewWorkflowExecutionContext(ctx, workflowId, metadataId, context)
	if err != nil {
		return err
	}

	txn, err := svc.ds.NewTransaction(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to start workflow transaction", slog.Any("error", err), slog.String("metadataId", metadataId), slog.String("workflowId", workflowId))
		if txn != nil {
			txn.Rollback()
		}
		return err
	}

	executionId, err := svc.ds.AddWorkflowExecution(ctx, txn, workflowId, metadataId, executionContext)
	if err != nil {
		slog.ErrorContext(ctx, "failed to add workflow execution", slog.Any("error", err), slog.String("metadataId", metadataId), slog.String("workflowId", workflowId))
		txn.Rollback()
		return err
	}

	executionContext, err = svc.ds.GetWorkflowExecutionContextForUpdate(ctx, txn, executionId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow execution for update", slog.Any("error", err), slog.String("metadataId", metadataId), slog.String("workflowId", workflowId))
		return err
	}

	complete, queues, err := svc.ds.QueueNextWorkflowJobs(ctx, txn, executionContext)
	if err != nil {
		slog.ErrorContext(ctx, "failed to queue workflow execution jobs", slog.Any("error", err), slog.String("metadataId", metadataId), slog.String("workflowId", workflowId))
		txn.Rollback()
		return err
	}

	err = txn.Commit()
	if err != nil {
		slog.ErrorContext(ctx, "failed to commit workflow execution", slog.Any("error", err), slog.String("metadataId", metadataId), slog.String("workflowId", workflowId))
		return err
	}

	if complete {
		queues = getAllQueues(executionContext)
		err = svc.ds.NotifyExecutionCompletion(ctx, queues, executionId)
		if err != nil {
			slog.ErrorContext(ctx, "failed to notify of execution completion", slog.Any("error", err), slog.String("metadataId", metadataId), slog.String("workflowId", workflowId))
			return err
		}
		return nil
	} else {
		err = svc.ds.NotifyJobAvailable(ctx, queues)
		if err != nil {
			slog.ErrorContext(ctx, "failed to notify of job availability", slog.Any("error", err), slog.String("metadataId", metadataId), slog.String("workflowId", workflowId))
			return err
		}
	}

	if waitForCompletion {
		err = svc.ds.WaitForExecutionCompletion(executionId)
		if err != nil {
			slog.ErrorContext(ctx, "failed to wait for execution completion", slog.Any("error", err), slog.String("metadataId", metadataId), slog.String("workflowId", workflowId))
		}
	}
	return err
}

func (svc *service) ExecuteWorkflow(ctx context.Context, request *grpc.WorkflowExecutionRequest) (*protobuf.Empty, error) {
	err := svc.executeWorkflow(ctx, request.MetadataId, request.WorkflowId, request.Context, request.WaitForCompletion)
	if err != nil {
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) SetWorkflowActivityJobStatus(ctx context.Context, request *grpc.WorkflowActivityJobStatus) (*protobuf.Empty, error) {
	txn, err := svc.ds.NewTransaction(ctx)
	if err != nil {
		return nil, err
	}

	executionContext, err := svc.ds.GetWorkflowExecutionContextForUpdate(ctx, txn, request.ExecutionId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow execution for update", slog.Any("error", err), slog.String("executionId", request.ExecutionId))
		txn.Rollback()
		return nil, err
	}

	isCurrentExecutionGroupComplete, err := svc.ds.IsWorkflowExecutionGroupComplete(ctx, txn, request.ExecutionId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to check if workflow execution group is complete", slog.Any("error", err), slog.String("executionId", request.ExecutionId))
		txn.Rollback()
		return nil, err
	}

	var complete bool
	var queues []string
	if isCurrentExecutionGroupComplete {
		complete, queues, err = svc.ds.QueueNextWorkflowJobs(ctx, txn, executionContext)
		if err != nil {
			slog.ErrorContext(ctx, "failed to queue workflow jobs", slog.Any("error", err), slog.String("executionId", request.ExecutionId))
			txn.Rollback()
			return nil, err
		}
	}

	err = txn.Commit()
	if err != nil {
		return nil, err
	}

	if complete {
		// TODO: probably notifying too much
		if queues == nil {
			queues = getAllQueues(executionContext)
		}
		err = svc.ds.NotifyExecutionCompletion(ctx, queues, request.ExecutionId)
		if err != nil {
			slog.ErrorContext(ctx, "failed to notify execution completion", slog.Any("error", err), slog.String("executionId", request.ExecutionId))
			return nil, err
		}
	} else if isCurrentExecutionGroupComplete {
		err = svc.ds.NotifyJobAvailable(ctx, queues)
		if err != nil {
			slog.ErrorContext(ctx, "failed to notify jobs are complete", slog.Any("error", err), slog.String("executionId", request.ExecutionId))
			return nil, err
		}
	}

	return &protobuf.Empty{}, nil
}
