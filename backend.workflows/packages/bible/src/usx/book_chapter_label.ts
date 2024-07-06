import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { Text, TextFactory } from './text'
import { BookChapterLabelStyle, BookChapterLabelStyles } from './styles'

export class BookChapterLabel extends UsxItemContainer<Text> {
  style: BookChapterLabelStyle

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as BookChapterLabelStyle
  }
}

export class BookChapterLabelFactory extends UsxItemFactory<BookChapterLabel> {

  static readonly instance = new BookChapterLabelFactory()

  private constructor() {
    super('para', new StyleFactoryFilter(BookChapterLabelStyles))
    this.register(TextFactory.instance)
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): BookChapterLabel {
    return new BookChapterLabel(context, attributes)
  }
}