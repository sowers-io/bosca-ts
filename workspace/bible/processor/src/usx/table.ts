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

import { Attributes, HtmlContext, UsxContext, UsxItem, UsxItemContainer, UsxItemFactory } from './item'
import { Footnote, FootnoteFactory } from './footnote'
import { Char, CharFactory } from './char'
import { Milestone, MilestoneFactory } from './milestone'
import { Figure, FigureFactory } from './figure'
import { Break, BreakFactory } from './break'
import { Text, TextFactory } from './text'
import { CrossReference, CrossReferenceFactory } from './cross_reference'
import { Verse } from './verse'
import { VerseStartFactory } from './verse_start'
import { VerseEndFactory } from './verse_end'

export class Table extends UsxItemContainer<Row> {

  vid: string

  constructor(context: UsxContext, parent: UsxItem | null, attributes: Attributes) {
    super(context, parent, attributes)
    this.vid = attributes.VID.toString()
  }

  get htmlClass(): string {
    return ''
  }

  toHtml(context: HtmlContext): string {
    return context.render('table', this)
  }
}

type RowType = Verse | TableContent

export class Row extends UsxItemContainer<RowType> {

  style: string

  constructor(context: UsxContext, parent: UsxItem | null, attributes: Attributes) {
    super(context, parent, attributes)
    this.style = attributes.STYLE.toString()
  }

  get htmlClass(): string {
    return this.style
  }

  toHtml(context: HtmlContext): string {
    return context.render('tr', this)
  }
}

type TableContentType = Footnote | CrossReference | Char | Milestone | Figure | Verse | Break | Text

export class TableContent extends UsxItemContainer<TableContentType> {
  style: string
  align: string
  colspan: string

  constructor(context: UsxContext, parent: UsxItem | null, attributes: Attributes) {
    super(context, parent, attributes)
    this.style = attributes.STYLE.toString()
    this.align = attributes.ALIGN.toString()
    this.colspan = attributes.COLSPAN?.toString()
  }

  get htmlClass(): string {
    return this.style
  }

  toHtml(context: HtmlContext): string {
    return context.render('td', this)
  }
}

export class TableFactory extends UsxItemFactory<Table> {

  static readonly instance = new TableFactory()

  private constructor() {
    super('table')
  }

  protected onInitialize() {
    this.register(RowFactory.instance)
    this.register(VerseStartFactory.instance)
    this.register(VerseEndFactory.instance)
  }

  create(context: UsxContext, parent: UsxItem | null, attributes: Attributes): Table {
    return new Table(context, parent, attributes)
  }
}

class RowFactory extends UsxItemFactory<Row> {

  static readonly instance = new RowFactory()

  private constructor() {
    super('row')
  }

  protected onInitialize() {
    this.register(ContentFactory.instance)
  }

  create(context: UsxContext, parent: UsxItem | null, attributes: Attributes): Row {
    return new Row(context, parent, attributes)
  }
}

class ContentFactory extends UsxItemFactory<TableContent> {

  static readonly instance = new ContentFactory()

  private constructor() {
    super('cell')
  }

  protected onInitialize() {
    this.register(FootnoteFactory.instance)
    this.register(CrossReferenceFactory.instance)
    this.register(CharFactory.instance)
    this.register(MilestoneFactory.instance)
    this.register(FigureFactory.instance)
    this.register(VerseStartFactory.instance)
    this.register(VerseEndFactory.instance)
    this.register(BreakFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, parent: UsxItem | null, attributes: Attributes): TableContent {
    return new TableContent(context, parent, attributes)
  }
}