import {
  Metadata,
  WorkflowEnqueueRequest,
  WorkflowEnqueueResponse,
  WorkflowJob,
  WorkflowParentJobId,
  WorkflowQueueService,
  WorkflowState,
} from '@bosca/protobufs'
import { WorkflowDataSource } from '../../datasources/workflow'
import { Code, ConnectError } from '@connectrpc/connect'
import { ContentDataSource } from '../../datasources/content'
import { Subject, useServiceAccountClient } from '@bosca/common'
import { setWorkflowStateComplete } from './metadata'

export const StatePending = 'pending'
export const StateProcessing = 'processing'
export const StateDraft = 'draft'
export const StatePublished = 'published'
export const StateFailed = 'failed'

export async function executeWorkflow(
  workflowDataSource: WorkflowDataSource,
  parent: WorkflowParentJobId | undefined | null,
  metadataId: string | undefined | null,
  collectionId: string | undefined | null,
  workflowId: string,
  context: { [key: string]: string } | undefined | null,
  waitForCompletion: boolean
): Promise<WorkflowEnqueueResponse> {
  const workflow = await workflowDataSource.getWorkflow(workflowId)
  if (!workflow) throw new ConnectError('missing workflow', Code.NotFound)
  const jobs: WorkflowJob[] = []
  const activities = await workflowDataSource.getWorkflowActivities(workflowId)
  for (const activity of activities) {
    const prompts = await workflowDataSource.getWorkflowActivityPrompts(Number(activity.workflowActivityId))
    const storageSystems = await workflowDataSource.getWorkflowActivityStorageSystems(
      Number(activity.workflowActivityId)
    )
    const models = await workflowDataSource.getWorkflowActivityModels(Number(activity.workflowActivityId))
    jobs.push(
      new WorkflowJob({
        workflowId: workflow.id,
        metadataId: metadataId || undefined,
        activity: activity,
        prompts: prompts,
        models: models,
        storageSystems: storageSystems,
        context: context || undefined,
      })
    )
  }
  return useServiceAccountClient(WorkflowQueueService).enqueue(
    new WorkflowEnqueueRequest({
      parent: parent || undefined,
      metadataId: metadataId || undefined,
      collectionId: collectionId || undefined,
      workflow: workflow,
      jobs: jobs,
      context: context || undefined,
      waitForCompletion: waitForCompletion,
    })
  )
}

export async function verifyTransitionExists(
  dataSource: WorkflowDataSource,
  metadata: Metadata,
  nextStateId: string
): Promise<void> {
  const transition = await dataSource.getWorkflowTransition(metadata.workflowStateId, nextStateId)
  if (!transition) {
    throw new ConnectError('missing transition', Code.NotFound)
  }
}

export async function verifyExitTransitionExecution(
  dataSource: WorkflowDataSource,
  metadata: Metadata,
  nextStateId: string
): Promise<void> {
  const nextState = await dataSource.getWorkflowState(nextStateId)
  if (!nextState) {
    throw new ConnectError('missing next state', Code.NotFound)
  }
  if (nextState.exitWorkflowId) {
    await executeWorkflow(dataSource, null, metadata.id, null, nextState.exitWorkflowId, null, true)
  }
}

export async function verifyEnterTransitionExecution(
  dataSource: WorkflowDataSource,
  metadata: Metadata,
  nextStateId: string
): Promise<WorkflowState | null> {
  const nextState = await dataSource.getWorkflowState(nextStateId)
  if (!nextState) {
    throw new ConnectError('missing next state', Code.NotFound)
  }
  if (nextState.entryWorkflowId) {
    await executeWorkflow(dataSource, null, metadata.id, null, nextState.entryWorkflowId, null, true)
  }
  return nextState
}

export async function transition(
  subject: Subject,
  contentDataSource: ContentDataSource,
  workflowDataSource: WorkflowDataSource,
  metadata: Metadata,
  nextState: WorkflowState,
  status: string,
  waitForCompletion: boolean
): Promise<void> {
  if (nextState.workflowId) {
    await contentDataSource.setWorkflowState(
      subject,
      metadata.id,
      metadata.workflowStateId,
      nextState.id,
      status,
      true,
      false
    )
    await executeWorkflow(workflowDataSource, null, metadata.id, null, nextState.workflowId, null, waitForCompletion)
  } else {
    await setWorkflowStateComplete(subject, contentDataSource, metadata.id, status)
  }
}

export async function completeTransitionWorkflow(
  subject: Subject,
  contentDataSource: ContentDataSource,
  metadata: Metadata,
  status: string
): Promise<void> {
  if (!metadata.workflowStatePendingId) {
    throw new ConnectError('no pending state', Code.FailedPrecondition)
  }
  await setWorkflowStateComplete(subject, contentDataSource, metadata.id, status)
}
