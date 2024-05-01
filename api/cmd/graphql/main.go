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

package main

import (
	graphqlConfig "bosca.io/api/graphql"
	"bosca.io/api/protobuf/content"
	"bosca.io/pkg/clients"
	"bosca.io/pkg/configuration"
	"context"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main() {
	cfg := configuration.NewServerConfiguration("", 0, 0)

	contentConnection := clients.NewClientConnection(cfg.ClientEndPoints.ContentApiAddress)
	defer contentConnection.Close()

	contentClient := content.NewContentServiceClient(contentConnection)

	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   "QueryRoot",
					Fields: graphqlConfig.NewQueryFields(contentClient),
				},
			),
			Mutation: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   "MutationRoot",
					Fields: graphqlConfig.NewMutationFields(contentClient),
				},
			),
		},
	)

	if err != nil {
		log.Fatalf("failed to create graphql schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema: &schema,
		Pretty: true,
		RootObjectFn: func(ctx context.Context, r *http.Request) map[string]interface{} {
			authorization := r.Header.Get("Authorization")
			return map[string]interface{}{
				"authorization": authorization,
			}
		},
	})

	http.Handle("/graphql", h)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start graphql server: %v", err)
	}
}
