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

package qdrant

import (
	"bosca.io/pkg/search"
	"context"
	"github.com/qdrant/go-client/qdrant"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms"
)

const MetadataIndex = "metadata"
const ContentPayload = "content"

type qdrantSearch struct {
	index    search.Index
	client   *Client
	model    llms.Model
	embedder embeddings.Embedder
}

type qdrantIndex struct {
}

func NewSearchClient(client *Client, model llms.Model) (search.SemanticClient, error) {
	embedder, err := embeddings.NewEmbedder(model.(embeddings.EmbedderClient))
	if err != nil {
		return nil, err
	}
	return &qdrantSearch{
		index:    &qdrantIndex{},
		client:   client,
		model:    model,
		embedder: embedder,
	}, nil
}

func (q *qdrantIndex) Name() string {
	return MetadataIndex
}

func (q *qdrantSearch) GetMetadataIndex() search.Index {
	return q.index
}

func (q *qdrantSearch) Search(ctx context.Context, index search.Index, query *search.SemanticQuery) (*search.SemanticResult, error) {
	vectors, err := q.embedder.EmbedQuery(ctx, query.Query)
	if err != nil {
		return nil, err
	}
	searchResults, err := q.client.Search(ctx, &go_client.SearchPoints{
		CollectionName: index.Name(),
		Vector:         vectors,
		Offset:         &query.Offset,
		Limit:          query.Limit,
	})
	if err != nil {
		return nil, err
	}
	metadataIds := make([]string, len(searchResults.Result))
	for i, result := range searchResults.Result {
		metadataIds[i] = result.Id.GetUuid()
	}
	return &search.SemanticResult{
		MetadataIds: metadataIds,
	}, nil
}
