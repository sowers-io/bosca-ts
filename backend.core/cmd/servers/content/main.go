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
	"bosca.io/api/content"
	protocontent "bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/datastore"
	"bosca.io/pkg/objectstore/factory"
	"bosca.io/pkg/security/spicedb"
	"bosca.io/pkg/server"
	"bosca.io/pkg/temporal"
	"bosca.io/pkg/util"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"log/slog"
	"os"
)

func main() {
	cfg := configuration.NewServerConfiguration("content", 5003, 5013)

	util.InitializeLogging(cfg)

	pool, err := datastore.NewDatabasePool(context.Background(), cfg)
	if err != nil {
		slog.Error("failed to connect to database", slog.Any("error", err))
		os.Exit(1)
	}

	objectStore, err := factory.NewObjectStore(cfg.Storage)
	if err != nil {
		slog.Error("failed to create object store", slog.Any("error", err))
		os.Exit(1)
	}

	ds := content.NewDataStore(stdlib.OpenDBFromPool(pool))
	permissions := spicedb.NewPermissionManager(spicedb.NewSpiceDBClient(cfg))

	temporalClient, err := temporal.NewClient(context.Background(), cfg.ClientEndPoints)
	if err != nil {
		slog.Error("failed to create temporal client", slog.Any("error", err))
		os.Exit(1)
	}

	svc := content.NewAuthorizationService(permissions, ds, content.NewService(cfg, ds, cfg.Security.ServiceAccountId, objectStore, permissions, temporalClient))
	err = server.StartServer(cfg, func(ctx context.Context, grpcSvr *grpc.Server, restSvr *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
		protocontent.RegisterContentServiceServer(grpcSvr, svc)
		err := protocontent.RegisterContentServiceHandlerFromEndpoint(ctx, restSvr, endpoint, opts)
		if err != nil {
			slog.Error("failed to register content", slog.Any("error", err))
			return err
		}
		return err
	})
	if err != nil {
		slog.Error("failed start server", slog.Any("error", err))
		os.Exit(1)
	}
}
