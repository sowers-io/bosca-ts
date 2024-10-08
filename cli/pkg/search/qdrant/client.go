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
	"bosca.io/pkg/clients"
	"google.golang.org/grpc"

	"context"
	"errors"

	"github.com/google/uuid"
	go_client "github.com/qdrant/go-client/qdrant"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/schema"
	"github.com/tmc/langchaingo/vectorstores"
)

type Client struct {
	go_client.PointsClient
	go_client.CollectionsClient

	qdrantConnection *grpc.ClientConn
}

func NewQdrantClient(endpoint string) (*Client, error) {
	qdrantConnection, err := clients.NewClientConnection(endpoint)
	if err != nil {
		return nil, err
	}
	qdrantPointsClient := go_client.NewPointsClient(qdrantConnection)
	qdrantCollectionsClient := go_client.NewCollectionsClient(qdrantConnection)
	client := &Client{
		PointsClient:      qdrantPointsClient,
		CollectionsClient: qdrantCollectionsClient,
		qdrantConnection:  qdrantConnection,
	}
	return client, nil
}

func (client *Client) Close() error {
	return client.qdrantConnection.Close()
}

type vectorStore struct {
	client   *Client
	embedder embeddings.Embedder
}

func NewVectorStore(client *Client, embedder embeddings.Embedder) vectorstores.VectorStore {
	return &vectorStore{
		client:   client,
		embedder: embedder,
	}
}

func (v *vectorStore) AddDocuments(ctx context.Context, docs []schema.Document, options ...vectorstores.Option) ([]string, error) {
	texts := make([]string, 0, len(docs))
	for _, doc := range docs {
		texts = append(texts, doc.PageContent)
	}

	vectors, err := v.embedder.EmbedDocuments(ctx, texts)
	if err != nil {
		return nil, err
	}

	if len(vectors) != len(docs) {
		return nil, errors.New("number of vectors from embedder does not match number of documents")
	}

	wait := true
	points := &go_client.UpsertPoints{
		CollectionName: MetadataIndex,
		Points:         make([]*go_client.PointStruct, 0, len(docs)),
		Wait:           &wait,
	}

	ids := make([]string, 0, len(docs))
	for i, doc := range docs {
		id := uuid.New().String()
		ids[i] = id
		payload := make(map[string]*go_client.Value)
		for k, m := range doc.Metadata {
			if m == nil {
				continue
			}
			v := go_client.Value{}
			switch m.(type) {
			case int64:
				v.Kind = &go_client.Value_IntegerValue{IntegerValue: m.(int64)}
			case float64:
				v.Kind = &go_client.Value_DoubleValue{DoubleValue: m.(float64)}
			case bool:
				v.Kind = &go_client.Value_BoolValue{BoolValue: m.(bool)}
			case string:
				v.Kind = &go_client.Value_StringValue{StringValue: m.(string)}
			}
			payload[k] = &v
		}
		points.Points[i] = &go_client.PointStruct{
			Id: &go_client.PointId{
				PointIdOptions: &go_client.PointId_Uuid{Uuid: id},
			},
			Payload: payload,
			Vectors: &go_client.Vectors{
				VectorsOptions: &go_client.Vectors_Vector{
					Vector: &go_client.Vector{
						Data: vectors[i],
					},
				},
			},
		}
	}

	result, err := v.client.Upsert(ctx, points)
	if err != nil {
		return nil, err
	}
	if result.Result.Status != go_client.UpdateStatus_Completed {
		return nil, errors.New("status not complete")
	}

	return ids, nil
}

func (v *vectorStore) SimilaritySearch(ctx context.Context, query string, numDocuments int, options ...vectorstores.Option) ([]schema.Document, error) {
	vectors, err := v.embedder.EmbedQuery(ctx, query)
	if err != nil {
		return nil, err
	}
	var threshold *float32
	ops := &vectorstores.Options{}
	if options != nil {
		for _, op := range options {
			op(ops)
		}
	}
	if ops.ScoreThreshold > 0.0 {
		threshold = &ops.ScoreThreshold
	}
	searchResults, err := v.client.Search(ctx, &go_client.SearchPoints{
		CollectionName: MetadataIndex,
		Vector:         vectors,
		ScoreThreshold: threshold,
		WithVectors: &go_client.WithVectorsSelector{
			SelectorOptions: &go_client.WithVectorsSelector_Enable{
				Enable: true,
			},
		},
		WithPayload: &go_client.WithPayloadSelector{
			SelectorOptions: &go_client.WithPayloadSelector_Enable{
				Enable: true,
			},
		},
		Limit: uint64(numDocuments),
	})
	if err != nil {
		return nil, err
	}
	documents := make([]schema.Document, len(searchResults.Result))

	for i, result := range searchResults.Result {
		doc := schema.Document{
			Metadata: make(map[string]interface{}),
		}
		for k, v := range result.Payload {
			if k == ContentPayload {
				doc.PageContent = v.GetStringValue()
				continue
			}
			switch v.Kind.(type) {
			case *go_client.Value_IntegerValue:
				doc.Metadata[k] = v.GetIntegerValue()
			case *go_client.Value_DoubleValue:
				doc.Metadata[k] = v.GetDoubleValue()
			case *go_client.Value_BoolValue:
				doc.Metadata[k] = v.GetBoolValue()
			case *go_client.Value_StringValue:
				doc.Metadata[k] = v.GetStringValue()
			}
		}
		doc.Score = result.Score
		documents[i] = doc
	}

	return documents, nil
}
