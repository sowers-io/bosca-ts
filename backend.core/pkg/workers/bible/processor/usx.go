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
	content2 "bosca.io/api/content"
	protobuf "bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/bible/usx"
	"bosca.io/pkg/workers/common"
	"context"
	"google.golang.org/grpc/status"
	"log/slog"
	"os"
	"strings"
)

func ProcessUSX(ctx context.Context, traitWorkflow *content.TraitWorkflow) error {
	svc := common.GetContentService(ctx)
	ctx = common.GetServiceAuthorizedContext(ctx)

	metadataDownloadUrl, err := svc.GetMetadataDownloadUrl(ctx, &protobuf.IdRequest{
		Id: traitWorkflow.Metadata.Id,
	})
	if err != nil {
		return err
	}

	if metadataDownloadUrl == nil {
		slog.WarnContext(ctx, "workflow download url is nil, nothing to do", slog.String("metadata_id", traitWorkflow.Metadata.Id))
		return nil
	}

	metadataFile, err := common.DownloadTemporaryFile(ctx, metadataDownloadUrl)
	if err != nil {
		return err
	}
	defer os.Remove(metadataFile.Name())

	bundle, err := usx.OpenBundle(metadataFile)
	for _, book := range bundle.Books() {
		for _, chapter := range book.FindChapterVerses() {
			text := strings.Builder{}
			chapterUsfm := chapter.Chapter.GetUsfm()

			// process verses
			for _, verse := range chapter.Verses {
				txt := verse.GetText()
				text.Write([]byte(txt))
				usfm := verse.GetUsfm()

				slog.DebugContext(ctx, "processing verse", slog.String("usfm", usfm))
				response, err := svc.AddMetadata(ctx, newAddMetadataRequest(bundle, "verse", usfm, int64(len(txt)), traitWorkflow.Metadata.Id, map[string]string{"translation.chapter.usfm": chapterUsfm}))
				if err != nil {
					if statusError, ok := status.FromError(err); ok {
						if statusError.Message() == "name must be unique" {
							continue
						}
					}
					return err
				}

				_, err = svc.AddMetadataRelationship(ctx, &content.AddMetadataRelationshipRequest{
					MetadataId1:  traitWorkflow.Metadata.Id,
					MetadataId2:  response.Id,
					Relationship: "usx-verse",
				})

				err = setText(ctx, svc, response, bundle, usfm, txt)
				if err != nil {
					return err
				}
			}

			if text.Len() == 0 {
				continue
			}

			// process chapter
			slog.DebugContext(ctx, "processing chapter", slog.String("usfm", chapterUsfm))
			response, err := svc.AddMetadata(ctx, newAddMetadataRequest(bundle, "chapter", chapterUsfm, int64(text.Len()), traitWorkflow.Metadata.Id, nil))
			if err != nil {
				if statusError, ok := status.FromError(err); ok {
					if statusError.Message() == "name must be unique" {
						continue
					}
				}
				return err
			}

			_, err = svc.AddMetadataRelationship(ctx, &content.AddMetadataRelationshipRequest{
				MetadataId1:  traitWorkflow.Metadata.Id,
				MetadataId2:  response.Id,
				Relationship: "usx-chapter",
			})

			err = setText(ctx, svc, response, bundle, chapterUsfm, text.String())
			if err != nil {
				return err
			}
		}
	}

	return nil
}

func newAddMetadataRequest(bundle *usx.Bundle, typeName string, usfm string, len int64, metadataId string, attributes map[string]string) *content.AddMetadataRequest {
	request := &content.AddMetadataRequest{
		Metadata: &content.Metadata{
			Name: usfm,
			Attributes: map[string]string{
				"translation":                       bundle.Metadata().Identification.SystemId[0].ID,
				"translation.name":                  bundle.Metadata().Identification.NameLocal,
				"translation.abbreviation":          bundle.Metadata().Identification.Abbreviation,
				"translation.abbreviation.local":    bundle.Metadata().Identification.AbbreviationLocal,
				"translation." + typeName + ".usfm": usfm,
				"translation.metadata.id":           metadataId,
			},
			TraitIds:      []string{"bible." + typeName + ".text"},
			LanguageTag:   bundle.Metadata().Language.Iso,
			ContentLength: len,
			ContentType:   "text/plain",
		},
	}
	if attributes != nil {
		for k, v := range attributes {
			request.Metadata.Attributes[k] = v
		}
	}
	return request
}

func setText(ctx context.Context, svc content.ContentServiceClient, id *protobuf.IdResponse, bundle *usx.Bundle, usfm string, text string) error {
	err := common.SetTextContent(ctx, id.Id, text)
	if err != nil {
		_, err2 := svc.CompleteTransitionWorkflow(ctx, &content.CompleteTransitionWorkflowRequest{
			MetadataId: id.Id,
			Status:     "failed to set text",
			Success:    false,
		})
		if err2 != nil {
			slog.ErrorContext(ctx, "failed to update workflow status to failed", slog.Any("error", err2))
		}
		return err
	} else {
		_, err = svc.BeginTransitionWorkflow(ctx, &content.TransitionWorkflowRequest{
			MetadataId: id.Id,
			StateId:    content2.WorkflowStateProcessing,
			Status:     "chapter uploaded",
		})
		if err != nil {
			slog.ErrorContext(ctx, "failed to update workflow status to processing", slog.String("translation", bundle.Metadata().Identification.SystemId[0].ID), slog.String("usfm", usfm), slog.Any("error", err))
		}
	}
	return nil
}
