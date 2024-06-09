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

package processor

import (
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/search/qdrant"
	"bosca.io/pkg/workers/common"
	"bosca.io/pkg/workers/textextractor"
	"context"
	"errors"
	pb "github.com/qdrant/go-client/qdrant"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"os"
)

func Vectorize(ctx context.Context, metadata *content.Metadata) error {
	cfg := common.GetConfiguration(ctx)
	llm, err := ollama.New(
		ollama.WithHTTPClient(common.GetHttpClient(ctx)),
		ollama.WithServerURL(cfg.ClientEndPoints.OllamaApiAddress),
		ollama.WithModel(cfg.AIConfiguration.OllamaLlmModel),
	)
	if err != nil {
		return err
	}

	embedder, err := embeddings.NewEmbedder(llm)
	if err != nil {
		return err
	}

	svc := common.GetContentService(ctx)

	signedUrl, err := svc.GetMetadataSupplementaryDownloadUrl(common.GetServiceAuthorizedContext(ctx), &content.SupplementaryIdRequest{
		Id:   metadata.Id,
		Type: textextractor.SupplementalTextType,
	})
	if err != nil {
		return err
	}

	file, err := common.DownloadTemporaryFile(ctx, signedUrl)
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	body, err := os.ReadFile(file.Name())
	if err != nil {
		return err
	}

	md := make(map[string]any)
	for k, v := range metadata.Attributes {
		md[k] = v
	}

	texts := make([]string, 1)
	texts[0] = string(body)

	vectors, err := embedder.EmbedDocuments(ctx, texts)
	if err != nil {
		return err
	}

	if len(vectors) != 1 {
		return errors.New("number of vectors from embedder does not match number of documents")
	}

	client := common.GetQdrantClient(ctx)
	wait := true
	points := &pb.UpsertPoints{
		CollectionName: qdrant.MetadataIndex,
		Points: []*pb.PointStruct{
			{
				Id: &pb.PointId{
					PointIdOptions: &pb.PointId_Uuid{Uuid: metadata.Id},
				},
				Payload: map[string]*pb.Value{
					qdrant.ContentPayload: {
						Kind: &pb.Value_StringValue{
							StringValue: texts[0],
						},
					},
				},
				Vectors: &pb.Vectors{
					VectorsOptions: &pb.Vectors_Vector{
						Vector: &pb.Vector{
							Data: vectors[0],
						},
					},
				},
			},
		},
		Wait: &wait,
	}
	result, err := client.Upsert(ctx, points)
	if err != nil {
		return err
	}
	if result.Result.Status != pb.UpdateStatus_Completed {
		return errors.New("status not complete")
	}
	return nil
}
