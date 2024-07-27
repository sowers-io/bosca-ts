import { Reference } from './reference';
import { Milestone } from './milestone';
import { Footnote } from './footnote';
import { Break } from './break';
import { Text } from './text';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { IntroCharStyle } from './styles';
import { Char } from './char';
type IntroCharType = Reference | Char | IntroChar | Milestone | Footnote | Break | Text;
export declare class IntroChar extends UsxItemContainer<IntroCharType> {
    style: IntroCharStyle;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class IntroCharFactory extends UsxItemFactory<IntroChar> {
    static readonly instance: IntroCharFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): IntroChar;
}
export {};
//# sourceMappingURL=intro_char.d.ts.map