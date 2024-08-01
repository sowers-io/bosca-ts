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

import { Activity, ActivityJobExecutor } from '../../activity'
import { Job } from 'bullmq'
import { ContentService, IdRequest, WorkflowJob } from '@bosca/protobufs'
import { useServiceAccountClient } from '@bosca/common/lib/service_client'
import { execute } from '../../../util/http'
import { getStorageSystem } from '@bosca/common/lib/storage/systems'

export class IndexText extends Activity {

  get id(): string {
    return 'metadata.text.index'
  }

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this, job, definition)
  }
}

class Executor extends ActivityJobExecutor<IndexText> {
  async execute() {
    const idRequest = new IdRequest({ id: this.definition.metadataId })
    const service = useServiceAccountClient(ContentService)
    const downloadUrl = await service.getMetadataDownloadUrl(idRequest)
    const payload = await execute(downloadUrl)
    const metadata = await service.getMetadata(idRequest)
    const system = await getStorageSystem(this.definition.storageSystems[0].configuration)
    await system.store(metadata, payload)
  }
}
