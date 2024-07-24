import { Permission, PermissionRelation, PermissionSubjectType } from '@bosca/protobufs'

export const AdministratorGroup = 'administrators'

export function newMetadataPermissions(serviceAccountId: string, userId: string, metadataId: string): Permission[] {
  return [
    new Permission({
      id: metadataId,
      subject: AdministratorGroup,
      subjectType: PermissionSubjectType.group,
      relation: PermissionRelation.owners,
    }),
    new Permission({
      id: metadataId,
      subject: serviceAccountId,
      subjectType: PermissionSubjectType.service_account,
      relation: PermissionRelation.owners,
    }),
    new Permission({
      id: metadataId,
      subject: userId,
      subjectType: PermissionSubjectType.user,
      relation: PermissionRelation.owners,
    }),
  ]
}
