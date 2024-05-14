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
	"bosca.io/pkg/objectstore"
	"bosca.io/pkg/security"
	"bosca.io/pkg/security/identity"
	"bosca.io/pkg/workers/metadata"
	"context"
	"errors"
	"go.temporal.io/sdk/client"
	"log/slog"
	"strings"
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

func (svc *service) CheckPermissions(ctx context.Context, request *grpc.PermissionCheckRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithSubjectIdError(ctx, request.SubjectType, request.Subject, request.ObjectType, request.Object, request.Action)
	if err != nil {
		slog.ErrorContext(ctx, "failed permission check", slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) GetRootCollectionItems(ctx context.Context, request *protobuf.Empty) (*grpc.CollectionItems, error) {
	return svc.GetCollectionItems(ctx, &protobuf.IdRequest{Id: RootCollectionId})
}

func (svc *service) AddCollection(ctx context.Context, request *grpc.AddCollectionRequest) (*protobuf.IdResponse, error) {
	userId, err := identity.GetAuthenticatedSubjectId(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get subject", slog.Any("request", request), slog.Any("error", err))
		return nil, err
	}
	id, err := svc.ds.AddCollection(ctx, request.Collection)
	if err != nil {
		slog.ErrorContext(ctx, "failed to add collection", slog.Any("request", request), slog.Any("error", err))
		return nil, err
	}
	err = svc.permissions.CreateRelationships(ctx, grpc.PermissionObjectType_collection_type, []*grpc.Permission{
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
		slog.ErrorContext(ctx, "failed to create relationships", slog.String("id", id), slog.Any("error", err))
		return nil, err
	}
	err = svc.ds.AddCollectionCollectionItems(ctx, request.Parent, []string{id})
	if err != nil {
		slog.ErrorContext(ctx, "failed to add to parent", slog.String("id", id), slog.String("parent_id", request.Parent), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.IdResponse{Id: id}, nil
}

func (svc *service) DeleteCollection(ctx context.Context, request *protobuf.IdRequest) (*protobuf.Empty, error) {
	err := svc.ds.DeleteCollection(ctx, request.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to delete collection", slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) GetCollectionItems(ctx context.Context, request *protobuf.IdRequest) (*grpc.CollectionItems, error) {
	collectionItemIds, err := svc.ds.GetCollectionCollectionItemIds(ctx, request.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get collection item ids", slog.String("id", request.Id), slog.Any("error", err))
		return nil, err
	}
	metadataItemIds, err := svc.ds.GetCollectionMetadataItemIds(ctx, request.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get metadata item ids", slog.String("id", request.Id), slog.Any("error", err))
		return nil, err
	}
	items := make([]*grpc.CollectionItem, 0, len(collectionItemIds)+len(metadataItemIds))
	for _, item := range collectionItemIds {
		collection, err := svc.ds.GetCollection(ctx, item)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get collection", slog.String("id", request.Id), slog.String("item", item), slog.Any("error", err))
			return nil, err
		}
		err = svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_collection_type, item, grpc.PermissionAction_view)
		if err != nil {
			slog.InfoContext(ctx, "permission check failed", slog.String("id", request.Id), slog.String("item", item), slog.Any("error", err))
			continue
		}
		items = append(items, &grpc.CollectionItem{
			Item: &grpc.CollectionItem_Collection{
				Collection: collection,
			},
		})
	}
	for _, item := range metadataItemIds {
		meta, err := svc.ds.GetMetadata(ctx, item)
		if err != nil {
			slog.ErrorContext(ctx, "failed to get metadata", slog.String("id", request.Id), slog.String("item", item), slog.Any("error", err))
			return nil, err
		}
		err = svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, item, grpc.PermissionAction_view)
		if err != nil {
			slog.InfoContext(ctx, "view permission check failed", slog.String("id", request.Id), slog.String("item", item), slog.Any("error", err))
			continue
		}
		if meta.Status == grpc.MetadataStatus_processing {
			err = svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, item, grpc.PermissionAction_edit)
			if err != nil {
				slog.InfoContext(ctx, "edit permission check failed", slog.String("id", request.Id), slog.String("item", item), slog.Any("error", err))
				continue
			}
		}
		items = append(items, &grpc.CollectionItem{
			Item: &grpc.CollectionItem_Metadata{
				Metadata: meta,
			},
		})
	}
	return &grpc.CollectionItems{
		Items: items,
	}, nil
}

func (svc *service) AddMetadata(ctx context.Context, request *grpc.AddMetadataRequest) (*protobuf.IdResponse, error) {
	userId, err := identity.GetAuthenticatedSubjectId(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get subject", slog.Any("request", request), slog.Any("error", err))
		return nil, err
	}
	if request.Metadata.ContentLength <= 0 {
		err := errors.New("content length must be greater than 0")
		slog.ErrorContext(ctx, "content length requirement failed", slog.Any("request", request), slog.Any("error", err))
		return nil, err
	}
	id, err := svc.ds.AddMetadata(ctx, request.Metadata)
	if err != nil {
		slog.ErrorContext(ctx, "failed to add metadata", slog.Any("request", request), slog.Any("error", err))
		return nil, err
	}
	err = svc.permissions.CreateRelationships(ctx, grpc.PermissionObjectType_metadata_type, []*grpc.Permission{
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
		slog.ErrorContext(ctx, "failed to add metadata permissions", slog.String("id", id), slog.Any("error", err))
		return nil, err
	}
	err = svc.ds.AddCollectionMetadataItems(ctx, request.Collection, []string{id})
	if err != nil {
		slog.ErrorContext(ctx, "failed to add to collection", slog.String("id", id), slog.String("collection", request.Collection), slog.Any("error", err))
		return nil, err
	}
	if request.Metadata.Source != nil && *request.Metadata.Source != "" {
		_, err := svc.ProcessMetadata(ctx, &protobuf.IdRequest{Id: id})
		if err != nil {
			slog.ErrorContext(ctx, "failed to process metadata", slog.String("id", id), slog.Any("error", err))
			return nil, err
		}
	}
	return &protobuf.IdResponse{
		Id: id,
	}, nil
}

func (svc *service) DeleteMetadata(ctx context.Context, request *protobuf.IdRequest) (*protobuf.Empty, error) {
	err := svc.os.Delete(ctx, request.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to delete file", slog.String("id", request.Id), slog.Any("error", err))
		return nil, err
	}
	err = svc.ds.DeleteMetadata(ctx, request.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to delete metadata", slog.String("id", request.Id), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) GetMetadata(ctx context.Context, request *protobuf.IdRequest) (*grpc.Metadata, error) {
	return svc.ds.GetMetadata(ctx, request.Id)
}

func (svc *service) GetMetadataDownloadUrl(ctx context.Context, request *protobuf.IdRequest) (*grpc.SignedUrl, error) {
	md, err := svc.ds.GetMetadata(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	id := request.Id
	if md.Source != nil && *md.Source != "" {
		id = strings.Split(*md.Source, "+")[0]
	}
	return svc.os.CreateDownloadUrl(ctx, id)
}

func (svc *service) AddMetadataSupplementary(ctx context.Context, request *grpc.AddSupplementaryRequest) (*grpc.SignedUrl, error) {
	id := request.Id + "." + request.Type
	return svc.os.CreateUploadUrl(ctx, id, request.Name, request.ContentType, request.ContentLength, nil)
}

func (svc *service) GetMetadataSupplementaryDownloadUrl(ctx context.Context, request *grpc.SupplementaryIdRequest) (*grpc.SignedUrl, error) {
	id := request.Id + "." + request.Type
	return svc.os.CreateDownloadUrl(ctx, id)
}

func (svc *service) DeleteMetadataSupplementary(ctx context.Context, request *grpc.SupplementaryIdRequest) (*protobuf.Empty, error) {
	id := request.Id + "." + request.Type
	err := svc.os.Delete(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to delete supplementary file", slog.String("id", request.Id), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) SetMetadataUploaded(ctx context.Context, request *protobuf.IdRequest) (*protobuf.Empty, error) {
	return svc.ProcessMetadata(ctx, request)
}

func (svc *service) ProcessMetadata(ctx context.Context, request *protobuf.IdRequest) (*protobuf.Empty, error) {
	_, err := svc.temporalClient.ExecuteWorkflow(ctx, client.StartWorkflowOptions{
		TaskQueue: metadata.TaskQueue,
	}, metadata.ProcessMetadata, request.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to execute process workflow", slog.String("id", request.Id), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) SetMetadataStatus(ctx context.Context, request *grpc.SetMetadataStatusRequest) (*protobuf.Empty, error) {
	err := svc.ds.SetMetadataStatus(ctx, request.Id, request.Status)
	if err != nil {
		slog.ErrorContext(ctx, "failed to set metadata status", slog.String("id", request.Id), slog.String("variant_id", request.VariantId), slog.String("status", request.Status.String()), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) GetMetadataPermissions(ctx context.Context, request *protobuf.IdRequest) (*grpc.Permissions, error) {
	return svc.permissions.GetPermissions(ctx, grpc.PermissionObjectType_metadata_type, request.Id)
}

func (svc *service) AddMetadataPermissions(ctx context.Context, permissions *grpc.Permissions) (*protobuf.Empty, error) {
	for _, permission := range permissions.Permissions {
		if permission.Id == "" {
			permission.Id = permissions.Id
		}
	}
	err := svc.permissions.CreateRelationships(ctx, grpc.PermissionObjectType_metadata_type, permissions.Permissions)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create metadata permission", slog.Any("permissions", permissions), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) AddMetadataPermission(ctx context.Context, permission *grpc.Permission) (*protobuf.Empty, error) {
	err := svc.permissions.CreateRelationship(ctx, grpc.PermissionObjectType_metadata_type, permission)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create metadata permission", slog.Any("permission", permission), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) GetCollectionPermissions(ctx context.Context, request *protobuf.IdRequest) (*grpc.Permissions, error) {
	return svc.permissions.GetPermissions(ctx, grpc.PermissionObjectType_collection_type, request.Id)
}

func (svc *service) AddCollectionPermission(ctx context.Context, permission *grpc.Permission) (*protobuf.Empty, error) {
	err := svc.permissions.CreateRelationship(ctx, grpc.PermissionObjectType_collection_type, permission)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create collection permission", slog.Any("permission", permission), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}
