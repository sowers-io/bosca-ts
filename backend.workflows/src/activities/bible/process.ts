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

import { USXProcessor, BibleMetadata, Book } from '@bosca/bible'
import {
  AddCollectionRequest,
  AddCollectionsRequest,
  Collection,
  FindCollectionRequest,
} from '../../generated/protobuf/bosca/content/collections_pb'
import { IdRequest, IdResponse } from '../../generated/protobuf/bosca/requests_pb'
import { useServiceClient } from '../../util/util'
import { ContentService } from '../../generated/protobuf/bosca/content/service_connect'
import { AddMetadataRequest, AddMetadatasRequest, Metadata } from '../../generated/protobuf/bosca/content/metadata_pb'
import { execute, toArrayBuffer } from '../../util/http'
import { protoInt64 } from '@bufbuild/protobuf'
import { Downloader } from '../../util/downloader'
import { getCollection, getMetadata, getMetadataUploadUrl } from '../../util/service'
import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import { Activity } from '../activity'
import { Queue } from '../../util/queue'
import { Retry } from '../../util/retry'
import { upload, uploadAll } from '../../util/uploader'
import { addCollections, addMetadatas } from '../../util/adder'
import { Source } from '../../generated/protobuf/bosca/content/sources_pb'

export class ProcessBibleActivity extends Activity {
  private readonly downloader: Downloader

  constructor(downloader: Downloader) {
    super()
    this.downloader = downloader
  }

  get id(): string {
    return 'bible.process'
  }

  private async createBibleCollection(metadata: BibleMetadata): Promise<Collection> {
    const version = await this.findNextVersion(metadata)
    const service = useServiceClient(ContentService)
    const addResponse = await service.addCollection(
      new AddCollectionRequest({
        collection: new Collection({
          name: metadata.identification.nameLocal,
          attributes: {
            'bible.type': 'bible',
            'bible.language': metadata.language.iso,
            'bible.system.id': metadata.identification.systemId.id,
            'bible.abbreviation': metadata.identification.abbreviationLocal,
            'bible.version': version.toString(),
          },
        }),
      })
    )
    return getCollection(new IdRequest({ id: addResponse.id }))
  }

  private async createBookCollections(
    source: Source,
    metadata: BibleMetadata,
    bible: Collection,
    books: Book[]
  ): Promise<Collection[]> {
    const addCollectionRequests: AddCollectionRequest[] = []
    const addMetadatasRequests: AddMetadataRequest[] = []
    const buffers: ArrayBuffer[] = []
    let order = 0

    // build bulk requests
    for (const book of books) {
      const attributes = {
        'bible.type': 'book',
        'bible.language': metadata.language.iso,
        'bible.version': bible.attributes['bible.version'],
        'bible.system.id': bible.attributes['bible.system.id'],
        'bible.abbreviation': bible.attributes['bible.abbreviation'],
        'bible.book.usfm': book.usfm,
        'bible.book.order': order.toString(),
      }
      addCollectionRequests.push(
        new AddCollectionRequest({
          parent: bible.id,
          collection: new Collection({
            name: book.name.short + ' Chapters',
            attributes: attributes,
          }),
        })
      )
      const buffer = toArrayBuffer(book.raw)
      buffers.push(buffer)
      addMetadatasRequests.push(
        new AddMetadataRequest({
          collection: bible.id,
          metadata: new Metadata({
            name: book.name.short,
            contentType: 'bible/usx-book',
            languageTag: metadata.language.iso,
            contentLength: protoInt64.parse(buffer.byteLength),
            attributes: attributes,
            sourceId: source.id,
            traitIds: ['bible.usx.book'],
          }),
        })
      )
      order++
    }

    // create collections
    const addCollectionResponses = await addCollections(addCollectionRequests)

    // create metadata
    await addMetadatas(addMetadatasRequests, buffers)

    // fetch created collections
    const collections: Collection[] = []
    for (const addResponse of addCollectionResponses.id) {
      if (addResponse.error) {
        throw new Error(addResponse.error)
      }
      // TODO: Add bulk getCollections
      const collection = await getCollection(new IdRequest({ id: addResponse.id }))
      collections.push(collection)
    }
    return collections
  }

  private async createChapters(source: Source, metadata: BibleMetadata, bookCollection: Collection, book: Book) {
    const requests: AddMetadataRequest[] = []
    const buffers: ArrayBuffer[] = []
    let order = 0
    for (const chapter of book.chapters) {
      const buffer = toArrayBuffer(book.raw.substring(chapter.position.start, chapter.position.end))
      buffers.push(buffer)
      requests.push(
        new AddMetadataRequest({
          collection: bookCollection.id,
          metadata: new Metadata({
            name: book.name.short + ' ' + chapter.number,
            contentType: 'bible/usx-chapter',
            contentLength: protoInt64.parse(buffer.byteLength),
            languageTag: metadata.language.iso,
            attributes: {
              'bible.type': 'chapter',
              'bible.language': metadata.language.iso,
              'bible.version': bookCollection.attributes['bible.version'],
              'bible.system.id': bookCollection.attributes['bible.system.id'],
              'bible.abbreviation': bookCollection.attributes['bible.abbreviation'],
              'bible.book.usfm': book.usfm,
              'bible.chapter.usfm': chapter.usfm,
              'bible.book.order': bookCollection.attributes['bible.book.order'],
              'bible.chapter.order': (order++).toString(),
            },
            sourceId: source.id,
            traitIds: ['bible.usx.chapter'],
          }),
        })
      )
    }

    await addMetadatas(requests, buffers)
  }

  private async findNextVersion(metadata: BibleMetadata): Promise<number> {
    let version = 1
    const collections = await useServiceClient(ContentService).findCollection(
      new FindCollectionRequest({
        attributes: {
          'bible.type': 'bible',
          'bible.system.id': metadata.identification.systemId.id,
        },
      })
    )
    for (const collection of collections.collections) {
      version = Math.max(parseInt(collection.attributes['bible.version']), version)
    }
    return version
  }

  private async enqueueCreateChapters(
    queue: Queue,
    source: Source,
    metadata: BibleMetadata,
    collection: Collection,
    book: Book
  ) {
    const creator = this
    queue.enqueue(() => creator.createChapters(source, metadata, collection, book))
    await queue.process()
  }

  async execute(activity: WorkflowActivityJob) {
    const source = await useServiceClient(ContentService).getSource(new IdRequest({ id: 'workflow' }))

    const file = await this.downloader.download(activity)
    try {
      const processor = new USXProcessor()
      await processor.process(file)

      const metadata = processor.metadata!
      const bibleCollection = await this.createBibleCollection(metadata)
      const bookCollections = await this.createBookCollections(source, metadata, bibleCollection, processor.books)

      const queue = new Queue(this.id, 4)
      for (let bookIndex = 0; bookIndex < processor.books.length; bookIndex++) {
        const book = processor.books[bookIndex]
        const collection = bookCollections[bookIndex]
        await this.enqueueCreateChapters(queue, source, metadata, collection, book)
      }
    } finally {
      await this.downloader.cleanup(file)
    }
  }
}
