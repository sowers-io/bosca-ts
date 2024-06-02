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
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/security/identity"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
)

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
