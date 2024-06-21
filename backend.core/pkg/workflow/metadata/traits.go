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
	"bosca.io/pkg/workflow/common"
	"go.temporal.io/sdk/workflow"
)

func ProcessTraits(ctx workflow.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	if executionContext.Metadata.TraitIds != nil && len(executionContext.Metadata.TraitIds) > 0 {
		contentService := common.GetWorkflowContentService(ctx)
		serviceCtx := common.GetWorkflowServiceAuthorizedContext(ctx)
		for _, traitId := range executionContext.Metadata.TraitIds {
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
