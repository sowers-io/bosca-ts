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

import pino from 'pino'
import { Code, ConnectError, Interceptor } from '@connectrpc/connect'

export const logger = pino({
  level: process.env.NODE_ENV === 'production' ? 'info' : 'debug',
  serializers: {
    err: pino.stdSerializers.err,
    error: pino.stdSerializers.err,
  },
})

export function newLoggingInterceptor(): Interceptor {
  return (next) => async (req) => {
    try {
      return await next(req)
    } catch (e) {
      if (e instanceof ConnectError) {
        if (e.code == Code.NotFound) {
          logger.trace({ error: e }, 'connect error')
        } else {
          logger.debug({ error: e }, 'connect error')
        }
      } else {
        logger.error({ error: e }, 'uncaught error')
      }
      throw e
    }
  }
}
