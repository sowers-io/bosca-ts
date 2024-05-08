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

import "github.com/meilisearch/meilisearch-go"
import "bosca.io/pkg/search"

type meilisearchSearch struct {
	client        *meilisearch.Client
	metadataIndex search.Index
}

type meilisearchIndex struct {
	name string
}

func (m *meilisearchIndex) Name() string {
	return m.name
}

func NewSearchClient(client *meilisearch.Client) search.Client {
	return &meilisearchSearch{
		client: client,
		metadataIndex: &meilisearchIndex{
			name: MetadataIndex,
		},
	}
}

func (m *meilisearchSearch) GetMetadataIndex() search.Index {
	return m.metadataIndex
}

func (m *meilisearchSearch) Index(index search.Index, doc *search.Document) error {
	ix := m.client.Index(index.Name())
	_, err := ix.AddDocuments([]*search.Document{doc})
	if err != nil {
		return err
	}
	return nil
}
