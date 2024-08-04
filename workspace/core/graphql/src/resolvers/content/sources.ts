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

import { Resolvers, Source as GSource } from '../../generated/resolvers'
import { GraphQLRequestContext, executeGraphQL, getGraphQLHeaders } from '@bosca/common'
import { useClient } from '@bosca/common'
import { ContentService, Empty, IdRequest } from '@bosca/protobufs'

export const resolvers: Resolvers<GraphQLRequestContext> = {
  Query: {
    sources: async (_, __, context) => {
      return await executeGraphQL<GSource[]>(async () => {
        const service = useClient(ContentService)
        const sources = await service.getSources(new Empty(), {
          headers: getGraphQLHeaders(context),
        })
        return sources.sources.map((s) => s.toJson()) as unknown as GSource[]
      })
    },
    source: async (_, { id }, context) => {
      return await executeGraphQL<GSource | null>(async () => {
        const service = useClient(ContentService)
        const source = await service.getSource(new IdRequest({ id: id }), {
          headers: getGraphQLHeaders(context),
        })
        if (!source) return null
        return source.toJson() as unknown as GSource
      })
    },
  },
}
