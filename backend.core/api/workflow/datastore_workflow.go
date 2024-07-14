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
	"database/sql"
	"encoding/json"
	"errors"
	"log/slog"
	"time"
)

func (ds *DataStore) AddWorkflow(ctx context.Context, workflow *workflow.Workflow) error {
	config, err := json.Marshal(workflow.Configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO workflows (id, name, description, queue, configuration) VALUES ($1, $2, $3, $4, ($5)::jsonb)",
		workflow.Id, workflow.Name, workflow.Description, workflow.Queue, string(config),
	)
	return err
}

func (ds *DataStore) GetAllQueues(ctx context.Context) ([]string, error) {
	rows, err := ds.db.QueryContext(ctx, "select distinct queue from workflows")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	queues := make([]string, 0)
	for rows.Next() {
		var queue string
		err = rows.Scan(&queue)
		if err != nil {
			return nil, err
		}
		queues = append(queues, queue)
	}
	return queues, nil
}

func (ds *DataStore) GetWorkflow(ctx context.Context, id string) (*workflow.Workflow, error) {
	row := ds.db.QueryRowContext(ctx, "SELECT id, name, description, queue, configuration FROM workflows WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var wf workflow.Workflow
	var configuration json.RawMessage
	err := row.Scan(&wf.Id, &wf.Name, &wf.Description, &wf.Queue, &configuration)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	err = json.Unmarshal(configuration, &wf.Configuration)
	if err != nil {
		return nil, err
	}
	return &wf, nil
}

func getWorkflowActivityParameterValue(inputs *sql.Rows) (string, string, error) {
	var name string
	var value string
	err := inputs.Scan(&name, &value)
	if err != nil {
		return "", "", err
	}
	return name, value, nil
}

func (ds *DataStore) getWorkflowActivity(ctx context.Context, rows *sql.Rows) (*workflow.WorkflowActivity, error) {
	activity := &workflow.WorkflowActivity{
		Inputs:  make(map[string]string),
		Outputs: make(map[string]string),
	}

	var configuration json.RawMessage
	err := rows.Scan(&activity.WorkflowActivityId, &activity.ActivityId, &activity.ChildWorkflowId, &activity.ExecutionGroup, &configuration, &activity.Queue)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(configuration, &activity.Configuration)
	if err != nil {
		return nil, err
	}

	// TODO: prepared statements
	inputs, err := ds.db.QueryContext(ctx, "select name, value from workflow_activity_inputs where activity_id = $1", activity.WorkflowActivityId)
	if err != nil {
		return nil, err
	}
	for inputs.Next() {
		name, input, err := getWorkflowActivityParameterValue(inputs)
		if err != nil {
			inputs.Close()
			return nil, err
		}
		activity.Inputs[name] = input
	}
	inputs.Close()

	outputs, err := ds.db.QueryContext(ctx, "select name, value from workflow_activity_outputs where activity_id = $1", activity.WorkflowActivityId)
	if err != nil {
		return nil, err
	}
	for outputs.Next() {
		name, output, err := getWorkflowActivityParameterValue(outputs)
		if err != nil {
			outputs.Close()
			return nil, err
		}
		activity.Outputs[name] = output
	}
	outputs.Close()
	return activity, nil
}

func (ds *DataStore) GetWorkflowActivity(ctx context.Context, id int64) (*workflow.WorkflowActivity, error) {
	rows, err := ds.db.QueryContext(ctx, "SELECT wa.id, wa.activity_id, a.child_workflow_id, wa.execution_group, wa.configuration, w.queue FROM workflow_activities wa inner join activities a on (wa.activity_id = a.id) inner join workflows w on (wa.workflow_id = w.id) WHERE wa.id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		return ds.getWorkflowActivity(ctx, rows)
	}
	return nil, nil
}

func (ds *DataStore) GetWorkflowActivities(ctx context.Context, workflowId string) ([]*workflow.WorkflowActivity, error) {
	rows, err := ds.db.QueryContext(ctx, "SELECT wa.id, wa.activity_id, a.child_workflow_id, wa.execution_group, wa.configuration, w.queue FROM workflow_activities wa inner join activities a on (wa.activity_id = a.id) inner join workflows w on (wa.workflow_id = w.id) where workflow_id = $1 order by wa.execution_group", workflowId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	activities := make([]*workflow.WorkflowActivity, 0)
	for rows.Next() {
		activity, err := ds.getWorkflowActivity(ctx, rows)
		if err != nil {
			return nil, err
		}
		activities = append(activities, activity)
	}
	return activities, nil
}

func (ds *DataStore) GetWorkflows(ctx context.Context) ([]*workflow.Workflow, error) {
	result, err := ds.db.QueryContext(ctx, "SELECT id, name, description, queue, configuration FROM workflows")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	workflows := make([]*workflow.Workflow, 0)
	for result.Next() {
		var wf workflow.Workflow
		var configuration json.RawMessage
		err := result.Scan(&wf.Id, &wf.Name, &wf.Description, &wf.Queue, &configuration)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, err
		}
		err = json.Unmarshal(configuration, &wf.Configuration)
		if err != nil {
			return nil, err
		}
		workflows = append(workflows, &wf)
	}
	return workflows, nil
}

func (ds *DataStore) AddActivity(ctx context.Context, activity *workflow.Activity) error {
	config, err := json.Marshal(activity.Configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO activities (id, name, description, child_workflow_id, configuration) VALUES ($1, $2, $3, $4, ($5)::jsonb)",
		activity.Id, activity.Name, activity.Description, activity.ChildWorkflowId, string(config),
	)
	if err != nil {
		return err
	}

	stmtInput, err := ds.db.PrepareContext(ctx, "INSERT INTO activity_inputs (activity_id, name, type) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmtInput.Close()

	for name, parameterType := range activity.Inputs {
		_, err := stmtInput.ExecContext(ctx, activity.Id, name, parameterType.String())
		if err != nil {
			return err
		}
	}

	stmtOutput, err := ds.db.PrepareContext(ctx, "INSERT INTO activity_outputs (activity_id, name, type) VALUES ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmtOutput.Close()

	for name, parameterType := range activity.Outputs {
		_, err := stmtOutput.ExecContext(ctx, activity.Id, name, parameterType.String())
		if err != nil {
			return err
		}
	}

	return nil
}

func (ds *DataStore) GetWorkflowTransition(ctx context.Context, fromStateId string, toStateId string) (*workflow.WorkflowStateTransition, error) {
	row := ds.db.QueryRowContext(ctx, "SELECT from_state_id, to_state_id FROM workflow_state_transitions WHERE from_state_id = $1 AND to_state_id = $2", fromStateId, toStateId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var transition workflow.WorkflowStateTransition
	err := row.Scan(&transition.FromStateId, &transition.ToStateId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &transition, nil
}

func (ds *DataStore) AddWorkflowState(ctx context.Context, state *workflow.WorkflowState) error {
	config, err := json.Marshal(state.Configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO workflow_states (id, name, description, type, configuration, workflow_id, exit_workflow_id, entry_workflow_id) VALUES ($1, $2, $3, $4, ($5)::jsonb, $6, $7, $8)",
		state.Id, state.Name, state.Description, state.Type.String(), string(config), state.WorkflowId, state.ExitWorkflowId, state.EntryWorkflowId,
	)
	return err
}

func (ds *DataStore) GetWorkflowStates(ctx context.Context) ([]*workflow.WorkflowState, error) {
	records, err := ds.db.QueryContext(ctx, "SELECT id, name, description, type, configuration, workflow_id, exit_workflow_id, entry_workflow_id FROM workflow_states")
	if err != nil {
		return nil, err
	}
	defer records.Close()
	states := make([]*workflow.WorkflowState, 0)
	for records.Next() {
		var state workflow.WorkflowState
		var configuration json.RawMessage
		var stateType string
		err := records.Scan(&state.Id, &state.Name, &state.Description, &stateType, &configuration, &state.WorkflowId, &state.ExitWorkflowId, &state.EntryWorkflowId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, err
		}
		state.Type = workflow.WorkflowStateType(workflow.WorkflowStateType_value[stateType])
		err = json.Unmarshal(configuration, &state.Configuration)
		if err != nil {
			return nil, err
		}
		states = append(states, &state)
	}
	return states, nil
}

func (ds *DataStore) GetWorkflowState(ctx context.Context, stateId string) (*workflow.WorkflowState, error) {
	row := ds.db.QueryRowContext(ctx, "SELECT id, name, description, type, configuration, workflow_id, exit_workflow_id, entry_workflow_id FROM workflow_states WHERE id = $1", stateId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var state workflow.WorkflowState
	var configuration json.RawMessage
	var stateType string
	err := row.Scan(&state.Id, &state.Name, &state.Description, &stateType, &configuration, &state.WorkflowId, &state.ExitWorkflowId, &state.EntryWorkflowId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	state.Type = workflow.WorkflowStateType(workflow.WorkflowStateType_value[stateType])
	err = json.Unmarshal(configuration, &state.Configuration)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func (ds *DataStore) AddWorkflowTransition(ctx context.Context, transition *workflow.WorkflowStateTransition) error {
	_, err := ds.db.ExecContext(ctx, "INSERT INTO workflow_state_transitions (from_state_id, to_state_id, description) VALUES ($1, $2, $3)",
		transition.FromStateId, transition.ToStateId, transition.Description,
	)
	return err
}

func (ds *DataStore) AddWorkflowActivityStorageSystem(ctx context.Context, workflowActivityId int64, storageSystemId string, configuration map[string]string) error {
	config, err := json.Marshal(configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO workflow_activity_storage_systems (activity_id, storage_system_id, configuration) VALUES ($1, $2, ($3)::jsonb)",
		workflowActivityId, storageSystemId, string(config),
	)
	return err
}

func (ds *DataStore) AddWorkflowActivityInstancePrompt(ctx context.Context, workflowActivityId int64, promptId string, configuration map[string]string) error {
	config, err := json.Marshal(configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO workflow_activity_prompts (activity_id, prompt_id, configuration) VALUES ($1, $2, ($3)::jsonb)",
		workflowActivityId, promptId, string(config),
	)
	return err
}

func (ds *DataStore) AddWorkflowActivityModel(ctx context.Context, workflowActivityId int64, modelId string, configuration map[string]string) error {
	config, err := json.Marshal(configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO workflow_activity_models (activity_id, model_id, configuration) VALUES ($1, $2, ($3)::jsonb)",
		workflowActivityId, modelId, string(config),
	)
	return err
}

func (ds *DataStore) AddWorkflowActivity(ctx context.Context, workflowId string, activity *workflow.WorkflowActivity) (int64, error) {
	config, err := json.Marshal(activity.Configuration)
	if err != nil {
		return 0, err
	}

	rows, err := ds.db.QueryContext(ctx, "INSERT INTO workflow_activities (workflow_id, activity_id, execution_group, configuration) VALUES ($1, $2, $3, ($4)::jsonb) returning id",
		workflowId, activity.ActivityId, activity.ExecutionGroup, string(config),
	)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var workflowActivityId int64
	if rows.Next() {
		err = rows.Scan(&workflowActivityId)
		if err != nil {
			return 0, err
		}
	}
	stmtInput, err := ds.db.PrepareContext(ctx, "INSERT INTO workflow_activity_inputs (activity_id, name, value) VALUES ($1, $2, $3)")
	if err != nil {
		return 0, err
	}
	defer stmtInput.Close()

	for name, value := range activity.Inputs {
		_, err = stmtInput.ExecContext(ctx, workflowActivityId, name, value)
		if err != nil {
			return 0, err
		}
	}

	stmtOutput, err := ds.db.PrepareContext(ctx, "INSERT INTO workflow_activity_outputs (activity_id, name, value) VALUES ($1, $2, $3)")
	if err != nil {
		return 0, err
	}
	defer stmtOutput.Close()

	for name, value := range activity.Outputs {
		_, err = stmtOutput.ExecContext(ctx, workflowActivityId, name, value)
		if err != nil {
			return 0, err
		}
	}
	return workflowActivityId, nil
}

func (ds *DataStore) AddWorkflowExecution(ctx context.Context, txn *sql.Tx, parentExecutionId *string, workflowId string, metadataId *string, collectionId *string, executionContext *workflow.WorkflowExecutionContext) (string, error) {
	executionContext.CurrentExecutionGroup = -1
	executionContextJson, err := json.Marshal(executionContext)
	if err != nil {
		return "", err
	}
	result, err := txn.QueryContext(ctx, "insert into workflow_executions (parent_execution_id, workflow_id, metadata_id, collection_id, context) values ($1, $2, $3, $4, ($5)::jsonb) returning id::varchar", parentExecutionId, workflowId, metadataId, collectionId, executionContextJson)
	if err != nil {
		return "", err
	}
	defer result.Close()
	if result.Next() {
		var executionId string
		err = result.Scan(&executionId)
		if err != nil {
			return "", err
		}
		return executionId, nil
	}
	return "", errors.New("adding execution failed")
}

func (ds *DataStore) AreChildWorkflowExecutionsComplete(ctx context.Context, txn *sql.Tx, executionId string) (bool, error) {
	result, err := txn.QueryContext(ctx, "select case when failed is not null then null else completed end from workflow_executions where parent_execution_id = $1::uuid", executionId)
	if err != nil {
		return false, err
	}
	defer result.Close()
	for result.Next() {
		var completed *time.Time
		err = result.Scan(&completed)
		if err != nil {
			return false, err
		}
		if completed == nil {
			return false, nil
		}
	}
	return true, nil
}

func (ds *DataStore) IsWorkflowExecutionGroupComplete(ctx context.Context, txn *sql.Tx, executionId string) (bool, error) {
	result, err := txn.QueryContext(ctx, "select case when failed is not null then null else completed end from workflow_execution_jobs where workflow_execution_id = $1::uuid", executionId)
	if err != nil {
		return false, err
	}
	defer result.Close()
	for result.Next() {
		var completed *time.Time
		err = result.Scan(&completed)
		if err != nil {
			return false, err
		}
		if completed == nil {
			return false, nil
		}
	}
	return true, nil
}

func (ds *DataStore) SetWorkflowExecutionJobStatus(ctx context.Context, txn *sql.Tx, executionId, jobId string, success bool, complete bool, errorStr *string) (bool, *string, error) {
	if complete {
		if success {
			_, err := txn.ExecContext(ctx, "update workflow_execution_jobs set failed = null, completed = now(), error = null, worker_id = null where id = $1::uuid", jobId)
			if err != nil {
				return false, nil, err
			}
			return true, nil, nil
		} else {
			_, err := txn.ExecContext(ctx, "update workflow_execution_jobs set failed = now(), completed = now(), error = $1, worker_id = null where id = $2::uuid", errorStr, jobId)
			if err != nil {
				return false, nil, err
			}
			_, err = txn.ExecContext(ctx, "update workflow_executions set completed = now(), failed = now(), error = $1 where id = $2::uuid", errorStr, executionId)
			if err != nil {
				return false, nil, err
			}
			return false, errorStr, err
		}
	} else if !success {
		_, err := txn.ExecContext(ctx, "update workflow_execution_jobs set failed = now(), completed = null, error = $1, worker_id = null where id = $2::uuid", errorStr, jobId)
		if err != nil {
			return false, nil, err
		}
		return false, errorStr, err
	}
	return true, nil, nil
}

func (ds *DataStore) AddWorkflowExecutionJob(ctx context.Context, txn *sql.Tx, executionContext *workflow.WorkflowExecutionContext, activity *workflow.WorkflowActivity) error {
	slog.InfoContext(ctx, "adding workflow execution job", slog.String("executionId", executionContext.ExecutionId), slog.String("workflowId", executionContext.WorkflowId), slog.Any("parentExecutionId", executionContext.ParentExecutionId), slog.String("activityId", activity.ActivityId))
	scheduled := time.Now()
	executionContextContext, err := json.Marshal(executionContext.Context)
	if err != nil {
		return err
	}
	_, err = txn.QueryContext(ctx, "insert into workflow_execution_jobs (workflow_execution_id, workflow_activity_id, activity_id, queue, scheduled, context) values ($1::uuid, $2, $3, $4, $5, $6::jsonb)", executionContext.ExecutionId, activity.WorkflowActivityId, activity.ActivityId, activity.Queue, scheduled, executionContextContext)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStore) GetWorkflowExecutionContextForUpdate(ctx context.Context, txn *sql.Tx, executionId string) (*workflow.WorkflowExecutionContext, error) {
	result, err := txn.QueryContext(ctx, "select parent_execution_id, context from workflow_executions where id = $1::uuid for update", executionId)
	if err != nil {
		return nil, err
	}

	var parentExecutionId *string
	var executionContextJson json.RawMessage
	if result.Next() {
		err = result.Scan(&parentExecutionId, &executionContextJson)
		if err != nil {
			result.Close()
			return nil, err
		}
	} else {
		result.Close()
		return nil, errors.New("missing execution workflow")
	}
	result.Close()

	executionContext := &workflow.WorkflowExecutionContext{}
	err = json.Unmarshal(executionContextJson, &executionContext)
	if err != nil {
		return nil, err
	}
	executionContext.ParentExecutionId = parentExecutionId
	executionContext.ExecutionId = executionId
	return executionContext, nil
}

func (ds *DataStore) QueueNextWorkflowJobs(ctx context.Context, txn *sql.Tx, executionContext *workflow.WorkflowExecutionContext) (bool, []string, error) {
	slog.InfoContext(ctx, "queuing next workflow job", slog.String("executionId", executionContext.ExecutionId), slog.String("workflowId", executionContext.WorkflowId), slog.Any("parentExecutionId", executionContext.ParentExecutionId))
	groups := calculateWorkflowExecutionGroups(executionContext)
	executionContext.CurrentExecutionGroup++
	if int(executionContext.CurrentExecutionGroup) >= len(groups) {
		_, err := txn.ExecContext(ctx, "update workflow_executions set completed = now() where id = $1", executionContext.ExecutionId)
		if err != nil {
			return false, nil, err
		}
		return true, nil, nil
	} else {
		executionContextJson, err := json.Marshal(executionContext)
		if err != nil {
			return false, nil, err
		}

		_, err = txn.ExecContext(ctx, "update workflow_executions set context = $1::jsonb where id = $2::uuid", executionContextJson, executionContext.ExecutionId)
		if err != nil {
			return false, nil, err
		}

		activities := groups[executionContext.CurrentExecutionGroup]
		queues := make([]string, 0)
		if len(activities) == 0 {
			return false, nil, errors.New("empty activities in execution group")
		} else {
			for _, activity := range activities {
				err = ds.AddWorkflowExecutionJob(ctx, txn, executionContext, activity)
				if err != nil {
					return false, nil, err
				}
				queues = append(queues, activity.Queue)
			}
		}
		return false, queues, nil
	}
}
