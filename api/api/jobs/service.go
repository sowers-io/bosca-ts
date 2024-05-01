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

package jobs

import (
	grpc "bosca.io/api/protobuf/jobs"
	"context"
	"github.com/craigpastro/pgmq-go"
	"strconv"
)

type service struct {
	grpc.UnimplementedJobsServiceServer
	queue *pgmq.PGMQ
}

func NewService(queue *pgmq.PGMQ) grpc.JobsServiceServer {
	return &service{
		queue: queue,
	}
}

func (svc *service) AddJobToQueue(ctx context.Context, request *grpc.JobQueueRequest) (*grpc.JobResponse, error) {
	id, err := svc.queue.Send(ctx, request.Queue, request.Json)
	if err != nil {
		return nil, err
	}
	return &grpc.JobResponse{
		Id: strconv.FormatInt(id, 10),
	}, nil
}
