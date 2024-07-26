import { Metadata, MetadataSupplementary, SignedUrl } from '@bosca/protobufs'

export interface ObjectStore {
  createUploadUrl(metadata: Metadata | MetadataSupplementary): Promise<SignedUrl>

  createDownloadUrl(metadata: Metadata | MetadataSupplementary): Promise<SignedUrl>

  delete(metadata: Metadata | MetadataSupplementary | string): Promise<void>
}
