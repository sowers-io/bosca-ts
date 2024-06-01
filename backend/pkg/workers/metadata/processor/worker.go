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

package processor

import (
	"bosca.io/api/protobuf"
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/search"
	"bosca.io/pkg/workers/common"
	"bosca.io/pkg/workers/textextractor"
	"context"
	"log/slog"
	"os"
)

func TransitionToDraft(ctx context.Context, metadata *content.Metadata) error {
	contentService := common.GetContentService(ctx)
	if metadata.WorkflowStateId == "processing" {
		_, err := contentService.TransitionWorkflow(common.GetServiceAuthorizedContext(ctx), &content.TransitionWorkflowRequest{
			MetadataId: metadata.Id,
			StateId:    "draft",
		})
		return err
	}
	slog.Warn("Metadata is not in processing state", slog.String("metadataId", metadata.Id), slog.String("stateId", metadata.WorkflowStateId))
	return nil
}

func AddToSearchIndex(ctx context.Context, metadata *content.Metadata) error {
	svc := common.GetContentService(ctx)

	signedUrl, err := svc.GetMetadataSupplementaryDownloadUrl(common.GetServiceAuthorizedContext(ctx), &content.SupplementaryIdRequest{
		Id:   metadata.Id,
		Type: textextractor.SupplementalTextType,
	})
	if err != nil {
		return err
	}

	file, err := common.DownloadTemporaryFile(ctx, signedUrl)
	if err != nil {
		return err
	}
	defer os.Remove(file.Name())

	body, err := os.ReadFile(file.Name())
	if err != nil {
		return err
	}

	client := common.GetSearchClient(ctx)
	return client.Index(client.GetMetadataIndex(), &search.Document{
		Id:   metadata.Id,
		Name: metadata.Name,
		Body: string(body),
	})
}

type TraitWorkflow struct {
	Id    string
	Queue string
}

func GetTraitWorkflows(ctx context.Context, metadata *content.Metadata) ([]*TraitWorkflow, error) {
	if metadata.TraitIds == nil || len(metadata.TraitIds) == 0 {
		return nil, nil
	}

	svc := common.GetContentService(ctx)

	workflows, err := svc.GetWorkflows(common.GetServiceAuthorizedContext(ctx), &protobuf.Empty{})
	if err != nil {
		return nil, err
	}

	traits, err := svc.GetTraits(common.GetServiceAuthorizedContext(ctx), &protobuf.Empty{})
	if err != nil {
		return nil, err
	}

	traitsById := make(map[string]*content.Trait)
	for _, trait := range traits.Traits {
		traitsById[trait.Id] = trait
	}

	workflowsById := make(map[string]*content.Workflow)
	for _, workflow := range workflows.Workflows {
		workflowsById[workflow.Id] = workflow
	}

	traitWorkflows := make([]*TraitWorkflow, 0, len(metadata.TraitIds))
	for _, id := range metadata.TraitIds {
		if trait, ok := traitsById[id]; ok {
			if workflow, ok := workflowsById[trait.WorkflowId]; ok {
				traitWorkflows = append(traitWorkflows, &TraitWorkflow{
					Id:    workflow.Id,
					Queue: workflow.Queue,
				})
			}
		}
	}

	return traitWorkflows, nil
}
