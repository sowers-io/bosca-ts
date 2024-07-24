import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { Text } from './text';
export declare class Reference extends UsxItemContainer<Text> {
    loc: string;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class ReferenceFactory extends UsxItemFactory<Reference> {
    static readonly instance: ReferenceFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): Reference;
}
//# sourceMappingURL=reference.d.ts.map