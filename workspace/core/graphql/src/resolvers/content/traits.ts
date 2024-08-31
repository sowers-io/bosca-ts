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

import { Resolvers, Trait as GTrait } from '../../generated/resolvers'
import { GraphQLRequestContext, executeGraphQL, getGraphQLHeaders } from '@bosca/common'
import { useClient } from '@bosca/common'
import { ContentService, Empty, IdRequest } from '@bosca/protobufs'

export const resolvers: Resolvers<GraphQLRequestContext> = {
  Query: {
    traits: async (_, __, context) => {
      return await executeGraphQL<GTrait[]>(async () => {
        const service = useClient(ContentService)
        const traits = await service.getTraits(new Empty(), {
          headers: getGraphQLHeaders(context),
        })
        return traits.traits.map((s) => s.toJson()) as unknown as GTrait[]
      })
    },
    trait: async (_, { id }, context) => {
      return await executeGraphQL<GTrait | null>(async () => {
        const service = useClient(ContentService)
        const trait = await service.getTrait(new IdRequest({ id: id }), {
          headers: getGraphQLHeaders(context),
        })
        return trait.toJson() as unknown as GTrait
      })
    },
  },
  Trait: {
    name: (parent) => {
      if (!parent.name) return parent.id
      return parent.name
    },
    workflowIds: (parent) => {
      if (!parent.workflowIds) return []
      return parent.workflowIds
    },
  },
}
