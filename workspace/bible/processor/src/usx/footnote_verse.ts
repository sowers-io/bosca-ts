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
  StringContext,
  StyleFactoryFilter,
  UsxContext,
  UsxItem,
  UsxItemContainer,
  UsxItemFactory,
} from './item'
import { Text } from './text'
import { FootnoteVerseStyle, FootnoteVerseStyles } from './styles'

export class FootnoteVerse extends UsxItemContainer<Text> {

  style: FootnoteVerseStyle

  constructor(context: UsxContext, parent: UsxItem | null, attributes: Attributes) {
    super(context, parent, attributes)
    this.style = attributes.STYLE.toString() as FootnoteVerseStyle
  }

  get htmlClass(): string {
    return this.style
  }

  toString(context: StringContext | undefined = undefined): string {
    const ctx = context || StringContext.defaultContext
    if (!ctx.includeFootNotes) return ''
    return super.toString(context)
  }
}

export class FootnoteVerseFactory extends UsxItemFactory<FootnoteVerse> {

  constructor() {
    super('char', new StyleFactoryFilter(FootnoteVerseStyles))
  }

  protected onInitialize() {
  }

  create(context: UsxContext, parent: UsxItem | null, attributes: Attributes): FootnoteVerse {
    return new FootnoteVerse(context, parent, attributes)
  }
}