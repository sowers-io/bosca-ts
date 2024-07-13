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

import {UsxContext, UsxItemContainer, UsxVerseItems} from './item'
import {Paragraph} from './paragraph'
import {List} from './list'
import {Table} from './table'
import {Footnote} from './footnote'
import {CrossReference} from './cross_reference'
import {Sidebar} from './sidebar'
import {ChapterStart} from './chapter_start'
import {ChapterEnd} from './chapter_end'
import {Book} from './book'

export type ChapterType = Paragraph | List | Table | Footnote | CrossReference | Sidebar | ChapterEnd

export class Chapter extends UsxItemContainer<ChapterType> {

    readonly number: string
    readonly usfm: string
    readonly verseItems: { [verse: string]: UsxVerseItems } = {}
    readonly start: ChapterStart
    private _end?: ChapterEnd

    constructor(context: UsxContext, book: Book, start: ChapterStart) {
        super(context, {})
        this.number = start.number
        this.usfm = book.usfm + '.' + start.number
        this.start = start
    }

    get end(): ChapterEnd {
        return this._end!
    }

    set end(end: ChapterEnd) {
        if (this._end) throw new Error('end already defined')
        this._end = end
    }

    addVerseItems(items: UsxVerseItems[]) {
        for (const item of items) {
            this.verseItems[item.usfm] = item
        }
    }

    addItem(item: ChapterType) {
        if (item instanceof ChapterEnd) return
        super.addItem(item)
    }
}