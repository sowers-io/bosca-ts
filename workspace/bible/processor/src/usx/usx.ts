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

import { Attributes, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { BookIdentification, BookIdentificationFactory } from './book_identification'
import { BookHeader, BookHeaderFactory } from './book_header'
import { BookTitle, BookTitleFactory } from './book_title'
import { BookIntroduction, BookIntroductionFactory, BookIntroductionTableFactory } from './book_introduction'
import { BookIntroductionEndTitle, BookIntroductionEndTitleFactory } from './book_introduction_end_titles'
import { BookChapterLabel, BookChapterLabelFactory } from './book_chapter_label'
import { ChapterStart, ChapterStartFactory } from './chapter_start'
import { ChapterEnd, ChapterEndFactory } from './chapter_end'
import { ParagraphFactory } from './paragraph'
import { ListFactory } from './list'
import { FootnoteFactory } from './footnote'
import { CrossReferenceFactory } from './cross_reference'
import { TextFactory } from './text'
import { TableFactory } from './table'

type UsxType =
  BookIdentification
  | BookHeader
  | BookTitle
  | BookIntroduction
  | BookIntroductionEndTitle
  | BookChapterLabel
  | ChapterStart
  | ChapterEnd

export class Usx extends UsxItemContainer<UsxType> {

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
  }

  get htmlClass(): string {
    return ''
  }
}

export class UsxFactory extends UsxItemFactory<Usx> {

  private static _instance?: UsxFactory

  static get instance(): UsxFactory {
    if (this._instance == null) {
      this._instance = new UsxFactory()
      this._instance.initialize()
    }
    return this._instance
  }

  private constructor() {
    super('usx')
  }

  protected onInitialize() {
    this.register(BookIdentificationFactory.instance)
    this.register(BookHeaderFactory.instance)
    this.register(BookTitleFactory.instance)
    this.register(BookIntroductionFactory.instance)
    this.register(BookIntroductionTableFactory.instance)
    this.register(BookIntroductionEndTitleFactory.instance)
    this.register(BookChapterLabelFactory.instance)
    this.register(ChapterStartFactory.instance)
    this.register(ChapterEndFactory.instance)
    this.register(ParagraphFactory.instance)
    this.register(ListFactory.instance)
    this.register(TableFactory.instance)
    this.register(FootnoteFactory.instance)
    this.register(CrossReferenceFactory.instance)
    // this.register(Sidebar)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): Usx {
    return new Usx(context, attributes)
  }
}