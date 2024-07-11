import { WorkflowActivityJob } from '../generated/protobuf/bosca/workflow/execution_context_pb'

export class ActivityResult {
  readonly childWorkflowExecutionIds: string[]

  constructor(childWorkflowExecutionIds: string[] = []) {
    this.childWorkflowExecutionIds = childWorkflowExecutionIds
  }
}

export abstract class Activity {

  abstract get id(): string

  abstract execute(activity: WorkflowActivityJob): Promise<ActivityResult | void>
}

export type ActivityRegistry = { [activityId: string]: Activity }
export type ActivityQueueRegistry = { [queue: string]: ActivityRegistry }
