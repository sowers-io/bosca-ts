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

package factory

import (
	"bosca.io/pkg/configuration"
	"bosca.io/pkg/search"
	"bosca.io/pkg/search/meilisearch"
	"fmt"
)

func NewSearch(cfg *configuration.SearchConfiguration) (search.StandardClient, error) {
	switch cfg.Type {
	case configuration.SearchTypeMeilisearch:
		client := meilisearch.NewMeilisearchClient(cfg.Endpoint, cfg.ApiKey)
		return meilisearch.NewSearchClient(client)
	default:
		return nil, fmt.Errorf("unknown storage type: %s", cfg.Type)
	}
}
