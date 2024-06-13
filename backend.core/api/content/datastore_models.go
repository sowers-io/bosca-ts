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
