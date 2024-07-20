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

export type PromiseFactory = () => Promise<any>

export class Queue {
  readonly name: string
  readonly maxActive: number

  private _active = 0
  private _processed = 0
  private _waited = 0
  private _waiting = 0
  private items: PromiseFactory[] = []
  private _waiter: Promise<void> | null = null
  private _waitNotifier: (() => void) | null = null
  private _blocking = false

  constructor(name: string, maxActive: number) {
    this.name = name
    this.maxActive = maxActive
  }

  get processed(): number {
    return this._processed
  }

  get waited(): number {
    return this._waited
  }

  get waiting(): number {
    return this._waiting
  }

  get active(): number {
    return this._active
  }

  get queued(): number {
    return this.items.length
  }

  get waiter(): Promise<void> | null {
    return this._waiter
  }

  enqueue(factory: PromiseFactory) {
    this.items.push(factory)
  }

  async process(): Promise<void> {
    let waited = false
    while (this._blocking) {
      if (!waited) {
        waited = true
        this._waited++
      }
      this._waiting++
      await this.waiter
      this._waiting--
    }
    const factory = this.items.shift()
    if (!factory) return
    await factory()
    this.processNextItem(factory)
  }

  private processNextItem(factory: PromiseFactory) {
    const queue = this
    queue._processed++
    queue._active++
    if (queue.active === queue.maxActive) {
      queue._blocking = true
      if (queue._waiter) {
        throw new Error('queue::processNextItem::waiter should be null')
      }
      queue._waiter = new Promise((resolve) => {
        queue._waitNotifier = resolve
      })
    }
    factory()
      .catch((e) => {
        console.error(
          'queue::processNextItem::error',
          queue.name,
          'active:',
          queue.active,
          'items:',
          queue.items.length,
          'error: ',
          e
        )
      })
      .finally(() => {
        queue._active--
        if (queue.active < queue.maxActive) {
          queue._blocking = false
          const notifier = queue._waitNotifier
          queue._waiter = null
          queue._waitNotifier = null
          if (notifier) {
            notifier()
          }
        } else {
          console.error(
            'queue::processNextItem::still full',
            queue.name,
            'active:',
            queue.active,
            'items:',
            queue.items.length
          )
        }
      })
  }
}
