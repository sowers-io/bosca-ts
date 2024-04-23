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
	"bosca.io/pkg/permissions"
	"context"
	ory "github.com/ory/client-go"
	"log"
)

type service struct {
	grpc.UnimplementedContentServiceServer

	ds *DataStore
	os ObjectStore

	oryClient *ory.APIClient
}

const RootCollectionId = "00000000-0000-0000-0000-000000000000"

func NewService(dataStore *DataStore, objectStore ObjectStore, oryClient *ory.APIClient) grpc.ContentServiceServer {
	svc := &service{
		ds:        dataStore,
		os:        objectStore,
		oryClient: oryClient,
	}

	ctx := context.Background()
	if added, err := svc.ds.AddRootCollection(ctx); added {
		if err != nil {
			log.Fatalf("error initializing root collection: %v", err)
		}
		_, err = svc.AddCollectionPermission(ctx, &grpc.Permission{
			Id:       RootCollectionId,
			Subject:  "administrators",
			Group:    true,
			Relation: grpc.PermissionRelation_owners,
		})
		if err != nil {
			_ = svc.ds.DeleteCollection(ctx, RootCollectionId)
			log.Fatalf("error initializing root collection permission: %v", err)
		}
	}

	return svc
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
	_, err = svc.AddMetadataPermission(ctx, &grpc.Permission{
		Id:       id,
		Subject:  userId,
		Relation: grpc.PermissionRelation_owners,
	})
	if err != nil {
		return nil, err
	}
	return svc.os.CreateUploadUrl(ctx, id, request.Metadata.Name, request.Metadata.ContentType, request.Metadata.Attributes)
}

func (svc *service) GetMetadataPermissions(ctx context.Context, request *protobuf.IdRequest) (*grpc.Permissions, error) {
	r := svc.oryClient.RelationshipAPI.GetRelationships(ctx)
	r.Namespace(permissions.PermissionsNamespace)
	r.Object(request.Id)
	r.PageSize(100)
	result, _, err := svc.oryClient.RelationshipAPI.GetRelationshipsExecute(r)
	if err != nil {
		return nil, err
	}
	mp := &grpc.Permissions{
		Permissions: make([]*grpc.Permission, 0),
	}
	for _, relation := range result.RelationTuples {
		action := grpc.PermissionRelation_viewers
		switch relation.Relation {
		case "viewers":
			action = grpc.PermissionRelation_viewers
			break
		case "editors":
			action = grpc.PermissionRelation_editors
			break
		case "owners":
			action = grpc.PermissionRelation_owners
			break
		}
		permission := &grpc.Permission{
			Id:       relation.Object,
			Relation: action,
		}
		if relation.HasSubjectId() {
			permission.Subject = *relation.SubjectId
		} else if relation.HasSubjectSet() {
			// KJB: TODO: I don't think this is right
			permission.Subject = (*relation.SubjectSet).Object
		}
		mp.Permissions = append(mp.Permissions, permission)
	}
	// TODO: handle paging
	return mp, nil
}

func (svc *service) AddMetadataPermission(ctx context.Context, permission *grpc.Permission) (*protobuf.Empty, error) {
	var subject string
	if permission.Group {
		subject = "Group:"
	} else {
		subject = "User:"
	}
	subject += permission.Subject
	err := svc.createRelationship(ctx, "Metadata:"+permission.Id, permission.Relation.String(), subject)
	if err != nil {
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) AddCollectionPermission(ctx context.Context, permission *grpc.Permission) (*protobuf.Empty, error) {
	var subject string
	if permission.Group {
		subject = "Group:"
	} else {
		subject = "User:"
	}
	subject += permission.Subject
	err := svc.createRelationship(ctx, "Collection:"+permission.Id, permission.Relation.String(), subject)
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

func (svc *service) createRelationship(ctx context.Context, object string, relation string, subject string) error {
	r := svc.oryClient.RelationshipAPI.CreateRelationship(ctx)
	ns := permissions.PermissionsNamespace
	body := ory.CreateRelationshipBody{
		Namespace: &ns,
		Object:    &object,
		Relation:  &relation,
	}
	body.SubjectId = &subject
	r = r.CreateRelationshipBody(body)
	_, _, err := svc.oryClient.RelationshipAPI.CreateRelationshipExecute(r)
	return err
}
