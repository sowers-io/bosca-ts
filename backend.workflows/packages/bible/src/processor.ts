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

import decompress from 'decompress'
import { parseStringPromise } from 'xml2js'
import { BibleMetadata } from './metadata'
import { BookProcessor } from './book_processor';
import { Book } from './usx/book'

export class USXProcessor {

  metadata!: BibleMetadata
  books: Book[] = []

  async process(file: string) {
    const files = await decompress(file)
    const filesMap: { [key: string]: decompress.File } = {}
    for (const file of files) {
      if (file.path == 'metadata.xml') {
        const metadata = await parseStringPromise(file.data)
        this.metadata = new BibleMetadata(metadata.DBLMetadata)
      } else {
        filesMap[file.path] = file
      }
    }

    const processor = new BookProcessor()
    for (const name of this.metadata.publication.names) {
      const content = this.metadata.publication.contents[name.id]
      const file = filesMap[content.file]
      const book = await processor.process(name, content, file.data.toString())
      this.books.push(book)
    }
  }
}