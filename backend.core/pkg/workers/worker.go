package workers

import (
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workers/bible"
	"bosca.io/pkg/workers/embeddings"
	"bosca.io/pkg/workers/textextractor"
	"context"
	"errors"
	"go.temporal.io/sdk/activity"
	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/temporal"
	"go.temporal.io/sdk/worker"
	"go.temporal.io/sdk/workflow"
	"time"
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
			if instance.Activity.ChildWorkflow {
				childQueue := executionContext.Workflow.Queue
				if instance.Activity.ChildWorkflowQueue != nil {
					childQueue = *instance.Activity.ChildWorkflowQueue
				}
				childCtx := workflow.WithChildOptions(ctx, workflow.ChildWorkflowOptions{
					TaskQueue: childQueue,
				})
				futures[i] = workflow.ExecuteChildWorkflow(childCtx, instance.Activity.Id, instanceExecutionContext)
			} else {
				futures[i] = workflow.ExecuteActivity(ctx, instance.Activity.Id, instanceExecutionContext)
			}
		}
		for _, future := range futures {
			var result *content.WorkflowActivityResult
			err := future.Get(ctx, &result)
			if err != nil {
				return err
			}
			for n, value := range result.Context {
				executionContext.Context[n] = value
			}
			for n, value := range result.Outputs {
				executionContext.Inputs[n] = value
			}
		}
	}

	return nil
}

func ProcessActivity(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	switch executionContext.Activity.Activity.Id {
	case "ProcessSupplementaryTraits":
		panic("TODO")
	case "ExtractToSupplementaryText":
		return textextractor.ExtractToSupplementaryText(ctx, executionContext)
	case "ExtractChaptersToMetadata":
		return bible.ExtractChaptersToMetadata(ctx, executionContext)
	case "ExtractVersesToMetadata":
		return bible.ExtractVersesToMetadata(ctx, executionContext)
	case "CreateSupplementaryVerseMarkdownTable":
		return bible.CreateSupplementaryVerseMarkdownTable(ctx, executionContext)
	case "CreatePendingEmbeddingsFromMarkdownTable":
		return embeddings.CreatePendingEmbeddingsFromMarkdownTable(ctx, executionContext)
	case "DeleteMetadata":
		panic("TODO")
	case "DeleteSupplementary":
		panic("TODO")
	default:
		return errors.New("unknown activity id: " + executionContext.Activity.Activity.Id)
	}
	return nil
}
