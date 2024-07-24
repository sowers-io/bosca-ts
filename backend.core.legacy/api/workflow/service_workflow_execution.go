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
	"bosca.io/api/protobuf/bosca/content"
	grpc "bosca.io/api/protobuf/bosca/workflow"
	"bosca.io/pkg/security/identity"
	"context"
)

func (svc *service) executeWorkflow(ctx context.Context, parent *grpc.WorkflowParentJobId, metadataId *string, collectionId *string, workflowId string, context map[string]string, waitForCompletion bool) (*grpc.WorkflowEnqueueResponse, error) {
	request, err := svc.newWorkflowEnqueueRequest(ctx, parent, workflowId, metadataId, collectionId, context, waitForCompletion)
	if err != nil {
		return nil, err
	}
	return svc.queueClient.Enqueue(ctx, request)
}

func (svc *service) ExecuteWorkflow(ctx context.Context, request *grpc.WorkflowExecutionRequest) (*grpc.WorkflowEnqueueResponse, error) {
	response, err := svc.executeWorkflow(ctx, request.Parent, request.MetadataId, request.CollectionId, request.WorkflowId, request.Context, request.WaitForCompletion)
	if err != nil {
		return nil, err
	}
	return response, nil
}

func (svc *service) FindAndExecuteWorkflow(ctx context.Context, request *grpc.FindAndWorkflowExecutionRequest) (*grpc.WorkflowEnqueueResponses, error) {
	responses := &grpc.WorkflowEnqueueResponses{
		Responses: make([]*grpc.WorkflowEnqueueResponse, 0),
	}
	if request.MetadataAttributes != nil {
		md, err := svc.contentClient.FindMetadata(identity.GetAuthenticatedContext(ctx), &content.FindMetadataRequest{
			Attributes: request.MetadataAttributes,
		})
		if err != nil {
			return nil, err
		}
		for _, m := range md.Metadata {
			response, err := svc.executeWorkflow(ctx, request.Parent, &m.Id, nil, request.WorkflowId, request.Context, request.WaitForCompletion)
			if err != nil {
				return nil, err
			}
			responses.Responses = append(responses.Responses, response)
		}
	}
	if request.CollectionAttributes != nil {
		md, err := svc.contentClient.FindCollection(identity.GetAuthenticatedContext(ctx), &content.FindCollectionRequest{
			Attributes: request.CollectionAttributes,
		})
		if err != nil {
			return nil, err
		}
		for _, m := range md.Collections {
			response, err := svc.executeWorkflow(ctx, request.Parent, nil, &m.Id, request.WorkflowId, request.Context, request.WaitForCompletion)
			if err != nil {
				return nil, err
			}
			responses.Responses = append(responses.Responses, response)
		}
	}
	return responses, nil
}
