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
	"bosca.io/api/protobuf/content"
	protosearch "bosca.io/api/protobuf/search"
	"bosca.io/api/search"
	"bosca.io/pkg/clients"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/search/factory"
	"bosca.io/pkg/search/qdrant"
	"bosca.io/pkg/security/spicedb"
	"bosca.io/pkg/server"
	"bosca.io/pkg/util"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tmc/langchaingo/llms/ollama"
	"google.golang.org/grpc"
	"log"
	"log/slog"
	"os"
)

func main() {
	cfg := configuration.NewServerConfiguration("search", 5005, 5015)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	connection, err := clients.NewClientConnection(cfg.ClientEndPoints.ContentApiAddress)
	if err != nil {
		log.Fatalf("failed to create content client connection: %v", err)
	}
	contentClient := content.NewContentServiceClient(connection)

	permissions := spicedb.NewPermissionManager(spicedb.NewSpiceDBClient(cfg))

	qdrantClient, err := qdrant.NewQdrantClient(cfg.ClientEndPoints.QdrantApiAddress)
	if err != nil {
		log.Fatalf("failed to create qdrant client: %v", err)
	}

	httpClient := util.NewDefaultHttpClient()

	llm, err := ollama.New(
		ollama.WithHTTPClient(httpClient),
		ollama.WithServerURL(cfg.ClientEndPoints.OllamaApiAddress),
		ollama.WithModel(cfg.AIConfiguration.DefaultLlmModel),
	)
	if err != nil {
		log.Fatalf("failed to create model: %v", err)
	}

	semanticClient, err := qdrant.NewSearchClient(qdrantClient, llm)
	if err != nil {
		log.Fatalf("failed to create semantic search client: %v", err)
	}

	standardClient, err := factory.NewSearch(cfg.Search)
	if err != nil {
		log.Fatalf("failed to create standard search client: %v", err)
	}

	svc := search.NewAuthorizationService(permissions, search.NewService(contentClient, semanticClient, standardClient))
	server.StartServer(cfg, func(ctx context.Context, grpcSvr *grpc.Server, restSvr *runtime.ServeMux, endpoint string, opts []grpc.DialOption) {
		protosearch.RegisterSearchServiceServer(grpcSvr, svc)
		err := protosearch.RegisterSearchServiceHandlerFromEndpoint(ctx, restSvr, endpoint, opts)
		if err != nil {
			log.Fatalf("failed to register search: %v", err)
		}
	})
}
