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

import { Book } from '@bosca/bible'
import { toArrayBuffer } from '../../../util/http'
import { WorkflowActivityJob } from '../../../generated/protobuf/bosca/workflow/execution_context_pb'
import { findFirstMetadata } from '../../../util/finder'
import { Source } from '../../../generated/protobuf/bosca/content/sources_pb'
import { uploadSupplementary } from '../../../util/uploader'
import { BookActivity } from './book_activity'
import { Metadata } from '../../../generated/protobuf/bosca/content/metadata_pb'

export class CreateVerseMarkdownTable extends BookActivity {
  get id(): string {
    return 'bible.book.verse.markdown.table'
  }

  async executeBook(
    source: Source,
    systemId: string,
    metadata: Metadata,
    activity: WorkflowActivityJob,
    book: Book
  ): Promise<void> {
    const table = [['USFM', 'Verse']]
    const bookMetadata = await findFirstMetadata({
      'bible.type': 'book',
      'bible.system.id': systemId,
      'bible.book.usfm': book.usfm,
    })
    for (const chapter of book.chapters) {
      for (const verse of chapter.getVerses(book)) {
        table.push([verse.usfm, verse.items.map((item) => item.toString().trim().replace('\r', '').replace('\n', '')).join(' ')])
      }
    }
    const key = activity.activity!.outputs['supplementaryId']

    let markdown = ''
    let lengths = [0, 0]
    for (let r = 0; r < table.length; r++) {
      const row = table[r]
      for (let c = 0; c < row.length; c++) {
        row[c] = row[c].replace('|', '\\|')
        lengths[c] = Math.max(lengths[c], row[c].length)
      }
    }

    for (let r = 0; r < table.length; r++) {
      const row = table[r]
      markdown += '|'
      for (let c = 0; c < row.length; c++) {
        markdown += ' ' + row[c].padEnd(lengths[c], ' ') + ' |'
      }
      markdown += '\r\n'
      if (r === 0) {
        markdown += '|'
        for (let c = 0; c < row.length; c++) {
          markdown += ' ' + '-'.repeat(lengths[c]) + ' |'
        }
        markdown += '\r\n'
      }
    }

    const buffer = toArrayBuffer(markdown)
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
