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
	grpcContent "bosca.io/api/protobuf/bosca/content"
	grpc "bosca.io/api/protobuf/bosca/workflow"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"time"
)

func (svc *service) GetWorkflow(ctx context.Context, request *protobuf.IdRequest) (*grpc.Workflow, error) {
	workflow, err := svc.ds.GetWorkflow(ctx, request.Id)
	return workflow, err
}

func (svc *service) GetWorkflows(ctx context.Context, _ *protobuf.Empty) (*grpc.Workflows, error) {
	workflows, err := svc.ds.GetWorkflows(ctx)
	if err != nil {
		return nil, err
	}
	return &grpc.Workflows{
		Workflows: workflows,
	}, err
}

func (svc *service) GetWorkflowActivity(ctx context.Context, request *protobuf.IdRequest) (*grpc.WorkflowActivities, error) {
	activities, err := svc.ds.GetWorkflowActivities(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &grpc.WorkflowActivities{
		Activities: activities,
	}, err
}

func (svc *service) getNewWorkflowExecutionContext(ctx context.Context, workflowId string, metadataId string) (*grpc.WorkflowExecutionContext, error) {
	metadata, err := svc.getMetadata(ctx, metadataId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get metadata", slog.String("id", metadataId), slog.String("workflowId", workflowId))
		return nil, err
	}
	workflow, err := svc.ds.GetWorkflow(ctx, workflowId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow", slog.String("id", metadata.Id), slog.String("workflowId", workflowId))
		return nil, err
	}
	if workflow == nil {
		slog.ErrorContext(ctx, "workflow not found", slog.String("id", metadata.Id), slog.String("workflowId", workflowId))
		return nil, status.Error(codes.Internal, "workflow not found")
	}
	jobs, err := svc.ds.GetWorkflowActivityJobs(ctx, workflowId)
	if err != nil {
		return nil, err
	}
	return &grpc.WorkflowExecutionContext{
		MetadataId:            metadataId,
		WorkflowId:            workflowId,
		Jobs:                  jobs,
		CurrentExecutionGroup: -1,
	}, nil
}

func (svc *service) executeEnterExitWorkflow(ctx context.Context, metadataId string, workflowId string, exit bool) error {
	return svc.executeWorkflow(ctx, metadataId, workflowId)
}

func (svc *service) executeWorkflow(ctx context.Context, metadataId string, workflowId string) error {
	executionContext, err := svc.getNewWorkflowExecutionContext(ctx, workflowId, metadataId)
	if err != nil {
		return err
	}
	txn, err := svc.ds.NewTransaction(ctx)
	if err != nil {
		if txn != nil {
			txn.Rollback()
		}
		return err
	}
	executionId, err := svc.ds.AddWorkflowExecution(ctx, txn, workflowId, metadataId, executionContext)
	if err != nil {
		txn.Rollback()
		return err
	}
	err = svc.ds.ExecuteWorkflow(ctx, txn, executionId)
	if err != nil {
		txn.Rollback()
		return err
	}
	return txn.Commit()
}

func (svc *service) verifyTransitionExists(ctx context.Context, metadata *grpcContent.Metadata, nextStateId string) error {
	transition, err := svc.ds.GetWorkflowTransition(ctx, metadata.WorkflowStateId, nextStateId)
	if err != nil {
		slog.WarnContext(ctx, "getting transition failed", slog.String("id", metadata.Id), slog.String("currentState", metadata.WorkflowStateId), slog.String("newState", nextStateId), slog.Any("error", err))
		return err
	}
	if transition == nil {
		slog.ErrorContext(ctx, "transition not found", slog.String("id", metadata.Id), slog.String("currentState", metadata.WorkflowStateId), slog.String("newState", nextStateId))
		return status.Error(codes.InvalidArgument, "invalid transition")
	}
	return nil
}

func (svc *service) verifyEnterTransition(ctx context.Context, metadata *grpcContent.Metadata, nextStateId string) (*grpc.WorkflowState, error) {
	nextState, err := svc.ds.GetWorkflowState(ctx, nextStateId)
	if err != nil {
		slog.ErrorContext(ctx, "getting next state transition failed", slog.String("id", metadata.Id), slog.String("currentState", metadata.WorkflowStateId), slog.String("newState", nextStateId), slog.Any("error", err))
		return nil, err
	}
	if nextState.EntryWorkflowId != nil {
		err = svc.executeEnterExitWorkflow(ctx, metadata.Id, *nextState.EntryWorkflowId, false)
		if err != nil {
			slog.ErrorContext(ctx, "next state transition failed", slog.String("id", metadata.Id), slog.String("currentState", metadata.WorkflowStateId), slog.String("newState", nextStateId), slog.Any("error", err))
			return nil, err
		}
	}
	return nextState, nil
}

func (svc *service) verifyExitTransition(ctx context.Context, metadata *grpcContent.Metadata, nextStateId string) error {
	currentState, err := svc.ds.GetWorkflowState(ctx, metadata.WorkflowStateId)
	if err != nil {
		slog.ErrorContext(ctx, "getting current state transition failed", slog.String("id", metadata.Id), slog.String("currentState", metadata.WorkflowStateId), slog.String("newState", nextStateId), slog.Any("error", err))
		return err
	}
	if currentState.ExitWorkflowId != nil {
		err = svc.executeEnterExitWorkflow(ctx, metadata.Id, *currentState.ExitWorkflowId, true)
		if err != nil {
			slog.ErrorContext(ctx, "current state transition failed", slog.String("id", metadata.Id), slog.String("currentState", metadata.WorkflowStateId), slog.String("newState", nextStateId), slog.Any("error", err))
			return err
		}
	}
	return nil
}

func (svc *service) transition(ctx context.Context, metadata *grpcContent.Metadata, nextState *grpc.WorkflowState, status string) error {
	var err error
	if nextState.WorkflowId != nil {
		err = svc.setBeginMetadataWorkflowState(ctx, metadata, nextState, status)
		if err != nil {
			return err
		}
		err = svc.executeWorkflow(ctx, metadata.Id, *nextState.WorkflowId)
		if err != nil {
			_ = svc.setCompleteMetadataWorkflowState(ctx, metadata, status, false)
			return err
		}
	} else {
		err = svc.setCompleteMetadataWorkflowState(ctx, metadata, status, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) BeginTransitionWorkflow(ctx context.Context, request *grpc.BeginTransitionWorkflowRequest) (*protobuf.Empty, error) {
	var md *grpcContent.Metadata
	var err error
	for tries := 0; tries < 10; tries++ {
		md, err = svc.getMetadata(ctx, request.MetadataId)
		if err != nil {
			if err.Error() == "rpc error: code = Unauthenticated desc = permission check failed" {
				time.Sleep(3 * time.Second)
			} else {
				slog.ErrorContext(ctx, "failed to begin transition workflow get metadata", slog.String("id", request.MetadataId), slog.Any("error", err))
				return nil, err
			}
		} else if md != nil {
			break
		}
	}

	if md == nil {
		slog.ErrorContext(ctx, "metadata not found", slog.String("id", request.MetadataId), slog.Any("error", err))
		return nil, status.Error(codes.NotFound, "metadata not found")
	}

	if md.WorkflowStateId == request.StateId {
		slog.WarnContext(ctx, "workflow already in state", slog.String("id", request.MetadataId), slog.String("state", request.StateId))
		return &protobuf.Empty{}, nil
	}

	if err = svc.verifyTransitionExists(ctx, md, request.StateId); err != nil {
		return nil, err
	}

	if err = svc.verifyExitTransition(ctx, md, request.StateId); err != nil {
		return nil, err
	}

	if nextState, err := svc.verifyEnterTransition(ctx, md, request.StateId); err != nil {
		return nil, err
	} else if err = svc.transition(ctx, md, nextState, request.StateId); err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func (svc *service) CompleteTransitionWorkflow(ctx context.Context, request *grpc.CompleteTransitionWorkflowRequest) (*protobuf.Empty, error) {
	md, err := svc.getMetadata(ctx, request.MetadataId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to complete transition workflow get metadata failed", slog.String("id", request.MetadataId), slog.Any("error", err))
		return nil, err
	}

	if md == nil {
		slog.ErrorContext(ctx, "workflow not found", slog.String("id", request.MetadataId), slog.Any("error", err))
		return nil, status.Error(codes.NotFound, "workflow not found")
	}

	if md.WorkflowStatePendingId == nil {
		slog.ErrorContext(ctx, "workflow no pending state", slog.String("id", request.MetadataId), slog.Any("error", err))
		return nil, status.Error(codes.FailedPrecondition, "workflow no pending state")
	}

	err = svc.setCompleteMetadataWorkflowState(ctx, md, request.Status, request.Success)
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func (svc *service) GetWorkflowStates(ctx context.Context, _ *protobuf.Empty) (*grpc.WorkflowStates, error) {
	states, err := svc.ds.GetWorkflowStates(ctx)
	if err != nil {
		return nil, err
	}
	return &grpc.WorkflowStates{
		States: states,
	}, err
}
