import { Activity } from '../../activity'
import {
  WorkflowActivityJob, WorkflowExecutionRequest
} from '../../../generated/protobuf/bosca/workflow/execution_context_pb'
import { useServiceClient } from '../../../util/util'
import { ContentService } from '../../../generated/protobuf/bosca/content/service_connect'
import { Empty } from '../../../generated/protobuf/bosca/empty_pb'
import { IdRequest } from '../../../generated/protobuf/bosca/requests_pb'
import { Trait } from '../../../generated/protobuf/bosca/content/traits_pb'
import { WorkflowService } from '../../../generated/protobuf/bosca/workflow/service_connect'

export class ProcessTraitsActivity extends Activity {

  get id(): string {
    return 'metadata.traits.process'
  }

  async execute(activity: WorkflowActivityJob) {
    const contentService = useServiceClient(ContentService)
    const workflowService = useServiceClient(WorkflowService)
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
        await workflowService.executeWorkflow(new WorkflowExecutionRequest({
          parentExecutionId: activity.executionId,
          workflowId: workflowId,
          metadataId: activity.metadataId,
        }))
      }
    }
  }
}