import { Activity } from '../activity'
import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import { ContentService } from '../../generated/protobuf/bosca/content/service_connect'
import { useServiceClient } from '../../util/util'
import {
  SetWorkflowStateCompleteRequest,
  SetWorkflowStateRequest
} from '../../generated/protobuf/bosca/content/workflows_pb'


export class TransitionToActivity extends Activity {

  get id(): string {
    return 'metadata.transition.to'
  }

  async execute(activity: WorkflowActivityJob) {
    const contentService = useServiceClient(ContentService)
    const state = activity.activity?.configuration['state']
    const status = activity.activity?.configuration['status']
    await contentService.setWorkflowStateComplete(new SetWorkflowStateCompleteRequest({
      metadataId: activity.metadataId,
      status: status
    }))
    await contentService.setWorkflowState(new SetWorkflowStateRequest({
      metadataId: activity.metadataId,
      stateId: state,
      status: status,
      immediate: true
    }))
  }
}