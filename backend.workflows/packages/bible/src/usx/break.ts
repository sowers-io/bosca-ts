import { Attributes, UsxContext, UsxItem, UsxItemFactory } from './item'
import { Position } from './position'

export class Break implements UsxItem {

  readonly verse: string | null
  readonly position: Position

  constructor(context: UsxContext) {
    this.position = context.position
    this.verse = context.addVerseItem(this)
  }
}

export class BreakFactory extends UsxItemFactory<Break> {

  static readonly instance = new BreakFactory()

  private constructor() {
    super('optbreak')
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): Break {
    return new Break(context)
  }
}