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

import { PermissionManager, Subject } from '@bosca/common'
import { PermissionAction, PermissionObjectType } from '@bosca/protobufs'

export const AdministratorGroup = 'administrators'

export type ValidIdMap = { [key: string]: boolean }

export async function toValidIds(
  subject: Subject,
  permissions: PermissionManager,
  resourceIdMap: ValidIdMap,
): Promise<ValidIdMap> {
  const resourceIds = Object.keys(resourceIdMap)
  const validIds = await permissions.bulkCheck(
    subject,
    PermissionObjectType.collection_type,
    resourceIds,
    PermissionAction.edit,
  )
  const validIdsMap: { [key: string]: boolean } = {}
  for (const validId of validIds) {
    validIdsMap[validId] = true
  }
  return validIdsMap
}
