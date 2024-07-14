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

export class Retry {
  private readonly retries: number
  private readonly backoff: number
  private readonly maxBackoff: number
  private currentBackoff: number

  constructor(retries: number, backoff: number = 500, maxBackoff: number = 20000) {
    this.retries = retries
    this.backoff = backoff
    this.maxBackoff = maxBackoff
    this.currentBackoff = 0
  }

  async execute<T>(fn: () => Promise<T>): Promise<T> {
    let tries = 0
    const retrier = this
    while (tries < this.retries) {
      tries++
      try {
        return await fn()
      } catch (error) {
        if (tries === this.retries - 1) {
          throw error
        }
        this.currentBackoff = Math.min(this.currentBackoff + this.backoff, this.maxBackoff)
        console.error('failed to execute function, retrying... backoff: ' + this.currentBackoff + 'ms', error)
        await new Promise((resolve) => setTimeout(resolve, retrier.currentBackoff))
      }
    }
    throw new Error('retry failed')
  }

  public static execute(retries: number, fn: () => Promise<any>) {
    return new Retry(retries).execute(fn)
  }
}
