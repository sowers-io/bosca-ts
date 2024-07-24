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

package tests

import (
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/datastore"
	"context"
	"database/sql"
	"github.com/jackc/pgx/v5/stdlib"
	"github.com/pressly/goose/v3"
	"os"
)

func Up(name string) *sql.DB {
	err := os.Setenv("BOSCA_TESTS_CONNECTION_STRING", "postgresql://bosca:bosca@localhost:5432/boscatest")
	if err != nil {
		panic(err)
	}
	cfg := configuration.NewServerConfiguration("tests", 6001, 6011)
	if pool, err := datastore.NewDatabasePool(context.Background(), cfg); err == nil {
		db := stdlib.OpenDBFromPool(pool)
		goose.SetBaseFS(os.DirFS("../../../database/" + name))
		if err := goose.SetDialect("postgres"); err != nil {
			panic(err)
		}
		_ = goose.Reset(db, ".")
		if err := goose.Up(db, "."); err != nil {
			panic(err)
		}
		return db
	} else {
		panic(err)
	}
}

func Down(db *sql.DB) {
	err := goose.Reset(db, ".")
	if err != nil {
		panic(err)
	}
}
