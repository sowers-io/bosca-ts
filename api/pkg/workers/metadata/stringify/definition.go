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
	"bosca.io/api/protobuf"
	"bosca.io/api/protobuf/content"
	"context"
	"go.temporal.io/sdk/workflow"
)

func GetMetadata(ctx workflow.Context, id string) (*content.Metadata, error) {
	return svc.GetMetadata(context.Background(), &protobuf.IdRequest{
		Id: id,
	})
}

func SetMetadataStatusReady(ctx workflow.Context, metadata *content.Metadata) error {
	_, err := svc.SetMetadataStatus(context.Background(), &content.SetMetadataStatusRequest{
		Id:     metadata.Id,
		Status: content.MetadataStatus_ready,
	})
	return err
}

func Stringify(ctx workflow.Context, content *content.Metadata) error {
	return nil
}
