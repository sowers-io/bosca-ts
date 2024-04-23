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
)

type DataStore struct {
	db *sql.DB
}

func NewDataStore(db *sql.DB) *DataStore {
	return &DataStore{
		db,
	}
}

func (ds *DataStore) AddRootCollection(ctx context.Context) (bool, error) {
	root, err := ds.GetMetadata(ctx, "00000000-0000-0000-0000-000000000000")
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		return false, err
	}
	if root != nil {
		return false, nil
	}
	_, err = ds.db.ExecContext(ctx, "insert into collections (id, name, type, workflow_state_id) values ('00000000-0000-0000-0000-000000000000', 'Root', 'root', 'published')")
	if err != nil {
		return false, err
	}
	return true, nil
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

func (ds *DataStore) GetMetadata(ctx context.Context, id string) (*content.Metadata, error) {
	var metadata content.Metadata
	err := ds.db.QueryRowContext(ctx, "SELECT * FROM metadata WHERE id = $1", id).Scan(
		&metadata.Id,
		&metadata.Name,
		&metadata.ContentType,
		&metadata.TraitIds,
		&metadata.CategoryIds,
		&metadata.Tags,
		&metadata.Attributes,
		&metadata.Created,
		&metadata.Modified,
		&metadata.Status,
	)
	if err != nil {
		return nil, err
	}
	return &metadata, nil
}

func (ds *DataStore) AddMetadata(ctx context.Context, metadata *content.Metadata) (string, error) {
	stmt, err := ds.db.PrepareContext(ctx, "INSERT INTO metadata (name, content_type, tags, attributes) VALUES ($1, $2, $3, ($4)::jsonb) returning id")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	tags := metadata.Tags
	if tags == nil {
		tags = make([]string, 0)
	}
	attributes := metadata.Attributes
	if attributes == nil {
		attributes = make(map[string]string)
	}

	result := stmt.QueryRowContext(ctx,
		metadata.Name,
		metadata.ContentType,
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

func (ds *DataStore) DeleteMetadata(ctx context.Context, id string) error {
	stmt, err := ds.db.PrepareContext(ctx, "DELETE FROM metadata WHERE id = $1::uuid")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, id)
	return err
}

func (ds *DataStore) SetMetadataWorkflowStateId(ctx context.Context, id string, stateId string) error {
	stmt, err := ds.db.PrepareContext(ctx, "UPDATE metadata set workflow_state_id = $1 WHERE id = $2::uuid")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, stateId, id)
	return err
}

func (ds *DataStore) SetMetadataStatus(ctx context.Context, id string, status content.MetadataStatus) error {
	stmt, err := ds.db.PrepareContext(ctx, "UPDATE metadata set status = ($1)::metadata_status WHERE id = $2")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.ExecContext(ctx, status.String(), id)
	return err
}
