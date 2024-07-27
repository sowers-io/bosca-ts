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

import { Activity, ActivityJobExecutor } from '../../activity'
import { Retry } from '../../../util/retry'
import {
  ContentService,
  Empty,
  IdRequest,
  Trait,
  WorkflowExecutionRequest,
  WorkflowJob,
  WorkflowParentJobId,
  WorkflowService,
} from '@bosca/protobufs'
import { Job } from 'bullmq/dist/esm/classes/job'
import { useServiceAccountClient } from '@bosca/common'
import { WaitingChildrenError } from 'bullmq'

export class ProcessTraitsActivity extends Activity {
  get id(): string {
    return 'metadata.traits.process'
  }

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, definition)
  }
}

class Executor extends ActivityJobExecutor<ProcessTraitsActivity> {
  private async executeWorkflow(workflowId: string) {
    await Retry.execute(100, async () => {
      const workflowService = useServiceAccountClient(WorkflowService)
      await workflowService.executeWorkflow(
        new WorkflowExecutionRequest({
          parent: new WorkflowParentJobId({
            id: this.job.id,
            queue: this.job.queueQualifiedName,
          }),
          workflowId: workflowId,
          metadataId: this.definition.metadataId,
        })
      )
    })
  }

  async execute() {
    const contentService = useServiceAccountClient(ContentService)
    const metadata = await Retry.execute(
      10,
      async () => await contentService.getMetadata(new IdRequest({ id: this.definition.metadataId }))
    )
    if (!metadata.traitIds || metadata.traitIds.length === 0) return

    const traits = await contentService.getTraits(new Empty())
    const traitsById: { [id: string]: Trait } = {}
    for (const trait of traits.traits) {
      traitsById[trait.id] = trait
    }

    if (!this.job.data.hasOwnProperty('executed')) {
      this.job.data.executed = []
    }
    let executed = false
    for (const traitId of metadata.traitIds) {
      const trait = traitsById[traitId]
      if (!trait.workflowIds || trait.workflowIds.length === 0) continue
      for (const workflowId of trait.workflowIds) {
        const key = traitId + '-' + workflowId
        if (this.job.data.executed.includes(key)) continue
        await this.executeWorkflow(workflowId)
        executed = true
        this.job.data.executed.push(key)
        await this.job.updateData(this.job.data)
      }
    }

    if (executed && (await this.job.moveToWaitingChildren(this.job.token!))) {
      throw new WaitingChildrenError()
    }
  }
}
