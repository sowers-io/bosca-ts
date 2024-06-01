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
	"bosca.io/api/ai"
	protoai "bosca.io/api/protobuf/ai"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/search/qdrant"
	"bosca.io/pkg/security/spicedb"
	"bosca.io/pkg/server"
	"bosca.io/pkg/util"
	"context"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"google.golang.org/grpc"
	"log/slog"
	"os"
)

func main() {
	cfg := configuration.NewServerConfiguration("ai", 5007, 5017)

	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	permissions := spicedb.NewPermissionManager(spicedb.NewSpiceDBClient(cfg))

	qdrantClient, err := qdrant.NewQdrantClient(cfg.ClientEndPoints.QdrantApiAddress)
	if err != nil {
		slog.Error("failed to create qdrant client", slog.Any("error", err))
		os.Exit(1)
	}

	httpClient := util.NewDefaultHttpClient()

	llm, err := ollama.New(
		ollama.WithHTTPClient(httpClient),
		ollama.WithServerURL(cfg.ClientEndPoints.OllamaApiAddress),
		ollama.WithModel(cfg.AIConfiguration.DefaultLlmModel),
	)
	if err != nil {
		slog.Error("failed to create model", slog.Any("error", err))
		os.Exit(1)
	}

	embedder, err := embeddings.NewEmbedder(llm)
	if err != nil {
		slog.Error("failed to create embedder", slog.Any("error", err))
		os.Exit(1)
	}

	store := qdrant.NewVectorStore(qdrantClient, embedder)

	svc := ai.NewAuthorizationService(permissions, ai.NewService(cfg.Security.ServiceAccountId, permissions, store, llm))
	err = server.StartServer(cfg, func(ctx context.Context, grpcSvr *grpc.Server, restSvr *runtime.ServeMux, endpoint string, opts []grpc.DialOption) error {
		protoai.RegisterAIServiceServer(grpcSvr, svc)
		err := protoai.RegisterAIServiceHandlerFromEndpoint(ctx, restSvr, endpoint, opts)
		if err != nil {
			slog.Error("failed to register ai endpoint", slog.Any("error", err))
			return err
		}
		return nil
	})
	if err != nil {
		slog.Error("failed to start server", slog.Any("error", err))
		os.Exit(1)
	}
}
