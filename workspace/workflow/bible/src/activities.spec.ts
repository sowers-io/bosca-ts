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

import { Activity, Downloader, FileName } from '@bosca/workflow-activities-api'
import { ProcessBibleActivity } from './process'
import { CreateVerseJsonTable } from './book/verse_table'
import { CreateVerses } from './book/verse_create'
import { WorkflowActivity, WorkflowJob } from '@bosca/protobufs'
import { Job, Queue, Worker } from 'bullmq'
import { test } from 'vitest'

class DummyDownloader implements Downloader {

  newTemporaryFile(_: string): Promise<FileName> {
    throw new Error('unimplemented')
  }

  async download(_: WorkflowJob): Promise<FileName> {
    return '../example-data/asv.zip'
  }

  async cleanup(_: FileName): Promise<void> {}
}

async function runTest(activity: Activity, definition: WorkflowJob) {
  const queue = new Queue('test')
  await queue.add('job', definition.toJson())
  const worker = new Worker('job', async (job: Job) => {
    await activity.newJobExecutor(job, definition).execute()
  })
  await worker.close()
}

test('process bible activity', async () => {
  const downloader = new DummyDownloader()
  const activity = new ProcessBibleActivity(downloader)
  const definition = new WorkflowJob({
    workflowId: 'wid',
    metadataId: 'mid',
    activity: new WorkflowActivity({
      activityId: activity.id,
    }),
  })
  await runTest(activity, definition)
}, 1200000)

test('create verse tables', async () => {
  const downloader = new DummyDownloader()
  const activity = new CreateVerseJsonTable(downloader)
  const definition = new WorkflowJob({
    workflowId: 'wid',
    metadataId: 'mid',
    activity: new WorkflowActivity({
      activityId: activity.id,
    }),
  })
  await runTest(activity, definition)
}, 1200000)

test('create verses', async () => {
  const downloader = new DummyDownloader()
  const activity = new CreateVerses(downloader)
  const definition = new WorkflowJob({
    workflowId: 'wid',
    metadataId: 'mid',
    activity: new WorkflowActivity({
      activityId: activity.id,
    }),
  })
  await runTest(activity, definition)
}, 1200000)
