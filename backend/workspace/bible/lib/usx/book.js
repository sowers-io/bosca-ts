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
exports.Book = void 0;
class Book {
    name;
    content;
    chapters = [];
    raw;
    constructor(name, content, raw) {
        this.name = name;
        this.content = content;
        this.raw = raw;
    }
    get usfm() {
        return this.content.usfm;
    }
    getRawContent(position) {
        return this.raw.substring(position.start, position.end);
    }
}
exports.Book = Book;
//# sourceMappingURL=book.js.map