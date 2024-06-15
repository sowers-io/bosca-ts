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
	"bosca.io/pkg/workers/common"
	"context"
)

func CompleteTransition(ctx context.Context, metadataId string) error {
	contentService := common.GetContentService(ctx)
	_, err := contentService.CompleteTransitionWorkflow(common.GetServiceAuthorizedContext(ctx), &content.CompleteTransitionWorkflowRequest{
		MetadataId: metadataId,
		Status:     "workflow transition",
		Success:    true,
	})
	return err
}

func TransitionTo(ctx context.Context, metadataId string, stateId string) error {
	contentService := common.GetContentService(ctx)
	_, err := contentService.BeginTransitionWorkflow(common.GetServiceAuthorizedContext(ctx), &content.TransitionWorkflowRequest{
		MetadataId: metadataId,
		StateId:    stateId,
		Status:     "workflow transition",
	})
	return err
}
