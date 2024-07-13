import { useServiceClient } from './util'
import { ContentService } from '../generated/protobuf/bosca/content/service_connect'
import { Collection, FindCollectionRequest } from '../generated/protobuf/bosca/content/collections_pb'
import { Retry } from './retry'
import { FindMetadataRequest, Metadata } from '../generated/protobuf/bosca/content/metadata_pb'

export async function findAllCollections(attributes: { [key: string]: string }): Promise<Collection[]> {
  return Retry.execute(10, async () => {
    const result = await useServiceClient(ContentService).findCollection(
        new FindCollectionRequest({ attributes: attributes })
    )
    return result.collections
  })
}

export async function findFirstCollection(attributes: { [key: string]: string }): Promise<Collection> {
  return Retry.execute(10, async () => {
    const result = await useServiceClient(ContentService).findCollection(
        new FindCollectionRequest({ attributes: attributes })
    )
    if (result.collections.length === 0) {
      throw new Error('Collection not found')
    }
    return result.collections[0]
  })
}

export async function findAllMetadatas(attributes: { [key: string]: string }): Promise<Metadata[]> {
  return Retry.execute(10, async () => {
    const result = await useServiceClient(ContentService).findMetadata(
        new FindMetadataRequest({ attributes: attributes })
    )
    return result.metadata
  })
}

export async function findFirstMetadata(attributes: { [key: string]: string }): Promise<Metadata> {
  return Retry.execute(10, async () => {
    const result = await useServiceClient(ContentService).findMetadata(
        new FindMetadataRequest({ attributes: attributes })
    )
    if (result.metadata.length === 0) {
      throw new Error('Metadata not found')
    }
    return result.metadata
  })
}
