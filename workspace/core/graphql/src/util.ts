import { Permission, PermissionRelation, Permissions, PermissionSubjectType } from '@bosca/protobufs'
import {
  Permission as GPermission,
  PermissionRelation as GPermissionRelation,
  PermissionSubjectType as GPermissionSubjectType,
} from './generated/resolvers'

export function toGraphPermissions(id: string, permissions: Permissions): GPermission[] {
  return permissions.permissions.map((p) => {
    const j = p.toJson() as unknown as GPermission
    switch (p.subjectType) {
      case PermissionSubjectType.group:
        j.subjectType = GPermissionSubjectType.Group
        break
      case PermissionSubjectType.user:
        j.subjectType = GPermissionSubjectType.User
        break
      case PermissionSubjectType.service_account:
        j.subjectType = GPermissionSubjectType.Serviceaccount
        break
    }
    switch (p.relation) {
      case PermissionRelation.viewers:
        j.relation = GPermissionRelation.Viewers
        break
      case PermissionRelation.discoverers:
        j.relation = GPermissionRelation.Discoverers
        break
      case PermissionRelation.editors:
        j.relation = GPermissionRelation.Editors
        break
      case PermissionRelation.managers:
        j.relation = GPermissionRelation.Managers
        break
      case PermissionRelation.serviceaccounts:
        j.relation = GPermissionRelation.Serviceaccounts
        break
      case PermissionRelation.owners:
        j.relation = GPermissionRelation.Owners
        break
    }
    return j
  })
}

export function toGrpcPermissions(id: string, permissions: GPermission[]): Permissions {
  return new Permissions({
    id: id,
    permissions: permissions.map((p) => {
      let relation = PermissionRelation.viewers
      let subjectType = PermissionSubjectType.unknown_subject_type
      switch (p.relation) {
        case GPermissionRelation.Owners:
          relation = PermissionRelation.owners
          break
        case GPermissionRelation.Editors:
          relation = PermissionRelation.editors
          break
        case GPermissionRelation.Serviceaccounts:
          relation = PermissionRelation.serviceaccounts
          break
        case GPermissionRelation.Managers:
          relation = PermissionRelation.managers
          break
        case GPermissionRelation.Discoverers:
          relation = PermissionRelation.discoverers
          break
        case GPermissionRelation.Viewers:
          relation = PermissionRelation.viewers
          break
      }
      switch (p.subjectType) {
        case GPermissionSubjectType.User:
          subjectType = PermissionSubjectType.user
          break
        case GPermissionSubjectType.Group:
          subjectType = PermissionSubjectType.group
          break
        case GPermissionSubjectType.Serviceaccount:
          subjectType = PermissionSubjectType.service_account
          break
      }
      return new Permission({
        id: id,
        relation: relation,
        subject: p.subject,
        subjectType: subjectType,
      })
    }),
  })
}