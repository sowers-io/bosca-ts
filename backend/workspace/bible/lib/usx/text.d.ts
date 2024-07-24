import { Attributes, UsxContext, UsxItem, UsxItemFactory } from './item';
import { Position } from './position';
export declare class Text implements UsxItem {
    text: string;
    readonly position: Position;
    readonly verse: string | null;
    constructor(context: UsxContext);
    toString(): string;
}
export declare class TextFactory extends UsxItemFactory<Text> {
    static readonly instance: TextFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): Text;
}
//# sourceMappingURL=text.d.ts.map