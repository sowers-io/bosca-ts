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
	"context"
	"errors"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/md"
	"os"
	"strconv"
)

func init() {
	registry.RegisterActivity("bible.chapter.verses.table", createChapterVerseTable)
}

func createChapterVerseTable(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	activity := executionContext.Activities[executionContext.CurrentActivityIndex]
	contentService := common.GetContentService(ctx)

	bibleRelationship, err := contentService.GetMetadataRelationships(common.GetServiceAuthorizedContext(ctx), &content.MetadataRelationshipIdRequest{
		Id:           executionContext.Metadata.Id,
		Relationship: "usx-bible",
	})
	if err != nil {
		return err
	}
	if len(bibleRelationship.Relationships) != 1 {
		return errors.New("expected only one relationship, got " + strconv.Itoa(len(bibleRelationship.Relationships)))
	}

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

	document := &ast.Document{}

	verseRelationships, err := contentService.GetMetadataRelationships(common.GetServiceAuthorizedContext(ctx), &content.MetadataRelationshipIdRequest{
		Id:           executionContext.Metadata.Id,
		Relationship: "usx-verse",
	})
	if err != nil {
		return err
	}
	verseIds := make(map[string]string)
	for _, relationship := range verseRelationships.Relationships {
		verseIds[relationship.Attributes["usfm"]] = relationship.MetadataId2
	}

	for _, book := range bundle.Books() {
		for _, chapter := range book.Chapters {
			if chapter.GetUsfm() != executionContext.Metadata.Attributes["usfm"] {
				continue
			}

			ast.AppendChild(document, &ast.Text{Leaf: ast.Leaf{Literal: []byte("Chapter USFM: " + chapter.GetUsfm())}})

			table := &ast.Table{}
			header := &ast.TableHeader{}
			headerRow := &ast.TableRow{}

			cell := &ast.TableCell{}
			ast.AppendChild(cell, &ast.Text{Leaf: ast.Leaf{Literal: []byte("Metadata ID")}})
			ast.AppendChild(headerRow, cell)

			cell = &ast.TableCell{}
			ast.AppendChild(cell, &ast.Text{Leaf: ast.Leaf{Literal: []byte("Verse USFM")}})
			ast.AppendChild(headerRow, cell)

			cell = &ast.TableCell{}
			ast.AppendChild(cell, &ast.Text{Leaf: ast.Leaf{Literal: []byte("Verse Content")}})
			ast.AppendChild(header, cell)

			ast.AppendChild(header, headerRow)
			ast.AppendChild(table, header)

			body := &ast.TableBody{}
			for _, verse := range chapter.FindVerses() {
				row := &ast.TableRow{}

				cell = &ast.TableCell{}
				ast.AppendChild(cell, &ast.Text{Leaf: ast.Leaf{Literal: []byte(verseIds[verse.GetUsfm()])}})
				ast.AppendChild(row, cell)

				cell = &ast.TableCell{}
				ast.AppendChild(cell, &ast.Text{Leaf: ast.Leaf{Literal: []byte(verse.GetUsfm())}})
				ast.AppendChild(row, cell)

				cell = &ast.TableCell{}
				ast.AppendChild(cell, &ast.Text{Leaf: ast.Leaf{Literal: []byte(verse.GetText())}})
				ast.AppendChild(row, cell)

				ast.AppendChild(body, row)
			}
			ast.AppendChild(table, body)
			ast.AppendChild(document, table)
			break
		}
	}

	output := string(markdown.Render(document, md.NewRenderer()))
	supplementaryId := activity.Outputs["supplementaryId"]
	return common.SetSupplementaryContent(ctx, executionContext, supplementaryId, "text/markdown", []byte(output))
}
