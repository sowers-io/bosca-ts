import { Attributes, UsxContext, UsxItem, UsxItemFactory } from './item';
import { Position } from './position';
export declare class ChapterStart implements UsxItem {
    readonly number: string;
    readonly sid: string;
    readonly altnumber?: string;
    readonly pubnumber?: string;
    readonly position: Position;
    readonly verse: string | null;
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class ChapterStartFactory extends UsxItemFactory<ChapterStart> {
    static readonly instance: ChapterStartFactory;
    constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): ChapterStart;
}
//# sourceMappingURL=chapter_start.d.ts.map