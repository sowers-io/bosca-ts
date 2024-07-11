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
	"errors"
	"log/slog"
	"sync"
	"time"
)

type DataStore struct {
	db *sql.DB

	executionChannels     map[string]chan *ExecutionNotification
	queueChannels         map[string]chan *ExecutionNotification
	executionChannelMutex sync.Mutex
}

func NewDataStore(db *sql.DB) *DataStore {
	ds := &DataStore{
		db:                db,
		executionChannels: make(map[string]chan *ExecutionNotification),
		queueChannels:     make(map[string]chan *ExecutionNotification),
	}
	go func() {
		ctx := context.Background()
		for {
			allQueues, err := ds.GetAllQueues(ctx)
			if err != nil {
				slog.Error("failed to get queues", slog.Any("error", err))
			}

			// warm up any connections
			err = ds.NotifyJobAvailable(ctx, allQueues)
			if err != nil {
				slog.Error("failed to notify queues", slog.Any("error", err))
			}

			ctxTimeout, ctxCancel := context.WithTimeout(ctx, 30 * time.Second)
			err = ds.ListenForCompletions(ctxTimeout)
			if err != nil && !errors.Is(err, context.DeadlineExceeded) {
				slog.Error("failed to listen for completions", slog.Any("error", err))
			}
			ctxCancel()
		}
	}()
	return ds
}

func (ds *DataStore) NewTransaction(ctx context.Context) (*sql.Tx, error) {
	return ds.db.BeginTx(ctx, nil)
}
