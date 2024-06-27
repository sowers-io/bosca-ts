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
	"bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/bible/usx"
	"bosca.io/pkg/workflow/common"
	"bosca.io/pkg/workflow/registry"
	"bytes"
	"context"
	"log/slog"
	"os"
	"time"
)

func init() {
	registry.RegisterActivity("bible.process", processBible)
}

func processBible(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	metadataFile, err := common.DownloadTemporaryMetadataFile(ctx, executionContext.Metadata.Id)
	if err != nil {
		slog.ErrorContext(ctx, "failed to download file", slog.Any("error", err))
		return err
	}
	defer metadataFile.Close()
	defer os.Remove(metadataFile.Name())

	bundle, err := usx.OpenBundle(metadataFile)
	if err != nil {
		slog.ErrorContext(ctx, "failed to open bundle", slog.Any("error", err))
		return err
	}

	svc := common.GetContentService(ctx)
	ctx = common.GetServiceAuthorizedContext(ctx)

	source, err := svc.GetSource(ctx, &bosca.IdRequest{
		Id: "workflow",
	})
	if err != nil {
		slog.ErrorContext(ctx, "failed to get source", slog.Any("error", err))
		return err
	}

	addBookRequests := make([]*content.AddCollectionRequest, 0)
	addChapterRequests := make(map[string][]*content.AddMetadataRequest)

	contents := make(map[string]string)

	bibleCollection, err := svc.AddCollection(ctx, &content.AddCollectionRequest{
		Collection: &content.Collection{
			Name: bundle.Metadata().Identification.NameLocal,
			Type: content.CollectionType_standard,
			Attributes: map[string]string{
				"bible.type": "bible",
			},
		},
	})
	if err != nil && err.Error() != "rpc error: code = Unknown desc = name must be unique" {
		slog.ErrorContext(ctx, "failed to add bible collections", slog.Any("error", err))
		return err
	}

	var newBookRequest = func(book *usx.USX) {
		request := &content.AddCollectionRequest{
			Parent: bibleCollection.Id,
			Collection: &content.Collection{
				Name: book.BookIdentification.Code.ToString(),
				Type: content.CollectionType_standard,
				Attributes: map[string]string{
					"bible.book.usfm": book.GetUsfm(),
					"bible.type":      "book",
				},
			},
		}
		addBookRequests = append(addBookRequests, request)
	}

	var newChapterRequest = func(book *usx.USX, chapter *usx.Chapter) {
		if chapter.GetUsfm() == "" {
			return
		}
		text := &bytes.Buffer{}
		for _, verse := range chapter.FindVerses() {
			text.WriteString(verse.GetText())
		}
		contentLength := int64(len(text.String()))
		request := &content.AddMetadataRequest{
			Metadata: &content.Metadata{
				Name:          chapter.GetUsfm(),
				SourceId:      &source.Id,
				TraitIds:      []string{"bible.usx.chapter"},
				ContentType:   "text/plain",
				ContentLength: &contentLength,
				Attributes: map[string]string{
					"bible.chapter.usfm": chapter.GetUsfm(),
					"bible.book.usfm":    book.GetUsfm(),
					"bible.type":         "chapter",
				},
				LanguageTag: bundle.Metadata().Language.Iso,
			},
		}
		contents[chapter.GetUsfm()] = text.String()
		if addChapterRequests[book.GetUsfm()] == nil {
			addChapterRequests[book.GetUsfm()] = make([]*content.AddMetadataRequest, 0)
		}
		addChapterRequests[book.GetUsfm()] = append(addChapterRequests[book.GetUsfm()], request)
	}

	for _, book := range bundle.Books() {
		newBookRequest(book)
		for _, chapter := range book.Chapters {
			if chapter.Number == "" {
				continue
			}
			newChapterRequest(book, chapter)
		}
	}

	bookCollectionIds, err := svc.AddCollections(ctx, &content.AddCollectionsRequest{
		Collections: addBookRequests,
	})
	if err != nil && err.Error() != "rpc error: code = Unknown desc = name must be unique" {
		slog.ErrorContext(ctx, "failed to add book collections", slog.Any("error", err))
		return err
	}

	for _, bookId := range bookCollectionIds.Id {
		_, err = svc.AddCollectionItem(ctx, &content.AddCollectionItemRequest{
			CollectionId: bibleCollection.Id,
			ItemId: &content.AddCollectionItemRequest_ChildCollectionId{
				ChildCollectionId: bookId.Id,
			},
		})
		if err != nil && err.Error() != "rpc error: code = Unknown desc = ERROR: duplicate key value violates unique constraint \"collection_collection_items_pkey\" (SQLSTATE 23505)" {
			slog.ErrorContext(ctx, "failed to add book to bible collection", slog.Any("error", err))
			return err
		}
	}

	for ix, book := range bundle.Books() {
		chapterIds, err := svc.AddMetadatas(ctx, &content.AddMetadatasRequest{
			Metadatas: addChapterRequests[book.GetUsfm()],
		})
		if err != nil {
			slog.ErrorContext(ctx, "failed to add metadata", slog.Any("error", err))
			return err
		}
		c := make(map[string]string)
		for i, request := range addChapterRequests[book.GetUsfm()] {
			c[chapterIds.Id[i].Id] = contents[request.Metadata.Attributes["bible.chapter.usfm"]]
		}
		for tries := 0; tries < 10; tries++ {
			err = addRelationships(ctx, executionContext.Metadata.Id, chapterIds.Id, c, svc)
			if err == nil {
				break
			} else {
				time.Sleep(1 * time.Second)
			}
		}
		if err != nil {
			slog.ErrorContext(ctx, "failed to add relationship", slog.Any("error", err))
			return err
		}
		for _, chapterId := range chapterIds.Id {
			_, err = svc.AddCollectionItem(ctx, &content.AddCollectionItemRequest{
				CollectionId: bookCollectionIds.Id[ix].Id,
				ItemId: &content.AddCollectionItemRequest_ChildMetadataId{
					ChildMetadataId: chapterId.Id,
				},
			})
			if err != nil {
				slog.ErrorContext(ctx, "failed to add chapter to book", slog.Any("error", err))
				return err
			}
		}
	}
	return nil
}

func addRelationships(ctx context.Context, bibleId string, ids []*bosca.IdResponsesId, contents map[string]string, svc content.ContentServiceClient) error {
	for _, id := range ids {
		_, err := svc.AddMetadataRelationship(ctx, &content.MetadataRelationship{
			MetadataId1:  bibleId,
			MetadataId2:  id.Id,
			Relationship: "bible.usx.chapter",
		})
		if err != nil {
			slog.ErrorContext(ctx, "failed to add relationship", slog.Any("error", err))
			return err
		}
		_, err = svc.AddMetadataRelationship(ctx, &content.MetadataRelationship{
			MetadataId1:  id.Id,
			MetadataId2:  bibleId,
			Relationship: "bible.usx.bible",
		})
		if err != nil {
			slog.ErrorContext(ctx, "failed to add relationship", slog.Any("error", err))
			return err
		}
		for tries := 0; tries < 10; tries++ {
			err = common.SetContent(ctx, id.Id, []byte(contents[id.Id]))
			if err == nil {
				break
			}
			time.Sleep(1 * time.Second)
		}
		if err != nil {
			slog.ErrorContext(ctx, "failed to add content", slog.Any("error", err))
			return err
		}
	}
	return nil
}
