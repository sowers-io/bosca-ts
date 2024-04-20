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
	cfg := configuration.NewServerConfiguration()
	cfg.DatabaseConnectionString = "postgresql://bosca:bosca@localhost:5432/boscatest"
	if pool, err := datastore.NewDatabasePool(context.Background(), cfg); err == nil {
		db := stdlib.OpenDBFromPool(pool)
		goose.SetBaseFS(os.DirFS("../../../database/" + name))
		if err := goose.SetDialect("postgres"); err != nil {
			panic(err)
		}
		if err := goose.Up(db, "."); err != nil {
			panic(err)
		}
		return db
	} else {
		panic(err)
	}
}

func Down(db *sql.DB) {
	err := goose.Down(db, ".")
	if err != nil {
		panic(err)
	}
}
