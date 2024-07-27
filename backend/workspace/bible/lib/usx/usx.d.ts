import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item';
import { BookIdentification } from './book_identification';
import { BookHeader } from './book_header';
import { BookTitle } from './book_title';
import { BookIntroduction } from './book_introduction';
import { BookIntroductionEndTitle } from './book_introduction_end_titles';
import { BookChapterLabel } from './book_chapter_label';
import { ChapterStart } from './chapter_start';
import { ChapterEnd } from './chapter_end';
type UsxType = BookIdentification | BookHeader | BookTitle | BookIntroduction | BookIntroductionEndTitle | BookChapterLabel | ChapterStart | ChapterEnd;
export declare class Usx extends UsxItemContainer<UsxType> {
    constructor(context: UsxContext, attributes: Attributes);
}
export declare class UsxFactory extends UsxItemFactory<Usx> {
    private static _instance?;
    static get instance(): UsxFactory;
    private constructor();
    protected onInitialize(): void;
    create(context: UsxContext, attributes: Attributes): Usx;
}
export {};
//# sourceMappingURL=usx.d.ts.map