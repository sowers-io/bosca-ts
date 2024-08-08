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

import { execute } from './http'
import { Retry } from './retry'
import { protoInt64 } from '@bufbuild/protobuf'
import {
  AddSupplementaryRequest,
  ContentService,
  IdRequest,
  IdResponses,
  MetadataSupplementary,
  SupplementaryIdRequest,
} from '@bosca/protobufs'
import { logger, useServiceAccountClient } from '@bosca/common'
import { Code, ConnectError } from '@connectrpc/connect'

let uploading = 0

export async function uploadAll(response: IdResponses, buffers: ArrayBuffer[]) {
  for (let ix = 0; ix < response.id.length; ix++) {
    const addResponse = response.id[ix]
    if (addResponse.error) {
      if (addResponse.error != 'name must be unique') {
        throw new Error(addResponse.error)
      }
      logger.error(addResponse.error)
    }
    await upload(addResponse.id, buffers[ix])
  }
}

export async function upload(id: string, buffer: ArrayBuffer) {
  return Retry.execute(10, async () => {
    uploading++
    logger.trace({ id, uploading, length: buffer.byteLength }, 'starting upload')
    try {
      const idRequest = new IdRequest({ id: id })
      const uploadUrl = await useServiceAccountClient(ContentService).getMetadataUploadUrl(idRequest)
      logger.trace({ id, uploading, length: buffer.byteLength, headers: uploadUrl.headers }, 'starting upload')
      await execute(uploadUrl, buffer)
      try {
        await useServiceAccountClient(ContentService).setMetadataReady(idRequest)
      } catch (e) {
        if (e instanceof ConnectError) {
          if (e.message === '[failed_precondition] workflow already in state') {
            logger.warn(
              {
                id,
                uploading,
                length: buffer.byteLength,
              },
              'finished upload, metadata workflow state was already set',
            )
            return
          }
        }
        throw e
      }
      logger.trace({ id, uploading, length: buffer.byteLength }, 'finished upload')
    } catch (e) {
      logger.error({ id, uploading, error: e, length: buffer.byteLength }, 'upload error')
      throw e
    } finally {
      uploading--
    }
  })
}

export async function getMetadataSupplementaryUploadUrl(id: SupplementaryIdRequest) {
  return Retry.execute(10, async () => {
    try {
      return useServiceAccountClient(ContentService).getMetadataSupplementaryUploadUrl(id)
    } catch (e) {
      if (e instanceof ConnectError && e.code == Code.NotFound) {
        return null
      }
      throw e
    }
  })
}

async function setMetadataSupplementaryUploaded(id: SupplementaryIdRequest) {
  return Retry.execute(10, () => useServiceAccountClient(ContentService).setMetadataSupplementaryReady(id))
}

export async function uploadSupplementary(
  metadataId: string,
  name: string,
  contentType: string,
  key: string,
  sourceId: string | undefined,
  sourceIdentifier: string | undefined,
  traitIds: string[] | undefined,
  buffer: ArrayBuffer,
) {
  const supplementary = await Retry.execute(10, async () => {
    const service = useServiceAccountClient(ContentService)
    const idRequest = new SupplementaryIdRequest({ id: metadataId, key: key })
    let supplementary: MetadataSupplementary | undefined
    try {
      supplementary = await service.getMetadataSupplementary(idRequest)
    } catch (e) {
      if (e instanceof ConnectError && e.code == Code.NotFound) {
        supplementary = undefined
      } else {
        throw e
      }
    }
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
      traitIds: traitIds,
    }
    try {
      return await service.addMetadataSupplementary(new AddSupplementaryRequest(request))
    } catch (e: any) {
      if (
        e.toString() ===
        'ConnectError: [unknown] ERROR: duplicate key value violates unique constraint "metadata_supplementary_pkey" (SQLSTATE 23505)'
      ) {
        await service.deleteMetadataSupplementary(idRequest)
      }
      throw e
    }
  })
  return Retry.execute(10, async () => {
    const idRequest = new SupplementaryIdRequest({ id: supplementary.metadataId, key: supplementary.key })
    const uploadUrl = await getMetadataSupplementaryUploadUrl(idRequest)
    await execute(uploadUrl, buffer)
    await setMetadataSupplementaryUploaded(idRequest)
  })
}
