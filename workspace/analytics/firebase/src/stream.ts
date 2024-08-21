import { AnalyticEventSink, AnalyticEvent } from '@bosca/analytics-client-api'

import { Analytics, logEvent  } from 'firebase/analytics'

export class FirebaseAnalyticsSink extends AnalyticEventSink {
  
  private analytics: Analytics

  constructor(analytics: Analytics) {
    super()
    this.analytics = analytics
  }

  async onAdd(event: AnalyticEvent): Promise<void> {
    logEvent(this.analytics, event.name, event.toParameters())
  }
}