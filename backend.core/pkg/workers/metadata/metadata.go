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
	content2 "bosca.io/api/content"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workers/common"
	"bosca.io/pkg/workers/metadata/processor"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

func ProcessMetadata(ctx workflow.Context, id string) error {
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second / 2,
		BackoffCoefficient:     1.5,
		MaximumInterval:        30 * time.Second,
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

	if metadata.TraitIds == nil || len(metadata.TraitIds) == 0 {
		// TODO: run activity to guess traits
	} else {
		traitsCtx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
			TaskQueue: TaskQueue,
		})
		err = workflow.ExecuteChildWorkflow(traitsCtx, ProcessTraits, id).Get(ctx, nil)
		if err != nil {
			return err
		}
	}

	err = workflow.ExecuteActivity(ctx, processor.CompleteTransition, metadata).Get(ctx, &metadata)
	if err != nil {
		return err
	}

	return workflow.ExecuteActivity(ctx, processor.TransitionTo, &processor.Transition{
		Metadata: metadata,
		StateId:  content2.WorkflowStateDraft,
	}).Get(ctx, &metadata)
}
