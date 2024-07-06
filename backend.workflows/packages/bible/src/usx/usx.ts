import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { BookIdentification, BookIdentificationFactory } from './book_identification'
import { BookHeader, BookHeaderFactory } from './book_header'
import { BookTitle, BookTitleFactory } from './book_title'
import { BookIntroduction, BookIntroductionFactory } from './book_introduction'
import { BookIntroductionEndTitle, BookIntroductionEndTitleFactory } from './book_introduction_end_titles'
import { BookChapterLabel, BookChapterLabelFactory } from './book_chapter_label'
import { ChapterStart, ChapterStartFactory } from './chapter_start'
import { ChapterEnd, ChapterEndFactory } from './chapter_end'
import { ParagraphFactory } from './paragraph'
import { ListFactory } from './list'
import { FootnoteFactory } from './footnote'
import { CrossReferenceFactory } from './cross_reference'
import { TextFactory } from './text'

/*
element usx {
        attribute version { xsd:string { minLength = "3" pattern = "\d+\.\d+(\.\d+)?"} },
        attribute xsi:noNamespaceSchemaLocation { text }?,

        BookIdentification,
        BookHeaders*,
        BookTitles+,
        BookIntroduction*,
        BookIntroductionEndTitles*,
        BookChapterLabel?,
        Chapter,
        # Chapter is used to separate intro from scripture text.
        # All books will have chapter including the single chapter books: OBA, PHM, 2JN, 3JN, JUD
        ChapterContent+
    }
 */

type UsxType =
  BookIdentification
  | BookHeader
  | BookTitle
  | BookIntroduction
  | BookIntroductionEndTitle
  | BookChapterLabel
  | ChapterStart
  | ChapterEnd

export class Usx extends UsxItemContainer<UsxType> {

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
  }
}

export class UsxFactory extends UsxItemFactory<Usx> {

  private static _instance?: UsxFactory

  static get instance(): UsxFactory {
    if (this._instance == null) {
      this._instance = new UsxFactory()
      this._instance.initialize()
    }
    return this._instance
  }

  private constructor() {
    super('usx')
  }

  protected onInitialize() {
    this.register(BookIdentificationFactory.instance)
    this.register(BookHeaderFactory.instance)
    this.register(BookTitleFactory.instance)
    this.register(BookIntroductionFactory.instance)
    this.register(BookIntroductionEndTitleFactory.instance)
    this.register(BookChapterLabelFactory.instance)
    this.register(ChapterStartFactory.instance)
    this.register(ChapterEndFactory.instance)
    this.register(ParagraphFactory.instance)
    this.register(ListFactory.instance)
    // this.register(Table)
    this.register(FootnoteFactory.instance)
    this.register(CrossReferenceFactory.instance)
    // this.register(Sidebar)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): Usx {
    return new Usx(context, attributes)
  }
}