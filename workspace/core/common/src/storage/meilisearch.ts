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
import { Metadata, PendingEmbeddings, WorkflowJob } from '@bosca/protobufs'
import { MeiliSearch, Index } from 'meilisearch'
import { logger } from '../logger';

export class MeilisearchStorageSystem implements IStorageSystem {
  private readonly client = new MeiliSearch({
    host: process.env.BOSCA_MEILISEARCH_API_ADDRESS!,
    apiKey: process.env.BOSCA_MEILI_MASTER_KEY!,
  })
  private readonly indexName: string
  private readonly primaryKey: string
  private index!: Index

  constructor(configuration: { [key: string]: string }) {
    this.indexName = configuration.indexName
    this.primaryKey = configuration.primaryKey
  }

  async register(): Promise<void> {
    try {
      await this.client.getIndex(this.indexName)
    } catch (e: any) {
      if (e.message === 'Index `metadata` not found.') {
        await this.client.createIndex(this.indexName, {
          primaryKey: this.primaryKey,
        })
      } else {
        throw e
      }
    }
  }

  async initialize(): Promise<void> {
    try {
      this.index = await this.client.getIndex(this.indexName)
    } catch (e: any) {
      if (e.message === 'Index `' + this.indexName + '` not found.') {
        const task = await this.client.createIndex(this.indexName, {
          primaryKey: this.primaryKey,
        })
        for (let tries = 0; tries < 10; tries++) {
          try {
            await this.client.tasks.waitForTask(task.taskUid)
            this.index = await this.client.getIndex(this.indexName)
            return
          } catch (e) {
            logger.warn({ error: e }, 'failed to create index')
          }
        }
        this.index = await this.client.getIndex(this.indexName)
      } else {
        throw e
      }
    }
  }

  async storeContent(definition: WorkflowJob, metadata: Metadata, content: Buffer): Promise<void> {
    const document: Record<string, any> = {
      content: content.toString(),
    }
    for (const key in metadata.attributes) {
      document[key.replace(/\./gi, '_')] = metadata.attributes[key]
    }
    document[this.primaryKey] = metadata.id
    await this.index.addDocuments([document])
  }

  async storePendingEmbeddings(
    _: WorkflowJob,
    __: Metadata,
    ___: PendingEmbeddings,
  ): Promise<void> {
    throw new Error('unsupported')
  }
}
