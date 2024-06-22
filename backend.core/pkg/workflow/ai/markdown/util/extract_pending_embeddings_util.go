package util

import (
	search "bosca.io/api/protobuf/bosca/ai"
	"errors"
	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
)

func ExtractPendingEmbeddingsFromMarkdown(body []byte, idColumn, contentColumn string) ([]*search.PendingEmbedding, error) {
	pending := make([]*search.PendingEmbedding, 0)
	for _, node := range markdown.Parse(body, parser.New()).GetChildren() {
		var table *ast.Table
		var ok bool
		if table, ok = node.(*ast.Table); !ok {
			continue
		}

		index := -1
		idColumnIndex := -1
		contentColumnIndex := -1
		children := table.AsContainer().GetChildren()
		for _, child := range children {
			if hdr, ok := child.(*ast.TableHeader); ok {
				index = 0
				row := hdr.Children[0].(*ast.TableRow)
				for _, c := range row.Children {
					container := c.AsContainer().Children[0]
					strBytes := container.AsLeaf().Literal
					if string(strBytes) == idColumn {
						idColumnIndex = index
					}
					if string(strBytes) == contentColumn {
						contentColumnIndex = index
					}
					index++
				}
			} else if body, ok := child.(*ast.TableBody); ok {
				if idColumnIndex == -1 {
					return nil, errors.New("failed to find header column")
				}
				for _, row := range body.GetChildren() {
					index = 0
					idValue := ""
					columnValue := ""
					for _, cell := range row.AsContainer().GetChildren() {
						container := cell.AsContainer().GetChildren()[0]
						strBytes := container.AsLeaf().Literal
						if index == idColumnIndex {
							idValue = string(strBytes)
						}
						if index == contentColumnIndex {
							columnValue = string(strBytes)
						}
						index++
					}
					pending = append(pending, &search.PendingEmbedding{
						Id:      idValue,
						Content: &columnValue,
					})
				}
			}
		}
	}
	return pending, nil
}
