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

import { Activity, ActivityJobExecutor, Downloader, Job } from '@bosca/workflow-activities-api'
import { Book, ManifestName, PublicationContent, USXProcessor } from '@bosca/bible-processor'
import { ContentService, IdRequest, Metadata, Source, WorkflowJob } from '@bosca/protobufs'
import { useServiceAccountClient } from '@bosca/common'

export abstract class BookActivity extends Activity {
  readonly downloader: Downloader

  constructor(downloader: Downloader) {
    super()
    this.downloader = downloader
  }

  abstract newBookExecutor(job: Job, definition: WorkflowJob): BookExecutor

  newJobExecutor(job: Job, definition: WorkflowJob): ActivityJobExecutor<any> {
    return new Executor(this.newBookExecutor(job, definition), this, job, definition)
  }
}

export abstract class BookExecutor {

  protected readonly job: Job
  protected readonly definition: WorkflowJob

  constructor(job: Job, definition: WorkflowJob) {
    this.job = job
    this.definition = definition
  }

  abstract execute(
    source: Source,
    systemId: string,
    version: string,
    metadata: Metadata,
    book: Book
  ): Promise<void>
}

class Executor extends ActivityJobExecutor<BookActivity> {
  private readonly executor: BookExecutor

  constructor(executor: BookExecutor, activity: BookActivity, job: Job, definition: WorkflowJob) {
    super(activity, job, definition)
    this.executor = executor
  }

  async execute() {
    const contentService = useServiceAccountClient(ContentService)
    const metadata = await contentService.getMetadata(new IdRequest({ id: this.definition.metadataId }))
    const systemId = metadata.attributes['bible.system.id']
    const version = metadata.attributes['bible.version']
    const file = await this.activity.downloader.download(this.definition)
    try {
      const source = await contentService.getSource(new IdRequest({ id: 'workflow' }))
      const processor = new USXProcessor()
      const name = new ManifestName({ long: metadata.name })
      const content = new PublicationContent({ $: { role: metadata.attributes['bible.book.usfm'] } })
      const book = await processor.processBook(name, content, file)
      await this.executor.execute(source, systemId, version, metadata, book)
    } finally {
      await this.activity.downloader.cleanup(file)
    }
  }
}
