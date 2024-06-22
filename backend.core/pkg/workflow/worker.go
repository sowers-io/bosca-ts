package workflow

import (
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workflow/registry"
	"context"
	"errors"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
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
		w.RegisterWorkflowWithOptions(ProcessWorkflow, workflow.RegisterOptions{
			Name: workflowId,
		})
	}
	for _, activityId := range activityIds {
		w.RegisterActivityWithOptions(ProcessActivity, activity.RegisterOptions{
			Name: activityId,
		})
	}
	return w, nil
}

func ProcessWorkflow(ctx workflow.Context, workflowInstance *content.WorkflowInstance) error {
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
	executionGroup := GetExecutionGroups(workflowInstance.Activities)
	workflowContext := make(map[string]*content.WorkflowActivityParameterValue)
	for _, group := range executionGroup {
		futures := make([]workflow.Future, len(group))
		for i, instance := range group {
			inputs := make(map[string]*content.WorkflowActivityParameterValue)
			for k, input := range instance.Inputs {
				inputs[k] = input
			}
			instanceExecutionContext := &content.WorkflowActivityExecutionContext{
				WorkflowId: workflowInstance.Workflow.Id,
				TraitId:    workflowInstance.TraitId,
				Metadata:   workflowInstance.Metadata,
				Activity:   instance,
				Context:    workflowContext,
				Inputs:     inputs,
			}
			if instance.ChildWorkflow {
				childQueue := "metadata"
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

func ProcessActivity(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	activityFn := registry.GetActivity(executionContext.Activity.Id)
	if activityFn == nil {
		return errors.New("TODO: " + executionContext.Activity.Id)
	}
	return activityFn(ctx, executionContext)
}
