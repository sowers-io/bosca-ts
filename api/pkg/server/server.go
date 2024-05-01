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

package server

import (
	"bosca.io/api/health"
	protohealth "bosca.io/api/protobuf/health"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/security"
	"bosca.io/pkg/security/ory"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func NewSessionInterceptor(interceptorType string) security.SessionInterceptor {
	switch interceptorType {
	case "ory":
		return ory.NewSessionInterceptor()
	default:
		log.Fatalf("failed to find interceptor type %s", interceptorType)
		return nil
	}
}

func StartServer(cfg *configuration.ServerConfiguration, register func(context.Context, *grpc.Server, *runtime.ServeMux, string, []grpc.DialOption)) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	endpoint := fmt.Sprintf("0.0.0.0:%d", cfg.GrpcPort)
	listen, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	interceptor := NewSessionInterceptor(cfg.Security.SessionEndpointType)
	interceptors := security.NewSecurityInterceptors(cfg.Security.SessionEndpoint, interceptor)

	grpcOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(interceptors.UnaryInterceptor()),
		grpc.StreamInterceptor(interceptors.StreamInterceptor()),
	}
	server := grpc.NewServer(grpcOpts...)
	mux := runtime.NewServeMux()

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	register(ctx, server, mux, endpoint, opts)

	protohealth.RegisterHealthServiceServer(server, health.NewHealthService())
	err = protohealth.RegisterHealthServiceHandlerFromEndpoint(ctx, mux, endpoint, opts)
	if err != nil {
		log.Fatalf("failed to register health endpoints: %v", err)
	}

	go func() {
		err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", cfg.RestPort), mux)
		if err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
	}()

	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("failed to start serving: %v", err)
	}
}
