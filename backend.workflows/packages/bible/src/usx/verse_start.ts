import { Attributes, StyleFactoryFilter, UsxContext, UsxItem, UsxItemFactory } from './item'
import { VerseStartStyle, VerseStartStyles } from './styles'
import { Position } from './position'

export class VerseStart implements UsxItem {

  readonly style: VerseStartStyle
  readonly number: string
  readonly altnumber?: string
  readonly pubnumber?: string
  readonly sid: string
  readonly verse: string | null
  readonly position: Position

  constructor(context: UsxContext, attributes: Attributes) {
    this.position = context.position
    this.style = attributes.STYLE.toString() as VerseStartStyle
    this.number = attributes.NUMBER.toString()
    this.altnumber = attributes.ALTNUMBER?.toString()
    this.pubnumber = attributes.PUBNUMBER?.toString()
    this.sid = attributes.SID.toString()
    this.verse = context.addVerseItem(this)
  }

  toString(): string {
    return this.number + ". "
  }
}

export class VerseStartFactory extends UsxItemFactory<VerseStart> {

  static readonly instance = new VerseStartFactory()

  constructor() {
    super('verse', new StyleFactoryFilter(VerseStartStyles))
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): VerseStart {
    return new VerseStart(context, attributes)
  }
}