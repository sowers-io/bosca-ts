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
exports.TableContent = exports.Row = exports.Table = void 0;
const item_1 = require("./item");
class Table extends item_1.UsxItemContainer {
    vid;
    get element() {
        return 'table';
    }
    constructor(context, attributes) {
        super(context, attributes);
        this.vid = attributes.VID.toString();
    }
}
exports.Table = Table;
class Row extends item_1.UsxItemContainer {
    style;
    get element() {
        return 'row';
    }
    constructor(context, attributes) {
        super(context, attributes);
        this.style = attributes.STYLE.toString();
    }
}
exports.Row = Row;
class TableContent extends item_1.UsxItemContainer {
    style;
    align;
    colspan;
    get element() {
        return 'cell';
    }
    constructor(context, attributes) {
        super(context, attributes);
        this.style = attributes.STYLE.toString();
        this.align = attributes.ALIGN.toString();
        this.colspan = attributes.COLSPAN?.toString();
    }
}
exports.TableContent = TableContent;
//# sourceMappingURL=table.js.map