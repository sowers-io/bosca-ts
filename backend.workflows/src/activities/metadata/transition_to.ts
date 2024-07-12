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