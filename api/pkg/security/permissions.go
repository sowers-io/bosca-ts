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
	grpc "bosca.io/api/protobuf/content"
	"bosca.io/pkg/identity"
	"context"
	pb "github.com/authzed/authzed-go/proto/authzed/api/v1"
	"github.com/authzed/authzed-go/v1"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
)

type ObjectType int

const (
	MetadataObject ObjectType = iota
	CollectionObject
	SystemResourceObject
)

type PermissionManager struct {
	permissionsClient *authzed.Client
}

func NewPermissionManager(permissionsClient *authzed.Client) *PermissionManager {
	return &PermissionManager{
		permissionsClient: permissionsClient,
	}
}

func (s *PermissionManager) CheckWithError(ctx context.Context, objectType ObjectType, resourceId string, action grpc.PermissionAction) error {
	userId, err := identity.GetAuthenticatedSubjectId(ctx)
	if err != nil {
		return err
	}
	return s.CheckWithUserIdError(ctx, userId, objectType, resourceId, action)
}

func (s *PermissionManager) CheckWithUserIdError(ctx context.Context, userId string, objectType ObjectType, resourceId string, action grpc.PermissionAction) error {
	r, err := s.permissionsClient.CheckPermission(ctx, &pb.CheckPermissionRequest{
		Resource: &pb.ObjectReference{
			ObjectType: s.getObjectType(objectType),
			ObjectId:   resourceId,
		},
		Permission: s.getAction(action),
		Subject: &pb.SubjectReference{
			Object: &pb.ObjectReference{
				ObjectType: "user",
				ObjectId:   userId,
			},
		},
	})
	if err != nil {
		return err
	}
	if r.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_NO_PERMISSION ||
		r.Permissionship == pb.CheckPermissionResponse_PERMISSIONSHIP_UNSPECIFIED {
		return status.Errorf(codes.Unauthenticated, "permission check failed")
	}
	return nil
}

/*

func (s *Manager) GetGroups(ctx context.Context) ([]string, error) {
	client, err := s.permissionsClient.LookupSubjects(ctx, &pb.LookupSubjectsRequest{
		Resource: "",
		SubjectObjectType: "group",
	})
	if err != nil {
		return nil, err
	}
	if err != nil {
		return nil, err
	}
	defer client.CloseSend()
	groups := make([]string, 0)
	for {
		result, err := client.Recv()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, err
		}
		groups = append(groups, result.Subject.SubjectObjectId)
	}
	return groups, nil
}

*/

func (s *PermissionManager) getObjectType(objectType ObjectType) string {
	switch objectType {
	case MetadataObject:
		return "metadata"
	case CollectionObject:
		return "collection"
	case SystemResourceObject:
		return "systemresource"
	}
	return ""
}

func (s *PermissionManager) getSubjectType(group bool) string {
	if group {
		return "group"
	} else {
		return "user"
	}
}

func (s *PermissionManager) getRelation(relation grpc.PermissionRelation) string {
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
		return "serviceaccounts"
	case grpc.PermissionRelation_owners:
		return "owners"
	}
	return ""
}

func (s *PermissionManager) getAction(relation grpc.PermissionAction) string {
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

func (s *PermissionManager) CreateRelationship(ctx context.Context, objectType ObjectType, permission *grpc.Permission) error {
	_, err := s.permissionsClient.WriteRelationships(ctx, &pb.WriteRelationshipsRequest{
		Updates: []*pb.RelationshipUpdate{
			{
				Operation: pb.RelationshipUpdate_OPERATION_CREATE,
				Relationship: &pb.Relationship{
					Resource: &pb.ObjectReference{
						ObjectType: s.getObjectType(objectType),
						ObjectId:   permission.Id,
					},
					Relation: s.getRelation(permission.Relation),
					Subject: &pb.SubjectReference{
						Object: &pb.ObjectReference{
							ObjectType: s.getSubjectType(permission.Group),
							ObjectId:   permission.Subject,
						},
					},
				},
			},
		},
	})
	if err != nil {
		return err
	}
	return nil
}

func (s *PermissionManager) GetPermissions(ctx context.Context, objectType ObjectType, resourceId string) (*grpc.Permissions, error) {
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
		case "serviceaccounts":
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
		permission.Group = relationship.Subject.Object.ObjectType == "group"
		mp.Permissions = append(mp.Permissions, permission)
	}
	return mp, nil
}
