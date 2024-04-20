package profiles

import (
	"bosca.io/api/protobuf/profiles"
	"bosca.io/tests"
	"context"
	_ "github.com/pressly/goose/v3"
	"testing"
)

func Test_AddProfile(t *testing.T) {
	db := tests.Up("profiles")
	defer tests.Down(db)

	profile := &profiles.Profile{
		Id:   "ASDF1234",
		Name: "User's Name",
	}

	ds := NewDataStore(db)
	err := ds.AddProfile(context.Background(), profile)
	if err != nil {
		panic(err)
	}
}
