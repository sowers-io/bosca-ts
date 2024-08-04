import { NodeSDK } from '@opentelemetry/sdk-node'
import { getNodeAutoInstrumentations } from '@opentelemetry/auto-instrumentations-node'
import { ConsoleMetricExporter, MeterProvider, PeriodicExportingMetricReader } from '@opentelemetry/sdk-metrics'
import { Resource } from '@opentelemetry/resources'
import { SemanticResourceAttributes } from '@opentelemetry/semantic-conventions'
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-http'
import { OTLPMetricExporter } from '@opentelemetry/exporter-metrics-otlp-http'
import { metrics } from '@opentelemetry/api'
import { PushMetricExporter } from '@opentelemetry/sdk-metrics/build/src/export/MetricExporter'
import { ConsoleSpanExporter } from '@opentelemetry/sdk-trace-node'
import { SpanExporter } from '@opentelemetry/sdk-trace-base'

let instrumentationSDK: NodeSDK

export function initializeInstrumentation(serviceName: string) {
  if (process.env.DISABLE_OTEL) return

  const resource = new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: serviceName,
  })

  const traceExporter: SpanExporter = process.env.OTLP_TRACE_ENDPOINT
    ? new OTLPTraceExporter({
      url: process.env.OTLP_TRACE_ENDPOINT + '/v1/traces',
    })
    : new ConsoleSpanExporter()

  const metricExporter: PushMetricExporter = process.env.OTLP_METRICS_ENDPOINT
    ? new OTLPMetricExporter({
      url: process.env.OTLP_METRICS_ENDPOINT + '/v1/metrics',
    })
    : new ConsoleMetricExporter()

  const metricReader = new PeriodicExportingMetricReader({
    exporter: metricExporter,
    exportIntervalMillis: 1000,
  })
  const meterProvider = new MeterProvider({
    resource: resource,
    readers: [metricReader],
  })
  metrics.setGlobalMeterProvider(meterProvider)

  instrumentationSDK = new NodeSDK({
    serviceName: serviceName,
    traceExporter: traceExporter,
    metricReader: new PeriodicExportingMetricReader({
      exporter: metricExporter,
    }),
    resource: resource,
    instrumentations: [getNodeAutoInstrumentations()],
  })
  instrumentationSDK.start()
  process.on('SIGTERM', () => {
    instrumentationSDK
      .shutdown()
      .then(() => console.log('Tracing terminated'))
      .catch((error) => console.log('Error terminating tracing', error))
      .finally(() => process.exit(0))
  })
}
