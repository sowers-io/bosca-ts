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

package util

import (
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/clients"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/search/factory"
	"bosca.io/pkg/search/qdrant"
	"bosca.io/pkg/temporal"
	"bosca.io/pkg/util"
	"bosca.io/pkg/workers/common"
	rootContext "context"
	"github.com/qdrant/go-client/qdrant"
	"go.temporal.io/sdk/client"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func NewAITemporalClient() (client.Client, error) {
	ctx := rootContext.Background()

	cfg := configuration.NewWorkerConfiguration()
	connection, err := clients.NewClientConnection(cfg.ClientEndPoints.ContentApiAddress)
	if err != nil {
		return nil, err
	}
	contentService := content.NewContentServiceClient(connection)

	searchClient, err := factory.NewSearch(cfg.Search)
	if err != nil {
		return nil, err
	}

	qdrantClient, err := qdrant.NewQdrantClient(cfg.ClientEndPoints.QdrantApiAddress)
	if err != nil {
		return nil, err
	}

	_, err = qdrantClient.CollectionsClient.Get(ctx, &go_client.GetCollectionInfoRequest{
		CollectionName: qdrant.MetadataIndex,
	})
	if err != nil {
		if s, ok := status.FromError(err); ok {
			if s.Code() == codes.NotFound {
				collection := &go_client.CreateCollection{
					CollectionName: qdrant.MetadataIndex,
					VectorsConfig: &go_client.VectorsConfig{
						Config: &go_client.VectorsConfig_Params{
							Params: &go_client.VectorParams{
								Size:     uint64(cfg.AIConfiguration.OllamaVectorSize),
								Distance: go_client.Distance_Cosine,
							},
						},
					},
				}
				result, err := qdrantClient.CollectionsClient.Create(ctx, collection)
				if err != nil {
					return nil, err
				}
				if !result.Result {
					return nil, err
				}
			}
		} else {
			return nil, err
		}
	}

	httpClient := util.NewDefaultHttpClient()
	propagator := common.NewContextPropagator(cfg, httpClient, contentService, searchClient, qdrantClient)
	return temporal.NewClientWithPropagator(ctx, cfg.ClientEndPoints, propagator)
}
