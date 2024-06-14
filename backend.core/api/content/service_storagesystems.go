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
	"bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/content"
	"context"
)

func (svc *service) GetStorageSystems(ctx context.Context, request *bosca.Empty) (*content.StorageSystems, error) {
	systems, err := svc.ds.GetStorageSystems(ctx)
	if err != nil {
		return nil, err
	}
	return &content.StorageSystems{
		Systems: systems,
	}, nil
}

func (svc *service) GetStorageSystem(ctx context.Context, request *bosca.IdRequest) (*content.StorageSystem, error) {
	system, err := svc.ds.GetStorageSystem(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return system, nil
}

func (svc *service) GetStorageSystemModels(ctx context.Context, request *bosca.IdRequest) (*content.StorageSystemModels, error) {
	models, err := svc.ds.GetStorageSystemModels(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &content.StorageSystemModels{
		Models: models,
	}, nil
}
