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

/* eslint-disable @typescript-eslint/no-unused-vars */
import { Observable, Subscription } from 'rxjs'

export enum PagingState { idle = 'idle', paging = 'paging', done = 'done' }

export interface IPageDataSource {
}

// noinspection JSUnusedLocalSymbols
export abstract class PageDataSource<T> implements IPageDataSource {
  get pageSize() {
    return 25
  }

  get maxPageBuffer() {
    return 5
  }

  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  isFetchable(item: T): boolean {
    return true
  }

  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  willFetch(pageIndex: number): boolean {
    return true
  }

  // eslint-disable-next-line @typescript-eslint/ban-ts-comment
  // @ts-ignore
  shouldFetch(pageIndex: number): boolean {
    return true
  }

  abstract fetchPage(pageIndex: number): Promise<boolean>

  abstract getPage(pageIndex: number): Promise<Observable<T[]> | undefined>

  close(): void {
  }
}

export interface IPageListener<T> {
  onPageChanged(pageIndex: number, items: Array<T>): void

  onSizeChanged(pageIndex: number, newSize: number): void

  onError(pageIndex: number, exception: any): void

  onStateChange(oldState: PagingState, newState: PagingState): void
}

export class PagedList<T> {

  private readonly dataSource: PageDataSource<T>
  private readonly pages = new Map<number, Array<T>>()
  private readonly fetchingPages = new Set<number>()
  private readonly fetchedPages = new Set<number>()
  private readonly subscriptions = new Map<number, Subscription>()

  private listener: IPageListener<T>
  private _state = PagingState.idle
  private _size: number = -1
  private startPage: number = -1
  private endPage: number = -1
  private currentPageIndex: number
  private activeFetches: number = 0
  private closed = false

  constructor(dataSource: PageDataSource<T>, listener: IPageListener<T>) {
    this.listener = listener
    this.dataSource = dataSource
    this.currentPageIndex = -1 * dataSource.maxPageBuffer
    this.moveTo(0)
    void this.fetchPage(0)
  }

  get state(): PagingState {
    return this._state
  }

  get size(): number {
    const s = this._size == -1 ? 0 : this._size
    if (this.state != PagingState.done) {
      return s + 1
    }
    return s
  }

  close(): void {
    this.closed = true
    for (const pageIndex of this.subscriptions.keys()) {
      this.subscriptions.get(pageIndex)?.unsubscribe()
    }
    this.subscriptions.clear()
    this.dataSource.close()
    this.pages.clear()
  }

  private cancel(pageIndex: number): void {
    this.pages.delete(pageIndex)
    this.fetchedPages.delete(pageIndex)
    this.subscriptions.get(pageIndex)?.unsubscribe()
    this.subscriptions.delete(pageIndex)
  }

  private cancelIfNecessary(pageIndex: number): boolean {
    if (this.closed || pageIndex < this.startPage || pageIndex > this.endPage) {
      this.cancel(pageIndex)
      return true
    }
    return false
  }

  private moveTo(pageIndex: number): void {
    if (pageIndex == this.currentPageIndex) return
    this.currentPageIndex = pageIndex
    const pageBuffer = ~~(this.dataSource.maxPageBuffer / 2)
    if (pageIndex <= this.startPage || pageIndex >= this.endPage) {
      this.startPage = pageIndex - pageBuffer
      this.endPage = pageIndex + pageBuffer
      if (this.startPage < 0) {
        this.endPage += pageBuffer
        this.startPage = 0
      }
      const pageKeys = Array.of(...this.pages.keys())
      for (const pageIndex of pageKeys) {
        this.cancelIfNecessary(pageIndex)
      }
      for (let index = this.startPage; index < this.endPage; index++) {
        void this.subscribeIfNecessary(index)
      }
    }
  }

  private hasMore = true

  private async fetchPage(pageIndex: number): Promise<void> {
    if (!this.hasMore || this.fetchingPages.has(pageIndex) || this.fetchedPages.has(pageIndex)) {
      return
    }
    this.fetchingPages.add(pageIndex)
    this.activeFetches++
    if (this.activeFetches === 1) {
      this.setState(PagingState.paging)
    }
    try {
      this.hasMore = await this.dataSource.fetchPage(pageIndex)
      this.fetchedPages.add(pageIndex)
    } catch (e) {
      this.listener.onError(pageIndex, e)
    } finally {
      this.activeFetches--
      if (this.activeFetches === 0) {
        this.setState(this.hasMore ? PagingState.idle : PagingState.done)
      }
      this.fetchingPages.delete(pageIndex)
    }
  }

  private setState(state: PagingState): void {
    const oldState = this.state
    this._state = state
    this.listener.onStateChange(oldState, state)
    if (state == PagingState.done && this.size == -1) {
      this._size = 0
      this.listener.onSizeChanged(0, 0)
    }
  }

  private isEqual(a: Array<T> | undefined, b: Array<T> | undefined): boolean {
    if (a == undefined || b == undefined) {
      return a === b
    }
    if (a.length !== b.length) {
      return false
    }
    for (let i = 0; i < a.length; i++) {
      if (a[i] !== b[i]) return false
    }
    return true
  }


  private updateSize(page: Array<T>, pageIndex: number) {
    if (page.length !== 0) {
      const newSize = (pageIndex * this.dataSource.pageSize) + page.length
      if (newSize > this.size) {
        this._size = newSize
        this.listener.onSizeChanged(pageIndex, newSize)
      }
    }
  }

  private async subscribeIfNecessary(pageIndex: number): Promise<void> {
    if (this.pages.get(pageIndex) != null) return
    this.pages.set(pageIndex, [])
    const page = await this.dataSource.getPage(pageIndex)
    if (this.cancelIfNecessary(pageIndex)) return
    if (!page) return
    // eslint-disable-next-line @typescript-eslint/no-this-alias
    const list = this
    const subscription = page.subscribe((page) => {
      if (list.closed) return
      const lastPage = list.pages.get(pageIndex)
      if (!list.isEqual(lastPage, page)) {
        list.pages.set(pageIndex, page)
        list.listener.onPageChanged(pageIndex, page)
      }
      list.updateSize(page, pageIndex)
    })
    this.subscriptions.set(pageIndex, subscription)
  }

  isLoadingItem(index: number): boolean {
    return this.state != PagingState.done && this.size - 1 == index
  }

  private fetchIfNeeded(pageIndex: number, pageItemIndex: number, item: T | null) {
    const diff = this.dataSource.pageSize - pageItemIndex
    if (diff <= 3 || pageItemIndex < 4) {
      if (this.dataSource.shouldFetch(pageIndex) &&
        (item == null || this.dataSource.isFetchable(item))) {
        this.dataSource.willFetch(pageIndex)
        setTimeout(() => this.fetchPage(pageIndex), 0)
      }
    }
  }

  get(index: number): T | null {
    const pageIndex = ~~(index / this.dataSource.pageSize)
    this.moveTo(pageIndex)
    const page = this.pages.get(pageIndex)
    const pageItemIndex = index % this.dataSource.pageSize
    const item = page == null
      ? null
      : page.length > pageItemIndex
        ? page[pageItemIndex]
        : null
    this.fetchIfNeeded(pageIndex, pageItemIndex, item)
    return item
  }

  range(start: number, end: number): (T | null)[] {
    const items: (T | null)[] = []
    for (let i = start; i < end; i++) {
      items.push(this.get(i))
    }
    return items
  }
}