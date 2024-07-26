import { SubjectFinder } from '@bosca/common/lib/authentication/subject_finder'
import { useServiceAccountClient } from '@bosca/common/lib/service_client'
import {
  ContentService,
  PermissionAction,
  PermissionCheckRequest,
  PermissionObjectType,
  PermissionSubjectType,
} from '@bosca/protobufs'
import { SubjectType } from '@bosca/common'
import { Code, ConnectError } from '@connectrpc/connect'

export async function verifyPermissions(
  fromCookie: boolean,
  authorization: string,
  collection: string,
  subjectFinder: SubjectFinder
) {
  const subject = await subjectFinder.findSubject(fromCookie, authorization)
  let subjectPermissionType = PermissionSubjectType.user
  if (subject.type == SubjectType.serviceaccount) {
    subjectPermissionType = PermissionSubjectType.service_account
  }
  const result = await useServiceAccountClient(ContentService).checkPermission(
    new PermissionCheckRequest({
      subject: subject.id,
      subjectType: subjectPermissionType,
      object: collection,
      objectType: PermissionObjectType.collection_type,
      action: PermissionAction.edit,
    })
  )
  if (result == null || !result.allowed) {
    throw new ConnectError('unauthorized', Code.PermissionDenied)
  }
}
