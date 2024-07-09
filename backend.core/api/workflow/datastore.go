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
	"context"
	"database/sql"
	"sync"
)

type DataStore struct {
	db *sql.DB

	executionChannels     map[string]chan error
	queueChannels         map[string]chan string
	executionChannelMutex sync.Mutex
}

func NewDataStore(db *sql.DB) *DataStore {
	return &DataStore{
		db:                db,
		executionChannels: make(map[string]chan error),
		queueChannels:     make(map[string]chan string),
	}
}

func (ds *DataStore) NewTransaction(ctx context.Context) (*sql.Tx, error) {
	return ds.db.BeginTx(ctx, nil)
}
