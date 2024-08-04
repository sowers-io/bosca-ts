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

import { UsxContext, UsxItemContainer, UsxVerseItems } from './item'
import { Paragraph } from './paragraph'
import { List } from './list'
import { Table } from './table'
import { Footnote } from './footnote'
import { CrossReference } from './cross_reference'
import { Sidebar } from './sidebar'
import { ChapterStart } from './chapter_start'
import { ChapterEnd } from './chapter_end'
import { Book } from './book'

export type ChapterType = Paragraph | List | Table | Footnote | CrossReference | Sidebar | ChapterEnd

export class ChapterVerse {
  readonly usfm: string
  readonly chapter: string
  readonly verse: string
  readonly items: UsxVerseItems[]
  readonly raw: string

  constructor(usfm: string, chapter: string, verse: string, items: UsxVerseItems[], raw: string) {
    this.usfm = usfm
    this.chapter = chapter
    this.verse = verse
    this.items = items
    this.raw = raw
  }

  toString(): string {
    let buf = ''
    for (const item of this.items) {
      buf += item.toString()
    }
    return buf
  }
}

export class Chapter extends UsxItemContainer<ChapterType> {

  readonly number: string
  readonly usfm: string
  readonly verseItems: { [verse: string]: UsxVerseItems[] } = {}
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
      let current = this.verseItems[item.usfm]
      if (!current) {
        current = []
        this.verseItems[item.usfm] = current
      }
      current.push(item)
    }
  }

  addItem(item: ChapterType) {
    if (item instanceof ChapterEnd) return
    super.addItem(item)
  }

  getVerses(book: Book): ChapterVerse[] {
    const verses: ChapterVerse[] = []
    for (const usfm in this.verseItems) {
      const items = this.verseItems[usfm]
      const usfmParts = usfm.split('.')
      let raw = ''
      for (const item of items) {
        raw += book.getRawContent(item.position)
      }
      verses.push(new ChapterVerse(
        usfm,
        this.number,
        usfmParts[usfmParts.length - 1],
        items,
        raw,
      ))
    }

    function getNumber(value: string): number {
      let number = parseInt(value)
      if (isNaN(number)) number = 0
      return number
    }

    verses.sort((a, b) => {
      const aChapter = getNumber(a.chapter)
      const bChapter = getNumber(b.chapter)
      if (aChapter > bChapter) return 1
      if (aChapter < bChapter) return -1

      const aVerse = getNumber(a.verse)
      const bVerse = getNumber(b.verse)
      if (aVerse > bVerse) return 1
      if (aVerse < bVerse) return -1
      return 0
    })

    return verses
  }
}