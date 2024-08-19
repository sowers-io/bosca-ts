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

import {
  Attributes,
  EndIdFactoryFilter,
  HtmlContext,
  NegateFactoryFilter,
  UsxContext,
  UsxItem,
  UsxItemFactory,
} from './item'
import { Position } from './position'

export class ChapterStart implements UsxItem {

  readonly number: string
  readonly sid: string
  readonly altnumber?: string
  readonly pubnumber?: string
  readonly position: Position

  readonly verse: string | null

  constructor(context: UsxContext, attributes: Attributes) {
    this.verse = context.addVerseItem(this)
    this.number = attributes.NUMBER.toString()
    this.sid = attributes.SID.toString()
    this.altnumber = attributes.ALTNUMBER?.toString()
    this.pubnumber = attributes.PUBNUMBER?.toString()
    this.position = context.position
  }

  get htmlClass(): string {
    return ''
  }

  get htmlAttributes(): { [p: string]: string } {
    return {}
  }

  toHtml(_: HtmlContext): string {
    return ''
  }

  toString(): string {
    return ''
  }
}

export class ChapterStartFactory extends UsxItemFactory<ChapterStart> {

  static readonly instance = new ChapterStartFactory()

  constructor() {
    super('chapter', new NegateFactoryFilter(new EndIdFactoryFilter()))
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): ChapterStart {
    return new ChapterStart(context, attributes)
  }
}