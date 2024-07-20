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
)

type DataStore struct {
	db     *sql.DB
	pubsub *PubSub
}

func NewDataStore(db *sql.DB, pubsub *PubSub) *DataStore {
	ds := &DataStore{
		db:     db,
		pubsub: pubsub,
	}
	return ds
}

func (ds *DataStore) NewTransaction(ctx context.Context) (*sql.Tx, error) {
	return ds.db.BeginTx(ctx, nil)
}

func (ds *DataStore) WaitForExecutionCompletion(ctx context.Context, queue string, executionId string) (*workflow.WorkflowExecutionNotification, error) {
	var notification *workflow.WorkflowExecutionNotification
	err := ds.pubsub.Listen(ctx, queue, func(e *workflow.WorkflowExecutionNotification) (bool, error) {
		if e.Type == workflow.WorkflowExecutionNotificationType_execution_completion && e.ExecutionId == executionId {
			notification = e
			return false, nil
		}
		return true, nil
	})
	return notification, err
}
