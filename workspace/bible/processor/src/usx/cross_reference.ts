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

import { CrossReferenceChar, CrossReferenceCharFactory } from './cross_reference_char'
import { Text, TextFactory } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { CrossReferenceStyle, CrossReferenceStyles } from './styles'

type CrossReferenceType = CrossReferenceChar | Text

export class CrossReference extends UsxItemContainer<CrossReferenceType> {

  style: CrossReferenceStyle
  caller: string

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as CrossReferenceStyle
    this.caller = attributes.CALLER.toString()
  }
}

export class CrossReferenceFactory extends UsxItemFactory<CrossReference> {

  static readonly instance = new CrossReferenceFactory()

  private constructor() {
    super('note', new StyleFactoryFilter(CrossReferenceStyles))
  }

  protected onInitialize() {
    this.register(CrossReferenceCharFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): CrossReference {
    return new CrossReference(context, attributes)
  }
}