import { FootnoteStyle } from './styles'
import { Text, TextFactory } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { FootnoteStyles } from './styles'
import { FootnoteChar, FootnoteCharFactory } from './footnote_char'

type FootnoteItem = FootnoteChar | Text

export class Footnote extends UsxItemContainer<FootnoteItem> {
  style: FootnoteStyle
  caller: string
  category?: string

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as FootnoteStyle
    this.caller = attributes.CALLER.toString()
    this.category = attributes.CATEGORY?.toString()
  }
}

export class FootnoteFactory extends UsxItemFactory<Footnote> {

  static readonly instance = new FootnoteFactory()

  private constructor() {
    super('note', new StyleFactoryFilter(FootnoteStyles))
  }

  protected onInitialize() {
    this.register(FootnoteCharFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): Footnote {
    return new Footnote(context, attributes)
  }
}

