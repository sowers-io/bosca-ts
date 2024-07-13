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
  AddSupplementaryRequest,
  FindMetadataRequest,
  Metadata
} from '../../generated/protobuf/bosca/content/metadata_pb'
import { execute, toArrayBuffer } from '../../util/http'
import { protoInt64 } from '@bufbuild/protobuf'
import { IdRequest, SupplementaryIdRequest } from '../../generated/protobuf/bosca/requests_pb'
import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'

interface Verse {
  usfm: string
  chapter: number
  verse: number
  items: UsxItem[]
}

export class CreateVerseMarkdownTable extends Activity {

  private readonly downloader: Downloader

  constructor(downloader: Downloader) {
    super()
    this.downloader = downloader
  }

  get id(): string {
    return 'bible.book.verse.markdown.table'
  }

  private async buildVerseMarkdownTable(chapter: Chapter, table: string[][]) {
    const verses: Verse[] = []
    for (const verse in chapter.verseItems) {
      const usfmSplit = verse.split('.')
      let chapterNumber = parseInt(usfmSplit[1])
      let verseNumber = parseInt(usfmSplit[2])
      if (isNaN(chapterNumber)) chapterNumber = 0
      if (isNaN(verseNumber)) verseNumber = 0
      verses.push({ usfm: verse, chapter: chapterNumber, verse: verseNumber, items: chapter.verseItems[verse] })
    }
    verses.sort((a, b) => {
      if (a.chapter > b.chapter) return 1
      if (a.chapter < b.chapter) return -1
      if (a.verse > b.verse) return 1
      if (a.verse < b.verse) return -1
      return 0
    })
    for (const verse of verses) {
      table.push([
        verse.usfm,
        verse.items.map((item) => item.toString().trim().replace('\r|\n', '')).join(' ')
      ])
    }
  }

  private async findBookMetadata(metadata: BibleMetadata, book: Book): Promise<Metadata> {
    const chapterMetadatas = await useServiceClient(ContentService).findMetadata(new FindMetadataRequest({
      attributes: {
        'bible.type': 'book',
        'bible.system.id': metadata.identification.systemId.id,
        'bible.book.usfm': book.usfm
      }
    }))
    if (chapterMetadatas.metadata.length === 0) {
      throw new Error('failed to find book: ' + book.usfm)
    }
    return chapterMetadatas.metadata[0]
  }

  private async createVerseTable(metadata: BibleMetadata, book: Book) {
    const table = [
      ['USFM', 'Verse']
    ]
    const service = useServiceClient(ContentService)
    const source = await service.getSource(new IdRequest({id: 'workflow'}))
    const bookMetadata = await this.findBookMetadata(metadata, book)
    for (const chapter of book.chapters) {
      await this.buildVerseMarkdownTable(chapter, table)
    }
    const { markdownTable } = await import('markdown-table');
    const markdown = markdownTable(table)
    const buffer = toArrayBuffer(markdown)
    const supplementary = await service.addMetadataSupplementary(new AddSupplementaryRequest({
      metadataId: bookMetadata.id,
      name: 'Verse Markdown Table',
      contentLength: protoInt64.parse(buffer.byteLength),
      contentType: 'text/markdown',
      key: 'verse-table-markdown',
      sourceId: source.id
    }))
    const uploadUrl = await service.getMetadataSupplementaryUploadUrl(new SupplementaryIdRequest({
      id: bookMetadata.id,
      key: supplementary.key
    }))
    const uploadResponse = await execute(uploadUrl, buffer)
    if (!uploadResponse.ok) {
      throw new Error('failed to upload verse table: ' + book.usfm + ': ' + await uploadResponse.text())
    }
    await service.setMetadataSupplementaryReady(new SupplementaryIdRequest({id: bookMetadata.id, key: 'verse-table-markdown'}))
  }

  async execute(activity: WorkflowActivityJob) {
    const file = await this.downloader.download(activity)
    try {
      const processor = new USXProcessor()
      await processor.process(file)
      for (const book of processor.books) {
        await this.createVerseTable(processor.metadata, book)
      }
    } finally {
      await this.downloader.cleanup(file)
    }
  }
}