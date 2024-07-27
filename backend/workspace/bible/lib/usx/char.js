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
exports.CharFactory = exports.Char = void 0;
const styles_1 = require("./styles");
const item_1 = require("./item");
const footnote_1 = require("./footnote");
const break_1 = require("./break");
const text_1 = require("./text");
const reference_1 = require("./reference");
const milestone_1 = require("./milestone");
class Char extends item_1.UsxItemContainer {
    style;
    //char.link?,
    //char.closed?,
    constructor(context, attributes) {
        super(context, attributes);
        this.style = attributes.STYLE.toString();
    }
}
exports.Char = Char;
class CharFactory extends item_1.UsxItemFactory {
    static instance = new CharFactory();
    constructor() {
        super('char', new item_1.StyleFactoryFilter(styles_1.CharStyles));
    }
    onInitialize() {
        this.register(reference_1.ReferenceFactory.instance);
        this.register(CharFactory.instance);
        this.register(milestone_1.MilestoneFactory.instance);
        this.register(footnote_1.FootnoteFactory.instance);
        this.register(break_1.BreakFactory.instance);
        this.register(text_1.TextFactory.instance);
    }
    create(context, attributes) {
        return new Char(context, attributes);
    }
}
exports.CharFactory = CharFactory;
//# sourceMappingURL=char.js.map