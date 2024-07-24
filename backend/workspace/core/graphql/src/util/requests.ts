import { GraphQLError } from 'graphql/error'
import { ConnectError, Code } from '@connectrpc/connect'
import { RequestContext } from '../context'

export function getHeaders(context: RequestContext): Record<string, string> {
  const headers: Record<string, string> = {}
  // @ts-expect-error
  const authorization = context.request.headers.headersInit!['authorization']
  if (authorization && authorization.length > 0) {
    headers['Authorization'] = authorization
  }
  return headers
}

export async function execute<T>(fn: () => Promise<T>): Promise<T | null> {
  try {
    return await fn()
  } catch (e: any) {
    if (e instanceof ConnectError) {
      if (e.code == Code.NotFound) {
        return null
      }
    }
    throw new GraphQLError(e.message, {
      originalError: e
    })
  }
}
