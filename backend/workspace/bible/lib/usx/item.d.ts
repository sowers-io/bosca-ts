import { Tag, QualifiedTag, QualifiedAttribute } from 'sax';
import { Position } from './position';
export type Attributes = {
    [key: string]: string | QualifiedAttribute;
};
export type UsxTag = Tag | QualifiedTag;
export interface UsxItem {
    readonly position: Position;
    readonly verse: string | null;
    toString(): string;
}
export declare abstract class UsxItemContainer<T extends UsxItem> implements UsxItem {
    items: T[];
    readonly position: Position;
    readonly verse: string | null;
    protected constructor(context: UsxContext, attributes: Attributes);
    addItem(item: T): void;
    toString(): string;
}
export interface ItemFactoryFilter {
    supports(context: UsxContext, attributes: Attributes): boolean;
}
export declare class StyleFactoryFilter<T> implements ItemFactoryFilter {
    private styles;
    constructor(styles: T[]);
    supports(context: UsxContext, attributes: Attributes): boolean;
}
export declare class CodeFactoryFilter<T> implements ItemFactoryFilter {
    private styles;
    constructor(styles: T[]);
    supports(context: UsxContext, attributes: Attributes): boolean;
}
export declare class EndIdFactoryFilter implements ItemFactoryFilter {
    supports(context: UsxContext, attributes: Attributes): boolean;
}
export declare class NegateFactoryFilter implements ItemFactoryFilter {
    private filter;
    constructor(filter: ItemFactoryFilter);
    supports(context: UsxContext, attributes: Attributes): boolean;
}
export declare class UsxVerseItems {
    readonly usfm: string;
    readonly verse: string;
    readonly items: UsxItem[];
    readonly position: Position;
    constructor(usfm: string, verse: string, position: Position);
    addItem(item: UsxItem): void;
    toString(): string;
}
export declare abstract class UsxContext {
    protected positions: Position[];
    private verses;
    get position(): Position;
    pushVerse(bookChapterUsfm: string, verse: string, position: Position): void;
    popVerse(): UsxVerseItems;
    get verse(): string | null;
    addVerseItem(item: UsxItem): string | null;
    supports(factory: UsxItemFactory<any>, tag: UsxTag): boolean;
}
export declare abstract class UsxItemFactory<T extends UsxItem> {
    readonly tagName: string;
    private filter;
    private factories;
    protected constructor(tagName: string, filter?: ItemFactoryFilter | null);
    private initialized;
    initialize(): void;
    protected abstract onInitialize(): void;
    protected register(factory: UsxItemFactory<any>): void;
    supports(context: UsxContext, attributes: Attributes): boolean;
    abstract create(context: UsxContext, attributes: Attributes): T;
    findChildFactory(context: UsxContext, tag: UsxTag): UsxItemFactory<any>;
}
//# sourceMappingURL=item.d.ts.map