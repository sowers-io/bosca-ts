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
exports.FigureFactory = exports.Figure = void 0;
const item_1 = require("./item");
class Figure extends item_1.UsxItemContainer {
    style;
    alt;
    file;
    size;
    loc;
    copy;
    ref;
    constructor(context, attributes) {
        super(context, attributes);
        this.style = attributes.STYLE.toString();
        this.alt = attributes.ALT?.toString();
        this.file = attributes.FILE.toString();
        this.size = attributes.SIZE?.toString();
        this.loc = attributes.LOC?.toString();
        this.copy = attributes.COPY?.toString();
        this.ref = attributes.REF?.toString();
    }
}
exports.Figure = Figure;
class FigureFactory extends item_1.UsxItemFactory {
    static instance = new FigureFactory();
    constructor() {
        super('figure');
    }
    onInitialize() {
    }
    create(context, attributes) {
        return new Figure(context, attributes);
    }
}
exports.FigureFactory = FigureFactory;
//# sourceMappingURL=figure.js.map