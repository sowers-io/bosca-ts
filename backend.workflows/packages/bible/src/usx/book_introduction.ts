import { BookIntroductionStyle, BookIntroductionStyles } from './styles'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { Text, TextFactory } from './text'
import { Milestone, MilestoneFactory } from './milestone'
import { Figure, FigureFactory } from './figure'
import { Reference, ReferenceFactory } from './reference'
import { Footnote, FootnoteFactory } from './footnote'
import { CrossReference, CrossReferenceFactory } from './cross_reference'
import { Char, CharFactory } from './char'
import { IntroChar, IntroCharFactory } from './intro_char'

type BookIntroductionType = Reference | Footnote | CrossReference | Char | IntroChar | Milestone | Figure | Text

export class BookIntroduction extends UsxItemContainer<BookIntroductionType> {
  style: BookIntroductionStyle

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as BookIntroductionStyle
  }
}

export class BookIntroductionFactory extends UsxItemFactory<BookIntroduction> {

  static readonly instance = new BookIntroductionFactory()

  private constructor() {
    super('para', new StyleFactoryFilter(BookIntroductionStyles))
  }

  protected onInitialize() {
    this.register(ReferenceFactory.instance)
    this.register(FootnoteFactory.instance)
    this.register(CrossReferenceFactory.instance)
    this.register(CharFactory.instance)
    this.register(IntroCharFactory.instance)
    this.register(MilestoneFactory.instance)
    this.register(FigureFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): BookIntroduction {
    return new BookIntroduction(context, attributes)
  }
}