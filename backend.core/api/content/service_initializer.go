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

package content

import (
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/search/meilisearch"
	"bosca.io/pkg/search/qdrant"
	"context"
	meilisearch2 "github.com/meilisearch/meilisearch-go"
	go_client "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"os"
	"strconv"
)

func initializeService(cfg *configuration.ServerConfiguration, dataStore *DataStore) {
	ctx := context.Background()

	if added, err := dataStore.AddRootCollection(ctx); added {
		if err != nil {
			slog.Error("error initializing root collection: %v", slog.Any("error", err))
			os.Exit(1)
		}
	} else if err != nil {
		slog.Error("failed to initialize root collection permission", slog.Any("error", err))
		os.Exit(1)
	} else {
		slog.Info("root collection already initialized")
	}

	systems, err := dataStore.GetStorageSystems(ctx)
	if err != nil {
		slog.Error("failed to get storage systems", slog.Any("error", err))
		os.Exit(1)
	}

	meilisearchClient := meilisearch.NewMeilisearchClient(cfg.Search.Endpoint, cfg.Search.ApiKey)
	qdrantClient, err := qdrant.NewQdrantClient(cfg.ClientEndPoints.QdrantApiAddress)
	if err != nil {
		slog.Error("failed to get qdrant client", slog.Any("error", err))
		os.Exit(1)
	}
	defer qdrantClient.Close()

	for _, system := range systems {
		if system.Type == content.StorageSystemType_vector_storage_system {
			initializeQdrant(ctx, qdrantClient, system)
		}
		if system.Type == content.StorageSystemType_search_storage_system {
			initializeMeilisearch(meilisearchClient, system)
		}
	}
}

func initializeQdrant(ctx context.Context, qdrantClient *qdrant.Client, system *content.StorageSystem) {
	_, err := qdrantClient.CollectionsClient.Get(ctx, &go_client.GetCollectionInfoRequest{
		CollectionName: system.Configuration["indexName"],
	})
	if err != nil {
		slog.Warn("error getting qdrant collection info, trying to create collection", slog.Any("error", err), slog.String("collectionName", system.Configuration["indexName"]))
		if s, ok := status.FromError(err); ok {
			if s.Code() == codes.NotFound {
				size, err := strconv.ParseInt(system.Configuration["vectorSize"], 0, 64)
				if err != nil {
					slog.Error("failed to parse vector size in system configuration", slog.Any("error", err))
					os.Exit(1)
				}
				collection := &go_client.CreateCollection{
					CollectionName: system.Configuration["indexName"],
					VectorsConfig: &go_client.VectorsConfig{
						Config: &go_client.VectorsConfig_Params{
							Params: &go_client.VectorParams{
								Size:     uint64(size),
								Distance: go_client.Distance_Cosine,
							},
						},
					},
				}
				result, err := qdrantClient.CollectionsClient.Create(ctx, collection)
				if err != nil {
					slog.Error("failed to create qdrant collection", slog.Any("error", err))
					os.Exit(1)
				}
				if !result.Result {
					slog.Error("failed to create qdrant collection")
					os.Exit(1)
				}
			}
		} else {
			slog.Error("failed to create qdrant collection", slog.Any("error", err))
			os.Exit(1)
		}
	} else {
		slog.Info("qdrant collection already exists", slog.Any("collectionName", system.Configuration["indexName"]))
	}
}

func initializeMeilisearch(client *meilisearch2.Client, system *content.StorageSystem) {
	ix, err := client.GetIndex(system.Configuration["indexName"])
	if ix != nil {
		slog.Info("meilisearch index already exists", slog.String("indexName", system.Configuration["indexName"]))
		return
	}
	slog.Warn("error getting meilisearch index info, trying to create index", slog.Any("error", err), slog.String("index", system.Configuration["indexName"]))
	_, err = client.CreateIndex(&meilisearch2.IndexConfig{
		Uid:        system.Configuration["indexName"],
		PrimaryKey: "id",
	})
	if err != nil {
		slog.Error("failed to create meilisearch index", slog.Any("error", err))
		os.Exit(1)
	}
}
