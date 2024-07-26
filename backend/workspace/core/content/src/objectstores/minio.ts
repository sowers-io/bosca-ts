import { ObjectStore } from './objectstore'
import { Metadata, MetadataSupplementary, SignedUrl, SignedUrlHeader } from '@bosca/protobufs'
import * as Minio from 'minio'
import { Code, ConnectError } from '@connectrpc/connect'

export class MinioObjectStore implements ObjectStore {
  private readonly client: Minio.Client
  private readonly bucket: string

  constructor() {
    const endpoint = process.env.BOSCA_S3_ENDPOINT!.split(':')
    this.bucket = process.env.BOSCA_S3_BUCKET || 'bosca'
    this.client = new Minio.Client({
      endPoint: endpoint[0],
      port: parseInt(endpoint[1]),
      useSSL: false,
      accessKey: process.env.BOSCA_S3_ACCESS_KEY_ID!,
      secretKey: process.env.BOSCA_S3_SECRET_ACCESS_KEY!,
    })
  }

  private getId(metadata: Metadata | MetadataSupplementary | string): string {
    if (metadata instanceof Metadata) {
      return metadata.id
    } else if (metadata instanceof MetadataSupplementary) {
      return metadata.metadataId + '.' + metadata.key
    } else {
      return metadata
    }
  }

  async createUploadUrl(metadata: Metadata | MetadataSupplementary): Promise<SignedUrl> {
    if (!metadata.contentLength) {
      throw new ConnectError('metadata does not have a content length', Code.FailedPrecondition)
    }
    const id = this.getId(metadata)
    const expires = new Date()
    expires.setMinutes(5)
    const policy = this.client.newPostPolicy()
    policy.setKey(id)
    policy.setBucket(this.bucket)
    policy.setContentType(metadata.contentType)
    policy.setContentLengthRange(Number(metadata.contentLength), Number(metadata.contentLength))
    policy.setExpires(expires)
    const url = await this.client.presignedPostPolicy(policy)
    const headers: SignedUrlHeader[] = []
    for (const key in url.formData) {
      headers.push(new SignedUrlHeader({ name: key, value: url.formData[key] }))
    }
    return new SignedUrl({
      id: id,
      headers: headers,
      method: 'POST',
      url: url.postURL,
    })
  }

  async createDownloadUrl(metadata: Metadata | MetadataSupplementary): Promise<SignedUrl> {
    const id = this.getId(metadata)
    const url = await this.client.presignedGetObject(this.bucket, id, 5 * 60)
    return new SignedUrl({ id: id, method: 'GET', url: url })
  }

  async delete(metadata: Metadata | MetadataSupplementary | string): Promise<void> {
    const id = this.getId(metadata)
    await this.client.removeObject(this.bucket, id)
  }
}
