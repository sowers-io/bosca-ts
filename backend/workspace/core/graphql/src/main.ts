import { fastify } from 'fastify'
import { createSchema, createYoga } from 'graphql-yoga'
import { loadFiles } from '@graphql-tools/load-files'
import { RequestContext } from './context'

async function main() {
  const server = fastify({
    logger: true,
  })
  const schema = createSchema<RequestContext>({
    typeDefs: await loadFiles('src/schema/**/*.graphql'),
    resolvers: await loadFiles(['src/resolvers/*.ts', 'src/resolvers/**/*.ts']),
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
  await server.listen({ host: '0.0.0.0', port: 9000 })
  console.log('server is listening at', server.addresses()[0].address + ':' + server.addresses()[0].port)
}

void main()
