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
exports.CrossReferenceFactory = exports.CrossReference = void 0;
const cross_reference_char_1 = require("./cross_reference_char");
const text_1 = require("./text");
const item_1 = require("./item");
const styles_1 = require("./styles");
class CrossReference extends item_1.UsxItemContainer {
    style;
    caller;
    constructor(context, attributes) {
        super(context, attributes);
        this.style = attributes.STYLE.toString();
        this.caller = attributes.CALLER.toString();
    }
}
exports.CrossReference = CrossReference;
class CrossReferenceFactory extends item_1.UsxItemFactory {
    static instance = new CrossReferenceFactory();
    constructor() {
        super('note', new item_1.StyleFactoryFilter(styles_1.CrossReferenceStyles));
    }
    onInitialize() {
        this.register(cross_reference_char_1.CrossReferenceCharFactory.instance);
        this.register(text_1.TextFactory.instance);
    }
    create(context, attributes) {
        return new CrossReference(context, attributes);
    }
}
exports.CrossReferenceFactory = CrossReferenceFactory;
//# sourceMappingURL=cross_reference.js.map