import { WorkflowActivityJob } from '../generated/protobuf/bosca/workflow/execution_context_pb'

export type FileName = string

export interface Downloader {

  download(activity: WorkflowActivityJob): Promise<FileName>
  cleanup(file: FileName): Promise<void>
}