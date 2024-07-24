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

package health

import (
	"context"
	grpc "google.golang.org/grpc/health/grpc_health_v1"
)

type healthService struct {
	grpc.UnimplementedHealthServer
}

func NewHealthService() grpc.HealthServer {
	return &healthService{}
}

func (svc *healthService) Check(context.Context, *grpc.HealthCheckRequest) (*grpc.HealthCheckResponse, error) {
	return &grpc.HealthCheckResponse{
		Status: grpc.HealthCheckResponse_SERVING,
	}, nil
}

func (svc *healthService) Watch(request *grpc.HealthCheckRequest, server grpc.Health_WatchServer) error {
	c, err := svc.Check(context.Background(), nil)
	if err != nil {
		return err
	}
	return server.SendMsg(c)
}
