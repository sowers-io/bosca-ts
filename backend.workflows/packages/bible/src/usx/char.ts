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

import { CharStyle, CharStyles } from './styles'
import { Attributes, StyleFactoryFilter, UsxContext, UsxItemContainer, UsxItemFactory } from './item'
import { Footnote, FootnoteFactory } from './footnote'
import { Break, BreakFactory } from './break'
import { Text, TextFactory } from './text'
import { Reference, ReferenceFactory } from './reference'
import { Milestone, MilestoneFactory } from './milestone'

type CharType = Reference | Char | Milestone | Footnote | Break | Text

export class Char extends UsxItemContainer<CharType> {

  style: CharStyle
  //char.link?,
  //char.closed?,

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as CharStyle
  }
}

export class CharFactory extends UsxItemFactory<Char> {

  static readonly instance = new CharFactory()

  private constructor() {
    super('char', new StyleFactoryFilter(CharStyles))
  }

  protected onInitialize() {
    this.register(ReferenceFactory.instance)
    this.register(CharFactory.instance)
    this.register(MilestoneFactory.instance)
    this.register(FootnoteFactory.instance)
    this.register(BreakFactory.instance)
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, attributes: Attributes): Char {
    return new Char(context, attributes)
  }
}
