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
import { MeiliSearch, Index } from 'meilisearch'

export class MeilisearchStorageSystem implements IStorageSystem {

  private readonly client = new MeiliSearch({
    host: process.env.BOSCA_MEILISEARCH_API_ADDRESS!,
    apiKey: process.env.BOSCA_MEILISEARCH_KEY!
  })
  private readonly indexName: string
  private readonly primaryKey: string
  private index!: Index

  constructor(configuration: { [key: string]: string }) {
    this.indexName = configuration.indexName
    this.primaryKey = configuration.primaryKey
  }

  async register(): Promise<void> {
    const index = await this.client.getIndex(this.indexName)
    if (index) return
    await this.client.createIndex(
      this.indexName,
      {
        primaryKey: this.primaryKey
      }
    )
  }

  async initialize(): Promise<void> {
    this.index = await this.client.getIndex(this.indexName)
  }

  async store(metadata: Metadata, content: Buffer): Promise<void> {
    const document: Record<string, any> = {
      'content': content.toString(),
      ...metadata.attributes
    }
    document[this.primaryKey] = metadata.id
    await this.index.addDocuments([document])
  }
}