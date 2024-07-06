import { BookHeaderStyle } from './styles'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItem, UsxItemContainer, UsxItemFactory } from './item'
import { Text, TextFactory } from './text'
import { BookHeaderStyles } from './styles'

export class BookHeader extends UsxItemContainer<Text> {
  style: BookHeaderStyle

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as BookHeaderStyle
  }
}

export class BookHeaderFactory extends UsxItemFactory<BookHeader> {

  static readonly instance = new BookHeaderFactory()

  private constructor() {
    super('para', new StyleFactoryFilter(BookHeaderStyles))
  }

  protected onInitialize() {
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): BookHeader {
    return new BookHeader(context, attributes)
  }
}