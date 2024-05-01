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
	"bosca.io/api/content/minio"
	"bosca.io/api/content/s3"
	protocontent "bosca.io/api/protobuf/content"
	"bosca.io/api/protobuf/jobs"
	"bosca.io/pkg/clients"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/datastore"
	"bosca.io/pkg/security/spicedb"
	"bosca.io/pkg/server"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/jackc/pgx/v5/stdlib"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cfg := configuration.NewServerConfiguration("content", 5003, 5013)
	pool, err := datastore.NewDatabasePool(context.Background(), cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	jobsConnection := clients.NewClientConnection(cfg.ClientEndPoints.JobsApiAddress)
	defer jobsConnection.Close()

	jobsClient := jobs.NewJobsServiceClient(jobsConnection)

	ds := content.NewDataStore(stdlib.OpenDBFromPool(pool))
	permissions := spicedb.NewPermissionManager(spicedb.NewSpiceDBClient(cfg))

	var os content.ObjectStore
	switch cfg.StorageType {
	case configuration.StorageTypeMinio:
		os = minio.NewMinioObjectStore(cfg)
		break
	case configuration.StorageTypeS3:
		os = s3.NewS3ObjectStore(cfg)
		break
	default:
		log.Fatalf("unknown storage type: %v", cfg.StorageType)
	}

	svc := content.NewAuthorizationService(permissions, ds, content.NewService(ds, os, permissions, jobsClient))
	server.StartServer(cfg, func(ctx context.Context, grpcSvr *grpc.Server, restSvr *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
		protocontent.RegisterContentServiceServer(grpcSvr, svc)
		err := protocontent.RegisterContentServiceHandlerFromEndpoint(ctx, restSvr, endpoint, opts)
		if err != nil {
			log.Fatalf("failed to register content: %v", err)
		}
	})
}
