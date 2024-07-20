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
  queue.enqueue(() => upload(id, buffer))
  await queue.process()
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
    console.log('starting upload:', id, 'length:', buffer.byteLength)
    try {
      const idRequest = new IdRequest({id: id})
      const uploadUrl = await getMetadataUploadUrl(idRequest)
      console.log('uploading:', id, 'length:', buffer.byteLength, 'headers:', uploadUrl.headers)
      const uploadResponse = await execute(uploadUrl, buffer)
      if (!uploadResponse.ok) {
        throw new Error('failed to upload content: ' + (await uploadResponse.text()))
      }
      await useServiceClient(ContentService).setMetadataReady(idRequest)
      console.log('finished upload:', id, 'length:', buffer.byteLength)
    } catch (e) {
      console.log('upload failed:', id, 'length:', buffer.byteLength)
      throw e
    }
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
  const supplementary = await Retry.execute(10, async () => {
    const service = useServiceClient(ContentService)
    const idRequest = new SupplementaryIdRequest({ id: metadataId, key: key })
    const supplementary = await service.getMetadataSupplementary(idRequest)
    if (supplementary && supplementary.metadataId && supplementary.metadataId.length > 0) {
      await service.deleteMetadataSupplementary(idRequest)
    }
    const request = {
      metadataId: metadataId,
      name: name,
      contentLength: protoInt64.parse(buffer.byteLength),
      contentType: contentType,
      key: key,
      sourceId: sourceId,
      sourceIdentifier: sourceIdentifier,
    }
    console.error(request)
    try {
      const result = await service.addMetadataSupplementary(new AddSupplementaryRequest(request))
      console.error(result)
      return result
    } catch (e: any) {
      if (e.toString() === 'ConnectError: [unknown] ERROR: duplicate key value violates unique constraint "metadata_supplementary_pkey" (SQLSTATE 23505)') {
        await service.deleteMetadataSupplementary(idRequest)
      }
      throw e
    }
  })
  return Retry.execute(10, async () => {
    const idRequest = new SupplementaryIdRequest({ id: supplementary.metadataId, key: supplementary.key })
    const uploadUrl = await getMetadataSupplementaryUploadUrl(idRequest)
    const uploadResponse = await execute(uploadUrl, buffer)
    if (!uploadResponse.ok) {
      throw new Error('failed to upload supplementary: ' + (await uploadResponse.text()))
    }
    await setMetadataSupplementaryUploaded(idRequest)
  })
}
