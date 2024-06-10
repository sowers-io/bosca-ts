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
	"bosca.io/api/protobuf"
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/bible/usx"
	"bosca.io/pkg/workers/common"
	"context"
	"google.golang.org/grpc/status"
	"log/slog"
	"os"
	"strings"
)

func ProcessUSX(ctx context.Context, metadata *content.Metadata) error {
	svc := common.GetContentService(ctx)
	ctx = common.GetServiceAuthorizedContext(ctx)

	metadataDownloadUrl, err := svc.GetMetadataDownloadUrl(ctx, &protobuf.IdRequest{
		Id: metadata.Id,
	})
	if err != nil {
		return err
	}

	if metadataDownloadUrl == nil {
		slog.WarnContext(ctx, "workflow download url is nil, nothing to do", slog.String("metadata_id", metadata.Id))
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
			for _, verse := range chapter.Verses {
				text.Write([]byte(verse.GetText()))
			}
			if text.Len() == 0 {
				continue
			}
			chapterUsfm := chapter.Chapter.GetUsfm()
			slog.DebugContext(ctx, "processing chapter", slog.String("usfm", chapterUsfm))
			response, err := svc.AddMetadata(ctx, &content.AddMetadataRequest{
				Metadata: &content.Metadata{
					Name: chapterUsfm,
					Attributes: map[string]string{
						"translation":                    bundle.Metadata().Identification.SystemId[0].ID,
						"translation.name":               bundle.Metadata().Identification.NameLocal,
						"translation.abbreviation":       bundle.Metadata().Identification.Abbreviation,
						"translation.abbreviation.local": bundle.Metadata().Identification.AbbreviationLocal,
						"translation.chapter.usfm":       chapterUsfm,
						"translation.workflow.id":        metadata.Id,
					},
					TraitIds:      []string{"bible.chapter.text"},
					LanguageTag:   bundle.Metadata().Language.Iso,
					ContentLength: int64(text.Len()),
					ContentType:   "text/plain",
				},
			})
			if err != nil {
				if statusError, ok := status.FromError(err); ok {
					if statusError.Message() == "name must be unique" {
						continue
					}
				}
				return err
			}
			_, err = svc.AddMetadataRelationship(ctx, &content.AddMetadataRelationshipRequest{
				MetadataId1:  metadata.Id,
				MetadataId2:  response.Id,
				Relationship: "usx-chapter",
			})
			err = common.SetTextContent(ctx, response.Id, text.String())
			if err != nil {
				_, err2 := svc.CompleteTransitionWorkflow(ctx, &content.CompleteTransitionWorkflowRequest{
					MetadataId: response.Id,
					Status:     "failed to set text",
					Success:    false,
				})
				if err2 != nil {
					slog.ErrorContext(ctx, "failed to update workflow status to failed", slog.Any("error", err2))
				}
				return err
			} else {
				_, err = svc.BeginTransitionWorkflow(ctx, &content.TransitionWorkflowRequest{
					MetadataId: response.Id,
					StateId:    content2.WorkflowStateProcessing,
					Status:     "chapter uploaded",
				})
				if err != nil {
					slog.ErrorContext(ctx, "failed to update workflow status to processing for chapter", slog.String("translation", bundle.Metadata().Identification.SystemId[0].ID), slog.String("usfm", chapterUsfm), slog.Any("error", err))
				}
			}
		}
	}

	return nil
}
