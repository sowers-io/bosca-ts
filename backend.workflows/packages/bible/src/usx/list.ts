import { ListStyle, ListStyles } from './styles'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { Reference, ReferenceFactory } from './reference'
import { Footnote, FootnoteFactory } from './footnote'
import { Char, CharFactory } from './char'
import { Milestone, MilestoneFactory } from './milestone'
import { Figure, FigureFactory } from './figure'
import { Verse } from './verse'
import { Break, BreakFactory } from './break'
import { Text, TextFactory } from './text'
import { CrossReference, CrossReferenceFactory } from './cross_reference'
import { ListChar, ListCharFactory } from './list_char'
import { VerseStartFactory } from './verse_start'
import { VerseEndFactory } from './verse_end'

type ListType = Reference | Footnote | CrossReference | Char | ListChar | Milestone | Figure | Verse | Break | Text

export class List extends UsxItemContainer<ListType> {

  style: ListStyle
  vid: string

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as ListStyle
    this.vid = attributes.VID.toString()
  }
}

export class ListFactory extends UsxItemFactory<List> {

  static readonly instance = new ListFactory()

  private constructor() {
    super('para', new StyleFactoryFilter(ListStyles))
  }

  protected onInitialize() {
    this.register(ReferenceFactory.instance)
    this.register(FootnoteFactory.instance)
    this.register(CrossReferenceFactory.instance)
    this.register(CharFactory.instance)
    this.register(ListCharFactory.instance)
    this.register(MilestoneFactory.instance)
    this.register(FigureFactory.instance)
    this.register(VerseStartFactory.instance)
    this.register(VerseEndFactory.instance)
    this.register(BreakFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): List {
    return new List(context, attributes)
  }
}