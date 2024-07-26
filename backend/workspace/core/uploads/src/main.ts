import { fastify } from 'fastify'
import { Server, FileKvStore } from '@tus/server'
import { S3Store } from '@tus/s3-store'
import { HttpSessionInterceptor, HttpSubjectFinder, useServiceAccountClient } from '@bosca/common'
import { verifyPermissions } from './authorization'
import { ContentService, IdRequest, Metadata } from '@bosca/protobufs'
import { protoInt64 } from '@bufbuild/protobuf'

async function main() {
  const server = fastify({
    logger: {
      level: 'debug',
    },
  })

  const sessionInterceptor = new HttpSessionInterceptor()
  const subjectFinder = new HttpSubjectFinder(
    process.env.BOSCA_SESSION_ENDPOINT!,
    process.env.BOSCA_SERVICE_ACCOUNT_ID!,
    process.env.BOSCA_SERVICE_ACCOUNT_TOKEN!,
    sessionInterceptor
  )

  const tusServer = new Server({
    path: '/files',
    datastore: new S3Store({
      s3ClientConfig: {
        bucket: process.env.BOSCA_BUCKET || 'bosca',
        region: process.env.BOSCA_REGION || 'us-east-1',
        endpoint: process.env.BOSCA_ENDPOINT || 'http://localhost:9000',
        credentials: {
          accessKeyId: process.env.BOSCA_ACCESS_KEY_ID!,
          secretAccessKey: process.env.BOSCA_SECRET_ACCESS_KEY!,
        },
      },
      cache: new FileKvStore(process.env.UPLOAD_DIR || '/tmp/uploads'),
    }),
    onUploadCreate: async (req, res) => {
      try {
        const url = new URL(req.url!)
        let collection = url.searchParams.get('collection')
        if (!collection) throw new Error('TODO: collection is required')
        if (req.headers.cookie) {
          await verifyPermissions(true, req.headers.cookie, collection, subjectFinder)
        } else if (req.headers.authorization) {
          await verifyPermissions(false, req.headers.authorization, collection, subjectFinder)
        }
      } catch (e) {
        res.statusCode = 401
      }
      return res
    },
    onUploadFinish: async (req, res, upload) => {
      const url = new URL(req.url!)
      let collection = url.searchParams.get('collection')
      if (!collection) throw new Error('TODO: collection is required')
      const traits = url.searchParams.get('trait')?.split(',')
      const metadata = new Metadata({
        name: upload.metadata!['name']!,
        contentType: upload.metadata!['filetype']!,
        traitIds: traits,
        contentLength: protoInt64.parse(upload.size!),
        sourceId: 'TODO',
        sourceIdentifier: upload.id,
      })
      const service = useServiceAccountClient(ContentService)
      const newMetadata = await service.addMetadata({
        collection: collection,
        metadata: metadata,
      })
      await service.setMetadataReady(new IdRequest({ id: newMetadata.id }))
      return res
    },
  })

  /**
   * add new content-type to fastify forewards request
   * without any parser to leave body untouched
   * @see https://www.fastify.io/docs/latest/Reference/ContentTypeParser/
   */
  server.addContentTypeParser('application/offset+octet-stream', (request, payload, done) => done(null))

  /**
   * let tus handle preparation and filehandling requests
   * fastify exposes raw nodejs http req/res via .raw property
   * @see https://www.fastify.io/docs/latest/Reference/Request/
   * @see https://www.fastify.io/docs/latest/Reference/Reply/#raw
   */
  server.all('/files', (req, res) => {
    // @ts-ignore
    tusServer.handle(req.raw, res.raw)
  })
  server.all('/files/*', (req, res) => {
    // @ts-ignore
    tusServer.handle(req.raw, res.raw)
  })

  await server.listen({ host: '0.0.0.0', port: 7001 })
  console.log('server is listening at', server.addresses()[0].address + ':' + server.addresses()[0].port)
}

void main()
