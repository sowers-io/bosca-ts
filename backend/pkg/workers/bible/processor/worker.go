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
	"bosca.io/api/protobuf"
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/bible/usx"
	"bosca.io/pkg/search"
	"bosca.io/pkg/search/qdrant"
	"bosca.io/pkg/workers/common"
	"bosca.io/pkg/workers/textextractor"
	"context"
	"encoding/xml"
	"errors"
	"github.com/google/uuid"
	pb "github.com/qdrant/go-client/qdrant"
	"github.com/tmc/langchaingo/embeddings"
	"github.com/tmc/langchaingo/llms/ollama"
	"log/slog"
	"os"
)

func AddToSearchIndex(ctx context.Context, metadata *content.Metadata) error {
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

	client := common.GetSearchClient(ctx)
	return client.Index(client.GetMetadataIndex(), &search.Document{
		Id:   metadata.Id,
		Name: metadata.Name,
		Body: string(body),
	})
}

func VectorizeBible(ctx context.Context, metadata *content.Metadata) error {
	cfg := common.GetConfiguration(ctx)
	llm, err := ollama.New(
		ollama.WithHTTPClient(common.GetHttpClient(ctx)),
		ollama.WithServerURL(cfg.ClientEndPoints.OllamaApiAddress),
		ollama.WithModel(cfg.AIConfiguration.DefaultLlmModel),
	)
	if err != nil {
		return err
	}

	embedder, err := embeddings.NewEmbedder(llm)
	if err != nil {
		return err
	}

	metadataDownloadUrl, err := common.GetContentService(ctx).GetMetadataDownloadUrl(common.GetServiceAuthorizedContext(ctx), &protobuf.IdRequest{
		Id: metadata.Id,
	})
	if err != nil {
		return err
	}
	if metadataDownloadUrl == nil {
		slog.WarnContext(ctx, "metadata download url is nil, nothing to do", slog.String("metadata_id", metadata.Id))
		return nil
	}
	metadataFile, err := common.DownloadTemporaryFile(ctx, metadataDownloadUrl)
	if err != nil {
		return err
	}
	defer os.Remove(metadataFile.Name())

	uuidMappingDownloadUrl, err := common.GetContentService(ctx).GetMetadataSupplementaryDownloadUrl(common.GetServiceAuthorizedContext(ctx), &content.SupplementaryIdRequest{
		Id:   metadata.Id,
		Type: "usfm_mapping",
	})
	if err != nil {
		return err
	}
	if metadataDownloadUrl == nil {
		slog.WarnContext(ctx, "metadata download url is nil, nothing to do", slog.String("metadata_id", metadata.Id))
		return nil
	}
	uuidMappingFile, err := common.DownloadTemporaryFile(ctx, uuidMappingDownloadUrl)
	if err != nil {
		return err
	}
	defer os.Remove(uuidMappingFile.Name())

	mapping := &UUIDMap{}
	mappingData, err := os.ReadFile(uuidMappingFile.Name())
	if err != nil {
		return err
	}
	if len(mappingData) > 0 {
		if err = xml.Unmarshal(mappingData, mapping); err != nil {
			return err
		}
	}

	bundle, err := usx.OpenBundle(metadataFile)

	for _, book := range bundle.Books() {
		for _, chapter := range book.FindChapterVerses() {
			verseText := make([]string, len(chapter.Verses))
			for i, verse := range chapter.Verses {
				verseText[i] = verse.GetText()
			}
			// TODO: chunk this more
			vectors, err := embedder.EmbedDocuments(ctx, verseText)
			if err != nil {
				return err
			}
			if err = addToVectorDatabase(ctx, metadata, mapping, bundle.Metadata().Identification.SystemId[0].ID, chapter.Verses, &vectors); err != nil {
				return err
			}
		}
	}

	// TODO: upload mappings file

	return nil
}

type UUIDMap struct {
	Mapping map[string]string `json:"mapping"`
}

func addToVectorDatabase(ctx context.Context, metadata *content.Metadata, mapping *UUIDMap, translationId string, verses []*usx.Verse, vectors *[][]float32) error {
	documents := make([]*pb.PointStruct, len(verses))
	for i, verse := range verses {
		usfm := verse.GetUsfm()
		id := mapping.Mapping[usfm]
		if id == "" {
			id = uuid.New().String()
			mapping.Mapping[usfm] = id
		}
		documents[i] = &pb.PointStruct{
			Id: &pb.PointId{
				PointIdOptions: &pb.PointId_Uuid{Uuid: id},
			},
			Payload: map[string]*pb.Value{
				"translation": {Kind: &pb.Value_StringValue{StringValue: translationId}},
				"usfm":        {Kind: &pb.Value_StringValue{StringValue: verse.GetUsfm()}},
				"metadata":    {Kind: &pb.Value_StringValue{StringValue: metadata.Id}},
			},
			Vectors: &pb.Vectors{
				VectorsOptions: &pb.Vectors_Vector{
					Vector: &pb.Vector{
						Data: (*vectors)[i],
					},
				},
			},
		}
	}
	client := common.GetQdrantClient(ctx)
	wait := true
	points := &pb.UpsertPoints{
		CollectionName: qdrant.MetadataIndex,
		Points:         documents,
		Wait:           &wait,
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
