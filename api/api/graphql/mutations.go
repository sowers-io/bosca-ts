package graphql

import "github.com/graphql-go/graphql"

var MutationFields = graphql.Fields{
	"addMetadata": &graphql.Field{
		Type: graphql.String,
		Args: graphql.FieldConfigArgument{
			"id": &graphql.ArgumentConfig{
				Type: graphql.NewNonNull(graphql.String),
			},
		},
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			// Insert your code here
			return "Metadata created", nil
		},
	},
}
