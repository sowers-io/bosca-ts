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
	grpc "bosca.io/api/protobuf/bosca/workflow"
	"context"
)

func (svc *service) RegisterWorker(ctx context.Context, request *grpc.RegisterWorkerRequest) (*bosca.IdResponse, error) {
	id, err := svc.ds.RegisterWorker(ctx, request.Queue, request.ActivityId)
	if err != nil {
		return nil, err
	}
	return &bosca.IdResponse{Id: id}, nil
}

func (svc *service) HeartbeatWorker(ctx context.Context, request *bosca.IdRequest) (*bosca.Empty, error) {
	return &bosca.Empty{}, nil
}

func (svc *service) UnregisterWorker(ctx context.Context, request *bosca.IdRequest) (*bosca.Empty, error) {
	err := svc.ds.UnregisterWorker(ctx, request.Id)
	if err != nil {
		return nil, err
	}
	return &bosca.Empty{}, nil
}
