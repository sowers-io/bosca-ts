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
	"bosca.io/api/protobuf"
	grpc "bosca.io/api/protobuf/jobs"
	"context"
	"errors"
	"fmt"
	"github.com/craigpastro/pgmq-go"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"strconv"
)

type service struct {
	grpc.UnimplementedJobsServiceServer
	notifier JobNotifier
	queue    *pgmq.PGMQ
}

func NewService(notifier JobNotifier, queue *pgmq.PGMQ) grpc.JobsServiceServer {
	return &service{
		notifier: notifier,
		queue:    queue,
	}
}

func (svc *service) Enqueue(ctx context.Context, request *grpc.QueueRequest) (*grpc.QueueResponse, error) {
	id, err := svc.queue.Send(ctx, request.Queue, request.Json)
	if err != nil {
		return nil, err
	}
	svc.notifier.Notify(ctx, request.Queue)
	return &grpc.QueueResponse{
		Id: strconv.FormatInt(id, 10),
	}, nil
}

func (svc *service) Poll(request *grpc.PollRequest, server grpc.JobsService_PollServer) error {
	ctx := context.Background()
	for {
		msg, err := svc.queue.Read(ctx, request.GetQueue(), request.Timeout)
		if err != nil {
			if errors.Is(err, pgmq.ErrNoRows) {
				svc.notifier.WaitForNotification(ctx, request.Queue)
				continue
			}
			return status.Errorf(codes.Internal, "failed to poll job: %v", err)
		}
		if msg == nil {
			return status.Errorf(codes.NotFound, "no job available")
		}
		err = server.Send(&grpc.Job{
			Id:   strconv.FormatInt(msg.MsgID, 10),
			Json: msg.Message,
		})
		if err != nil {
			return status.Errorf(codes.Internal, "failed to send job: %v", err)
		}
	}
}

func (svc *service) Finish(ctx context.Context, request *grpc.FinishRequest) (*protobuf.Empty, error) {
	id, err := strconv.ParseInt(request.Id, 10, 64)
	if err != nil {
		return nil, err
	}
	success, err := svc.queue.Delete(ctx, request.Queue, id)
	if err != nil {
		return nil, err
	}
	if !success {
		return nil, status.Error(codes.Unavailable, fmt.Sprintf("failed to finish job: %s", request.Id))
	}
	return &protobuf.Empty{}, nil
}
