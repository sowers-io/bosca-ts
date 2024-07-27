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
exports.IntroCharFactory = exports.IntroChar = void 0;
const item_1 = require("./item");
const styles_1 = require("./styles");
class IntroChar extends item_1.UsxItemContainer {
    style;
    // char.closed?
    constructor(context, attributes) {
        super(context, attributes);
        this.style = attributes.STYLE.toString();
    }
}
exports.IntroChar = IntroChar;
class IntroCharFactory extends item_1.UsxItemFactory {
    static instance = new IntroCharFactory();
    constructor() {
        super('char', new item_1.StyleFactoryFilter(styles_1.IntroCharStyles));
    }
    onInitialize() {
    }
    create(context, attributes) {
        return new IntroChar(context, attributes);
    }
}
exports.IntroCharFactory = IntroCharFactory;
//# sourceMappingURL=intro_char.js.map