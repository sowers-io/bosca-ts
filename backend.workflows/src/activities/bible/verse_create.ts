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
import { BibleMetadata, Book, Chapter, UsxItem, USXProcessor } from '@bosca/bible/lib'
import { Downloader } from '../../util/downloader'
import { useServiceClient } from '../../util/util'
import { ContentService } from '../../generated/protobuf/bosca/content/service_connect'
import {
  AddMetadataRequest,
  AddSupplementaryRequest,
  FindMetadataRequest,
  Metadata
} from '../../generated/protobuf/bosca/content/metadata_pb'
import { execute, toArrayBuffer } from '../../util/http'
import { protoInt64 } from '@bufbuild/protobuf'
import { IdRequest, SupplementaryIdRequest } from '../../generated/protobuf/bosca/requests_pb'
import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import {
  AddCollectionRequest,
  Collection,
  FindCollectionRequest
} from "../../generated/protobuf/bosca/content/collections_pb";

interface Verse {
  usfm: string
  chapter: number
  verse: number
  verseString: string
  items: UsxItem[]
  raw: string
}

export class CreateVerseMarkdownTable extends Activity {

  private readonly downloader: Downloader

  constructor(downloader: Downloader) {
    super()
    this.downloader = downloader
  }

  get id(): string {
    return 'bible.chapter.verses.create'
  }

  private async findBookCollection(metadata: BibleMetadata, book: Book): Promise<Collection> {
    const result = await useServiceClient(ContentService).findCollection(new FindCollectionRequest({
      attributes: {
        'bible.type': 'book',
        'bible.system.id': metadata.identification.systemId.id,
        'bible.book.usfm': book.usfm
      }
    }))
    if (result.collections.length === 0) {
      throw new Error('failed to find book: ' + book.usfm)
    }
    return result.collections[0]
  }

  private async buildVerses(book: Book, chapter: Chapter): Promise<Verse[]> {
    const verses: Verse[] = []
    for (const verse in chapter.verseItems) {
      const usfmSplit = verse.split('.')
      let chapterNumber = parseInt(usfmSplit[1])
      let verseNumber = parseInt(usfmSplit[2])
      if (isNaN(chapterNumber)) chapterNumber = 0
      if (isNaN(verseNumber)) verseNumber = 0
      const verseItems = chapter.verseItems[verse]
      const rawVerse = book.getRawContent(verseItems.position)
      verses.push({ usfm: verse, chapter: chapterNumber, verse: verseNumber, verseString: usfmSplit[2], items: verseItems, raw: rawVerse })
    }
    verses.sort((a, b) => {
      if (a.chapter > b.chapter) return 1
      if (a.chapter < b.chapter) return -1
      if (a.verse > b.verse) return 1
      if (a.verse < b.verse) return -1
      return 0
    })
    return verses
  }

  private async createVerses(metadata: BibleMetadata, book: Book) {
    const service = useServiceClient(ContentService)
    const source = await service.getSource(new IdRequest({id: 'workflow'}))
    for (const chapter of book.chapters) {
      const bookCollection = await this.findBookCollection(metadata, book)
      const verses = await this.buildVerses(book, chapter)

      const collection = await service.addCollection(new AddCollectionRequest({
        parent: bookCollection.id,
        collection: new Collection({
          name: book.name.short + ' ' + chapter.number + ' Verses',
          attributes: {
            'bible.type': 'verses',
            'bible.system.id': metadata.identification.systemId.id,
            'bible.abbreviation': metadata.identification.abbreviationLocal,
            'bible.book.usfm': book.usfm,
            'bible.chapter.usfm': chapter.usfm
          }
        })
      }))

      let order = 0
      for (const verse of verses) {
        const buffer = toArrayBuffer(verse.raw)
        const verseMetadata = await service.addMetadata(new AddMetadataRequest({
          collection: bookCollection.id,
          metadata: new Metadata({
            name: book.name.short + ' ' + chapter.number + ':' + verse.verseString,
            contentType: 'bible/usx-verse',
            languageTag: metadata.language.iso,
            contentLength: protoInt64.parse(buffer.byteLength),
            sourceId: source.id,
            attributes: {
              'bible.type': 'verse',
              'bible.system.id': metadata.identification.systemId.id,
              'bible.abbreviation': metadata.identification.abbreviationLocal,
              'bible.book.usfm': book.usfm,
              'bible.chapter.usfm': chapter.usfm,
              'bible.verse.usfm': verse.usfm
            }
          })
        }))
        const uploadUrl = await service.getMetadataUploadUrl(new IdRequest({id: verseMetadata.id}))
        const uploadResponse = await execute(uploadUrl, buffer)
        if (!uploadResponse.ok) {
          throw new Error('failed to upload verse table: ' + book.usfm + ': ' + await uploadResponse.text())
        }
        await service.setMetadataReady(new IdRequest({id: verseMetadata.id}))
      }
    }
  }

  async execute(activity: WorkflowActivityJob) {
    const file = await this.downloader.download(activity)
    try {
      const processor = new USXProcessor()
      await processor.process(file)
      for (const book of processor.books) {
        await this.createVerses(processor.metadata, book)
      }
    } finally {
      await this.downloader.cleanup(file)
    }
  }
}