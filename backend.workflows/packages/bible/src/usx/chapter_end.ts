import { Attributes, EndIdFactoryFilter, UsxContext, UsxItem, UsxItemFactory } from './item'
import { Position } from './position'

export class ChapterEnd implements UsxItem {

  readonly eid: string
  readonly position: Position

  readonly verse: string | null

  constructor(context: UsxContext, attributes: Attributes) {
    this.eid = attributes.EID.toString()
    this.verse = context.addVerseItem(this)
    this.position = context.position
  }
}

export class ChapterEndFactory extends UsxItemFactory<ChapterEnd> {

  static readonly instance = new ChapterEndFactory()

  constructor() {
    super('chapter', new EndIdFactoryFilter())
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): ChapterEnd {
    return new ChapterEnd(context, attributes)
  }
}