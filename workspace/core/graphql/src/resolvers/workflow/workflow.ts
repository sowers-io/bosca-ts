import {
  Resolvers,
  Workflow,
  WorkflowActivity,
  WorkflowActivityModel, WorkflowActivityPrompt,
  WorkflowActivityStorageSystem,
  Workflows,
} from '../../generated/resolvers'
import { executeGraphQL, getGraphQLHeaders, GraphQLRequestContext, useClient } from '@bosca/common'
import {
  Empty,
  IdRequest,
  WorkflowActivityIdIntRequest,
  WorkflowExecutionRequest,
  WorkflowService,
} from '@bosca/protobufs'
import { protoInt64 } from '@bufbuild/protobuf'
import { GraphQLError } from 'graphql'

export const resolvers: Resolvers<GraphQLRequestContext> = {
  Query: {
    workflows: () => {
      const workflows: Workflow[] = []
      return {
        workflows: workflows,
      } as Workflows
    },
  },
  Mutation: {
    executeWorkflow: async (_, args, context) => {
      if (!args || !args.request || !args.request.workflowId || !args.request.metadataId) throw new GraphQLError('Invalid request')
      const ctx: { [key: string]: string } = {}
      if (args.request.context) {
        for (const c of args.request.context) {
          ctx[c.key] = c.value
        }
      }
      return await executeGraphQL(async () => {
        const service = useClient(WorkflowService)
        const request = new WorkflowExecutionRequest({ 
          workflowId: args.request!.workflowId!,
          metadataId: args.request!.metadataId!,
          context: ctx,
        })
        const response = await service.executeWorkflow(request, {
          headers: await getGraphQLHeaders(context),
        })
        if (response.error) {
          throw new GraphQLError(response.error)
        }
        return response.jobId
      })
    },
  },
  Model: {
    configuration: (parent) => {
      const configuration = []
      if (parent.configuration) {
        for (const k in parent.configuration) {
          configuration.push({
            key: k,
            value: parent.configuration[k] as unknown as string,
          })
        }
      }
      return configuration
    },
  },
  StorageSystem: {
    configuration: (parent) => {
      parent.configuration = []
      if (parent.configuration) {
        for (const k in parent.configuration) {
          parent.configuration.push({
            key: k,
            value: parent.configuration[k] as unknown as string,
          })
        }
      }
      return parent.configuration
    },
  },
  Workflows: {
    workflow: async (_, args, context) => {
      return await executeGraphQL<Workflow>(async () => {
        const service = useClient(WorkflowService)
        const workflow = await service.getWorkflow(new IdRequest({ id: args.id }), {
          headers: await getGraphQLHeaders(context),
        })
        const j = workflow.toJson() as unknown as Workflow
        if (!j.name) j.name = workflow.id
        return j
      })
    },
    workflows: async (_, __, context) => {
      return await executeGraphQL<Workflow[]>(async () => {
        const service = useClient(WorkflowService)
        const response = await service.getWorkflows(new Empty(), {
          headers: await getGraphQLHeaders(context),
        })
        return response.workflows.map((w) => {
          const j = w.toJson() as unknown as Workflow
          if (!j.name) j.name = w.id
          return j
        }) as unknown as Workflow[]
      })
    },
  },
  Workflow: {
    activities: async (parent, _, context) => {
      return await executeGraphQL<WorkflowActivity[]>(async () => {
        const service = useClient(WorkflowService)
        const response = await service.getWorkflowActivities(new IdRequest({ id: parent.id }), {
          headers: await getGraphQLHeaders(context),
        })
        return response.activities.map((a) => {
          const j = a.toJson() as unknown as WorkflowActivity
          // @ts-ignore
          j.id = j.workflowActivityId
          j.configuration = []
          if (a.configuration) {
            for (const k in a.configuration) {
              j.configuration.push({
                key: k,
                value: a.configuration[k],
              })
            }
          }
          j.inputs = []
          if (a.inputs) {
            for (const k in a.inputs) {
              j.inputs.push({
                key: k,
                value: a.inputs[k],
              })
            }
          }
          j.outputs = []
          if (a.outputs) {
            for (const k in a.outputs) {
              j.outputs.push({
                key: k,
                value: a.outputs[k],
              })
            }
          }
          return j
        }) as unknown as WorkflowActivity[]
      })
    },
  },
  WorkflowActivity: {
    prompts: async (parent, _, context) => {
      return await executeGraphQL<WorkflowActivityPrompt[]>(async () => {
        const service = useClient(WorkflowService)
        // @ts-ignore
        const request = new WorkflowActivityIdIntRequest({ workflowId: parent.workflowId, workflowActivityId: protoInt64.parse(parent.id) })
        const response = await service.getWorkflowActivityPrompts(request, {
          headers: await getGraphQLHeaders(context),
        })
        return response.prompts.map((w) => {
          const j = w.toJson() as unknown as WorkflowActivityPrompt
          j.configuration = []
          if (w.configuration) {
            for (const k in w.configuration) {
              j.configuration.push({
                key: k,
                value: w.configuration[k],
              })
            }
          }
          return j
        }) as unknown as WorkflowActivityPrompt[]
      })
    },
    models: async (parent, _, context) => {
      return await executeGraphQL<WorkflowActivityModel[]>(async () => {
        const service = useClient(WorkflowService)
        // @ts-ignore
        const request = new WorkflowActivityIdIntRequest({ workflowId: parent.workflowId, workflowActivityId: protoInt64.parse(parent.id) })
        const response = await service.getWorkflowActivityModels(request, {
          headers: await getGraphQLHeaders(context),
        })
        return response.models.map((w) => {
          const j = w.toJson() as unknown as WorkflowActivityModel
          j.configuration = []
          if (w.configuration) {
            for (const k in w.configuration) {
              j.configuration.push({
                key: k,
                value: w.configuration[k],
              })
            }
          }
          return j
        }) as unknown as WorkflowActivityModel[]
      })
    },
    storageSystems: async (parent, _, context) => {
      return await executeGraphQL<WorkflowActivityStorageSystem[]>(async () => {
        const service = useClient(WorkflowService)
        // @ts-ignore
        const request = new WorkflowActivityIdIntRequest({ workflowId: parent.workflowId, workflowActivityId: protoInt64.parse(parent.id) })
        const response = await service.getWorkflowActivityStorageSystems(request, {
          headers: await getGraphQLHeaders(context),
        })
        return response.systems.map((w) => {
          const j = w.toJson() as unknown as WorkflowActivityStorageSystem
          j.configuration = []
          if (w.configuration) {
            for (const k in w.configuration) {
              j.configuration.push({
                key: k,
                value: w.configuration[k],
              })
            }
          }
          return j
        }) as unknown as WorkflowActivityStorageSystem[]
      })
    },
  },
}