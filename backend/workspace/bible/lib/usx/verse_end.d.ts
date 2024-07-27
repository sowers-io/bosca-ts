import { Attributes, UsxContext, UsxItem, UsxItemFactory } from './item';
import { Position } from './position';
export declare class VerseEnd implements UsxItem {
    eid: string;
    readonly verse: string | null;
    readonly position: Position;
    constructor(context: UsxContext, attributes: Attributes);
    toString(): string;
}
export declare class VerseEndFactory extends UsxItemFactory<VerseEnd> {
    static readonly instance: VerseEndFactory;
    constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): VerseEnd;
}
//# sourceMappingURL=verse_end.d.ts.map