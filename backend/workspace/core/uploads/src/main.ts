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

import { fastify } from 'fastify'
import { Server, FileKvStore } from '@tus/server'
import { S3Store } from '@tus/s3-store'
import { HttpSessionInterceptor, HttpSubjectFinder, logger, useServiceAccountClient } from '@bosca/common'
import { verifyPermissions } from './authorization'
import { ContentService, IdRequest, Metadata } from '@bosca/protobufs'
import { protoInt64 } from '@bufbuild/protobuf'
import { Code, ConnectError } from '@connectrpc/connect'
import http, { ServerResponse } from 'node:http'

async function main() {
  const server = fastify({
    logger: {
      level: 'debug',
    },
  })
  server.setErrorHandler((error, request, reply) => {
    logger.error({ error, request }, 'uncaught error')
    reply.status(500).send({ ok: false })
  })

  const sessionInterceptor = new HttpSessionInterceptor()
  const subjectFinder = new HttpSubjectFinder(
    process.env.BOSCA_SESSION_ENDPOINT!,
    process.env.BOSCA_SERVICE_ACCOUNT_ID!,
    process.env.BOSCA_SERVICE_ACCOUNT_TOKEN!,
    sessionInterceptor
  )

  function onError(res: ServerResponse<http.IncomingMessage>, e: any, rethrow: boolean) {
    if (e instanceof ConnectError) {
      switch (e.code) {
        case Code.Unauthenticated:
          res.statusCode = 401
          break
        case Code.PermissionDenied:
          res.statusCode = 403
          break
        default:
          if (rethrow) {
            throw e
          }
          res.statusCode = 500
      }
    } else {
      if (rethrow) {
        throw e
      }
      res.statusCode = 500
    }
  }

  const tusServer = new Server({
    path: '/files',
    datastore: new S3Store({
      s3ClientConfig: {
        bucket: process.env.BOSCA_S3_BUCKET || 'bosca',
        region: process.env.BOSCA_S3_REGION || 'us-east-1',
        endpoint: process.env.BOSCA_S3_ENDPOINT || 'http://127.0.0.1:9010',
        credentials: {
          accessKeyId: process.env.BOSCA_S3_ACCESS_KEY_ID!,
          secretAccessKey: process.env.BOSCA_S3_SECRET_ACCESS_KEY!,
        },
      },
      cache: new FileKvStore(process.env.UPLOAD_DIR || '/tmp/uploads'),
    }),
    onUploadCreate: async (req, res, upload) => {
      try {
        let collection = upload.metadata!['collection']
        if (!collection) {
          collection = '00000000-0000-0000-0000-000000000000'
        }
        if (req.headers.cookie) {
          await verifyPermissions(true, req.headers.cookie, collection, subjectFinder)
        } else if (req.headers.authorization) {
          await verifyPermissions(false, req.headers.authorization, collection, subjectFinder)
        }
      } catch (e) {
        onError(res, e, true)
      }
      return res
    },
    onResponseError: async (req, res, error) => {
      logger.error({ error }, 'failed to upload')
      onError(res, error, false)
    },
    onUploadFinish: async (req, res, upload) => {
      try {
        let collection = upload.metadata!['collection']
        if (!collection) {
          collection = '00000000-0000-0000-0000-000000000000'
        }
        const service = useServiceAccountClient(ContentService)
        const source = await service.getSource(new IdRequest({ id: 'uploader' }))
        let traits = upload.metadata!['traits']?.split(',') || []
        const metadata = new Metadata({
          name: upload.metadata!['name']!,
          contentType: upload.metadata!['filetype']!,
          traitIds: traits,
          contentLength: protoInt64.parse(upload.size!),
          languageTag: upload.metadata!['language'] || 'en',
          sourceId: source.id,
          sourceIdentifier: upload.id,
        })
        const newMetadata = await service.addMetadata({
          collection: collection,
          metadata: metadata,
        })
        await service.setMetadataReady(new IdRequest({ id: newMetadata.id }))
      } catch (e) {
        onError(res, e, true)
      }
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
}

void main()
