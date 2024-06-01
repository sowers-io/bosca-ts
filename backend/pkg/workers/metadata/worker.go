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

package metadata

import (
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/workers/common"
	"bosca.io/pkg/workers/metadata/processor"
	"bosca.io/pkg/workers/textextractor"
	"bosca.io/pkg/workers/vectors/vectorizer"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"time"
)

const TaskQueue = "metadata"

func NewWorker(client client.Client) worker.Worker {
	w := worker.New(client, TaskQueue, worker.Options{})
	w.RegisterWorkflow(ProcessMetadata)
	w.RegisterActivity(common.GetMetadata)
	w.RegisterActivity(processor.AddToSearchIndex)
	w.RegisterActivity(vectorizer.Vectorize)
	w.RegisterActivity(processor.TransitionToDraft)
	return w
}

func ProcessMetadata(ctx workflow.Context, id string) error {
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second,
		BackoffCoefficient:     2.0,
		MaximumInterval:        100 * time.Second,
		MaximumAttempts:        500, // 0 is unlimited retries
		NonRetryableErrorTypes: []string{"MissingMetadata"},
	}

	options := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
		RetryPolicy:         retryPolicy,
	}

	ctx = workflow.WithActivityOptions(ctx, options)

	var metadata *content.Metadata
	err := workflow.ExecuteActivity(ctx, common.GetMetadata, id).Get(ctx, &metadata)
	if err != nil {
		return err
	}

	var workflows []*processor.TraitWorkflow
	err = workflow.ExecuteActivity(ctx, processor.GetTraitWorkflows, metadata).Get(ctx, &workflows)
	if err != nil {
		return err
	}

	if workflows != nil {
		for _, trait := range workflows {
			traitWorkflowCtx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
				TaskQueue: trait.Queue,
			})
			err = workflow.ExecuteChildWorkflow(traitWorkflowCtx, trait.Id, metadata).Get(ctx, nil)
			if err != nil {
				return err
			}
		}
	}

	textExtractorCtx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
		TaskQueue: textextractor.TaskQueue,
	})
	err = workflow.ExecuteChildWorkflow(textExtractorCtx, textextractor.TextExtractor, metadata).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, processor.AddToSearchIndex, metadata).Get(ctx, nil)
	if err != nil {
		return err
	}

	err = workflow.ExecuteActivity(ctx, vectorizer.Vectorize, metadata).Get(ctx, &metadata)
	if err != nil {
		return err
	}

	return workflow.ExecuteActivity(ctx, processor.TransitionToDraft, metadata).Get(ctx, &metadata)
}
