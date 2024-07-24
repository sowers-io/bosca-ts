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
exports.Chapter = exports.ChapterVerse = void 0;
const item_1 = require("./item");
const chapter_end_1 = require("./chapter_end");
class ChapterVerse {
    usfm;
    chapter;
    verse;
    items;
    raw;
    constructor(usfm, chapter, verse, items, raw) {
        this.usfm = usfm;
        this.chapter = chapter;
        this.verse = verse;
        this.items = items;
        this.raw = raw;
    }
}
exports.ChapterVerse = ChapterVerse;
class Chapter extends item_1.UsxItemContainer {
    number;
    usfm;
    verseItems = {};
    start;
    _end;
    constructor(context, book, start) {
        super(context, {});
        this.number = start.number;
        this.usfm = book.usfm + '.' + start.number;
        this.start = start;
    }
    get end() {
        return this._end;
    }
    set end(end) {
        if (this._end)
            throw new Error('end already defined');
        this._end = end;
    }
    addVerseItems(items) {
        for (const item of items) {
            let current = this.verseItems[item.usfm];
            if (!current) {
                current = [];
                this.verseItems[item.usfm] = current;
            }
            current.push(item);
        }
    }
    addItem(item) {
        if (item instanceof chapter_end_1.ChapterEnd)
            return;
        super.addItem(item);
    }
    getVerses(book) {
        const verses = [];
        for (const usfm in this.verseItems) {
            const items = this.verseItems[usfm];
            const usfmParts = usfm.split('.');
            let raw = '';
            for (const item of items) {
                raw += book.getRawContent(item.position);
            }
            verses.push(new ChapterVerse(usfm, this.number, usfmParts[usfmParts.length - 1], items, raw));
        }
        function getNumber(value) {
            let number = parseInt(value);
            if (isNaN(number))
                number = 0;
            return number;
        }
        verses.sort((a, b) => {
            const aChapter = getNumber(a.chapter);
            const bChapter = getNumber(b.chapter);
            if (aChapter > bChapter)
                return 1;
            if (aChapter < bChapter)
                return -1;
            const aVerse = getNumber(a.verse);
            const bVerse = getNumber(b.verse);
            if (aVerse > bVerse)
                return 1;
            if (aVerse < bVerse)
                return -1;
            return 0;
        });
        return verses;
    }
}
exports.Chapter = Chapter;
//# sourceMappingURL=chapter.js.map