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
	"bosca.io/pkg/configuration"
	"context"
	"fmt"
	oath "github.com/ory/oathkeeper/middleware"
	"google.golang.org/grpc"
	"log"
	"net"
)

func StartServer(cfg *configuration.ServerConfiguration, register func(*grpc.Server)) {
	listen, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", cfg.Port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	ctx := context.Background()

	oathmw, err := oath.New(ctx, oath.WithConfigFile(cfg.OathKeeperConfiguration))
	if err != nil {
		log.Fatalf("failed to create oath middleware: %v", err)
	}

	opts := []grpc.ServerOption{
		grpc.UnaryInterceptor(oathmw.UnaryInterceptor()),
		grpc.StreamInterceptor(oathmw.StreamInterceptor()),
	}
	server := grpc.NewServer(opts...)

	register(server)

	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("failed to start serving: %v", err)
	}
}
