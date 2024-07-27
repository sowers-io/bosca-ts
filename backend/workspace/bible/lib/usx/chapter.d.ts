import { UsxContext, UsxItemContainer, UsxVerseItems } from './item';
import { Paragraph } from './paragraph';
import { List } from './list';
import { Table } from './table';
import { Footnote } from './footnote';
import { CrossReference } from './cross_reference';
import { Sidebar } from './sidebar';
import { ChapterStart } from './chapter_start';
import { ChapterEnd } from './chapter_end';
import { Book } from './book';
export type ChapterType = Paragraph | List | Table | Footnote | CrossReference | Sidebar | ChapterEnd;
export declare class ChapterVerse {
    readonly usfm: string;
    readonly chapter: string;
    readonly verse: string;
    readonly items: UsxVerseItems[];
    readonly raw: string;
    constructor(usfm: string, chapter: string, verse: string, items: UsxVerseItems[], raw: string);
}
export declare class Chapter extends UsxItemContainer<ChapterType> {
    readonly number: string;
    readonly usfm: string;
    readonly verseItems: {
        [verse: string]: UsxVerseItems[];
    };
    readonly start: ChapterStart;
    private _end?;
    constructor(context: UsxContext, book: Book, start: ChapterStart);
    get end(): ChapterEnd;
    set end(end: ChapterEnd);
    addVerseItems(items: UsxVerseItems[]): void;
    addItem(item: ChapterType): void;
    getVerses(book: Book): ChapterVerse[];
}
//# sourceMappingURL=chapter.d.ts.map