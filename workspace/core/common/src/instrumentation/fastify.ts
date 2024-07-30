// TypeScript port from https://github.com/autotelic/fastify-opentelemetry

import fp from 'fastify-plugin'
import {
  Context,
  context,
  defaultTextMapGetter,
  defaultTextMapSetter,
  propagation,
  Span,
  SpanStatusCode,
  TextMapSetter,
  trace,
  Tracer,
} from '@opentelemetry/api'
import { FastifyPluginCallback, FastifyReply, FastifyRequest } from 'fastify'
import { FastifyInstance } from 'fastify/types/instance'
import { FastifyPluginOptions } from 'fastify/types/plugin'
import { TextMapGetter } from '@opentelemetry/api/build/src/propagation/TextMapPropagator'

export interface FastifyOpenTelemtryOptions extends FastifyPluginOptions {
  moduleName: string
  moduleVersion: string
}

export interface FastifyOpenTelemetry {
  get activeSpan(): Span | null

  get context(): Context

  get tracer(): Tracer

  inject(carrier: any, setter: TextMapSetter): void

  inject(carrier: any): void

  extract(carrier: any, getter: TextMapGetter): Context

  extract(carrier: any): Context
}

export interface OpenTelemetryFastifyRequest extends FastifyRequest {
  openTelemetry(): FastifyOpenTelemetry
}

function defaultFormatSpanName(request: FastifyRequest) {
  const { method } = request
  let path
  if (request.routeOptions) {
    path = request.routeOptions.url
  } else {
    path = request.routerPath
  }
  return path ? `${method} ${path}` : method
}

const defaultFormatSpanAttributes = {
  request(request: FastifyRequest) {
    return {
      'req.method': request.raw.method,
      'req.url': request.raw.url,
    }
  },
  reply(reply: FastifyReply) {
    return {
      'reply.statusCode': reply.statusCode,
    }
  },
  error(error: any) {
    return {
      'error.name': error.name,
      'error.message': error.message,
      'error.stack': error.stack,
    }
  },
}

async function openTelemetryPluginCallback(fastify: FastifyInstance, opts: FastifyPluginOptions = {}) {
  const {
    wrapRoutes,
    exposeApi = true,
    formatSpanName = defaultFormatSpanName,
    ignoreRoutes = [],
    propagateToReply = false,
  } = opts

  const shouldIgnoreRoute =
    typeof ignoreRoutes === 'function' ? ignoreRoutes : (path: any) => ignoreRoutes.includes(path)

  const formatSpanAttributes = {
    ...defaultFormatSpanAttributes,
    ...(opts.formatSpanAttributes || {}),
  }

  function getContext(request: FastifyRequest) {
    return contextMap.get(request) || context.active()
  }

  function openTelemetry() {
    // @ts-ignore
    const request: FastifyRequest = this as FastifyRequest
    return {
      get activeSpan() {
        return trace.getSpan(getContext(request))
      },
      get context() {
        return getContext(request)
      },
      get tracer() {
        return tracer
      },
      inject(carrier: any, setter = defaultTextMapSetter) {
        return propagation.inject(getContext(request), carrier, setter)
      },
      extract(carrier: any, getter = defaultTextMapGetter) {
        return propagation.extract(getContext(request), carrier, getter)
      },
    }
  }

  if (exposeApi) {
    fastify.decorateRequest('openTelemetry', openTelemetry)
  }

  const contextMap = new WeakMap()
  const tracer = trace.getTracer(opts.moduleName, opts.moduleVersion)

  async function onRequest(request: FastifyRequest) {
    if (shouldIgnoreRoute(request.url, request.method)) return

    let activeContext = context.active()

    // if not running within a local span then extract the context from the headers carrier
    if (!trace.getSpan(activeContext)) {
      activeContext = propagation.extract(activeContext, request.headers)
    }

    const span = tracer.startSpan(formatSpanName(request), {}, activeContext)
    span.setAttributes(formatSpanAttributes.request(request))
    contextMap.set(request, trace.setSpan(activeContext, span))
  }

  function onRequestWrapRoutes(request: FastifyRequest, reply: FastifyReply, done: (...args: any) => any) {
    if (
      !shouldIgnoreRoute(request.url, request.method) &&
      (wrapRoutes === true || (Array.isArray(wrapRoutes) && wrapRoutes.includes(request.url)))
    ) {
      context.with(getContext(request), done)
    } else {
      done()
    }
  }

  async function onResponse(request: FastifyRequest, reply: FastifyReply) {
    if (shouldIgnoreRoute(request.url, request.method)) return

    const activeContext = getContext(request)
    const span = trace.getSpan(activeContext)

    if (!span) {
      return
    }

    const spanStatus = { code: SpanStatusCode.OK }

    if (reply.statusCode >= 400) {
      spanStatus.code = SpanStatusCode.ERROR
    }

    span.setAttributes(formatSpanAttributes.reply(reply))
    span.setStatus(spanStatus)
    span.end()
    contextMap.delete(request)
  }

  async function onError(request: FastifyRequest, reply: FastifyReply, error: any) {
    if (shouldIgnoreRoute(request.url, request.method)) return

    const activeContext = getContext(request)
    const span = trace.getSpan(activeContext)

    if (!span) {
      return
    }

    span.setAttributes(formatSpanAttributes.error(error))
  }

  async function onSend(request: FastifyRequest, reply: FastifyReply, payload: any) {
    const { inject } = (request as OpenTelemetryFastifyRequest).openTelemetry()
    const propagationHeaders = {}
    inject(propagationHeaders)
    reply.headers(propagationHeaders)
    return payload
  }

  fastify.addHook('onRequest', onRequest)
  if (wrapRoutes) fastify.addHook('onRequest', onRequestWrapRoutes)
  fastify.addHook('onResponse', onResponse)
  fastify.addHook('onError', onError)
  if (propagateToReply) fastify.addHook('onSend', onSend)
}

export const openTelemetryPlugin: FastifyPluginCallback<FastifyOpenTelemtryOptions> = fp<FastifyOpenTelemtryOptions>(
  openTelemetryPluginCallback,
  {
    fastify: '4.x',
    name: 'fastify-opentelemetry-ts',
  }
)
