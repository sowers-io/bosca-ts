import { Resolvers, Source as GSource } from '../../generated/resolvers'
import { RequestContext } from '../../context'
import { useClient } from '@bosca/common'
import { ContentService, Empty } from '@bosca/protobufs'
import { execute, getHeaders } from '../../util/requests'

export const resolvers: Resolvers<RequestContext> = {
  Query: {
    sources: async (_, __, context) => {
      return await execute(async () => {
        const service = useClient(ContentService)
        const sources = await service.getSources(new Empty(), {
          headers: getHeaders(context),
        })
        return sources.sources.map((s) => s.toJson())
      }) as GSource[]
    },
  },
}
