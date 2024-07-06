import { USXProcessor } from './processor';

test('USX Processor', async () => {
  const processor = new USXProcessor()
  await processor.process('../../../example-data/asv.zip')

  console.log(processor.metadata)

  const book = processor.books[0]
  const chapter = book.chapters[0]

  console.log(chapter.toString())

  console.log(book.raw.substring(chapter.position.start, chapter.position.end))
})