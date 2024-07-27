import { CharStyle } from './styles';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { Footnote } from './footnote';
import { Break } from './break';
import { Text } from './text';
import { Reference } from './reference';
import { Milestone } from './milestone';
type CharType = Reference | Char | Milestone | Footnote | Break | Text;
export declare class Char extends UsxItemContainer<CharType> {
    style: CharStyle;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class CharFactory extends UsxItemFactory<Char> {
    static readonly instance: CharFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): Char;
}
export {};
//# sourceMappingURL=char.d.ts.map