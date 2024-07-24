import { BibleMetadata, ManifestName, PublicationContent } from './metadata';
import { Book } from './usx/book';
import { Chapter } from './usx/chapter';
export declare class USXProcessor {
    metadata?: BibleMetadata;
    books: Book[];
    processBook(name: ManifestName, content: PublicationContent, file: string): Promise<Book>;
    processChapter(name: ManifestName, content: PublicationContent, file: string): Promise<Chapter>;
    process(file: string): Promise<void>;
}
//# sourceMappingURL=processor.d.ts.map