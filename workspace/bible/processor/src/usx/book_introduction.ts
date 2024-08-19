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

import { BookIntroductionStyle, BookIntroductionStyles } from './styles'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { Text, TextFactory } from './text'
import { Milestone, MilestoneFactory } from './milestone'
import { Figure, FigureFactory } from './figure'
import { Reference, ReferenceFactory } from './reference'
import { Footnote, FootnoteFactory } from './footnote'
import { CrossReference, CrossReferenceFactory } from './cross_reference'
import { Char, CharFactory } from './char'
import { IntroChar, IntroCharFactory } from './intro_char'
import { TableFactory } from './table'

type BookIntroductionType = Reference | Footnote | CrossReference | Char | IntroChar | Milestone | Figure | Text

export class BookIntroduction extends UsxItemContainer<BookIntroductionType> {
  style: BookIntroductionStyle

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as BookIntroductionStyle
  }

  get htmlClass(): string {
    return 'book-introduction ' + this.style
  }
}

export class BookIntroductionFactory extends UsxItemFactory<BookIntroduction> {

  static readonly instance = new BookIntroductionFactory()

  private constructor() {
    super('para', new StyleFactoryFilter(BookIntroductionStyles))
  }

  protected onInitialize() {
    this.register(ReferenceFactory.instance)
    this.register(FootnoteFactory.instance)
    this.register(CrossReferenceFactory.instance)
    this.register(CharFactory.instance)
    this.register(IntroCharFactory.instance)
    this.register(MilestoneFactory.instance)
    this.register(FigureFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): BookIntroduction {
    return new BookIntroduction(context, attributes)
  }
}

export class BookIntroductionTableFactory extends UsxItemFactory<BookIntroduction> {

  static readonly instance = new BookIntroductionTableFactory()

  private constructor() {
    super('table', new StyleFactoryFilter(BookIntroductionStyles))
  }

  protected onInitialize() {
    this.register(TableFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): BookIntroduction {
    return new BookIntroduction(context, attributes)
  }
}