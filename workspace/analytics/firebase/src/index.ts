import { addSink } from '@bosca/analytics-client-api'
import { FirebaseAnalyticsSink } from './stream'
import { Analytics } from 'firebase/analytics'

export * from './stream'

let initialized = false

export function initialize(analytics: Analytics) {
  if (initialized) return
  initialized = true
  addSink(new FirebaseAnalyticsSink(analytics))
}
