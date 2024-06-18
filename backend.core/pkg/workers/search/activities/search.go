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

package activities

import (
	"bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/search"
	"bosca.io/pkg/workers/common"
	"context"
	"os"
)

func AddToSearchIndex(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	svc := common.GetContentService(ctx)

	signedUrl, err := svc.GetMetadataDownloadUrl(common.GetServiceAuthorizedContext(ctx), &bosca.IdRequest{
		Id: executionContext.Metadata.Id,
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
		Id:   executionContext.Metadata.Id,
		Name: executionContext.Metadata.Name,
		Body: string(body),
	})
}
