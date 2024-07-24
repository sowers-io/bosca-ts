import { BookIntroductionStyle } from './styles';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { Text } from './text';
import { Milestone } from './milestone';
import { Figure } from './figure';
import { Reference } from './reference';
import { Footnote } from './footnote';
import { CrossReference } from './cross_reference';
import { Char } from './char';
import { IntroChar } from './intro_char';
type BookIntroductionType = Reference | Footnote | CrossReference | Char | IntroChar | Milestone | Figure | Text;
export declare class BookIntroduction extends UsxItemContainer<BookIntroductionType> {
    style: BookIntroductionStyle;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class BookIntroductionFactory extends UsxItemFactory<BookIntroduction> {
    static readonly instance: BookIntroductionFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): BookIntroduction;
}
export {};
//# sourceMappingURL=book_introduction.d.ts.map