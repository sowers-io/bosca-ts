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

package registry

import (
	"bosca.io/api/protobuf/bosca/content"
	"context"
	"go.temporal.io/sdk/workflow"
)

var activityRegistry = make(map[string]func(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error)
var workflowRegistry = make(map[string]func(ctx workflow.Context, executionContext *content.WorkflowActivityExecutionContext) error)

func GetActivity(name string) func(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	return activityRegistry[name]
}

func GetWorkflow(name string) func(ctx workflow.Context, executionContext *content.WorkflowActivityExecutionContext) error {
	return workflowRegistry[name]
}

func RegisterActivity(name string, fn func(ctx context.Context, executionContext *content.WorkflowActivityExecutionContext) error) {
	if _, ok := activityRegistry[name]; ok {
		panic("duplicate activity name: " + name)
	}
	activityRegistry[name] = fn
}

func RegisterWorkflow(name string, fn func(ctx workflow.Context, executionContext *content.WorkflowActivityExecutionContext) error) {
	if _, ok := workflowRegistry[name]; ok {
		panic("duplicate workflow name: " + name)
	}
	workflowRegistry[name] = fn
}
