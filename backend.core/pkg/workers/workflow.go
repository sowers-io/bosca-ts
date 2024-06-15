package workers

import (
	"bosca.io/api/protobuf/bosca/content"
)

type Workflow struct {
	ctx *content.WorkflowActivityExecutionContext
}

func NewWorkflow(ctx *content.WorkflowActivityExecutionContext) *Workflow {
	return &Workflow{
		ctx: ctx,
	}
}

func (w *Workflow) GetExecutionGroups() [][]*content.WorkflowActivityInstance {
	executionGroups := make([][]*content.WorkflowActivityInstance, 0, len(w.ctx.Workflow.Activities))
	executionGroupIndex := int32(0)
	executionGroup := make([]*content.WorkflowActivityInstance, 0, 1)
	for _, activity := range w.ctx.Workflow.Activities {
		if activity.ExecutionGroup != executionGroupIndex {
			executionGroups = append(executionGroups, executionGroup)
			executionGroup = make([]*content.WorkflowActivityInstance, 0, 1)
		}
		executionGroup = append(executionGroup, activity)
	}
	if len(executionGroup) > 0 {
		executionGroups = append(executionGroups, executionGroup)
	}
	return executionGroups
}
