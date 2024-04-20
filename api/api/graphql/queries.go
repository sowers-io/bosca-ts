package graphql

import (
	object "bosca.io/api/graphql/content"
	model "bosca.io/api/protobuf/content"
	"github.com/graphql-go/graphql"
)

var QueryFields = graphql.Fields{
	"metadata": &graphql.Field{
		Type:        graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(object.MetadataObjectConfig))),
		Description: "Get Metadata",
		Resolve: func(p graphql.ResolveParams) (interface{}, error) {
			metadata := make([]*model.Metadata, 0)
			metadata = append(metadata, &model.Metadata{
				Id: "hi",
			})
			return metadata, nil
		},
	},
}
