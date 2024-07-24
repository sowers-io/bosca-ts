import { Activity, ActivityJobExecutor } from '../../activity'
import { Job } from 'bullmq'
import { WorkflowJob } from '@bosca/protobufs'

export class IndexText extends Activity {
  get id(): string {
    return 'metadata.text.index'
  }

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, definition)
  }
}

class Executor extends ActivityJobExecutor<IndexText> {
  async execute() {}
}
