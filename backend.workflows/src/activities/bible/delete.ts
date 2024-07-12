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

import { Activity } from '../activity'
import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import { useServiceClient } from '../../util/util'
import { ContentService } from '../../generated/protobuf/bosca/content/service_connect'
import { FindMetadataRequest } from '../../generated/protobuf/bosca/content/metadata_pb'
import { IdRequest } from '../../generated/protobuf/bosca/requests_pb'
import { USXProcessor } from '@bosca/bible'
import { Downloader } from '../../util/downloader'

export class DeleteBibleActivity extends Activity {

  private downloader: Downloader

  constructor(downloader: Downloader) {
    super()
    this.downloader = downloader
  }

  get id(): string {
    return 'bible.delete'
  }

  async execute(activity: WorkflowActivityJob) {
    const file = await this.downloader.download(activity)
    try {
      const processor = new USXProcessor()
      await processor.process(file)
      const contentService = useServiceClient(ContentService)
      const metadatas = await contentService.findMetadata(new FindMetadataRequest({
        attributes: {
          'bible.system.id': processor.metadata.identification.systemId.id
        }
      }))
      for (const metadata of metadatas.metadata) {
        await contentService.deleteMetadata(new IdRequest({ id: metadata.id }))
      }
      const collections = await contentService.findCollection(new FindMetadataRequest({
        attributes: {
          'bible.system.id': processor.metadata.identification.systemId.id
        }
      }))
      for (const collection of collections.collections) {
        await contentService.deleteCollection(new IdRequest({ id: collection.id }))
      }
    } finally {
      await this.downloader.cleanup(file)
    }
  }
}