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
	"bosca.io/api/protobuf/content"
	"bosca.io/tests"
	"context"
	_ "github.com/pressly/goose/v3"
	"testing"
)

func newCollection(t *testing.T, ctx context.Context, ds *DataStore) string {
	col := &content.Collection{
		Name: "MD 1",
		Type: content.CollectionType_root,
	}

	id, err := ds.AddCollection(ctx, col)
	if err != nil {
		panic(err)
	}
	if len(id) == 0 {
		t.Error("Expected id to not be empty")
	}

	err = ds.SetCollectionWorkflowStateId(ctx, id, "published")
	if err != nil {
		panic(err)
	}
	return id
}

func newMetadata(t *testing.T, ctx context.Context, ds *DataStore) string {
	metadata := &content.Metadata{
		Name:        "MD 1",
		ContentType: "text/plain",
	}

	id, err := ds.AddMetadata(ctx, metadata)
	if err != nil {
		panic(err)
	}
	if len(id) == 0 {
		t.Error("Expected id to not be empty")
	}

	//err = ds.SetMetadataStatus(ctx, id, content.MetadataStatus_ready)
	//if err != nil {
	//	panic(err)
	//}
	//
	//err = ds.SetMetadataWorkflowStateId(ctx, id, "published")
	//if err != nil {
	//	panic(err)
	//}
	return id
}

func Test_AddMetadata(t *testing.T) {
	db := tests.Up("content")
	defer tests.Down(db)

	ds := NewDataStore(db)
	newMetadata(t, context.Background(), ds)
}

func Test_AddCollection(t *testing.T) {
	db := tests.Up("content")
	defer tests.Down(db)

	ds := NewDataStore(db)
	newCollection(t, context.Background(), ds)
}

func Test_AddCollectionItems(t *testing.T) {
	db := tests.Up("content")
	defer tests.Down(db)

	ds := NewDataStore(db)
	parent := newCollection(t, context.Background(), ds)
	childCollection := newCollection(t, context.Background(), ds)

	err := ds.AddCollectionCollectionItems(context.Background(), parent, []string{childCollection})
	if err != nil {
		t.Error(err)
	}

	childMetadata := newMetadata(t, context.Background(), ds)

	err = ds.AddCollectionMetadataItems(context.Background(), parent, []string{childMetadata})
	if err != nil {
		t.Error(err)
	}
}
