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

import { Attributes, EndIdFactoryFilter, UsxContext, UsxItem, UsxItemFactory } from './item'
import { Position } from './position'

export class VerseEnd implements UsxItem {

  eid: string
  readonly verse: string | null
  readonly position: Position

  constructor(context: UsxContext, attributes: Attributes) {
    this.position = context.position
    this.eid = attributes.EID.toString()
    this.verse = context.verse
  }

  toString(): string {
    return ''
  }
}

export class VerseEndFactory extends UsxItemFactory<VerseEnd> {

  static readonly instance = new VerseEndFactory()

  constructor() {
    super('verse', new EndIdFactoryFilter())
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): VerseEnd {
    return new VerseEnd(context, attributes)
  }
}