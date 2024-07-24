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

import { useServiceClient } from './util'
import { Retry } from './retry'
import { Collection, ContentService, FindCollectionRequest, FindMetadataRequest, Metadata } from '@bosca/protobufs'

export async function findAllCollections(attributes: { [key: string]: string }): Promise<Collection[]> {
  return Retry.execute(10, async () => {
    const result = await useServiceClient(ContentService).findCollection(
        new FindCollectionRequest({ attributes: attributes })
    )
    return result.collections
  })
}

export async function findFirstCollection(attributes: { [key: string]: string }): Promise<Collection> {
  return Retry.execute(10, async () => {
    const result = await useServiceClient(ContentService).findCollection(
        new FindCollectionRequest({ attributes: attributes })
    )
    if (result.collections.length === 0) {
      throw new Error('Collection not found')
    }
    return result.collections[0]
  })
}

export async function findAllMetadatas(attributes: { [key: string]: string }): Promise<Metadata[]> {
  return Retry.execute(10, async () => {
    const result = await useServiceClient(ContentService).findMetadata(
        new FindMetadataRequest({ attributes: attributes })
    )
    return result.metadata
  })
}

export async function findFirstMetadata(attributes: { [key: string]: string }): Promise<Metadata> {
  return Retry.execute(10, async () => {
    const result = await useServiceClient(ContentService).findMetadata(
        new FindMetadataRequest({ attributes: attributes })
    )
    if (result.metadata.length === 0) {
      throw new Error('Metadata not found')
    }
    return result.metadata[0]
  })
}
