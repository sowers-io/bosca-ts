import { metrics } from '@opentelemetry/api'

const meter = metrics.getMeter('workflow_workers')

export const workerCount = meter.createCounter('bosca_worker_count', {
  description: 'Counts total number of workers',
})

export const jobAddedCount = meter.createCounter('bosca_job_added_count', {
  description: 'Counts total number of jobs',
})

export const jobStartedCount = meter.createCounter('bosca_job_started_count', {
  description: 'Counts total number of started jobs',
})

export const jobFailedCount = meter.createCounter('bosca_job_failed_count', {
  description: 'Counts total number of failed jobs',
})

export const jobErrorCount = meter.createCounter('bosca_job_error_count', {
  description: 'Counts total number of error jobs',
})

export const jobFinishedCount = meter.createCounter('bosca_job_finished_count', {
  description: 'Counts total number of finished jobs',
})