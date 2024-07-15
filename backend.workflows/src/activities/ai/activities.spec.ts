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

import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import {
  WorkflowActivity,
  WorkflowActivityModel,
  WorkflowActivityPrompt,
} from '../../generated/protobuf/bosca/workflow/activities_pb'
import { PromptActivity } from './prompt'
import { Model } from '../../generated/protobuf/bosca/workflow/models_pb'
import { Prompt } from '../../generated/protobuf/bosca/workflow/prompts_pb'
import fs from 'node:fs'

test('test prompt', async () => {
  const activity = new PromptActivity()
  const activityJob = new WorkflowActivityJob({
    workflowId: 'wid',
    metadataId: '540d85d3-7cfe-4b2c-b707-86e5ed2090a8',
    models: [
      new WorkflowActivityModel({
        model: new Model({ name: 'gemma2:27b' }),
      }),
    ],
    prompts: [
      new WorkflowActivityPrompt({
        prompt: new Prompt({
          systemPrompt: '',
          userPrompt: fs.readFileSync('../prompts/verselabeller', 'utf-8').toString(),
        }),
      }),
    ],
    activity: new WorkflowActivity({
      activityId: activity.id,
      inputs: {
        supplementaryId: 'chapter-verse-table',
      },
      outputs: {
        supplementaryId: 'chapter-verse-table-prompted',
      },
    }),
  })
  await activity.execute(activityJob)
}, 1200000)
