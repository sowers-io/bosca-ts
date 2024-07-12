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
	"bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/content"
	"context"
)

func (svc *service) SetWorkflowState(ctx context.Context, request *content.SetWorkflowStateRequest) (*bosca.Empty, error) {
	metadata, err := svc.ds.GetMetadata(ctx, request.MetadataId)
	if err != nil {
		return nil, err
	}
	err = svc.ds.SetMetadataWorkflowState(ctx, metadata, request.StateId, request.Status, true, request.Immediate)
	if err != nil {
		return nil, err
	}
	return &bosca.Empty{}, nil
}

func (svc *service) SetWorkflowStateComplete(ctx context.Context, request *content.SetWorkflowStateCompleteRequest) (*bosca.Empty, error) {
	metadata, err := svc.ds.GetMetadata(ctx, request.MetadataId)
	if err != nil {
		return nil, err
	}
	state := metadata.WorkflowStateId
	if metadata.WorkflowStatePendingId != nil {
		state = *metadata.WorkflowStatePendingId
	}
	err = svc.ds.SetMetadataWorkflowState(ctx, metadata, state, request.Status, true, true)
	if err != nil {
		return nil, err
	}
	return &bosca.Empty{}, nil
}
