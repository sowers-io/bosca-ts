import { Reference } from './reference';
import { Milestone } from './milestone';
import { Footnote } from './footnote';
import { Break } from './break';
import { Text } from './text';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { ListCharStyle } from './styles';
import { Char } from './char';
type ListCharType = Reference | Char | Milestone | Footnote | Break | Text;
export declare class ListChar extends UsxItemContainer<ListCharType> {
    style: ListCharStyle;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class ListCharFactory extends UsxItemFactory<ListChar> {
    static readonly instance: ListCharFactory;
    constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): ListChar;
}
export {};
//# sourceMappingURL=list_char.d.ts.map