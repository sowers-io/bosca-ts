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

import { useServiceClient } from './util/util'
import { WorkflowService } from './generated/protobuf/bosca/workflow/service_connect'
import {
  RegisterWorkerRequest,
  WorkflowActivityJob, WorkflowActivityJobRequest,
  WorkflowActivityJobStatus,
  WorkflowExecutionNotification,
  WorkflowExecutionNotificationType
} from './generated/protobuf/bosca/workflow/execution_context_pb'
import { Activity, ActivityQueueRegistry } from './activities/activity'
import { Queue } from './util/queue'
import { Retry } from './util/retry'
import { createClient } from 'redis'
import { RedisClientType } from '@redis/client/dist/lib/client'

export class Executor {
  private activityRegistry: ActivityQueueRegistry = {}
  private queueRegistry: { [key: string]: Queue } = {}

  private listening = 0
  private executingActivities = 0
  private completedActivities = 0
  private receivedJobs = 0
  private totalActivityErrors = 0
  private jobStatusErrors = 0
  private activityExecutions: { [key: string]: number } = {}
  private activityCompletions: { [key: string]: number } = {}
  private activityErrors: { [key: string]: number } = {}
  private activityActive: { [key: string]: number } = {}
  private redis: RedisClientType

  constructor() {
    const executor = this
    setInterval(() => executor.printStats(), 10000)
    this.redis = createClient({
      url: process.env.BOSCA_REDIS_ADDRESS
    })
  }

  private printStats() {
    console.warn('===========================================================')
    console.warn(
      'listening:',
      this.listening,
      'executingActivities:',
      this.executingActivities,
      'completedActivities:',
      this.completedActivities,
      'receivedJobs:',
      this.receivedJobs,
      'totalActivityErrors:',
      this.totalActivityErrors,
      'jobStatusErrors:',
      this.jobStatusErrors,
      'retries:',
      Retry.retries
    )
    console.warn('-----------------------------------------------------------')
    for (const key in this.queueRegistry) {
      const queue = this.queueRegistry[key]
      console.warn(
        'queue:',
        queue.name,
        'queued:',
        queue.queued,
        'active:',
        queue.active,
        'max:',
        queue.maxActive,
        'waiting:',
        queue.waiting,
        'processed:',
        queue.processed
      )
    }
    console.warn('-----------------------------------------------------------')
    for (const key in this.activityExecutions) {
      console.warn(
        'activity:',
        key,
        'executions:',
        this.activityExecutions[key],
        'completions:',
        this.activityCompletions[key],
        'errors:',
        this.activityErrors[key],
        'active:',
        this.activityActive[key]
      )
    }
    console.warn('===========================================================')
  }

  registerQueue(queue: string, maxConcurrent: number) {
    this.queueRegistry[queue] = new Queue(queue, maxConcurrent)
    console.log('registering queue:', queue, ':: with maxConcurrent: ', maxConcurrent)
  }

  registerActivity(queue: string, activity: Activity) {
    let activities = this.activityRegistry[queue]
    if (!activities) {
      activities = {}
      this.activityRegistry[queue] = activities
    }
    activities[activity.id] = activity
    this.activityExecutions[activity.id] = 0
    this.activityActive[activity.id] = 0
    this.activityErrors[activity.id] = 0
    this.activityCompletions[activity.id] = 0
    console.log('registering activity:', activity.id, ':: in queue:', queue)
  }

  async listen(queue: string, activityIds: string[]): Promise<void> {
    try {
      this.listening++
      const service = useServiceClient(WorkflowService)
      const activityQueue = this.queueRegistry[queue]
      const supportedActivities: { [activityId: string]: boolean } = {}
      for (const activityId of activityIds) {
        supportedActivities[activityId] = true
      }

      const registration = await service.registerWorker(new RegisterWorkerRequest({
        queue: queue,
        activityId: activityIds
      }))

      const subscriber = this.redis.duplicate()
      await subscriber.connect()

      await new Promise((_, reject) => {
        subscriber.subscribe(queue, async (notificationJson) => {
          const notification = WorkflowExecutionNotification.fromJsonString(notificationJson)
          if (notification.type === WorkflowExecutionNotificationType.job_available) {
            try {
              const activityJob = await service.getWorkflowActivityJob(new WorkflowActivityJobRequest({
                activityId: activityIds,
                queue: queue,
                workerId: registration.id,
              }))

              if (!activityJob || !activityJob.activity) return

              this.receivedJobs++
              console.log(
                'received job: ',
                activityJob.jobId,
                ':: for activity: ',
                activityJob.activity?.activityId,
                ':: in queue: ',
                queue
              )
              const activity = this.activityRegistry[queue][activityJob.activity!.activityId]
              try {
                this.activityActive[activityJob.activity!.activityId] =
                  this.activityActive[activityJob.activity!.activityId] + 1
                await this.executeActivity(activityQueue, activity, activityJob)
              } finally {
                this.activityActive[activityJob.activity!.activityId] =
                  this.activityActive[activityJob.activity!.activityId] - 1
              }
            } catch (e) {
              console.error('error processing job', e)
            }
          }
        }).catch(reject)
      })
    } catch (e) {
      console.log('failed to subscribe and listen', e)
      throw e
    } finally {
      this.listening--
    }
  }

  private async executeActivity(queue: Queue, activity: Activity, job: WorkflowActivityJob) {
    const executor = this
    queue.enqueue(async () => {
      try {
        executor.executingActivities++
        executor.activityExecutions[job.activity!.activityId] =
          executor.activityExecutions[job.activity!.activityId] + 1
        console.log(
          'executing job: ',
          job.jobId,
          ':: for activity: ',
          job.activity?.activityId,
          ':: in queue: ',
          queue.name
        )
        await activity.execute(job)
        executor.completedActivities++
        executor.activityCompletions[job.activity!.activityId] =
          executor.activityCompletions[job.activity!.activityId] + 1
        await executor.updateJobStatus(job, true, true)
        console.log(
          'finished job:  ',
          job.jobId,
          ':: for activity: ',
          job.activity?.activityId,
          ':: in queue: ',
          queue.name
        )
      } catch (e) {
        console.error('failed to execute activity', job, e)
        await executor.updateJobStatus(job, false, false, e)
        executor.totalActivityErrors++
        executor.activityErrors[job.activity!.activityId] = executor.activityErrors[job.activity!.activityId] + 1
      } finally {
        executor.executingActivities--
      }
    })
    await queue.process()
  }

  async updateJobStatus(
    activityJob: WorkflowActivityJob,
    complete: boolean,
    success: boolean,
    error?: any
  ): Promise<void> {
    const executor = this
    await Retry.execute(1000, async () => {
      try {
        const service = useServiceClient(WorkflowService)
        await service.setWorkflowActivityJobStatus(
          new WorkflowActivityJobStatus({
            jobId: activityJob.jobId,
            executionId: activityJob.executionId,
            workflowActivityId: activityJob.activity?.workflowActivityId,
            complete: complete,
            success: success,
            error: error?.toString()
          })
        )
      } catch (e) {
        executor.jobStatusErrors++
        console.error('failed to update job status', activityJob, e)
        throw e
      }
    })
  }

  async execute() {
    const listeners: Promise<void>[] = []
    for (const queue in this.activityRegistry) {
      const activityIds: string[] = []
      for (const activityId in this.activityRegistry[queue]) {
        activityIds.push(activityId)
      }
      listeners.push(this.listen(queue, activityIds))
    }
    try {
      await Promise.any(listeners)
    } catch (e) {
      if (e instanceof AggregateError) {
        for (const err of e.errors) {
          console.error(err)
        }
      }
      console.error(e)
      process.exit(1)
    }
  }
}
