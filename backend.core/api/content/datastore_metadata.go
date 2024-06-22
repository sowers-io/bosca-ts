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
	if metadata.WorkflowStateId != WorkflowStateDraft {
		return nil, errors.New("not in a draft state")
	}
	_, err = ds.db.ExecContext(ctx, "insert into metadata_traits (metadata_id, trait_id) values ($1, $2)", id, trait)
	if err != nil {
		return nil, err
	}
	metadata.TraitIds = append(metadata.TraitIds, trait)
	return metadata, nil
}

func (ds *DataStore) GetMetadatas(ctx context.Context, id []string) ([]*content.Metadata, error) {
	if len(id) == 0 {
		return nil, nil
	}

	m := pgtype.NewMap()

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
