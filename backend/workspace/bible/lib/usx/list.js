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
exports.ListFactory = exports.List = void 0;
const styles_1 = require("./styles");
const item_1 = require("./item");
const reference_1 = require("./reference");
const footnote_1 = require("./footnote");
const char_1 = require("./char");
const milestone_1 = require("./milestone");
const figure_1 = require("./figure");
const break_1 = require("./break");
const text_1 = require("./text");
const cross_reference_1 = require("./cross_reference");
const list_char_1 = require("./list_char");
const verse_start_1 = require("./verse_start");
const verse_end_1 = require("./verse_end");
class List extends item_1.UsxItemContainer {
    style;
    vid;
    constructor(context, attributes) {
        super(context, attributes);
        this.style = attributes.STYLE.toString();
        this.vid = attributes.VID.toString();
    }
}
exports.List = List;
class ListFactory extends item_1.UsxItemFactory {
    static instance = new ListFactory();
    constructor() {
        super('para', new item_1.StyleFactoryFilter(styles_1.ListStyles));
    }
    onInitialize() {
        this.register(reference_1.ReferenceFactory.instance);
        this.register(footnote_1.FootnoteFactory.instance);
        this.register(cross_reference_1.CrossReferenceFactory.instance);
        this.register(char_1.CharFactory.instance);
        this.register(list_char_1.ListCharFactory.instance);
        this.register(milestone_1.MilestoneFactory.instance);
        this.register(figure_1.FigureFactory.instance);
        this.register(verse_start_1.VerseStartFactory.instance);
        this.register(verse_end_1.VerseEndFactory.instance);
        this.register(break_1.BreakFactory.instance);
        this.register(text_1.TextFactory.instance);
    }
    create(context, attributes) {
        return new List(context, attributes);
    }
}
exports.ListFactory = ListFactory;
//# sourceMappingURL=list.js.map