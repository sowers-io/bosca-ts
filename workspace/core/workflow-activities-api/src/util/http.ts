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

import { SignedUrl } from '@bosca/protobufs'
import * as http from 'node:http'
import { logger } from '@bosca/common'

let executing = 0

export function toArrayBuffer(value: string): ArrayBuffer {
  const enc = new TextEncoder()
  return enc.encode(value).buffer
}

export async function execute(signedUrl: SignedUrl, body?: ArrayBuffer | null): Promise<Buffer> {
  const headers: { [key: string]: string } = {}
  for (const header of signedUrl.headers) {
    headers[header.name] = header.value
  }
  const url = signedUrl.url
  if (!executing) {
    executing = 0
  }
  executing++
  try {
    logger.trace({ url, executing }, 'executing request')
    const buffer = await new Promise<Buffer>((resolve, reject) => {
      const options = {
        method: signedUrl.method,
        headers: headers,
      }
      const request = http.request(signedUrl.url, options, function (res) {
        if (!res.statusCode || res.statusCode < 200 || res.statusCode >= 300) {
          return reject(new Error('Request Failed: ' + res.statusCode))
        }
        const body: any = []
        res.on('data', function (chunk) {
          body.push(chunk)
        })
        res.on('end', function () {
          request.destroy()
          try {
            resolve(Buffer.concat(body))
          } catch (e) {
            reject(e)
          }
        })
      })
      request.on('error', function (err) {
        reject(err)
      })
      if (body) {
        request.write(Buffer.from(body))
      }
      request.end()

      setTimeout(() => {
        if (!request.destroyed) {
          logger.error({ url, executing }, 'request timeout')
          request.destroy(new Error('timeout'))
        }
      }, 15000)
    })
    executing--
    logger.trace({ url, executing }, 'request complete')
    return buffer
  } catch (e) {
    executing--
    logger.error({ url, executing, error: e }, 'request failed')
    throw e
  }
}
