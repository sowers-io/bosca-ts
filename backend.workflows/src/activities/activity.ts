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
