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
	oath "bosca.io/pkg/identity/middleware"
	"context"
	"fmt"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
	"net"
	"net/http"
)

func StartServer(cfg *configuration.ServerConfiguration, register func(context.Context, *grpc.Server, *runtime.ServeMux, string, []grpc.DialOption)) {
	ctx := context.Background()
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()

	endpoint := fmt.Sprintf("0.0.0.0:%d", cfg.GrpcPort)
	listen, err := net.Listen("tcp", endpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	oathmw, err := oath.New(ctx, oath.WithConfigFile(cfg.OathKeeperConfiguration))
	if err != nil {
		log.Fatalf("failed to create oath middleware: %v", err)
	}

	grpcOpts := []grpc.ServerOption{
		grpc.UnaryInterceptor(oathmw.UnaryInterceptor()),
		grpc.StreamInterceptor(oathmw.StreamInterceptor()),
	}
	server := grpc.NewServer(grpcOpts...)
	mux := runtime.NewServeMux()

	register(ctx, server, mux, endpoint, []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})

	go func() {
		authentication := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			print(r.Header.Get("Authorization"))
			mux.ServeHTTP(w, r)
		})
		err = http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", cfg.RestPort), authentication)
		if err != nil {
			log.Fatalf("failed to start HTTP server: %v", err)
		}
	}()

	err = server.Serve(listen)
	if err != nil {
		log.Fatalf("failed to start serving: %v", err)
	}
}
