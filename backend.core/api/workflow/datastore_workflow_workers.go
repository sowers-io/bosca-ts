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
	grpc "bosca.io/api/protobuf/bosca/workflow"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"time"
)

func (ds *DataStore) RegisterWorker(ctx context.Context, request *grpc.WorkflowActivityJobRequest) (string, error) {
	cfg, err := json.Marshal(request)
	if err != nil {
		return "", err
	}
	result, err := ds.db.QueryContext(ctx, "insert into workflow_workers (registered, configuration) values (now(), $1::jsonb) returning id::varchar", cfg)
	if err != nil {
		return "", err
	}
	defer result.Close()
	if result.Next() {
		var id string
		err = result.Scan(&id)
		if err != nil {
			return "", err
		}
		return id, nil
	}
	return "", errors.New("failed to register worker")
}

func (ds *DataStore) UnregisterWorker(ctx context.Context, workerId string) error {
	_, err := ds.db.ExecContext(ctx, "delete from workflow_workers where id = $1::uuid", workerId)
	return err
}

func (ds *DataStore) ClaimNextJob(ctx context.Context, txn *sql.Tx, workerId string, queue string, activityIds []string) (*grpc.WorkflowActivityJob, error) {
	result, err := txn.QueryContext(ctx, "select j.id::varchar, j.workflow_execution_id::varchar, we.workflow_id, we.metadata_id::varchar, j.workflow_activity_id, j.context, j.completed, j.failed, j.error from claim_next_job($1, $2, $3) as j inner join workflow_executions we on (j.workflow_execution_id = we.id)", workerId, queue, activityIds)
	if err != nil {
		return nil, err
	}
	defer result.Close()
	if result.Next() {
		var jobId string
		var executionId string
		var workflowId string
		var metadataId string
		var workflowActivityId int64
		var contextJson json.RawMessage
		var complete *time.Time
		var failed *time.Time
		var errorStr *string
		err = result.Scan(&jobId, &executionId, &workflowId, &metadataId, &workflowActivityId, &contextJson, &complete, &failed, &errorStr)
		if err != nil {
			return nil, err
		}
		activity, err := ds.GetWorkflowActivity(ctx, workflowActivityId)
		if err != nil {
			return nil, err
		}
		jobContext := make(map[string]string)
		err = json.Unmarshal(contextJson, &jobContext)
		if err != nil {
			return nil, err
		}
		job := &grpc.WorkflowActivityJob{
			JobId:       jobId,
			ExecutionId: executionId,
			WorkflowId:  workflowId,
			MetadataId:  metadataId,
			Activity:    activity,
			Context:     jobContext,
			Complete:    complete != nil,
			Success:     complete != nil && failed == nil,
			Error:       errorStr,
		}
		job.Prompts, err = ds.GetWorkflowActivityPrompts(ctx, activity.WorkflowActivityId)
		if err != nil {
			return nil, err
		}
		job.StorageSystems, err = ds.GetWorkflowActivityStorageSystems(ctx, activity.WorkflowActivityId)
		if err != nil {
			return nil, err
		}
		job.Models, err = ds.GetWorkflowActivityModels(ctx, activity.WorkflowActivityId)
		if err != nil {
			return nil, err
		}
		return job, nil
	}
	return nil, nil
}
