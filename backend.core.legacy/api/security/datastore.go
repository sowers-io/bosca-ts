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

package security

import (
	grpc "bosca.io/api/protobuf/bosca/security"
	"context"
	"database/sql"
)

type DataStore struct {
	db *sql.DB
}

func NewDataStore(db *sql.DB) *DataStore {
	return &DataStore{
		db,
	}
}

func (ds *DataStore) GetGroups(ctx context.Context) ([]*grpc.Group, error) {
	result, err := ds.db.QueryContext(ctx, "select name, description from groups where enabled = true")
	if err != nil {
		return nil, err
	}
	defer result.Close()
	groups := make([]*grpc.Group, 0)

	for result.Next() {
		var group grpc.Group
		err = result.Scan(&group.Name, &group.Description)
		if err != nil {
			return nil, err
		}
		groups = append(groups, &group)
	}

	return groups, nil
}

func (ds *DataStore) AddGroup(ctx context.Context, name, description string) error {
	tx, err := ds.db.BeginTx(ctx, &sql.TxOptions{Isolation: sql.LevelReadCommitted})
	if err != nil {
		return err
	}
	stmt, err := ds.db.Prepare("insert into groups (name, description) values ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(name, description)
	if err != nil {
		return err
	}
	return tx.Commit()
}
