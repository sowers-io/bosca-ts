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

import {
  ContentService,
  Metadata,
  SetWorkflowStateCompleteRequest,
  SetWorkflowStateRequest,
  WorkflowEnqueueRequest,
  WorkflowEnqueueResponse,
  WorkflowJob,
  WorkflowParentJobId,
  WorkflowQueueService,
  WorkflowState,
} from '@bosca/protobufs'
import { WorkflowDataSource } from '../../datasources/workflow'
import { Code, ConnectError } from '@connectrpc/connect'
import { useServiceAccountClient } from '@bosca/common'

export async function executeWorkflow(
  workflowDataSource: WorkflowDataSource,
  parent: WorkflowParentJobId | undefined | null,
  metadataId: string | undefined | null,
  supplementaryId: string | undefined | null,
  collectionId: string | undefined | null,
  workflowId: string,
  context: { [key: string]: string } | undefined | null,
  waitForCompletion: boolean,
): Promise<WorkflowEnqueueResponse> {
  const workflow = await workflowDataSource.getWorkflow(workflowId)
  if (!workflow) throw new ConnectError('missing workflow', Code.NotFound)
  const jobs: WorkflowJob[] = []
  const activities = await workflowDataSource.getWorkflowActivities(workflowId)
  for (const activity of activities) {
    const prompts = await workflowDataSource.getWorkflowActivityPrompts(Number(activity.workflowActivityId))
    const storageSystems = await workflowDataSource.getWorkflowActivityStorageSystems(
      Number(activity.workflowActivityId),
    )
    const models = await workflowDataSource.getWorkflowActivityModels(Number(activity.workflowActivityId))
    jobs.push(
      new WorkflowJob({
        workflowId: workflow.id,
        metadataId: metadataId || undefined,
        supplementaryId: supplementaryId || undefined,
        activity: activity,
        prompts: prompts,
        models: models,
        storageSystems: storageSystems,
        context: context || undefined,
      }),
    )
  }
  return useServiceAccountClient(WorkflowQueueService).enqueue(
    new WorkflowEnqueueRequest({
      parent: parent || undefined,
      metadataId: metadataId || undefined,
      supplementaryId: supplementaryId || undefined,
      collectionId: collectionId || undefined,
      workflow: workflow,
      jobs: jobs,
      context: context || undefined,
      waitForCompletion: waitForCompletion,
    }),
  )
}

export async function verifyTransitionExists(
  dataSource: WorkflowDataSource,
  metadata: Metadata,
  nextStateId: string,
): Promise<void> {
  const transition = await dataSource.getWorkflowTransition(metadata.workflowStateId, nextStateId)
  if (!transition) {
    throw new ConnectError('missing transition', Code.NotFound)
  }
}

export async function verifyExitTransitionExecution(
  dataSource: WorkflowDataSource,
  metadata: Metadata,
  nextStateId: string,
  context: { [key: string]: string },
): Promise<void> {
  const nextState = await dataSource.getWorkflowState(nextStateId)
  if (!nextState) {
    throw new ConnectError('missing next state', Code.NotFound)
  }
  if (nextState.exitWorkflowId) {
    await executeWorkflow(dataSource, null, metadata.id, null, null, nextState.exitWorkflowId, context, true)
  }
}

export async function verifyEnterTransitionExecution(
  dataSource: WorkflowDataSource,
  metadata: Metadata,
  nextStateId: string,
  context: { [key: string]: string },
): Promise<WorkflowState | null> {
  const nextState = await dataSource.getWorkflowState(nextStateId)
  if (!nextState) {
    throw new ConnectError('missing next state', Code.NotFound)
  }
  if (nextState.entryWorkflowId) {
    await executeWorkflow(dataSource, null, metadata.id, null, null, nextState.entryWorkflowId, context, true)
  }
  return nextState
}

export async function transition(
  workflowDataSource: WorkflowDataSource,
  metadata: Metadata,
  nextState: WorkflowState,
  status: string,
  waitForCompletion: boolean,
  context: { [key: string]: string },
): Promise<void> {
  if (nextState.workflowId) {
    await useServiceAccountClient(ContentService).setWorkflowState(
      new SetWorkflowStateRequest({
        metadataId: metadata.id,
        status: status,
        stateId: nextState.id,
      }),
    )
    await executeWorkflow(workflowDataSource, null, metadata.id, null, null, nextState.workflowId, context, waitForCompletion)
  } else {
    await useServiceAccountClient(ContentService).setWorkflowStateComplete(
      new SetWorkflowStateCompleteRequest({
        metadataId: metadata.id,
        status: status,
      }),
    )
  }
}

export async function completeTransitionWorkflow(metadata: Metadata, status: string): Promise<void> {
  if (!metadata.workflowStatePendingId) {
    throw new ConnectError('no pending state', Code.FailedPrecondition)
  }
  await useServiceAccountClient(ContentService).setWorkflowStateComplete(
    new SetWorkflowStateCompleteRequest({
      metadataId: metadata.id,
      status: status,
    }),
  )
}
