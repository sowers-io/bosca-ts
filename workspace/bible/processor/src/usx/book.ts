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

import { ManifestName, PublicationContent } from '../metadata'
import { Chapter } from './chapter'
import { Position } from './position'

export class Book {

  readonly name: ManifestName
  readonly content: PublicationContent
  readonly chapters: Chapter[] = []
  readonly raw: string

  constructor(name: ManifestName, content: PublicationContent, raw: string) {
    this.name = name
    this.content = content
    this.raw = raw
  }

  get usfm(): string {
    return this.content.usfm
  }

  getRawContent(position: Position): string {
    return this.raw.substring(position.start, position.end)
  }
}
