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

import { Downloader, FileName } from '../../util/downloader'
import { ProcessBibleActivity } from './process'
import { CreateVerseMarkdownTable } from './verse_table'
import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import { WorkflowActivity } from '../../generated/protobuf/bosca/workflow/activities_pb'
import { CreateVerses } from './verse_create'

class DummyDownloader implements Downloader {

  async download(activity: WorkflowActivityJob): Promise<FileName> {
    return '../example-data/asv.zip'
  }

  async cleanup(file: FileName): Promise<void> {
  }
}

test('process bible activity', async () => {
  const downloader = new DummyDownloader()
  const activity = new ProcessBibleActivity(downloader)
  const activityJob = new WorkflowActivityJob({
    workflowId: 'wid',
    metadataId: 'mid',
    activity: new WorkflowActivity({
      activityId: activity.id
    })
  })
  await activity.execute(activityJob)
}, 1200000)

test('create verse tables', async () => {
  const downloader = new DummyDownloader()
  const activity = new CreateVerseMarkdownTable(downloader)
  const activityJob = new WorkflowActivityJob({
    workflowId: 'wid',
    metadataId: 'mid',
    activity: new WorkflowActivity({
      activityId: activity.id
    })
  })
  await activity.execute(activityJob)
}, 1200000)

test('create verses', async () => {
  const downloader = new DummyDownloader()
  const activity = new CreateVerses(downloader)
  const activityJob = new WorkflowActivityJob({
    workflowId: 'wid',
    metadataId: 'mid',
    activity: new WorkflowActivity({
      activityId: activity.id
    })
  })
  await activity.execute(activityJob)
}, 1200000)