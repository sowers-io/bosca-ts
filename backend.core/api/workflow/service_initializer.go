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

package workflow

import (
	workflow2 "bosca.io/api/protobuf/bosca/workflow"
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/search/meilisearch"
	"bosca.io/pkg/search/qdrant"
	"bosca.io/pkg/workflow"
	"context"
	meilisearch2 "github.com/meilisearch/meilisearch-go"
	goclient "github.com/qdrant/go-client/qdrant"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
	"os"
	"strconv"
)

func initializeService(cfg *configuration.ServerConfiguration, dataStore *DataStore) {
	ctx := context.Background()

	initializeEmbeddedWorkflow(ctx, dataStore)

	systems, err := dataStore.GetStorageSystems(ctx)
	if err != nil {
		slog.Error("failed to get storage systems", slog.Any("error", err))
		os.Exit(1)
	}

	meilisearchClient := meilisearch.NewMeilisearchClient(cfg.Search.Endpoint, cfg.Search.ApiKey)
	qdrantClient, err := qdrant.NewQdrantClient(cfg.ClientEndPoints.QdrantApiAddress)
	if err != nil {
		slog.Error("failed to get qdrant client", slog.Any("error", err))
		os.Exit(1)
	}
	defer qdrantClient.Close()

	for _, system := range systems {
		if system.Type == workflow2.StorageSystemType_vector_storage_system {
			initializeQdrant(ctx, qdrantClient, system)
		}
		if system.Type == workflow2.StorageSystemType_search_storage_system {
			initializeMeilisearch(meilisearchClient, system)
		}
	}
}

func initializeEmbeddedWorkflow(ctx context.Context, dataStore *DataStore) {
	defaultWorkflow, err := dataStore.GetWorkflow(ctx, "metadata.process")
	if err != nil {
		slog.Error("failed to get default model", slog.Any("error", err))
		os.Exit(1)
	}
	if defaultWorkflow != nil {
		return
	}
	cfg := workflow.GetEmbeddedConfiguration()
	models := make(map[string]string)
	prompts := make(map[string]string)
	storageSystems := make(map[string]string)
	for configId, model := range cfg.Models {
		id, err := dataStore.AddModel(ctx, &workflow2.Model{
			Name:          model.Name,
			Type:          model.Type,
			Description:   model.Name,
			Configuration: model.Configuration,
		})
		if err != nil {
			slog.Error("failed to create model", slog.Any("error", err))
			os.Exit(1)
		}
		models[configId] = id
	}
	for configId, prompt := range cfg.Prompts {
		id, err := dataStore.AddPrompt(ctx, &workflow2.Prompt{
			Name:         prompt.Name,
			Description:  prompt.Description,
			SystemPrompt: prompt.SystemPrompt,
			UserPrompt:   prompt.UserPrompt,
		})
		if err != nil {
			slog.Error("failed to create prompt", slog.Any("error", err))
			os.Exit(1)
		}
		prompts[configId] = id
	}
	for configId, storageSystem := range cfg.StorageSystems {
		id, err := dataStore.AddStorageSystem(ctx, &workflow2.StorageSystem{
			Name:          storageSystem.Name,
			Type:          workflow2.StorageSystemType(workflow2.StorageSystemType_value[storageSystem.Type+"_storage_system"]),
			Description:   storageSystem.Description,
			Configuration: storageSystem.Configuration,
		})
		if err != nil {
			slog.Error("failed to create storage system", slog.Any("error", err))
			os.Exit(1)
		}
		for modelId, model := range storageSystem.Models {
			err = dataStore.AddStorageSystemModel(ctx, id, models[modelId], model.Configuration)
			if err != nil {
				slog.Error("failed to create storage system model", slog.Any("error", err))
				os.Exit(1)
			}
		}
		storageSystems[configId] = id
	}
	for id, activity := range cfg.Workflows.Activities {
		inputs := make(map[string]workflow2.WorkflowActivityParameterType)
		for key, _ := range activity.Inputs {
			inputs[key] = workflow2.WorkflowActivityParameterType_supplementary
		}
		outputs := make(map[string]workflow2.WorkflowActivityParameterType)
		for key, _ := range activity.Outputs {
			outputs[key] = workflow2.WorkflowActivityParameterType_supplementary
		}
		err := dataStore.AddActivity(ctx, &workflow2.Activity{
			Id:              id,
			Name:            activity.Name,
			Description:     activity.Description,
			ChildWorkflowId: activity.ChildWorkflowId,
			Inputs:          inputs,
			Outputs:         outputs,
			Configuration:   activity.Configuration,
		})
		if err != nil {
			slog.Error("failed to create activity", slog.Any("error", err))
			os.Exit(1)
		}
	}
	for id, wf := range cfg.Workflows.Workflows {
		err := dataStore.AddWorkflow(ctx, &workflow2.Workflow{
			Id:            id,
			Name:          wf.Name,
			Description:   wf.Description,
			Queue:         wf.Queue,
			Configuration: wf.Configuration,
		})
		for aid, activity := range wf.Activities {
			inputs := make(map[string]string)
			for key, val := range activity.Inputs {
				inputs[key] = val
			}
			outputs := make(map[string]string)
			for key, val := range activity.Outputs {
				outputs[key] = val
			}
			id, err := dataStore.AddWorkflowActivity(ctx, id, &workflow2.WorkflowActivity{
				ActivityId:     aid,
				Queue:          wf.Queue,
				ExecutionGroup: activity.ExecutionGroup,
				Configuration:  activity.Configuration,
				Inputs:         inputs,
				Outputs:        outputs,
			})
			for systemId, system := range activity.StorageSystems {
				err = dataStore.AddWorkflowActivityStorageSystem(ctx, id, storageSystems[systemId], system.Configuration)
				if err != nil {
					slog.Error("failed to add workflow storage system", slog.Any("error", err))
					os.Exit(1)
				}
			}
			for promptId, prompt := range activity.Prompts {
				err = dataStore.AddWorkflowActivityInstancePrompt(ctx, id, prompts[promptId], prompt.Configuration)
				if err != nil {
					slog.Error("failed to add workflow prompt", slog.Any("error", err))
					os.Exit(1)
				}
			}
			for modelId, model := range activity.Models {
				err = dataStore.AddWorkflowActivityModel(ctx, id, models[modelId], model.Configuration)
				if err != nil {
					slog.Error("failed to add workflow prompt", slog.Any("error", err))
					os.Exit(1)
				}
			}
			if err != nil {
				slog.Error("failed to create workflow activity instance", slog.Any("error", err))
				os.Exit(1)
			}
		}
		if err != nil {
			slog.Error("failed to create workflow", slog.Any("error", err))
			os.Exit(1)
		}
	}
	for id, state := range cfg.Workflows.States {
		var wf *string
		var entryWf *string
		var exitWf *string
		if state.Workflow != "" {
			wf = &state.Workflow
		}
		if state.EntryWorkflow != "" {
			entryWf = &state.EntryWorkflow
		}
		if state.ExitWorkflow != "" {
			exitWf = &state.ExitWorkflow
		}
		err := dataStore.AddWorkflowState(ctx, &workflow2.WorkflowState{
			Id:              id,
			Name:            state.Name,
			Type:            workflow2.WorkflowStateType(workflow2.WorkflowStateType_value[state.Type]),
			WorkflowId:      wf,
			EntryWorkflowId: entryWf,
			ExitWorkflowId:  exitWf,
			Configuration:   state.Configuration,
		})
		if err != nil {
			slog.Error("failed to create workflow activity instance", slog.Any("error", err))
			os.Exit(1)
		}
	}
	for _, transition := range cfg.Workflows.Transitions {
		err := dataStore.AddWorkflowTransition(ctx, &workflow2.WorkflowStateTransition{
			ToStateId:   transition.To,
			FromStateId: transition.From,
			Description: transition.Description,
		})
		if err != nil {
			slog.Error("failed to create transition", slog.Any("error", err))
			os.Exit(1)
		}
	}
}

func initializeQdrant(ctx context.Context, qdrantClient *qdrant.Client, system *workflow2.StorageSystem) {
	_, err := qdrantClient.CollectionsClient.Get(ctx, &goclient.GetCollectionInfoRequest{
		CollectionName: system.Configuration["indexName"],
	})
	if err != nil {
		slog.Warn("error getting qdrant collection info, trying to create collection", slog.Any("error", err), slog.String("collectionName", system.Configuration["indexName"]))
		if s, ok := status.FromError(err); ok {
			if s.Code() == codes.NotFound {
				size, err := strconv.ParseInt(system.Configuration["vectorSize"], 0, 64)
				if err != nil {
					slog.Error("failed to parse vector size in system configuration", slog.Any("error", err))
					os.Exit(1)
				}
				collection := &goclient.CreateCollection{
					CollectionName: system.Configuration["indexName"],
					VectorsConfig: &goclient.VectorsConfig{
						Config: &goclient.VectorsConfig_Params{
							Params: &goclient.VectorParams{
								Size:     uint64(size),
								Distance: goclient.Distance_Cosine,
							},
						},
					},
				}
				result, err := qdrantClient.CollectionsClient.Create(ctx, collection)
				if err != nil {
					slog.Error("failed to create qdrant collection", slog.Any("error", err))
					os.Exit(1)
				}
				if !result.Result {
					slog.Error("failed to create qdrant collection")
					os.Exit(1)
				}
			}
		} else {
			slog.Error("failed to create qdrant collection", slog.Any("error", err))
			os.Exit(1)
		}
	} else {
		slog.Info("qdrant collection already exists", slog.Any("collectionName", system.Configuration["indexName"]))
	}
}

func initializeMeilisearch(client *meilisearch2.Client, system *workflow2.StorageSystem) {
	ix, err := client.GetIndex(system.Configuration["indexName"])
	if ix != nil {
		slog.Info("meilisearch index already exists", slog.String("indexName", system.Configuration["indexName"]))
		return
	}
	slog.Warn("error getting meilisearch index info, trying to create index", slog.Any("error", err), slog.String("index", system.Configuration["indexName"]))
	_, err = client.CreateIndex(&meilisearch2.IndexConfig{
		Uid:        system.Configuration["indexName"],
		PrimaryKey: "id",
	})
	if err != nil {
		slog.Error("failed to create meilisearch index", slog.Any("error", err))
		os.Exit(1)
	}
}
