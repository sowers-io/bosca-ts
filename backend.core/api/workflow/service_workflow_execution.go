/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package workflow

import (
	protobuf "bosca.io/api/protobuf/bosca"
	grpc "bosca.io/api/protobuf/bosca/workflow"
	"context"
	"database/sql"
	"log/slog"
)

func (svc *service) executeWorkflow(ctx context.Context, parentExecutionId *string, metadataId *string, collectionId *string, workflowId string, context map[string]string, waitForCompletion bool) (*grpc.WorkflowExecutionResponse, error) {
	executionContext, err := svc.getNewWorkflowExecutionContext(ctx, workflowId, metadataId, collectionId, context)
	if err != nil {
		return nil, err
	}

	txn, err := svc.ds.NewTransaction(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to start workflow transaction", slog.Any("error", err), slog.Any("metadataId", metadataId), slog.Any("collectionId", collectionId), slog.String("workflowId", workflowId))
		if txn != nil {
			txn.Rollback()
		}
		return nil, err
	}

	executionId, err := svc.ds.AddWorkflowExecution(ctx, txn, parentExecutionId, workflowId, metadataId, collectionId, executionContext)
	if err != nil {
		slog.ErrorContext(ctx, "failed to add workflow execution", slog.Any("error", err), slog.Any("metadataId", metadataId), slog.Any("collectionId", collectionId), slog.String("workflowId", workflowId))
		txn.Rollback()
		return nil, err
	}

	executionContext, err = svc.ds.GetWorkflowExecutionContextForUpdate(ctx, txn, executionId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow execution for update", slog.Any("error", err), slog.Any("metadataId", metadataId), slog.Any("collectionId", collectionId), slog.String("workflowId", workflowId))
		return nil, err
	}

	complete, queues, err := svc.ds.QueueNextWorkflowJobs(ctx, txn, executionContext)
	if err != nil {
		slog.ErrorContext(ctx, "failed to queue workflow execution jobs", slog.Any("error", err), slog.Any("metadataId", metadataId), slog.Any("collectionId", collectionId), slog.String("workflowId", workflowId))
		txn.Rollback()
		return nil, err
	}

	err = txn.Commit()
	if err != nil {
		slog.ErrorContext(ctx, "failed to commit workflow execution", slog.Any("error", err), slog.Any("metadataId", metadataId), slog.Any("collectionId", collectionId), slog.String("workflowId", workflowId))
		return nil, err
	}

	response := &grpc.WorkflowExecutionResponse{
		ExecutionId: executionId,
		Context:     executionContext.Context,
		Complete:    complete,
	}

	if complete {
		queues = getAllQueues(executionContext)
		err = svc.ds.NotifyExecutionCompletion(ctx, queues, executionId, true)
		if err != nil {
			slog.ErrorContext(ctx, "failed to notify of execution completion", slog.Any("error", err), slog.Any("metadataId", metadataId), slog.Any("collectionId", collectionId), slog.String("workflowId", workflowId))
			return nil, err
		}
	} else {
		err = svc.ds.NotifyJobAvailable(ctx, queues)
		if err != nil {
			slog.ErrorContext(ctx, "failed to notify of job availability", slog.Any("error", err), slog.Any("metadataId", metadataId), slog.Any("collectionId", collectionId), slog.String("workflowId", workflowId))
			return nil, err
		}
	}

	if waitForCompletion {
		notification := svc.ds.WaitForExecutionCompletion(executionId)
		if notification.Error != "" {
			response.Error = &notification.Error
			slog.ErrorContext(ctx, "failed to wait for execution completion", slog.Any("error", err), slog.Any("metadataId", metadataId), slog.Any("collectionId", collectionId), slog.String("workflowId", workflowId))
		}
		response.Success = notification.Success
		response.Complete = true
	}

	return response, err
}

func (svc *service) ExecuteWorkflow(ctx context.Context, request *grpc.WorkflowExecutionRequest) (*grpc.WorkflowExecutionResponse, error) {
	response, err := svc.executeWorkflow(ctx, request.ParentExecutionId, request.MetadataId, request.CollectionId, request.WorkflowId, request.Context, request.WaitForCompletion)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (svc *service) SetWorkflowActivityJobStatus(ctx context.Context, request *grpc.WorkflowActivityJobStatus) (*protobuf.Empty, error) {
	err := svc.processWorkflowExecution(ctx, request.ExecutionId, request)
	if err != nil {
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) processWorkflowActivityQueue(ctx context.Context, txn *sql.Tx, executionContext *grpc.WorkflowExecutionContext) (bool, bool, []string, error) {
	areChildExecutionsComplete, err := svc.ds.AreChildWorkflowExecutionsComplete(ctx, txn, executionContext.ExecutionId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to check if child workflow executions are complete", slog.Any("error", err), slog.String("executionId", executionContext.ExecutionId))
		txn.Rollback()
		return false, false, nil, err
	}

	if !areChildExecutionsComplete {
		return false, false, nil, nil
	}

	isCurrentExecutionGroupComplete, err := svc.ds.IsWorkflowExecutionGroupComplete(ctx, txn, executionContext.ExecutionId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to check if workflow execution group is complete", slog.Any("error", err), slog.String("executionId", executionContext.ExecutionId))
		txn.Rollback()
		return false, false, nil, err
	}

	var isWorkflowExecutionComplete bool
	var queues []string
	if isCurrentExecutionGroupComplete && areChildExecutionsComplete {
		isWorkflowExecutionComplete, queues, err = svc.ds.QueueNextWorkflowJobs(ctx, txn, executionContext)
		if err != nil {
			slog.ErrorContext(ctx, "failed to queue workflow jobs", slog.Any("error", err), slog.String("executionId", executionContext.ExecutionId))
			txn.Rollback()
			return false, false, nil, err
		}
	}

	return isWorkflowExecutionComplete, isCurrentExecutionGroupComplete, queues, nil
}

func (svc *service) processWorkflowExecution(ctx context.Context, executionId string, originalRequest *grpc.WorkflowActivityJobStatus) error {
	txn, err := svc.ds.NewTransaction(ctx)
	if err != nil {
		return err
	}

	executionContext, err := svc.ds.GetWorkflowExecutionContextForUpdate(ctx, txn, executionId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow execution for update", slog.Any("error", err), slog.String("executionId", executionId))
		txn.Rollback()
		return err
	}

	var success bool
	if originalRequest != nil {
		var errorMsg *string
		success, errorMsg, err = svc.ds.SetWorkflowExecutionJobStatus(ctx, txn, originalRequest.ExecutionId, originalRequest.JobId, originalRequest.Success, originalRequest.Complete, originalRequest.Error)
		if err != nil {
			slog.ErrorContext(ctx, "failed to update job status", slog.Any("error", err), slog.String("executionId", executionContext.ExecutionId))
			txn.Rollback()
			return err
		}
		if errorMsg != nil {
			slog.ErrorContext(ctx, "updated job status with provided error", slog.Any("error", err), slog.String("executionId", executionContext.ExecutionId))
			err = txn.Commit()
			if err != nil {
				return err
			}
			err := svc.ds.NotifyJobAvailable(ctx, getAllQueues(executionContext))
			if err != nil {
				slog.ErrorContext(ctx, "failed to notify jobs are complete", slog.Any("error", err), slog.String("executionId", executionContext.ExecutionId))
				return err
			}
			return nil
		}
	}

	isWorkflowExecutionComplete, isCurrentExecutionGroupComplete, queues, err := svc.processWorkflowActivityQueue(ctx, txn, executionContext)

	err = txn.Commit()
	if err != nil {
		return err
	}

	if isWorkflowExecutionComplete {
		// TODO: probably notifying too much
		if queues == nil {
			queues = getAllQueues(executionContext)
		}
		err := svc.ds.NotifyExecutionCompletion(ctx, queues, executionContext.ExecutionId, true)
		if err != nil {
			slog.ErrorContext(ctx, "failed to notify execution completion", slog.Any("error", err), slog.String("executionId", executionContext.ExecutionId))
			return err
		}
		if executionContext.ParentExecutionId != nil {
			return svc.processWorkflowExecution(ctx, *executionContext.ParentExecutionId, nil)
		}
	} else if isCurrentExecutionGroupComplete {
		err := svc.ds.NotifyJobAvailable(ctx, queues)
		if err != nil {
			slog.ErrorContext(ctx, "failed to notify jobs are complete", slog.Any("error", err), slog.String("executionId", executionContext.ExecutionId))
			return err
		}
	}

	if !success {
		// TODO: probably notifying too much
		if queues == nil {
			queues = getAllQueues(executionContext)
		}
		err = svc.ds.NotifyExecutionFailed(ctx, queues, executionContext.ExecutionId, err)
		if err != nil {
			return err
		}
	}

	return nil
}
