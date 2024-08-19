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

import { ManifestName, PublicationContent } from './metadata'
import { parser } from 'sax'
import { UsxContext, UsxItem, UsxItemContainer, UsxItemFactory, UsxTag, UsxVerseItems } from './usx/item'
import { Book } from './usx/book'
import { Usx, UsxFactory } from './usx/usx'
import { BookIdentificationFactory } from './usx/book_identification'
import { BookTitleFactory } from './usx/book_title'
import { BookHeaderFactory } from './usx/book_header'
import { Text } from './usx/text'
import { ChapterEnd } from './usx/chapter_end'
import { Chapter } from './usx/chapter'
import { ChapterStart } from './usx/chapter_start'
import { VerseStart } from './usx/verse_start'
import { VerseEnd } from './usx/verse_end'
import { Position } from './usx/position'
import { BookIntroductionFactory } from './usx/book_introduction'
import { BookChapterLabelFactory } from './usx/book_chapter_label'
import { BookIntroductionEndTitleFactory } from './usx/book_introduction_end_titles'
import {
  BookChapterLabelStyles,
  BookHeaderStyles,
  BookIntroductionEndTitleStyles,
  BookIntroductionStyles,
  BookTitleStyles,
} from './usx/styles'

export class UsxNode {
  readonly factory: UsxItemFactory<any>
  readonly item: UsxItem
  readonly position: number

  constructor(factory: UsxItemFactory<any>, item: UsxItem, position: number) {
    this.factory = factory
    this.item = item
    this.position = position
  }
}

enum CompletedBookTag {
  identification,
  headers,
  titles,
  introduction,
  endIntroductionTitles,
  label,
  chapter,
}

enum BookTagResult {
  supported,
  unsupported,
  unknown,
}

class BookProcessorContext extends UsxContext {

  private readonly book: Book
  private readonly chapters: Chapter[] = []
  private readonly items: UsxNode[] = []
  private readonly completed: CompletedBookTag[] = []
  private lastFactory: UsxItemFactory<any> = BookIdentificationFactory.instance
  private progression = true
  private verseItems: UsxVerseItems[] = []

  constructor(book: Book) {
    super()
    this.book = book
  }

  private readonly tags = [
    { factory: BookIdentificationFactory.instance, tag: CompletedBookTag.identification, maxStyles: 1 },
    { factory: BookHeaderFactory.instance, tag: CompletedBookTag.headers, maxStyles: BookHeaderStyles.length },
    { factory: BookTitleFactory.instance, tag: CompletedBookTag.titles, maxStyles: BookTitleStyles.length },
    {
      factory: BookIntroductionFactory.instance,
      tag: CompletedBookTag.introduction,
      maxStyles: BookIntroductionStyles.length,
    },
    {
      factory: BookIntroductionEndTitleFactory.instance,
      tag: CompletedBookTag.endIntroductionTitles,
      maxStyles: BookIntroductionEndTitleStyles.length,
    },
    { factory: BookChapterLabelFactory.instance, tag: CompletedBookTag.label, maxStyles: BookChapterLabelStyles.length },
  ]

  private supportsInternal(factory: UsxItemFactory<any>, parent: UsxNode, tag: UsxTag): BookTagResult {
    for (let tagIndex = 0; tagIndex < this.tags.length; tagIndex++) {
      const tagFactory = this.tags[tagIndex]
      if (this.completed.includes(tagFactory.tag)) {
        if (tagFactory.factory === factory) {
          return BookTagResult.unsupported
        }
        continue
      }
      if (tagFactory.factory === factory) {
        if (tagIndex > 0) {
          const tag = this.tags[tagIndex - 1]
          if (!this.completed.includes(tag.tag)) {
            return BookTagResult.unsupported
          }
        }
        for (let i = 0; i < tagFactory.maxStyles; i++) {
          if (super.supports(tagFactory.factory, parent, tag, i)) {
            if (i + 1 === tagFactory.maxStyles) {
              this.completed.push(tagFactory.tag)
            }
            return BookTagResult.supported
          }
        }
        if (factory === this.lastFactory) {
          this.completed.push(tagFactory.tag)
        }
        return BookTagResult.unsupported
      }
    }
    return BookTagResult.unknown
  }

  supports(factory: UsxItemFactory<any>, parent: UsxNode, tag: UsxTag): boolean {
    if (tag.name.toLowerCase() === 'chapter') {
      this.progression = false
    }
    if (!this.progression || tag.name === '#text' || parent.factory != UsxFactory.instance) {
      if (!this.progression) {
        for (const tag of this.tags) {
          if (tag.factory === factory) return false
        }
      }
      return super.supports(factory, parent, tag)
    }
    switch (this.supportsInternal(factory, parent, tag)) {
      case BookTagResult.supported: {
        this.lastFactory = factory
        return true
      }
      case BookTagResult.unsupported:
        return false
      case BookTagResult.unknown:
        return !this.progression && this.supports(factory, parent, tag)
    }
  }

  addText(text: string, position: number) {
    this.positions.push(new Position(position))
    if (this.items.length === 0) return
    const node = this.push({ name: '#text', attributes: {}, isSelfClosing: true }, position)
    if (node instanceof Text) {
      node.text = text
    } else {
      throw new Error('unsupported text node')
    }
    this.pop(position + text.length)
  }

  push(tag: UsxTag, position: number): UsxItem {
    this.positions.push(new Position(position))
    if (tag.name.toLowerCase() === 'usx') {
      const factory = UsxFactory.instance
      const item = factory.create(this, tag.attributes)
      this.items.push(new UsxNode(factory, item, position))
      return item
    }
    if (this.items.length === 0) {
      throw new Error('empty stack, invalid state')
    }
    const node = this.items[this.items.length - 1]
    const factory = node.factory.findChildFactory(this, node, tag)
    const item = factory.create(this, tag.attributes)
    if (item instanceof ChapterStart) {
      this.progression = false
      this.positions.push(new Position(position))
      const chapter = new Chapter(this, this.book, item)
      this.positions.pop()
      this.chapters.push(chapter)
      this.items.push(new UsxNode(node.factory, chapter, position))
    } else if (item instanceof VerseStart) {
      this.pushVerse(this.chapters[this.chapters.length - 1].usfm, item, new Position(item.position.start))
    }
    if (node.item instanceof UsxItemContainer) {
      node.item.addItem(item)
    }
    this.items.push(new UsxNode(factory, item, position))
    return item
  }

  pop(position: number) {
    const lastPosition = this.positions.pop()!
    lastPosition.end = position
    const node = this.items.pop()
    if (node) {
      if (node.item instanceof VerseEnd) {
        const verse = this.popVerse()
        verse.position.end = position
        this.verseItems.push(verse)
      } else if (node.item instanceof ChapterEnd) {
        const chapterNode = this.items.pop()
        if (!chapterNode) throw new Error('missing chapter node')
        const chapter = chapterNode.item as Chapter
        chapter.end = node.item
        chapter.position.start = chapter.start.position.start
        chapter.position.end = node.item.position.end
        chapter.addVerseItems(this.verseItems)
        this.verseItems = []
        this.book.chapters.push(chapter)
      }
    }
  }
}

export class BookProcessor {

  async process(name: ManifestName, content: PublicationContent, data: string): Promise<Book> {
    const book = new Book(name, content, data)
    const context = new BookProcessorContext(book)
    await new Promise((resolve, reject) => {
      const saxParser = parser()
      saxParser.onerror = reject
      saxParser.onopentag = (tag) => {
        context.push(tag, saxParser.startTagPosition - 1)
      }
      saxParser.ontext = (text) => {
        context.addText(text, saxParser.startTagPosition)
      }
      saxParser.onclosetag = (_) => {
        context.pop(saxParser.position)
      }
      saxParser.onend = () => {
        resolve(null)
      }
      saxParser.write(data).close()
    })
    return book
  }
}

