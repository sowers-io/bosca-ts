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

import { Retry } from './retry'
import { uploadAll } from './uploader'
import {
  AddCollectionRequest,
  AddCollectionsRequest,
  AddMetadataRequest, AddMetadatasRequest,
  ContentService,
  IdResponses,
} from '@bosca/protobufs'
import { useServiceAccountClient } from '@bosca/common'

export async function addCollections(addCollectionRequests: AddCollectionRequest[]): Promise<IdResponses> {
  return Retry.execute(10, () =>
    useServiceAccountClient(ContentService).addCollections(
      new AddCollectionsRequest({
        collections: addCollectionRequests,
      }),
    ),
  )
}

export async function addMetadatas(
  addMetadataRequests: AddMetadataRequest[],
  buffers: ArrayBuffer[] | null = null,
): Promise<IdResponses> {
  return await Retry.execute(10, async () => {
    const responses = await useServiceAccountClient(ContentService).addMetadatas(
      new AddMetadatasRequest({
        metadatas: addMetadataRequests,
      }),
    )
    if (buffers) {
      await uploadAll(responses, buffers)
    }
    return responses
  })
}
