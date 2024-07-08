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

package workflow

import (
	"bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/workflow"
	"context"
)

func (svc *service) GetModels(ctx context.Context, request *bosca.Empty) (*workflow.Models, error) {
	models, err := svc.ds.GetModels(ctx)
	if err != nil {
		return nil, err
	}
	return &workflow.Models{
		Models: models,
	}, nil
}

func (svc *service) GetModel(ctx context.Context, request *bosca.IdRequest) (*workflow.Model, error) {
	return svc.ds.GetModel(ctx, request.Id)
}
