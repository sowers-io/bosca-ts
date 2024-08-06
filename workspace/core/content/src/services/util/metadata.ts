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

import {
  BeginTransitionWorkflowRequest,
  IdResponse,
  IdResponsesId,
  Metadata,
  Permission,
  PermissionAction,
  PermissionObjectType,
  PermissionRelation,
  PermissionSubjectType, WorkflowExecutionRequest,
  WorkflowService,
} from '@bosca/protobufs'
import { AdministratorGroup } from './permissions'
import { ContentDataSource } from '../../datasources/content'
import { PermissionManager, StateProcessing, Subject, useServiceAccountClient } from '@bosca/common'
import { Code, ConnectError } from '@connectrpc/connect'
import { findNonUniqueId } from './collections'

export async function addMetadata(
  dataSource: ContentDataSource,
  permissions: PermissionManager,
  serviceAccountId: string,
  subject: Subject,
  parentId: string,
  metadata: Metadata,
): Promise<IdResponse> {
  if (metadata.name.trim().length === 0) {
    throw new ConnectError('name must not be empty', Code.InvalidArgument)
  }
  if (parentId && parentId.length > 0) {
    const id = await findNonUniqueId(dataSource, parentId, metadata.name)
    if (id && id != metadata.id) {
      return new IdResponsesId({ id: id, error: 'name must be unique' })
    }
  }
  const { metadataId, version } = await dataSource.addMetadata(metadata)
  if (version === 1) {
    const newPermissions = newMetadataPermissions(serviceAccountId, subject.id, metadataId)
    await permissions.createRelationships(PermissionObjectType.metadata_type, newPermissions)
    if (parentId && parentId.length > 0) {
      await dataSource.addCollectionMetadataItem(parentId, metadataId)
    }
    await permissions.waitForPermissions(PermissionObjectType.metadata_type, newPermissions)
  }
  return new IdResponsesId({ id: metadataId })
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
      relation: PermissionRelation.serviceaccounts,
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
  metadataId: string,
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
    }),
  )
}

export async function setMetadataSupplementaryReady(
  dataSource: ContentDataSource,
  permissions: PermissionManager,
  subject: Subject,
  metadataId: string,
  supplementaryId: string,
) {
  const supplementary = await dataSource.getMetadataSupplementary(metadataId, supplementaryId)
  if (!supplementary) throw new ConnectError('missing metadata supplementary', Code.NotFound)
  await permissions.checkWithError(subject, PermissionObjectType.metadata_type, supplementary.metadataId, PermissionAction.manage)
  await dataSource.setMetadataSupplementaryReady(metadataId, supplementaryId)
  if (!supplementary.traitIds || supplementary.traitIds.length === 0) return
  const workflowService = useServiceAccountClient(WorkflowService)
  for (const traitId of supplementary.traitIds) {
    const workflowIds = await dataSource.getTraitWorkflowIds(traitId)
    for (const workflowId of workflowIds) {
      await workflowService.executeWorkflow(
        new WorkflowExecutionRequest({
          workflowId: workflowId,
          metadataId: metadataId,
          supplementaryId: supplementaryId,
        }),
      )
    }
  }
}

export async function setWorkflowStateComplete(
  subject: Subject,
  dataSource: ContentDataSource,
  metadataId: string,
  status: string,
) {
  const metadata = await dataSource.getMetadata(metadataId)
  if (!metadata) {
    throw new ConnectError('missing metadata', Code.NotFound)
  }
  let state = metadata.workflowStateId
  if (metadata.workflowStatePendingId) {
    state = metadata.workflowStatePendingId
  }
  await dataSource.setWorkflowState(subject, metadata.id, metadata.workflowStateId, state, status, true, true)
}
