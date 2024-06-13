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
	"bosca.io/pkg/security/identity"
	"context"
	"errors"
	"log/slog"
	"strings"
)

func (svc *service) GetRootCollectionItems(ctx context.Context, request *protobuf.Empty) (*grpc.CollectionItems, error) {
	return svc.GetCollectionItems(ctx, &protobuf.IdRequest{Id: RootCollectionId})
}

func (svc *service) GetCollection(ctx context.Context, request *protobuf.IdRequest) (*grpc.Collection, error) {
	return svc.ds.GetCollection(ctx, request.Id)
}

func (svc *service) AddCollection(ctx context.Context, request *grpc.AddCollectionRequest) (*protobuf.IdResponse, error) {
	userId, err := identity.GetAuthenticatedSubjectId(ctx)
	if err != nil {
		slog.ErrorContext(ctx, "failed to get subject", slog.Any("request", request), slog.Any("error", err))
		return nil, err
	}
	if strings.Trim(request.Collection.Name, " ") == "" {
		err := errors.New("name must not be empty")
		slog.ErrorContext(ctx, "name requirement failed", slog.Any("request", request), slog.Any("error", err))
		return nil, err
	}
	if unique, err := svc.verifyUniqueName(ctx, request.Parent, request.Collection.Name); !unique || err != nil {
		if err == nil {
			return nil, errors.New("name must be unique")
		}
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
		slog.ErrorContext(ctx, "failed to get workflow item ids", slog.String("id", request.Id), slog.Any("error", err))
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
			slog.ErrorContext(ctx, "failed to get workflow", slog.String("id", request.Id), slog.String("item", item), slog.Any("error", err))
			return nil, err
		}
		if meta.WorkflowStateId != WorkflowStatePublished {
			err = svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, item, grpc.PermissionAction_edit)
			if err != nil {
				slog.InfoContext(ctx, "edit permission check failed, not returning workflow", slog.String("id", request.Id), slog.String("item", item), slog.Any("error", err))
				continue
			}
		} else {
			err = svc.permissions.CheckWithError(ctx, grpc.PermissionObjectType_metadata_type, item, grpc.PermissionAction_view)
			if err != nil {
				slog.InfoContext(ctx, "view permission check failed", slog.String("id", request.Id), slog.String("item", item), slog.Any("error", err))
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
