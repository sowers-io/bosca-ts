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

import { SubjectFinder } from '@bosca/common'
import { useServiceAccountClient } from '@bosca/common'
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
  subjectFinder: SubjectFinder,
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
    }),
  )
  if (result == null || !result.allowed) {
    throw new ConnectError('unauthorized', Code.PermissionDenied)
  }
}
