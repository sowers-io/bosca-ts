import { ObjectStore } from './objectstore'
import { Metadata, SignedUrl, SignedUrlHeader } from '@bosca/protobufs'
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

  async createUploadUrl(metadata: Metadata): Promise<SignedUrl> {
    if (!metadata.contentLength) {
      throw new ConnectError('metadata does not have a content length', Code.FailedPrecondition)
    }
    const expires = new Date()
    expires.setMinutes(5)
    const policy = this.client.newPostPolicy()
    policy.setKey(metadata.id)
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
      id: metadata.id,
      headers: headers,
      method: 'POST',
      url: url.postURL,
    })
  }

  async createDownloadUrl(id: string): Promise<SignedUrl> {
    const url = await this.client.presignedGetObject(this.bucket, id, 5 * 60)
    return new SignedUrl({ id: id, method: 'GET', url: url })
  }

  async delete(id: string): Promise<void> {
    await this.client.removeObject(this.bucket, id)
  }
}
