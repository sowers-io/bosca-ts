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

import { GraphQLError } from 'graphql/error'
import { ConnectError, Code } from '@connectrpc/connect'
import { GraphQLRequestContext } from './graphql'
import { logger } from '../logger'

export function getAuthenticationToken(context: GraphQLRequestContext): string | null {
  // @ts-ignore
  const authorization = context.request.headers.headersInit!['authorization']
  if (authorization && authorization.length > 0) {
    if (authorization.startsWith('Bearer ')) {
      return authorization.toString().substring('Bearer '.length)
    }
  }
  return null
}

export function getGraphQLHeaders(context: GraphQLRequestContext): Record<string, string> {
  const headers: Record<string, string> = {}
  // @ts-ignore
  const authorization = context.request.headers.headersInit!['authorization']
  if (authorization && authorization.length > 0) {
    headers['Authorization'] = authorization
  }
  return headers
}

export async function executeGraphQL<T>(fn: () => Promise<T>): Promise<T> {
  try {
    return await fn()
  } catch (e: any) {
    logger.error({ error: e }, 'failed to execute graphql request')
    if (e instanceof ConnectError) {
      if (e.code == Code.NotFound) {
        return null as T
      }
    }
    throw new GraphQLError(e.message, {
      originalError: e,
    })
  }
}
