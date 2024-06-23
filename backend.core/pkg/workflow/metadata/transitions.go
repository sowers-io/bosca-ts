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
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workflow/common"
	"bosca.io/pkg/workflow/registry"
	"context"
)

func init() {
	registry.RegisterActivity("metadata.transition.complete", completeTransition)
	registry.RegisterActivity("metadata.transition.fail", failTransition)
	registry.RegisterActivity("metadata.transition.to", transitionTo)
}

func completeTransition(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	contentService := common.GetContentService(ctx)
	_, err := contentService.CompleteTransitionWorkflow(common.GetServiceAuthorizedContext(ctx), &content.CompleteTransitionWorkflowRequest{
		MetadataId: executionContext.Metadata.Id,
		Status:     executionContext.Activities[executionContext.CurrentActivityIndex].Configuration["status"],
		Success:    true,
	})
	return err
}

func failTransition(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	contentService := common.GetContentService(ctx)
	_, err := contentService.CompleteTransitionWorkflow(common.GetServiceAuthorizedContext(ctx), &content.CompleteTransitionWorkflowRequest{
		MetadataId: executionContext.Metadata.Id,
		Status:     executionContext.Activities[executionContext.CurrentActivityIndex].Configuration["status"],
		Success:    false,
	})
	return err
}

func transitionTo(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	activity := executionContext.Activities[executionContext.CurrentActivityIndex]
	contentService := common.GetContentService(ctx)
	_, err := contentService.BeginTransitionWorkflow(common.GetServiceAuthorizedContext(ctx), &content.TransitionWorkflowRequest{
		MetadataId: executionContext.Metadata.Id,
		StateId:    activity.Configuration["state"],
		Status:     activity.Configuration["status"],
	})
	return err
}
