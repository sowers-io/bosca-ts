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
exports.ChapterEndFactory = exports.ChapterEnd = void 0;
const item_1 = require("./item");
class ChapterEnd {
    eid;
    position;
    verse;
    constructor(context, attributes) {
        this.eid = attributes.EID.toString();
        this.verse = context.addVerseItem(this);
        this.position = context.position;
    }
}
exports.ChapterEnd = ChapterEnd;
class ChapterEndFactory extends item_1.UsxItemFactory {
    static instance = new ChapterEndFactory();
    constructor() {
        super('chapter', new item_1.EndIdFactoryFilter());
    }
    onInitialize() {
    }
    create(context, attributes) {
        return new ChapterEnd(context, attributes);
    }
}
exports.ChapterEndFactory = ChapterEndFactory;
//# sourceMappingURL=chapter_end.js.map