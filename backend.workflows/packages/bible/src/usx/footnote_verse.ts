import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { Text } from './text'
import { FootnoteVerseStyle, FootnoteVerseStyles } from './styles'

export class FootnoteVerse extends UsxItemContainer<Text> {

  style: FootnoteVerseStyle

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as FootnoteVerseStyle
  }
}

export class FootnoteVerseFactory extends UsxItemFactory<FootnoteVerse> {

  constructor() {
    super('char', new StyleFactoryFilter(FootnoteVerseStyles))
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): FootnoteVerse {
    return new FootnoteVerse(context, attributes)
  }
}