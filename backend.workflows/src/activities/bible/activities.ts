import { WorkflowActivityExecutionContext } from '../../generated/protobuf/bosca/content/workflows_pb'
import { Activity } from '../../workflow/workflow'
import { USXProcessor, BibleMetadata, Book } from '@bosca/bible'
import {
  AddCollectionRequest,
  AddCollectionsRequest,
  Collection
} from '../../generated/protobuf/bosca/content/collections_pb'
import { IdRequest } from '../../generated/protobuf/bosca/requests_pb'
import { useServiceClient } from '../../util/util'
import { ContentService } from '../../generated/protobuf/bosca/content/content_connect'
import { AddMetadataRequest, AddMetadatasRequest, Metadata } from '../../generated/protobuf/bosca/content/metadata_pb'
import { execute, toArrayBuffer } from '../../util/http'
import { protoInt64 } from '@bufbuild/protobuf'
import { SignedUrl } from '../../generated/protobuf/bosca/content/url_pb'

export interface ProcessBibleDownloader {

  download(executionContext: WorkflowActivityExecutionContext): Promise<string>

  cleanup(file: string): Promise<void>
}

export class ProcessBibleActivity extends Activity {

  private readonly downloader: ProcessBibleDownloader

  constructor(downloader: ProcessBibleDownloader) {
    super()
    this.downloader = downloader
  }

  get id(): string {
    return 'bible.process'
  }

  async createBibleCollection(metadata: BibleMetadata): Promise<Collection> {
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
    return this.getCollection(new IdRequest({ id: addResponse.id }))
  }

  async createBookCollections(metadata: BibleMetadata, bible: Collection, books: Book[]): Promise<Collection[]> {
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
          attributes: attributes
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
      const uploadUrl = await this.getMetadataUploadUrl(idRequest)
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
      const collection = await this.getCollection(new IdRequest({ id: addResponse.id }))
      collections.push(collection)
    }

    return collections
  }

  async getCollection(id: IdRequest): Promise<Collection> {
    try {
      return await useServiceClient(ContentService).getCollection(id)
    } catch (e: any) {
      if (e.toString().indexOf('permission check failed') !== -1) {
        await new Promise((resolve) => setTimeout(resolve, 5))
        return await this.getCollection(id)
      }
      throw e
    }
  }

  async getMetadata(id: IdRequest): Promise<Metadata> {
    try {
      return await useServiceClient(ContentService).getMetadata(id)
    } catch (e: any) {
      if (e.toString().indexOf('permission check failed') !== -1) {
        await new Promise((resolve) => setTimeout(resolve, 5))
        return await this.getMetadata(id)
      }
      throw e
    }
  }

  async getMetadataUploadUrl(id: IdRequest): Promise<SignedUrl> {
    try {
      return await useServiceClient(ContentService).getMetadataUploadUrl(id)
    } catch (e: any) {
      if (e.toString().indexOf('permission check failed') !== -1) {
        await new Promise((resolve) => setTimeout(resolve, 5))
        return await this.getMetadataUploadUrl(id)
      }
      throw e
    }
  }

  async createChapters(metadata: BibleMetadata, bookCollection: Collection, book: Book): Promise<Metadata[]> {
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
            'bible.usfm': chapter.usfm,
            'bible.chapter.order': (order++).toString()
          }
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
      const metadata = await this.getMetadata(idRequest)
      const uploadUrl = await this.getMetadataUploadUrl(idRequest)
      const uploadResponse = await execute(uploadUrl, buffers[chapterIndex])
      if (!uploadResponse.ok) {
        throw new Error('failed to upload chapter: ' + book.chapters[chapterIndex].usfm + ': ' + await uploadResponse.text())
      }
      await service.setMetadataUploaded(idRequest)
      metadatas.push(metadata)
    }
    return metadatas
  }

  async execute(executionContext: WorkflowActivityExecutionContext) {
    const file = await this.downloader.download(executionContext)
    try {
      const processor = new USXProcessor()
      await processor.process(file)
      const bibleCollection = await this.createBibleCollection(processor.metadata)
      const bookCollections = await this.createBookCollections(processor.metadata, bibleCollection, processor.books)

      for (let bookIndex = 0; bookIndex < processor.books.length; bookIndex++) {
        const book = processor.books[bookIndex]
        const collection = bookCollections[bookIndex]

        await this.createChapters(processor.metadata, collection, book)
      }

    } finally {
      await this.downloader.cleanup(file)
    }
  }
}

// export const create = (client: PromiseClient<typeof ContentService>) => ({
//   'bible.process': async (context: WorkflowActivityExecutionContext): Promise<void> => {
//
//   },
//   'bible.chapters.create': async (context: WorkflowActivityExecutionContext): Promise<void> => {
//
//   },
//   'bible.chapter.verses.create': async (context: WorkflowActivityExecutionContext): Promise<void> => {
//
//   },
//   'bible.chapter.verses.table': async (context: WorkflowActivityExecutionContext): Promise<void> => {
//
//   }
// })
