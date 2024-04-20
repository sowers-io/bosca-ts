package main

import (
	graphqlConfig "bosca.io/api/graphql"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
	"net/http"
)

func main() {
	schema, err := graphql.NewSchema(
		graphql.SchemaConfig{
			Query: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   "QueryRoot",
					Fields: graphqlConfig.QueryFields,
				},
			),
			Mutation: graphql.NewObject(
				graphql.ObjectConfig{
					Name:   "MutationRoot",
					Fields: graphqlConfig.MutationFields,
				},
			),
		},
	)

	if err != nil {
		log.Fatalf("failed to create graphql schema: %v", err)
	}

	h := handler.New(&handler.Config{
		Schema:   &schema,
		Pretty:   true,
		GraphiQL: true,
	})

	http.Handle("/graphql", h)

	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("failed to start graphql server: %v", err)
	}
}
