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

package markdown

import (
	search "bosca.io/api/protobuf/bosca/ai"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workflow/ai/markdown/util"
	"bosca.io/pkg/workflow/common"
	"bosca.io/pkg/workflow/registry"
	"context"
	"google.golang.org/protobuf/proto"
	"io"
	"os"
)

func init() {
	registry.RegisterActivity("ai.embeddings.pending.from-markdown-table", createPendingEmbeddingsFromMarkdownTable)
}

func createPendingEmbeddingsFromMarkdownTable(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	activity := executionContext.Activities[executionContext.CurrentActivityIndex]
	inputSupplementaryId := activity.Inputs["supplementaryId"]
	file, err := common.DownloadTemporarySupplementaryFile(ctx, executionContext.Metadata.Id, inputSupplementaryId)
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	bytes, err := io.ReadAll(file)
	if err != nil {
		return err
	}

	idColumn := activity.Inputs["idColumn"].GetSingleValue()
	contentColumn := activity.Inputs["contentColumn"].GetSingleValue()

	pending, err := util.ExtractPendingEmbeddingsFromMarkdown(bytes, idColumn, contentColumn)
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

	outputSupplementaryId := activity.Outputs["supplementaryId"]
	return common.SetSupplementaryContent(ctx, executionContext, outputSupplementaryId, "application/protobuf", data)
}
