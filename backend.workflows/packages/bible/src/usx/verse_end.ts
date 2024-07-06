import { Attributes, EndIdFactoryFilter, UsxContext, UsxItem, UsxItemFactory } from './item'
import { Position } from './position'

export class VerseEnd implements UsxItem {

  eid: string
  readonly verse: string | null
  readonly position: Position

  constructor(context: UsxContext, attributes: Attributes) {
    this.position = context.position
    this.eid = attributes.EID.toString()
    this.verse = context.verse
  }

  toString(): string {
    return ''
  }
}

export class VerseEndFactory extends UsxItemFactory<VerseEnd> {

  static readonly instance = new VerseEndFactory()

  constructor() {
    super('verse', new EndIdFactoryFilter())
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): VerseEnd {
    return new VerseEnd(context, attributes)
  }
}