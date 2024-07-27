"use strict";
/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.BookProcessor = void 0;
const sax_1 = require("sax");
const item_1 = require("./usx/item");
const book_1 = require("./usx/book");
const usx_1 = require("./usx/usx");
const book_identification_1 = require("./usx/book_identification");
const book_title_1 = require("./usx/book_title");
const book_introduction_1 = require("./usx/book_introduction");
const book_header_1 = require("./usx/book_header");
const book_introduction_end_titles_1 = require("./usx/book_introduction_end_titles");
const book_chapter_label_1 = require("./usx/book_chapter_label");
const text_1 = require("./usx/text");
const chapter_end_1 = require("./usx/chapter_end");
const chapter_1 = require("./usx/chapter");
const chapter_start_1 = require("./usx/chapter_start");
const verse_start_1 = require("./usx/verse_start");
const verse_end_1 = require("./usx/verse_end");
const position_1 = require("./usx/position");
class UsxNode {
    factory;
    item;
    position;
    constructor(factory, item, position) {
        this.factory = factory;
        this.item = item;
        this.position = position;
    }
}
var BookSections;
(function (BookSections) {
    BookSections[BookSections["identification"] = 0] = "identification";
    BookSections[BookSections["headers"] = 1] = "headers";
    BookSections[BookSections["titles"] = 2] = "titles";
    BookSections[BookSections["introduction"] = 3] = "introduction";
    BookSections[BookSections["endTitles"] = 4] = "endTitles";
    BookSections[BookSections["label"] = 5] = "label";
})(BookSections || (BookSections = {}));
class BookProcessorContext extends item_1.UsxContext {
    book;
    chapters = [];
    items = [];
    sections = [];
    verseItems = [];
    constructor(book) {
        super();
        this.book = book;
    }
    supports(factory, tag) {
        if (factory === book_identification_1.BookIdentificationFactory.instance) {
            if (this.sections.includes(BookSections.identification)) {
                return false;
            }
        }
        else if (factory === book_header_1.BookHeaderFactory.instance) {
            if (!this.sections.includes(BookSections.identification)) {
                return false;
            }
        }
        else if (factory === book_title_1.BookTitleFactory.instance) {
            if (!this.sections.includes(BookSections.headers)) {
                return false;
            }
        }
        else if (factory === book_introduction_1.BookIntroductionFactory.instance) {
            if (!this.sections.includes(BookSections.titles)) {
                return false;
            }
        }
        else if (factory === book_introduction_end_titles_1.BookIntroductionEndTitleFactory.instance) {
            if (!this.sections.includes(BookSections.introduction)) {
                return false;
            }
        }
        else if (factory === book_chapter_label_1.BookChapterLabelFactory.instance) {
            if (!this.sections.includes(BookSections.introduction)) {
                return false;
            }
        }
        return super.supports(factory, tag);
    }
    addText(text, position) {
        this.positions.push(new position_1.Position(position));
        if (this.items.length === 0)
            return;
        const node = this.push({ name: '#text', attributes: {}, isSelfClosing: true }, position);
        if (node instanceof text_1.Text) {
            node.text = text;
        }
        else {
            throw new Error('unsupported text node');
        }
        this.pop(position + text.length);
    }
    onFactoryPush(factory) {
        if (factory === book_identification_1.BookIdentificationFactory.instance) {
            this.sections.push(BookSections.identification);
        }
        else if (factory === book_header_1.BookHeaderFactory.instance) {
            this.sections.push(BookSections.headers);
        }
        else if (factory === book_title_1.BookTitleFactory.instance) {
            this.sections.push(BookSections.titles);
        }
        else if (factory === book_introduction_1.BookIntroductionFactory.instance) {
            this.sections.push(BookSections.introduction);
        }
        else if (factory === book_introduction_end_titles_1.BookIntroductionEndTitleFactory.instance) {
            this.sections.push(BookSections.endTitles);
        }
        else if (factory === book_chapter_label_1.BookChapterLabelFactory.instance) {
            this.sections.push(BookSections.label);
        }
    }
    push(tag, position) {
        this.positions.push(new position_1.Position(position));
        if (tag.name.toLowerCase() === 'usx') {
            const factory = usx_1.UsxFactory.instance;
            const item = factory.create(this, tag.attributes);
            this.items.push(new UsxNode(factory, item, position));
            return item;
        }
        if (this.items.length === 0) {
            throw new Error('empty stack, invalid state');
        }
        const node = this.items[this.items.length - 1];
        const factory = node.factory.findChildFactory(this, tag);
        const item = factory.create(this, tag.attributes);
        if (item instanceof chapter_start_1.ChapterStart) {
            this.positions.push(new position_1.Position(position));
            const chapter = new chapter_1.Chapter(this, this.book, item);
            this.positions.pop();
            this.chapters.push(chapter);
            this.items.push(new UsxNode(node.factory, chapter, position));
        }
        else if (item instanceof verse_start_1.VerseStart) {
            this.pushVerse(this.chapters[this.chapters.length - 1].usfm, item.number, new position_1.Position(item.position.start));
        }
        if (node.item instanceof item_1.UsxItemContainer) {
            node.item.addItem(item);
            this.onFactoryPush(factory);
        }
        this.items.push(new UsxNode(factory, item, position));
        return item;
    }
    pop(position) {
        const lastPosition = this.positions.pop();
        lastPosition.end = position;
        const node = this.items.pop();
        if (node) {
            if (node.item instanceof verse_end_1.VerseEnd) {
                const verse = this.popVerse();
                verse.position.end = position;
                this.verseItems.push(verse);
            }
            else if (node.item instanceof chapter_end_1.ChapterEnd) {
                const chapterNode = this.items.pop();
                if (!chapterNode)
                    throw new Error('missing chapter node');
                const chapter = chapterNode.item;
                chapter.end = node.item;
                chapter.position.start = chapter.start.position.start;
                chapter.position.end = node.item.position.end;
                chapter.addVerseItems(this.verseItems);
                this.verseItems = [];
                this.book.chapters.push(chapter);
            }
        }
    }
}
class BookProcessor {
    async process(name, content, data) {
        const book = new book_1.Book(name, content, data);
        const context = new BookProcessorContext(book);
        await new Promise((resolve, reject) => {
            const saxParser = (0, sax_1.parser)();
            saxParser.onerror = reject;
            saxParser.onopentag = (tag) => {
                context.push(tag, saxParser.startTagPosition - 1);
            };
            saxParser.ontext = (text) => {
                context.addText(text, saxParser.startTagPosition);
            };
            saxParser.onclosetag = (_) => {
                context.pop(saxParser.position);
            };
            saxParser.onend = () => {
                resolve(null);
            };
            saxParser.write(data).close();
        });
        return book;
    }
}
exports.BookProcessor = BookProcessor;
//# sourceMappingURL=book_processor.js.map