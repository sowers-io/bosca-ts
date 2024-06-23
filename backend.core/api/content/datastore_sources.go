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
	"database/sql"
	"encoding/json"
)

func getSource(rows *sql.Rows) (*content.Source, error) {
	source := &content.Source{}
	var cfgStr string
	cfg := make(map[string]string)
	err := rows.Scan(&source.Id, &source.Name, &source.Description, &cfgStr)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(cfgStr), &cfg)
	if err != nil {
		return nil, err
	}
	return source, nil
}

func (ds *DataStore) GetSources(ctx context.Context) ([]*content.Source, error) {
	query := "select id, name, description, configuration::varchar from sources"
	rows, err := ds.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	sources := make([]*content.Source, 0)
	for rows.Next() {
		source, err := getSource(rows)
		if err != nil {
			return nil, err
		}
		sources = append(sources, source)
	}
	return sources, nil
}

func (ds *DataStore) GetSource(ctx context.Context, id string) (*content.Source, error) {
	rows, err := ds.db.QueryContext(ctx, "select id, name, description, configuration::varchar from sources where name = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	if rows.Next() {
		return getSource(rows)
	}
	rowsById, err := ds.db.QueryContext(ctx, "select id, name, description, configuration::varchar from sources where id = $1::uuid", id)
	if err != nil {
		return nil, err
	}
	defer rowsById.Close()
	if rowsById.Next() {
		return getSource(rowsById)
	}
	return nil, nil
}
