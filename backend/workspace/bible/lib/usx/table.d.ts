import { Attributes, UsxContext, UsxItemContainer } from './item';
import { Footnote } from './footnote';
import { Char } from './char';
import { Milestone } from './milestone';
import { Figure } from './figure';
import { Break } from './break';
import { Text } from './text';
import { CrossReference } from './cross_reference';
import { Verse } from './verse';
export declare class Table extends UsxItemContainer<Row> {
    vid: string;
    get element(): string;
    constructor(context: UsxContext, attributes: Attributes);
}
type RowType = Verse | TableContent;
export declare class Row extends UsxItemContainer<RowType> {
    style: string;
    get element(): string;
    constructor(context: UsxContext, attributes: Attributes);
}
type TableContentType = Footnote | CrossReference | Char | Milestone | Figure | Verse | Break | Text;
export declare class TableContent extends UsxItemContainer<TableContentType> {
    style: string;
    align: string;
    colspan: string;
    get element(): string;
    constructor(context: UsxContext, attributes: Attributes);
}
export {};
//# sourceMappingURL=table.d.ts.map