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

func ProcessWorkflow(ctx workflow.Context, executionContext *content.WorkflowActivityExecutionContext) error {
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
	wf := NewWorkflow(executionContext)
	executionGroup := wf.GetExecutionGroups()
	for _, group := range executionGroup {
		futures := make([]workflow.Future, len(group))
		for i, instance := range group {
			inputs := make(map[string]*content.WorkflowActivityParameterValue)
			for k, input := range instance.Inputs {
				inputs[k] = input
			}
			for k, input := range executionContext.Inputs {
				inputs[k] = input
			}
			instanceExecutionContext := &content.WorkflowActivityExecutionContext{
				Metadata: executionContext.Metadata,
				Workflow: executionContext.Workflow,
				Activity: instance,
				Context:  executionContext.Context,
				Inputs:   inputs,
			}
			if instance.ChildWorkflow {
				childQueue := executionContext.Workflow.Queue
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
			err := future.Get(ctx, nil)
			if err != nil {
				return err
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
