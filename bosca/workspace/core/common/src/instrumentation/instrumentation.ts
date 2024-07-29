import { NodeSDK } from '@opentelemetry/sdk-node'
import { getNodeAutoInstrumentations } from '@opentelemetry/auto-instrumentations-node'
import { MeterProvider, PeriodicExportingMetricReader } from '@opentelemetry/sdk-metrics'
import { Resource } from '@opentelemetry/resources'
import { SemanticResourceAttributes } from '@opentelemetry/semantic-conventions'
import { OTLPTraceExporter } from '@opentelemetry/exporter-trace-otlp-http'
import { OTLPMetricExporter } from '@opentelemetry/exporter-metrics-otlp-http'
import { metrics } from '@opentelemetry/api';

let instrumentationSDK: NodeSDK

export function initializeInstrumentation(serviceName: string) {
  const exporterOptions = {
    url: 'https://ingest.{region}.signoz.cloud:443/v1/traces',
  }
  const resource = new Resource({
    [SemanticResourceAttributes.SERVICE_NAME]: serviceName,
  })
  const traceExporter = new OTLPTraceExporter(exporterOptions)
  const metricExporter = new OTLPMetricExporter(exporterOptions)
  const metricReader = new PeriodicExportingMetricReader({
    exporter: metricExporter,
  })
  const meterProvider = new MeterProvider({
    resource: resource,
    readers: [metricReader],
  })
  metrics.setGlobalMeterProvider(meterProvider);
  instrumentationSDK = new NodeSDK({
    traceExporter: traceExporter,
    metricReader: metricReader,
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
