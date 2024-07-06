import { Attributes, UsxContext, UsxItem, UsxItemFactory } from './item'
import { Position } from './position'

export class Milestone implements UsxItem {

  readonly style: string
  readonly sid: string
  readonly eid: string

  readonly position: Position
  readonly verse: string | null

  constructor(context: UsxContext, attributes: Attributes) {
    this.position = context.position
    this.style = attributes.STYLE.toString()
    this.sid = attributes.SID.toString()
    this.eid = attributes.EID.toString()
    this.verse = context.addVerseItem(this)
  }
}

export class MilestoneFactory extends UsxItemFactory<Milestone> {

  static readonly instance = new MilestoneFactory()

  private constructor() {
    super('ms')
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): Milestone {
    return new Milestone(context, attributes)
  }
}