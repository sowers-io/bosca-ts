import { Workflow } from '../../workflow/workflow'
import { ProcessBibleActivity, ProcessBibleDownloader } from './activities'
import {
  WorkflowActivityExecutionContext,
  WorkflowActivityInstance
} from '../../generated/protobuf/bosca/content/workflows_pb'
import { Metadata } from '../../generated/protobuf/bosca/content/metadata_pb'

class DummyDownloader implements ProcessBibleDownloader {

  async download(executionContext: WorkflowActivityExecutionContext): Promise<string> {
    return '../example-data/asv.zip'
  }

  async cleanup(file: string): Promise<void> {
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