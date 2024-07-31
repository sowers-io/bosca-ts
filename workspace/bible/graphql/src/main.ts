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
import { createSchema, createYoga } from 'graphql-yoga'
import { loadFiles, LoadFilesOptions } from '@graphql-tools/load-files'
import { RequestContext } from './context'
import { logger } from '@bosca/common'
import url from 'url'

async function main() {
  const server = fastify({
    logger: {
      level: process.env.NODE_ENV === 'production' ? 'info' : 'debug',
    },
  })
  const options: LoadFilesOptions = {
    ignoreIndex: true,
    requireMethod: async (path: any) => {
      return await import(url.pathToFileURL(path).toString());
    },
  }
  const schema = createSchema<RequestContext>({
    typeDefs: await loadFiles('src/schema/**/*.graphql', options),
    resolvers: await loadFiles(['src/resolvers/*.ts', 'src/resolvers/**/*.ts'], options),
  })
  const yoga = createYoga<RequestContext>({
    schema: schema,
    logging: {
      debug: (...args) => args.forEach((arg) => server.log.debug(arg)),
      info: (...args) => args.forEach((arg) => server.log.info(arg)),
      warn: (...args) => args.forEach((arg) => server.log.warn(arg)),
      error: (...args) => args.forEach((arg) => server.log.error(arg)),
    },
  })
  server.route({
    url: yoga.graphqlEndpoint,
    method: ['GET', 'POST', 'OPTIONS'],
    handler: async (request, reply) => {
      const response = await yoga.handleNodeRequestAndResponse(request, reply, { request: request, reply: reply })
      response.headers.forEach((value, key) => {
        reply.header(key, value)
      })
      reply.status(response.status)
      reply.send(response.body)
      return reply
    },
  })
  await server.listen({ host: '0.0.0.0', port: 2000 })
}

void main()
