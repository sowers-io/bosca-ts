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
  WorkflowActivityJob,
  WorkflowActivityJobRequest,
  WorkflowActivityJobStatus,
} from './generated/protobuf/bosca/workflow/execution_context_pb'
import { Activity, ActivityQueueRegistry } from './activities/activity'
import { Queue } from './util/queue'
import { Retry } from './util/retry'

export class Executor {
  private activityRegistry: ActivityQueueRegistry = {}
  private queueRegistry: { [key: string]: Queue } = {}

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
    console.log('registering activity:', activity.id, ':: in queue:', queue)
  }

  async listen(queue: string, activityIds: string[]): Promise<void> {
    const service = useServiceClient(WorkflowService)
    const listenRequest = new WorkflowActivityJobRequest({
      queue: queue,
      activityId: activityIds,
    })
    const activityQueue = this.queueRegistry[queue]
    for await (const activityJob of service.getWorkflowActivityJobs(listenRequest)) {
      console.log('received job: ', activityJob.jobId, ':: for activity: ', activityJob.activity?.activityId, ':: in queue: ', queue)
      const activity = this.activityRegistry[queue][activityJob.activity!.activityId]
      await this.executeActivity(activityQueue, activity, activityJob)
    }
  }

  private async executeActivity(queue: Queue, activity: Activity, job: WorkflowActivityJob) {
    await queue.enqueue(async () => {
      try {
        console.log('executing job: ', job.jobId, ':: for activity: ', job.activity?.activityId, ':: in queue: ', queue.name)
        await activity.execute(job)
        await this.updateJobStatus(job, true, true)
        console.log('finished job:  ', job.jobId, ':: for activity: ', job.activity?.activityId, ':: in queue: ', queue.name)
      } catch (e) {
        console.error('failed to execute activity', job, e)
        await this.updateJobStatus(job, false, false, e)
      }
    })
  }

  async updateJobStatus(
    activityJob: WorkflowActivityJob,
    complete: boolean,
    success: boolean,
    error?: any
  ): Promise<void> {
    await Retry.execute(1000, async () => {
      const service = useServiceClient(WorkflowService)
      await service.setWorkflowActivityJobStatus(
        new WorkflowActivityJobStatus({
          jobId: activityJob.jobId,
          executionId: activityJob.executionId,
          workflowActivityId: activityJob.activity?.workflowActivityId,
          complete: complete,
          success: success,
          error: error?.toString(),
        })
      )
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
