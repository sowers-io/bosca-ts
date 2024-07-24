import { Activity, ActivityJobExecutor } from '../activity'
import { Job } from 'bullmq'
import { WorkflowJob } from '@bosca/protobufs'

export class CreatePendingEmbeddingsIndex extends Activity {
  get id(): string {
    return 'ai.embeddings.pending.index'
  }

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, definition)
  }
}

class Executor extends ActivityJobExecutor<CreatePendingEmbeddingsIndex> {
  async execute() {}
}
