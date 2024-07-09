import { WorkflowActivityJob } from '../generated/protobuf/bosca/workflow/execution_context_pb'

export abstract class Activity {

  abstract get id(): string
  abstract execute(activity: WorkflowActivityJob): Promise<void>
}

export type ActivityRegistry = { [activityId: string]: Activity }
export type ActivityQueueRegistry = { [queue: string]: ActivityRegistry }
