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
	protosearch "bosca.io/api/protobuf/bosca/search"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/server"
	"bosca.io/pkg/util"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"log/slog"
	"os"
)

func main() {
	cfg := configuration.NewServerConfiguration("search", 5005, 5015)

	util.InitializeLogging(cfg)

	cfg.GrpcDisabled = true
	err := server.StartServer(cfg, func(ctx context.Context, grpcSvr *grpc.Server, restSvr *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
		err := protosearch.RegisterSearchServiceHandlerFromEndpoint(ctx, restSvr, endpoint, opts)
		if err != nil {
			slog.Error("failed to register search", slog.Any("error", err))
			return err
		}
		return nil
	})
	if err != nil {
		slog.Error("failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}
