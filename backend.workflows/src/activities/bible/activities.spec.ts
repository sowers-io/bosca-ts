import { Downloader, FileName } from '../../util/downloader'
import { ProcessBibleActivity } from './process'
import { CreateVerseMarkdownTable } from './verse_table'
import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import { WorkflowActivity } from '../../generated/protobuf/bosca/workflow/activities_pb'

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