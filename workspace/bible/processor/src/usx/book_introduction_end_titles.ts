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

import { Footnote, FootnoteFactory } from './footnote'
import { CrossReference, CrossReferenceFactory } from './cross_reference'
import { Char, CharFactory } from './char'
import { Milestone, MilestoneFactory } from './milestone'
import { Break, BreakFactory } from './break'
import { Text, TextFactory } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItem, UsxItemContainer, UsxItemFactory } from './item'
import { BookIntroductionEndTitleStyle, BookIntroductionEndTitleStyles } from './styles'


type BookIntroductionEndTitleType = Footnote | CrossReference | Char | Milestone | Break | Text

export class BookIntroductionEndTitle extends UsxItemContainer<BookIntroductionEndTitleType> {
  style: BookIntroductionEndTitleStyle

  constructor(context: UsxContext, parent: UsxItem | null, attributes: Attributes) {
    super(context, parent, attributes)
    this.style = attributes.STYLE.toString() as BookIntroductionEndTitleStyle
  }

  get htmlClass(): string {
    return this.style
  }
}

export class BookIntroductionEndTitleFactory extends UsxItemFactory<BookIntroductionEndTitle> {

  static readonly instance = new BookIntroductionEndTitleFactory()

  private constructor() {
    super('para', new StyleFactoryFilter(BookIntroductionEndTitleStyles))
  }

  protected onInitialize() {
    this.register(FootnoteFactory.instance)
    this.register(CrossReferenceFactory.instance)
    this.register(CharFactory.instance)
    this.register(MilestoneFactory.instance)
    this.register(BreakFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, parent: UsxItem | null, attributes: Attributes): BookIntroductionEndTitle {
    return new BookIntroductionEndTitle(context, parent, attributes)
  }
}