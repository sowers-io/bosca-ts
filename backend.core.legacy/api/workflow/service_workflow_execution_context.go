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

func (svc *service) newWorkflowEnqueueRequest(ctx context.Context, parent *grpc.WorkflowParentJobId, workflowId string, metadataId *string, collectionId *string, context map[string]string, waitForCompletion bool) (*grpc.WorkflowEnqueueRequest, error) {
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
	jobs := make([]*grpc.WorkflowJob, len(activities))
	for i, activity := range activities {
		prompts, err := svc.ds.GetWorkflowActivityPrompts(ctx, activity.WorkflowActivityId)
		if err != nil {
			return nil, err
		}
		storageSystems, err := svc.ds.GetWorkflowActivityStorageSystems(ctx, activity.WorkflowActivityId)
		if err != nil {
			return nil, err
		}
		models, err := svc.ds.GetWorkflowActivityModels(ctx, activity.WorkflowActivityId)
		if err != nil {
			return nil, err
		}
		jobs[i] = &grpc.WorkflowJob{
			WorkflowId:     workflow.Id,
			MetadataId:     metadataId,
			Activity:       activity,
			Prompts:        prompts,
			Models:         models,
			StorageSystems: storageSystems,
			Context:        make(map[string]string),
		}
	}
	return &grpc.WorkflowEnqueueRequest{
		Parent:            parent,
		MetadataId:        metadataId,
		CollectionId:      collectionId,
		Workflow:          workflow,
		Jobs:              jobs,
		Context:           context,
		WaitForCompletion: waitForCompletion,
	}, nil
}
