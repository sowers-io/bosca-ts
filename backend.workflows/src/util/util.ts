import { createPromiseClient, Interceptor } from '@connectrpc/connect'
import { createGrpcTransport } from '@connectrpc/connect-node'
import { PromiseClient, Transport } from '@connectrpc/connect'
import { ServiceType } from '@bufbuild/protobuf'

const clients: { [typeName: string]: any } = {}

export function useServiceClient<T extends ServiceType>(service: T): PromiseClient<T> {
  let client = clients[service.typeName]
  if (client != null) {
    return client as PromiseClient<T>
  }
  const serviceAuthorization: Interceptor = (next) => async (req) => {
    req.header.set('authorization', 'Token ' + process.env.BOSCA_SERVICE_ACCOUNT_TOKEN)
    return await next(req)
  }
  let transport: Transport
  switch (service.typeName) {
    case 'bosca.content.ContentService':
      transport = createGrpcTransport({
        baseUrl: 'http://' + process.env.BOSCA_CONTENT_API_ADDRESS,
        httpVersion: '2',
        interceptors: [serviceAuthorization]
      })
      break
    case 'bosca.workflow.WorkflowService':
      transport = createGrpcTransport({
        baseUrl: 'http://' + process.env.BOSCA_WORKFLOW_API_ADDRESS,
        httpVersion: '2',
        interceptors: [serviceAuthorization]
      })
      break
    default:
      throw new Error('unsupported transport: ' + service.typeName)
  }
  client = createPromiseClient(service, transport)
  clients[service.typeName] = client
  return client
}