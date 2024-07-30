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
import { fastifyConnectPlugin } from '@connectrpc/connect-fastify'
import {
  HttpSessionInterceptor,
  HttpSubjectFinder,
  newAuthenticationInterceptor,
  newLoggingInterceptor,
  openTelemetryPlugin,
} from '@bosca/common'
import routes from './services/routes'
import { logger } from '@bosca/common'

async function main() {
  const grpcServer = fastify({
    http2: true,
    logger: {
      level: process.env.NODE_ENV === 'production' ? 'info' : 'debug',
    },
  })
  grpcServer.setErrorHandler((error, request, reply) => {
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
  await grpcServer.register(openTelemetryPlugin)
  await grpcServer.register(fastifyConnectPlugin, {
    routes,
    interceptors: [newLoggingInterceptor(), newAuthenticationInterceptor(subjectFinder)],
  })
  await grpcServer.listen({ host: '0.0.0.0', port: 7000 })
}

void main()
