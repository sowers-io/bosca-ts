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
	"bosca.io/pkg/security"
	"context"
	"errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strings"
)

// This service authorizes all calls to ensure the principal is authorized to perform the request.
//
//	If the principal is authorized, it forwards the request to the underlying service interface.
type authorizationService struct {
	grpc.UnimplementedContentServiceServer

	service     grpc.ContentServiceServer
	permissions security.PermissionManager
}

func NewAuthorizationService(permissions security.PermissionManager, dataSource *DataStore, service grpc.ContentServiceServer) grpc.ContentServiceServer {
	svc := &authorizationService{
		service:     service,
		permissions: permissions,
	}
	return svc
}

func (svc *authorizationService) GetSources(ctx context.Context, request *protobuf.Empty) (*grpc.Sources, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetSources(ctx, request)
}

func (svc *authorizationService) GetSource(ctx context.Context, request *protobuf.IdRequest) (*grpc.Source, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetSource(ctx, request)
}

func (svc *authorizationService) GetPrompts(ctx context.Context, request *protobuf.Empty) (*grpc.Prompts, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetPrompts(ctx, request)
}

func (svc *authorizationService) GetPrompt(ctx context.Context, request *protobuf.IdRequest) (*grpc.Prompt, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetPrompt(ctx, request)
}

func (svc *authorizationService) CheckPermission(ctx context.Context, request *grpc.PermissionCheckRequest) (*grpc.PermissionCheckResponse, error) {
	err := svc.permissions.CheckWithSubjectIdError(ctx, request.SubjectType, request.Subject, request.ObjectType, request.Object, request.Action)
	if err != nil {
		return &grpc.PermissionCheckResponse{Allowed: false}, nil
	}
	return &grpc.PermissionCheckResponse{Allowed: true}, nil
}

func (svc *authorizationService) GetRootCollectionItems(ctx context.Context, request *protobuf.Empty) (*grpc.CollectionItems, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, RootCollectionId, grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetRootCollectionItems(ctx, request)
}

func (svc *authorizationService) GetCollection(ctx context.Context, request *protobuf.IdRequest) (*grpc.Collection, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, request.Id, grpc.PermissionAction_view)
	if err != nil {
		return nil, err
	}
	return svc.service.GetCollection(ctx, request)
}

func (svc *authorizationService) GetCollectionItems(ctx context.Context, request *protobuf.IdRequest) (*grpc.CollectionItems, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, request.Id, grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetCollectionItems(ctx, request)
}

func (svc *authorizationService) AddCollection(ctx context.Context, request *grpc.AddCollectionRequest) (*protobuf.IdResponse, error) {
	if request.Collection == nil {
		return nil, errors.New("collection is required")
	}
	parentId := request.Parent
	if len(strings.Trim(request.Parent, " ")) == 0 {
		if request.Collection.Type == grpc.CollectionType_folder {
			request.Parent = RootCollectionId
		}
		parentId = RootCollectionId
	}
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, parentId, grpc.PermissionAction_edit)
	if err != nil {
		return nil, err
	}
	return svc.service.AddCollection(ctx, request)
}

func (svc *authorizationService) AddCollections(ctx context.Context, request *grpc.AddCollectionsRequest) (*protobuf.IdResponses, error) {
	parentIds := make([]string, 0)
	for _, collection := range request.Collections {
		if collection.Collection == nil {
			return nil, errors.New("collection is required")
		}
		if len(strings.Trim(collection.Parent, " ")) == 0 {
			collection.Parent = RootCollectionId
		}
		parentIds = append(parentIds, collection.Parent)
	}
	ids, err := svc.permissions.BulkCheck(ctx, grpc.PermissionObjectType_collection_type, parentIds, grpc.PermissionAction_edit)
	if err != nil {
		return nil, err
	}
	if len(ids) != len(parentIds) {
		return nil, status.Errorf(codes.Unauthenticated, "permission check failed")
	}
	return svc.service.AddCollections(ctx, request)
}

func (svc *authorizationService) AddCollectionItem(ctx context.Context, request *grpc.AddCollectionItemRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, request.CollectionId, grpc.PermissionAction_edit)
	if err != nil {
		return nil, err
	}
	return svc.service.AddCollectionItem(ctx, request)
}

func (svc *authorizationService) DeleteCollection(ctx context.Context, request *protobuf.IdRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, request.Id, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.DeleteCollection(ctx, request)
}

func (svc *authorizationService) GetMetadata(ctx context.Context, request *protobuf.IdRequest) (*grpc.Metadata, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_view)
	if err != nil {
		return nil, err
	}
	return svc.service.GetMetadata(ctx, request)
}

func (svc *authorizationService) GetMetadatas(ctx context.Context, request *protobuf.IdsRequest) (*grpc.Metadatas, error) {
	return svc.service.GetMetadatas(ctx, request)
}

func (svc *authorizationService) GetMetadataUploadUrl(ctx context.Context, request *protobuf.IdRequest) (*grpc.SignedUrl, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_edit)
	if err != nil {
		err = svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_edit)
		if err != nil {
			return nil, err
		}
	}
	return svc.service.GetMetadataUploadUrl(ctx, request)
}

func (svc *authorizationService) GetMetadataDownloadUrl(ctx context.Context, request *protobuf.IdRequest) (*grpc.SignedUrl, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_view)
	if err != nil {
		return nil, err
	}
	return svc.service.GetMetadataDownloadUrl(ctx, request)
}

func (svc *authorizationService) AddMetadataRelationship(ctx context.Context, request *grpc.MetadataRelationship) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.MetadataId1, grpc.PermissionAction_edit)
	if err != nil {
		return nil, err
	}
	err = svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.MetadataId2, grpc.PermissionAction_edit)
	if err != nil {
		return nil, err
	}
	return svc.service.AddMetadataRelationship(ctx, request)
}

func (svc *authorizationService) AddMetadata(ctx context.Context, request *grpc.AddMetadataRequest) (*protobuf.IdResponse, error) {
	if request.Metadata == nil {
		return nil, errors.New("workflow is required")
	}
	collection := RootCollectionId
	if request.Collection != nil {
		collection = *request.Collection
	}
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, collection, grpc.PermissionAction_edit)
	if err != nil {
		return nil, err
	}
	return svc.service.AddMetadata(ctx, request)
}

func (svc *authorizationService) AddMetadatas(ctx context.Context, request *grpc.AddMetadatasRequest) (*protobuf.IdResponses, error) {
	collections := make([]string, len(request.Metadatas))
	for i, metadata := range request.Metadatas {
		if metadata.Metadata == nil {
			return nil, errors.New("workflow is required")
		}
		collection := RootCollectionId
		if metadata.Collection != nil {
			collection = *metadata.Collection
		}
		collections[i] = collection
	}
	ids, err := svc.permissions.BulkCheck(ctx, grpc.PermissionObjectType_collection_type, collections, grpc.PermissionAction_edit)
	if err != nil {
		return nil, err
	}
	if len(ids) != len(request.Metadatas) {
		return nil, status.Errorf(codes.Unauthenticated, "permission check failed")
	}
	return svc.service.AddMetadatas(ctx, request)
}

func (svc *authorizationService) DeleteMetadata(ctx context.Context, request *protobuf.IdRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.DeleteMetadata(ctx, request)
}

func (svc *authorizationService) GetMetadataSupplementaryDownloadUrl(ctx context.Context, request *grpc.SupplementaryIdRequest) (*grpc.SignedUrl, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_service)
	if err != nil {
		return nil, err
	}
	return svc.service.GetMetadataSupplementaryDownloadUrl(ctx, request)
}

func (svc *authorizationService) AddMetadataSupplementary(ctx context.Context, request *grpc.AddSupplementaryRequest) (*grpc.SignedUrl, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_service)
	if err != nil {
		return nil, err
	}
	return svc.service.AddMetadataSupplementary(ctx, request)
}

func (svc *authorizationService) DeleteMetadataSupplementary(ctx context.Context, request *grpc.SupplementaryIdRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_service)
	if err != nil {
		return nil, err
	}
	return svc.service.DeleteMetadataSupplementary(ctx, request)
}

func (svc *authorizationService) SetMetadataUploaded(ctx context.Context, request *protobuf.IdRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_service)
	if err != nil {
		return nil, err
	}
	return svc.service.SetMetadataUploaded(ctx, request)
}

func (svc *authorizationService) ProcessMetadata(ctx context.Context, request *protobuf.IdRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_edit)
	if err != nil {
		return nil, err
	}
	return svc.service.SetMetadataUploaded(ctx, request)
}

func (svc *authorizationService) GetMetadataPermissions(ctx context.Context, request *protobuf.IdRequest) (*grpc.Permissions, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.GetMetadataPermissions(ctx, request)
}

func (svc *authorizationService) AddMetadataPermissions(ctx context.Context, permission *grpc.Permissions) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, permission.Id, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.AddMetadataPermissions(ctx, permission)
}

func (svc *authorizationService) AddMetadataPermission(ctx context.Context, permission *grpc.Permission) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, permission.Id, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.AddMetadataPermission(ctx, permission)
}

func (svc *authorizationService) GetCollectionPermissions(ctx context.Context, request *protobuf.IdRequest) (*grpc.Permissions, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, request.Id, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.GetCollectionPermissions(ctx, request)
}

func (svc *authorizationService) AddCollectionPermission(ctx context.Context, permission *grpc.Permission) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, permission.Id, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.AddCollectionPermission(ctx, permission)
}

func (svc *authorizationService) GetTraits(ctx context.Context, request *protobuf.Empty) (*grpc.Traits, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetTraits(ctx, request)
}

func (svc *authorizationService) GetTrait(ctx context.Context, request *protobuf.IdRequest) (*grpc.Trait, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetTrait(ctx, request)
}

func (svc *authorizationService) GetModels(ctx context.Context, request *protobuf.Empty) (*grpc.Models, error) {
	// TODO: permission type?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetModels(ctx, request)
}

func (svc *authorizationService) GetModel(ctx context.Context, request *protobuf.IdRequest) (*grpc.Model, error) {
	// TODO: permission type?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetModel(ctx, request)
}

func (svc *authorizationService) GetStorageSystems(ctx context.Context, request *protobuf.Empty) (*grpc.StorageSystems, error) {
	// TODO: permission type?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetStorageSystems(ctx, request)
}

func (svc *authorizationService) GetStorageSystem(ctx context.Context, request *protobuf.IdRequest) (*grpc.StorageSystem, error) {
	// TODO: permission type?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetStorageSystem(ctx, request)
}

func (svc *authorizationService) GetStorageSystemModels(ctx context.Context, request *protobuf.IdRequest) (*grpc.StorageSystemModels, error) {
	// TODO: permission type?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetStorageSystemModels(ctx, request)
}

func (svc *authorizationService) GetWorkflows(ctx context.Context, request *protobuf.Empty) (*grpc.Workflows, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflows(ctx, request)
}

func (svc *authorizationService) GetWorkflow(ctx context.Context, request *protobuf.IdRequest) (*grpc.Workflow, error) {
	// TODO
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflow(ctx, request)
}

func (svc *authorizationService) GetWorkflowActivityInstances(ctx context.Context, request *protobuf.IdRequest) (*grpc.WorkflowActivityInstances, error) {
	// TODO
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflowActivityInstances(ctx, request)
}

func (svc *authorizationService) GetWorkflowActivityStorageSystems(ctx context.Context, request *grpc.WorkflowActivityIdRequest) (*grpc.WorkflowActivityStorageSystems, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflowActivityStorageSystems(ctx, request)
}

func (svc *authorizationService) GetWorkflowActivityPrompts(ctx context.Context, request *grpc.WorkflowActivityIdRequest) (*grpc.WorkflowActivityPrompts, error) {
	// TODO: maybe a trait permission object?
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflowActivityPrompts(ctx, request)
}

func (svc *authorizationService) GetWorkflowStates(ctx context.Context, request *protobuf.Empty) (*grpc.WorkflowStates, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_workflow_state_type, "all", grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetWorkflowStates(ctx, request)
}

func (svc *authorizationService) BeginTransitionWorkflow(ctx context.Context, request *grpc.TransitionWorkflowRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.MetadataId, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.BeginTransitionWorkflow(ctx, request)
}

func (svc *authorizationService) CompleteTransitionWorkflow(ctx context.Context, request *grpc.CompleteTransitionWorkflowRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.MetadataId, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.CompleteTransitionWorkflow(ctx, request)
}
