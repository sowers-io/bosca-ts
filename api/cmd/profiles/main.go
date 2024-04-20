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

//go:generate protoc -I ../../../protobuf --go_out=../../api/protobuf --go_opt=paths=source_relative --go-grpc_out=../../api/protobuf --go-grpc_opt=paths=source_relative ../../../protobuf/profiles/profiles.proto ../../../protobuf/empty.proto ../../../protobuf/requests.proto
package main

import (
	"bosca.io/api/profiles"
	protoprofiles "bosca.io/api/protobuf/profiles"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/datastore"
	"bosca.io/pkg/server"
	"google.golang.org/grpc"
	"log"
)

func main() {
	cfg := configuration.NewServerConfiguration()
	db, err := datastore.NewDatabase(cfg)
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
	ds := profiles.NewDataStore(db)
	svc := profiles.NewService(ds)
	server.StartServer(cfg, func(svr *grpc.Server) {
		protoprofiles.RegisterProfilesServiceServer(svr, svc)
	})
}
