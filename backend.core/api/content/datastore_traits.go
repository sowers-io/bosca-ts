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
	"context"
)

func (ds *DataStore) AddTrait(ctx context.Context, trait *content.Trait) error {
	_, err := ds.db.ExecContext(ctx, "INSERT INTO traits (id, name, description) VALUES ($1, $2, $3)",
		trait.Id, trait.Name, trait.Description,
	)
	if err != nil {
		return err
	}
	stmtInput, err := ds.db.PrepareContext(ctx, "INSERT INTO trait_workflows (trait_id, workflow_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmtInput.Close()
	for _, workflowId := range trait.WorkflowIds {
		_, err = stmtInput.ExecContext(ctx, trait.Id, workflowId)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ds *DataStore) GetTraits(ctx context.Context) ([]*content.Trait, error) {
	query := "select id, name, description from traits"
	rows, err := ds.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	traits := make([]*content.Trait, 0)
	for rows.Next() {
		trait := &content.Trait{}
		err := rows.Scan(&trait.Id, &trait.Name, &trait.Description)
		if err != nil {
			return nil, err
		}
		traitWorkflowIds, err := ds.GetTraitWorkflowIds(ctx, trait.Id)
		if err != nil {
			return nil, err
		}
		trait.WorkflowIds = traitWorkflowIds
		traits = append(traits, trait)
	}
	return traits, nil
}

func (ds *DataStore) GetTraitWorkflowIds(ctx context.Context, id string) ([]string, error) {
	query := "select workflow_id from trait_workflows where trait_id = $1"
	rows, err := ds.db.QueryContext(ctx, query, id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ids := make([]string, 0)
	for rows.Next() {
		var workflowId string
		err := rows.Scan(&workflowId)
		if err != nil {
			return nil, err
		}
		ids = append(ids, workflowId)
	}
	return ids, nil
}

func (ds *DataStore) GetWorkflowActivityStorageSystemIds(ctx context.Context, workflowId, activityId string) ([]string, error) {
	query := "select storage_system_id from workflow_activity_storage_systems where workflow_id = $1 and activity_id = $2"
	rows, err := ds.db.QueryContext(ctx, query, workflowId, activityId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	ids := make([]string, 0)
	for rows.Next() {
		var storageSystemId string
		err := rows.Scan(&storageSystemId)
		if err != nil {
			return nil, err
		}
		ids = append(ids, storageSystemId)
	}
	return ids, nil
}
