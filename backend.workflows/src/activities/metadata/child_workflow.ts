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
import {
  WorkflowActivityJob,
  WorkflowExecutionRequest
} from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import { ContentService } from '../../generated/protobuf/bosca/content/service_connect'
import { useServiceClient } from '../../util/util'
import {
  SetWorkflowStateCompleteRequest,
  SetWorkflowStateRequest
} from '../../generated/protobuf/bosca/content/workflows_pb'
import { Retry } from '../../util/retry'
import { WorkflowService } from '../../generated/protobuf/bosca/workflow/service_connect'


export class ChildWorkflow extends Activity {

  get id(): string {
    return 'metadata.child.workflow'
  }

  async execute(activity: WorkflowActivityJob) {
    const workflowId = activity.activity?.configuration['workflowId']
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
}