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
