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
        if (!args.query) return []
        const request = new FindMetadataRequest({ attributes: {} })
        for (const attribute of args.query.attributes) {
          if (!attribute.name) continue
          request.attributes[attribute.name] = attribute.value || ''
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
