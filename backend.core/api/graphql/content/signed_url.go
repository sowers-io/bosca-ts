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

package content

import (
	"bosca.io/api/graphql/common"
	"bosca.io/api/protobuf/bosca/content"
	"github.com/graphql-go/graphql"
)

var SignedUrlObjectConfig = graphql.NewObject(graphql.ObjectConfig{
	Name: "SignedUrl",
	Fields: graphql.Fields{
		"url": &graphql.Field{
			Type: graphql.NewNonNull(graphql.String),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var url = p.Source.(*content.SignedUrl)
				return url.Url, nil
			},
		},
		"headers": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(common.KeyValueObjectConfig))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var url = p.Source.(*content.SignedUrl)
				headers := make([]common.KeyValue, 0)
				for _, hdr := range url.Headers {
					headers = append(headers, common.KeyValue{
						Key:   hdr.Name,
						Value: hdr.Value,
					})
				}
				return headers, nil
			},
		},
		"attributes": &graphql.Field{
			Type: graphql.NewNonNull(graphql.NewList(graphql.NewNonNull(common.KeyValueObjectConfig))),
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				var url = p.Source.(*content.SignedUrl)
				headers := make([]common.KeyValue, 0)
				for key, val := range url.Attributes {
					headers = append(headers, common.KeyValue{
						Key:   key,
						Value: val,
					})
				}
				return headers, nil
			},
		},
	},
})
