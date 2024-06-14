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
	"database/sql"
	"encoding/json"
)

func getStorageSystem(results *sql.Rows) (*model.StorageSystem, error) {
	i := &model.StorageSystem{}
	var configurationJson json.RawMessage
	var storageType string
	err := results.Scan(&i.Id, &i.Name, &i.Description, &storageType, &configurationJson)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal(configurationJson, &i.Configuration)
	if err != nil {
		return nil, err
	}
	switch storageType {
	case "vector":
		i.Type = model.StorageSystemType_vector_storage_system
	case "search":
		i.Type = model.StorageSystemType_search_storage_system
	case "metadata":
		i.Type = model.StorageSystemType_metadata_storage_system
	case "supplementary":
		i.Type = model.StorageSystemType_supplementary_storage_system
	default:
		i.Type = model.StorageSystemType_unknown_storage_system
	}
	return i, nil
}

func (ds *DataStore) GetStorageSystems(ctx context.Context) ([]*model.StorageSystem, error) {
	results, err := ds.db.QueryContext(ctx, "select id::varchar, name, description, type, configuration from storage_systems")
	if err != nil {
		return nil, err
	}
	defer results.Close()
	systems := make([]*model.StorageSystem, 0)
	for results.Next() {
		system, err := getStorageSystem(results)
		if err != nil {
			return nil, err
		}
		systems = append(systems, system)
	}
	return systems, nil
}

func (ds *DataStore) GetStorageSystem(ctx context.Context, id string) (*model.StorageSystem, error) {
	results, err := ds.db.QueryContext(ctx, "select id::varchar, name, description, type, configuration from storage_systems where id = $1::uuid", id)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	if results.Next() {
		return getStorageSystem(results)
	}
	return nil, nil
}

func (ds *DataStore) GetStorageSystemModels(ctx context.Context, id string) ([]*model.StorageSystemModel, error) {
	results, err := ds.db.QueryContext(ctx, "select model_id::varchar, type from storage_system_models where system_id = $1::uuid", id)
	if err != nil {
		return nil, err
	}
	defer results.Close()
	systems := make([]*model.StorageSystemModel, 0)
	for results.Next() {
		i := &model.StorageSystemModel{}
		var modelId string
		err = results.Scan(&modelId, &i.Type)
		if err != nil {
			return nil, err
		}
		i.Model, err = ds.GetModel(ctx, modelId)
		if err != nil {
			return nil, err
		}
		systems = append(systems, i)
	}
	return systems, nil
}
