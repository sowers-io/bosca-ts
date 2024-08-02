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
import { QdrantClient, GetCollectionInfoRequest, CreateCollection, Distance } from '@qdrant/js-client-grpc'
import { OllamaEmbeddings } from '@langchain/community/embeddings/ollama'
import { protoInt64 } from '@bufbuild/protobuf'
import { Document } from 'langchain/document'
import { QdrantVectorStore } from '@langchain/qdrant'
import { OpenAIEmbeddings } from '@langchain/openai'
import { RecursiveCharacterTextSplitter } from '@langchain/textsplitters'

export class QdrantStorageSystem implements IStorageSystem {
  private readonly indexName: string
  private readonly vectorSize: number

  constructor(configuration: { [key: string]: string }) {
    this.indexName = configuration.indexName
    this.vectorSize = parseInt(configuration.vectorSize)
  }

  async initialize(): Promise<void> {
    const client = new QdrantClient({
      host: process.env.BOSCA_QDRANT_API_ADDRESS!.split(':')[0],
      port: parseInt(process.env.BOSCA_QDRANT_API_ADDRESS!.split(':')[1]),
    })
    try {
      await client.api('collections').get(new GetCollectionInfoRequest({ collectionName: this.indexName }))
    } catch (e) {
      await client.api('collections').create(
        new CreateCollection({
          collectionName: this.indexName,
          vectorsConfig: {
            config: {
              case: 'params',
              value: {
                size: protoInt64.parse(this.vectorSize),
                distance: Distance.Cosine,
              },
            },
          },
        })
      )
    }
  }

  async storeContent(definition: WorkflowJob, metadata: Metadata, content: Buffer): Promise<void> {
    const document = new Document({
      pageContent: content.toString(),
      metadata: { metadataId: metadata.id, ...metadata.attributes },
    })
    const embeddings = new OpenAIEmbeddings({
      apiKey: process.env.OPENAI_KEY!,
    })
    const store = await QdrantVectorStore.fromExistingCollection(embeddings, {
      url: process.env.BOSCA_QDRANT_REST_API_ADDRESS,
      collectionName: this.indexName,
    })
    const textSplitter = new RecursiveCharacterTextSplitter({
      chunkSize: 1000,
      chunkOverlap: 200,
    })
    const chunks = await textSplitter.splitDocuments([document])
    await store.addDocuments(chunks)
  }

  async storePendingEmbeddings(
    definition: WorkflowJob,
    metadata: Metadata,
    embeddings: PendingEmbeddings
  ): Promise<void> {
    const documents = embeddings.embedding.map((e) => {
      return new Document({
        pageContent: e.content || '',
        metadata: { embeddingId: e.id, metadataId: metadata.id, ...metadata.attributes },
      })
    })
    const vectorEmbeddings = new OpenAIEmbeddings({
      apiKey: process.env.OPENAI_KEY!,
    })
    const store = await QdrantVectorStore.fromExistingCollection(vectorEmbeddings, {
      url: process.env.BOSCA_QDRANT_REST_API_ADDRESS,
      collectionName: this.indexName,
    })
    const textSplitter = new RecursiveCharacterTextSplitter({
      chunkSize: 1000,
      chunkOverlap: 200,
    })
    const chunks = await textSplitter.splitDocuments(documents)
    await store.addDocuments(chunks)
  }
}
