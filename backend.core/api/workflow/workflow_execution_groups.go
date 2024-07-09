package workflow

import "bosca.io/api/protobuf/bosca/workflow"

func getAllQueues(executionContext *workflow.WorkflowExecutionContext) []string {
	queuesMap := make(map[string]bool)
	for _, activity := range executionContext.Activities {
		queuesMap[activity.Queue] = true
	}
	queues := make([]string, len(queuesMap))
	i := 0
	for queue, _ := range queuesMap {
		queues[i] = queue
		i++
	}
	return queues
}

func calculateWorkflowExecutionGroups(executionContext *workflow.WorkflowExecutionContext) [][]*workflow.WorkflowActivity {
	executionGroups := make([][]*workflow.WorkflowActivity, 0, len(executionContext.Activities))
	executionGroupIndex := int32(0)
	executionGroup := make([]*workflow.WorkflowActivity, 0, 1)
	for _, activity := range executionContext.Activities {
		if activity.ExecutionGroup != executionGroupIndex {
			executionGroupIndex = activity.ExecutionGroup
			if len(executionGroup) > 0 {
				executionGroups = append(executionGroups, executionGroup)
				executionGroup = make([]*workflow.WorkflowActivity, 0, 1)
			}
		}
		executionGroup = append(executionGroup, activity)
	}
	if len(executionGroup) > 0 {
		executionGroups = append(executionGroups, executionGroup)
	}
	return executionGroups
}
