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

import {
  Resolvers,
  Collection as GCollection,
  CollectionItem as GCollectionItem,
} from '../../generated/resolvers'
import { GraphQLRequestContext, executeGraphQL, getGraphQLHeaders } from '@bosca/common'
import { useClient } from '@bosca/common'
import {
  ContentService,
  IdRequest,
  Collection,
  AddCollectionRequest,
} from '@bosca/protobufs'
import { transformMetadata } from './metadata'
import { GraphQLError } from 'graphql'
import { toGraphPermissions, toGrpcPermissions } from '../../util'

export function transformCollection(collection: Collection): GCollection {
  const c = collection.toJson() as unknown as GCollection
  c.__typename = 'Collection'
  if (collection.attributes) {
    c.attributes = []
    for (const key in collection.attributes) {
      c.attributes.push({
        key: key,
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
  Mutation: {
    addCollection: async (_, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const response = await service.addCollection(
          new AddCollectionRequest({
            parent: args.parent || '00000000-0000-0000-0000-000000000000',
            collection: {
              name: args.collection.name,
            },
          }), {
            headers: getGraphQLHeaders(context),
          },
        )
        let lastError: any | null = null
        for (let tries = 0; tries < 100; tries++) {
          try {
            const collection = await service.getCollection(new IdRequest({ id: response.id }), {
              headers: getGraphQLHeaders(context),
            })
            return transformCollection(collection)
          } catch (e) {
            lastError = e
            await new Promise((resolve) => setTimeout(resolve, 100))
          }
        }
        if (lastError) {
          throw lastError
        }
        throw new GraphQLError('failed to get metadata after it was created')
      })
    },
    deleteCollection: async (_, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        await service.deleteCollection(new IdRequest({ id: args.id! }), {
          headers: getGraphQLHeaders(context),
        })
        return true
      })
    },
    addCollectionPermissions: async (_, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        await service.addCollectionPermissions(toGrpcPermissions(args.id, args.permissions), {
          headers: getGraphQLHeaders(context),
        })
        const metadata = await service.getCollection(new IdRequest({ id: args.id }), {
          headers: getGraphQLHeaders(context),
        })
        return transformCollection(metadata)
      })
    },
    deleteCollectionPermissions: async (_, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        await service.deleteCollectionPermissions(toGrpcPermissions(args.id, args.permissions), {
          headers: getGraphQLHeaders(context),
        })
        const collection = await service.getCollection(new IdRequest({ id: args.id }), {
          headers: getGraphQLHeaders(context),
        })
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
        return items.items.filter((item) => {
          if (args.filter) {
            if (args.filter.created) {
              // TODO: move this filter to getCollectionItems
              const dt = new Date(Date.parse(args.filter.created))
              let created: Date | undefined = undefined
              if (item.Item.case === 'collection') {
                created = item.Item.value.created?.toDate()
              } else if (item.Item.case === 'metadata') {
                created = item.Item.value.created?.toDate()
              }
              if (!created) return false
              return created.getTime() > dt.getTime()
            }
          }
          return true
        }).map((item) => {
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
    permissions: async (parent, _, context) => {
      return (await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const request = new IdRequest({ id: parent.id })
        const response = await service.getCollectionPermissions(request, {
          headers: getGraphQLHeaders(context),
        })
        return toGraphPermissions(parent.id, response)
      }))!
    },
  },
}
