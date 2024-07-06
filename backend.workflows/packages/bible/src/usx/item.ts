import { Tag, QualifiedTag, QualifiedAttribute } from 'sax'
import { Position } from './position'

export type Attributes = { [key: string]: string | QualifiedAttribute }

export type UsxTag = Tag | QualifiedTag

export interface UsxItem {

  readonly position: Position
  readonly verse: string | null

  toString(): string
}

export abstract class UsxItemContainer<T extends UsxItem> implements UsxItem {
  items: T[] = []

  readonly position: Position
  readonly verse: string | null

  protected constructor(context: UsxContext, attributes: Attributes) {
    this.position = context.position
    this.verse = context.addVerseItem(this)
  }

  addItem(item: T) {
    this.items.push(item)
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
  supports(context: UsxContext, attributes: Attributes): boolean
}

export class StyleFactoryFilter<T> implements ItemFactoryFilter {
  private styles: T[]

  constructor(styles: T[]) {
    this.styles = styles
  }

  supports(context: UsxContext, attributes: Attributes): boolean {
    if (!attributes.STYLE) return false
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

  supports(context: UsxContext, attributes: Attributes): boolean {
    return !this.filter.supports(context, attributes)
  }
}

export class UsxVerseItems {
  readonly usfm: string
  readonly verse: string
  readonly items: UsxItem[] = []

  constructor(usfm: string, verse: string) {
    this.usfm = usfm
    this.verse = verse
  }

  addItem(item: UsxItem) {
    this.items.push(item)
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

  pushVerse(bookChapterUsfm: string, verse: string) {
    this.verses.push(new UsxVerseItems(bookChapterUsfm + '.' + verse, verse))
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

  supports(factory: UsxItemFactory<any>, tag: UsxTag): boolean {
    return factory.supports(this, tag.attributes)
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

  supports(context: UsxContext, attributes: Attributes): boolean {
    return this.filter?.supports(context, attributes) ?? true
  }

  abstract create(context: UsxContext, attributes: Attributes): T

  findChildFactory(context: UsxContext, tag: UsxTag): UsxItemFactory<any> {
    const factories = this.factories[tag.name.toLowerCase()]
    if (!factories) {
      throw new Error('unsupported tag: ' + tag.name + ' in ' + this.tagName)
    }
    const supported = factories.filter((factory) => context.supports(factory, tag))
    if (supported.length === 0) {
      throw new Error('zero supported items')
    } else if (supported.length > 1) {
      throw new Error('multiple supported items')
    }
    return supported[0]
  }
}