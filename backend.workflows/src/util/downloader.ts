import { WorkflowActivityJob } from '../generated/protobuf/bosca/workflow/execution_context_pb'
import { useServiceClient } from './util'
import { ContentService } from '../generated/protobuf/bosca/content/service_connect'
import * as fs from 'node:fs'
import path, { sep } from 'node:path'
import { execute } from './http'
import { IdRequest } from '../generated/protobuf/bosca/requests_pb'
import { Readable } from 'node:stream'
import { promisify } from 'node:util'
import { ReadableStream } from 'node:stream/web'

export type FileName = string

const mkdtemp = promisify(fs.mkdtemp)
const deleteFile = promisify(fs.unlink)
const deleteDirectory = promisify(fs.rmdir)

export interface Downloader {

  download(activity: WorkflowActivityJob): Promise<FileName>

  cleanup(file: FileName): Promise<void>
}

export class DefaultDownloader implements Downloader {

  async download(activity: WorkflowActivityJob): Promise<FileName> {
    const service = useServiceClient(ContentService)
    const directory = await mkdtemp(`${activity.metadataId}${sep}`)
    const url = await service.getMetadataDownloadUrl(new IdRequest({ id: activity.metadataId }))
    const response = await execute(url)
    if (!response.ok) {
      throw new Error('failed to download file: ' + activity.metadataId)
    }
    const fileName = `${directory}${sep}${activity.metadataId}`
    const fileStream = fs.createWriteStream(fileName)
    const readable = Readable.fromWeb(response.body! as ReadableStream)
    await new Promise<void>((resolve, reject) => {
      readable.pipe(fileStream)
      readable.on('error', reject)
      fileStream.on('finish', resolve)
    })
    return fileName
  }

  async cleanup(file: FileName): Promise<void> {
    await deleteFile(file)
    const temporaryDirectory = path.dirname(file)
    await deleteDirectory(temporaryDirectory)
  }
}