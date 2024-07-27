import {
  Collection,
  CollectionItem,
  CollectionItems,
  IdResponse,
  IdResponsesId,
  Permission,
  PermissionAction,
  PermissionObjectType,
  PermissionRelation,
  PermissionSubjectType,
  WorkflowExecutionRequest,
  WorkflowService,
} from '@bosca/protobufs'
import { logger, PermissionManager, Subject, useServiceAccountClient } from '@bosca/common'
import { ContentDataSource, IdName } from '../../datasources/content'
import { Code, ConnectError } from '@connectrpc/connect'

export async function getCollectionItems(
  dataSource: ContentDataSource,
  permissions: PermissionManager,
  subject: Subject,
  collectionId: string
): Promise<CollectionItems> {
  // TODO: paging?
  let collectionItemIds = await dataSource.getCollectionCollectionItemIds(collectionId)
  collectionItemIds = await permissions.bulkCheck(
    subject,
    PermissionObjectType.collection_type,
    collectionItemIds,
    PermissionAction.view
  )
  let metadataItemIds = await dataSource.getCollectionMetadataItemIds(collectionId)
  metadataItemIds = await permissions.bulkCheck(
    subject,
    PermissionObjectType.metadata_type,
    metadataItemIds,
    PermissionAction.view
  )
  const items: CollectionItem[] = []
  for (const id of collectionItemIds) {
    const collection = await dataSource.getCollection(id)
    if (!collection) {
      logger.error({ collectionId, id }, 'failed to get collection')
      continue
    }
    items.push(new CollectionItem({ Item: { case: 'collection', value: collection } }))
  }
  for (const id of metadataItemIds) {
    const metadata = await dataSource.getMetadata(id)
    if (!metadata) {
      logger.error({ collectionId, metadataId: id }, 'failed to get metadata')
      continue
    }
    items.push(new CollectionItem({ Item: { case: 'metadata', value: metadata } }))
  }
  return new CollectionItems({ items: items })
}

export async function findNonUniqueId(
  dataSource: ContentDataSource,
  parentId: string,
  name: string
): Promise<string | null> {
  let found: IdName[]
  found = await dataSource.getCollectionIdName(parentId, name)
  if (found.length > 0) {
    return found[0].id
  }
  found = await dataSource.getMetadataIdName(parentId, name)
  if (found.length > 0) {
    return found[0].id
  }
  return null
}

export async function addCollection(
  dataSource: ContentDataSource,
  permissions: PermissionManager,
  serviceAccountId: string,
  subject: Subject,
  parentId: string,
  collection: Collection
): Promise<IdResponse> {
  if (collection.name.trim().length === 0) {
    throw new ConnectError('name must not be empty', Code.InvalidArgument)
  }
  if (parentId && parentId.length > 0) {
    const id = await findNonUniqueId(dataSource, parentId, collection.name)
    if (id) {
      return new IdResponsesId({ id: id, error: 'name must be unique' })
    }
  }
  const id = await dataSource.addCollection(collection)
  const newPermissions = newCollectionPermissions(serviceAccountId, subject.id, id)
  await permissions.createRelationships(PermissionObjectType.collection_type, newPermissions)
  if (parentId && parentId.length) {
    await dataSource.addCollectionCollectionItem(parentId, id)
  }
  await permissions.waitForPermissions(PermissionObjectType.collection_type, newPermissions)
  return new IdResponsesId({ id: id })
}

export function newCollectionPermissions(serviceAccountId: string, userId: string, id: string): Permission[] {
  return [
    new Permission({
      id: id,
      subject: 'administrators',
      subjectType: PermissionSubjectType.group,
      relation: PermissionRelation.owners,
    }),
    new Permission({
      id: id,
      subject: serviceAccountId,
      subjectType: PermissionSubjectType.service_account,
      relation: PermissionRelation.serviceaccounts,
    }),
    new Permission({
      id: id,
      subject: userId,
      subjectType: PermissionSubjectType.user,
      relation: PermissionRelation.owners,
    }),
  ]
}

export async function setCollectionReady(
  dataSource: ContentDataSource,
  permissions: PermissionManager,
  subject: Subject,
  collectionId: string
) {
  const collection = await dataSource.getCollection(collectionId)
  if (!collection) throw new ConnectError('missing collection', Code.NotFound)
  await permissions.checkWithError(
    subject,
    PermissionObjectType.collection_type,
    collection.id,
    PermissionAction.manage
  )
  if (!collection.traitIds || collection.traitIds.length === 0) return
  const workflowService = useServiceAccountClient(WorkflowService)
  for (const traitId of collection.traitIds) {
    const workflowIds = await dataSource.getTraitWorkflowIds(traitId)
    for (const workflowId of workflowIds) {
      await workflowService.executeWorkflow(
        new WorkflowExecutionRequest({
          collectionId: collection.id,
          workflowId: workflowId,
        })
      )
    }
  }
}
