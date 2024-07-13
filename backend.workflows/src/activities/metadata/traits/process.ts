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

import { Activity } from '../../activity'
import {
  WorkflowActivityJob,
  WorkflowExecutionRequest,
} from '../../../generated/protobuf/bosca/workflow/execution_context_pb'
import { useServiceClient } from '../../../util/util'
import { ContentService } from '../../../generated/protobuf/bosca/content/service_connect'
import { Empty } from '../../../generated/protobuf/bosca/empty_pb'
import { IdRequest } from '../../../generated/protobuf/bosca/requests_pb'
import { Trait } from '../../../generated/protobuf/bosca/content/traits_pb'
import { WorkflowService } from '../../../generated/protobuf/bosca/workflow/service_connect'
import { Retry } from '../../../util/retry'

export class ProcessTraitsActivity extends Activity {
  get id(): string {
    return 'metadata.traits.process'
  }

  private async executeWorkflow(workflowId: string, activity: WorkflowActivityJob) {
    await Retry.execute(100, async () => {
      const workflowService = useServiceClient(WorkflowService)
      await workflowService.executeWorkflow(
        new WorkflowExecutionRequest({
          parentExecutionId: activity.executionId,
          workflowId: workflowId,
          metadataId: activity.metadataId,
        })
      )
    })
  }

  async execute(activity: WorkflowActivityJob) {
    const contentService = useServiceClient(ContentService)
    const metadata = await contentService.getMetadata(new IdRequest({ id: activity.metadataId }))
    if (!metadata.traitIds || metadata.traitIds.length === 0) return

    const traits = await contentService.getTraits(new Empty())
    const traitsById: { [id: string]: Trait } = {}
    for (const trait of traits.traits) {
      traitsById[trait.id] = trait
    }

    for (const traitId of metadata.traitIds) {
      const trait = traitsById[traitId]
      if (!trait.workflowIds || trait.workflowIds.length === 0) continue

      for (const workflowId of trait.workflowIds) {
        await this.executeWorkflow(workflowId, activity)
      }
    }
  }
}
