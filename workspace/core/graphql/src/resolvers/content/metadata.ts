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

import { Resolvers, Metadata as GMetadata, SignedUrl as GSignedUrl } from '../../generated/resolvers'
import { GraphQLRequestContext, executeGraphQL, getGraphQLHeaders } from '@bosca/common'
import { useClient } from '@bosca/common'
import { AddMetadataRequest, ContentService, IdRequest, Metadata } from '@bosca/protobufs'

export function transformMetadata(metadata: Metadata): GMetadata {
  const m = metadata.toJson() as unknown as GMetadata
  m.__typename = 'Metadata'
  if (metadata.attributes) {
    m.attributes = []
    for (const key in metadata.attributes) {
      m.attributes.push({
        name: key,
        value: metadata.attributes[key],
      })
    }
  }
  m.workflowState = {
    id: metadata.workflowStateId,
    pendingId: metadata.workflowStatePendingId,
    deleteWorkflowId: metadata.deleteWorkflowId,
  }
  return m
}

export const resolvers: Resolvers<GraphQLRequestContext> = {
  Query: {
    metadata: async (parent, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const metadata = await service.getMetadata(new IdRequest({ id: args.id }), {
          headers: getGraphQLHeaders(context),
        })
        return transformMetadata(metadata)
      })
    },
  },
  Metadata: {
    uploadUrl: async (parent, args, context) => {
      return await executeGraphQL<GSignedUrl>(async () => {
        const service = useClient(ContentService)
        const url = await service.getMetadataUploadUrl(new IdRequest({ id: parent.id }), {
          headers: getGraphQLHeaders(context),
        })
        return url.toJson() as unknown as GSignedUrl
      })
    },
    downloadUrl: async (parent, args, context) => {
      return (await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const url = await service.getMetadataDownloadUrl(new IdRequest({ id: parent.id }), {
          headers: getGraphQLHeaders(context),
        })
        return url.toJson() as unknown as GSignedUrl
      }))!
    },
  },
  Mutation: {
    addMetadata: async (parent, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const response = await service.addMetadata(
          new AddMetadataRequest({
            metadata: {
              name: args.metadata.name,
              // todo
            },
          })
        )
        const metadata = await service.getMetadata(new IdRequest({ id: response.id }), {
          headers: getGraphQLHeaders(context),
        })
        return transformMetadata(metadata)
      })
    },
  },
}
