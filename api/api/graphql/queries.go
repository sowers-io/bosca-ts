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

package graphql

import (
	content "bosca.io/api/graphql/content"
	grpc "bosca.io/api/protobuf/content"
	"github.com/graphql-go/graphql"
)

func NewQueryFields(client grpc.ContentServiceClient) graphql.Fields {
	return graphql.Fields{
		"metadata": &graphql.Field{
			Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(content.MetadataObjectConfig))),
			Description: "Get Metadata",
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				metadata := make([]*grpc.Metadata, 0)
				metadata = append(metadata, &grpc.Metadata{
					Id: "hi",
				})
				return metadata, nil
			},
		},
	}
}
