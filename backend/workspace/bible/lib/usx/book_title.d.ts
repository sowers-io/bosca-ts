import { Footnote } from './footnote';
import { CrossReference } from './cross_reference';
import { Char } from './char';
import { Break } from './break';
import { Text } from './text';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { BookTitleStyle } from './styles';
type BookTitleType = Footnote | CrossReference | Char | Break | Text;
export declare class BookTitle extends UsxItemContainer<BookTitleType> {
    style: BookTitleStyle;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class BookTitleFactory extends UsxItemFactory<BookTitle> {
    static readonly instance: BookTitleFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): BookTitle;
}
export {};
//# sourceMappingURL=book_title.d.ts.map