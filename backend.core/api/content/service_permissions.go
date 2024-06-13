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
	"context"
	"log/slog"
)

func (svc *service) CheckPermissions(ctx context.Context, request *grpc.PermissionCheckRequest) (*protobuf.Empty, error) {
	err := svc.permissions.CheckWithSubjectIdError(ctx, request.SubjectType, request.Subject, request.ObjectType, request.Object, request.Action)
	if err != nil {
		slog.ErrorContext(ctx, "failed permission check", slog.Any("error", err))
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
		slog.ErrorContext(ctx, "failed to create workflow permission", slog.Any("permissions", permissions), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) AddMetadataPermission(ctx context.Context, permission *grpc.Permission) (*protobuf.Empty, error) {
	err := svc.permissions.CreateRelationship(ctx, grpc.PermissionObjectType_metadata_type, permission)
	if err != nil {
		slog.ErrorContext(ctx, "failed to create workflow permission", slog.Any("permission", permission), slog.Any("error", err))
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
