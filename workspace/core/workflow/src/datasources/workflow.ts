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

import { DataSource, logger, getIStorageSystem } from '@bosca/common'
import { promisify } from 'node:util'
import fs from 'node:fs'
import { parse } from 'yaml'
import {
  Activity,
  Model,
  Prompt,
  StorageSystem,
  StorageSystemModel,
  StorageSystemType,
  Trait,
  Workflow,
  WorkflowActivity, WorkflowActivityModel,
  WorkflowActivityParameterType,
  WorkflowActivityPrompt,
  WorkflowActivityStorageSystem,
  WorkflowState, WorkflowStateTransition,
  WorkflowStateType,
} from '@bosca/protobufs'
import { proto3 } from '@bufbuild/protobuf'

const readFile = promisify(fs.readFile)

export class WorkflowDataSource extends DataSource {
  async getModels(): Promise<Model[]> {
    return await this.queryAndMap(() => new Model(), 'select * from models')
  }

  async getModel(id: string): Promise<Model | null> {
    return await this.queryAndMapFirst(() => new Model(), 'select * from models where id::uuid = $1', [id])
  }

  async addModel(model: Model): Promise<string> {
    const records = await this.query(
      'insert into models (type, name, description, configuration) values ($1, $2, $3, ($4)::jsonb) returning id::varchar',
      [model.type, model.name, model.description, JSON.stringify(model.configuration)],
    )
    return records.rows[0].id
  }

  async getStorageSystems(): Promise<StorageSystem[]> {
    return await this.queryAndMap(() => new StorageSystem(), 'select * from storage_systems')
  }

  async getStorageSystem(id: string): Promise<StorageSystem | null> {
    return await this.queryAndMapFirst(() => new StorageSystem(), 'select * from storage_systems where id = $1', [id])
  }

  async addStorageSystem(storageSystem: StorageSystem): Promise<string> {
    const StorageSystemTypeEnum = proto3.getEnumType(StorageSystemType)
    const records = await this.query(
      'insert into storage_systems (type, name, description, configuration) values ($1, $2, $3, ($4)::jsonb) returning id::varchar',
      [
        StorageSystemTypeEnum.findNumber(storageSystem.type)?.name.replace('_storage_system', ''),
        storageSystem.name,
        storageSystem.description,
        JSON.stringify(storageSystem.configuration),
      ],
    )
    storageSystem.id = records.rows[0].id
    await getIStorageSystem(storageSystem)
    return storageSystem.id
  }

  async getStorageSystemModels(id: string): Promise<StorageSystemModel[]> {
    const records = await this.query('select model_id, configuration from storage_system_models where system_id = $1', [
      id,
    ])
    const models: StorageSystemModel[] = []
    for (const record of records.rows) {
      const model = await this.getModel(record.model_id)
      if (model == null) throw new Error('missing model')
      models.push(
        new StorageSystemModel({
          model: model,
          configuration: record.configuration,
        }),
      )
    }
    return models
  }

  async addStorageSystemModel(systemId: string, model: StorageSystemModel) {
    await this.query(
      'insert into storage_system_models (system_id, model_id, configuration) values ($1, $2, ($3)::jsonb)',
      [systemId, model.model!.id, JSON.stringify(model.configuration)],
    )
  }

  async addTrait(trait: Trait) {
    await this.query('insert into traits (id, name, description) values ($1, $2, $3)', [
      trait.id,
      trait.name,
      trait.description,
    ])
    for (const workflowId of trait.workflowIds) {
      await this.query('insert into trait_workflows (trait_id, workflow_id) values ($1, $2)', [trait.id, workflowId])
    }
  }

  async getPrompts(): Promise<Prompt[]> {
    return await this.queryAndMap(() => new Prompt(), 'select * from prompts')
  }

  async getPrompt(id: string): Promise<Prompt | null> {
    return await this.queryAndMapFirst(() => new Prompt(), 'select * from prompts where id = $1::uuid', [id])
  }

  async addPrompt(prompt: Prompt): Promise<string> {
    const records = await this.query(
      'insert into prompts (name, description, system_prompt, user_prompt, input_type, output_type) values ($1, $2, $3, $4, $5, $6) returning id::varchar',
      [prompt.name, prompt.description, prompt.systemPrompt, prompt.userPrompt, prompt.inputType, prompt.outputType],
    )
    return records.rows[0].id
  }

  async addActivity(activity: Activity): Promise<void> {
    await this.query(
      'insert into activities (id, name, description, child_workflow_id, configuration) values ($1, $2, $3, $4, ($5)::jsonb)',
      [
        activity.id,
        activity.name,
        activity.description,
        activity.childWorkflowId,
        JSON.stringify(activity.configuration),
      ],
    )

    const WorkflowActivityParameterTypeEnum = proto3.getEnumType(WorkflowActivityParameterType)
    if (activity.inputs) {
      for (const inputId in activity.inputs) {
        await this.query('insert into activity_inputs (activity_id, name, type) values ($1, $2, $3)', [
          activity.id,
          inputId,
          WorkflowActivityParameterTypeEnum.findNumber(activity.inputs[inputId])?.name,
        ])
      }
    }
    if (activity.outputs) {
      for (const outputId in activity.outputs) {
        await this.query('insert into activity_outputs (activity_id, name, type) values ($1, $2, $3)', [
          activity.id,
          outputId,
          WorkflowActivityParameterTypeEnum.findNumber(activity.outputs[outputId])?.name,
        ])
      }
    }
  }

  async getWorkflows(): Promise<Workflow[]> {
    return this.queryAndMap(() => new Workflow(), 'select * from workflows')
  }

  async getWorkflow(id: string): Promise<Workflow | null> {
    return this.queryAndMapFirst(() => new Workflow(), 'select * from workflows where id = $1', [id])
  }

  async getWorkflowStates(): Promise<WorkflowState[]> {
    return this.queryAndMap(() => new WorkflowState(), 'select * from workflow_states')
  }

  async getWorkflowState(id: string): Promise<WorkflowState | null> {
    return this.queryAndMapFirst(() => new WorkflowState(), 'select * from workflow_states where id = $1', [id])
  }

  async addWorkflow(workflow: Workflow) {
    await this.query(
      'insert into workflows (id, name, description, queue, configuration) values ($1, $2, $3, $4, ($5)::jsonb)',
      [workflow.id, workflow.name, workflow.description, workflow.queue, JSON.stringify(workflow.configuration)],
    )
  }

  private async processActivity(activity: WorkflowActivity) {
    activity.inputs = {}
    activity.outputs = {}
    const inputs = await this.query('select name, value from workflow_activity_inputs where activity_id = $1', [
      activity.workflowActivityId,
    ])
    for (const input of inputs.rows) {
      activity.inputs[input.name] = input.value
    }
    const outputs = await this.query('select name, value from workflow_activity_outputs where activity_id = $1', [
      activity.workflowActivityId,
    ])
    for (const output of outputs.rows) {
      activity.outputs[output.name] = output.value
    }
  }

  async getWorkflowActivities(workflowId: string): Promise<WorkflowActivity[]> {
    const activities = await this.queryAndMap(
      () => new WorkflowActivity(),
      'select id as workflow_activity_id, * from workflow_activities where workflow_id = $1 order by execution_group asc',
      [workflowId],
    )
    for (const activity of activities) {
      await this.processActivity(activity)
    }
    return activities
  }

  async getWorkflowActivity(id: number): Promise<WorkflowActivity | null> {
    const activity = await this.queryAndMapFirst(
      () => new WorkflowActivity(),
      'select id as workflow_activity_id, * from workflow_activities where id = $1',
      [id],
    )
    if (activity) {
      await this.processActivity(activity)
    }
    return activity
  }

  async addWorkflowActivity(workflowId: string, activity: WorkflowActivity): Promise<number> {
    const records = await this.query(
      'insert into workflow_activities (workflow_id, activity_id, execution_group, queue, configuration) values ($1, $2, $3, $4, ($5)::jsonb) returning id',
      [workflowId, activity.activityId, activity.executionGroup, activity.queue, JSON.stringify(activity.configuration)],
    )
    for (const inputKey in activity.inputs) {
      await this.query('insert into workflow_activity_inputs (activity_id, name, value) values ($1, $2, $3)', [
        records.rows[0].id,
        inputKey,
        activity.inputs[inputKey],
      ])
    }
    for (const outputsKey in activity.outputs) {
      await this.query('insert into workflow_activity_outputs (activity_id, name, value) values ($1, $2, $3)', [
        records.rows[0].id,
        outputsKey,
        activity.outputs[outputsKey],
      ])
    }
    return records.rows[0].id
  }

  async getWorkflowActivityStorageSystems(activityId: number): Promise<WorkflowActivityStorageSystem[]> {
    const records = await this.query('select * from workflow_activity_storage_systems where activity_id = $1', [
      activityId,
    ])
    const storageSystems: WorkflowActivityStorageSystem[] = []
    for (const record of records.rows) {
      storageSystems.push(
        new WorkflowActivityStorageSystem({
          storageSystem: (await this.getStorageSystem(record.storage_system_id))!,
          models: await this.getStorageSystemModels(record.storage_system_id),
          configuration: record.configuration,
        }),
      )
    }
    return storageSystems
  }

  async addWorkflowActivityStorageSystem(
    activityId: number,
    storageSystemId: string,
    configuration: {
      [key: string]: string
    },
  ) {
    await this.query(
      'insert into workflow_activity_storage_systems (activity_id, storage_system_id, configuration) values ($1, $2, ($3)::jsonb)',
      [activityId, storageSystemId, JSON.stringify(configuration || '{}')],
    )
  }

  async getWorkflowActivityPrompts(activityId: number): Promise<WorkflowActivityPrompt[]> {
    const records = await this.query('select * from workflow_activity_prompts where activity_id = $1', [activityId])
    const prompts: WorkflowActivityPrompt[] = []
    for (const record of records.rows) {
      prompts.push(
        new WorkflowActivityPrompt({
          prompt: (await this.getPrompt(record.prompt_id))!,
          configuration: record.configuration,
        }),
      )
    }
    return prompts
  }

  async addWorkflowActivityPrompt(
    activityId: number,
    promptId: string,
    configuration: {
      [key: string]: string
    },
  ) {
    await this.query(
      'insert into workflow_activity_prompts (activity_id, prompt_id, configuration) values ($1, $2, ($3)::jsonb)',
      [activityId, promptId, JSON.stringify(configuration || '{}')],
    )
  }

  async getWorkflowActivityModels(activityId: number): Promise<WorkflowActivityModel[]> {
    const records = await this.query('select * from workflow_activity_models where activity_id = $1', [
      activityId,
    ])
    const models: WorkflowActivityModel[] = []
    for (const record of records.rows) {
      models.push(
        new WorkflowActivityModel({
          model: (await this.getModel(record.model_id))!,
          configuration: record.configuration,
        }),
      )
    }
    return models
  }

  async addWorkflowActivityModel(
    activityId: number,
    modelId: string,
    configuration: {
      [key: string]: string
    },
  ) {
    await this.query(
      'insert into workflow_activity_models (activity_id, model_id, configuration) values ($1, $2, ($3)::jsonb)',
      [activityId, modelId, JSON.stringify(configuration || '{}')],
    )
  }

  async getWorkflowTransition(fromStateId: string, toStateId: String): Promise<WorkflowStateTransition | null> {
    return await this.queryAndMapFirst(
      () => new WorkflowStateTransition(),
      'select * from workflow_state_transitions where from_state_id = $1 and to_state_id = $2',
      [fromStateId, toStateId],
    )
  }

  async addTransition(fromStateId: string, toStateId: string, description: string) {
    await this.query(
      'insert into workflow_state_transitions (from_state_id, to_state_id, description) values ($1, $2, $3)',
      [fromStateId, toStateId, description],
    )
  }

  async addWorkflowState(state: WorkflowState) {
    const WorkflowStateTypeEnum = proto3.getEnumType(WorkflowStateType)
    await this.query(
      'insert into workflow_states (id, name, description, type, configuration, workflow_id, exit_workflow_id, entry_workflow_id) values ($1, $2, $3, $4, ($5)::jsonb, $6, $7, $8)',
      [
        state.id,
        state.name,
        state.description,
        WorkflowStateTypeEnum.findNumber(state.type)?.name,
        JSON.stringify(state.configuration),
        state.workflowId,
        state.exitWorkflowId,
        state.entryWorkflowId,
      ],
    )
  }

  async initialize() {
    const current = await this.getWorkflows()
    if (current.length > 0) {
      logger.info('initial workflows already initialized')
      return
    }
    const workflows = parse(await readFile('workflows.yaml', 'utf8'))
    const storageSystemIds: { [id: string]: string } = {}
    const models: { [id: string]: Model } = {}
    const promptIds: { [id: string]: string } = {}
    for (const modelId of Object.keys(workflows.models)) {
      const m = workflows.models[modelId]
      for (const key of Object.keys(m.configuration)) {
        m.configuration[key] = m.configuration[key].toString()
      }
      const model = Model.fromJson(m, { ignoreUnknownFields: true })
      model.id = await this.addModel(model)
      models[modelId] = model
    }
    for (const storageSystemId of Object.keys(workflows.storageSystems)) {
      const s = workflows.storageSystems[storageSystemId]
      for (const key of Object.keys(s.configuration)) {
        s.configuration[key] = s.configuration[key].toString()
      }
      const ss = StorageSystem.fromJson(s, { ignoreUnknownFields: true })
      switch (s.type) {
        case 'search':
          ss.type = StorageSystemType.search_storage_system
          break
        case 'vector':
          ss.type = StorageSystemType.vector_storage_system
          break
        case 'metadata':
          ss.type = StorageSystemType.metadata_storage_system
          break
        case 'supplementary':
          ss.type = StorageSystemType.supplementary_storage_system
          break
      }
      storageSystemIds[storageSystemId] = await this.addStorageSystem(ss)
      if (s.models) {
        for (const modelId of Object.keys(s.models)) {
          const sm = StorageSystemModel.fromJson(s.models[modelId], { ignoreUnknownFields: true })
          sm.model = models[modelId]
          await this.addStorageSystemModel(storageSystemIds[storageSystemId], sm)
        }
      }
    }
    for (const promptId of Object.keys(workflows.prompts)) {
      const prompt = Prompt.fromJson(workflows.prompts[promptId], { ignoreUnknownFields: true })
      promptIds[promptId] = await this.addPrompt(prompt)
    }
    for (const activityId of Object.keys(workflows.workflows.activities)) {
      const a = workflows.workflows.activities[activityId]
      if (a.configuration) {
        for (const key of Object.keys(a.configuration)) {
          a.configuration[key] = a.configuration[key].toString()
        }
      }
      const activity = Activity.fromJson(a, { ignoreUnknownFields: true })
      activity.id = activityId
      await this.addActivity(activity)
    }
    for (const workflowId of Object.keys(workflows.workflows.workflows)) {
      const workflow = Workflow.fromJson(workflows.workflows.workflows[workflowId], { ignoreUnknownFields: true })
      workflow.id = workflowId
      await this.addWorkflow(workflow)

      if (workflows.workflows.workflows[workflowId].activities) {
        for (const activityId of Object.keys(workflows.workflows.workflows[workflowId].activities)) {
          const a = workflows.workflows.workflows[workflowId].activities[activityId]
          if (!a.queue || a.queue === '') {
            a.queue = workflow.queue
          }
          if (a.configuration) {
            for (const key in a.configuration) {
              if (a.configuration[key]) {
                a.configuration[key] = a.configuration[key].toString()
              } else {
                delete a.configuration[key]
              }
            }
          }
          const w = WorkflowActivity.fromJson(a, { ignoreUnknownFields: true })
          w.activityId = activityId
          const workflowActivityId = await this.addWorkflowActivity(workflowId, w)
          if (a.storageSystems) {
            for (const storageSystemId of Object.keys(a.storageSystems)) {
              const s = a.storageSystems[storageSystemId]
              if (s.configuration) {
                for (const key of Object.keys(s.configuration)) {
                  s.configuration[key] = s.configuration[key].toString()
                }
              }
              await this.addWorkflowActivityStorageSystem(
                workflowActivityId,
                storageSystemIds[storageSystemId],
                s.configuration,
              )
            }
          }
          if (a.prompts) {
            for (const promptId of Object.keys(a.prompts)) {
              const s = a.prompts[promptId]
              if (s.configuration) {
                for (const key of Object.keys(s.configuration)) {
                  s.configuration[key] = s.configuration[key].toString()
                }
              }
              await this.addWorkflowActivityPrompt(workflowActivityId, promptIds[promptId], s.configuration)
            }
          }
          if (a.models) {
            for (const modelId of Object.keys(a.models)) {
              const s = a.models[modelId]
              if (s.configuration) {
                for (const key of Object.keys(s.configuration)) {
                  s.configuration[key] = s.configuration[key].toString()
                }
              }
              await this.addWorkflowActivityModel(workflowActivityId, models[modelId].id, s.configuration)
            }
          }
        }
      }
    }
    for (const traitId of Object.keys(workflows.traits)) {
      const trait = Trait.fromJson(workflows.traits[traitId], { ignoreUnknownFields: true })
      trait.id = traitId
      await this.addTrait(trait)
    }
    for (const stateId of Object.keys(workflows.workflows.states)) {
      const state = WorkflowState.fromJson(workflows.workflows.states[stateId], { ignoreUnknownFields: true })
      state.id = stateId
      await this.addWorkflowState(state)
    }
    for (const transition of workflows.workflows.transitions) {
      await this.addTransition(transition.from, transition.to, transition.description)
    }
  }
}
