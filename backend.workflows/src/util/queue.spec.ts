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

import { Queue } from './queue'

test('test await enqueue', async () => {
  const queue = new Queue('test', 4)
  const items = 10000

  for (let i = 0; i < items; i++) {
    expect(queue.active).toBeLessThanOrEqual(queue.maxActive)
    expect(queue.queued).toBeLessThanOrEqual(1)
    queue.enqueue(() => new Promise((resolve) => setTimeout(resolve, 10)))
    await queue.process()
  }

  expect(queue.processed).toEqual(items)
  expect(queue.waited).toEqual(items - queue.maxActive)
  expect(queue.waiting).toEqual(0)
}, 60000)

test('test no await enqueue', async () => {
  const queue = new Queue('test', 4)
  const items = 100

  function enqueue(timeout: number) {
    queue.enqueue(() => new Promise((resolve) => setTimeout(resolve, timeout)))
  }

  const processing = []
  for (let i = 0; i < items; i++) {
    expect(queue.active).toBeLessThanOrEqual(queue.maxActive)
    enqueue(i < 4 ? 100 : 1)
    processing.push(queue.process())
  }

  expect(queue.queued).toEqual(items - queue.maxActive)
  expect(queue.active).toEqual(queue.maxActive)
  expect(queue.waiting).toEqual(items - queue.maxActive)

  await Promise.all(processing)

  expect(queue.processed).toEqual(items)
  expect(queue.waited).toEqual(items - queue.maxActive)
}, 60000)
