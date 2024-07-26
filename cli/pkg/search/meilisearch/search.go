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

package meilisearch

import (
	"errors"
	"github.com/meilisearch/meilisearch-go"
)
import "bosca.io/pkg/search"

const MetadataIndex = "metadata"

type meilisearchSearch struct {
	client        *meilisearch.Client
	metadataIndex search.Index
}

type meilisearchIndex struct {
	index *meilisearch.Index
}

func (m *meilisearchIndex) Name() string {
	return m.index.UID
}

func NewSearchClient(client *meilisearch.Client) (search.StandardClient, error) {
	ix, err := client.GetIndex(MetadataIndex)
	if err != nil {
		return nil, err
	}
	if ix == nil {
		_, err = client.CreateIndex(&meilisearch.IndexConfig{
			Uid:        MetadataIndex,
			PrimaryKey: "id",
		})
		if err != nil {
			return nil, err
		}
	}
	s := &meilisearchSearch{
		client:        client,
		metadataIndex: &meilisearchIndex{index: ix},
	}
	return s, nil
}

func (m *meilisearchSearch) GetMetadataIndex() search.Index {
	return m.metadataIndex
}

func (m *meilisearchSearch) Index(index search.Index, doc *search.Document) error {
	ix := index.(*meilisearchIndex).index
	taskInfo, err := ix.AddDocuments([]*search.Document{doc})
	if err != nil {
		return err
	}
	task, err := ix.WaitForTask(taskInfo.TaskUID)
	if err != nil {
		return err
	}
	if task.Status != meilisearch.TaskStatusSucceeded {
		return errors.New(task.Error.Message)
	}
	return nil
}
