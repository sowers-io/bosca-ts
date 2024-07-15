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
