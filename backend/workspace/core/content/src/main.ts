import { fastify } from 'fastify'
import { fastifyConnectPlugin } from '@connectrpc/connect-fastify'
import routes from './service'
import { HttpSessionInterceptor, HttpSubjectFinder, newAuthenticationInterceptor } from '@bosca/common'

async function main() {
  const grpcServer = fastify({
    http2: true,
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
  const authenticationInterceptor = newAuthenticationInterceptor(subjectFinder)
  await grpcServer.register(fastifyConnectPlugin, {
    routes,
    interceptors: [authenticationInterceptor],
  })
  await grpcServer.listen({ host: '0.0.0.0', port: 7000 })
  console.log('server is listening at', grpcServer.addresses()[0].address + ':' + grpcServer.addresses()[0].port)
}

void main()
