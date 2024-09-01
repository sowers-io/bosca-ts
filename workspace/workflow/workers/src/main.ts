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
import { newLoggingInterceptor } from '@bosca/common'
import { ConnectionOptions } from 'bullmq'
import { getActivities } from './activities'
import { start, newConfiguration } from '@bosca/workflow-worker-api'
import routes from './routes'
import fs from 'node:fs'

async function main() {
  const configuration = newConfiguration(fs.readFileSync(process.env.CONFIGURATION_FILE || 'configuration.json', 'utf8'))
  const connection: ConnectionOptions = {
    host: (process.env.BOSCA_REDIS_HOST || 'localhost'),
    port: parseInt(process.env.BOSCA_REDIS_PORT || '6379'),
  }
  const activities = getActivities()
  await start(connection, configuration, activities)
  const server = fastify({ http2: true })
  await server.register(fastifyConnectPlugin, {
    routes,
    interceptors: [newLoggingInterceptor()],
  })
  await server.listen({ host: '0.0.0.0', port: 7800 })
}

main().catch((e) => {
  console.error(e)
  process.exit(1)
})
