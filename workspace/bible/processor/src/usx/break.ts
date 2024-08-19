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

import { Attributes, HtmlContext, UsxContext, UsxItem, UsxItemFactory } from './item'
import { Position } from './position'

export class Break implements UsxItem {

  readonly verse: string | null
  readonly position: Position

  constructor(context: UsxContext, parent: UsxItem | null) {
    this.position = context.position
    this.verse = context.addVerseItem(parent, this)
  }

  get htmlClass(): string {
    return ''
  }

  get htmlAttributes(): { [p: string]: string } {
    return {}
  }

  toHtml(context: HtmlContext): string {
    return context.render('br', this)
  }

  toString(): string {
    return ''
  }
}

export class BreakFactory extends UsxItemFactory<Break> {

  static readonly instance = new BreakFactory()

  private constructor() {
    super('optbreak')
  }

  protected onInitialize() {
  }

  create(context: UsxContext, parent: UsxItem | null, _: Attributes): Break {
    return new Break(context, parent)
  }
}