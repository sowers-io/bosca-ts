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

package processor

import (
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/search"
	"bosca.io/pkg/workers/common"
	"bosca.io/pkg/workers/textextractor"
	"context"
	"os"
)

func SetMetadataStatusReady(ctx context.Context, metadata *content.Metadata) error {
	contentService := common.GetContentService(ctx)
	_, err := contentService.SetMetadataStatus(common.GetServiceAuthorizedContext(ctx), &content.SetMetadataStatusRequest{
		Id:     metadata.Id,
		Status: content.MetadataStatus_ready,
	})
	return err
}

func AddToSearchIndex(ctx context.Context, metadata *content.Metadata) error {
	svc := common.GetContentService(ctx)

	signedUrl, err := svc.GetMetadataSupplementaryDownloadUrl(common.GetServiceAuthorizedContext(ctx), &content.SupplementaryIdRequest{
		Id:   metadata.Id,
		Type: textextractor.SupplementalTextType,
	})
	if err != nil {
		return err
	}

	file, err := common.DownloadTemporaryFile(ctx, signedUrl)
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	body, err := os.ReadFile(file.Name())
	if err != nil {
		return err
	}

	client := common.GetSearchClient(ctx)
	return client.Index(client.GetMetadataIndex(), &search.Document{
		Id:    metadata.Id,
		Title: metadata.Name,
		Body:  string(body),
	})
}
