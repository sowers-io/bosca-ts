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

import { Char } from './char'
import { Reference } from './reference'
import { Text, TextFactory } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { FootnoteCharStyle, FootnoteCharStyles } from './styles'
import { FootnoteVerse } from './footnote_verse'

type FootnoteCharType = Char | FootnoteChar | FootnoteVerse | Reference | Text

export class FootnoteChar extends UsxItemContainer<FootnoteCharType> {

  style: FootnoteCharStyle
  // char.link?,
  // char.closed?,

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as FootnoteCharStyle
  }

  get htmlClass(): string {
    return this.style
  }
}

export class FootnoteCharFactory extends UsxItemFactory<FootnoteChar> {

  static readonly instance = new FootnoteCharFactory()

  private constructor() {
    super('char', new StyleFactoryFilter(FootnoteCharStyles))
  }

  protected onInitialize() {
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): FootnoteChar {
    return new FootnoteChar(context, attributes)
  }
}
