import { Attributes, UsxContext, UsxItem, UsxItemFactory } from './item'
import { Position } from './position'

export class Text implements UsxItem {
  text: string = ''

  readonly position: Position
  readonly verse: string | null

  constructor(context: UsxContext) {
    this.position = context.position
    this.verse = context.addVerseItem(this)
  }

  toString(): string {
    return this.text
  }
}

export class TextFactory extends UsxItemFactory<Text> {

  static readonly instance = new TextFactory()

  private constructor() {
    super('#text')
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): Text {
    return new Text(context)
  }
}