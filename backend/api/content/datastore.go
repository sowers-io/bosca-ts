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
	"bosca.io/pkg/security/identity"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
	"log/slog"
	"strings"
	"time"
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
	root, err := ds.GetCollection(ctx, RootCollectionId)
	if err != nil && !errors.Is(err, sql.ErrNoRows) {
		slog.Error("failed to get metadata for root collection", slog.Any("error", err))
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

func (ds *DataStore) GetWorkflow(ctx context.Context, id string) (*content.Workflow, error) {
	row := ds.db.QueryRowContext(ctx, "SELECT id, name, description, queue, configuration FROM workflows WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var workflow content.Workflow
	var configuration json.RawMessage
	err := row.Scan(&workflow.Id, &workflow.Name, &workflow.Description, &workflow.Queue, &configuration)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	err = json.Unmarshal(configuration, &workflow.Configuration)
	if err != nil {
		return nil, err
	}
	return &workflow, nil
}

func (ds *DataStore) GetWorkflowTransition(ctx context.Context, fromStateId string, toStateId string) (*content.WorkflowStateTransition, error) {
	row := ds.db.QueryRowContext(ctx, "SELECT from_state_id, to_state_id FROM workflow_state_transitions WHERE from_state_id = $1 AND to_state_id = $2", fromStateId, toStateId)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var transition content.WorkflowStateTransition
	err := row.Scan(&transition.FromStateId, &transition.ToStateId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &transition, nil
}

func (ds *DataStore) GetWorkflowState(ctx context.Context, id string) (*content.WorkflowState, error) {
	row := ds.db.QueryRowContext(ctx, "SELECT id, name, description, type, configuration, workflow_id, exist_workflow_id, entry_workflow_id FROM workflow_states WHERE id = $1", id)
	if row.Err() != nil {
		return nil, row.Err()
	}
	var state content.WorkflowState
	var configuration json.RawMessage
	err := row.Scan(&state.Id, &state.Name, &state.Description, &state.Type, &configuration, &state.WorkflowId, &state.ExitWorkflowId, &state.EntryWorkflowId)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	err = json.Unmarshal(configuration, &state.Configuration)
	if err != nil {
		return nil, err
	}
	return &state, nil
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
	metadatas, err := ds.GetMetadatas(ctx, []string{id})
	if err != nil {
		return nil, err
	}
	if len(metadatas) == 0 {
		return nil, nil
	}
	return metadatas[0], nil
}

func (ds *DataStore) GetMetadatas(ctx context.Context, id []string) ([]*content.Metadata, error) {
	if len(id) == 0 {
		return nil, nil
	}

	m := pgtype.NewMap()

	queryString := &strings.Builder{}
	queryString.WriteString("SELECT id, name, tags, attributes, content_type, content_length, created, modified, source, language_tag, workflow_state_id FROM metadata WHERE id = $1")
	if len(id) > 1 {
		for i := 1; i < len(id); i++ {
			queryString.WriteString(fmt.Sprintf(" OR id = $%d", i+1))
		}
	}
	metadataQuery, err := ds.db.PrepareContext(ctx, queryString.String())
	if err != nil {
		return nil, err
	}
	defer metadataQuery.Close()

	traitsQuery, err := ds.db.PrepareContext(ctx, "select trait_id from metadata_traits where metadata_id = $1")
	if err != nil {
		return nil, err
	}
	defer traitsQuery.Close()

	categoriesQuery, err := ds.db.PrepareContext(ctx, "select category_id from metadata_categories where metadata_id = $1")
	if err != nil {
		return nil, err
	}
	defer categoriesQuery.Close()

	args := make([]any, len(id))
	for i, v := range id {
		u, err := uuid.Parse(v)
		if err != nil {
			return nil, err
		}
		args[i] = u
	}

	result, err := metadataQuery.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}

	metadatas := make([]*content.Metadata, 0)

	for result.Next() {
		var metadata content.Metadata
		var created time.Time
		var modified time.Time
		var tags []string
		var attributesJson json.RawMessage

		err = result.Scan(
			&metadata.Id,
			&metadata.Name,
			m.SQLScanner(&tags),
			&attributesJson,
			&metadata.ContentType,
			&metadata.ContentLength,
			&created,
			&modified,
			&metadata.Source,
			&metadata.LanguageTag,
			&metadata.WorkflowStateId,
		)

		if err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return nil, nil
			}
			return nil, err
		}

		err = json.Unmarshal(attributesJson, &metadata.Attributes)
		if err != nil {
			return nil, err
		}

		metadata.Created = timestamppb.New(created)
		metadata.Modified = timestamppb.New(modified)
		metadata.Tags = tags

		result, err := traitsQuery.QueryContext(ctx, metadata.Id)
		if err != nil {
			return nil, err
		}
		traits := make([]string, 0)
		for result.Next() {
			var trait string
			err = result.Scan(&trait)
			if err != nil {
				result.Close()
				return nil, err
			}
			traits = append(traits, trait)
		}
		metadata.TraitIds = traits
		result.Close()

		result, err = categoriesQuery.QueryContext(ctx, metadata.Id)
		if err != nil {
			return nil, err
		}
		categories := make([]string, 0)
		for result.Next() {
			var category string
			err = result.Scan(&category)
			if err != nil {
				result.Close()
				return nil, err
			}
			categories = append(categories, category)
		}
		metadata.CategoryIds = categories
		result.Close()

		metadatas = append(metadatas, &metadata)
	}

	return metadatas, nil
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

func (ds *DataStore) AddMetadata(ctx context.Context, metadata *content.Metadata) (string, error) {
	stmt, err := ds.db.PrepareContext(ctx, "INSERT INTO metadata (name, content_type, content_length, tags, attributes, source, language_tag) VALUES ($1, $2, $3, $4, ($5)::jsonb, $6, $7) returning id")
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

	if metadata.LanguageTag == "" {
		// TODO: pull from some default setting
		metadata.LanguageTag = "en"
	}

	result := stmt.QueryRowContext(ctx,
		metadata.Name,
		metadata.ContentType,
		metadata.ContentLength,
		tags,
		attributes,
		metadata.Source,
		metadata.LanguageTag,
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

func (ds *DataStore) TransitionMetadataWorkflowStateId(ctx context.Context, metadata *content.Metadata, toState *content.WorkflowState, status string, success bool, complete bool) error {
	subjectId, err := identity.GetSubjectId(ctx)
	if err != nil {
		return err
	}

	txn, err := ds.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	_, err = txn.ExecContext(ctx, "insert into metadata_workflow_state_history (metadata_id, from_state_id, to_state_id, subject, status, success, complete) values ($1::uuid, $2, $3, $4, $5, $6)", metadata.Id, metadata.WorkflowStateId, toState.Id, subjectId, status, success, complete)
	if err != nil {
		txn.Rollback()
		return err
	}
	if !success {
		_, err = txn.ExecContext(ctx, "update metadata set workflow_state_pending_id = null where id = $1::uuid", metadata.Id)
		if err != nil {
			txn.Rollback()
			return err
		}
	} else {
		if complete {
			_, err = txn.ExecContext(ctx, "update metadata set workflow_state_id = $1, workflow_state_pending_id = null where id = $2::uuid", toState.Id, metadata.Id)
			if err != nil {
				txn.Rollback()
				return err
			}
		} else {
			_, err = txn.ExecContext(ctx, "update metadata set workflow_state_pending_id = $1 where id = $2::uuid", toState.Id, metadata.Id)
			if err != nil {
				txn.Rollback()
				return err
			}
		}
	}
	return txn.Commit()
}
