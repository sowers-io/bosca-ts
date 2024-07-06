import { Footnote, FootnoteFactory } from './footnote'
import { CrossReference, CrossReferenceFactory } from './cross_reference'
import { Char, CharFactory } from './char'
import { Milestone, MilestoneFactory } from './milestone'
import { Break, BreakFactory } from './break'
import { Text, TextFactory } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { BookIntroductionEndTitleStyle, BookIntroductionEndTitleStyles } from './styles'


type BookIntroductionEndTitleType = Footnote | CrossReference | Char | Milestone | Break | Text

export class BookIntroductionEndTitle extends UsxItemContainer<BookIntroductionEndTitleType> {
  style: BookIntroductionEndTitleStyle

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as BookIntroductionEndTitleStyle
  }
}

export class BookIntroductionEndTitleFactory extends UsxItemFactory<BookIntroductionEndTitle> {

  static readonly instance = new BookIntroductionEndTitleFactory()

  private constructor() {
    super('para', new StyleFactoryFilter(BookIntroductionEndTitleStyles))
  }

  protected onInitialize() {
    this.register(FootnoteFactory.instance)
    this.register(CrossReferenceFactory.instance)
    this.register(CharFactory.instance)
    this.register(MilestoneFactory.instance)
    this.register(BreakFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): BookIntroductionEndTitle {
    return new BookIntroductionEndTitle(context, attributes)
  }
}