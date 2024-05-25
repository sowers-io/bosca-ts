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
	"bosca.io/api/protobuf"
	grpc "bosca.io/api/protobuf/content"
	"bosca.io/pkg/security"
	"context"
	"errors"
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
	if len(strings.Trim(request.Parent, " ")) == 0 {
		request.Parent = RootCollectionId
	}
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, request.Parent, grpc.PermissionAction_edit)
	if err != nil {
		return nil, err
	}
	return svc.service.AddCollection(ctx, request)
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

func (svc *authorizationService) GetMetadataDownloadUrl(ctx context.Context, request *protobuf.IdRequest) (*grpc.SignedUrl, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_view)
	if err != nil {
		return nil, err
	}
	return svc.service.GetMetadataDownloadUrl(ctx, request)
}

func (svc *authorizationService) AddMetadata(ctx context.Context, request *grpc.AddMetadataRequest) (*protobuf.IdResponse, error) {
	if request.Metadata == nil {
		return nil, errors.New("metadata is required")
	}
	if len(strings.Trim(request.Collection, " ")) == 0 {
		request.Collection = RootCollectionId
	}
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, request.Collection, grpc.PermissionAction_edit)
	if err != nil {
		return nil, err
	}
	return svc.service.AddMetadata(ctx, request)
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

func (svc *authorizationService) SetMetadataStatus(ctx context.Context, request *grpc.SetMetadataStatusRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_service)
	if err != nil {
		return nil, err
	}
	return svc.service.SetMetadataStatus(ctx, request)
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
