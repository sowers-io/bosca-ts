import { Reference, ReferenceFactory } from './reference'
import { Footnote, FootnoteFactory } from './footnote'
import { Char, CharFactory } from './char'
import { Break, BreakFactory } from './break'
import { Text, TextFactory } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory, UsxTag } from './item'
import { Verse } from './verse'
import { Milestone, MilestoneFactory } from './milestone'
import { Figure, FigureFactory } from './figure'
import { ParaStyle, ParaStyles } from './styles'
import { CrossReference, CrossReferenceFactory } from './cross_reference'
import { VerseStartFactory } from './verse_start'
import { VerseEndFactory } from './verse_end'

type ParagraphType = Reference | Footnote | CrossReference | Char | Milestone | Figure | Verse | Break | Text

export class Paragraph extends UsxItemContainer<ParagraphType> {

  style: ParaStyle
  vid?: string

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as ParaStyle
    this.vid = attributes.VID?.toString()
  }
}

export class ParagraphFactory extends UsxItemFactory<Paragraph> {

  static readonly instance = new ParagraphFactory()

  constructor() {
    super('para', new StyleFactoryFilter(ParaStyles))
  }

  protected onInitialize() {
    this.register(ReferenceFactory.instance)
    this.register(FootnoteFactory.instance)
    this.register(CrossReferenceFactory.instance)
    this.register(CharFactory.instance)
    this.register(MilestoneFactory.instance)
    this.register(FigureFactory.instance)
    this.register(VerseStartFactory.instance)
    this.register(VerseEndFactory.instance)
    this.register(BreakFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): Paragraph {
    return new Paragraph(context, attributes)
  }
}
