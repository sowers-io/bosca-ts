import { metrics } from '@opentelemetry/api';

const meter = metrics.getMeter('workflow_workers')

export const workerCount = meter.createCounter("worker_count", {
  description: "Counts total number of workers",
});

export const jobCount = meter.createCounter("job_count", {
  description: "Counts total number of jobs"
});