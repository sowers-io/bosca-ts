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

import { Attributes, HtmlContext, StyleFactoryFilter, UsxContext, UsxItem, UsxItemFactory } from './item'
import { VerseStartStyle, VerseStartStyles } from './styles'
import { Position } from './position'

export class VerseStart implements UsxItem {

  readonly style: VerseStartStyle
  readonly number: string
  readonly altnumber?: string
  readonly pubnumber?: string
  readonly sid: string
  readonly verse: string | null
  readonly position: Position

  constructor(context: UsxContext, attributes: Attributes) {
    this.position = context.position
    this.style = attributes.STYLE.toString() as VerseStartStyle
    this.number = attributes.NUMBER.toString()
    this.altnumber = attributes.ALTNUMBER?.toString()
    this.pubnumber = attributes.PUBNUMBER?.toString()
    this.sid = attributes.SID.toString()
    this.verse = context.addVerseItem(this)
  }

  get htmlClass(): string {
    return this.style
  }

  get htmlAttributes(): { [p: string]: string } {
    return {
      'data-verse': this.number
    }
  }

  toHtml(context: HtmlContext): string {
    if (context.includeVerseNumbers) return context.render('span', this, this.number)
    return ''
  }

  toString(): string {
    return this.number + '. '
  }
}

export class VerseStartFactory extends UsxItemFactory<VerseStart> {

  static readonly instance = new VerseStartFactory()

  constructor() {
    super('verse', new StyleFactoryFilter(VerseStartStyles))
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): VerseStart {
    return new VerseStart(context, attributes)
  }
}