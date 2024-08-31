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
  Empty,
  IdRequest,
  Models,
  Prompts,
  StorageSystemModels,
  StorageSystems,
  WorkflowActivities,
  WorkflowActivityPrompts,
  WorkflowActivityStorageSystems,
  Workflows,
  WorkflowService,
  WorkflowStates,
  FindCollectionRequest,
  FindMetadataRequest,
  WorkflowEnqueueResponses, PermissionObjectType, PermissionAction, WorkflowActivityModels,
} from '@bosca/protobufs'
import { logger, PermissionManager, SubjectKey, SubjectType, useServiceAccountClient } from '@bosca/common'
import { Code, ConnectError, type ConnectRouter } from '@connectrpc/connect'
import { WorkflowDataSource } from '../datasources/workflow'
import {
  completeTransitionWorkflow,
  executeWorkflow,
  transition,
  verifyEnterTransitionExecution,
  verifyExitTransitionExecution,
  verifyTransitionExists,
} from './util/workflows'

export function workflow(
  router: ConnectRouter,
  permissions: PermissionManager,
  dataSource: WorkflowDataSource,
): ConnectRouter {
  return router.service(WorkflowService, {
    async getModels(_, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      return new Models({ models: await dataSource.getModels() })
    },
    async getModel(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      const model = await dataSource.getModel(request.id)
      if (!model) {
        throw new ConnectError('missing model', Code.NotFound)
      }
      return model
    },
    async getPrompts(_, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      return new Prompts({ prompts: await dataSource.getPrompts() })
    },
    async getPrompt(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      const model = await dataSource.getModel(request.id)
      if (!model) {
        throw new ConnectError('missing model', Code.NotFound)
      }
      return model
    },
    async getStorageSystems(_, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      return new StorageSystems({ systems: await dataSource.getStorageSystems() })
    },
    async getStorageSystem(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      const system = await dataSource.getStorageSystem(request.id)
      if (!system) {
        throw new ConnectError('missing storage system', Code.NotFound)
      }
      return system
    },
    async getStorageSystemModels(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      return new StorageSystemModels({
        models: await dataSource.getStorageSystemModels(request.id),
      })
    },
    async getWorkflows(_, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      return new Workflows({ workflows: await dataSource.getWorkflows() })
    },
    async getWorkflow(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      const workflow = await dataSource.getWorkflow(request.id)
      if (!workflow) {
        logger.warn({ id: request.id }, 'missing workflow')
        throw new ConnectError('missing workflow', Code.NotFound)
      }
      return workflow
    },
    async getWorkflowStates(_, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      return new WorkflowStates({ states: await dataSource.getWorkflowStates() })
    },
    async getWorkflowState(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      const state = await dataSource.getWorkflowState(request.id)
      if (!state) {
        logger.warn({ id: request.id }, 'missing workflow state')
        throw new ConnectError('missing workflow state', Code.NotFound)
      }
      return state
    },
    async getWorkflowActivities(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      return new WorkflowActivities({
        activities: await dataSource.getWorkflowActivities(request.id),
      })
    },
    async getWorkflowActivityStorageSystems(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      return new WorkflowActivityStorageSystems({
        systems: await dataSource.getWorkflowActivityStorageSystems(Number(request.workflowActivityId)),
      })
    },
    async getWorkflowActivityPrompts(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      return new WorkflowActivityPrompts({
        prompts: await dataSource.getWorkflowActivityPrompts(Number(request.workflowActivityId)),
      })
    },
    async getWorkflowActivityModels(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.view,
      )
      return new WorkflowActivityModels({
        models: await dataSource.getWorkflowActivityModels(Number(request.workflowActivityId)),
      })
    },
    async beginTransitionWorkflow(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.metadataId,
        PermissionAction.manage,
      )
      const metadata = await useServiceAccountClient(ContentService).getMetadata(
        new IdRequest({ id: request.metadataId }),
      )
      if (!metadata) throw new ConnectError('missing metadata', Code.NotFound)
      if (metadata.workflowStateId === request.stateId) {
        throw new ConnectError('workflow already in state', Code.FailedPrecondition)
      }
      const ctx: { [key: string]: string } = { subject: subject.id, subjectType: subject.type }
      await verifyTransitionExists(dataSource, metadata, request.stateId)
      await verifyExitTransitionExecution(dataSource, metadata, request.stateId, ctx)
      const nextState = await verifyEnterTransitionExecution(dataSource, metadata, request.stateId, ctx)
      if (!nextState) {
        throw new ConnectError('missing next state', Code.NotFound)
      }
      await transition(dataSource, metadata, nextState, request.stateId, request.waitForCompletion, ctx)
      return new Empty()
    },
    async completeTransitionWorkflow(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.workflow_type,
        'all',
        PermissionAction.manage,
      )
      const metadata = await useServiceAccountClient(ContentService).getMetadata(
        new IdRequest({ id: request.metadataId }),
      )
      if (!metadata) throw new ConnectError('missing metadata', Code.NotFound)
      await completeTransitionWorkflow(metadata, request.status)
    },
    async executeWorkflow(request, context) {
      const subject = context.values.get(SubjectKey)
      if (subject.type != SubjectType.serviceaccount) {
        throw new ConnectError('unauthorized', Code.PermissionDenied)
      }
      return await executeWorkflow(
        dataSource,
        request.parent,
        request.metadataId,
        request.supplementaryId,
        request.collectionId,
        request.workflowId,
        request.context,
        request.waitForCompletion,
      )
    },
    async findAndExecuteWorkflow(request, context) {
      const subject = context.values.get(SubjectKey)
      if (subject.type != SubjectType.serviceaccount) {
        throw new ConnectError('unauthorized', Code.PermissionDenied)
      }
      const service = useServiceAccountClient(ContentService)
      const responses = new WorkflowEnqueueResponses({ responses: [] })
      const requests = []
      if (request.metadataId) {
        requests.push({ collection: null, metadata: request.metadataId })
      }
      if (request.collectionId) {
        requests.push({ collection: request.collectionId, metadata: null })
      }
      if (request.collectionAttributes) {
        const collections = await service.findCollection(
          new FindCollectionRequest({ attributes: request.collectionAttributes }),
        )
        for (const collection of collections.collections) {
          requests.push({ collection: collection.id, metadata: null })
        }
      }
      if (request.metadataAttributes) {
        const metadatas = await service.findMetadata(
          new FindMetadataRequest({ attributes: request.metadataAttributes }),
        )
        for (const metadata of metadatas.metadata) {
          requests.push({ collection: null, metadata: metadata.id })
        }
      }
      for (const r of requests) {
        const response = await executeWorkflow(
          dataSource,
          null,
          r.metadata,
          null,
          r.collection,
          request.workflowId,
          request.context,
          request.waitForCompletion,
        )
        responses.responses.push(response)
      }
      return responses
    },
  })
}
