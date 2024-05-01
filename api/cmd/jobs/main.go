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

package main

import (
	"bosca.io/api/jobs"
	protojobs "bosca.io/api/protobuf/jobs"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/server"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cfg := configuration.NewServerConfiguration("jobs", 5004, 5014)
	svc := jobs.NewService()
	server.StartServer(cfg, func(ctx context.Context, grpcSvr *grpc.Server, restSvr *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
		protojobs.RegisterJobsServiceServer(grpcSvr, svc)
		err := protojobs.RegisterJobsServiceHandlerFromEndpoint(ctx, restSvr, endpoint, opts)
		if err != nil {
			log.Fatalf("failed to register profiles: %v", err)
		}
	})
}
