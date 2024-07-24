import { Attributes, UsxContext, UsxItem, UsxItemFactory } from './item';
import { Position } from './position';
export declare class ChapterEnd implements UsxItem {
    readonly eid: string;
    readonly position: Position;
    readonly verse: string | null;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class ChapterEndFactory extends UsxItemFactory<ChapterEnd> {
    static readonly instance: ChapterEndFactory;
    constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): ChapterEnd;
}
//# sourceMappingURL=chapter_end.d.ts.map