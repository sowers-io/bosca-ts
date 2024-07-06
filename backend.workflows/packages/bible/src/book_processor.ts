import { ManifestName, PublicationContent } from './metadata'
import { parser } from 'sax'
import { UsxContext, UsxItem, UsxItemContainer, UsxItemFactory, UsxTag, UsxVerseItems } from './usx/item'
import { Book } from './usx/book'
import { UsxFactory } from './usx/usx'
import { BookIdentificationFactory } from './usx/book_identification'
import { BookTitleFactory } from './usx/book_title'
import { BookIntroductionFactory } from './usx/book_introduction'
import { BookHeaderFactory } from './usx/book_header'
import { BookIntroductionEndTitleFactory } from './usx/book_introduction_end_titles'
import { BookChapterLabelFactory } from './usx/book_chapter_label'
import { Text } from './usx/text'
import { ChapterEnd } from './usx/chapter_end'
import { Chapter } from './usx/chapter'
import { ChapterStart } from './usx/chapter_start'
import { VerseStart } from './usx/verse_start'
import { VerseEnd } from './usx/verse_end'
import { Position } from './usx/position'

class UsxNode {
  readonly factory: UsxItemFactory<any>
  readonly item: UsxItem
  readonly position: number

  constructor(factory: UsxItemFactory<any>, item: UsxItem, position: number) {
    this.factory = factory
    this.item = item
    this.position = position
  }
}

enum BookSections {
  identification,
  headers,
  titles,
  introduction,
  endTitles,
  label
}

class BookProcessorContext extends UsxContext {

  private readonly book: Book
  private readonly chapters: Chapter[] = []
  private readonly items: UsxNode[] = []
  private readonly sections: BookSections[] = []
  private readonly verseItems: UsxVerseItems[] = []

  constructor(book: Book) {
    super()
    this.book = book
  }

  supports(factory: UsxItemFactory<any>, tag: UsxTag): boolean {
    if (factory === BookIdentificationFactory.instance) {
      if (this.sections.includes(BookSections.identification)) {
        return false
      }
    } else if (factory === BookHeaderFactory.instance) {
      if (!this.sections.includes(BookSections.identification)) {
        return false
      }
    } else if (factory === BookTitleFactory.instance) {
      if (!this.sections.includes(BookSections.headers)) {
        return false
      }
    } else if (factory === BookIntroductionFactory.instance) {
      if (!this.sections.includes(BookSections.titles)) {
        return false
      }
    } else if (factory === BookIntroductionEndTitleFactory.instance) {
      if (!this.sections.includes(BookSections.introduction)) {
        return false
      }
    } else if (factory === BookChapterLabelFactory.instance) {
      if (!this.sections.includes(BookSections.introduction)) {
        return false
      }
    }
    return super.supports(factory, tag)
  }

  addText(text: string, position: number) {
    this.positions.push(new Position(position))
    if (this.items.length === 0) return
    const node = this.push({ name: '#text', attributes: {}, isSelfClosing: true }, position)
    if (node instanceof Text) {
      node.text = text
    } else {
      throw new Error('unsupported text node')
    }
    this.pop(position + text.length)
  }

  private onFactoryPush(factory: UsxItemFactory<any>) {
    if (factory === BookIdentificationFactory.instance) {
      this.sections.push(BookSections.identification)
    } else if (factory === BookHeaderFactory.instance) {
      this.sections.push(BookSections.headers)
    } else if (factory === BookTitleFactory.instance) {
      this.sections.push(BookSections.titles)
    } else if (factory === BookIntroductionFactory.instance) {
      this.sections.push(BookSections.introduction)
    } else if (factory === BookIntroductionEndTitleFactory.instance) {
      this.sections.push(BookSections.endTitles)
    } else if (factory === BookChapterLabelFactory.instance) {
      this.sections.push(BookSections.label)
    }
  }

  push(tag: UsxTag, position: number): UsxItem {
    this.positions.push(new Position(position))
    if (tag.name.toLowerCase() === 'usx') {
      const factory = UsxFactory.instance
      const item = factory.create(this, tag.attributes)
      this.items.push(new UsxNode(factory, item, position))
      return item
    }
    if (this.items.length === 0) {
      throw new Error('empty stack, invalid state')
    }
    const node = this.items[this.items.length - 1]
    const factory = node.factory.findChildFactory(this, tag)
    const item = factory.create(this, tag.attributes)
    if (item instanceof ChapterStart) {
      this.positions.push(new Position(position))
      const chapter = new Chapter(this, this.book, item)
      this.positions.pop()
      this.chapters.push(chapter)
      this.items.push(new UsxNode(node.factory, chapter, position))
    } else if (item instanceof VerseStart) {
      this.pushVerse(this.chapters[this.chapters.length - 1].usfm, item.number)
    } else if (item instanceof VerseEnd) {
      this.verseItems.push(this.popVerse())
    }
    if (node.item instanceof UsxItemContainer) {
      node.item.addItem(item)
      this.onFactoryPush(factory)
    }
    this.items.push(new UsxNode(factory, item, position))
    return item
  }

  pop(position: number) {
    const lastPosition = this.positions.pop()!
    lastPosition.end = position
    const node = this.items.pop()
    if (node && node.item instanceof ChapterEnd) {
      const chapterNode = this.items.pop()
      if (!chapterNode) throw new Error('missing chapter node')
      const chapter = chapterNode.item as Chapter
      chapter.end = node.item
      chapter.position.start = chapter.start.position.start
      chapter.position.end = node.item.position.end
      chapter.addVerseItems(this.verseItems.slice(0, this.verseItems.length))
      this.book.chapters.push(chapter)
    }
  }
}

export class BookProcessor {

  async process(name: ManifestName, content: PublicationContent, data: string): Promise<Book> {
    const book = new Book(name, content, data)
    const context = new BookProcessorContext(book)
    await new Promise((resolve, reject) => {
      const saxParser = parser()
      saxParser.onerror = reject
      saxParser.onopentag = (tag) => {
        context.push(tag, saxParser.startTagPosition - 1)
      }
      saxParser.ontext = (text) => {
        context.addText(text, saxParser.startTagPosition)
      }
      saxParser.onclosetag = (_) => {
        context.pop(saxParser.position)
      }
      saxParser.onend = () => {
        resolve(null)
      }
      saxParser.write(data).close()
    })
    return book
  }
}

