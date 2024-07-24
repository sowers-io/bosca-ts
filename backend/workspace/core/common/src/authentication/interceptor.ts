import type { Interceptor } from '@connectrpc/connect'
import { createContextKey } from '@connectrpc/connect'
import { SubjectType } from '../permissions/permissions'
import { SubjectFinder, Subject } from './subject_finder'
import { StreamRequest, UnaryRequest } from '@connectrpc/connect/dist/cjs/interceptor'

export const SubjectKey = createContextKey<Subject>(
  {
    id: 'anonymous',
    type: SubjectType.user,
  } // Default value
)

export function newAuthenticationInterceptor(subjectFinder: SubjectFinder): Interceptor {
  async function authenticate(
    req: UnaryRequest | StreamRequest,
    fromCookie: boolean,
    authorization: string | null
  ): Promise<boolean> {
    if (authorization && authorization.length > 0) {
      const subject = await subjectFinder.findSubject(fromCookie, authorization)
      req.contextValues.set(SubjectKey, subject)
      return true
    }
    return false
  }

  return (next) => async (req) => {
    if (await authenticate(req, false, req.header.get('Authorization'))) {
      return await next(req)
    }
    if (await authenticate(req, false, req.header.get('X-Service-Authorization'))) {
      return await next(req)
    }
    if (await authenticate(req, true, req.header.get('Cookie'))) {
      return await next(req)
    }
    return await next(req)
  }
}
