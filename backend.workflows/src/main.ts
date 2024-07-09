import { WorkflowExecutor } from './workflow_executor'
import { ProcessBibleActivity } from './activities/bible/process'
import { DefaultDownloader } from './util/downloader'
import { CreateVerseMarkdownTable } from './activities/bible/verse_table'
import { Activity } from './activities/activity'
import { ProcessTraitsActivity } from './activities/metadata/traits/process'

const downloader = new DefaultDownloader()

const activities = [
  new ProcessTraitsActivity(),
  new ProcessBibleActivity(downloader),
  new CreateVerseMarkdownTable(downloader)
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