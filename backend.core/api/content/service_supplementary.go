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
	"errors"
	"log/slog"
)

func (svc *service) AddMetadataSupplementary(ctx context.Context, request *grpc.AddSupplementaryRequest) (*grpc.MetadataSupplementary, error) {
	return nil, errors.New("TODO")
}

func (svc *service) GetMetadataSupplementaryUploadUrl(ctx context.Context, request *protobuf.SupplementaryIdRequest) (*grpc.SignedUrl, error) {
	id := request.Id + "." + request.Key
	return svc.objectStore.CreateDownloadUrl(ctx, id)
}

func (svc *service) GetMetadataSupplementaryDownloadUrl(ctx context.Context, request *protobuf.SupplementaryIdRequest) (*grpc.SignedUrl, error) {
	id := request.Id + "." + request.Key
	return svc.objectStore.CreateDownloadUrl(ctx, id)
}

func (svc *service) DeleteMetadataSupplementary(ctx context.Context, request *protobuf.SupplementaryIdRequest) (*protobuf.Empty, error) {
	id := request.Id + "." + request.Key
	err := svc.objectStore.Delete(ctx, id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to delete supplementary file", slog.String("id", request.Id), slog.Any("error", err))
		return nil, err
	}
	return &protobuf.Empty{}, nil
}
