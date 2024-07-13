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
