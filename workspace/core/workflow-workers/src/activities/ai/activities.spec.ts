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
import * as fs from 'node:fs'
import {
  Model,
  Prompt,
  WorkflowActivity,
  WorkflowActivityModel,
  WorkflowActivityPrompt,
  WorkflowJob,
} from '@bosca/protobufs'
import { Activity } from '../activity'
import { Job, Queue, Worker } from 'bullmq'
import { test } from 'vitest';

async function runTest(activity: Activity, definition: WorkflowJob) {
  const connection = {
    host: 'localhost',
    port: 6379,
  }
  const queue = new Queue('test', {
    connection: connection,
  })
  await queue.add('job', definition.toJson())
  await new Promise((resolve, reject) => {
    new Worker(
      'test',
      async (job: Job) => {
        try {
          await activity.newJobExecutor(job, definition).execute()
          resolve(null)
        } catch (e) {
          reject(e)
        }
      },
      {
        connection: connection,
      },
    )
  })
}

test('test prompt', async () => {
  const activity = new PromptActivity()
  const definition = new WorkflowJob({
    workflowId: 'wid',
    metadataId: '5323d949-5eda-4edd-802e-ba7a2a886059',
    models: [
      new WorkflowActivityModel({
        model: new Model({
          // type: 'google-llm',
          // name: 'gemini-1.5-pro',
          type: 'openai-llm',
          // name: 'gpt-4o-mini',
          name: 'gpt-4o',
          // type: 'ollama-llm',
          // name: 'llama3.1:8b-instruct-q8_0',
          configuration: {
            temperature: '0',
          },
        }),
      }),
    ],
    prompts: [
      new WorkflowActivityPrompt({
        prompt: new Prompt({
          systemPrompt: fs.readFileSync('../../../example-data/prompts/verselabeller/system', 'utf-8').toString(),
          userPrompt: fs.readFileSync('../../../example-data/prompts/verselabeller/user', 'utf-8').toString(),
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
