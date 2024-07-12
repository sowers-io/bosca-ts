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

import { SidebarStyle } from './styles'
import { Attributes, UsxContext, UsxItemContainer } from './item'
import { Paragraph } from './paragraph'
import { List } from './list'
import { Table } from './table'
import { Footnote } from './footnote'
import { CrossReference } from './cross_reference'

type SidebarType = Paragraph | List | Table | Footnote | CrossReference

export class Sidebar extends UsxItemContainer<SidebarType> {

  style: SidebarStyle
  category?: String

  constructor(context: UsxContext, attributes: Attributes) {
    super(context, attributes)
    this.style = attributes.STYLE.toString() as SidebarStyle
    this.category = attributes.CATEGORY?.toString()
  }
}