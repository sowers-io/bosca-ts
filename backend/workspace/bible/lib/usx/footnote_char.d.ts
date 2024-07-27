import { Char } from './char';
import { Reference } from './reference';
import { Text } from './text';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { FootnoteCharStyle } from './styles';
import { FootnoteVerse } from './footnote_verse';
type FootnoteCharType = Char | FootnoteChar | FootnoteVerse | Reference | Text;
export declare class FootnoteChar extends UsxItemContainer<FootnoteCharType> {
    style: FootnoteCharStyle;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class FootnoteCharFactory extends UsxItemFactory<FootnoteChar> {
    static readonly instance: FootnoteCharFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): FootnoteChar;
}
export {};
//# sourceMappingURL=footnote_char.d.ts.map