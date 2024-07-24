import { Metadata, SignedUrl } from '@bosca/protobufs'

export interface ObjectStore {
  createUploadUrl(metadata: Metadata): Promise<SignedUrl>

  createDownloadUrl(id: string): Promise<SignedUrl>

  delete(id: string): Promise<void>
}
