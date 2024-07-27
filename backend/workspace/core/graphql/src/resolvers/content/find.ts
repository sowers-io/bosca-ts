import {
  Resolvers,
  Metadata as GMetadata,
  Collection as GCollection,
} from '../../generated/resolvers'
import { RequestContext } from '../../context'
import { useClient } from '@bosca/common'
import { AddMetadataRequest, ContentService, FindMetadataRequest, IdRequest, Metadata } from '@bosca/protobufs'
import { execute, getHeaders } from '../../util/requests'
import { transformMetadata } from './metadata'
import { transformCollection } from './collection'

export const resolvers: Resolvers<RequestContext> = {
  Query: {
    find: async (_) => {
      return {
        __typename: 'Find',
        metadata: [],
        collections: [],
      }
    },
  },
  Find: {
    collections: async (parent, args, context, info) => {
      return await execute<GCollection[]>(async () => {
        const request = new FindMetadataRequest({ attributes: {} })
        for (const attribute of args.query!.attributes) {
          request.attributes[attribute.name!] = attribute.value || ''
        }
        const service = useClient(ContentService)
        const collections = await service.findCollection(request, {
          headers: getHeaders(context),
        })
        return collections.collections.map((c) => {
          return transformCollection(c)
        })
      })
    },
    metadata: async (parent, args, context, info) => {
      return await execute<GMetadata[]>(async () => {
        const request = new FindMetadataRequest({ attributes: {} })
        for (const attribute of args.query!.attributes) {
          request.attributes[attribute.name!] = attribute.value || ''
        }
        const service = useClient(ContentService)
        const metadata = await service.findMetadata(request, {
          headers: getHeaders(context),
        })
        return metadata.metadata.map((m) => {
          return transformMetadata(m)
        })
      })
    },
  },
}
