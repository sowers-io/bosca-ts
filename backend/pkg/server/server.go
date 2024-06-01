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
	securityFactory "bosca.io/pkg/security/factory"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"log/slog"
	"net"
	"net/http"
)

func StartServer(cfg *configuration.ServerConfiguration, register func(context.Context, *grpc.Server, *runtime.ServeMux, string, []grpc.DialOption) error) error {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	endpoint := fmt.Sprintf("0.0.0.0:%d", cfg.GrpcPort)
	listen, err := net.Listen("tcp", endpoint)
	if err != nil {
		slog.Error("failed to listen", slog.Int("grpc_port", cfg.GrpcPort), slog.Any("error", err))
		return nil
	}

	sessionInterceptor := securityFactory.NewSessionInterceptor(cfg.Security.SessionEndpointType)
	subjectFinder := security.NewSubjectFinder(cfg.Security.SessionEndpoint, cfg.Security.ServiceAccountId, cfg.Security.ServiceAccountToken, sessionInterceptor)
	interceptors := security.NewSecurityInterceptors(subjectFinder)

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
		slog.Error("failed to register health endpoints", slog.Any("error", err))
		return nil
	}

	go func() {
		err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", cfg.RestPort), mux)
		if err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
	}()

	slog.Info("running server", slog.Int("grpc_port", cfg.GrpcPort), slog.Int("rest_port", cfg.RestPort))

	err = server.Serve(listen)
	if err != nil {
		slog.Error("failed to start serving", slog.Any("error", err))
		return err
	}
	return nil
}
