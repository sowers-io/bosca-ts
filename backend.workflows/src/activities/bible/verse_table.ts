import { Activity } from '../../workflow/workflow'
import { WorkflowActivityExecutionContext } from '../../generated/protobuf/bosca/content/workflows_pb'
import { markdownTable } from 'markdown-table'
import { BibleMetadata, Book, Chapter, UsxItem, USXProcessor } from '@bosca/bible/lib'
import { Downloader } from '../../util/downloader'
import { useServiceClient } from '../../util/util'
import { ContentService } from '../../generated/protobuf/bosca/content/content_connect'
import {
  AddSupplementaryRequest,
  FindMetadataRequest,
  Metadata
} from '../../generated/protobuf/bosca/content/metadata_pb'
import { execute, toArrayBuffer } from '../../util/http'
import { protoInt64 } from '@bufbuild/protobuf'

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
    return 'bible.verse.markdown.table'
  }

  private buildVerseMarkdownTable(chapter: Chapter) {
    const table = [
      ['USFM', 'Verse']
    ]
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
    return markdownTable(table)
  }

  private async findChapterMetadata(metadata: BibleMetadata, chapter: Chapter): Promise<Metadata> {
    const chapterMetadatas = await useServiceClient(ContentService).findMetadata(new FindMetadataRequest({
      attributes: {
        'bible.system.id': metadata.identification.systemId.id,
        'bible.chapter.usfm': chapter.usfm
      }
    }))
    if (chapterMetadatas.metadata.length === 0) {
      throw new Error('failed to find chapter: ' + chapter.usfm)
    }
    return chapterMetadatas.metadata[0]
  }

  private async createVerseTable(metadata: BibleMetadata, book: Book) {
    const service = useServiceClient(ContentService)
    for (const chapter of book.chapters) {
      console.log('generating table for ' + chapter.usfm)
      const chapterMetadata = await this.findChapterMetadata(metadata, chapter)
      const markdown = this.buildVerseMarkdownTable(chapter)
      const buffer = toArrayBuffer(markdown)
      const uploadUrl = await service.addMetadataSupplementary(new AddSupplementaryRequest({
        id: chapterMetadata.id,
        name: 'Verse Markdown Table',
        contentLength: protoInt64.parse(buffer.byteLength),
        contentType: 'text/markdown',
        type: 'verse-table-markdown'
      }))
      const uploadResponse = await execute(uploadUrl, buffer)
      if (!uploadResponse.ok) {
        throw new Error('failed to upload verse table: ' + chapter.usfm + ': ' + await uploadResponse.text())
      }
    }
  }

  async execute(executionContext: WorkflowActivityExecutionContext): Promise<void> {
    const file = await this.downloader.download(executionContext)
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