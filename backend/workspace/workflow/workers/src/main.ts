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

import { ProcessBibleActivity } from './activities/bible/process'
import { DefaultDownloader } from './util/downloader'
import { Activity } from './activities/activity'
import { ProcessTraitsActivity } from './activities/metadata/traits/process'
import { DeleteBibleActivity } from './activities/bible/delete'
import { TransitionToActivity } from './activities/metadata/transition_to'
import { getConfiguration } from './configuration'
import { CreateVerseMarkdownTable } from './activities/bible/book/verse_table'
import { CreateVerses } from './activities/bible/book/verse_create'
import { PromptActivity } from './activities/ai/prompt'
import { ChildWorkflow } from './activities/metadata/child_workflow'
import { ConnectionOptions, QueueEvents, Worker } from 'bullmq'
import { WorkflowJob } from '@bosca/protobufs'
import { CreatePendingEmbeddingsFromMarkdownTable } from './activities/ai/create_pending_embeddings_markdown'
import { CreateTextEmbeddings } from './activities/ai/create_text_embeddings'
import { CreatePendingEmbeddingsIndex } from './activities/ai/create_pending_embeddings_index'
import { IndexText } from './activities/metadata/text'
import { initializeUploadLimiter } from './util/uploader'

const downloader = new DefaultDownloader()

function getAvailableActivities(): { [id: string]: Activity } {
  const activities = [
    new ProcessTraitsActivity(),
    new ProcessBibleActivity(downloader),
    new DeleteBibleActivity(downloader),
    new CreateVerseMarkdownTable(downloader),
    new CreateVerses(downloader),
    new TransitionToActivity(),
    new PromptActivity(),
    new ChildWorkflow(),
    new CreatePendingEmbeddingsFromMarkdownTable(),
    new CreatePendingEmbeddingsIndex(),
    new CreateTextEmbeddings(),
    new IndexText(),
  ]
  const activitiesById: { [id: string]: Activity } = {}
  for (const activity of activities) {
    activitiesById[activity.id] = activity
  }
  return activitiesById
}

async function main() {
  const configuration = getConfiguration()
  const activities = getAvailableActivities()

  initializeUploadLimiter(configuration.maxUploadConcurrency)

  const connection: ConnectionOptions = {
    host: process.env.BOSCA_REDIS_HOST!,
    port: parseInt(process.env.BOSCA_REDIS_PORT!),
  }

  for (const queueConfigurationId in configuration.queues) {
    const queueConfiguration = configuration.queues[queueConfigurationId]
    const worker = new Worker(
      queueConfigurationId,
      async (job) => {
        console.log('new job', job.id, job.name)
        const definition = WorkflowJob.fromJson(job.data)
        if (!definition.activity) return
        const activity = activities[definition.activity.activityId]
        if (!activity) {
          throw new Error('activity not found: ' + definition.activity.activityId)
        }
        const executor = activity.newJobExecutor(job, definition)
        return await executor.execute()
      },
      {
        connection: connection,
        concurrency: queueConfiguration.maxConcurrency,
        stalledInterval: 600000,
        lockDuration: 600000,
        maxStalledCount: 10,
      }
    )
    worker.on('completed', (job) => {
      console.log(`${job.id} has completed!`)
    })
    worker.on('failed', (job, err) => {
      console.log(`${job.id} has failed with ${err.message}`)
    })

    const events = new QueueEvents(queueConfigurationId, { connection })
    events.on('added', async (job) => {
      console.log('job added', job.jobId, job.name)
    })
    events.on('error', async (err) => {
      console.log('job error', err)
    })
  }
  console.log('running...')
}

main().catch((e) => {
  console.error(e)
  process.exit(1)
})
