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

import { BookActivity, BookExecutor } from './book_activity'
import { Book } from '@bosca/bible-processor'
import {
  addCollections,
  addMetadatas,
  findFirstCollection,
  toArrayBuffer,
  uploadAll,
  uploadSupplementary,
} from '@bosca/workflow-activities-api'
import { AddCollectionRequest, AddMetadataRequest, Collection, Metadata, Source, WorkflowJob } from '@bosca/protobufs'
import { Job } from 'bullmq/dist/esm/classes/job'
import { protoInt64 } from '@bufbuild/protobuf';

export class CreateVerses extends BookActivity {
  get id(): string {
    return 'bible.chapter.verses.create'
  }

  newBookExecutor(job: Job, definition: WorkflowJob): BookExecutor {
    return new Executor(job, definition)
  }
}

class Executor extends BookExecutor {
  private async createVersesCollectionRequests(
    systemId: string,
    version: string,
    abbreviation: string,
    book: Book,
  ): Promise<AddCollectionRequest[]> {
    const requests: AddCollectionRequest[] = []
    const bookCollection = await findFirstCollection({
      'bible.type': 'book',
      'bible.version': version,
      'bible.system.id': systemId,
      'bible.book.usfm': book.usfm,
    })
    for (const chapter of book.chapters) {
      requests.push(
        new AddCollectionRequest({
          parent: bookCollection.id,
          collection: new Collection({
            name: bookCollection.name + ' ' + chapter.number + ' Verses',
            attributes: {
              'bible.type': 'verses',
              'bible.system.id': systemId,
              'bible.version': version,
              'bible.abbreviation': abbreviation,
              'bible.book.usfm': book.usfm,
              'bible.chapter.usfm': chapter.usfm,
            },
          }),
        }),
      )
    }
    return requests
  }

  override async execute(source: Source, systemId: string, version: string, metadata: Metadata, book: Book): Promise<void> {
    const addCollectionRequests = await this.createVersesCollectionRequests(
      systemId,
      version,
      metadata.attributes['bible.abbreviation'],
      book,
    )
    const collections = await addCollections(addCollectionRequests)

    const addMetadataRequests: AddMetadataRequest[] = []
    const textBuffers: ArrayBuffer[] = []
    const rawBuffers: ArrayBuffer[] = []
    for (let i = 0; i < book.chapters.length; i++) {
      const chapter = book.chapters[i]
      const verses = chapter.getVerses(book)
      const collection = collections.id[i]
      for (let v = 0; v < verses.length; v++) {
        const verse = verses[v]
        const buffer = toArrayBuffer(verse.raw)
        rawBuffers.push(buffer)
        textBuffers.push(toArrayBuffer(verse.toString()))
        addMetadataRequests.push(
          new AddMetadataRequest({
            collection: collection.id,
            metadata: new Metadata({
              name: book.name.short + ' ' + chapter.number + ':' + verse.verse,
              contentType: 'bible/usx-verse',
              languageTag: metadata.languageTag,
              contentLength: protoInt64.parse(buffer.byteLength),
              sourceId: source.id,
              traitIds: ['bible.usx.verse'],
              attributes: {
                'bible.type': 'verse',
                'bible.system.id': systemId,
                'bible.version': version,
                'bible.abbreviation': metadata.attributes['bible.abbreviation'],
                'bible.book.usfm': book.usfm,
                'bible.chapter.usfm': chapter.usfm,
                'bible.verse.usfm': verse.usfm,
                'bible.verse.order': v.toString(),
              },
            }),
          }),
        )
      }
    }
    const responses = await addMetadatas(addMetadataRequests)

    for (let i = 0; i < responses.id.length; i++) {
      if (responses.id[i].id === '') {
        throw new Error('not all metadata were created')
      }
    }

    for (let i = 0; i < responses.id.length; i++) {
      await uploadSupplementary(
        responses.id[i].id,
        'Verse Text',
        'text/plain',
        'text',
        source.id,
        undefined,
        undefined,
        textBuffers[i],
      )
    }
    await uploadAll(responses, rawBuffers)
  }
}
