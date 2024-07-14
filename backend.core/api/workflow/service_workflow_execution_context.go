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
	grpc "bosca.io/api/protobuf/bosca/workflow"
	"context"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

func (svc *service) getNewWorkflowExecutionContext(ctx context.Context, workflowId string, metadataId *string, collectionId *string, context map[string]string) (*grpc.WorkflowExecutionContext, error) {
	workflow, err := svc.ds.GetWorkflow(ctx, workflowId)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow", slog.Any("metadataId", metadataId), slog.Any("collectionId", collectionId), slog.String("workflowId", workflowId))
		return nil, err
	}
	if workflow == nil {
		slog.ErrorContext(ctx, "workflow not found", slog.Any("metadataId", metadataId), slog.Any("collectionId", collectionId), slog.String("workflowId", workflowId))
		return nil, status.Error(codes.Internal, "workflow not found")
	}
	activities, err := svc.ds.GetWorkflowActivities(ctx, workflowId)
	if err != nil {
		return nil, err
	}
	return &grpc.WorkflowExecutionContext{
		MetadataId:            metadataId,
		CollectionId:          collectionId,
		WorkflowId:            workflowId,
		Activities:            activities,
		Context:               context,
		CurrentExecutionGroup: -1,
	}, nil
}
