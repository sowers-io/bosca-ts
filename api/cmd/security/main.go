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
	protosecurity "bosca.io/api/protobuf/security"
	api "bosca.io/api/security"
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
	cfg := configuration.NewServerConfiguration("security", 5006, 5016)
	pool, err := datastore.NewDatabasePool(context.Background(), cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	ds := api.NewDataStore(stdlib.OpenDBFromPool(pool))
	permissionsClient := spicedb.NewSpiceDBClient(cfg)
	mgr := spicedb.NewPermissionManager(permissionsClient)
	svc := api.NewAuthorizationService(mgr, api.NewService(ds))
	server.StartServer(cfg, func(ctx context.Context, grpcSvr *grpc.Server, restSvr *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
		protosecurity.RegisterSecurityServiceServer(grpcSvr, svc)
		err := protosecurity.RegisterSecurityServiceHandlerFromEndpoint(ctx, restSvr, endpoint, opts)
		if err != nil {
			log.Fatalf("failed to register profiles: %v", err)
		}
	})
}
