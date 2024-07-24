import { Attributes, UsxContext, UsxItem, UsxItemFactory } from './item';
import { Position } from './position';
export declare class Break implements UsxItem {
    readonly verse: string | null;
    readonly position: Position;
    constructor(context: UsxContext);
}
export declare class BreakFactory extends UsxItemFactory<Break> {
    static readonly instance: BreakFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): Break;
}
//# sourceMappingURL=break.d.ts.map