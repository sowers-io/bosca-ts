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

import { Activity, ActivityJobExecutor, execute } from '@bosca/workflow-activities-api'
import { Job } from 'bullmq'
import { ContentService, IdRequest, SignedUrl, SupplementaryIdRequest, WorkflowJob } from '@bosca/protobufs'
import { getIStorageSystem, useServiceAccountClient } from '@bosca/common'

export class CreateTextEmbeddings extends Activity {
  get id(): string {
    return 'ai.embeddings.text'
  }

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, definition)
  }
}

class Executor extends ActivityJobExecutor<CreateTextEmbeddings> {
  async execute() {
    const idRequest = new IdRequest({ id: this.definition.metadataId })
    const service = useServiceAccountClient(ContentService)
    const metadata = await service.getMetadata(idRequest)
    let downloadUrl: SignedUrl
    if (metadata.contentType === 'text/plain' && !this.definition.activity?.inputs['supplementaryId']) {
      downloadUrl = await service.getMetadataDownloadUrl(idRequest)
    } else {
      downloadUrl = await service.getMetadataSupplementaryDownloadUrl(
        new SupplementaryIdRequest({
          id: this.definition.metadataId,
          key: this.definition.activity?.inputs['supplementaryId'] ?? 'text',
        }),
      )
    }
    const payload = await execute(downloadUrl)
    const system = await getIStorageSystem(this.definition.storageSystems[0].storageSystem!)
    await system.storeContent(this.definition, metadata, payload)
  }
}
