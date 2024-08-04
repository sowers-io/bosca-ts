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

import { Activity, ActivityJobExecutor } from '../activity'

import { WorkflowExecutionRequest, WorkflowJob, WorkflowParentJobId, WorkflowService } from '@bosca/protobufs'
import { Job, WaitingChildrenError } from 'bullmq'
import { useServiceAccountClient } from '@bosca/common'

export class ChildWorkflow extends Activity {
  get id(): string {
    return 'metadata.child.workflow'
  }

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, definition)
  }
}

class Executor extends ActivityJobExecutor<ChildWorkflow> {
  async execute() {
    // eslint-disable-next-line no-prototype-builtins
    if (this.job.data.hasOwnProperty('executed') && this.job.data.executed) {
      return
    }

    const workflowId = this.definition.activity?.configuration['workflowId']
    const workflowService = useServiceAccountClient(WorkflowService)

    await workflowService.executeWorkflow(
      new WorkflowExecutionRequest({
        parent: new WorkflowParentJobId({
          id: this.job.id,
          queue: this.job.queueQualifiedName,
        }),
        workflowId: workflowId,
        metadataId: this.definition.metadataId,
      }),
    )

    this.job.data.executed = true
    await this.job.updateData(this.job.data)

    if (await this.job.moveToWaitingChildren(this.job.token!)) {
      throw new WaitingChildrenError()
    }
  }
}
