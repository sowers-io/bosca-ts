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
import { toArrayBuffer } from '../../util/http'
import { IdRequest } from '../../generated/protobuf/bosca/requests_pb'
import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import { findFirstMetadata } from '../../util/finder'
import { Source } from '../../generated/protobuf/bosca/content/sources_pb'
import { uploadSupplementary } from '../../util/uploader'
import { Retry } from '../../util/retry'
import { Queue } from '../../util/queue'

export class CreateVerseMarkdownTable extends Activity {
  private readonly downloader: Downloader

  constructor(downloader: Downloader) {
    super()
    this.downloader = downloader
  }

  get id(): string {
    return 'bible.book.verse.markdown.table'
  }

  private async createVerseTable(source: Source, metadata: BibleMetadata, book: Book, key: string) {
    await Retry.execute(10, async () => {
      const table = [['USFM', 'Verse']]
      const bookMetadata = await findFirstMetadata({
        'bible.type': 'book',
        'bible.system.id': metadata.identification.systemId.id,
        'bible.book.usfm': book.usfm,
      })
      for (const chapter of book.chapters) {
        for (const verse of chapter.getVerses(book)) {
          table.push([verse.usfm, verse.items.map((item) => item.toString().trim().replace('\r|\n', '')).join(' ')])
        }
      }
      const { markdownTable } = await import('markdown-table')
      const markdown = markdownTable(table)
      const buffer = toArrayBuffer(markdown)
      await uploadSupplementary(
        bookMetadata.id,
        'Verse Markdown Table',
        'text/markdown',
        'verse-table-markdown',
        source.id,
        undefined,
        buffer
      )
    })
  }

  private async enqueueBook(queue: Queue, source: Source, processor: USXProcessor, book: Book, key: string) {
    const creator = this
    await queue.enqueue(() => creator.createVerseTable(source, processor.metadata, book, key))
  }

  async execute(activity: WorkflowActivityJob) {
    const file = await this.downloader.download(activity)
    try {
      const source = await useServiceClient(ContentService).getSource(new IdRequest({ id: 'workflow' }))
      const key = activity.activity!.outputs['supplementaryId']!.value.value as string
      const processor = new USXProcessor()
      await processor.process(file)
      const queue = new Queue(this.id, 4)
      for (const book of processor.books) {
        await this.enqueueBook(queue, source, processor, book, key)
      }
    } finally {
      await this.downloader.cleanup(file)
    }
  }
}
