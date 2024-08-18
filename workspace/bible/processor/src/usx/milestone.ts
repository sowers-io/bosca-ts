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

import { Attributes, HtmlContext, UsxContext, UsxItem, UsxItemFactory } from './item'
import { Position } from './position'

export class Milestone implements UsxItem {

  readonly style: string
  readonly sid: string
  readonly eid: string

  readonly position: Position
  readonly verse: string | null

  constructor(context: UsxContext, attributes: Attributes) {
    this.position = context.position
    this.style = attributes.STYLE.toString()
    this.sid = attributes.SID.toString()
    this.eid = attributes.EID.toString()
    this.verse = context.addVerseItem(this)
  }

  get htmlClass(): string {
    return this.style
  }

  get htmlAttributes(): { [p: string]: string } {
    return {}
  }

  toHtml(context: HtmlContext): string {
    return context.render('milestone', this)
  }

  toString(): string {
    return ''
  }
}

export class MilestoneFactory extends UsxItemFactory<Milestone> {

  static readonly instance = new MilestoneFactory()

  private constructor() {
    super('ms')
  }

  protected onInitialize() {
  }

  create(context: UsxContext, attributes: Attributes): Milestone {
    return new Milestone(context, attributes)
  }
}