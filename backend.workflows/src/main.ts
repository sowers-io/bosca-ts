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

import { Executor } from './executor'
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
    new ChildWorkflow()
  ]
  const activitiesById: { [id: string]: Activity } = {}
  for (const activity of activities) {
    activitiesById[activity.id] = activity
  }
  return activitiesById
}

async function main() {
  const executor = new Executor()
  const configuration = getConfiguration()
  const activities = getAvailableActivities()

  for (const queueConfigurationId in configuration.queues) {
    const queueConfiguration = configuration.queues[queueConfigurationId]
    executor.registerQueue(queueConfigurationId, queueConfiguration.maxConcurrency)
  }

  for (const activityConfigurationId in configuration.activities) {
    const activityConfiguration = configuration.activities[activityConfigurationId]
    const activity = activities[activityConfigurationId]
    if (!activity) throw new Error("couldn't find activity: " + activityConfigurationId)
    executor.registerActivity(activityConfiguration.queue, activity)
  }

  await executor.execute()
}

main().catch((e) => {
  console.error(e)
  process.exit(1)
})
