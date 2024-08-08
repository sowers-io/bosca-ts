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

import { Activity, ActivityJobExecutor, Downloader, uploadSupplementary } from '@bosca/workflow-activities-api'
import { Job } from 'bullmq'
import { ContentService, IdRequest, WorkflowJob } from '@bosca/protobufs'
import { toArrayBuffer, useServiceAccountClient } from '@bosca/common'
import { downloadHLS, extractMP3, transcribe } from './ffmpeg'

export class Transcribe extends Activity {

  readonly downloader: Downloader

  constructor(downloader: Downloader) {
    super()
    this.downloader = downloader
  }

  get id(): string {
    return 'media.transcribe'
  }

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, definition)
  }
}

class Executor extends ActivityJobExecutor<Transcribe> {
  async execute() {
    const source = await useServiceAccountClient(ContentService).getSource(new IdRequest({ id: 'workflow' }))
    const idRequest = new IdRequest({ id: this.definition.metadataId })
    const service = useServiceAccountClient(ContentService)
    const metadata = await service.getMetadata(idRequest)
    let transcribed: any | null = null

    switch (metadata.contentType.toLowerCase()) {
      case 'video/mp4': {
        const fileName = await this.activity.downloader.download(this.definition)
        const mp3File = await this.activity.downloader.newTemporaryFile(this.definition.metadataId + '-mp3')
        try {
          await extractMP3(fileName, mp3File)
          transcribed = await transcribe(mp3File, metadata.languageTag)
        } finally {
          await this.activity.downloader.cleanup(mp3File)
          await this.activity.downloader.cleanup(fileName)
        }
        break
      }
      case 'application/x-mpegurl':
      case 'vnd.apple.mpegurl':
      case 'application/vnd.apple.mpegurl': {
        const mp4File = await this.activity.downloader.newTemporaryFile(this.definition.metadataId + '-mp4')
        const mp3File = await this.activity.downloader.newTemporaryFile(this.definition.metadataId + '-mp3')
        try {
          await downloadHLS(metadata.sourceIdentifier!, mp4File)
          await extractMP3(mp4File, mp3File)
          transcribed = await transcribe(mp3File, metadata.languageTag)
        } finally {
          await this.activity.downloader.cleanup(mp3File)
          await this.activity.downloader.cleanup(mp4File)
        }
        break
      }
      default: {
        const fileName = await this.activity.downloader.download(this.definition)
        try {
          transcribed = await transcribe(fileName, metadata.languageTag)
        } finally {
          await this.activity.downloader.cleanup(fileName)
        }
        break
      }
    }

    await uploadSupplementary(
      this.definition.metadataId!,
      'Transcription',
      'text/json',
      this.definition.supplementaryId
        ? this.definition.activity!.outputs['supplementaryId'] + this.definition.supplementaryId
        : this.definition.activity!.outputs['supplementaryId'],
      source.id,
      undefined,
      undefined,
      toArrayBuffer(JSON.stringify(transcribed)),
    )

    await uploadSupplementary(
      this.definition.metadataId!,
      'Transcription',
      'text/json',
      this.definition.supplementaryId
        ? this.definition.activity!.outputs['supplementaryTextId'] + this.definition.supplementaryId
        : this.definition.activity!.outputs['supplementaryTextId'],
      source.id,
      undefined,
      undefined,
      toArrayBuffer(transcribed!.text),
    )
  }
}
