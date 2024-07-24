import { CrossReferenceChar } from './cross_reference_char';
import { Text } from './text';
import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { CrossReferenceStyle } from './styles';
type CrossReferenceType = CrossReferenceChar | Text;
export declare class CrossReference extends UsxItemContainer<CrossReferenceType> {
    style: CrossReferenceStyle;
    caller: string;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class CrossReferenceFactory extends UsxItemFactory<CrossReference> {
    static readonly instance: CrossReferenceFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): CrossReference;
}
export {};
//# sourceMappingURL=cross_reference.d.ts.map