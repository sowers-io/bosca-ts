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
	"bosca.io/pkg/clients"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/search/factory"
	"bosca.io/pkg/search/qdrant"
	"bosca.io/pkg/temporal"
	"bosca.io/pkg/util"
	"bosca.io/pkg/workers/common"
	"bosca.io/pkg/workers/textextractor"
	rootContext "context"
	"go.temporal.io/sdk/worker"
	"log/slog"
	"os"
)

func main() {

	ctx := rootContext.Background()

	cfg := configuration.NewWorkerConfiguration()
	util.InitializeLogging(cfg)

	connection, err := clients.NewClientConnection(cfg.ClientEndPoints.ContentApiAddress)
	if err != nil {
		slog.Error("failed to get content service connection", slog.Any("error", err))
		os.Exit(1)
	}

	searchClient, err := factory.NewSearch(cfg.Search)
	if err != nil {
		slog.Error("failed to get search client", slog.Any("error", err))
		os.Exit(1)
	}

	qdrantClient, err := qdrant.NewQdrantClient(cfg.ClientEndPoints.QdrantApiAddress)
	if err != nil {
		slog.Error("failed to get qdrant client", slog.Any("error", err))
		os.Exit(1)
	}

	httpClient := util.NewDefaultHttpClient()
	propagator := common.NewContextPropagator(cfg, httpClient, content.NewContentServiceClient(connection), searchClient, qdrantClient)
	client, err := temporal.NewClientWithPropagator(ctx, cfg.ClientEndPoints, propagator)
	if err != nil {
		slog.Error("Unable to create temporal client", slog.Any("error", err))
		os.Exit(1)
	}
	defer client.Close()

	err = textextractor.NewWorker(client).Run(worker.InterruptCh())
	if err != nil {
		slog.Error("error starting worker", slog.Any("error", err))
		os.Exit(1)
	}
}
