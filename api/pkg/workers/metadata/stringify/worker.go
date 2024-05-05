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

package stringify

import (
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/workers/common"
	"context"
)

func SetMetadataStatusReady(ctx context.Context, metadata *content.Metadata) error {
	contentService := common.GetContentService(ctx)
	_, err := contentService.SetMetadataStatus(common.GetServiceAuthorizedContext(ctx), &content.SetMetadataStatusRequest{
		Id:     metadata.Id,
		Status: content.MetadataStatus_ready,
	})
	return err
}

func Stringify(ctx context.Context, content *content.Metadata) error {
	return nil
}
