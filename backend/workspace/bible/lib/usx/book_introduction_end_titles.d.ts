import { Footnote } from './footnote';
import { CrossReference } from './cross_reference';
import { Char } from './char';
import { Milestone } from './milestone';
import { Break } from './break';
import { Text } from './text';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { BookIntroductionEndTitleStyle } from './styles';
type BookIntroductionEndTitleType = Footnote | CrossReference | Char | Milestone | Break | Text;
export declare class BookIntroductionEndTitle extends UsxItemContainer<BookIntroductionEndTitleType> {
    style: BookIntroductionEndTitleStyle;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class BookIntroductionEndTitleFactory extends UsxItemFactory<BookIntroductionEndTitle> {
    static readonly instance: BookIntroductionEndTitleFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): BookIntroductionEndTitle;
}
export {};
//# sourceMappingURL=book_introduction_end_titles.d.ts.map