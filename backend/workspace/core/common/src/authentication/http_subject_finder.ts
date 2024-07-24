import { SessionInterceptor, Subject, SubjectFinder } from './subject_finder'
import { PermissionError, SubjectType } from '../permissions/permissions'

export class HttpSubjectFinder implements SubjectFinder {
  private readonly endpoint: string
  private readonly serviceAccountId: string
  private readonly serviceAccountTokenHeader: string
  private readonly interceptor: SessionInterceptor

  constructor(
    endpoint: string,
    serviceAccountId: string,
    serviceAccountToken: string,
    interceptor: SessionInterceptor
  ) {
    this.endpoint = endpoint
    this.serviceAccountId = serviceAccountId
    this.serviceAccountTokenHeader = 'Token ' + serviceAccountToken
    this.interceptor = interceptor
  }

  async findSubject(fromCookie: boolean, authorization: string): Promise<Subject> {
    if (authorization == this.serviceAccountTokenHeader) {
      return { id: this.serviceAccountId, type: SubjectType.serviceaccount }
    }
    const response = await fetch(this.endpoint, {
      method: 'GET',
      headers: {
        ...(fromCookie ? { Cookie: authorization } : { Authorization: authorization }),
      },
    })
    if (!response.ok) {
      throw new PermissionError('failed to get session: ' + response.status)
    }
    const subjectId = await this.interceptor.getSubjectId(response)
    if (!subjectId) {
      return { id: 'anonymous', type: SubjectType.user }
    }
    return { id: subjectId, type: SubjectType.user }
  }
}

export class HttpSessionInterceptor implements SessionInterceptor {
  async getSubjectId(response: Response): Promise<string | null> {
    const session = await response.json()
    if (!session || !session.identity || !session.identity.id) {
      return null
    }
    return session.identity.id
  }
}
