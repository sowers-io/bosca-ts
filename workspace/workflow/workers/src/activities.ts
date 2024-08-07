import {
  CreateVerseJsonTable,
  CreateVerses,
  DeleteBibleActivity,
  ProcessBibleActivity,
} from '@bosca/workflow-bible-activities';
import {
  CreatePendingEmbeddingsFromJsonTable,
  CreatePendingEmbeddingsIndex,
  CreateTextEmbeddings,
  PromptActivity,
} from '@bosca/workflow-ai-activities';
import { ProcessTraitsActivity, ChildWorkflow, IndexText, TransitionToActivity } from '@bosca/workflow-metadata-activities';
import { Activity, DefaultDownloader } from '@bosca/workflow-activities-api'

const downloader = new DefaultDownloader()

export function getActivities(): { [id: string]: Activity } {
  const activities = [
    new ProcessTraitsActivity(),
    new ProcessBibleActivity(downloader),
    new DeleteBibleActivity(downloader),
    new CreateVerseJsonTable(downloader),
    new CreateVerses(downloader),
    new TransitionToActivity(),
    new PromptActivity(),
    new ChildWorkflow(),
    new CreatePendingEmbeddingsFromJsonTable(),
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
