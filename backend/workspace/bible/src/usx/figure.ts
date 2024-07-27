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
import { Text } from './text'

export class Figure extends UsxItemContainer<Text> {

  style: string
  alt?: string
  file: string
  size?: string
  loc?: string
  copy?: string
  ref?: string

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString()
    this.alt = attributes.ALT?.toString()
    this.file = attributes.FILE.toString()
    this.size = attributes.SIZE?.toString()
    this.loc = attributes.LOC?.toString()
    this.copy = attributes.COPY?.toString()
    this.ref = attributes.REF?.toString()
  }
}

export class FigureFactory extends UsxItemFactory<Figure> {

  static readonly instance = new FigureFactory()

  private constructor() {
    super('figure')
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): Figure {
    return new Figure(context, attributes)
  }
}