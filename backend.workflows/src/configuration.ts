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

import fs from 'node:fs'

export interface Configuration {
  queues: { [id: string]: QueueConfiguration }
  activities: { [id: string]: ActivityConfiguration }
}

export interface QueueConfiguration {
  id: string
  maxConcurrency: number
}

export interface ActivityConfiguration {
  id: string
  queue: string
}

export function getConfiguration(): Configuration {
  const configuration: Configuration = JSON.parse(fs.readFileSync('configuration.json', 'utf8'))

  for (const queueId in configuration.queues) {
    const queue = configuration.queues[queueId]
    queue.id = queueId
  }

  for (const activityId in configuration.activities) {
    const activity = configuration.activities[activityId]
    activity.id = activityId

    if (!configuration.queues[activity.queue]) {
      throw new Error('Queue (' + activity.queue + ') not found for activity: ' + activityId)
    }
  }

  return configuration
}
