import { BookHeaderStyle } from './styles';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { Text } from './text';
export declare class BookHeader extends UsxItemContainer<Text> {
    style: BookHeaderStyle;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class BookHeaderFactory extends UsxItemFactory<BookHeader> {
    static readonly instance: BookHeaderFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): BookHeader;
}
//# sourceMappingURL=book_header.d.ts.map