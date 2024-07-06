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