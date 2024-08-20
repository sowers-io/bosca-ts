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
import { BibleMetadata, ManifestName, PublicationContent } from './metadata'
import { BookProcessor } from './book_processor'
import { Book } from './usx/book'
import { promisify } from 'node:util'
import * as fs from 'node:fs'
import { Chapter } from './usx/chapter'
import { Bible } from './bible'

const readFile = promisify(fs.readFile)

export class USXProcessor {
  metadata?: BibleMetadata
  books: Book[] = []
  booksByUsfm: { [usfm: string]: Book } = {}

  async processBook(name: ManifestName, content: PublicationContent, file: string): Promise<Book> {
    const data = await readFile(file, 'utf-8')
    const processor = new BookProcessor()
    return await processor.process(name, content, data.toString())
  }

  async processChapter(name: ManifestName, content: PublicationContent, file: string): Promise<Chapter> {
    const data = await readFile(file, 'utf-8')
    const processor = new BookProcessor()
    const book = await processor.process(name, content, data.toString())
    return book.chapters[0]
  }

  async process(file: string): Promise<Bible> {
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
    for (const name of this.metadata!.publication.names) {
      const content = this.metadata!.publication.contents[name.id]
      const file = filesMap[content.file]
      const book = await processor.process(name, content, file.data.toString())
      this.books.push(book)
      this.booksByUsfm[book.usfm] = book
    }

    return new Bible(this.books, this.booksByUsfm)
  }
}
