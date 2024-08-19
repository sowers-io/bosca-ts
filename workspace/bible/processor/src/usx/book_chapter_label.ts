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

import { Attributes, StyleFactoryFilter, UsxContext, UsxItem, UsxItemContainer, UsxItemFactory } from './item'
import { Text, TextFactory } from './text'
import { BookChapterLabelStyle, BookChapterLabelStyles } from './styles'

export class BookChapterLabel extends UsxItemContainer<Text> {
  style: BookChapterLabelStyle

  constructor(context: UsxContext, parent: UsxItem | null, attributes: Attributes) {
    super(context, parent, attributes)
    this.style = attributes.STYLE.toString() as BookChapterLabelStyle
  }

  get htmlClass(): string {
    return this.style
  }
}

export class BookChapterLabelFactory extends UsxItemFactory<BookChapterLabel> {

  static readonly instance = new BookChapterLabelFactory()

  private constructor() {
    super('para', new StyleFactoryFilter(BookChapterLabelStyles))
    this.register(TextFactory.instance)
  }

  protected onInitialize() {
  }

  create(context: UsxContext, parent: UsxItem | null, attributes: Attributes): BookChapterLabel {
    return new BookChapterLabel(context, parent, attributes)
  }
}