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
	"log"
	"strings"
)

type serviceSecurity struct {
	grpc.UnimplementedContentServiceServer

	service     grpc.ContentServiceServer
	permissions *security.PermissionManager
}

func NewServiceSecurity(permissions *security.PermissionManager, dataSource *DataStore, service grpc.ContentServiceServer) grpc.ContentServiceServer {
	svc := &serviceSecurity{
		service:     service,
		permissions: permissions,
	}

	ctx := context.Background()
	if added, err := dataSource.AddRootCollection(ctx); added {
		if err != nil {
			log.Fatalf("error initializing root collection: %v", err)
		}
		err = svc.permissions.CreateRelationship(ctx, security.CollectionObject, &grpc.Permission{
			Id:       RootCollectionId,
			Subject:  "administrators",
			Group:    true,
			Relation: grpc.PermissionRelation_owners,
		})
		if err != nil {
			_ = dataSource.DeleteCollection(ctx, RootCollectionId)
			log.Fatalf("error initializing root collection permission: %v", err)
		}
	}

	return svc
}

func (svc *serviceSecurity) GetRootCollectionItems(ctx context.Context, request *protobuf.Empty) (*grpc.CollectionItems, error) {
	err := svc.permissions.CheckWithError(ctx, security.CollectionObject, RootCollectionId, grpc.PermissionAction_list)
	if err != nil {
		return nil, err
	}
	return svc.service.GetRootCollectionItems(ctx, request)
}

func (svc *serviceSecurity) AddMetadata(ctx context.Context, request *grpc.AddMetadataRequest) (*grpc.SignedUrl, error) {
	if request.Metadata == nil {
		return nil, errors.New("metadata is required")
	}
	if len(strings.Trim(request.Collection, " ")) == 0 {
		return nil, errors.New("collection is required")
	}
	err := svc.permissions.CheckWithError(ctx, security.CollectionObject, request.Collection, grpc.PermissionAction_edit)
	if err != nil {
		return nil, err
	}
	return svc.service.AddMetadata(ctx, request)
}

func (svc *serviceSecurity) SetMetadataStatus(ctx context.Context, request *grpc.SetMetadataStatusRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, security.MetadataObject, request.Id, grpc.PermissionAction_service)
	if err != nil {
		return nil, err
	}
	return svc.service.SetMetadataStatus(ctx, request)
}

func (svc *serviceSecurity) GetMetadataPermissions(ctx context.Context, request *protobuf.IdRequest) (*grpc.Permissions, error) {
	err := svc.permissions.CheckWithError(ctx, security.MetadataObject, request.Id, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.GetMetadataPermissions(ctx, request)
}

func (svc *serviceSecurity) AddMetadataPermission(ctx context.Context, permission *grpc.Permission) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, security.MetadataObject, permission.Id, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.AddMetadataPermission(ctx, permission)
}

func (svc *serviceSecurity) GetCollectionPermissions(ctx context.Context, request *protobuf.IdRequest) (*grpc.Permissions, error) {
	err := svc.permissions.CheckWithError(ctx, security.CollectionObject, request.Id, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.GetCollectionPermissions(ctx, request)
}

func (svc *serviceSecurity) AddCollectionPermission(ctx context.Context, permission *grpc.Permission) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithError(ctx, security.CollectionObject, permission.Id, grpc.PermissionAction_manage)
	if err != nil {
		return nil, err
	}
	return svc.service.AddCollectionPermission(ctx, permission)
}
