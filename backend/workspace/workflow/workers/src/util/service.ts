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
import { Collection, ContentService, IdRequest, Metadata, SignedUrl } from '@bosca/protobufs'
import { useServiceAccountClient } from '@bosca/common'

export async function getCollection(id: IdRequest): Promise<Collection> {
  return Retry.execute(10, () => useServiceAccountClient(ContentService).getCollection(id))
}

export async function getMetadata(id: IdRequest): Promise<Metadata> {
  return Retry.execute(10, () => useServiceAccountClient(ContentService).getMetadata(id))
}

export async function getMetadataUploadUrl(id: IdRequest): Promise<SignedUrl> {
  return Retry.execute(10, () => useServiceAccountClient(ContentService).getMetadataUploadUrl(id))
}
