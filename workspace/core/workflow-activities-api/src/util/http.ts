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
  executing++
  try {
    logger.trace({ url, executing }, 'executing request')
    const response = await fetch(url, {
      method: signedUrl.method,
      headers: headers,
      body: body ? Buffer.from(body) : undefined,
    })
    if (!response.ok) {
      throw new Error(`Request failed: ${response.status}: ${await response.text()}`)
    }
    const responseBody = await response.arrayBuffer()
    executing--
    logger.trace({ url, executing }, 'request complete')
    return Buffer.from(responseBody)
  } catch (e) {
    executing--
    logger.error({ url, executing, error: e }, 'request failed')
    throw e
  }
}
