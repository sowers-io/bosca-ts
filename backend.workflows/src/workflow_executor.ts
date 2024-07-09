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
        await this.updateJobStatus(activityJob, true, true)
      } catch (e) {
        await this.updateJobStatus(activityJob, true, false, e)
      }
    }
  }

  async updateJobStatus(activityJob: WorkflowActivityJob, complete: boolean, success: boolean, error?: any): Promise<void> {
    // TODO: keep trying to update the status, in case of network failures
    const service = useServiceClient(WorkflowService)
    await service.setWorkflowActivityJobStatus(new WorkflowActivityJobStatus({
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