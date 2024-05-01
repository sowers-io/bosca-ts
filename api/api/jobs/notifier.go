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

package jobs

import (
	"context"
	"errors"
	"github.com/jackc/pgx/v5/pgxpool"
	"log"
	"time"
)

type JobNotifier interface {
	Notify(context.Context, string)
	WaitForNotification(context.Context, string)
}

type notifier struct {
	pool *pgxpool.Pool
}

func NewNotifier(pool *pgxpool.Pool) JobNotifier {
	return &notifier{
		pool: pool,
	}
}

func (n *notifier) Notify(ctx context.Context, id string) {
	conn, err := n.pool.Acquire(ctx)
	if err != nil {
		log.Printf("failed to acquire connection: %v", err)
		return
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "select pg_notify($1, '')", id)
	if err != nil {
		log.Printf("failed to notify: %v", err)
	}
}

func (n *notifier) WaitForNotification(ctx context.Context, id string) {
	conn, err := n.pool.Acquire(ctx)
	if err != nil {
		log.Printf("failed to acquire connection: %v", err)
		return
	}
	defer conn.Release()
	_, err = conn.Exec(ctx, "select listen($1)", id)
	if err != nil {
		log.Printf("failed to listen: %v", err)
	}
	defer func() {
		_, _ = conn.Exec(ctx, "select unlisten($1)", id)
	}()
	timeout, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()
	_, err = conn.Conn().WaitForNotification(timeout)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Print("timed out waiting on notification, continuing")
		} else {
			log.Printf("failed waiting for notification: %v", err)
		}
	} else {
		log.Println("received notification, continuing")
	}
}
