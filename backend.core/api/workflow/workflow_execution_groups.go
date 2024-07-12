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
