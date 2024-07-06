import { SidebarStyle } from './styles'
import { Attributes, UsxContext, UsxItemContainer } from './item'
import { Paragraph } from './paragraph'
import { List } from './list'
import { Table } from './table'
import { Footnote } from './footnote'
import { CrossReference } from './cross_reference'

type SidebarType = Paragraph | List | Table | Footnote | CrossReference

export class Sidebar extends UsxItemContainer<SidebarType> {

  style: SidebarStyle
  category?: String

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as SidebarStyle
    this.category = attributes.CATEGORY?.toString()
  }
}