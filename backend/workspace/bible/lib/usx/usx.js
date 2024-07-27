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
exports.UsxFactory = exports.Usx = void 0;
const item_1 = require("./item");
const book_identification_1 = require("./book_identification");
const book_header_1 = require("./book_header");
const book_title_1 = require("./book_title");
const book_introduction_1 = require("./book_introduction");
const book_introduction_end_titles_1 = require("./book_introduction_end_titles");
const book_chapter_label_1 = require("./book_chapter_label");
const chapter_start_1 = require("./chapter_start");
const chapter_end_1 = require("./chapter_end");
const paragraph_1 = require("./paragraph");
const list_1 = require("./list");
const footnote_1 = require("./footnote");
const cross_reference_1 = require("./cross_reference");
const text_1 = require("./text");
class Usx extends item_1.UsxItemContainer {
    constructor(context, attributes) {
        super(context, attributes);
    }
}
exports.Usx = Usx;
class UsxFactory extends item_1.UsxItemFactory {
    static _instance;
    static get instance() {
        if (this._instance == null) {
            this._instance = new UsxFactory();
            this._instance.initialize();
        }
        return this._instance;
    }
    constructor() {
        super('usx');
    }
    onInitialize() {
        this.register(book_identification_1.BookIdentificationFactory.instance);
        this.register(book_header_1.BookHeaderFactory.instance);
        this.register(book_title_1.BookTitleFactory.instance);
        this.register(book_introduction_1.BookIntroductionFactory.instance);
        this.register(book_introduction_end_titles_1.BookIntroductionEndTitleFactory.instance);
        this.register(book_chapter_label_1.BookChapterLabelFactory.instance);
        this.register(chapter_start_1.ChapterStartFactory.instance);
        this.register(chapter_end_1.ChapterEndFactory.instance);
        this.register(paragraph_1.ParagraphFactory.instance);
        this.register(list_1.ListFactory.instance);
        // this.register(Table)
        this.register(footnote_1.FootnoteFactory.instance);
        this.register(cross_reference_1.CrossReferenceFactory.instance);
        // this.register(Sidebar)
        this.register(text_1.TextFactory.instance);
    }
    create(context, attributes) {
        return new Usx(context, attributes);
    }
}
exports.UsxFactory = UsxFactory;
//# sourceMappingURL=usx.js.map