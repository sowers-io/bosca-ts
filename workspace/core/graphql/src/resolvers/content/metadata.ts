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

import { Resolvers, Metadata as GMetadata, SignedUrl as GSignedUrl, Supplementary } from '../../generated/resolvers'
import { GraphQLRequestContext, executeGraphQL, getGraphQLHeaders, toArrayBuffer } from '@bosca/common'
import { useClient, executeHttpRequest } from '@bosca/common'
import { AddMetadataRequest, ContentService, IdRequest, Metadata, MetadataReadyRequest, SupplementaryIdRequest } from '@bosca/protobufs'
import { protoBase64, protoInt64 } from '@bufbuild/protobuf'
import { GraphQLError } from 'graphql'

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
    supplementary: async (parent, args, context) => {
      return (await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const request = new SupplementaryIdRequest({ id: parent.id, key: args.key })
        const response = await service.getMetadataSupplementary(request, {
          headers: getGraphQLHeaders(context),
        })
        return response?.toJson() as unknown as Supplementary
      }))
    },
    supplementaries: async (parent, args, context) => {
      return (await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const request = new IdRequest({ id: parent.id })
        const response = await service.getMetadataSupplementaries(request, {
          headers: getGraphQLHeaders(context),
        })
        return response.supplementaries.map((s) => s.toJson()) as unknown as Supplementary[]
      }))!
    },
    content: async (parent, args, context) => {
      const type = parent.contentType.split(';')[0].trim()
      if (type === 'text/plain' || type === 'text/json') {
        const url = await executeGraphQL(async () => {
          const service = useClient(ContentService)
          const url = await service.getMetadataDownloadUrl(new IdRequest({ id: parent.id }), {
            headers: getGraphQLHeaders(context),
          })
          return url
        })
        if (!url) return null
        const content = await executeHttpRequest(url)
        if (type === 'text/json') {
          return {
            json: JSON.parse(content.toString())
          }
        }
        return {
          text: content.toString()
        }
      }
      return null
    },
  },
  Supplementary: {
    uploadUrl: async (parent, args, context) => {
      return await executeGraphQL<GSignedUrl>(async () => {
        const service = useClient(ContentService)
        const url = await service.getMetadataSupplementaryUploadUrl(new SupplementaryIdRequest({ id: parent.metadataId, key: parent.key }), {
          headers: getGraphQLHeaders(context),
        })
        return url.toJson() as unknown as GSignedUrl
      })
    },
    downloadUrl: async (parent, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const request = new SupplementaryIdRequest({ id: parent.metadataId, key: parent.key })
        const url = await service.getMetadataSupplementaryDownloadUrl(request, {
          headers: getGraphQLHeaders(context),
        })
        return url.toJson() as unknown as GSignedUrl
      })
    },
    content: async (parent, args, context) => {
      const type = parent.contentType.split(';')[0].trim()
      if (type === 'text/plain' || type === 'text/json') {
        const url = await executeGraphQL(async () => {
          const service = useClient(ContentService)
          const request = new SupplementaryIdRequest({ id: parent.metadataId, key: parent.key })
          const url = await service.getMetadataSupplementaryDownloadUrl(request, {
            headers: getGraphQLHeaders(context),
          })
          return url
        })
        if (!url) return null
        const content = await executeHttpRequest(url)
        if (type === 'text/json') {
          return {
            json: JSON.parse(content.toString())
          }
        }
        return {
          text: content.toString()
        }
      }
      return null
    },
  },
  Mutation: {
    addMetadata: async (parent, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const response = await service.addMetadata(
          new AddMetadataRequest({
            collection: '00000000-0000-0000-0000-000000000000',
            metadata: {
              name: args.metadata.name,
              contentType: args.metadata.contentType,
              traitIds: args.metadata.traitIds?.map((s) => s),
              contentLength: args.metadata.contentLength ? protoInt64.parse(args.metadata.contentLength) : undefined,
              languageTag: args.metadata.languageTag,
            },
          }), {
            headers: getGraphQLHeaders(context),
          }
        )
        let lastError: any | null = null
        for (let tries = 0; tries < 100; tries++) {
          try {
            const metadata = await service.getMetadata(new IdRequest({ id: response.id }), {
              headers: getGraphQLHeaders(context),
            })
            return transformMetadata(metadata)
          } catch (e) {
            lastError = e
            await new Promise((resolve) => setTimeout(resolve, 100))
          }
        }
        if (lastError) {
          throw lastError
        }
        throw new GraphQLError('failed to get metadata after it was created')
      })
    },
    setMetadataReady: async (_, args, context) => {
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const response = await service.setMetadataReady(
          new MetadataReadyRequest({
            id: args.id,
          }), {
            headers: getGraphQLHeaders(context),
          }
        )
        const metadata = await service.getMetadata(new IdRequest({ id: args.id }), {
          headers: getGraphQLHeaders(context),
        })
        return transformMetadata(metadata)
      })
    },
    setMetadataJSONContent: async (_, args, context) => {
      if (!args.json) throw new GraphQLError('missing json')
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const idRequest = new IdRequest({ id: args.id })
        const url = await service.getMetadataUploadUrl(idRequest, {
          headers: getGraphQLHeaders(context),
        })
        const content = typeof args.json === 'string' ? args.json : JSON.stringify(args.json)
        await executeHttpRequest(url, toArrayBuffer(content))
        const metadata = await service.getMetadata(idRequest, {
          headers: getGraphQLHeaders(context),
        })
        return transformMetadata(metadata)
      })
    },
    setMetadataTextContent: async (_, args, context) => {
      if (!args.text) throw new GraphQLError('missing json')
      return await executeGraphQL(async () => {
        const service = useClient(ContentService)
        const idRequest = new IdRequest({ id: args.id })
        const url = await service.getMetadataUploadUrl(idRequest, {
          headers: getGraphQLHeaders(context),
        })
        await executeHttpRequest(url, toArrayBuffer(args.text!))
        const metadata = await service.getMetadata(idRequest, {
          headers: getGraphQLHeaders(context),
        })
        return transformMetadata(metadata)
      })
    },
  },
}
