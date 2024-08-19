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

import { BookIdentificationCode, BookIdentificationCodes } from '../identification'
import { Attributes, CodeFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { Text, TextFactory } from './text'

export class BookIdentification extends UsxItemContainer<Text> {
  id!: string
  code: BookIdentificationCode

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.id = attributes.STYLE.toString()
    this.code = attributes.CODE.toString() as BookIdentificationCode
  }

  get htmlClass(): string {
    return 'book-identification'
  }

  get htmlAttributes(): { [p: string]: string } {
    return {
      'data-id': this.id,
      'data-code': this.code,
      ...super.htmlAttributes,
    }
  }
}

export class BookIdentificationFactory extends UsxItemFactory<BookIdentification> {

  static readonly instance = new BookIdentificationFactory()

  private constructor() {
    super('book', new CodeFactoryFilter<BookIdentificationCode>(BookIdentificationCodes))
  }

  protected onInitialize() {
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): BookIdentification {
    return new BookIdentification(context, attributes)
  }
}