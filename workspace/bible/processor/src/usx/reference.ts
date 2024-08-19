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

import { Attributes, UsxContext, UsxItem, UsxItemContainer, UsxItemFactory } from './item'
import { Text, TextFactory } from './text'

export class Reference extends UsxItemContainer<Text> {

  loc: string

  constructor(context: UsxContext, parent: UsxItem | null, attributes: Attributes) {
    super(context, parent, attributes)
    this.loc = attributes.LOC.toString()
  }

  get htmlClass(): string {
    return ''
  }
}

export class ReferenceFactory extends UsxItemFactory<Reference> {

  static readonly instance = new ReferenceFactory()

  private constructor() {
    super('ref')
  }

  protected onInitialize() {
    this.register(TextFactory.instance)
  }

  create(context: UsxContext, parent: UsxItem | null, attributes: Attributes): Reference {
    return new Reference(context, parent, attributes)
  }
}
