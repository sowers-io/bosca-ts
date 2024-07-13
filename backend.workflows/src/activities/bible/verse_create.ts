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

import { Activity } from '../activity'
import { BibleMetadata, Book, USXProcessor } from '@bosca/bible/lib'
import { Downloader } from '../../util/downloader'
import { useServiceClient } from '../../util/util'
import { ContentService } from '../../generated/protobuf/bosca/content/service_connect'
import { AddMetadataRequest, AddMetadatasRequest, Metadata } from '../../generated/protobuf/bosca/content/metadata_pb'
import { execute, toArrayBuffer } from '../../util/http'
import { protoInt64 } from '@bufbuild/protobuf'
import { IdRequest } from '../../generated/protobuf/bosca/requests_pb'
import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import {
  AddCollectionRequest,
  AddCollectionsRequest,
  Collection,
  FindCollectionRequest,
} from '../../generated/protobuf/bosca/content/collections_pb'
import { Queue } from '../../util/queue'
import { findFirstCollection } from '../../util/finder'
import { addCollections, addMetadatas } from '../../util/adder'
import { uploadAll } from '../../util/uploader'
import { Source } from '../../generated/protobuf/bosca/content/sources_pb'

export class CreateVerses extends Activity {
  private readonly downloader: Downloader

  constructor(downloader: Downloader) {
    super()
    this.downloader = downloader
  }

  get id(): string {
    return 'bible.chapter.verses.create'
  }

  private async createVersesCollectionRequests(metadata: BibleMetadata, book: Book): Promise<AddCollectionRequest[]> {
    const requests: AddCollectionRequest[] = []
    for (const chapter of book.chapters) {
      const bookCollection = await findFirstCollection({
        'bible.type': 'book',
        'bible.system.id': metadata.identification.systemId.id,
        'bible.book.usfm': book.usfm,
      })
      requests.push(
        new AddCollectionRequest({
          parent: bookCollection.id,
          collection: new Collection({
            name: book.name.short + ' ' + chapter.number + ' Verses',
            attributes: {
              'bible.type': 'verses',
              'bible.system.id': metadata.identification.systemId.id,
              'bible.abbreviation': metadata.identification.abbreviationLocal,
              'bible.book.usfm': book.usfm,
              'bible.chapter.usfm': chapter.usfm,
            },
          }),
        })
      )
    }
    return requests
  }

  private async createVerses(source: Source, metadata: BibleMetadata, book: Book) {
    const addCollectionRequests = await this.createVersesCollectionRequests(metadata, book)
    const collections = await addCollections(addCollectionRequests)

    const addMetadataRequests: AddMetadataRequest[] = []
    const buffers: ArrayBuffer[] = []
    for (let i = 0; i < book.chapters.length; i++) {
      const chapter = book.chapters[i]
      const verses = chapter.getVerses(book)
      const collection = collections.id[i]
      for (let v = 0; v < verses.length; v++) {
        const verse = verses[v]
        const buffer = toArrayBuffer(verse.raw)
        buffers.push(buffer)
        addMetadataRequests.push(
          new AddMetadataRequest({
            collection: collection.id,
            metadata: new Metadata({
              name: book.name.short + ' ' + chapter.number + ':' + verse.verse,
              contentType: 'bible/usx-verse',
              languageTag: metadata.language.iso,
              contentLength: protoInt64.parse(buffer.byteLength),
              sourceId: source.id,
              traitIds: ['bible.usx.verse'],
              attributes: {
                'bible.type': 'verse',
                'bible.system.id': metadata.identification.systemId.id,
                'bible.abbreviation': metadata.identification.abbreviationLocal,
                'bible.book.usfm': book.usfm,
                'bible.chapter.usfm': chapter.usfm,
                'bible.verse.usfm': verse.usfm,
                'bible.verse.order': v.toString(),
              },
            }),
          })
        )
      }
    }

    await addMetadatas(addMetadataRequests, buffers)
  }

  private async enqueue(queue: Queue, source: Source, metadata: BibleMetadata, book: Book) {
    const creator = this
    await queue.enqueue(() => creator.createVerses(source, metadata, book))
  }

  async execute(activity: WorkflowActivityJob) {
    const source = await useServiceClient(ContentService).getSource(new IdRequest({ id: 'workflow' }))
    const file = await this.downloader.download(activity)
    try {
      const processor = new USXProcessor()
      await processor.process(file)
      const queue = new Queue(this.id, 4)
      for (const book of processor.books) {
        await this.enqueue(queue, source, processor.metadata, book)
      }
    } finally {
      await this.downloader.cleanup(file)
    }
  }
}
