import { WorkflowActivityExecutionContext } from '../generated/protobuf/bosca/content/workflows_pb'

export abstract class Activity {

  abstract get id(): string

  abstract execute(executionContext: WorkflowActivityExecutionContext): Promise<void>
}

export class Workflow {

  private activities: { [id: string]: Activity } = {}

  register(activity: Activity) {
    this.activities[activity.id] = activity
  }

  async execute(executionContext: WorkflowActivityExecutionContext) {
    while (executionContext.currentActivityIndex < executionContext.activities.length) {
      const activityDefinition = executionContext.activities[executionContext.currentActivityIndex]
      executionContext.currentActivityIndex++

      const activity = this.activities[activityDefinition.id]
      await activity.execute(executionContext)
    }
  }
}