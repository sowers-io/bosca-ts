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
  WorkflowActivityJobRequest, WorkflowActivityJobStatus
} from './generated/protobuf/bosca/workflow/execution_context_pb'
import { Activity, ActivityQueueRegistry } from './activities/activity'

export class WorkflowExecutor {

  private registry: ActivityQueueRegistry = {}

  register(queue: string, activity: Activity) {
    let activities = this.registry[queue]
    if (!activities) {
      activities = {}
      this.registry[queue] = activities
    }
    activities[activity.id] = activity
  }

  async listen(queue: string, activityIds: string[]): Promise<void> {
    const service = useServiceClient(WorkflowService)
    const listenRequest = new WorkflowActivityJobRequest({
      queue: queue,
      activityId: activityIds
    })
    for await (const activityJob of service.getWorkflowActivityJobs(listenRequest)) {
      const activity = this.registry[queue][activityJob.activity!.activityId]
      try {
        await activity.execute(activityJob)
        await this.updateJobStatus(activityJob, queue, true, true)
        console.log('job finished', activityJob)
      } catch (e) {
        console.error('failed to execute activity', activityJob, e)
        await this.updateJobStatus(activityJob, queue, false, false, e)
      }
    }
  }

  async updateJobStatus(activityJob: WorkflowActivityJob, queue: string, complete: boolean, success: boolean, error?: any): Promise<void> {
    // TODO: keep trying to update the status, in case of network failures
    const service = useServiceClient(WorkflowService)
    await service.setWorkflowActivityJobStatus(new WorkflowActivityJobStatus({
      jobId: activityJob.jobId,
      executionId: activityJob.executionId,
      workflowActivityId: activityJob.activity?.workflowActivityId,
      complete: complete,
      success: success,
      error: error?.toString()
    }))
  }

  async execute() {
    const listeners: Promise<void>[] = []
    for (const queue in this.registry) {
      const activityIds: string[] = []
      for (const activityId in this.registry[queue]) {
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