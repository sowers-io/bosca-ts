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

package content

import (
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/pkg/security/identity"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
)

func (ds *DataStore) AddWorkflow(ctx context.Context, workflow *content.Workflow) error {
	config, err := json.Marshal(workflow.Configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO workflows (id, name, description, queue, configuration) VALUES ($1, $2, $3, $4, ($5)::jsonb)",
		workflow.Id, workflow.Name, workflow.Description, workflow.Queue, string(config),
	)
	return err
}

func (ds *DataStore) GetWorkflow(ctx context.Context, id string) (*content.Workflow, error) {
	row := ds.db.QueryRowContext(ctx, "SELECT id, name, description, queue, configuration FROM workflows WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var workflow content.Workflow
	var configuration json.RawMessage
	err := row.Scan(&workflow.Id, &workflow.Name, &workflow.Description, &workflow.Queue, &configuration)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	err = json.Unmarshal(configuration, &workflow.Configuration)
	if err != nil {
		return nil, err
	}
	return &workflow, nil
}

func getWorkflowActivityParameterValue(inputs *sql.Rows) (string, *content.WorkflowActivityParameterValue, error) {
	var name string
	var value json.RawMessage
	input := &content.WorkflowActivityParameterValue{}
	err := inputs.Scan(&name, &value)
	if err != nil {
		return "", nil, err
	}
	err = json.Unmarshal(value, &input)
	if err != nil {
		return name, nil, err
	}
	return name, input, nil
}

func (ds *DataStore) GetWorkflowActivityInstances(ctx context.Context, workflowId string) ([]*content.WorkflowActivityInstance, error) {
	rows, err := ds.db.QueryContext(ctx, "SELECT wi.id, wi.activity_id, wa.child_workflow, wa.child_workflow_queue, execution_group, wi.configuration FROM workflow_activity_instances wi inner join workflow_activities wa on (wi.activity_id = wa.id) WHERE workflow_id = $1", workflowId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	instances := make([]*content.WorkflowActivityInstance, 0)
	for rows.Next() {
		instance := &content.WorkflowActivityInstance{
			Inputs:  make(map[string]*content.WorkflowActivityParameterValue),
			Outputs: make(map[string]*content.WorkflowActivityParameterValue),
		}

		var configuration json.RawMessage
		err = rows.Scan(&instance.InstanceId, &instance.Id, &instance.ChildWorkflow, &instance.ChildWorkflowQueue, &instance.ExecutionGroup, &configuration)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(configuration, &instance.Configuration)
		if err != nil {
			return nil, err
		}

		inputs, err := ds.db.QueryContext(ctx, "select name, value::varchar from workflow_activity_instance_inputs where instance_id = $1", instance.InstanceId)
		if err != nil {
			return nil, err
		}
		defer inputs.Close()
		for inputs.Next() {
			name, input, err := getWorkflowActivityParameterValue(inputs)
			if err != nil {
				return nil, err
			}
			instance.Inputs[name] = input
		}

		outputs, err := ds.db.QueryContext(ctx, "select name, value::varchar from workflow_activity_instance_outputs where instance_id = $1", instance.InstanceId)
		if err != nil {
			return nil, err
		}
		defer outputs.Close()
		for outputs.Next() {
			name, output, err := getWorkflowActivityParameterValue(outputs)
			if err != nil {
				return nil, err
			}
			instance.Outputs[name] = output
		}

		instance.Prompts, err = ds.GetWorkflowActivityInstancePrompts(ctx, instance.InstanceId)
		if err != nil {
			return nil, err
		}
		instance.StorageSystems, err = ds.GetWorkflowActivityInstanceStorageSystems(ctx, instance.InstanceId)
		if err != nil {
			return nil, err
		}
		instance.Models, err = ds.GetWorkflowActivityInstanceModels(ctx, instance.InstanceId)
		if err != nil {
			return nil, err
		}

		instances = append(instances, instance)
	}
	return instances, nil
}

func (ds *DataStore) GetWorkflows(ctx context.Context) ([]*content.Workflow, error) {
	result, err := ds.db.QueryContext(ctx, "SELECT id, name, description, queue, configuration FROM workflows")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	workflows := make([]*content.Workflow, 0)
	for result.Next() {
		var workflow content.Workflow
		var configuration json.RawMessage
		err := result.Scan(&workflow.Id, &workflow.Name, &workflow.Description, &workflow.Queue, &configuration)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, err
		}
		err = json.Unmarshal(configuration, &workflow.Configuration)
		if err != nil {
			return nil, err
		}
		workflows = append(workflows, &workflow)
	}
	return workflows, nil
}

func (ds *DataStore) AddWorkflowActivity(ctx context.Context, activity *content.WorkflowActivity) error {
	config, err := json.Marshal(activity.Configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO workflow_activities (id, name, description, child_workflow, child_workflow_queue, configuration) VALUES ($1, $2, $3, $4, $5, ($6)::jsonb)",
		activity.Id, activity.Name, activity.Description, activity.ChildWorkflow, activity.ChildWorkflowQueue, string(config),
	)
	if err != nil {
		return err
	}

	stmtInput, err := ds.db.PrepareContext(ctx, "INSERT INTO workflow_activity_inputs (activity_id, name, type) VALUES ($1, $2, $3)")
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

	stmtOutput, err := ds.db.PrepareContext(ctx, "INSERT INTO workflow_activity_outputs (activity_id, name, type) VALUES ($1, $2, $3)")
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

func (ds *DataStore) GetWorkflowTransition(ctx context.Context, fromStateId string, toStateId string) (*content.WorkflowStateTransition, error) {
	row := ds.db.QueryRowContext(ctx, "SELECT from_state_id, to_state_id FROM workflow_state_transitions WHERE from_state_id = $1 AND to_state_id = $2", fromStateId, toStateId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var transition content.WorkflowStateTransition
	err := row.Scan(&transition.FromStateId, &transition.ToStateId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &transition, nil
}

func (ds *DataStore) AddWorkflowState(ctx context.Context, state *content.WorkflowState) error {
	config, err := json.Marshal(state.Configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO workflow_states (id, name, description, type, configuration, workflow_id, exit_workflow_id, entry_workflow_id) VALUES ($1, $2, $3, $4, ($5)::jsonb, $6, $7, $8)",
		state.Id, state.Name, state.Description, state.Type.String(), string(config), state.WorkflowId, state.ExitWorkflowId, state.EntryWorkflowId,
	)
	return err
}

func (ds *DataStore) GetWorkflowStates(ctx context.Context) ([]*content.WorkflowState, error) {
	records, err := ds.db.QueryContext(ctx, "SELECT id, name, description, type, configuration, workflow_id, exit_workflow_id, entry_workflow_id FROM workflow_states")
	if err != nil {
		return nil, err
	}
	defer records.Close()
	states := make([]*content.WorkflowState, 0)
	for records.Next() {
		var state content.WorkflowState
		var configuration json.RawMessage
		var stateType string
		err := records.Scan(&state.Id, &state.Name, &state.Description, &stateType, &configuration, &state.WorkflowId, &state.ExitWorkflowId, &state.EntryWorkflowId)
		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, err
		}
		state.Type = content.WorkflowStateType(content.WorkflowStateType_value[stateType])
		err = json.Unmarshal(configuration, &state.Configuration)
		if err != nil {
			return nil, err
		}
		states = append(states, &state)
	}
	return states, nil
}

func (ds *DataStore) GetWorkflowState(ctx context.Context, id string) (*content.WorkflowState, error) {
	row := ds.db.QueryRowContext(ctx, "SELECT id, name, description, type, configuration, workflow_id, exit_workflow_id, entry_workflow_id FROM workflow_states WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var state content.WorkflowState
	var configuration json.RawMessage
	var stateType string
	err := row.Scan(&state.Id, &state.Name, &state.Description, &stateType, &configuration, &state.WorkflowId, &state.ExitWorkflowId, &state.EntryWorkflowId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	state.Type = content.WorkflowStateType(content.WorkflowStateType_value[stateType])
	err = json.Unmarshal(configuration, &state.Configuration)
	if err != nil {
		return nil, err
	}
	return &state, nil
}

func (ds *DataStore) TransitionMetadataWorkflowStateId(ctx context.Context, metadata *content.Metadata, toState *content.WorkflowState, status string, success bool, complete bool) error {
	subjectId, err := identity.GetSubjectId(ctx)
	if err != nil {
		return err
	}

	txn, err := ds.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = txn.ExecContext(ctx, "insert into metadata_workflow_transition_history (metadata_id, from_state_id, to_state_id, subject, status, success, complete) values ($1::uuid, $2, $3, $4, $5, $6, $7)", metadata.Id, metadata.WorkflowStateId, toState.Id, subjectId, status, success, complete)
	if err != nil {
		txn.Rollback()
		return err
	}
	if !success {
		_, err = txn.ExecContext(ctx, "update metadata set workflow_state_pending_id = null where id = $1::uuid", metadata.Id)
		if err != nil {
			txn.Rollback()
			return err
		}
	} else {
		if complete {
			_, err = txn.ExecContext(ctx, "update metadata set workflow_state_id = $1, workflow_state_pending_id = null where id = $2::uuid", toState.Id, metadata.Id)
			if err != nil {
				txn.Rollback()
				return err
			}
		} else {
			_, err = txn.ExecContext(ctx, "update metadata set workflow_state_pending_id = $1 where id = $2::uuid", toState.Id, metadata.Id)
			if err != nil {
				txn.Rollback()
				return err
			}
		}
	}
	return txn.Commit()
}

func (ds *DataStore) AddWorkflowTransition(ctx context.Context, transition *content.WorkflowStateTransition) error {
	_, err := ds.db.ExecContext(ctx, "INSERT INTO workflow_state_transitions (from_state_id, to_state_id, description) VALUES ($1, $2, $3)",
		transition.FromStateId, transition.ToStateId, transition.Description,
	)
	return err
}

func (ds *DataStore) AddWorkflowActivityInstanceStorageSystem(ctx context.Context, instanceId int64, storageSystemId string, configuration map[string]string) error {
	config, err := json.Marshal(configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO workflow_activity_instance_storage_systems (instance_id, storage_system_id, configuration) VALUES ($1, $2, ($3)::jsonb)",
		instanceId, storageSystemId, string(config),
	)
	return err
}

func (ds *DataStore) AddWorkflowActivityInstancePrompt(ctx context.Context, instanceId int64, promptId string, configuration map[string]string) error {
	config, err := json.Marshal(configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO workflow_activity_instance_prompts (instance_id, prompt_id, configuration) VALUES ($1, $2, ($3)::jsonb)",
		instanceId, promptId, string(config),
	)
	return err
}

func (ds *DataStore) AddWorkflowActivityInstanceModel(ctx context.Context, instanceId int64, modelId string, configuration map[string]string) error {
	config, err := json.Marshal(configuration)
	if err != nil {
		return err
	}
	_, err = ds.db.ExecContext(ctx, "INSERT INTO workflow_activity_instance_models (instance_id, model_id, configuration) VALUES ($1, $2, ($3)::jsonb)",
		instanceId, modelId, string(config),
	)
	return err
}

func (ds *DataStore) AddWorkflowActivityInstance(ctx context.Context, workflowId string, instance *content.WorkflowActivityInstance) (int64, error) {
	config, err := json.Marshal(instance.Configuration)
	if err != nil {
		return 0, err
	}

	rows, err := ds.db.QueryContext(ctx, "INSERT INTO workflow_activity_instances (workflow_id, activity_id, execution_group, configuration) VALUES ($1, $2, $3, ($4)::jsonb) returning id",
		workflowId, instance.Id, instance.ExecutionGroup, string(config),
	)
	if err != nil {
		return 0, err
	}
	defer rows.Close()
	var id int64
	if rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return 0, err
		}
	}
	stmtInput, err := ds.db.PrepareContext(ctx, "INSERT INTO workflow_activity_instance_inputs (instance_id, name, value) VALUES ($1, $2, ($3)::jsonb)")
	if err != nil {
		return 0, err
	}
	defer stmtInput.Close()

	for name, parameterType := range instance.Inputs {
		str, err := json.Marshal(parameterType)
		if err != nil {
			return 0, err
		}
		_, err = stmtInput.ExecContext(ctx, id, name, string(str))
		if err != nil {
			return 0, err
		}
	}

	stmtOutput, err := ds.db.PrepareContext(ctx, "INSERT INTO workflow_activity_instance_outputs (instance_id, name, value) VALUES ($1, $2, $3)")
	if err != nil {
		return 0, err
	}
	defer stmtOutput.Close()

	for name, parameterType := range instance.Outputs {
		str, err := json.Marshal(parameterType)
		if err != nil {
			return 0, err
		}
		_, err = stmtOutput.ExecContext(ctx, id, name, string(str))
		if err != nil {
			return 0, err
		}
	}
	return id, nil
}
