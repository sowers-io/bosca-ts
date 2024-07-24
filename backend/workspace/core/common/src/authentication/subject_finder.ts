import { SubjectType } from '../permissions/permissions'

export interface Subject {
  get id(): string

  get type(): SubjectType
}

export interface SubjectFinder {
  findSubject(fromCookie: boolean, authorization: string): Promise<Subject>
}

export interface SessionInterceptor {
  getSubjectId(response: Response): Promise<string | null>
}