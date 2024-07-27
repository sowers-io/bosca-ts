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
exports.MilestoneFactory = exports.Milestone = void 0;
const item_1 = require("./item");
class Milestone {
    style;
    sid;
    eid;
    position;
    verse;
    constructor(context, attributes) {
        this.position = context.position;
        this.style = attributes.STYLE.toString();
        this.sid = attributes.SID.toString();
        this.eid = attributes.EID.toString();
        this.verse = context.addVerseItem(this);
    }
}
exports.Milestone = Milestone;
class MilestoneFactory extends item_1.UsxItemFactory {
    static instance = new MilestoneFactory();
    constructor() {
        super('ms');
    }
    onInitialize() {
    }
    create(context, attributes) {
        return new Milestone(context, attributes);
    }
}
exports.MilestoneFactory = MilestoneFactory;
//# sourceMappingURL=milestone.js.map