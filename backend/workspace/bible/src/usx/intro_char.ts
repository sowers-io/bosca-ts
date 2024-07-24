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

import { Reference } from './reference'
import { Milestone } from './milestone'
import { Footnote } from './footnote'
import { Break } from './break'
import { Text } from './text'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { IntroCharStyle, IntroCharStyles } from './styles'
import { Char } from './char'

type IntroCharType = Reference | Char | IntroChar | Milestone | Footnote | Break | Text

export class IntroChar extends UsxItemContainer<IntroCharType> {

  style: IntroCharStyle

  // char.closed?

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as IntroCharStyle
  }
}

export class IntroCharFactory extends UsxItemFactory<IntroChar> {
  static readonly instance = new IntroCharFactory()

  private constructor() {
    super('char', new StyleFactoryFilter(IntroCharStyles))
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): IntroChar {
    return new IntroChar(context, attributes)
  }
}