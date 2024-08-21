import { AnalyticEvent } from './event'

export interface AnalyticEventInterceptor {

  intercept(event: AnalyticEvent): Promise<AnalyticEvent>
}

export abstract class AnalyticEventSink {

  private interceptors: AnalyticEventInterceptor[] = []

  addInterceptor(interceptor: AnalyticEventInterceptor) {
    this.interceptors.push(interceptor)
  }

  protected abstract onAdd(original: AnalyticEvent, event: AnalyticEvent): Promise<void>

  async add(event: AnalyticEvent): Promise<void> {
    const original = event
    for (const interceptor of this.interceptors) {
      event = await interceptor.intercept(event.clone())
    }
    await this.onAdd(original, event)
  }
}