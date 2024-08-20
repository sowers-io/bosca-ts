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

import { Reference, ReferenceFactory } from './reference'
import { Text, TextFactory } from './text'
import {
  Attributes,
  HtmlContext,
  StringContext,
  StyleFactoryFilter,
  UsxContext,
  UsxItem,
  UsxItemContainer,
  UsxItemFactory,
} from './item'
import { CrossReferenceCharStyle, CrossReferenceCharStyles } from './styles'
import { Char, CharFactory } from './char'

type CrossReferenceCharType = Char | CrossReferenceChar | Reference | Text

export class CrossReferenceChar extends UsxItemContainer<CrossReferenceCharType> {

  style: CrossReferenceCharStyle
  //char.link?
  //char.closed?

  constructor(context: UsxContext, parent: UsxItem | null, attributes: Attributes) {
    super(context, parent, attributes)
    this.style = attributes.STYLE.toString() as CrossReferenceCharStyle
  }

  get htmlClass(): string {
    return this.style
  }

  toHtml(context: HtmlContext): string {
    if (!context.includeCrossReferences) return ''
    return super.toHtml(context)
  }

  toString(context: StringContext | undefined = undefined): string {
    const ctx = context || StringContext.defaultContext
    if (!ctx.includeCrossReferences) return ''
    return super.toString(context)
  }
}

export class CrossReferenceCharFactory extends UsxItemFactory<CrossReferenceChar> {

  static readonly instance = new CrossReferenceCharFactory()

  private constructor() {
    super('char', new StyleFactoryFilter(CrossReferenceCharStyles))
  }

  protected onInitialize() {
    this.register(CharFactory.instance)
    this.register(CrossReferenceCharFactory.instance)
    this.register(ReferenceFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, parent: UsxItem | null, attributes: Attributes): CrossReferenceChar {
    return new CrossReferenceChar(context, parent, attributes)
  }
}