import { Resolvers, Trait as GTrait } from '../../generated/resolvers'
import { RequestContext } from '../../context'
import { useClient } from '@bosca/common'
import { ContentService, Empty, IdRequest } from '@bosca/protobufs'
import { execute, getHeaders } from '../../util/requests'

export const resolvers: Resolvers<RequestContext> = {
  Query: {
    traits: async (_, __, context) => {
      return await execute<GTrait[]>(async () => {
        const service = useClient(ContentService)
        const traits = await service.getTraits(new Empty(), {
          headers: getHeaders(context),
        })
        return traits.traits.map((s) => s.toJson()) as unknown as GTrait[]
      })
    },
    trait: async (_, { id }, context) => {
      return await execute<GTrait | null>(async () => {
        const service = useClient(ContentService)
        const trait = await service.getTrait(new IdRequest({ id: id }), {
          headers: getHeaders(context),
        })
        return trait.toJson() as unknown as GTrait
      })
    },
  },
}
