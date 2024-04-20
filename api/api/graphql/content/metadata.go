package content

import (
	"bosca.io/api/protobuf/content"
	"github.com/graphql-go/graphql"
)

var MetadataObjectConfig = graphql.NewObject(graphql.ObjectConfig{
	Name: "Metadata",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Name: "id",
			Type: graphql.ID,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var metadata = p.Source.(*content.Metadata)
				return metadata.Id, nil
			},
		},
	},
})
