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
	"bosca.io/api/graphql/common"
	"bosca.io/api/graphql/content"
	grpc "bosca.io/api/protobuf/bosca/content"
	"github.com/graphql-go/graphql"
	opts "google.golang.org/grpc"
)

func NewMutationFields(client grpc.ContentServiceClient) graphql.Fields {
	return graphql.Fields{
		"addMetadata": &graphql.Field{
			Type: graphql.NewNonNull(content.SignedUrlObjectConfig),
			Args: graphql.FieldConfigArgument{
				"collection": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"name": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"contentType": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.String),
				},
				"attributes": &graphql.ArgumentConfig{
					Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(common.KeyValueObjectConfig))),
				},
			},
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				attributes := make(map[string]string)

				for _, attribute := range p.Args["attributes"].([]interface{}) {
					kv := attribute.(*common.KeyValue)
					attributes[kv.Key] = kv.Value
				}

				metadata := &grpc.Metadata{
					Name:        p.Args["name"].(string),
					ContentType: p.Args["contentType"].(string),
					Attributes:  attributes,
				}

				rv := p.Info.RootValue.(map[string]interface{})
				authorization := rv["authorization"].(string)

				url, err := client.AddMetadata(
					p.Context,
					&grpc.AddMetadataRequest{
						Collection: p.Args["collection"].(string),
						Metadata:   metadata,
					},
					opts.PerRPCCredsCallOption{Creds: &common.Authorization{HeaderValue: authorization}},
				)

				return url, err
			},
		},
	}
}
