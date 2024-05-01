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

package security

import (
	"bosca.io/api/protobuf/content"
	"context"
)

type ObjectType int

const AdministratorGroup = "administrators"

const (
	MetadataObject ObjectType = iota
	CollectionObject
	SystemResourceObject
)

type PermissionManager interface {
	CheckWithError(ctx context.Context, objectType ObjectType, resourceId string, action content.PermissionAction) error
	CheckWithUserIdError(ctx context.Context, userId string, objectType ObjectType, resourceId string, action content.PermissionAction) error
	CreateRelationships(ctx context.Context, objectType ObjectType, permissions []*content.Permission) error
	CreateRelationship(ctx context.Context, objectType ObjectType, permission *content.Permission) error
	GetPermissions(ctx context.Context, objectType ObjectType, resourceId string) (*content.Permissions, error)
}
