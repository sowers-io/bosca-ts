import { Activity, ActivityJobExecutor } from '../activity'
import { Job } from 'bullmq'
import { WorkflowJob } from '@bosca/protobufs'

export class CreateTextEmbeddings extends Activity {
  get id(): string {
    return 'ai.embeddings.text'
  }

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, definition)
  }
}

class Executor extends ActivityJobExecutor<CreateTextEmbeddings> {
  async execute() {}
}
