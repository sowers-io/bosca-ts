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
)

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

func (svc *service) verifyEnterTransitionExecution(ctx context.Context, metadata *grpcContent.Metadata, nextStateId string) (*grpc.WorkflowState, error) {
	nextState, err := svc.ds.GetWorkflowState(ctx, nextStateId)
	if err != nil {
		slog.ErrorContext(ctx, "getting next state transition failed", slog.String("id", metadata.Id), slog.String("currentState", metadata.WorkflowStateId), slog.String("newState", nextStateId), slog.Any("error", err))
		return nil, err
	}
	if nextState.EntryWorkflowId != nil {
		_, err = svc.executeWorkflow(ctx, nil, metadata.Id, *nextState.EntryWorkflowId, nil, true)
		if err != nil {
			slog.ErrorContext(ctx, "next state transition failed", slog.String("id", metadata.Id), slog.String("currentState", metadata.WorkflowStateId), slog.String("newState", nextStateId), slog.Any("error", err))
			return nil, err
		}
	}
	return nextState, nil
}

func (svc *service) verifyExitTransitionExecution(ctx context.Context, metadata *grpcContent.Metadata, nextStateId string) error {
	currentState, err := svc.ds.GetWorkflowState(ctx, metadata.WorkflowStateId)
	if err != nil {
		slog.ErrorContext(ctx, "getting current state transition failed", slog.String("id", metadata.Id), slog.String("currentState", metadata.WorkflowStateId), slog.String("newState", nextStateId), slog.Any("error", err))
		return err
	}
	if currentState.ExitWorkflowId != nil {
		_, err = svc.executeWorkflow(ctx, nil, metadata.Id, *currentState.ExitWorkflowId, nil, true)
		if err != nil {
			slog.ErrorContext(ctx, "current state transition failed", slog.String("id", metadata.Id), slog.String("currentState", metadata.WorkflowStateId), slog.String("newState", nextStateId), slog.Any("error", err))
			return err
		}
	}
	return nil
}

func (svc *service) transition(ctx context.Context, metadata *grpcContent.Metadata, nextState *grpc.WorkflowState, status string, waitForCompletion bool) error {
	var err error
	if nextState.WorkflowId != nil {
		err = svc.setBeginMetadataWorkflowState(ctx, metadata, nextState, status)
		if err != nil {
			return err
		}
		response, err := svc.executeWorkflow(ctx, nil, metadata.Id, *nextState.WorkflowId, nil, waitForCompletion)
		if err != nil || response.Error != nil {
			return err
		}
	} else {
		err = svc.setCompleteMetadataWorkflowState(ctx, metadata, status)
		if err != nil {
			return err
		}
	}
	return nil
}

func (svc *service) BeginTransitionWorkflow(ctx context.Context, request *grpc.BeginTransitionWorkflowRequest) (*protobuf.Empty, error) {
	metadata, err := svc.getMetadata(ctx, request.MetadataId)
	if err != nil {
		slog.ErrorContext(ctx, "couldn't begin transition, metadata error", slog.String("id", request.MetadataId), slog.Any("error", err))
		return nil, err
	}
	if metadata == nil {
		slog.ErrorContext(ctx, "couldn't begin transition, missing metadata", slog.String("id", request.MetadataId), slog.Any("error", err))
		return nil, status.Error(codes.NotFound, "metadata not found")
	}

	if metadata.WorkflowStateId == request.StateId {
		slog.WarnContext(ctx, "workflow already in state", slog.String("id", request.MetadataId), slog.String("state", request.StateId))
		_ = svc.setCompleteMetadataWorkflowState(ctx, metadata, "duplicate transition")
		return &protobuf.Empty{}, nil
	}

	if err = svc.verifyTransitionExists(ctx, metadata, request.StateId); err != nil {
		return nil, err
	}

	if err = svc.verifyExitTransitionExecution(ctx, metadata, request.StateId); err != nil {
		return nil, err
	}

	if nextState, err := svc.verifyEnterTransitionExecution(ctx, metadata, request.StateId); err != nil {
		return nil, err
	} else if err = svc.transition(ctx, metadata, nextState, request.StateId, false); err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}

func (svc *service) CompleteTransitionWorkflow(ctx context.Context, request *grpc.CompleteTransitionWorkflowRequest) (*protobuf.Empty, error) {
	metadata, err := svc.getMetadata(ctx, request.MetadataId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to complete transition workflow get metadata failed", slog.String("id", request.MetadataId), slog.Any("error", err))
		return nil, err
	}

	if metadata == nil {
		slog.ErrorContext(ctx, "workflow not found", slog.String("id", request.MetadataId), slog.Any("error", err))
		return nil, status.Error(codes.NotFound, "workflow not found")
	}

	if metadata.WorkflowStatePendingId == nil {
		slog.ErrorContext(ctx, "workflow no pending state", slog.String("id", request.MetadataId), slog.Any("error", err))
		return nil, status.Error(codes.FailedPrecondition, "workflow no pending state")
	}

	err = svc.setCompleteMetadataWorkflowState(ctx, metadata, request.Status)
	if err != nil {
		return nil, err
	}

	return &protobuf.Empty{}, nil
}
