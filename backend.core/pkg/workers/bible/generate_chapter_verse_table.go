package bible

import (
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/bible/usx"
	"bosca.io/pkg/workers/common"
	"context"
	"errors"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/md"
	"os"
	"strconv"
)

func CreateSupplementaryVerseMarkdownTable(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	contentService := common.GetContentService(ctx)

	relationships, err := contentService.GetMetadataRelationships(common.GetServiceAuthorizedContext(ctx), &content.MetadataRelationshipIdRequest{
		Id:           executionContext.Metadata.Id,
		Relationship: "usx-bible",
	})
	if err != nil {
		return err
	}
	if len(relationships.Relationships) != 1 {
		return errors.New("expected only one relationship, got " + strconv.Itoa(len(relationships.Relationships)))
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

	for _, book := range bundle.Books() {
		for _, chapter := range book.Chapters {
			if chapter.GetUsfm() != executionContext.Metadata.Name {
				continue
			}

			ast.AppendChild(document, &ast.Text{Leaf: ast.Leaf{Literal: []byte("Chapter USFM: " + chapter.GetUsfm())}})

			table := &ast.Table{}
			header := &ast.TableHeader{}
			headerRow := &ast.TableRow{}

			cell := &ast.TableCell{}
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
	supplementaryId := executionContext.Activity.Outputs["supplementaryId"]
	return common.SetSupplementaryContent(ctx, executionContext, supplementaryId, "text/markdown", []byte(output))
}
