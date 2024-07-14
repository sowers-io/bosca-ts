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
	"bosca.io/api/protobuf/bosca/content"
	"bosca.io/api/workflow"
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
	"strings"
	"time"
)

func (ds *DataStore) AddMetadataRelationship(ctx context.Context, metadataId1 string, metadataId2 string, relationship string) error {
	_, err := ds.db.ExecContext(ctx, "INSERT INTO metadata_relationship (metadata1_id, metadata2_id, relationship) values ($1::uuid, $2::uuid, $3)", metadataId1, metadataId2, relationship)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStore) GetMetadataSupplementary(ctx context.Context, metadataId, key string) (*content.MetadataSupplementary, error) {
	row := ds.db.QueryRowContext(ctx, "select name, content_type, content_length, source_id, source_identifier from metadata_supplementary where metadata_id = $1::uuid and \"key\" = $2", metadataId, key)
	if row.Err() != nil {
		return nil, row.Err()
	}
	s := &content.MetadataSupplementary{
		MetadataId: metadataId,
		Key:        key,
	}
	err := row.Scan(&s.Name, &s.ContentType, &s.ContentLength, &s.SourceId, &s.SourceIdentifier)
	if err != nil {
		return nil, err
	}
	return s, nil
}

func (ds *DataStore) AddMetadataSupplementary(ctx context.Context, metadataId, key, name, contentType string, contentLength int64, traitIds []string, sourceId, sourceIdentifier *string) error {
	_, err := ds.db.ExecContext(ctx, "INSERT INTO metadata_supplementary (metadata_id, \"key\", name, content_type, content_length, source_id, source_identifier) values ($1::uuid, $2, $3, $4, $5, $6, $7)", metadataId, key, name, contentType, contentLength, sourceId, sourceIdentifier)
	if err != nil {
		return err
	}

	stmt, err := ds.db.PrepareContext(ctx, "insert into metadata_supplementary_traits (metadata_id, key, trait_id) values ($1, $2, $3)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, traitId := range traitIds {
		_, err = stmt.ExecContext(ctx, metadataId, key, traitId)
		if err != nil {
			return err
		}
	}

	return nil
}

func (ds *DataStore) SetMetadataSupplementaryReady(ctx context.Context, metadataId, key string) error {
	_, err := ds.db.ExecContext(ctx, "update metadata_supplementary set uploaded = now() where metadata_id = $1::uuid and \"key\" = $2", metadataId, key)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStore) DeleteMetadataSupplementary(ctx context.Context, metadataId, key string) error {
	_, err := ds.db.ExecContext(ctx, "delete from metadata_supplementary where metadata_id = $1::uuid and \"key\" = $2", metadataId, key)
	if err != nil {
		return err
	}
	return nil
}

func (ds *DataStore) AddMetadata(ctx context.Context, metadata *content.Metadata) (string, error) {
	stmt, err := ds.db.PrepareContext(ctx, "INSERT INTO metadata (name, content_type, content_length, labels, attributes, source_id, source_identifier, language_tag) VALUES ($1, $2, $3, $4, ($5)::jsonb, $6, $7, $8) returning id")
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	labels := metadata.Labels
	if labels == nil {
		labels = make([]string, 0)
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
		labels,
		attributes,
		metadata.SourceId,
		metadata.SourceIdentifier,
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

	if metadata.TraitIds != nil {
		for _, traitId := range metadata.TraitIds {
			_, err = ds.db.ExecContext(ctx, "insert into metadata_traits (metadata_id, trait_id) values ($1, $2)", id, traitId)
			if err != nil {
				return id, err
			}
		}
	}

	if metadata.CategoryIds != nil {
		for _, categoryId := range metadata.CategoryIds {
			_, err = ds.db.ExecContext(ctx, "insert into metadata_categories (metadata_id, category_id) values ($1, $2)", id, categoryId)
			if err != nil {
				return id, err
			}
		}
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

func (ds *DataStore) AddMetadataTrait(ctx context.Context, id string, trait string) (*content.Metadata, error) {
	metadata, err := ds.GetMetadata(ctx, id)
	if err != nil {
		return nil, err
	}
	if metadata.WorkflowStateId != workflow.StateDraft {
		return nil, errors.New("not in a draft state")
	}
	_, err = ds.db.ExecContext(ctx, "insert into metadata_traits (metadata_id, trait_id) values ($1, $2)", id, trait)
	if err != nil {
		return nil, err
	}
	metadata.TraitIds = append(metadata.TraitIds, trait)
	return metadata, nil
}

func (ds *DataStore) FindMetdata(ctx context.Context, request *content.FindMetadataRequest) ([]*content.Metadata, error) {
	if request.Attributes == nil {
		return nil, errors.New("a request argument is required")
	}
	queryString := &strings.Builder{}
	queryString.WriteString("SELECT id, name, labels, attributes, content_type, content_length, created, modified, source_id, source_identifier, language_tag, workflow_state_id, workflow_state_pending_id FROM metadata")
	where := &strings.Builder{}
	if request.Attributes != nil {
		i := 1
		for _ = range request.Attributes {
			if where.Len() > 0 {
				where.WriteString(" AND ")
			}
			where.WriteString(fmt.Sprintf("attributes->>$%d = $%d", i, i+1))
			i += 2
		}
	}
	queryString.WriteString(" WHERE ")
	queryString.WriteString(where.String())
	metadataQuery, err := ds.db.PrepareContext(ctx, queryString.String())
	if err != nil {
		return nil, err
	}
	defer metadataQuery.Close()
	args := make([]any, 0, len(request.Attributes)*2)
	for i, v := range request.Attributes {
		args = append(args, i)
		args = append(args, v)
	}
	return ds.getMetadatas(ctx, metadataQuery, args)
}

func (ds *DataStore) GetMetadatas(ctx context.Context, id []string) ([]*content.Metadata, error) {
	if len(id) == 0 {
		return nil, nil
	}
	queryString := &strings.Builder{}
	queryString.WriteString("SELECT id, name, labels, attributes, content_type, content_length, created, modified, source_id, source_identifier, language_tag, workflow_state_id, workflow_state_pending_id FROM metadata WHERE id = $1")
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
	args := make([]any, len(id))
	for i, v := range id {
		u, err := uuid.Parse(v)
		if err != nil {
			return nil, err
		}
		args[i] = u
	}
	return ds.getMetadatas(ctx, metadataQuery, args)
}

func (ds *DataStore) getMetadatas(ctx context.Context, metadataQuery *sql.Stmt, args []any) ([]*content.Metadata, error) {
	m := pgtype.NewMap()
	result, err := metadataQuery.QueryContext(ctx, args...)
	if err != nil {
		return nil, err
	}
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
	metadatas := make([]*content.Metadata, 0)
	for result.Next() {
		var metadata content.Metadata
		var created time.Time
		var modified time.Time
		var labels []string
		var attributesJson json.RawMessage
		err = result.Scan(
			&metadata.Id,
			&metadata.Name,
			m.SQLScanner(&labels),
			&attributesJson,
			&metadata.ContentType,
			&metadata.ContentLength,
			&created,
			&modified,
			&metadata.SourceId,
			&metadata.SourceIdentifier,
			&metadata.LanguageTag,
			&metadata.WorkflowStateId,
			&metadata.WorkflowStatePendingId,
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
		metadata.Labels = labels
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
