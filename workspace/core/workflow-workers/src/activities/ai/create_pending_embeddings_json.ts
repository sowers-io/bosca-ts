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

import { Activity, ActivityJobExecutor } from '../activity'
import { Job } from 'bullmq'
import { useServiceAccountClient } from '@bosca/common'
import {
  ContentService,
  WorkflowJob,
  SupplementaryIdRequest,
  PendingEmbedding,
  PendingEmbeddings,
  IdRequest,
} from '@bosca/protobufs'
import { execute } from '../../util/http'
import { uploadSupplementary } from '../../util/uploader'

export class CreatePendingEmbeddingsFromJsonTable extends Activity {
  get id(): string {
    return 'ai.embeddings.pending.from-json-table'
  }

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, definition)
  }
}

class Executor extends ActivityJobExecutor<CreatePendingEmbeddingsFromJsonTable> {
  async execute() {
    const service = useServiceAccountClient(ContentService)
    const source = await service.getSource(new IdRequest({ id: 'workflow' }))
    const key = this.definition.supplementaryId
        ? this.definition.activity!.inputs!['supplementaryId'] + this.definition.supplementaryId
        : this.definition.activity!.inputs!['supplementaryId']
    const signedUrl = await service.getMetadataSupplementaryDownloadUrl(
      new SupplementaryIdRequest({
        id: this.definition.metadataId,
        key: key,
      })
    )
    const payload = await execute(signedUrl)
    const table = JSON.parse(payload.toString())
    const idColumn = this.definition.activity!.configuration['idColumn']!
    const contentColumn = this.definition.activity!.configuration['contentColumn']!
    const embeddings: PendingEmbedding[] = []
    for (const row of table) {
      let content = row[contentColumn]
      if (Array.isArray(content)) {
        content = content.join(', ')
      }
      if (content.trim().length === 0) {
        continue
      }
      embeddings.push(
        new PendingEmbedding({
          id: row[idColumn],
          content: content,
        })
      )
    }
    const buffer = new PendingEmbeddings({ embedding: embeddings }).toBinary()
    await uploadSupplementary(
      this.definition.metadataId!,
      'Pending Embeddings',
      'application/protobuf',
      this.definition.supplementaryId
        ? this.definition.activity!.outputs['supplementaryId'] + this.definition.supplementaryId
        : this.definition.activity!.outputs['supplementaryId'],
      source.id,
      undefined,
      undefined,
      buffer
    )
  }
}
