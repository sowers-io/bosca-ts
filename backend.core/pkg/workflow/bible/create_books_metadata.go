package bible

import (
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workflow/registry"
	"context"
	"errors"
)

func init() {
	registry.RegisterActivity("bible.books.create", createBooksMetadata)
}

func createBooksMetadata(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	return errors.New("TODO")
}
