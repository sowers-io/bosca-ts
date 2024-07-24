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
exports.BookIdentificationFactory = exports.BookIdentification = void 0;
const identification_1 = require("../identification");
const item_1 = require("./item");
const text_1 = require("./text");
class BookIdentification extends item_1.UsxItemContainer {
    id;
    code;
    constructor(context, attributes) {
        super(context, attributes);
        this.id = attributes.STYLE.toString();
        this.code = attributes.CODE.toString();
    }
}
exports.BookIdentification = BookIdentification;
class BookIdentificationFactory extends item_1.UsxItemFactory {
    static instance = new BookIdentificationFactory();
    constructor() {
        super('book', new item_1.CodeFactoryFilter(identification_1.BookIdentificationCodes));
    }
    onInitialize() {
        this.register(text_1.TextFactory.instance);
    }
    create(context, attributes) {
        return new BookIdentification(context, attributes);
    }
}
exports.BookIdentificationFactory = BookIdentificationFactory;
//# sourceMappingURL=book_identification.js.map