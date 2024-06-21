package markdown

import (
	search "bosca.io/api/protobuf/bosca/ai"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workflow/common"
	"bosca.io/pkg/workflow/registry"
	"context"
	"google.golang.org/protobuf/proto"
	"io"
	"os"
)

func init() {
	registry.RegisterActivity("ai.embeddings.pending.from-table", createPendingEmbeddingsFromMarkdownTable)
}

func createPendingEmbeddingsFromMarkdownTable(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	inputSupplementaryId := executionContext.Inputs["supplementaryId"]
	file, err := common.DownloadTemporarySupplementaryFile(ctx, executionContext.Metadata.Id, inputSupplementaryId)
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	idColumn := executionContext.Inputs["idColumn"].GetSingleValue()
	contentColumn := executionContext.Inputs["contentColumn"].GetSingleValue()

	pending, err := extractPendingEmbeddingsFromMarkdown(bytes, idColumn, contentColumn)
	if err != nil {
		return err
	}

	pendings := &search.PendingEmbeddings{
		Embedding: pending,
	}

	data, err := proto.Marshal(pendings)
	if err != nil {
		return err
	}

	outputSupplementaryId := executionContext.Activity.Outputs["supplementaryId"]
	return common.SetSupplementaryContent(ctx, executionContext, outputSupplementaryId, "application/protobuf", data)
}
