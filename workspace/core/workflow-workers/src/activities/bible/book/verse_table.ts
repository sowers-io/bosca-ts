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

import { Book } from '@bosca/bible-processor'
import { toArrayBuffer } from '../../../util/http'
import { findFirstMetadata } from '../../../util/finder'
import { uploadSupplementary } from '../../../util/uploader'
import { BookActivity, BookExecutor } from './book_activity'
import { Metadata, Source, WorkflowJob } from '@bosca/protobufs'
import { Job } from 'bullmq/dist/esm/classes/job'

export class CreateVerseMarkdownTable extends BookActivity {
  get id(): string {
    return 'bible.book.verse.table.create'
  }

  newBookExecutor(job: Job, definition: WorkflowJob): BookExecutor {
    return new Executor(job, definition)
  }
}

class Executor extends BookExecutor {
  async execute(
    source: Source,
    systemId: string,
    metadata: Metadata,
    book: Book
  ): Promise<void> {
    const verses = []
    const bookMetadata = await findFirstMetadata({
      'bible.type': 'book',
      'bible.system.id': systemId,
      'bible.book.usfm': book.usfm,
    })
    for (const chapter of book.chapters) {
      for (const verse of chapter.getVerses(book)) {
        verses.push({
          usfm: verse.usfm,
          verse: verse.items.map((item) => item.toString().trim().replace('\r', '').replace('\n', '')).join(' '),
        })
      }
    }
    const key = this.definition.activity!.outputs['supplementaryId']
    const buffer = toArrayBuffer(JSON.stringify(verses))
    await uploadSupplementary(
      bookMetadata.id,
      'Verse Markdown Table',
      'text/markdown',
      key,
      source.id,
      undefined,
      buffer
    )
  }
}
