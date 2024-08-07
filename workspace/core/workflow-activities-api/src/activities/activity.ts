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

import { WorkflowJob } from '@bosca/protobufs'
import { Job } from 'bullmq'

export abstract class Activity {
  abstract get id(): string

  abstract newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any>
}

export abstract class ActivityJobExecutor<T extends Activity> {
  protected readonly activity: T
  protected readonly job: Job
  protected readonly definition: WorkflowJob

  constructor(activity: T, job: Job, definition: WorkflowJob) {
    this.activity = activity
    this.job = job
    this.definition = definition
  }

  abstract execute(): Promise<any>
}
