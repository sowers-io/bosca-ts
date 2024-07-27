import { Reference } from './reference';
import { Text } from './text';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { CrossReferenceCharStyle } from './styles';
import { Char } from './char';
type CrossReferenceCharType = Char | CrossReferenceChar | Reference | Text;
export declare class CrossReferenceChar extends UsxItemContainer<CrossReferenceCharType> {
    style: CrossReferenceCharStyle;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class CrossReferenceCharFactory extends UsxItemFactory<CrossReferenceChar> {
    static readonly instance: CrossReferenceCharFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): CrossReferenceChar;
}
export {};
//# sourceMappingURL=cross_reference_char.d.ts.map