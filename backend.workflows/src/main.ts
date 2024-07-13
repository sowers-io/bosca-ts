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

import { WorkflowExecutor } from './workflow_executor'
import { ProcessBibleActivity } from './activities/bible/process'
import { DefaultDownloader } from './util/downloader'
import { CreateVerseMarkdownTable } from './activities/bible/verse_table'
import { Activity } from './activities/activity'
import { ProcessTraitsActivity } from './activities/metadata/traits/process'
import { DeleteBibleActivity } from './activities/bible/delete'
import { TransitionToActivity } from './activities/metadata/transition_to'
import { CreateVerses } from './activities/bible/verse_create'

const downloader = new DefaultDownloader()

const activities = [
  new ProcessTraitsActivity(),
  new ProcessBibleActivity(downloader),
  new DeleteBibleActivity(downloader),
  new CreateVerseMarkdownTable(downloader),
  new CreateVerses(downloader),
  new TransitionToActivity()
]

async function main() {
  const executor = new WorkflowExecutor()
  const activitiesAndQueues = process.env.ACTIVITIES?.split(',')
  if (!activitiesAndQueues) {
    throw new Error('Environment variable `ACTIVITIES` is required')
  }

  const activitiesById: { [id: string]: Activity } = {}
  for (const activity of activities) {
    activitiesById[activity.id] = activity
  }

  for (const activityAndQueueStr of activitiesAndQueues) {
    if (activityAndQueueStr.trim() === '') continue
    const activityAndQueue = activityAndQueueStr.split(':')
    const activityId = activityAndQueue[0].trim()
    if (activityAndQueue.length === 1) {
      throw new Error('missing queue definition for ' + activityId + ' in the form of activity:queue')
    }
    const queue = activityAndQueue[1].trim()

    if (queue === '') {
      throw new Error('missing queue definition for ' + activityId + ' in the form of activity:queue')
    }

    console.log('registering activity ' + activityId + ' for queue ' + queue)

    const activity = activitiesById[activityId]
    if (!activity) throw new Error('couldn\'t find activity: ' + activityId)
    executor.register(queue, activity)
  }

  await executor.execute()
}

main().catch((e) => {
  console.error(e)
  process.exit(1)
})