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
	"bosca.io/api/protobuf/bosca/content"
	protoworkflow "bosca.io/api/protobuf/bosca/workflow"
	"bosca.io/api/workflow"
	"bosca.io/pkg/clients"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/datastore"
	"bosca.io/pkg/objectstore/factory"
	"bosca.io/pkg/security/spicedb"
	"bosca.io/pkg/server"
	"bosca.io/pkg/util"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"log/slog"
	"os"
)

func main() {
	cfg := configuration.NewServerConfiguration("workflow", 5001, 5011)

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

	pubsub := workflow.NewPubSub(cfg.ClientEndPoints)
	ds := workflow.NewDataStore(stdlib.OpenDBFromPool(pool), pubsub)
	permissions := spicedb.NewPermissionManager(spicedb.NewSpiceDBClient(cfg))

	contentConnection, err := clients.NewClientConnection(cfg.ClientEndPoints.ContentApiAddress)
	if err != nil {
		slog.Error("failed to get content connection", slog.Any("error", err))
		os.Exit(1)
	}
	contentClient := content.NewContentServiceClient(contentConnection)

	svc := workflow.NewAuthorizationService(permissions, ds, workflow.NewService(cfg, ds, cfg.Security.ServiceAccountId, cfg.Security.ServiceAccountToken, objectStore, permissions, contentClient))
	err = server.StartServer(cfg, func(ctx context.Context, grpcSvr *grpc.Server, restSvr *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
		protoworkflow.RegisterWorkflowServiceServer(grpcSvr, svc)
		err := protoworkflow.RegisterWorkflowServiceHandlerFromEndpoint(ctx, restSvr, endpoint, opts)
		if err != nil {
			slog.Error("failed to register workflow", slog.Any("error", err))
			return err
		}
		return err
	})
	if err != nil {
		slog.Error("failed start server", slog.Any("error", err))
		os.Exit(1)
	}
}
