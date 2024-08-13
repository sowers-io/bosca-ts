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

import { fastify } from 'fastify'
import {
  openTelemetryPlugin,
} from '@bosca/common'
import { logger } from '@bosca/common'
import sharp from 'sharp'
import { Readable } from 'node:stream'
import { encode } from 'blurhash'

interface QueryOpts {
  u: string | undefined
  w: string | undefined
  h: string | undefined
  f: string | undefined
  ch: string | undefined
  cw: string | undefined
  q: string | undefined
}

type Resize = {
  width: number | undefined
  height: number | undefined
}

async function main() {
  const supportedUrls = process.env.SUPPORTED_URLS?.split(',').map((u) => new RegExp(u)) || []
  const server = fastify()
  server.setErrorHandler((error, request, reply) => {
    logger.error({ error, request }, 'uncaught error')
    reply.status(500).send({ ok: false })
  })
  await server.register(openTelemetryPlugin)
  server.get('/', {}, async function (request, reply) {
    const opts = request.query as QueryOpts
    if (!opts.u) {
      reply.code(400).send()
      return
    }
    const imageUrl = new URL(opts.u)
    let isSupported = false
    for (const supported of supportedUrls) {
      if (supported.test(opts.u)) {
        isSupported = true
        break
      }
    }
    if (!isSupported) {
      reply.code(401).send()
      return
    }
    const response = await fetch(imageUrl)
    if (!response.ok || !response.body) {
      reply.code(500).send()
      return
    }
    let transformer = sharp()
    // @ts-ignore
    if (opts.w || opts.h) {
      let resize: Resize = { width: undefined, height: undefined }
      if (opts.w) {
        resize.width = parseInt(opts.w)
        if (isNaN(resize.width)) {
          return reply.code(400).send()
        }
      }
      if (opts.h) {
        resize.height = parseInt(opts.h)
        if (isNaN(resize.height)) {
          return reply.code(400).send()
        }
      }
      transformer = transformer.resize(resize)
    }
    switch (opts.f) {
      case 'blurhash': {
        transformer = transformer.raw().ensureAlpha()
        // @ts-ignore
        const result = await Readable.fromWeb(response.body).pipe(transformer).toBuffer({ resolveWithObject: true })
        const img = Uint8ClampedArray.from(result.data)
        const cWidth = opts.cw ? parseInt(opts.cw) : 4
        const cHeight = opts.ch ? parseInt(opts.ch) : 4
        const blurhash = encode(img, result.info.width, result.info.height, cWidth, cHeight)
        return reply
          .header('Content-Type', 'text/plain')
          .send(blurhash)
      }
      case 'jpeg': {
        const quality = opts.q ? parseInt(opts.q) : 80
        if (isNaN(quality)) {
          return reply.code(400).send()
        }
        transformer = transformer.toFormat(sharp.format.jpeg, {
          quality: quality,
        })
        break
      }
    }

    // @ts-ignore
    await reply.send(Readable.fromWeb(response.body).pipe(transformer))
  })
  await server.listen({ host: '0.0.0.0', port: 8002 })
  logger.info('server listening on 0.0.0.0:8002')
}

void main()
