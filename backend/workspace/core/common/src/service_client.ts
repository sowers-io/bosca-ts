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

import { createPromiseClient, Interceptor } from '@connectrpc/connect'
import { createGrpcTransport } from '@connectrpc/connect-node'
import { PromiseClient, Transport } from '@connectrpc/connect'
import { ServiceType } from '@bufbuild/protobuf'

const serviceAccountClients: { [typeName: string]: any } = {}
const clients: { [typeName: string]: any } = {}

function newTransport(typeName: string, serviceAuthorization: Interceptor | undefined = undefined): Transport {
  const interceptors: Interceptor[] = []
  if (serviceAuthorization) {
    interceptors.push(serviceAuthorization)
  }
  switch (typeName) {
    case 'bosca.content.ContentService':
      return createGrpcTransport({
        baseUrl: 'http://' + process.env.BOSCA_CONTENT_API_ADDRESS,
        httpVersion: '2',
        interceptors: interceptors,
      })
    case 'bosca.workflow.WorkflowService':
      return createGrpcTransport({
        baseUrl: 'http://' + process.env.BOSCA_WORKFLOW_API_ADDRESS,
        httpVersion: '2',
        interceptors: interceptors,
      })
    default:
      throw new Error('unsupported transport: ' + typeName)
  }
}

export function useServiceAccountClient<T extends ServiceType>(service: T): PromiseClient<T> {
  let client = serviceAccountClients[service.typeName]
  if (client != null) {
    return client as PromiseClient<T>
  }
  const serviceAuthorization: Interceptor = (next) => async (req) => {
    req.header.set('authorization', 'Token ' + process.env.BOSCA_SERVICE_ACCOUNT_TOKEN!)
    return await next(req)
  }
  const transport = newTransport(service.typeName, serviceAuthorization)
  client = createPromiseClient(service, transport)
  serviceAccountClients[service.typeName] = client
  return client
}

export function useClient<T extends ServiceType>(service: T): PromiseClient<T> {
  let client = clients[service.typeName]
  if (client != null) {
    return client as PromiseClient<T>
  }
  const transport = newTransport(service.typeName)
  client = createPromiseClient(service, transport)
  clients[service.typeName] = client
  return client
}
