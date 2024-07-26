import {
  Empty,
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
} from '@bosca/protobufs'
import { Code, ConnectError, type ConnectRouter } from '@connectrpc/connect'
import { PermissionManager, SubjectKey } from '@bosca/common'
import { WorkflowDataSource } from '../datasources/workflow'
import { ContentDataSource } from '../datasources/content'
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
  contentDataSource: ContentDataSource
): ConnectRouter {
  return router.service(WorkflowService, {
    async getModels() {
      return new Models({ models: await dataSource.getModels() })
    },
    async getModel(request) {
      const model = await dataSource.getModel(request.id)
      if (!model) {
        throw new ConnectError('missing model', Code.NotFound)
      }
      return model
    },
    async getPrompts() {
      return new Prompts({ prompts: await dataSource.getPrompts() })
    },
    async getPrompt(request) {
      const model = await dataSource.getModel(request.id)
      if (!model) {
        throw new ConnectError('missing model', Code.NotFound)
      }
      return model
    },
    async getStorageSystems() {
      return new StorageSystems({ systems: await dataSource.getStorageSystems() })
    },
    async getStorageSystem(request) {
      const system = await dataSource.getStorageSystem(request.id)
      if (!system) {
        throw new ConnectError('missing storage system', Code.NotFound)
      }
      return system
    },
    async getStorageSystemModels(request) {
      return new StorageSystemModels({
        models: await dataSource.getStorageSystemModels(request.id),
      })
    },
    async getWorkflows() {
      return new Workflows({ workflows: await dataSource.getWorkflows() })
    },
    async getWorkflow(request) {
      const workflow = await dataSource.getWorkflow(request.id)
      if (!workflow) {
        throw new ConnectError('missing workflow', Code.NotFound)
      }
      return workflow
    },
    async getWorkflowStates() {
      return new WorkflowStates({ states: await dataSource.getWorkflowStates() })
    },
    async getWorkflowState(request) {
      const state = await dataSource.getWorkflowState(request.id)
      if (!state) {
        throw new ConnectError('missing workflow state', Code.NotFound)
      }
      return state
    },
    async getWorkflowActivities(request) {
      return new WorkflowActivities({
        activities: await dataSource.getWorkflowActivities(request.id),
      })
    },
    async getWorkflowActivityStorageSystems(request) {
      return new WorkflowActivityStorageSystems({
        systems: await dataSource.getWorkflowActivityStorageSystems(Number(request.workflowActivityId)),
      })
    },
    async getWorkflowActivityPrompts(request) {
      return new WorkflowActivityPrompts({
        prompts: await dataSource.getWorkflowActivityPrompts(Number(request.workflowActivityId)),
      })
    },
    async beginTransitionWorkflow(request, context) {
      const subject = context.values.get(SubjectKey)
      const metadata = await contentDataSource.getMetadata(request.metadataId)
      if (!metadata) throw new ConnectError('missing metadata', Code.NotFound)
      if (metadata.workflowStateId === request.stateId) {
        throw new ConnectError('workflow already in state', Code.FailedPrecondition)
      }
      await verifyTransitionExists(dataSource, metadata, request.stateId)
      await verifyExitTransitionExecution(dataSource, metadata, request.stateId)
      const nextState = await verifyEnterTransitionExecution(dataSource, metadata, request.stateId)
      if (!nextState) {
        throw new ConnectError('missing next state', Code.NotFound)
      }
      await transition(subject, contentDataSource, dataSource, metadata, nextState, request.stateId, false)
      return new Empty()
    },
    async completeTransitionWorkflow(request, context) {
      const subject = context.values.get(SubjectKey)
      const metadata = await contentDataSource.getMetadata(request.metadataId)
      if (!metadata) throw new ConnectError('missing metadata', Code.NotFound)
      await completeTransitionWorkflow(subject, contentDataSource, metadata, request.status)
    },
    async executeWorkflow(request) {
      return await executeWorkflow(
        dataSource,
        request.parent,
        request.metadataId,
        request.collectionId,
        request.workflowId,
        request.context,
        request.waitForCompletion
      )
    },
    async findAndExecuteWorkflow() {
      throw new Error('unimplemented')
    },
  })
}
