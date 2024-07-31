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

import { Resolvers, Collection as GCollection, CollectionItem as GCollectionItem } from '../../generated/resolvers'
import { GraphQLRequestContext, executeGraphQL, getGraphQLHeaders } from '@bosca/common'
import { useClient } from '@bosca/common'
import { ContentService, IdRequest, Collection } from '@bosca/protobufs'
import { transformMetadata } from './metadata'

export function transformCollection(collection: Collection): GCollection {
  const c = collection.toJson() as unknown as GCollection
  c.__typename = 'Collection'
  if (collection.attributes) {
    c.attributes = []
    for (const key in collection.attributes) {
      c.attributes.push({
        name: key,
        value: collection.attributes[key],
      })
    }
  }
  return c
}

export const resolvers: Resolvers<GraphQLRequestContext> = {
  Query: {
    collection: async (_, { id }, context) => {
      return await executeGraphQL<GCollection | null>(async () => {
        const service = useClient(ContentService)
        const collection = await service.getCollection(new IdRequest({ id: id }), {
          headers: getGraphQLHeaders(context),
        })
        if (!collection) return null
        return transformCollection(collection)
      })
    },
  },
  Collection: {
    items: async (parent, args, context) => {
      return await executeGraphQL<GCollectionItem[]>(async () => {
        const service = useClient(ContentService)
        const items = await service.getCollectionItems(new IdRequest({ id: parent.id }), {
          headers: getGraphQLHeaders(context),
        })
        return items.items.map((item) => {
          if (item.Item.case === 'collection') {
            return transformCollection(item.Item.value)
          } else if (item.Item.case === 'metadata') {
            return transformMetadata(item.Item.value)
          } else {
            throw new Error('unsupported type: ' + item.Item.case)
          }
        }) as GCollectionItem[]
      })
    },
  },
}
