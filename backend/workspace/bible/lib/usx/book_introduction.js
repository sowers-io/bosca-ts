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
exports.BookIntroductionFactory = exports.BookIntroduction = void 0;
const styles_1 = require("./styles");
const item_1 = require("./item");
const text_1 = require("./text");
const milestone_1 = require("./milestone");
const figure_1 = require("./figure");
const reference_1 = require("./reference");
const footnote_1 = require("./footnote");
const cross_reference_1 = require("./cross_reference");
const char_1 = require("./char");
const intro_char_1 = require("./intro_char");
class BookIntroduction extends item_1.UsxItemContainer {
    style;
    constructor(context, attributes) {
        super(context, attributes);
        this.style = attributes.STYLE.toString();
    }
}
exports.BookIntroduction = BookIntroduction;
class BookIntroductionFactory extends item_1.UsxItemFactory {
    static instance = new BookIntroductionFactory();
    constructor() {
        super('para', new item_1.StyleFactoryFilter(styles_1.BookIntroductionStyles));
    }
    onInitialize() {
        this.register(reference_1.ReferenceFactory.instance);
        this.register(footnote_1.FootnoteFactory.instance);
        this.register(cross_reference_1.CrossReferenceFactory.instance);
        this.register(char_1.CharFactory.instance);
        this.register(intro_char_1.IntroCharFactory.instance);
        this.register(milestone_1.MilestoneFactory.instance);
        this.register(figure_1.FigureFactory.instance);
        this.register(text_1.TextFactory.instance);
    }
    create(context, attributes) {
        return new BookIntroduction(context, attributes);
    }
}
exports.BookIntroductionFactory = BookIntroductionFactory;
//# sourceMappingURL=book_introduction.js.map