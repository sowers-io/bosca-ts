import { Char } from './char'
import { Reference } from './reference'
import { Text, TextFactory } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { FootnoteCharStyle, FootnoteCharStyles } from './styles'
import { FootnoteVerse } from './footnote_verse'

type FootnoteCharType = Char | FootnoteChar | FootnoteVerse | Reference | Text

export class FootnoteChar extends UsxItemContainer<FootnoteCharType> {

  style: FootnoteCharStyle
  // char.link?,
  // char.closed?,

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as FootnoteCharStyle
  }
}

export class FootnoteCharFactory extends UsxItemFactory<FootnoteChar> {

  static readonly instance = new FootnoteCharFactory()

  private constructor() {
    super('char', new StyleFactoryFilter(FootnoteCharStyles))
  }

  protected onInitialize() {
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): FootnoteChar {
    return new FootnoteChar(context, attributes)
  }
}
