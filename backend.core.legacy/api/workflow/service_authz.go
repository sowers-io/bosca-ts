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
	"bosca.io/pkg/security"
	"context"
)

// This service authorizes all calls to ensure the principal is authorized to perform the request.
//
//	If the principal is authorized, it forwards the request to the underlying service interface.
type authorizationService struct {
	grpc.UnimplementedWorkflowServiceServer

	service     grpc.WorkflowServiceServer
	permissions security.PermissionManager
}

func NewAuthorizationService(permissions security.PermissionManager, dataSource *DataStore, service grpc.WorkflowServiceServer) grpc.WorkflowServiceServer {
	svc := &authorizationService{
		service:     service,
		permissions: permissions,
	}
	return svc
}

func (svc *authorizationService) GetPrompts(ctx context.Context, request *protobuf.Empty) (*grpc.Prompts, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetPrompts(ctx, request)
}

func (svc *authorizationService) GetPrompt(ctx context.Context, request *protobuf.IdRequest) (*grpc.Prompt, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetPrompt(ctx, request)
}

func (svc *authorizationService) GetModels(ctx context.Context, request *protobuf.Empty) (*grpc.Models, error) {
	// TODO: permission type?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetModels(ctx, request)
}

func (svc *authorizationService) GetModel(ctx context.Context, request *protobuf.IdRequest) (*grpc.Model, error) {
	// TODO: permission type?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetModel(ctx, request)
}

func (svc *authorizationService) GetStorageSystems(ctx context.Context, request *protobuf.Empty) (*grpc.StorageSystems, error) {
	// TODO: permission type?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetStorageSystems(ctx, request)
}

func (svc *authorizationService) GetStorageSystem(ctx context.Context, request *protobuf.IdRequest) (*grpc.StorageSystem, error) {
	// TODO: permission type?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetStorageSystem(ctx, request)
}

func (svc *authorizationService) GetStorageSystemModels(ctx context.Context, request *protobuf.IdRequest) (*grpc.StorageSystemModels, error) {
	// TODO: permission type?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetStorageSystemModels(ctx, request)
}

func (svc *authorizationService) GetWorkflows(ctx context.Context, request *protobuf.Empty) (*grpc.Workflows, error) {
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflows(ctx, request)
}

func (svc *authorizationService) GetWorkflow(ctx context.Context, request *protobuf.IdRequest) (*grpc.Workflow, error) {
	// TODO
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflow(ctx, request)
}

func (svc *authorizationService) GetWorkflowActivities(ctx context.Context, request *protobuf.IdRequest) (*grpc.WorkflowActivities, error) {
	// TODO
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflowActivities(ctx, request)
}

func (svc *authorizationService) GetWorkflowActivityStorageSystems(ctx context.Context, request *grpc.WorkflowActivityIdRequest) (*grpc.WorkflowActivityStorageSystems, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflowActivityStorageSystems(ctx, request)
}

func (svc *authorizationService) GetWorkflowActivityPrompts(ctx context.Context, request *grpc.WorkflowActivityIdRequest) (*grpc.WorkflowActivityPrompts, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflowActivityPrompts(ctx, request)
}

func (svc *authorizationService) GetWorkflowStates(ctx context.Context, request *protobuf.Empty) (*grpc.WorkflowStates, error) {
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_state_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflowStates(ctx, request)
}

func (svc *authorizationService) BeginTransitionWorkflow(ctx context.Context, request *grpc.BeginTransitionWorkflowRequest) (*protobuf.Empty, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_state_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.BeginTransitionWorkflow(ctx, request)
}

func (svc *authorizationService) CompleteTransitionWorkflow(ctx context.Context, request *grpc.CompleteTransitionWorkflowRequest) (*protobuf.Empty, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_state_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.CompleteTransitionWorkflow(ctx, request)
}

func (svc *authorizationService) ExecuteWorkflow(ctx context.Context, request *grpc.WorkflowExecutionRequest) (*grpc.WorkflowEnqueueResponse, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_state_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.ExecuteWorkflow(ctx, request)
}

func (svc *authorizationService) FindAndExecuteWorkflow(ctx context.Context, request *grpc.FindAndWorkflowExecutionRequest) (*grpc.WorkflowEnqueueResponses, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpcContent.PermissionObjectType_workflow_state_type, "all", grpcContent.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.FindAndExecuteWorkflow(ctx, request)
}
