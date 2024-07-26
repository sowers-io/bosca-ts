import {
  BeginTransitionWorkflowRequest,
  IdResponse,
  IdResponsesId,
  Metadata,
  Permission,
  PermissionAction,
  PermissionObjectType,
  PermissionRelation,
  PermissionSubjectType,
  WorkflowService,
} from '@bosca/protobufs'
import { AdministratorGroup } from './permissions'
import { ContentDataSource } from '../../datasources/content'
import { PermissionManager, Subject, useServiceAccountClient } from '@bosca/common'
import { Code, ConnectError } from '@connectrpc/connect'
import { findNonUniqueId } from './collections'
import { StateProcessing } from './workflows'

export async function addMetadata(
  dataSource: ContentDataSource,
  permissions: PermissionManager,
  serviceAccountId: string,
  subject: Subject,
  parentId: string,
  metadata: Metadata
): Promise<IdResponse> {
  if (metadata.name.trim().length === 0) {
    throw new ConnectError('name must not be empty', Code.InvalidArgument)
  }
  if (parentId && parentId.length > 0) {
    const id = await findNonUniqueId(dataSource, parentId, metadata.name)
    if (id) {
      return new IdResponsesId({ id: id, error: 'name must be unique' })
    }
  }
  const id = await dataSource.addMetadata(metadata)
  await permissions.createRelationships(
    PermissionObjectType.metadata_type,
    newMetadataPermissions(serviceAccountId, subject.id, id)
  )
  if (parentId && parentId.length > 0) {
    await dataSource.addCollectionMetadataItem(parentId, id)
  }
  return new IdResponsesId({ id: id })
}

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

export async function setMetadataReady(
  dataSource: ContentDataSource,
  permissions: PermissionManager,
  subject: Subject,
  metadataId: string
) {
  const metadata = await dataSource.getMetadata(metadataId)
  if (!metadata) throw new ConnectError('missing metadata', Code.NotFound)
  await permissions.checkWithError(subject, PermissionObjectType.metadata_type, metadata.id, PermissionAction.manage)
  if (!metadata.traitIds || metadata.traitIds.length === 0) return
  const workflowService = useServiceAccountClient(WorkflowService)
  await workflowService.beginTransitionWorkflow(
    new BeginTransitionWorkflowRequest({
      metadataId: metadataId,
      stateId: StateProcessing,
    })
  )
}
