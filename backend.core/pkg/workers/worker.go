package workers

import (
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/workers/bible"
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
			instanceExecutionContextClone := make(map[string]*content.WorkflowActivityParameterValue)
			for key, value := range executionContext.Context {
				instanceExecutionContextClone[key] = value
			}
			instanceExecutionInputClone := make(map[string]*content.WorkflowActivityParameterValue)
			for key, value := range executionContext.Inputs {
				instanceExecutionInputClone[key] = value
			}
			instanceExecutionContext := &content.WorkflowActivityExecutionContext{
				Workflow: executionContext.Workflow,
				Activity: instance,
				Context:  instanceExecutionContextClone,
				Inputs:   instanceExecutionInputClone,
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
	case "GenerateChapters":
		return bible.GenerateChapters(ctx, executionContext)
	/*
			('ProcessTraits', 'Process Metadata Traits', 'Execute workflows based on metadata traits', '{}'::jsonb),
		       ('ProcessSupplementaryTraits', 'Process Supplementary Metadata Traits', 'Execute workflows based on metadata supplementary traits', '{}'::jsonb),
		       ('GeneratePendingEmbeddingsFromTable', 'Generate Pending Embeddings from a Table', 'Generate Pending Embeddings from a Table, the supplied `column` will be the data', '{}'::jsonb),
		       ('ProcessUSX', 'Process USX', '', '{}'::jsonb),
		       ('GenerateChapters', 'Generate chapters', '', '{}'::jsonb),
		       ('GenerateChapterVerses', 'Generate verses', '', '{}'::jsonb),
		       ('GenerateChapterVerseTable', 'Generate Chapter Verse Table', 'Generate chapter verse table for the purposes of generating verse labels', '{}'::jsonb),
		       ('GenerateVerseLabelPendingEmbeddings', 'Generate Verse Label Pending Embeddings', '', '{}'::jsonb),
		       ('ExtractText', 'Extract Text from Content', 'Extract text from the main content', '{}'::jsonb),
		       ('ExecuteTablePrompt', 'Execute a table prompt', 'Execute a prompt by using the supplementary table data and save the results as supplementary data.  Uses context to leverage `inSupplementaryId` and `outSupplementaryId`', '{}'::jsonb),
		       ('GenerateTextEmbeddings', 'Generate Embeddings', 'Generate embeddings based on main content', '{}'::jsonb),
		       ('GenerateSupplementaryPendingEmbeddings', 'Generate Supplementary Embeddings', 'Generate embeddings based on supplementary content', '{}'::jsonb),
		       ('AddToVectorIndex', 'Add to Vector Index', 'Add pending embeddings to vector index', '{}'::jsonb),
		       ('AddToSearchIndex', 'Add to Search Index', 'Add text to search index', '{}'::jsonb),
		       ('DeleteMetadata', 'Delete Metadata', 'Delete metadata', '{}'::jsonb),
		       ('DeleteSupplementary', 'Del
	*/
	default:
		return errors.New("unknown activity id: " + executionContext.Activity.Activity.Id)
	}
}
