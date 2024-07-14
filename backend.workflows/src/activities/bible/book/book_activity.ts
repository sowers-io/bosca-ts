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

import { Activity } from '../../activity'
import { Book, ManifestName, PublicationContent, USXProcessor } from '@bosca/bible'
import { Downloader } from '../../../util/downloader'
import { useServiceClient } from '../../../util/util'
import { ContentService } from '../../../generated/protobuf/bosca/content/service_connect'
import { IdRequest } from '../../../generated/protobuf/bosca/requests_pb'
import { WorkflowActivityJob } from '../../../generated/protobuf/bosca/workflow/execution_context_pb'
import { Source } from '../../../generated/protobuf/bosca/content/sources_pb'
import { Metadata } from '../../../generated/protobuf/bosca/content/metadata_pb'

export abstract class BookActivity extends Activity {
  private readonly downloader: Downloader

  constructor(downloader: Downloader) {
    super()
    this.downloader = downloader
  }

  abstract executeBook(
    source: Source,
    systemId: string,
    metadata: Metadata,
    activity: WorkflowActivityJob,
    book: Book
  ): Promise<void>

  async execute(activity: WorkflowActivityJob) {
    const contentService = useServiceClient(ContentService)
    const metadata = await contentService.getMetadata(new IdRequest({ id: activity.metadataId }))
    const systemId = metadata.attributes['bible.system.id']
    const file = await this.downloader.download(activity)
    try {
      const source = await contentService.getSource(new IdRequest({ id: 'workflow' }))
      const processor = new USXProcessor()
      const name = new ManifestName({ short: metadata.name })
      const content = new PublicationContent({ $: { role: metadata.attributes['bible.book.usfm'] } })
      const book = await processor.processBook(name, content, file)
      await this.executeBook(source, systemId, metadata, activity, book)
    } finally {
      await this.downloader.cleanup(file)
    }
  }
}
