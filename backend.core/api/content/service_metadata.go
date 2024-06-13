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
	"fmt"
	"log/slog"
	"slices"
	"strings"
)

func (svc *service) verifyUniqueName(ctx context.Context, collectionId string, name string) (bool, error) {
	collectionNames, err := svc.ds.GetCollectionCollectionItemNames(ctx, collectionId)
	if err != nil {
		return false, err
	}
	if slices.Contains(collectionNames, name) {
		return false, nil
	}
	metadataNames, err := svc.ds.GetCollectionMetadataItemNames(ctx, collectionId)
	if err != nil {
		return false, err
	}
	if slices.Contains(metadataNames, name) {
		return false, nil
	}
	return true, nil
}

func (svc *service) AddMetadataRelationship(ctx context.Context, request *grpc.AddMetadataRelationshipRequest) (*protobuf.Empty, error) {
	err := svc.ds.AddMetadataRelationship(ctx, request.MetadataId1, request.MetadataId2, request.Relationship)
	if err != nil {
		return nil, err
	}
	return &protobuf.Empty{}, nil
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
	if strings.Trim(request.Metadata.Name, " ") == "" {
		err := errors.New("name must not be empty")
		slog.ErrorContext(ctx, "name requirement failed", slog.Any("request", request), slog.Any("error", err))
		return nil, err
	}
	if unique, err := svc.verifyUniqueName(ctx, request.Collection, request.Metadata.Name); !unique || err != nil {
		if err == nil {
			return nil, errors.New("name must be unique")
		}
		return nil, err
	}
	id, err := svc.ds.AddMetadata(ctx, request.Metadata)
	if err != nil {
		slog.ErrorContext(ctx, "failed to add workflow", slog.Any("request", request), slog.Any("error", err))
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
		slog.ErrorContext(ctx, "failed to add workflow permissions", slog.String("id", id), slog.Any("error", err))
		return nil, err
	}
	err = svc.ds.AddCollectionMetadataItems(ctx, request.Collection, []string{id})
	if err != nil {
		slog.ErrorContext(ctx, "failed to add to collection", slog.String("id", id), slog.String("collection", request.Collection), slog.Any("error", err))
		return nil, err
	}
	if request.Metadata.Source != nil && *request.Metadata.Source != "" {
		_, err := svc.BeginTransitionWorkflow(ctx, &grpc.TransitionWorkflowRequest{MetadataId: id, StateId: WorkflowStateProcessing})
		if err != nil {
			slog.ErrorContext(ctx, "failed to process workflow", slog.String("id", id), slog.Any("error", err))
			return nil, err
		}
	}
	return &protobuf.IdResponse{
		Id: id,
	}, nil
}

func (svc *service) DeleteMetadata(ctx context.Context, request *protobuf.IdRequest) (*protobuf.Empty, error) {
	md, err := svc.ds.GetMetadata(ctx, request.Id)
	if md == nil {
		return &protobuf.Empty{}, nil
	}
	if err != nil {
		slog.ErrorContext(ctx, "failed to get workflow to delete", slog.String("id", request.Id), slog.Any("error", err))
		return nil, err
	}
	if md.Source != nil && *md.Source != "" {
		err = svc.os.Delete(ctx, strings.Split(*md.Source, "+")[0])
		if err != nil {
			slog.ErrorContext(ctx, "failed to delete file", slog.String("id", request.Id), slog.Any("error", err))
			return nil, err
		}
	}
	err = svc.os.Delete(ctx, request.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to delete file", slog.String("id", request.Id), slog.Any("error", err))
		return nil, err
	}
	err = svc.ds.DeleteMetadata(ctx, request.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to delete workflow", slog.String("id", request.Id), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) GetMetadataUploadUrl(ctx context.Context, request *protobuf.IdRequest) (*grpc.SignedUrl, error) {
	metadata, err := svc.ds.GetMetadata(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	if metadata.WorkflowStateId != WorkflowStatePending {
		return nil, errors.New("invalid workflow state")
	}
	return svc.os.CreateUploadUrl(ctx, metadata.Id, metadata.Name, metadata.ContentType, metadata.ContentLength, nil)
}

func (svc *service) AddMetadataTrait(ctx context.Context, request *grpc.AddMetadataTraitRequest) (*grpc.Metadata, error) {
	metadata, err := svc.ds.AddMetadataTrait(ctx, request.MetadataId, request.TraitId)
	if err != nil {
		return nil, err
	}
	_, err = svc.BeginTransitionWorkflow(ctx, &grpc.TransitionWorkflowRequest{MetadataId: request.MetadataId, Status: fmt.Sprintf("adding trait: %s", request.TraitId), StateId: WorkflowStateProcessing})
	if err != nil {
		return nil, err
	}
	return metadata, nil
}

func (svc *service) GetMetadata(ctx context.Context, request *protobuf.IdRequest) (*grpc.Metadata, error) {
	return svc.ds.GetMetadata(ctx, request.Id)
}

func (svc *service) GetMetadatas(ctx context.Context, request *protobuf.IdsRequest) (*grpc.Metadatas, error) {
	ids, err := svc.permissions.BulkCheck(ctx, grpc.PermissionObjectType_metadata_type, request.Id, grpc.PermissionAction_view)
	if err != nil {
		return nil, err
	}
	m, err := svc.ds.GetMetadatas(ctx, ids)
	if err != nil {
		return nil, err
	}
	return &grpc.Metadatas{Metadata: m}, nil
}

func (svc *service) GetMetadataDownloadUrl(ctx context.Context, request *protobuf.IdRequest) (*grpc.SignedUrl, error) {
	md, err := svc.ds.GetMetadata(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	if md == nil {
		return nil, nil
	}
	id := request.Id
	if md.Source != nil && *md.Source != "" {
		id = strings.Split(*md.Source, "+")[0]
	}
	return svc.os.CreateDownloadUrl(ctx, id)
}

func (svc *service) SetMetadataUploaded(ctx context.Context, request *protobuf.IdRequest) (*protobuf.Empty, error) {
	return svc.BeginTransitionWorkflow(ctx, &grpc.TransitionWorkflowRequest{MetadataId: request.Id, StateId: WorkflowStateProcessing})
}
