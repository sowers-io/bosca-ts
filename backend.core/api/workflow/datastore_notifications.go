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
	"bosca.io/api/protobuf/bosca/workflow"
	"context"
	"log/slog"
)

func (ds *DataStore) notify(ctx context.Context, queues []string, notification *workflow.WorkflowExecutionNotification) error {
	slog.InfoContext(ctx, "notify", slog.Any("queues", queues), slog.Any("notification", notification))
	for _, queue := range queues {
		err := ds.pubsub.Publish(ctx, queue, notification)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ds *DataStore) NotifyJobAvailable(ctx context.Context, queues []string) error {
	notification := &workflow.WorkflowExecutionNotification{
		Type: workflow.WorkflowExecutionNotificationType_job_available,
	}
	return ds.notify(ctx, queues, notification)
}

func (ds *DataStore) NotifyExecutionCompletion(ctx context.Context, queues []string, executionId string, success bool) error {
	notification := &workflow.WorkflowExecutionNotification{
		Type:        workflow.WorkflowExecutionNotificationType_execution_completion,
		ExecutionId: executionId,
		Success:     success,
	}
	return ds.notify(ctx, queues, notification)
}

func (ds *DataStore) NotifyExecutionFailed(ctx context.Context, queues []string, executionId string, err error) error {
	notification := &workflow.WorkflowExecutionNotification{
		Type:        workflow.WorkflowExecutionNotificationType_execution_failed,
		ExecutionId: executionId,
	}
	if err != nil {
		e := err.Error()
		notification.Error = &e
	}
	return ds.notify(ctx, queues, notification)
}
