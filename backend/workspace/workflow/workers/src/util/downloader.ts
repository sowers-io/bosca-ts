/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

import * as fs from 'node:fs'
import path, { sep } from 'node:path'
import { execute } from './http'
import { promisify } from 'node:util'
import { tmpdir } from 'node:os'
import { ContentService, IdRequest, WorkflowJob } from '@bosca/protobufs'
import { useServiceAccountClient } from '@bosca/common'

export type FileName = string

const mkdtemp = promisify(fs.mkdtemp)
const writeFile = promisify(fs.writeFile)
const deleteFile = promisify(fs.unlink)
const deleteDirectory = promisify(fs.rmdir)

export interface Downloader {

  download(definition: WorkflowJob): Promise<FileName>

  cleanup(file: FileName): Promise<void>
}

export class DefaultDownloader implements Downloader {

  async download(definition: WorkflowJob): Promise<FileName> {
    const service = useServiceAccountClient(ContentService)
    const temporaryDirectory = tmpdir()
    const directory = await mkdtemp(`${temporaryDirectory}${sep}`)
    const url = await service.getMetadataDownloadUrl(new IdRequest({ id: definition.metadataId }))
    const buffer = await execute(url)
    const fileName = `${directory}${sep}${definition.metadataId}`
    await writeFile(fileName, buffer)
    return fileName
  }

  async cleanup(file: FileName): Promise<void> {
    await deleteFile(file)
    const temporaryDirectory = path.dirname(file)
    await deleteDirectory(temporaryDirectory)
  }
}