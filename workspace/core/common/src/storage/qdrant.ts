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
import { Metadata } from '@bosca/protobufs'
import { QdrantClient } from '@qdrant/js-client-rest'

export class QdrantStorageSystem implements IStorageSystem {
  private readonly client = new QdrantClient({ url: process.env.BOSCA_QDRANT_API_ADDRESS! })
  private readonly indexName: string
  private readonly vectorSize: number

  constructor(configuration: { [key: string]: string }) {
    this.indexName = configuration.indexName
    this.vectorSize = parseInt(configuration.vectorSize)
  }

  async register(): Promise<void> {
    const collection = await this.client.getCollection(this.indexName)
    if (collection) return
    await this.client.createCollection(
      this.indexName,
      {
        vectors: {
          size: this.vectorSize,
          distance: 'Cosine'
        }
      })
  }

  async initialize(): Promise<void> {
  }

  async store(metadata: Metadata, content: Buffer): Promise<void> {
    throw new Error('unsupported')
  }
}