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
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workflow/registry"
	"errors"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"slices"
	"strings"
	"time"

	_ "bosca.io/pkg/workflow/ai/markdown"
	_ "bosca.io/pkg/workflow/bible"
	_ "bosca.io/pkg/workflow/common"
	_ "bosca.io/pkg/workflow/metadata"
	_ "bosca.io/pkg/workflow/registry"
)

func NewWorker(client client.Client, workflowIds []string, activityIds []string, queue string) (worker.Worker, error) {
	w := worker.New(client, queue, worker.Options{})
	for _, workflowId := range workflowIds {
		name := strings.Trim(workflowId, " ")
		if name == "" {
			continue
		}
		workflowFunc := registry.GetWorkflow(name)
		if workflowFunc == nil {
			workflowFunc = processWorkflow
		}
		w.RegisterWorkflowWithOptions(workflowFunc, workflow.RegisterOptions{
			Name: name,
		})
	}
	for _, activityId := range activityIds {
		name := strings.Trim(activityId, " ")
		if name == "" {
			continue
		}
		activityFn := registry.GetActivity(name)
		if activityFn == nil {
			return nil, errors.New("missing activity: " + name)
		}
		w.RegisterActivityWithOptions(activityFn, activity.RegisterOptions{
			Name: name,
		})
	}
	return w, nil
}

func processWorkflow(ctx workflow.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	retryPolicy := &temporal.RetryPolicy{
		InitialInterval:        time.Second / 2,
		BackoffCoefficient:     1.5,
		MaximumInterval:        30 * time.Second,
		MaximumAttempts:        500, // 0 is unlimited retries
		NonRetryableErrorTypes: []string{"MissingMetadata", "rpc error: code = NotFound desc = metadata not found"},
	}
	options := workflow.ActivityOptions{
		StartToCloseTimeout: 10 * time.Minute,
		RetryPolicy:         retryPolicy,
	}
	ctx = workflow.WithActivityOptions(ctx, options)
	executionGroup := GetExecutionGroups(executionContext.Activities)
	workflowContext := make(map[string]*content.WorkflowActivityParameterValue)
	for _, group := range executionGroup {
		futures := make([]workflow.Future, len(group))
		for i, instance := range group {
			inputs := make(map[string]*content.WorkflowActivityParameterValue)
			for k, input := range instance.Inputs {
				inputs[k] = input
			}
			instanceExecutionContext := &content.WorkflowActivityExecutionContext{
				WorkflowId:           executionContext.WorkflowId,
				Metadata:             executionContext.Metadata,
				Activities:           executionContext.Activities,
				CurrentActivityIndex: int32(slices.Index(executionContext.Activities, instance)),
				Context:              workflowContext,
			}
			if instance.ChildWorkflow {
				childQueue := "default"
				if instance.ChildWorkflowQueue != nil {
					childQueue = *instance.ChildWorkflowQueue
				}
				childCtx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
					TaskQueue: childQueue,
				})
				futures[i] = workflow.ExecuteChildWorkflow(childCtx, instance.Id, instanceExecutionContext)
			} else {
				futures[i] = workflow.ExecuteActivity(ctx, instance.Id, instanceExecutionContext)
			}
		}
		for _, future := range futures {
			updatedContext := make(map[string]*content.WorkflowActivityParameterValue)
			err := future.Get(ctx, &updatedContext)
			if err != nil {
				return err
			}
			if updatedContext != nil {
				for s, value := range updatedContext {
					workflowContext[s] = value
				}
			}
		}
	}

	return nil
}
