export type PromiseFactory = () => Promise<any>

export class Queue {
  private readonly name: string
  private readonly max: number
  private active: number = 0
  private pending: Promise<any> = Promise.resolve()
  private pendingResolve: () => void = () => {}

  constructor(name: string, max: number) {
    this.name = name
    this.max = max
  }

  async enqueue(factory: PromiseFactory): Promise<void> {
    await this.pending
    const queue = this
    if (this.active >= this.max) {
      this.pending = new Promise<void>((resolve) => {
        if (queue.active < queue.max) {
          resolve()
        } else {
          console.log('queue (' + queue.name + ') is full, waiting...')
          queue.pendingResolve = resolve
        }
      })
    } else {
      this.active++
      factory().finally(() => {
        queue.active--
        if (queue.active < queue.max) {
          console.log('queue (' + queue.name + ') is not full, continuing...')
          queue.pendingResolve()
          queue.pendingResolve = () => {}
        }
      })
    }
  }
}
