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

import { SignedUrl } from '../generated/protobuf/bosca/content/url_pb'

export function toArrayBuffer(value: string): ArrayBuffer {
  const enc = new TextEncoder()
  return enc.encode(value).buffer
}

export async function execute(signedUrl: SignedUrl, body?: BodyInit | null): Promise<Response> {
  const headers: { [key: string]: string } = {}
  for (const header of signedUrl.headers) {
    headers[header.name] = header.value
  }
  const url = signedUrl.url
  try {
    const response = await fetch(url, {
      method: signedUrl.method,
      headers: headers,
      body: body,
      redirect: 'follow',
      cache: 'no-cache',
      keepalive: false,
      // @ts-ignore
      duplex: 'half',
      signal: AbortSignal.timeout(10000)
    })
    if (!response.ok) {
      console.error('failed to execute request: ' + await response.text())
    }
    return response
  } catch (e) {
    throw e
  }
}