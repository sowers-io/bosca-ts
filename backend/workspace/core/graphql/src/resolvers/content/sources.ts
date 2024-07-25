import { Resolvers, Source as GSource } from '../../generated/resolvers'
import { RequestContext } from '../../context'
import { useClient } from '@bosca/common'
import { ContentService, Empty, IdRequest } from '@bosca/protobufs'
import { execute, getHeaders } from '../../util/requests'

export const resolvers: Resolvers<RequestContext> = {
  Query: {
    sources: async (_, __, context) => {
      return await execute<GSource[]>(async () => {
        const service = useClient(ContentService)
        const sources = await service.getSources(new Empty(), {
          headers: getHeaders(context),
        })
        return sources.sources.map((s) => s.toJson()) as unknown as GSource[]
      })
    },
    source: async (_, { id }, context) => {
      return await execute<GSource | null>(async () => {
        const service = useClient(ContentService)
        const source = await service.getSource(new IdRequest({id: id}), {
          headers: getHeaders(context),
        })
        if (!source) return null
        return source.toJson() as unknown as GSource
      })
    }
  },
}
