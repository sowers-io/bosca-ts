import { Activity, ActivityJobExecutor } from '../activity'
import { Job } from 'bullmq'
import { WorkflowJob } from '@bosca/protobufs'

export class CreatePendingEmbeddingsFromMarkdownTable extends Activity {
  get id(): string {
    return 'ai.embeddings.pending.from-markdown-table'
  }

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, definition)
  }
}

class Executor extends ActivityJobExecutor<CreatePendingEmbeddingsFromMarkdownTable> {
  async execute() {}
}
