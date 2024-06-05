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
	"context"
	"database/sql"
	"errors"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"time"
)

func (ds *DataStore) AddRootCollection(ctx context.Context) (bool, error) {
	root, err := ds.GetCollection(ctx, RootCollectionId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		slog.Error("failed to get workflow for root collection", slog.Any("error", err))
		return false, err
	}
	if root != nil {
		return false, nil
	}
	id, err := uuid.Parse(RootCollectionId)
	if err != nil {
		return false, err
	}
	_, err = ds.db.ExecContext(ctx, "insert into collections (id, name, type, workflow_state_id) values ($1, 'Root', 'root', 'published')", id)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (ds *DataStore) GetCollectionCollectionItemIds(ctx context.Context, collectionId string) ([]string, error) {
	stmt, err := ds.db.PrepareContext(ctx, "SELECT child_id FROM collection_collection_items WHERE collection_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, collectionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ids := make([]string, 0)
	id := ""
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (ds *DataStore) GetCollectionMetadataItemIds(ctx context.Context, collectionId string) ([]string, error) {
	stmt, err := ds.db.PrepareContext(ctx, "SELECT metadata_id FROM collection_metadata_items WHERE collection_id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.QueryContext(ctx, collectionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ids := make([]string, 0)
	id := ""
	for rows.Next() {
		err = rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (ds *DataStore) GetCollection(ctx context.Context, id string) (*content.Collection, error) {
	var collection content.Collection

	m := pgtype.NewMap()

	var created time.Time
	var modified time.Time
	var collectionType string
	var tags []string
	var attributes string
	err := ds.db.QueryRowContext(ctx, "SELECT name, type, tags, attributes, created, modified FROM collections WHERE id = $1", id).Scan(
		&collection.Name,
		&collectionType,
		m.SQLScanner(&tags),
		&attributes,
		&created,
		&modified,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	switch collectionType {
	case "root":
		collection.Type = content.CollectionType_root
	case "folder":
		collection.Type = content.CollectionType_folder
	case "standard":
		collection.Type = content.CollectionType_standard
	}

	collection.Id = id
	collection.Tags = tags
	collection.Created = timestamppb.New(created)
	collection.Modified = timestamppb.New(modified)
	//collection.Attributes = attributes
	return &collection, nil
}

func (ds *DataStore) AddCollection(ctx context.Context, collection *content.Collection) (string, error) {
	stmt, err := ds.db.PrepareContext(ctx, "INSERT INTO collections (name, type, tags, attributes) VALUES ($1, $2, $3, ($4)::jsonb) returning id")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	tags := collection.Tags
	if tags == nil {
		tags = make([]string, 0)
	}
	attributes := collection.Attributes
	if attributes == nil {
		attributes = make(map[string]string)
	}

	result := stmt.QueryRowContext(ctx,
		collection.Name,
		collection.Type.String(),
		tags,
		attributes,
	)
	if result.Err() != nil {
		return "", result.Err()
	}

	var id string
	err = result.Scan(&id)
	if err != nil {
		return "", err
	}
	return id, nil
}

func (ds *DataStore) DeleteCollection(ctx context.Context, id string) error {
	stmt, err := ds.db.PrepareContext(ctx, "DELETE FROM collections WHERE id = $1::uuid")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	return err
}

func (ds *DataStore) AddCollectionCollectionItems(ctx context.Context, collectionId string, collectionIds []string) error {
	stmt, err := ds.db.PrepareContext(ctx, "INSERT INTO collection_collection_items (collection_id, child_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, id := range collectionIds {
		_, err = stmt.ExecContext(ctx, collectionId, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ds *DataStore) RemoveCollectionCollectionItems(ctx context.Context, collectionId string, collectionIds []string) error {
	stmt, err := ds.db.PrepareContext(ctx, "DELETE FROM collection_collection_items WHERE collection_id = $1 AND child_id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, id := range collectionIds {
		_, err = stmt.ExecContext(ctx, collectionId, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ds *DataStore) AddCollectionMetadataItems(ctx context.Context, collectionId string, metadataIds []string) error {
	stmt, err := ds.db.PrepareContext(ctx, "INSERT INTO collection_metadata_items (collection_id, metadata_id) VALUES ($1, $2)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, id := range metadataIds {
		_, err = stmt.ExecContext(ctx, collectionId, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ds *DataStore) RemoveCollectionMetadataItem(ctx context.Context, collectionId string, metadataIds []string) error {
	stmt, err := ds.db.PrepareContext(ctx, "DELETE FROM collection_metadata_items WHERE collection_id = $1 AND metadata_id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()
	for _, id := range metadataIds {
		_, err = stmt.ExecContext(ctx, collectionId, id)
		if err != nil {
			return err
		}
	}
	return nil
}

func (ds *DataStore) SetCollectionWorkflowStateId(ctx context.Context, id string, stateId string) error {
	stmt, err := ds.db.PrepareContext(ctx, "UPDATE collections set workflow_state_id = $1 WHERE id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, stateId, id)
	return err
}

func (ds *DataStore) GetCollectionCollectionItemNames(ctx context.Context, collectionId string) ([]string, error) {
	stmt, err := ds.db.PrepareContext(ctx, "SELECT name FROM collections WHERE id in (SELECT child_id FROM collection_collection_items WHERE collection_id = $1)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, collectionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	names := make([]string, 0)
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}

func (ds *DataStore) GetCollectionMetadataItemNames(ctx context.Context, collectionId string) ([]string, error) {
	stmt, err := ds.db.PrepareContext(ctx, "SELECT name FROM metadata WHERE id in (SELECT metadata_id FROM collection_metadata_items WHERE collection_id = $1)")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	rows, err := stmt.QueryContext(ctx, collectionId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	names := make([]string, 0)
	var name string
	for rows.Next() {
		err = rows.Scan(&name)
		if err != nil {
			return nil, err
		}
		names = append(names, name)
	}
	return names, nil
}
