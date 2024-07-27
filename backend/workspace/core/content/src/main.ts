import { fastify } from 'fastify'
import { fastifyConnectPlugin } from '@connectrpc/connect-fastify'
import {
  HttpSessionInterceptor,
  HttpSubjectFinder,
  newAuthenticationInterceptor,
  newLoggingInterceptor,
} from '@bosca/common'
import routes from './services/routes'
import { logger } from '@bosca/common'

async function main() {
  const grpcServer = fastify({
    http2: true,
    logger: {
      level: 'debug',
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
  await grpcServer.register(fastifyConnectPlugin, {
    routes,
    interceptors: [newLoggingInterceptor(), newAuthenticationInterceptor(subjectFinder)],
  })
  await grpcServer.listen({ host: '0.0.0.0', port: 7000 })
}

void main()
