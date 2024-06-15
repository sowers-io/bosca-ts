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
	"bosca.io/api/protobuf/bosca"
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workers/common"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/workflow"
	"time"
)

func ProcessTraits(ctx workflow.Context, id string) error {
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

	if metadata.TraitIds != nil && len(metadata.TraitIds) > 0 {
		contentService := common.GetWorkflowContentService(ctx)
		serviceCtx := common.GetWorkflowServiceAuthorizedContext(ctx)
		for _, traitId := range metadata.TraitIds {
			trait, err := contentService.GetTrait(serviceCtx, &bosca.IdRequest{
				Id: traitId,
			})
			if err != nil {
				return err
			}
			for _, workflowId := range trait.WorkflowIds {
				workflowInstance, err := contentService.GetTraitWorkflowInstance(
					serviceCtx,
					&content.TraitWorkflowIdRequest{
						TraitId:    traitId,
						WorkflowId: workflowId,
					},
				)
				if err != nil {
					return err
				}
				workflowInstance.Metadata = metadata
				traitWorkflowCtx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
					TaskQueue: workflowInstance.Queue,
				})
				err = workflow.ExecuteChildWorkflow(traitWorkflowCtx, workflowId, workflowInstance).Get(ctx, nil)
				if err != nil {
					return err
				}
			}
		}
	}

	return nil
}
