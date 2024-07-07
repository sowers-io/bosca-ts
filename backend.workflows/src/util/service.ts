import { IdRequest } from '../generated/protobuf/bosca/requests_pb'
import { Collection } from '../generated/protobuf/bosca/content/collections_pb'
import { useServiceClient } from './util'
import { ContentService } from '../generated/protobuf/bosca/content/content_connect'
import { Metadata } from '../generated/protobuf/bosca/content/metadata_pb'
import { SignedUrl } from '../generated/protobuf/bosca/content/url_pb'

export async function getCollection(id: IdRequest): Promise<Collection> {
  try {
    return await useServiceClient(ContentService).getCollection(id)
  } catch (e: any) {
    if (e.toString().indexOf('permission check failed') !== -1) {
      await new Promise((resolve) => setTimeout(resolve, 5))
      return await getCollection(id)
    }
    throw e
  }
}

export async function getMetadata(id: IdRequest): Promise<Metadata> {
  try {
    return await useServiceClient(ContentService).getMetadata(id)
  } catch (e: any) {
    if (e.toString().indexOf('permission check failed') !== -1) {
      await new Promise((resolve) => setTimeout(resolve, 5))
      return await getMetadata(id)
    }
    throw e
  }
}

export async function getMetadataUploadUrl(id: IdRequest): Promise<SignedUrl> {
  try {
    return await useServiceClient(ContentService).getMetadataUploadUrl(id)
  } catch (e: any) {
    if (e.toString().indexOf('permission check failed') !== -1) {
      await new Promise((resolve) => setTimeout(resolve, 5))
      return await getMetadataUploadUrl(id)
    }
    throw e
  }
}
