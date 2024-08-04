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

import type { Interceptor } from '@connectrpc/connect'
import { createContextKey } from '@connectrpc/connect'
import { SubjectType } from '../permissions/permissions'
import { SubjectFinder, Subject } from './subject_finder'
import { StreamRequest, UnaryRequest } from '@connectrpc/connect/dist/cjs/interceptor'

export const SubjectKey = createContextKey<Subject>(
  {
    id: 'anonymous',
    type: SubjectType.user,
  }, // Default value
)

export function newAuthenticationInterceptor(subjectFinder: SubjectFinder): Interceptor {
  async function authenticate(
    req: UnaryRequest | StreamRequest,
    fromCookie: boolean,
    authorization: string | null,
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
