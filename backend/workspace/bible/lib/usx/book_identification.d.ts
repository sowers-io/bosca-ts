import { BookIdentificationCode } from '../identification';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { Text } from './text';
export declare class BookIdentification extends UsxItemContainer<Text> {
    id: string;
    code: BookIdentificationCode;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class BookIdentificationFactory extends UsxItemFactory<BookIdentification> {
    static readonly instance: BookIdentificationFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): BookIdentification;
}
//# sourceMappingURL=book_identification.d.ts.map