import { FootnoteStyle } from './styles';
import { Text } from './text';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { FootnoteChar } from './footnote_char';
type FootnoteItem = FootnoteChar | Text;
export declare class Footnote extends UsxItemContainer<FootnoteItem> {
    style: FootnoteStyle;
    caller: string;
    category?: string;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class FootnoteFactory extends UsxItemFactory<Footnote> {
    static readonly instance: FootnoteFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): Footnote;
}
export {};
//# sourceMappingURL=footnote.d.ts.map