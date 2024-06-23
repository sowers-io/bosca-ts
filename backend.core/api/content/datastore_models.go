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
	model "bosca.io/api/protobuf/bosca/content"
	"context"
	"encoding/json"
)

func (ds *DataStore) AddModel(ctx context.Context, model *model.Model) (string, error) {
	config, err := json.Marshal(model.Configuration)
	if err != nil {
		return "", err
	}
	rows, err := ds.db.QueryContext(ctx, "INSERT INTO models (type, name, description, configuration) VALUES ($1, $2, $3, ($4)::jsonb) returning id::varchar",
		model.Type, model.Name, model.Description, string(config),
	)
	if err != nil {
		return "", err
	}
	defer rows.Close()
	if rows.Next() {
		var id string
		err = rows.Scan(&id)
		return id, err
	}
	return "", rows.Err()
}

func (ds *DataStore) GetModels(ctx context.Context) ([]*model.Model, error) {
	results, err := ds.db.QueryContext(ctx, "select id::varchar, type, name, description, configuration from models")
	if err != nil {
		return nil, err
	}
	defer results.Close()
	models := make([]*model.Model, 0)
	for results.Next() {
		m := &model.Model{}
		var configurationJson json.RawMessage
		err = results.Scan(&m.Id, &m.Type, &m.Name, &m.Description, &configurationJson)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(configurationJson, &m.Configuration)
		if err != nil {
			return nil, err
		}
		models = append(models, m)
	}
	return models, nil
}

func (ds *DataStore) GetModel(ctx context.Context, id string) (*model.Model, error) {
	results, err := ds.db.QueryContext(ctx, "select id, type, name, description, configuration from models where id = $1::uuid", id)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	if results.Next() {
		m := &model.Model{}
		var configurationJson json.RawMessage
		err = results.Scan(&m.Id, &m.Type, &m.Name, &m.Description, &configurationJson)
		if err != nil {
			return nil, err
		}
		err = json.Unmarshal(configurationJson, &m.Configuration)
		if err != nil {
			return nil, err
		}
		return m, nil
	}
	return nil, nil
}

func (ds *DataStore) GetWorkflowActivityInstanceModels(ctx context.Context, instanceId int64) ([]*model.WorkflowActivityModel, error) {
	rows, err := ds.db.QueryContext(ctx, "select model_id, configuration from workflow_activity_instance_models where instance_id = $1", instanceId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	models := make([]*model.WorkflowActivityModel, 0)
	for rows.Next() {
		var id string
		var configuration json.RawMessage
		err = rows.Scan(&id, &configuration)
		m, err := ds.GetModel(ctx, id)
		if err != nil {
			return nil, err
		}
		activityModel := &model.WorkflowActivityModel{
			Model: m,
		}
		err = json.Unmarshal(configuration, &activityModel.Configuration)
		if err != nil {
			return nil, err
		}
		models = append(models, activityModel)
	}
	return models, nil
}
