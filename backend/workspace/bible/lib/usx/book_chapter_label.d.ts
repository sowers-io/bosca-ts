import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { Text } from './text';
import { BookChapterLabelStyle } from './styles';
export declare class BookChapterLabel extends UsxItemContainer<Text> {
    style: BookChapterLabelStyle;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class BookChapterLabelFactory extends UsxItemFactory<BookChapterLabel> {
    static readonly instance: BookChapterLabelFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): BookChapterLabel;
}
//# sourceMappingURL=book_chapter_label.d.ts.map