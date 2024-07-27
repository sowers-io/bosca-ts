import {
  Permissions,
  Permission,
  PermissionAction,
  PermissionObjectType,
  PermissionSubjectType,
} from '@bosca/protobufs'
import { Subject } from '../authentication/subject_finder'
import { Code, ConnectError } from '@connectrpc/connect'

export enum SubjectType {
  user = 'user',
  group = 'group',
  serviceaccount = 'serviceaccount',
}

export class PermissionError extends ConnectError {
  constructor(message: string) {
    super(message, Code.PermissionDenied)
  }
}

export interface PermissionManager {
  bulkCheck(
    subject: Subject,
    objectType: PermissionObjectType,
    resourceId: string[],
    action: PermissionAction
  ): Promise<string[]>

  checkWithError(
    subject: Subject,
    objectType: PermissionObjectType,
    resourceId: string,
    action: PermissionAction
  ): Promise<void>

  checkWithSubjectIdError(
    subjectType: PermissionSubjectType,
    subjectId: string,
    objectType: PermissionObjectType,
    resourceId: string,
    action: PermissionAction
  ): Promise<void>

  createRelationships(objectType: PermissionObjectType, permissions: Permission[]): Promise<void>

  createRelationship(objectType: PermissionObjectType, permission: Permission): Promise<void>

  waitForPermissions(objectType: PermissionObjectType, permissions: Permission[]): Promise<void>

  getPermissions(objectType: PermissionObjectType, resourceId: string): Promise<Permissions>
}
