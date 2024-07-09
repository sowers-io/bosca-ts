import { USXProcessor, BibleMetadata, Book } from '@bosca/bible'
import {
  AddCollectionRequest,
  AddCollectionsRequest,
  Collection
} from '../../generated/protobuf/bosca/content/collections_pb'
import { IdRequest } from '../../generated/protobuf/bosca/requests_pb'
import { useServiceClient } from '../../util/util'
import { ContentService } from '../../generated/protobuf/bosca/content/service_connect'
import { AddMetadataRequest, AddMetadatasRequest, Metadata } from '../../generated/protobuf/bosca/content/metadata_pb'
import { execute, toArrayBuffer } from '../../util/http'
import { protoInt64 } from '@bufbuild/protobuf'
import { Downloader } from '../../util/downloader'
import { getCollection, getMetadata, getMetadataUploadUrl } from '../../util/service'
import { WorkflowActivityJob } from '../../generated/protobuf/bosca/workflow/execution_context_pb'
import { Activity } from '../activity'

export class ProcessBibleActivity extends Activity {

  private readonly downloader: Downloader

  constructor(downloader: Downloader) {
    super()
    this.downloader = downloader
  }

  get id(): string {
    return 'bible.process'
  }

  private async createBibleCollection(metadata: BibleMetadata): Promise<Collection> {
    const service = useServiceClient(ContentService)
    const addResponse = await service.addCollection(new AddCollectionRequest({
      collection: new Collection({
        name: metadata.identification.nameLocal,
        attributes: {
          'bible.system.id': metadata.identification.systemId.id,
          'bible.abbreviation': metadata.identification.abbreviationLocal
        }
      })
    }))
    return getCollection(new IdRequest({ id: addResponse.id }))
  }

  private async createBookCollections(workflowId: string, metadata: BibleMetadata, bible: Collection, books: Book[]): Promise<Collection[]> {
    const service = useServiceClient(ContentService)
    const addCollectionRequests: AddCollectionRequest[] = []
    const addMetadatasRequests: AddMetadataRequest[] = []
    const buffers: ArrayBuffer[] = []
    let order = 0

    // build bulk requests
    for (const book of books) {
      const attributes = {
        'bible.system.id': bible.attributes['bible.system.id'],
        'bible.abbreviation': bible.attributes['bible.abbreviation'],
        'bible.usfm': book.usfm,
        'bible.book.order': order.toString()
      }
      addCollectionRequests.push(new AddCollectionRequest({
        parent: bible.id,
        collection: new Collection({
          name: book.name.short + ' Chapters',
          attributes: attributes
        })
      }))
      const buffer = toArrayBuffer(book.raw)
      buffers.push(buffer)
      addMetadatasRequests.push(new AddMetadataRequest({
        collection: bible.id,
        metadata: new Metadata({
          name: book.name.short,
          contentType: 'bible/usx-book',
          languageTag: metadata.language.iso,
          contentLength: protoInt64.parse(buffer.byteLength),
          attributes: attributes,
          sourceId: 'workflow',
          sourceIdentifier: workflowId
        })
      }))
      order++
    }

    // create metadata
    const addMetadataResponses = await service.addMetadatas(new AddMetadatasRequest({
      metadatas: addMetadatasRequests
    }))

    // upload books
    for (let bookIndex = 0; bookIndex < books.length; bookIndex++) {
      const addResponse = addMetadataResponses.id[bookIndex]
      if (addResponse.error) {
        throw new Error(addResponse.error)
      }
      const idRequest = new IdRequest({ id: addResponse.id })
      const uploadUrl = await getMetadataUploadUrl(idRequest)
      const uploadResponse = await execute(uploadUrl, buffers[bookIndex])
      if (!uploadResponse.ok) {
        throw new Error('failed to upload book: ' + books[bookIndex].usfm + ': ' + await uploadResponse.text())
      }
      await service.setMetadataUploaded(idRequest)
    }

    // create collections
    const addCollectionResponses = await service.addCollections(new AddCollectionsRequest({
      collections: addCollectionRequests
    }))

    // fetch created collections
    const collections: Collection[] = []
    for (const addResponse of addCollectionResponses.id) {
      if (addResponse.error) {
        throw new Error(addResponse.error)
      }
      // TODO: Add bulk getCollections
      const collection = await getCollection(new IdRequest({ id: addResponse.id }))
      collections.push(collection)
    }

    return collections
  }

  private async createChapters(workflowId: string, metadata: BibleMetadata, bookCollection: Collection, book: Book): Promise<Metadata[]> {
    const service = useServiceClient(ContentService)
    const requests: AddMetadataRequest[] = []
    const buffers: ArrayBuffer[] = []
    let order = 0
    for (const chapter of book.chapters) {
      const buffer = toArrayBuffer(book.raw.substring(chapter.position.start, chapter.position.end))
      buffers.push(buffer)
      requests.push(new AddMetadataRequest({
        collection: bookCollection.id,
        metadata: new Metadata({
          name: book.name.short + ' ' + chapter.number,
          contentType: 'bible/usx-chapter',
          contentLength: protoInt64.parse(buffer.byteLength),
          languageTag: metadata.language.iso,
          attributes: {
            'bible.system.id': bookCollection.attributes['bible.system.id'],
            'bible.abbreviation': bookCollection.attributes['bible.abbreviation'],
            'bible.book.usfm': book.usfm,
            'bible.chapter.usfm': chapter.usfm,
            'bible.chapter.order': (order++).toString()
          },
          sourceId: 'workflow',
          sourceIdentifier: workflowId
        })
      }))
    }
    const response = await service.addMetadatas(new AddMetadatasRequest({
      metadatas: requests
    }))
    const metadatas: Metadata[] = []
    for (let chapterIndex = 0; chapterIndex < response.id.length; chapterIndex++) {
      const addResponse = response.id[chapterIndex]
      if (addResponse.error) {
        throw new Error(addResponse.error)
      }
      const idRequest = new IdRequest({ id: addResponse.id })
      const metadata = await getMetadata(idRequest)
      const uploadUrl = await getMetadataUploadUrl(idRequest)
      const uploadResponse = await execute(uploadUrl, buffers[chapterIndex])
      if (!uploadResponse.ok) {
        throw new Error('failed to upload chapter: ' + book.chapters[chapterIndex].usfm + ': ' + await uploadResponse.text())
      }
      await service.setMetadataUploaded(idRequest)
      metadatas.push(metadata)
    }
    return metadatas
  }

  async execute(activity: WorkflowActivityJob) {
    const file = await this.downloader.download(activity)
    try {
      const processor = new USXProcessor()
      await processor.process(file)
      const bibleCollection = await this.createBibleCollection(processor.metadata)
      const bookCollections = await this.createBookCollections(activity.workflowId, processor.metadata, bibleCollection, processor.books)

      for (let bookIndex = 0; bookIndex < processor.books.length; bookIndex++) {
        const book = processor.books[bookIndex]
        const collection = bookCollections[bookIndex]

        await this.createChapters(activity.workflowId, processor.metadata, collection, book)
      }

    } finally {
      await this.downloader.cleanup(file)
    }
  }
}
