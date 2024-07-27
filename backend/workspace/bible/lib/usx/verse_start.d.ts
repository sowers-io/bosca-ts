import { Attributes, UsxContext, UsxItem, UsxItemFactory } from './item';
import { VerseStartStyle } from './styles';
import { Position } from './position';
export declare class VerseStart implements UsxItem {
    readonly style: VerseStartStyle;
    readonly number: string;
    readonly altnumber?: string;
    readonly pubnumber?: string;
    readonly sid: string;
    readonly verse: string | null;
    readonly position: Position;
    constructor(context: UsxContext, attributes: Attributes);
    toString(): string;
}
export declare class VerseStartFactory extends UsxItemFactory<VerseStart> {
    static readonly instance: VerseStartFactory;
    constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): VerseStart;
}
//# sourceMappingURL=verse_start.d.ts.map