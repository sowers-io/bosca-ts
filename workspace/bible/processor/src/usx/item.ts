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

import { Tag, QualifiedTag, QualifiedAttribute } from 'sax'
import { Position } from './position'
import { VerseStart } from './verse_start'
import { UsxNode } from '../book_processor'

export type Attributes = { [key: string]: string | QualifiedAttribute }

export type UsxTag = Tag | QualifiedTag

export class HtmlContext {

  pretty: boolean
  indent: number
  includeVerseNumbers: boolean

  constructor(pretty: boolean, indent: number, includeVerseNumbers: boolean) {
    this.pretty = pretty
    this.indent = indent
    this.includeVerseNumbers = includeVerseNumbers
  }

  addIndent() {
    this.indent += 2
  }

  removeIndent() {
    this.indent -= 2
  }

  render(tag: string, item: UsxItem | UsxVerseItems, text: string | undefined = undefined): string {
    let html = ''
    if (this.pretty) html += ' '.repeat(this.indent)
    html += '<' + tag
    let attrs = {
      ...item.htmlAttributes,
    }
    if (item.htmlClass != '') attrs['class'] = item.htmlClass
    for (const key in attrs) {
      html += ' '
      html += key + '="' + attrs[key] + '"'
    }
    html += '>'
    if (this.pretty) html += '\n'
    this.addIndent()
    let childHtml = ''
    if (text) {
      if (this.pretty) html += ' '.repeat(this.indent)
      childHtml += text
    } else if (item instanceof UsxItemContainer) {
      for (const child of item.items) {
        childHtml += child.toHtml(this)
      }
    } else if (item instanceof UsxVerseItems) {
      for (const child of item.items) {
        childHtml += child.toHtml(this)
      }
    } else {
      childHtml += item.toHtml(this)
    }
    html += childHtml
    if (this.pretty && !html.endsWith('\n')) html += '\n'
    this.removeIndent()
    if (this.pretty) html += ' '.repeat(this.indent)
    html += '</' + tag + '>'
    if (this.pretty) html += '\n'
    return html
  }
}

export interface UsxItem {

  readonly position: Position
  readonly verse: string | null

  get htmlClass(): string

  get htmlAttributes(): { [key: string]: string }

  toHtml(context: HtmlContext): string

  toString(): string
}

export abstract class UsxItemContainer<T extends UsxItem> implements UsxItem {
  items: T[] = []

  readonly position: Position
  readonly verse: string | null

  protected constructor(context: UsxContext, _: Attributes) {
    this.position = context.position
    this.verse = context.addVerseItem(this)
  }

  abstract get htmlClass(): string

  get htmlAttributes(): { [key: string]: string } {
    return {}
  }

  addItem(item: T) {
    this.items.push(item)
  }

  toHtml(context: HtmlContext): string {
    return context.render('div', this)
  }

  toString(): string {
    let verseContent = ''
    for (const item of this.items) {
      verseContent += item.toString()
    }
    return verseContent
  }
}

export interface ItemFactoryFilter {
  supports(context: UsxContext, attributes: Attributes, progression: number | null): boolean
}

export class StyleFactoryFilter<T> implements ItemFactoryFilter {
  private styles: T[]

  constructor(styles: T[]) {
    this.styles = styles
  }

  supports(context: UsxContext, attributes: Attributes, progression: number | null): boolean {
    if (!attributes.STYLE) return false
    if (progression != null) {
      return this.styles[progression] === attributes.STYLE.toString()
    }
    return this.styles.includes(attributes.STYLE.toString() as T)
  }
}

export class CodeFactoryFilter<T> implements ItemFactoryFilter {
  private styles: T[]

  constructor(styles: T[]) {
    this.styles = styles
  }

  supports(context: UsxContext, attributes: Attributes): boolean {
    if (!attributes.CODE) return false
    return this.styles.includes(attributes.CODE.toString() as T)
  }
}

export class EndIdFactoryFilter implements ItemFactoryFilter {

  supports(context: UsxContext, attributes: Attributes): boolean {
    if (!attributes.EID) return false
    return true
  }
}

export class NegateFactoryFilter implements ItemFactoryFilter {
  private filter: ItemFactoryFilter

  constructor(filter: ItemFactoryFilter) {
    this.filter = filter
  }

  supports(context: UsxContext, attributes: Attributes, progression: number | null): boolean {
    return !this.filter.supports(context, attributes, progression)
  }
}

export class UsxVerseItems {
  readonly usfm: string
  readonly verse: string
  readonly items: UsxItem[] = []
  readonly position: Position

  constructor(usfm: string, verse: VerseStart, position: Position) {
    this.usfm = usfm
    this.verse = verse.number
    this.position = position
    this.items.push(verse)
  }

  get htmlClass(): string {
    return 'verses'
  }

  get htmlAttributes(): { [key: string]: string } {
    return {}
  }

  addItem(item: UsxItem) {
    this.items.push(item)
  }

  toHtml(context: HtmlContext) {
    return context.render('div', this)
  }

  toString(): string {
    let verseContent = ''
    for (const item of this.items) {
      verseContent += item.toString()
    }
    return verseContent
  }
}

export abstract class UsxContext {

  protected positions: Position[] = []

  private verses: UsxVerseItems[] = []

  get position(): Position {
    return this.positions[this.positions.length - 1]
  }

  pushVerse(bookChapterUsfm: string, verse: VerseStart, position: Position) {
    this.verses.push(new UsxVerseItems(bookChapterUsfm + '.' + verse.number, verse, position))
  }

  popVerse(): UsxVerseItems {
    return this.verses.pop()!
  }

  get verse(): string | null {
    if (this.verses.length === 0) return null
    return this.verses[this.verses.length - 1].verse
  }

  addVerseItem(item: UsxItem): string | null {
    if (this.verses.length === 0) return null
    const verses = this.verses[this.verses.length - 1]
    verses.addItem(item)
    return verses.verse
  }

  supports(factory: UsxItemFactory<any>, parent: UsxNode, tag: UsxTag, progression: number | null = null): boolean {
    return factory.supports(this, tag.attributes, progression)
  }
}

export abstract class UsxItemFactory<T extends UsxItem> {

  readonly tagName: string
  private filter: ItemFactoryFilter | null
  private factories: { [tag: string]: UsxItemFactory<any>[] } = {}

  protected constructor(tagName: string, filter: ItemFactoryFilter | null = null) {
    this.tagName = tagName
    this.filter = filter
  }

  private initialized = false

  initialize() {
    if (this.initialized) return
    this.initialized = true
    this.onInitialize()
    for (const tag in this.factories) {
      for (const factory of this.factories[tag]) {
        factory.initialize()
      }
    }
  }

  protected abstract onInitialize(): void

  protected register(factory: UsxItemFactory<any>) {
    let factories = this.factories[factory.tagName]
    if (!factories) {
      factories = []
      this.factories[factory.tagName] = factories
    }
    factories.push(factory)
  }

  supports(context: UsxContext, attributes: Attributes, progression: number | null): boolean {
    return this.filter?.supports(context, attributes, progression) ?? true
  }

  abstract create(context: UsxContext, attributes: Attributes): T

  findChildFactory(context: UsxContext, parent: UsxNode, tag: UsxTag): UsxItemFactory<any> {
    const factories = this.factories[tag.name.toLowerCase()]
    if (!factories) {
      throw new Error('unsupported tag: ' + tag.name + ' in ' + this.tagName)
    }
    const supported = factories.filter((factory) => context.supports(factory, parent, tag))
    if (supported.length === 0) {
      throw new Error('zero supported items')
    } else if (supported.length > 1) {
      throw new Error('multiple supported items')
    }
    return supported[0]
  }
}