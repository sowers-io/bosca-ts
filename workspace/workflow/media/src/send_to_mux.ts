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

import { Activity, ActivityJobExecutor, Downloader, Retry, uploadSupplementary } from '@bosca/workflow-activities-api'
import { Job } from 'bullmq'
import { AddMetadataAttributesRequest, ContentService, IdRequest, WorkflowJob } from '@bosca/protobufs'
import { Mux } from '@mux/mux-node'
import * as fs from 'node:fs'
import { promisify } from 'node:util'
import { logger, toArrayBuffer, useServiceAccountClient } from '@bosca/common'

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
      const mux = new Mux({
        tokenId: process.env.MUX_TOKEN_ID,
        tokenSecret: process.env.MUX_TOKEN_SECRET,
      })

      let uploadUrl = this.job.data['uploadUrl']
      let uploadId = this.job.data['uploadId']
      if (!uploadUrl) {
        const initialUpload = await mux.video.uploads.create({
          cors_origin: '*',
          new_asset_settings: {
            playback_policy: ['public'],
            encoding_tier: 'smart',
            test: true,
          },
        })
        uploadUrl = initialUpload.url
        uploadId = initialUpload.id
        await this.job.updateData({
          ...this.job.data,
          'uploadUrl': uploadUrl,
          'uploadId': uploadId,
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
      const { url, assetId } = await Retry.execute(50, async () => {
        const result = await mux.video.uploads.retrieve(uploadId)
        if (!result.asset_id) throw new Error('missing asset id')
        const asset = await mux.video.assets.retrieve(result.asset_id!)
        const playbackId = await mux.video.assets.createPlaybackId(asset.id, {
          policy: 'public',
        })
        return { url: 'https://stream.mux.com/' + playbackId.id + '.m3u8', assetId: result.asset_id! }
      })
      const service = useServiceAccountClient(ContentService)
      const source = await service.getSource(new IdRequest({ id: 'workflow' }))
      await service.addMetadataAttributes(new AddMetadataAttributesRequest({
        id: this.definition.metadataId!,
        attributes: {
          'mux.hls.url': url,
          'mux.asset.id': assetId,
        },
      }))
      await uploadSupplementary(
        this.definition.metadataId!,
        'Mux HLS',
        'text/json',
        this.definition.supplementaryId
          ? this.definition.activity!.outputs['supplementaryId'] + this.definition.supplementaryId
          : this.definition.activity!.outputs['supplementaryId'],
        source.id,
        undefined,
        undefined,
        toArrayBuffer(JSON.stringify({
          'uploadId': uploadId,
          'assetId': assetId,
          'hls.url': url,
        })),
      )
    } finally {
      await this.activity.downloader.cleanup(fileName)
    }
  }
}
