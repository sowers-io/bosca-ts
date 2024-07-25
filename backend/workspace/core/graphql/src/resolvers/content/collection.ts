import { Resolvers, Collection as GCollection } from '../../generated/resolvers'
import { RequestContext } from '../../context'
import { useClient } from '@bosca/common'
import { ContentService, IdRequest, Collection } from '@bosca/protobufs'
import { execute, getHeaders } from '../../util/requests'

function transformCollection(collection: Collection): GCollection {
  const c = collection.toJson() as unknown as GCollection
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
}
