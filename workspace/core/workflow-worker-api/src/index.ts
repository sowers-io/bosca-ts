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

import { logger } from '@bosca/common'
import { Activity, Job } from '@bosca/workflow-activities-api'
import { Configuration } from './configuration'
import { ConnectionOptions, QueueEvents, WaitingChildrenError, Worker } from 'bullmq'
import { WorkflowJob } from '@bosca/protobufs'
import { jobStartedCount, jobErrorCount, jobFinishedCount, jobAddedCount, jobFailedCount, workerCount } from './metrics'
import { ConnectError } from '@connectrpc/connect'

export { Configuration, QueueConfiguration, newConfiguration } from './configuration'

async function runJob(job: Job, definition: WorkflowJob, activities: { [id: string]: Activity }): Promise<any> {
  if (!definition.activity) return
  const activity = activities[definition.activity.activityId]
  if (!activity) {
    throw new Error('activity not found: ' + definition.activity.activityId)
  }
  const executor = activity.newJobExecutor(job, definition)
  try {
    const result = await executor.execute()
    logger.trace({ jobId: job.id, jobName: job.name }, 'finished job')
    return result
  } catch (e) {
    if (e instanceof WaitingChildrenError) {
      logger.trace({ jobId: job.id, jobName: job.name }, 'job waiting on children')
    } else {
      logger.error({ jobId: job.id, jobName: job.name, error: e }, 'error running job')
    }
    throw e
  }
}

export async function start(connection: ConnectionOptions, configuration: Configuration, activities: { [id: string]: Activity }) {
  const workers: Worker[] = []
  for (const queueConfigurationId in configuration.queues) {
    const queueConfiguration = configuration.queues[queueConfigurationId]
    const worker = new Worker(
      queueConfigurationId,
      async (job) => {
        logger.trace({ jobId: job.id, jobName: job.name }, 'running job')
        switch (job.data.type) {
          case 'job': {
            jobStartedCount.add(1)
            const definition = WorkflowJob.fromJson(job.data.job)
            return await runJob(job, definition, activities)
          }
          case 'workflow':
            break
        }
      },
      {
        connection: connection,
        concurrency: queueConfiguration.maxConcurrency,
        stalledInterval: 60000,
        lockDuration: 60000,
        maxStalledCount: 50,
      },
    )
    workers.push(worker)
    worker.on('completed', (job) => {
      if (job.data.type === 'job') {
        jobFinishedCount.add(1)
      }
      logger.trace({ jobId: job.id }, 'job completed')
    })
    worker.on('failed', (job, err) => {
      jobFailedCount.add(1)
      logger.error({ jobId: job?.id, error: err }, 'job failed')
    })
    const events = new QueueEvents(queueConfigurationId, { connection })
    events.on('added', async (job) => {
      jobAddedCount.add(1)
      logger.trace({ jobId: job.jobId, jobName: job.name }, 'job added')
    })
    events.on('error', async (error) => {
      if (error instanceof WaitingChildrenError) {
        logger.trace('job waiting on children')
      } else {
        jobErrorCount.add(1)
        logger.error({ error }, 'job error')
      }
    })
    workerCount.add(1)
    logger.info({ queue: queueConfigurationId, concurrency: queueConfiguration.maxConcurrency }, 'worker started')
  }
  logger.info('running...')
  const shutdown = async () => {
    workerCount.add(-workers.length)
    await Promise.all(workers.map((worker) => worker.close()))
    process.exit(0)
  }
  process
    .on('SIGTERM', shutdown)
    .on('SIGINT', shutdown)
    .on('unhandledRejection', (reason, p) => {
      logger.error({ reason, p }, 'unhandled rejection')
    })
    .on('uncaughtException', err => {
      logger.error({ error: err }, 'uncaught exception')
      if (err instanceof ConnectError) {
        return
      }
      process.exit(1)
    })
}