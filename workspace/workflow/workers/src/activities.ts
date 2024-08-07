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
