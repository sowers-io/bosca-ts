import { Footnote } from './footnote'
import { CrossReference } from './cross_reference'
import { Char } from './char'
import { Break } from './break'
import { Text, TextFactory } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { BookTitleStyle, BookTitleStyles } from './styles'

type BookTitleType = Footnote | CrossReference | Char | Break | Text

export class BookTitle extends UsxItemContainer<BookTitleType> {
  style: BookTitleStyle

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as BookTitleStyle
  }
}

export class BookTitleFactory extends UsxItemFactory<BookTitle> {

  static readonly instance = new BookTitleFactory()

  private constructor() {
    super('para', new StyleFactoryFilter(BookTitleStyles))
  }

  protected onInitialize() {
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): BookTitle {
    return new BookTitle(context, attributes)
  }
}