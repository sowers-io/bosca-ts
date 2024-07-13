import { IdRequest, IdResponses, SupplementaryIdRequest } from '../generated/protobuf/bosca/requests_pb'
import { getMetadataUploadUrl } from './service'
import { useServiceClient } from './util'
import { ContentService } from '../generated/protobuf/bosca/content/service_connect'
import { execute } from './http'
import { Retry } from './retry'
import { Queue } from './queue'
import { AddSupplementaryRequest } from '../generated/protobuf/bosca/content/metadata_pb'
import { protoInt64 } from '@bufbuild/protobuf'

async function uploadInQueue(queue: Queue, id: string, buffer: ArrayBuffer) {
  await queue.enqueue(() => upload(id, buffer))
}

export async function uploadAll(response: IdResponses, buffers: ArrayBuffer[]) {
  const queue = new Queue('uploadAll', 4)
  for (let ix = 0; ix < response.id.length; ix++) {
    const addResponse = response.id[ix]
    if (addResponse.error) {
      throw new Error(addResponse.error)
    }
    await uploadInQueue(queue, addResponse.id, buffers[ix])
  }
}

export async function upload(id: string, buffer: ArrayBuffer) {
  return Retry.execute(10, async () => {
    const idRequest = new IdRequest({ id: id })
    const uploadUrl = await getMetadataUploadUrl(idRequest)
    const uploadResponse = await execute(uploadUrl, buffer)
    if (!uploadResponse.ok) {
      throw new Error('failed to upload content: ' + (await uploadResponse.text()))
    }
    await useServiceClient(ContentService).setMetadataReady(idRequest)
  })
}

export async function getMetadataSupplementaryUploadUrl(id: SupplementaryIdRequest) {
  return Retry.execute(10, () => useServiceClient(ContentService).getMetadataSupplementaryUploadUrl(id))
}

async function setMetadataSupplementaryUploaded(id: SupplementaryIdRequest) {
  return Retry.execute(10, () => useServiceClient(ContentService).setMetadataSupplementaryReady(id))
}

export async function uploadSupplementary(
  metadataId: string,
  name: string,
  contentType: string,
  key: string,
  sourceId: string | undefined,
  sourceIdentifier: string | undefined,
  buffer: ArrayBuffer
) {
  return Retry.execute(10, async () => {
    const service = useServiceClient(ContentService)
    const request = {
      metadataId: metadataId,
      name: name,
      contentLength: protoInt64.parse(buffer.byteLength),
      contentType: contentType,
      key: key,
      sourceId: sourceId,
      sourceIdentifier: sourceIdentifier,
    }
    const supplementary = await service.addMetadataSupplementary(new AddSupplementaryRequest(request))
    const idRequest = new SupplementaryIdRequest({ id: supplementary.id, key: supplementary.key })
    const uploadUrl = await getMetadataSupplementaryUploadUrl(idRequest)
    const uploadResponse = await execute(uploadUrl, buffer)
    if (!uploadResponse.ok) {
      throw new Error('failed to upload supplementary: ' + (await uploadResponse.text()))
    }
    await setMetadataSupplementaryUploaded(idRequest)
  })
}
