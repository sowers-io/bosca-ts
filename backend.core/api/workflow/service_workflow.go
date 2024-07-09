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

func (svc *service) GetWorkflowStates(ctx context.Context, _ *protobuf.Empty) (*grpc.WorkflowStates, error) {
	states, err := svc.ds.GetWorkflowStates(ctx)
	if err != nil {
		return nil, err
	}
	return &grpc.WorkflowStates{
		States: states,
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

