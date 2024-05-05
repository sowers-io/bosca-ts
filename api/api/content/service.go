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
	"bosca.io/pkg/identity"
	"bosca.io/pkg/objectstore"
	"bosca.io/pkg/security"
	"bosca.io/pkg/workers/metadata"
	"context"
	"go.temporal.io/sdk/client"
)

type service struct {
	grpc.UnimplementedContentServiceServer

	ds *DataStore
	os objectstore.ObjectStore

	serviceAccountId string
	permissions      security.PermissionManager
	temporalClient   client.Client
}

const RootCollectionId = "00000000-0000-0000-0000-000000000000"

func NewService(dataStore *DataStore, serviceAccountId string, objectStore objectstore.ObjectStore, permissions security.PermissionManager, temporalClient client.Client) grpc.ContentServiceServer {
	return &service{
		ds:               dataStore,
		os:               objectStore,
		serviceAccountId: serviceAccountId,
		permissions:      permissions,
		temporalClient:   temporalClient,
	}
}

func (svc *service) GetRootCollectionItems(context.Context, *protobuf.Empty) (*grpc.CollectionItems, error) {
	//TODO implement me
	panic("implement me")
}

func (svc *service) AddMetadata(ctx context.Context, request *grpc.AddMetadataRequest) (*grpc.SignedUrl, error) {
	userId, err := identity.GetAuthenticatedSubjectId(ctx)
	if err != nil {
		return nil, err
	}
	id, err := svc.ds.AddMetadata(ctx, request.Metadata)
	if err != nil {
		return nil, err
	}
	err = svc.permissions.CreateRelationships(ctx, security.MetadataObject, []*grpc.Permission{
		{
			Id:          id,
			Subject:     security.AdministratorGroup,
			SubjectType: grpc.PermissionSubjectType_group,
			Relation:    grpc.PermissionRelation_owners,
		},
		{
			Id:          id,
			Subject:     svc.serviceAccountId,
			SubjectType: grpc.PermissionSubjectType_service_account,
			Relation:    grpc.PermissionRelation_serviceaccounts,
		},
		{
			Id:          id,
			Subject:     userId,
			SubjectType: grpc.PermissionSubjectType_user,
			Relation:    grpc.PermissionRelation_owners,
		},
	})
	if err != nil {
		return nil, err
	}
	err = svc.ds.AddCollectionMetadataItems(ctx, request.Collection, []string{id})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	return svc.os.CreateUploadUrl(ctx, id, request.Metadata.Name, request.Metadata.ContentType, request.Metadata.Attributes)
}

func (svc *service) GetMetadata(ctx context.Context, request *protobuf.IdRequest) (*grpc.Metadata, error) {
	return svc.ds.GetMetadata(ctx, request.Id)
}

func (svc *service) GetMetadataDownloadUrl(ctx context.Context, request *protobuf.IdRequest) (*grpc.SignedUrl, error) {
	return svc.os.CreateDownloadUrl(ctx, request.Id)
}

func (svc *service) SetMetadataUploaded(ctx context.Context, request *protobuf.IdRequest) (*protobuf.Empty, error) {
	_, err := svc.temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		TaskQueue: metadata.TaskQueue,
	}, metadata.Stringify, request.Id)
	if err != nil {
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) SetMetadataStatus(ctx context.Context, request *grpc.SetMetadataStatusRequest) (*protobuf.Empty, error) {
	err := svc.ds.SetMetadataStatus(ctx, request.Id, request.Status)
	if err != nil {
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) GetMetadataPermissions(ctx context.Context, request *protobuf.IdRequest) (*grpc.Permissions, error) {
	return svc.permissions.GetPermissions(ctx, security.MetadataObject, request.Id)
}

func (svc *service) AddMetadataPermissions(ctx context.Context, permissions *grpc.Permissions) (*protobuf.Empty, error) {
	for _, permission := range permissions.Permissions {
		if permission.Id == "" {
			permission.Id = permissions.Id
		}
	}
	err := svc.permissions.CreateRelationships(ctx, security.MetadataObject, permissions.Permissions)
	if err != nil {
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) AddMetadataPermission(ctx context.Context, permission *grpc.Permission) (*protobuf.Empty, error) {
	err := svc.permissions.CreateRelationship(ctx, security.MetadataObject, permission)
	if err != nil {
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) GetCollectionPermissions(ctx context.Context, request *protobuf.IdRequest) (*grpc.Permissions, error) {
	return svc.permissions.GetPermissions(ctx, security.CollectionObject, request.Id)
}

func (svc *service) AddCollectionPermission(ctx context.Context, permission *grpc.Permission) (*protobuf.Empty, error) {
	err := svc.permissions.CreateRelationship(ctx, security.CollectionObject, permission)
	if err != nil {
		return nil, err
	}
	return &protobuf.Empty{}, nil
}
