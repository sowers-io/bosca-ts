import { ManifestName, PublicationContent } from '../metadata';
import { Chapter } from './chapter';
import { Position } from "./position";
export declare class Book {
    readonly name: ManifestName;
    readonly content: PublicationContent;
    readonly chapters: Chapter[];
    readonly raw: string;
    constructor(name: ManifestName, content: PublicationContent, raw: string);
    get usfm(): string;
    getRawContent(position: Position): string;
}
//# sourceMappingURL=book.d.ts.map