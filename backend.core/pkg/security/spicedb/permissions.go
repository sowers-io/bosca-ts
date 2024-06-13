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

package spicedb

import (
	grpc "bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/security"
	"bosca.io/pkg/security/identity"
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log/slog"
	"time"
)

const SubjectTypeGroup = "group"
const SubjectTypeUser = "user"
const SubjectTypeServiceAccount = "serviceaccount"

type permissionManager struct {
	permissionsClient *authzed.Client
}

func NewPermissionManager(permissionsClient *authzed.Client) security.PermissionManager {
	return &permissionManager{
		permissionsClient: permissionsClient,
	}
}

func (s *permissionManager) BulkCheck(ctx context.Context, objectType grpc.PermissionObjectType, resourceId []string, action grpc.PermissionAction) ([]string, error) {
	subjectId, err := identity.GetSubjectId(ctx)
	if err != nil {
		return nil, err
	}
	subjectType := grpc.PermissionSubjectType_user
	if identity.GetSubjectType(ctx) == identity.SubjectTypeServiceAccount {
		subjectType = grpc.PermissionSubjectType_service_account
	}
	items := make([]*pb.CheckBulkPermissionsRequestItem, 0, len(resourceId))
	for _, id := range resourceId {
		items = append(items, &pb.CheckBulkPermissionsRequestItem{
			Subject: &pb.SubjectReference{
				Object: &pb.ObjectReference{
					ObjectType: s.getSubjectType(subjectType),
					ObjectId:   subjectId,
				},
			},
			Resource: &pb.ObjectReference{
				ObjectType: s.getObjectType(objectType),
				ObjectId:   id,
			},
			Permission: s.getAction(action),
		})
	}

	r, err := s.permissionsClient.CheckBulkPermissions(ctx, &pb.CheckBulkPermissionsRequest{
		Items: items,
	})
	if err != nil {
		return nil, err
	}
	ids := make([]string, 0, len(resourceId))
	for i, item := range r.Pairs {
		pair := item.GetItem()
		if pair.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_HAS_PERMISSION ||
			pair.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_CONDITIONAL_PERMISSION {
			ids = append(ids, resourceId[i])
		}
	}
	return ids, nil
}

func (s *permissionManager) CheckWithError(ctx context.Context, objectType grpc.PermissionObjectType, resourceId string, action grpc.PermissionAction) error {
	subjectId, err := identity.GetSubjectId(ctx)
	if err != nil {
		return err
	}
	subjectType := grpc.PermissionSubjectType_user
	if identity.GetSubjectType(ctx) == identity.SubjectTypeServiceAccount {
		subjectType = grpc.PermissionSubjectType_service_account
	}
	return s.CheckWithSubjectIdError(ctx, subjectType, subjectId, objectType, resourceId, action)
}

func (s *permissionManager) CheckWithSubjectIdError(ctx context.Context, subjectType grpc.PermissionSubjectType, subjectId string, objectType grpc.PermissionObjectType, objectId string, action grpc.PermissionAction) error {
	objectTypeString := s.getObjectType(objectType)
	subjectTypeString := s.getSubjectType(subjectType)
	actionString := s.getAction(action)

	logAttrs := []any{
		slog.String("objectType", objectTypeString),
		slog.String("objectId", objectId),
		slog.String("subjectType", subjectTypeString),
		slog.String("subjectId", subjectId),
		slog.String("action", actionString),
	}

	slog.DebugContext(ctx, "checking permissions", logAttrs...)

	check := &pb.CheckPermissionRequest{
		Resource: &pb.ObjectReference{
			ObjectType: objectTypeString,
			ObjectId:   objectId,
		},
		Permission: actionString,
		Subject: &pb.SubjectReference{
			Object: &pb.ObjectReference{
				ObjectType: subjectTypeString,
				ObjectId:   subjectId,
			},
		},
	}
	r, err := s.permissionsClient.CheckPermission(ctx, check)
	if err != nil {
		slog.DebugContext(ctx, "error checking permissions", logAttrs...)
		return err
	}
	if r.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_NO_PERMISSION ||
		r.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_UNSPECIFIED {
		slog.DebugContext(ctx, "permission check failed", logAttrs...)
		return status.Errorf(codes.Unauthenticated, "permission check failed")
	}
	slog.DebugContext(ctx, "permission check successful", logAttrs...)
	return nil
}

func (s *permissionManager) getObjectType(objectType grpc.PermissionObjectType) string {
	switch objectType {
	case grpc.PermissionObjectType_metadata_type:
		return "metadata"
	case grpc.PermissionObjectType_collection_type:
		return "collection"
	case grpc.PermissionObjectType_system_resource_type:
		return "systemresource"
	case grpc.PermissionObjectType_workflow_type:
		return "workflow"
	case grpc.PermissionObjectType_workflow_state_type:
		return "workflowstate"
	}
	return ""
}

func (s *permissionManager) getSubjectType(objectType grpc.PermissionSubjectType) string {
	switch objectType {
	case grpc.PermissionSubjectType_user:
		return "user"
	case grpc.PermissionSubjectType_group:
		return "group"
	case grpc.PermissionSubjectType_service_account:
		return "serviceaccount"
	}
	return ""
}

func (s *permissionManager) getRelation(relation grpc.PermissionRelation) string {
	switch relation {
	case grpc.PermissionRelation_viewers:
		return "viewers"
	case grpc.PermissionRelation_discoverers:
		return "discoverers"
	case grpc.PermissionRelation_editors:
		return "editors"
	case grpc.PermissionRelation_managers:
		return "managers"
	case grpc.PermissionRelation_serviceaccounts:
		return "servicers"
	case grpc.PermissionRelation_owners:
		return "owners"
	}
	return ""
}

func (s *permissionManager) getAction(relation grpc.PermissionAction) string {
	switch relation {
	case grpc.PermissionAction_view:
		return "view"
	case grpc.PermissionAction_list:
		return "list"
	case grpc.PermissionAction_edit:
		return "edit"
	case grpc.PermissionAction_manage:
		return "manage"
	case grpc.PermissionAction_service:
		return "service"
	case grpc.PermissionAction_delete:
		return "delete"
	}
	return ""
}

func (s *permissionManager) CreateRelationships(ctx context.Context, objectType grpc.PermissionObjectType, permissions []*grpc.Permission) error {
	updates := make([]*pb.RelationshipUpdate, 0)
	for _, permission := range permissions {
		subjectType := ""
		switch permission.SubjectType {
		case grpc.PermissionSubjectType_user:
			subjectType = SubjectTypeUser
			break
		case grpc.PermissionSubjectType_group:
			subjectType = SubjectTypeGroup
			break
		case grpc.PermissionSubjectType_service_account:
			subjectType = SubjectTypeServiceAccount
			break
		}
		updates = append(updates, &pb.RelationshipUpdate{
			Operation: pb.RelationshipUpdate_OPERATION_CREATE,
			Relationship: &pb.Relationship{
				Resource: &pb.ObjectReference{
					ObjectType: s.getObjectType(objectType),
					ObjectId:   permission.Id,
				},
				Relation: s.getRelation(permission.Relation),
				Subject: &pb.SubjectReference{
					Object: &pb.ObjectReference{
						ObjectType: subjectType,
						ObjectId:   permission.Subject,
					},
				},
			},
		})
	}
	_, err := s.permissionsClient.WriteRelationships(ctx, &pb.WriteRelationshipsRequest{Updates: updates})

	if err != nil {
		return err
	}

	for _, permission := range permissions {
		for {
			current, err := s.GetPermissions(ctx, objectType, permission.Id)
			if err != nil {
				return err
			}
			if current != nil {
				match := false
				for _, p := range current.Permissions {
					if permission.Id == p.Id && permission.SubjectType == p.SubjectType && permission.Subject == p.Subject && permission.Relation == p.Relation {
						match = true
						break
					}
				}
				if match {
					break
				} else {
					slog.Debug("permission not found while creating relationships, trying again in 500ms")
					time.Sleep(500 * time.Millisecond)
				}
			}
		}
	}

	return err
}

func (s *permissionManager) CreateRelationship(ctx context.Context, objectType grpc.PermissionObjectType, permission *grpc.Permission) error {
	return s.CreateRelationships(ctx, objectType, []*grpc.Permission{permission})
}

func (s *permissionManager) GetPermissions(ctx context.Context, objectType grpc.PermissionObjectType, resourceId string) (*grpc.Permissions, error) {
	client, err := s.permissionsClient.ReadRelationships(ctx, &pb.ReadRelationshipsRequest{
		RelationshipFilter: &pb.RelationshipFilter{
			ResourceType:       s.getObjectType(objectType),
			OptionalResourceId: resourceId,
		},
	})
	if err != nil {
		return nil, err
	}
	defer client.CloseSend()
	mp := &grpc.Permissions{
		Permissions: make([]*grpc.Permission, 0),
	}
	for {
		result, err := client.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		relationship := result.Relationship
		action := grpc.PermissionRelation_viewers
		switch relationship.Relation {
		case "viewers":
			action = grpc.PermissionRelation_viewers
			break
		case "discoverers":
			action = grpc.PermissionRelation_discoverers
			break
		case "editors":
			action = grpc.PermissionRelation_editors
			break
		case "managers":
			action = grpc.PermissionRelation_managers
			break
		case "servicers":
			action = grpc.PermissionRelation_serviceaccounts
			break
		case "owners":
			action = grpc.PermissionRelation_owners
			break
		}
		permission := &grpc.Permission{
			Id:       relationship.Resource.ObjectId,
			Relation: action,
		}
		permission.Subject = relationship.Subject.Object.ObjectId
		switch relationship.Subject.Object.ObjectType {
		case SubjectTypeUser:
			permission.SubjectType = grpc.PermissionSubjectType_user
			break
		case SubjectTypeGroup:
			permission.SubjectType = grpc.PermissionSubjectType_group
			break
		case SubjectTypeServiceAccount:
			permission.SubjectType = grpc.PermissionSubjectType_service_account
			break
		}
		mp.Permissions = append(mp.Permissions, permission)
	}
	return mp, nil
}
