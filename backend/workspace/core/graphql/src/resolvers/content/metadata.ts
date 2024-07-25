import { Resolvers, Metadata as GMetadata, SignedUrl as GSignedUrl } from '../../generated/resolvers'
import { RequestContext } from '../../context'
import { useClient } from '@bosca/common'
import { AddMetadataRequest, ContentService, IdRequest, Metadata } from '@bosca/protobufs'
import { execute, getHeaders } from '../../util/requests'

function transformMetadata(metadata: Metadata): GMetadata {
  const m = metadata.toJson() as unknown as GMetadata
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

export const resolvers: Resolvers<RequestContext> = {
  Query: {
    metadata: async (parent, args, context) => {
      return await execute(async () => {
        const service = useClient(ContentService)
        const metadata = await service.getMetadata(new IdRequest({ id: args.id }), {
          headers: getHeaders(context),
        })
        return transformMetadata(metadata)
      })
    },
  },
  Metadata: {
    uploadUrl: async (parent, args, context) => {
      return await execute<GSignedUrl>(async () => {
        const service = useClient(ContentService)
        const url = await service.getMetadataUploadUrl(new IdRequest({ id: parent.id }), {
          headers: getHeaders(context),
        })
        return url.toJson() as unknown as GSignedUrl
      })
    },
    downloadUrl: async (parent, args, context) => {
      return (await execute(async () => {
        const service = useClient(ContentService)
        const url = await service.getMetadataDownloadUrl(new IdRequest({ id: parent.id }), {
          headers: getHeaders(context),
        })
        return url.toJson() as unknown as GSignedUrl
      }))!
    },
  },
  Mutation: {
    addMetadata: async (parent, args, context) => {
      return await execute(async () => {
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
          headers: getHeaders(context),
        })
        return transformMetadata(metadata)
      })
    },
  },
}
