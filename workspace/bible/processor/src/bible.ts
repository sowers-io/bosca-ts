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

import { Book } from './usx/book'
import { UsxVerseItems } from './usx/item'

export class Bible {

  readonly books: Book[]
  readonly booksByUsfm: { [usfm: string]: Book }

  constructor(books: Book[], booksByUsfm: { [usfm: string]: Book }) {
    this.books = books
    this.booksByUsfm = booksByUsfm
  }

  getVerseItems(reference: BibleReference | BibleReference[]): UsxVerseItems[] {
    const allItems: UsxVerseItems[] = []
    const allReferences = Array.isArray(reference) ? reference.flatMap((r) => r.references) : reference.references
    for (const ref of allReferences) {
      const book = this.booksByUsfm[ref.bookUsfm]
      if (book) {
        const chapter = book.chaptersByUsfm[ref.chapterUsfm]
        if (chapter) {
          const items = chapter.verseItems[ref.usfm]
          if (items) {
            allItems.push(...items)
          }
        }
      }
    }
    return allItems
  }
}

export class BibleReference {
  readonly usfm: string
  private _references: BibleReference[] | undefined
  private _bookUsfm: string | undefined
  private _chapter: string | undefined
  private _chapterUsfm: string | undefined

  constructor(usfm: string) {
    this.usfm = usfm
  }

  get bookUsfm(): string {
    if (this._bookUsfm) return this._bookUsfm
    this._bookUsfm = this.usfm.split('.')[0]
    return this._bookUsfm
  }

  get chapterUsfm(): string {
    if (this._chapterUsfm) return this._chapterUsfm
    const parts = this.usfm.split('.')
    this._chapterUsfm = parts[0] + '.' + parts[1]
    return this._chapterUsfm
  }

  get chapter(): string {
    if (this._chapter) return this._chapter
    this._chapter = this.usfm.split('.')[1]
    return this._chapter
  }

  get references(): BibleReference[] {
    if (this._references) return this._references
    this._references = this.usfm.split('+').map((u) => new BibleReference(u))
    return this._references
  }

  join(other: BibleReference): BibleReference {
    return new BibleReference(this.usfm + '+' + other.usfm)
  }

  static parse(bible: Bible, human: string): BibleReference[] {
    const parts = human.split(',')
    const references: BibleReference[] = []
    for (const part of parts) {
      const reference = BibleReference.parseSingle(bible, part.trim())
      if (reference) {
        references.push(reference)
      }
    }
    const joinedReferences = new Map<string, string>()
    for (const reference of references) {
      let usfms = joinedReferences.get(reference.chapter)
      if (!usfms) {
        usfms = reference.usfm
      } else {
        usfms += '+' + reference.usfm
      }
      joinedReferences.set(reference.chapter, usfms)
    }
    const finalReferences: BibleReference[] = []
    for (const usfm of joinedReferences.values()) {
      finalReferences.push(new BibleReference(usfm))
    }
    return finalReferences
  }

  private static parseSingle(bible: Bible, human: string): BibleReference | null {
    const humanLower = human.toLocaleLowerCase()
    let book: Book | null = null
    let nonBook: string | null = null
    for (const b of bible.books) {
      if (humanLower.indexOf(b.name?.long?.toLocaleLowerCase() || '') === 0) {
        book = b
        nonBook = humanLower.substring(b.name.long.length).trim()
      } else if (humanLower.indexOf(b.name?.short?.toLocaleLowerCase() || '') === 0) {
        book = b
        nonBook = humanLower.substring(b.name.short.length).trim()
      } else if (humanLower.indexOf(b.name.abbreviation?.toLocaleLowerCase() || '') === 0) {
        book = b
        nonBook = humanLower.substring(b.name.abbreviation.length).trim()
      }
    }
    if (!book || !nonBook) {
      return null
    }
    if (nonBook.length === 0) {
      return new BibleReference(book.usfm)
    }
    const numberParts = nonBook.split(':')
    const chapter = book.chapters.find((c) => c.number.toLocaleLowerCase() === numberParts[0].toLocaleLowerCase())
    if (!chapter) {
      return new BibleReference(book.usfm)
    }
    if (numberParts.length === 1) {
      return new BibleReference(chapter.usfm)
    }
    if (numberParts[1].includes('–')) {
      numberParts[1] = numberParts[1].replace('–', '-')
    }
    if (numberParts[1].includes('-')) {
      const rangeParts = numberParts[1].split('-')
      if (rangeParts.length === 2) {
        const start = parseInt(rangeParts[0])
        const end = parseInt(rangeParts[1])
        const usfms: string[] = []
        for (let i = start; i <= end; i++) {
          usfms.push(chapter.usfm + '.' + i)
        }
        return new BibleReference(usfms.join('+'))
      } else {
        numberParts[1] = rangeParts[0]
      }
    }
    return new BibleReference(chapter.usfm + '.' + numberParts[1])
  }
}
