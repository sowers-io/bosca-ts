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

import { IStorageSystem } from './storagesystem'
import { MeilisearchStorageSystem } from './meilisearch'
import { QdrantStorageSystem } from './qdrant'

export async function getStorageSystem(configuration: { [key: string]: string }): Promise<IStorageSystem> {
  let system: IStorageSystem
  switch (configuration.type) {
    case 'meilisearch':
      system = new MeilisearchStorageSystem(configuration)
      break
    case 'qdrant':
      system = new QdrantStorageSystem(configuration)
      break
    default:
      throw new Error('unsupported storage system: ' + configuration.type)
  }
  await system.initialize()
  return system
}