import { ListStyle } from './styles';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { Reference } from './reference';
import { Footnote } from './footnote';
import { Char } from './char';
import { Milestone } from './milestone';
import { Figure } from './figure';
import { Verse } from './verse';
import { Break } from './break';
import { Text } from './text';
import { CrossReference } from './cross_reference';
import { ListChar } from './list_char';
type ListType = Reference | Footnote | CrossReference | Char | ListChar | Milestone | Figure | Verse | Break | Text;
export declare class List extends UsxItemContainer<ListType> {
    style: ListStyle;
    vid: string;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class ListFactory extends UsxItemFactory<List> {
    static readonly instance: ListFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): List;
}
export {};
//# sourceMappingURL=list.d.ts.map