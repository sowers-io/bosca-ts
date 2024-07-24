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

import { USXProcessor } from './processor';

test('USX Processor', async () => {
  const processor = new USXProcessor()
  await processor.process('../../../example-data/asv.zip')

  console.log(processor.metadata)

  const book = processor.books[0]
  const chapter = book.chapters[0]

  console.log(chapter.toString())

  console.log(book.raw.substring(chapter.position.start, chapter.position.end))
})