import pino from 'pino'
import { Interceptor } from '@connectrpc/connect'

export const logger = pino({
  level: 'debug',
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
      logger.error({ error: e }, 'uncaught error')
      throw e
    }
  }
}