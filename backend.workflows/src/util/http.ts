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
      duplex: 'half'
    })
    if (!response.ok) {
      console.error('failed to execute request: ' + await response.text())
    }
    return response
  } catch (e) {
    throw e
  }
}