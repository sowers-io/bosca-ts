import { Reference } from './reference';
import { Footnote } from './footnote';
import { Char } from './char';
import { Break } from './break';
import { Text } from './text';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { Verse } from './verse';
import { Milestone } from './milestone';
import { Figure } from './figure';
import { ParaStyle } from './styles';
import { CrossReference } from './cross_reference';
type ParagraphType = Reference | Footnote | CrossReference | Char | Milestone | Figure | Verse | Break | Text;
export declare class Paragraph extends UsxItemContainer<ParagraphType> {
    style: ParaStyle;
    vid?: string;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class ParagraphFactory extends UsxItemFactory<Paragraph> {
    static readonly instance: ParagraphFactory;
    constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): Paragraph;
}
export {};
//# sourceMappingURL=paragraph.d.ts.map