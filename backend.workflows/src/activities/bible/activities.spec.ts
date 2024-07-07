import { Workflow } from '../../workflow/workflow'
import {
  WorkflowActivityExecutionContext,
  WorkflowActivityInstance
} from '../../generated/protobuf/bosca/content/workflows_pb'
import { Metadata } from '../../generated/protobuf/bosca/content/metadata_pb'
import { Downloader, FileName } from '../../util/downloader'
import { ProcessBibleActivity } from './process'
import { CreateVerseMarkdownTable } from './verse_table'

class DummyDownloader implements Downloader {

  async download(executionContext: WorkflowActivityExecutionContext): Promise<FileName> {
    return '../example-data/asv.zip'
  }

  async cleanup(file: FileName): Promise<void> {
  }
}

test('process bible activity', async () => {
  const downloader = new DummyDownloader()
  const activity = new ProcessBibleActivity(downloader)
  const executionContext = new WorkflowActivityExecutionContext()
  executionContext.metadata = new Metadata({ id: 'asdf' })
  executionContext.activities = [
    new WorkflowActivityInstance({ id: activity.id })
  ]

  const workflow = new Workflow()
  workflow.register(activity)
  await workflow.execute(executionContext)
}, 1200000)

test('create verse tables', async () => {
  const downloader = new DummyDownloader()
  const activity = new CreateVerseMarkdownTable(downloader)
  const executionContext = new WorkflowActivityExecutionContext()
  executionContext.metadata = new Metadata({ id: 'asdf' })
  executionContext.activities = [
    new WorkflowActivityInstance({ id: activity.id })
  ]

  const workflow = new Workflow()
  workflow.register(activity)
  await workflow.execute(executionContext)
}, 1200000)