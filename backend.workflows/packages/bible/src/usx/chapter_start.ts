import { Attributes, EndIdFactoryFilter, NegateFactoryFilter, UsxContext, UsxItem, UsxItemFactory } from './item'
import { Position } from './position'

export class ChapterStart implements UsxItem {

  readonly number: string
  readonly sid: string
  readonly altnumber?: string
  readonly pubnumber?: string
  readonly position: Position

  readonly verse: string | null

  constructor(context: UsxContext, attributes: Attributes) {
    this.verse = context.addVerseItem(this)
    this.number = attributes.NUMBER.toString()
    this.sid = attributes.SID.toString()
    this.altnumber = attributes.ALTNUMBER?.toString()
    this.pubnumber = attributes.PUBNUMBER?.toString()
    this.position = context.position
  }
}

export class ChapterStartFactory extends UsxItemFactory<ChapterStart> {

  static readonly instance = new ChapterStartFactory()

  constructor() {
    super('chapter', new NegateFactoryFilter(new EndIdFactoryFilter()))
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): ChapterStart {
    return new ChapterStart(context, attributes)
  }
}