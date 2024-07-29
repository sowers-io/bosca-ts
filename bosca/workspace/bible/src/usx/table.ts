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

import { Attributes, UsxContext, UsxItemContainer } from './item'
import { Footnote } from './footnote'
import { Char } from './char'
import { Milestone } from './milestone'
import { Figure } from './figure'
import { Break } from './break'
import { Text } from './text'
import { CrossReference } from './cross_reference'
import { Verse } from './verse'

export class Table extends UsxItemContainer<Row> {

  vid: string

  get element(): string {
    return 'table'
  }

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.vid = attributes.VID.toString()
  }
}

type RowType = Verse | TableContent

export class Row extends UsxItemContainer<RowType> {

  style: string

  get element(): string {
    return 'row'
  }

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString()
  }
}

type TableContentType = Footnote | CrossReference | Char | Milestone | Figure | Verse | Break | Text

export class TableContent extends UsxItemContainer<TableContentType> {
  style: string
  align: string
  colspan: string

  get element(): string {
    return 'cell'
  }

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString()
    this.align = attributes.ALIGN.toString()
    this.colspan = attributes.COLSPAN?.toString()
  }
}