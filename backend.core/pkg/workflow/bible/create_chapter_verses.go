package bible

import (
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workflow/registry"
	"context"
	"errors"
)

func init() {
	registry.RegisterActivity("bible.chapter.verses.create", createChapterVerses)
}

func createChapterVerses(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	return errors.New("TODO")
}
