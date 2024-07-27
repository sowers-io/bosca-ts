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

import { PromptActivity } from './prompt'
import fs from 'node:fs'
import {
  Model,
  Prompt,
  WorkflowActivity,
  WorkflowActivityModel,
  WorkflowActivityPrompt,
  WorkflowJob
} from '@bosca/protobufs'
import { Activity } from '../activity'
import { Job, Queue, Worker } from 'bullmq'

async function runTest(activity: Activity, definition: WorkflowJob) {
  const queue = new Queue('test')
  await queue.add('job', definition.toJson())
  const worker = new Worker('job', async (job: Job) => {
    await activity.newJobExecutor(job, definition).execute()
  })
  await worker.close()
}

test('test prompt', async () => {
  const activity = new PromptActivity()
  const definition = new WorkflowJob({
    workflowId: 'wid',
    metadataId: '540d85d3-7cfe-4b2c-b707-86e5ed2090a8',
    models: [
      new WorkflowActivityModel({
        model: new Model({
          type: 'openai-llm',
          name: 'gpt-4o',
          configuration: {
            temperature: '0',
          }
        }),
      }),
    ],
    prompts: [
      new WorkflowActivityPrompt({
        prompt: new Prompt({
          systemPrompt: fs.readFileSync('../prompts/verselabeller/system', 'utf-8').toString(),
          userPrompt: fs.readFileSync('../prompts/verselabeller/user', 'utf-8').toString(),
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
  await runTest(activity, definition)
}, 1200000)
