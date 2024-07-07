import { WorkflowActivityExecutionContext } from '../generated/protobuf/bosca/content/workflows_pb'

export type FileName = string

export interface Downloader {

  download(executionContext: WorkflowActivityExecutionContext): Promise<FileName>

  cleanup(file: FileName): Promise<void>
}