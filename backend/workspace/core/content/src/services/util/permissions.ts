import { PermissionManager, Subject } from '@bosca/common'
import { PermissionAction, PermissionObjectType } from '@bosca/protobufs'

export const AdministratorGroup = 'administrators'

export type ValidIdMap = { [key: string]: boolean }

export async function toValidIds(
  subject: Subject,
  permissions: PermissionManager,
  resourceIdMap: ValidIdMap
): Promise<ValidIdMap> {
  const resourceIds = Object.keys(resourceIdMap)
  const validIds = await permissions.bulkCheck(
    subject,
    PermissionObjectType.collection_type,
    resourceIds,
    PermissionAction.edit
  )
  const validIdsMap: { [key: string]: boolean } = {}
  for (const validId of validIds) {
    validIdsMap[validId] = true
  }
  return validIdsMap
}
