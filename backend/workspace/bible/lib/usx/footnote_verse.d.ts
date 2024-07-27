import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { Text } from './text';
import { FootnoteVerseStyle } from './styles';
export declare class FootnoteVerse extends UsxItemContainer<Text> {
    style: FootnoteVerseStyle;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class FootnoteVerseFactory extends UsxItemFactory<FootnoteVerse> {
    constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): FootnoteVerse;
}
//# sourceMappingURL=footnote_verse.d.ts.map