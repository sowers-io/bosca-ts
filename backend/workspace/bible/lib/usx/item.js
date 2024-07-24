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
exports.UsxItemFactory = exports.UsxContext = exports.UsxVerseItems = exports.NegateFactoryFilter = exports.EndIdFactoryFilter = exports.CodeFactoryFilter = exports.StyleFactoryFilter = exports.UsxItemContainer = void 0;
class UsxItemContainer {
    items = [];
    position;
    verse;
    constructor(context, attributes) {
        this.position = context.position;
        this.verse = context.addVerseItem(this);
    }
    addItem(item) {
        this.items.push(item);
    }
    toString() {
        let verseContent = '';
        for (const item of this.items) {
            verseContent += item.toString();
        }
        return verseContent;
    }
}
exports.UsxItemContainer = UsxItemContainer;
class StyleFactoryFilter {
    styles;
    constructor(styles) {
        this.styles = styles;
    }
    supports(context, attributes) {
        if (!attributes.STYLE)
            return false;
        return this.styles.includes(attributes.STYLE.toString());
    }
}
exports.StyleFactoryFilter = StyleFactoryFilter;
class CodeFactoryFilter {
    styles;
    constructor(styles) {
        this.styles = styles;
    }
    supports(context, attributes) {
        if (!attributes.CODE)
            return false;
        return this.styles.includes(attributes.CODE.toString());
    }
}
exports.CodeFactoryFilter = CodeFactoryFilter;
class EndIdFactoryFilter {
    supports(context, attributes) {
        if (!attributes.EID)
            return false;
        return true;
    }
}
exports.EndIdFactoryFilter = EndIdFactoryFilter;
class NegateFactoryFilter {
    filter;
    constructor(filter) {
        this.filter = filter;
    }
    supports(context, attributes) {
        return !this.filter.supports(context, attributes);
    }
}
exports.NegateFactoryFilter = NegateFactoryFilter;
class UsxVerseItems {
    usfm;
    verse;
    items = [];
    position;
    constructor(usfm, verse, position) {
        this.usfm = usfm;
        this.verse = verse;
        this.position = position;
    }
    addItem(item) {
        this.items.push(item);
    }
    toString() {
        let verseContent = '';
        for (const item of this.items) {
            verseContent += item.toString();
        }
        return verseContent;
    }
}
exports.UsxVerseItems = UsxVerseItems;
class UsxContext {
    positions = [];
    verses = [];
    get position() {
        return this.positions[this.positions.length - 1];
    }
    pushVerse(bookChapterUsfm, verse, position) {
        this.verses.push(new UsxVerseItems(bookChapterUsfm + '.' + verse, verse, position));
    }
    popVerse() {
        return this.verses.pop();
    }
    get verse() {
        if (this.verses.length === 0)
            return null;
        return this.verses[this.verses.length - 1].verse;
    }
    addVerseItem(item) {
        if (this.verses.length === 0)
            return null;
        const verses = this.verses[this.verses.length - 1];
        verses.addItem(item);
        return verses.verse;
    }
    supports(factory, tag) {
        return factory.supports(this, tag.attributes);
    }
}
exports.UsxContext = UsxContext;
class UsxItemFactory {
    tagName;
    filter;
    factories = {};
    constructor(tagName, filter = null) {
        this.tagName = tagName;
        this.filter = filter;
    }
    initialized = false;
    initialize() {
        if (this.initialized)
            return;
        this.initialized = true;
        this.onInitialize();
        for (const tag in this.factories) {
            for (const factory of this.factories[tag]) {
                factory.initialize();
            }
        }
    }
    register(factory) {
        let factories = this.factories[factory.tagName];
        if (!factories) {
            factories = [];
            this.factories[factory.tagName] = factories;
        }
        factories.push(factory);
    }
    supports(context, attributes) {
        return this.filter?.supports(context, attributes) ?? true;
    }
    findChildFactory(context, tag) {
        const factories = this.factories[tag.name.toLowerCase()];
        if (!factories) {
            throw new Error('unsupported tag: ' + tag.name + ' in ' + this.tagName);
        }
        const supported = factories.filter((factory) => context.supports(factory, tag));
        if (supported.length === 0) {
            throw new Error('zero supported items');
        }
        else if (supported.length > 1) {
            throw new Error('multiple supported items');
        }
        return supported[0];
    }
}
exports.UsxItemFactory = UsxItemFactory;
//# sourceMappingURL=item.js.map