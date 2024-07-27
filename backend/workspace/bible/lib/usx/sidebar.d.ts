import { SidebarStyle } from './styles';
import { Attributes, UsxContext, UsxItemContainer } from './item';
import { Paragraph } from './paragraph';
import { List } from './list';
import { Table } from './table';
import { Footnote } from './footnote';
import { CrossReference } from './cross_reference';
type SidebarType = Paragraph | List | Table | Footnote | CrossReference;
export declare class Sidebar extends UsxItemContainer<SidebarType> {
    style: SidebarStyle;
    category?: String;
    constructor(context: UsxContext, attributes: Attributes);
}
export {};
//# sourceMappingURL=sidebar.d.ts.map