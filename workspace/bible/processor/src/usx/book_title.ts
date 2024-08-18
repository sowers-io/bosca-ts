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

import { Footnote } from './footnote'
import { CrossReference } from './cross_reference'
import { Char } from './char'
import { Break } from './break'
import { Text, TextFactory } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { BookTitleStyle, BookTitleStyles } from './styles'

type BookTitleType = Footnote | CrossReference | Char | Break | Text

export class BookTitle extends UsxItemContainer<BookTitleType> {
  style: BookTitleStyle

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as BookTitleStyle
  }

  get htmlClass(): string {
    return this.style
  }
}

export class BookTitleFactory extends UsxItemFactory<BookTitle> {

  static readonly instance = new BookTitleFactory()

  private constructor() {
    super('para', new StyleFactoryFilter(BookTitleStyles))
  }

  protected onInitialize() {
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): BookTitle {
    return new BookTitle(context, attributes)
  }
}