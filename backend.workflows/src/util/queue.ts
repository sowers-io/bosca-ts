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
  private readonly maxActive: number
  private active = 0
  private items: PromiseFactory[] = []
  private waiter: Promise<void> | null = null
  private waitNotifier: (() => void) | null = null

  constructor(name: string, maxActive: number) {
    this.name = name
    this.maxActive = maxActive
  }

  async enqueue(factory: PromiseFactory): Promise<void> {
    this.items.push(factory)
    console.log('queue::enqueue', this.name, 'active:', this.active, 'items:', this.items.length)
    await this.processNext()
  }

  private async processNext(): Promise<void> {
    while (this.waiter) {
      console.log('queue::processNextItem::waiting', this.name, 'active:', this.active, 'items:', this.items.length)
      await this.waiter
      console.log('queue::processNextItem::not waiting', this.name, 'active:', this.active, 'items:', this.items.length)
    }
    const factory = this.items.shift()
    if (!factory) return
    this.processNextItem(factory)
  }

  private processNextItem(factory: PromiseFactory) {
    const queue = this
    queue.active++
    console.log('queue::processNextItem::start', this.name, 'active:', queue.active, 'items:', this.items.length)
    if (queue.active === queue.maxActive) {
      queue.waiter = new Promise((resolve) => {
        queue.waitNotifier = resolve
      })
    }
    factory()
      .catch((e) => {
        console.error(
          'queue::processNextItem::error',
          this.name,
          'active:',
          queue.active,
          'items:',
          this.items.length,
          'error: ',
          e
        )
      })
      .finally(() => {
        queue.active--
        console.log('queue::processNextItem::done', this.name, 'active:', queue.active, 'items:', this.items.length)
        if (queue.active < queue.maxActive) {
          const notifier = queue.waitNotifier
          queue.waiter = null
          queue.waitNotifier = null
          if (notifier) {
            notifier()
          }
        } else {
          console.error(
            'queue::processNextItem::still full',
            this.name,
            'active:',
            queue.active,
            'items:',
            this.items.length
          )
        }
      })
  }
}
