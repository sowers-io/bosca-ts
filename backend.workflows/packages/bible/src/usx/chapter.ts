import { UsxContext, UsxItem, UsxItemContainer, UsxVerseItems } from './item'
import { Paragraph } from './paragraph'
import { List } from './list'
import { Table } from './table'
import { Footnote } from './footnote'
import { CrossReference } from './cross_reference'
import { Sidebar } from './sidebar'
import { ChapterStart } from './chapter_start'
import { ChapterEnd } from './chapter_end'
import { Book } from './book'

export type ChapterType = Paragraph | List | Table | Footnote | CrossReference | Sidebar | ChapterEnd

export class Chapter extends UsxItemContainer<ChapterType> {

  readonly number: string
  readonly usfm: string
  readonly verseItems: { [verse: string]: UsxItem[] } = {}
  readonly start: ChapterStart
  private _end?: ChapterEnd

  constructor(context: UsxContext, book: Book, start: ChapterStart) {
    super(context, {})
    this.number = start.number
    this.usfm = book.usfm + '.' + start.number
    this.start = start
  }

  get end(): ChapterEnd {
    return this._end!
  }

  set end(end: ChapterEnd) {
    if (this._end) throw new Error('end already defined')
    this._end = end
  }

  addVerseItems(items: UsxVerseItems[]) {
    for (const item of items) {
      this.verseItems[item.usfm] = item.items
    }
  }

  addItem(item: ChapterType) {
    if (item instanceof ChapterEnd) return
    super.addItem(item)
  }
}