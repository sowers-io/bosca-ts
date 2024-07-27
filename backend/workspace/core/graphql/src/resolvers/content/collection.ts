import { Resolvers, Collection as GCollection, CollectionItem as GCollectionItem } from '../../generated/resolvers'
import { RequestContext } from '../../context'
import { useClient } from '@bosca/common'
import { ContentService, IdRequest, Collection, Collections, CollectionItem } from '@bosca/protobufs'
import { execute, getHeaders } from '../../util/requests'
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

export const resolvers: Resolvers<RequestContext> = {
  Query: {
    collection: async (_, { id }, context) => {
      return await execute<GCollection | null>(async () => {
        const service = useClient(ContentService)
        const collection = await service.getCollection(new IdRequest({ id: id }), {
          headers: getHeaders(context),
        })
        if (!collection) return null
        return transformCollection(collection)
      })
    },
  },
  Collection: {
    items: async (parent, args, context) => {
      return await execute<GCollectionItem[]>(async () => {
        const service = useClient(ContentService)
        const items = await service.getCollectionItems(new IdRequest({ id: parent.id }), {
          headers: getHeaders(context),
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
