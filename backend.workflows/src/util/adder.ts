import { AddCollectionRequest, AddCollectionsRequest } from '../generated/protobuf/bosca/content/collections_pb'
import { IdResponses } from '../generated/protobuf/bosca/requests_pb'
import { Retry } from './retry'
import { useServiceClient } from './util'
import { ContentService } from '../generated/protobuf/bosca/content/service_connect'
import { AddMetadataRequest, AddMetadatasRequest } from '../generated/protobuf/bosca/content/metadata_pb'
import { uploadAll } from './uploader'

export async function addCollections(addCollectionRequests: AddCollectionRequest[]): Promise<IdResponses> {
  return Retry.execute(10, () =>
    useServiceClient(ContentService).addCollections(
      new AddCollectionsRequest({
        collections: addCollectionRequests,
      })
    )
  )
}

export async function addMetadatas(
  addMetadataRequests: AddMetadataRequest[],
  buffers: ArrayBuffer[] | null = null
): Promise<IdResponses> {
  const responses = await Retry.execute(10, () =>
    useServiceClient(ContentService).addMetadatas(
      new AddMetadatasRequest({
        metadatas: addMetadataRequests,
      })
    )
  )
  if (buffers) {
    await uploadAll(responses, buffers)
  }
  return responses
}
