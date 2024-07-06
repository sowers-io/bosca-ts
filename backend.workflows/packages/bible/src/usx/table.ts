import { Attributes, UsxContext, UsxItemContainer } from './item'
import { Footnote } from './footnote'
import { Char } from './char'
import { Milestone } from './milestone'
import { Figure } from './figure'
import { Break } from './break'
import { Text } from './text'
import { CrossReference } from './cross_reference'
import { Verse } from './verse'

export class Table extends UsxItemContainer<Row> {

  vid: string

  get element(): string {
    return 'table'
  }

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.vid = attributes.VID.toString()
  }
}

type RowType = Verse | TableContent

export class Row extends UsxItemContainer<RowType> {

  style: string

  get element(): string {
    return 'row'
  }

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString()
  }
}

type TableContentType = Footnote | CrossReference | Char | Milestone | Figure | Verse | Break | Text

export class TableContent extends UsxItemContainer<TableContentType> {
  style: string
  align: string
  colspan: string

  get element(): string {
    return 'cell'
  }

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString()
    this.align = attributes.ALIGN.toString()
    this.colspan = attributes.COLSPAN?.toString()
  }
}