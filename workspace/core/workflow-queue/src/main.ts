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
import routes from './service'
import { logger, newLoggingInterceptor } from '@bosca/common'

async function main() {
  const server = fastify({
    http2: true,
  })
  server.setErrorHandler((error, request, reply) => {
    logger.error({ error, request }, 'uncaught error')
    reply.status(500).send({ ok: false })
  })
  await server.register(fastifyConnectPlugin, {
    routes,
    interceptors: [newLoggingInterceptor()],
  })
  await server.listen({ host: '0.0.0.0', port: 7200 })
}

void main()
