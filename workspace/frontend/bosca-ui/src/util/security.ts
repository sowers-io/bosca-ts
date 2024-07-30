import {cookies} from "next/headers";

export function getAuthorizationToken(): string | undefined {
  const cookie = cookies().get('ory_kratos_session')
  if (cookie) {
    return cookie.value
  }
  return undefined
}

export function getHost() {
  if (process.env.NODE_ENV == 'development') {
    return 'http://localhost:4433'
  }
  return 'https://accounts.bosca.io'
}
