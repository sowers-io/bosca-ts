import { useServiceClient } from '../util/util'
import { WorkflowService } from '../generated/protobuf/bosca/workflow/service_connect'
import {
  WorkflowActivityJob,
  WorkflowActivityJobRequest
} from '../generated/protobuf/bosca/workflow/execution_context_pb'

export abstract class Activity {

  abstract get id(): string
  abstract execute(activity: WorkflowActivityJob): Promise<void>
}

type ActivityRegistry = { [activityId: string]: Activity }
type ActivityQueueRegistry = { [queue: string]: ActivityRegistry }

export class Workflow {

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
      await activity.execute(activityJob)
    }
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
    await Promise.any(listeners)
  }
}