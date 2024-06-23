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

package content

import (
	protobuf "bosca.io/api/protobuf/bosca"
	grpc "bosca.io/api/protobuf/bosca/content"
	"context"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (svc *service) GetTraits(ctx context.Context, request *protobuf.Empty) (*grpc.Traits, error) {
	traits, err := svc.ds.GetTraits(ctx)
	if err != nil {
		return nil, err
	}
	return &grpc.Traits{
		Traits: traits,
	}, err
}

func (svc *service) GetWorkflow(ctx context.Context, request *protobuf.IdRequest) (*grpc.Workflow, error) {
	workflow, err := svc.ds.GetWorkflow(ctx, request.Id)
	return workflow, err
}

func (svc *service) GetWorkflows(ctx context.Context, request *protobuf.Empty) (*grpc.Workflows, error) {
	workflows, err := svc.ds.GetWorkflows(ctx)
	if err != nil {
		return nil, err
	}
	return &grpc.Workflows{
		Workflows: workflows,
	}, err
}

func (svc *service) GetWorkflowActivityInstances(ctx context.Context, request *protobuf.IdRequest) (*grpc.WorkflowActivityInstances, error) {
	instances, err := svc.ds.GetWorkflowActivityInstances(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &grpc.WorkflowActivityInstances{
		Instances: instances,
	}, err
}

func (svc *service) executeEnterExitWorkflow(ctx context.Context, metadataId string, workflowId string, exit bool) error {
	workflow, err := svc.ds.GetWorkflow(ctx, workflowId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow", slog.String("id", metadataId), slog.String("workflowId", workflowId))
		return err
	}
	if workflow == nil {
		slog.ErrorContext(ctx, "workflow not found", slog.String("id", metadataId), slog.String("workflowId", workflowId))
		return status.Error(codes.Internal, "workflow not found")
	}
	var stateMemo = "enter-test"
	if exit {
		stateMemo = "exit-test"
	}
	run, err := svc.temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		TaskQueue: workflow.Queue,
		Memo: map[string]interface{}{
			"state": stateMemo,
		},
	}, workflow.Name, metadataId)
	if err != nil {
		return err
	}
	var result bool
	err = run.Get(ctx, &result)
	if err != nil {
		return err
	}
	if !result {
		return status.Error(codes.InvalidArgument, "transition validation failed")
	}
	return nil
}

func (svc *service) executeWorkflow(ctx context.Context, metadata *grpc.Metadata, workflowId string) error {
	workflow, err := svc.ds.GetWorkflow(ctx, workflowId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow", slog.String("id", metadata.Id), slog.String("workflowId", workflowId))
		return err
	}
	if workflow == nil {
		slog.ErrorContext(ctx, "workflow not found", slog.String("id", metadata.Id), slog.String("workflowId", workflowId))
		return status.Error(codes.Internal, "workflow not found")
	}
	executionContext := &grpc.WorkflowActivityExecutionContext{
		WorkflowId: workflowId,
		Metadata:   metadata,
	}
	instances, err := svc.ds.GetWorkflowActivityInstances(ctx, workflowId)
	if err != nil {
		return err
	}
	executionContext.Activities = instances
	_, err = svc.temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		TaskQueue: workflow.Queue,
	}, workflow.Id, executionContext)
	if err != nil {
		return err
	}
	return nil
}

func (svc *service) verifyTransitionExists(ctx context.Context, metadata *grpc.Metadata, nextStateId string) error {
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

func (svc *service) verifyEnterTransition(ctx context.Context, metadata *grpc.Metadata, nextStateId string) (*grpc.WorkflowState, error) {
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

func (svc *service) verifyExitTransition(ctx context.Context, metadata *grpc.Metadata, nextStateId string) error {
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

func (svc *service) transition(ctx context.Context, metadata *grpc.Metadata, nextState *grpc.WorkflowState, status string) error {
	var err error
	if nextState.WorkflowId != nil {
		err = svc.ds.TransitionMetadataWorkflowStateId(ctx, metadata, nextState, status, true, false)
		if err != nil {
			return err
		}
		err = svc.executeWorkflow(ctx, metadata, *nextState.WorkflowId)
		if err != nil {
			_ = svc.ds.TransitionMetadataWorkflowStateId(ctx, metadata, nextState, status, false, true)
			return err
		}
	} else {
		err = svc.ds.TransitionMetadataWorkflowStateId(ctx, metadata, nextState, status, true, true)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) BeginTransitionWorkflow(ctx context.Context, request *grpc.TransitionWorkflowRequest) (*protobuf.Empty, error) {
	md, err := svc.ds.GetMetadata(ctx, request.MetadataId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow", slog.String("id", request.MetadataId), slog.Any("error", err))
		return nil, err
	}

	if md == nil {
		slog.ErrorContext(ctx, "workflow not found", slog.String("id", request.MetadataId), slog.Any("error", err))
		return nil, status.Error(codes.NotFound, "workflow not found")
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
	md, err := svc.ds.GetMetadata(ctx, request.MetadataId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow", slog.String("id", request.MetadataId), slog.Any("error", err))
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

	state, err := svc.ds.GetWorkflowState(ctx, *md.WorkflowStatePendingId)
	if err != nil {
		slog.ErrorContext(ctx, "error getting state", slog.String("id", request.MetadataId), slog.String("stateId", *md.WorkflowStatePendingId), slog.Any("error", err))
		return nil, err
	}

	err = svc.ds.TransitionMetadataWorkflowStateId(ctx, md, state, request.Status, request.Success, true)
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func (svc *service) GetWorkflowStates(ctx context.Context, request *protobuf.Empty) (*grpc.WorkflowStates, error) {
	states, err := svc.ds.GetWorkflowStates(ctx)
	if err != nil {
		return nil, err
	}
	return &grpc.WorkflowStates{
		States: states,
	}, err
}
