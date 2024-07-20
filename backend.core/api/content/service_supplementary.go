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

func (svc *service) getSupplementaryId(metadataId, key string) string {
	return metadataId + "." + key
}

func (svc *service) GetMetadataSupplementary(ctx context.Context, request *protobuf.SupplementaryIdRequest) (*grpc.MetadataSupplementary, error) {
	return svc.ds.GetMetadataSupplementary(ctx, request.Id, request.Key)
}

func (svc *service) GetMetadataSupplementaries(ctx context.Context, request *protobuf.IdRequest) (*grpc.MetadataSupplementaries, error) {
	s, err := svc.ds.GetMetadataSupplementaries(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &grpc.MetadataSupplementaries{
		Supplementaries: s,
	}, nil
}

func (svc *service) AddMetadataSupplementary(ctx context.Context, request *grpc.AddSupplementaryRequest) (*grpc.MetadataSupplementary, error) {
	err := svc.ds.AddMetadataSupplementary(ctx, request.MetadataId, request.Key, request.Name, request.ContentType, request.ContentLength, request.TraitIds, request.SourceId, request.SourceIdentifier)
	if err != nil {
		return nil, err
	}
	return &grpc.MetadataSupplementary{
		Key:              request.Key,
		Name:             request.Name,
		ContentType:      request.ContentType,
		ContentLength:    request.ContentLength,
		TraitIds:         request.TraitIds,
		MetadataId:       request.MetadataId,
		SourceIdentifier: request.SourceIdentifier,
		SourceId:         request.SourceId,
	}, err
}

func (svc *service) SetMetadataSupplementaryReady(ctx context.Context, request *protobuf.SupplementaryIdRequest) (*protobuf.Empty, error) {
	err := svc.ds.SetMetadataSupplementaryReady(ctx, request.Id, request.Key)
	if err != nil {
		return nil, err
	}
	return &protobuf.Empty{}, nil
}

func (svc *service) GetMetadataSupplementaryUploadUrl(ctx context.Context, request *protobuf.SupplementaryIdRequest) (*grpc.SignedUrl, error) {
	id := svc.getSupplementaryId(request.Id, request.Key)
	supplementary, err := svc.ds.GetMetadataSupplementary(ctx, request.Id, request.Key)
	if err != nil {
		return nil, err
	}
	if supplementary == nil {
		return nil, nil
	}
	return svc.objectStore.CreateUploadUrl(ctx, id, supplementary.Name, supplementary.ContentType, supplementary.ContentLength, nil)
}

func (svc *service) GetMetadataSupplementaryDownloadUrl(ctx context.Context, request *protobuf.SupplementaryIdRequest) (*grpc.SignedUrl, error) {
	id := svc.getSupplementaryId(request.Id, request.Key)
	supplementary, err := svc.ds.GetMetadataSupplementary(ctx, request.Id, request.Key)
	if err != nil {
		return nil, err
	}
	if supplementary == nil {
		return nil, nil
	}
	return svc.objectStore.CreateDownloadUrl(ctx, id)
}

func (svc *service) DeleteMetadataSupplementary(ctx context.Context, request *protobuf.SupplementaryIdRequest) (*protobuf.Empty, error) {
	id := svc.getSupplementaryId(request.Id, request.Key)
	err := svc.ds.DeleteMetadataSupplementary(ctx, request.Id, request.Key)
	if err != nil {
		return nil, err
	}
	err = svc.objectStore.Delete(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to delete supplementary file", slog.String("id", request.Id), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}
