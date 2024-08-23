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

import { test } from 'vitest'
import { BibleReference } from './bible'
import { USXProcessor } from './processor'

test('John 1:1', async () => {
  const processor = new USXProcessor()
  const bible = await processor.process('../../../example-data/asv.zip')
  const references = BibleReference.parse(bible, 'John 1:1')

  expect(references.length).toBe(1)
  expect(references[0].usfm).toBe('JHN.1.1')
})

test('John 1:1, John 1:2, John 2:1-3', async () => {
  const processor = new USXProcessor()
  const bible = await processor.process('../../../example-data/asv.zip')
  const references = BibleReference.parse(bible, 'John 1:1, John 1:2, John 2:1-3')

  expect(references.length).toBe(2)
  expect(references[0].usfm).toBe('JHN.1.1+JHN.1.2')
  expect(references[1].usfm).toBe('JHN.2.1+JHN.2.2+JHN.2.3')
})

test('Isaiah 2:2–3', async () => {
  const processor = new USXProcessor()
  const bible = await processor.process('../../../example-data/asv.zip')
  const references = BibleReference.parse(bible, 'Isaiah 2:2–3')
  expect(references.length).toBe(1)
  expect(references[0].usfm).toBe('ISA.2.2+ISA.2.3')
})

test('Get Verses', async () => {
  const processor = new USXProcessor()
  const bible = await processor.process('../../../example-data/asv.zip')
  const references = BibleReference.parse(bible, 'John 1:1, John 1:2, John 2:1-3')

  const items = bible.getVerseItems(references)

  expect(items.length).toBe(5)
})