package profiles

import (
	"bosca.io/api/protobuf/profiles"
	"context"
	"database/sql"
	"embed"
	"github.com/pressly/goose/v3"
	_ "github.com/pressly/goose/v3"
	"testing"
)

//go:embed ../../../database/profiles/*.sql
var embedMigrations embed.FS

func Up() *sql.DB {
	if db, err := sql.Open("postgres", ""); err == nil {
		goose.SetBaseFS(embedMigrations)
		if err := goose.SetDialect("postgres"); err != nil {
			panic(err)
		}
		if err := goose.Up(db, "../../../database/profiles"); err != nil {
			panic(err)
		}
		return db
	} else {
		panic(err)
	}
}

func Down(db *sql.DB) {
	err := goose.Down(db, "../../../database/profiles")
	if err != nil {
		panic(err)
	}
}

func Test_AddProfile(t *testing.T) {
	db := Up()
	defer Down(db)

	profile := &profiles.Profile{}

	datastore := NewDataStore(db)
	err := datastore.AddProfile(context.Background(), profile)
	if err != nil {
		panic(err)
	}
}
