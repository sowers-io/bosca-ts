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

package bible

import (
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/bible/usx"
	"bosca.io/pkg/workflow/common"
	"bosca.io/pkg/workflow/registry"
	"bytes"
	"context"
	"os"
)

func init() {
	registry.RegisterActivity("bible.process", processBible)
}

func processBible(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	metadataFile, err := common.DownloadTemporaryMetadataFile(ctx, executionContext.Metadata.Id)
	if err != nil {
		return err
	}
	defer metadataFile.Close()
	defer os.Remove(metadataFile.Name())

	bundle, err := usx.OpenBundle(metadataFile)
	if err != nil {
		return err
	}

	svc := common.GetContentService(ctx)

	sourceId := "workflow"

	addBookRequest := &content.AddMetadataRequest{
		Metadata: &content.Metadata{
			SourceId:    &sourceId,
			TraitIds:    []string{"bible.usx.book"},
			ContentType: "text/plain",
			LanguageTag: bundle.Metadata().Language.Iso,
		},
	}

	addChapterRequest := &content.AddMetadataRequest{
		Metadata: &content.Metadata{
			SourceId:    &sourceId,
			TraitIds:    []string{"bible.usx.chapter"},
			ContentType: "text/plain",
			LanguageTag: bundle.Metadata().Language.Iso,
		},
	}

	bookCollection, err := svc.AddCollection(ctx, &content.AddCollectionRequest{
		Collection: &content.Collection{
			Name: executionContext.Metadata.Name,
			Type: content.CollectionType_standard,
			Attributes: map[string]string{
				"description": "Bible Books",
			},
		},
	})

	for _, book := range bundle.Books() {
		addBookRequest.Metadata.Name = book.BookIdentification.Text
		addBookRequest.Metadata.Attributes["bible.usfm"] = book.BookIdentification.Text
		addBookRequest.Metadata.Attributes["bible.type"] = "book"
		bookResponse, err := svc.AddMetadata(ctx, addBookRequest)
		if err != nil {
			return err
		}

		_, err = svc.AddCollectionItem(ctx, &content.AddCollectionItemRequest{
			CollectionId: bookCollection.Id,
			ItemId: &content.AddCollectionItemRequest_ChildMetadataId{
				ChildMetadataId: bookResponse.Id,
			},
		})
		if err != nil {
			return nil
		}

		chapterCollection, err := svc.AddCollection(ctx, &content.AddCollectionRequest{
			Collection: &content.Collection{
				Name: book.BookChapterLabel.Text,
				Type: content.CollectionType_standard,
				Attributes: map[string]string{
					"description": "Bible Chapters",
					"bible.usfm":  book.BookIdentification.Text,
				},
			},
		})
		if err != nil {
			return err
		}

		for _, chapter := range book.Chapters {
			text := &bytes.Buffer{}
			for _, verse := range chapter.FindVerses() {
				text.WriteString(verse.GetText())
			}

			addChapterRequest.Metadata.Name = chapter.GetUsfm()
			addChapterRequest.Metadata.ContentLength = int64(text.Len())
			addChapterRequest.Metadata.Attributes["bible.usfm"] = chapter.GetUsfm()
			addChapterRequest.Metadata.Attributes["bible.type"] = "chapter"
			addChapterRequest.Metadata.Attributes["bible.book.usfm"] = book.BookIdentification.Text
			response, err := svc.AddMetadata(ctx, addChapterRequest)
			if err != nil {
				return err
			}

			err = common.SetContent(ctx, response.Id, []byte(text.String()))
			if err != nil {
				return err
			}

			_, err = svc.AddMetadataRelationship(ctx, &content.MetadataRelationship{
				MetadataId1:  executionContext.Metadata.Id,
				MetadataId2:  response.Id,
				Relationship: "bible.usx.chapter",
				Attributes:   map[string]string{"bible.book.usfm": book.BookIdentification.Text, "usfm": chapter.GetUsfm()},
			})
			if err != nil {
				return nil
			}

			_, err = svc.AddMetadataRelationship(ctx, &content.MetadataRelationship{
				MetadataId1:  response.Id,
				MetadataId2:  executionContext.Metadata.Id,
				Relationship: "bible.usx.bible",
			})
			if err != nil {
				return nil
			}

			_, err = svc.AddCollectionItem(ctx, &content.AddCollectionItemRequest{
				CollectionId: chapterCollection.Id,
				ItemId: &content.AddCollectionItemRequest_ChildMetadataId{
					ChildMetadataId: response.Id,
				},
			})
			if err != nil {
				return nil
			}
		}
	}

	return nil
}
