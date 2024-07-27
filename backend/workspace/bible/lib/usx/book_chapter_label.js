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
exports.BookChapterLabelFactory = exports.BookChapterLabel = void 0;
const item_1 = require("./item");
const text_1 = require("./text");
const styles_1 = require("./styles");
class BookChapterLabel extends item_1.UsxItemContainer {
    style;
    constructor(context, attributes) {
        super(context, attributes);
        this.style = attributes.STYLE.toString();
    }
}
exports.BookChapterLabel = BookChapterLabel;
class BookChapterLabelFactory extends item_1.UsxItemFactory {
    static instance = new BookChapterLabelFactory();
    constructor() {
        super('para', new item_1.StyleFactoryFilter(styles_1.BookChapterLabelStyles));
        this.register(text_1.TextFactory.instance);
    }
    onInitialize() {
    }
    create(context, attributes) {
        return new BookChapterLabel(context, attributes);
    }
}
exports.BookChapterLabelFactory = BookChapterLabelFactory;
//# sourceMappingURL=book_chapter_label.js.map