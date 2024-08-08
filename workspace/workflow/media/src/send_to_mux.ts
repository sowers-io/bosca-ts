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

import { Activity, ActivityJobExecutor, Downloader, Retry } from '@bosca/workflow-activities-api'
import { Job } from 'bullmq'
import { WorkflowJob } from '@bosca/protobufs'
import { Mux } from '@mux/mux-node'
import * as fs from 'node:fs'
import { promisify } from 'node:util'
import { logger } from '@bosca/common'

export class SendToMux extends Activity {

  readonly downloader: Downloader

  constructor(downloader: Downloader) {
    super()
    this.downloader = downloader
  }

  get id(): string {
    return 'video.send_to_mux'
  }

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, definition)
  }
}

const readFile = promisify(fs.readFile)

class Executor extends ActivityJobExecutor<SendToMux> {
  async execute() {
    const fileName = await this.activity.downloader.download(this.definition)
    try {
      let uploadUrl = await this.job.data['uploadUrl']
      if (!uploadUrl) {
        const mux = new Mux({
          tokenId: process.env.MUX_TOKEN_ID,
          tokenSecret: process.env.MUX_TOKEN_SECRET,
        })
        const initialUpload = await mux.video.uploads.create({
          cors_origin: '*',
          new_asset_settings: {
            playback_policy: ['public'],
            encoding_tier: 'smart',
            test: true,
          },
        })
        uploadUrl = initialUpload.url
        await this.job.updateData({
          ...this.job.data,
          'uploadUrl': uploadUrl,
        })
      }
      const fileBuffer = await readFile(fileName)
      const chunkSize = 1024 * 1024 * 30
      for (let offset = 0; offset < fileBuffer.length; offset += chunkSize) {
        const chunk = fileBuffer.subarray(offset, offset + chunkSize)
        logger.debug(`bytes ${offset}-${offset + chunk.length - 1}/${fileBuffer.length}`)
        const headers = new Headers()
        headers.set('Content-Length', chunk.length.toString())
        headers.set('Content-Range', `bytes ${offset}-${offset + chunk.length - 1}/${fileBuffer.length}`)
        await Retry.execute(10, async () => {
          const response = await fetch(uploadUrl, {
            body: chunk,
            method: 'PUT',
            headers: headers,
          })
          if (response.status == 308) {
            return
          }
          if (!response.ok) {
            throw new Error('failed to upload: ' + response.status + ': ' + await response.text())
          }
        })
      }
    } finally {
      await this.activity.downloader.cleanup(fileName)
    }
  }
}
