// noinspection JSUnusedGlobalSymbols

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
  Collections,
  ContentService,
  Empty,
  IdRequest,
  IdResponse,
  IdResponses,
  IdResponsesId,
  MetadataRelationships,
  Metadatas,
  MetadataSupplementaries,
  MetadataSupplementary,
  Permission,
  PermissionAction,
  PermissionCheckResponse,
  PermissionObjectType,
  PermissionSubjectType,
  SignedUrl,
  Sources,
  Traits,
  WorkflowService,
} from '@bosca/protobufs'
import { Code, ConnectError, type ConnectRouter } from '@connectrpc/connect'
import {
  logger,
  PermissionManager,
  StateProcessing,
  SubjectKey,
  SubjectType,
  useServiceAccountClient,
} from '@bosca/common'
import { ContentDataSource, RootCollectionId } from '../datasources/content'
import { ObjectStore } from '../objectstores/objectstore'
import { addCollection, getCollectionItems, setCollectionReady } from './util/collections'
import { addMetadata, setMetadataReady, setMetadataSupplementaryReady, setWorkflowStateComplete } from './util/metadata'
import { toValidIds } from './util/permissions'

export function content(
  router: ConnectRouter,
  serviceAccountId: string,
  permissions: PermissionManager,
  dataSource: ContentDataSource,
  objectStore: ObjectStore,
): ConnectRouter {
  return router.service(ContentService, {
    async getSources() {
      return new Sources({ sources: await dataSource.getSources() })
    },
    async getSource(request) {
      const source = await dataSource.getSource(request.id)
      if (!source) throw new ConnectError('missing source', Code.NotFound)
      return source
    },
    async getTraits() {
      return new Traits({ traits: await dataSource.getTraits() })
    },
    async getTrait(request) {
      const trait = await dataSource.getTrait(request.id)
      if (!trait) throw new ConnectError('missing trait', Code.NotFound)
      return trait
    },
    async getRootCollectionItems(_, context) {
      const subject = context.values.get(SubjectKey)
      return getCollectionItems(dataSource, permissions, subject, RootCollectionId)
    },
    async getCollectionItems(request, context) {
      const subject = context.values.get(SubjectKey)
      return getCollectionItems(dataSource, permissions, subject, request.id)
    },
    async addCollections(request, context) {
      const resourceIdMap: { [key: string]: boolean } = {}
      for (const collection of request.collections) {
        resourceIdMap[collection.parent] = true
      }
      const subject = context.values.get(SubjectKey)
      const validIdsMap = await toValidIds(subject, permissions, resourceIdMap)
      const ids: IdResponsesId[] = []
      for (const addRequest of request.collections) {
        try {
          if (!addRequest.collection) {
            // noinspection ExceptionCaughtLocallyJS
            throw new ConnectError('missing collection', Code.InvalidArgument)
          }
          if (!validIdsMap[addRequest.parent]) {
            // noinspection ExceptionCaughtLocallyJS
            throw new ConnectError('permission check failed', Code.PermissionDenied)
          }
          ids.push(
            await addCollection(
              dataSource,
              permissions,
              serviceAccountId,
              subject,
              addRequest.parent,
              addRequest.collection,
            ),
          )
        } catch (error: any) {
          logger.error({ error }, 'error adding collection')
          ids.push(new IdResponsesId({ error: error.message }))
        }
      }
      return new IdResponses({ id: ids })
    },
    async addCollection(request, context) {
      const subject = context.values.get(SubjectKey)
      if (!request.collection) throw new ConnectError('missing collection', Code.InvalidArgument)
      let parent = RootCollectionId
      if (request.parent && request.parent.length > 0) {
        parent = request.parent
      }
      await permissions.checkWithError(subject, PermissionObjectType.collection_type, parent, PermissionAction.edit)
      const id = await addCollection(
        dataSource,
        permissions,
        serviceAccountId,
        subject,
        parent,
        request.collection,
      )
      return new IdResponse({ id: id.id })
    },
    async deleteCollection(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.collection_type,
        request.id,
        PermissionAction.delete,
      )
      await dataSource.deleteCollection(request.id)
      return new Empty()
    },
    async getCollectionPermissions(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.collection_type,
        request.id,
        PermissionAction.manage,
      )
      return await permissions.getPermissions(PermissionObjectType.collection_type, request.id)
    },
    async addCollectionPermission(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.collection_type,
        request.id,
        PermissionAction.manage,
      )
      await permissions.createRelationship(
        PermissionObjectType.collection_type,
        new Permission({
          id: request.id,
          subject: request.subject,
          relation: request.relation,
          subjectType: request.subjectType,
        }),
      )
      return new Empty()
    },
    async addCollectionItem(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.collection_type,
        request.collectionId,
        PermissionAction.manage,
      )
      switch (request.itemId.case) {
        case 'childCollectionId':
          await dataSource.addCollectionItemId(request.collectionId, request.itemId.value, null)
          break
        case 'childMetadataId':
          await dataSource.addCollectionItemId(request.collectionId, null, request.itemId.value)
          break
      }
    },
    async setCollectionReady(request, context) {
      const subject = context.values.get(SubjectKey)
      await setCollectionReady(dataSource, permissions, subject, request.id)
      return new Empty()
    },
    async getCollection(request, context) {
      const collection = await dataSource.getCollection(request.id)
      if (!collection) throw new ConnectError('missing collection', Code.NotFound)
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.collection_type,
        collection.id,
        PermissionAction.view,
      )
      return collection
    },
    async findCollection(request, context) {
      const collections = await dataSource.findCollection(request.attributes)
      const collectionIds = collections.map((m) => m.id)
      const subject = context.values.get(SubjectKey)
      const validIds = await permissions.bulkCheck(
        subject,
        PermissionObjectType.collection_type,
        collectionIds,
        PermissionAction.view,
      )
      const validIdsMap: { [key: string]: boolean } = {}
      for (const id of validIds) {
        validIdsMap[id] = true
      }
      return new Collections({
        collections: collections.filter((c) => validIdsMap[c.id]),
      })
    },
    async checkPermission(request) {
      try {
        let type = SubjectType.user
        if (request.subjectType === PermissionSubjectType.service_account) {
          type = SubjectType.serviceaccount
        } else if (request.subjectType === PermissionSubjectType.group) {
          type = SubjectType.group
        }
        await permissions.checkWithError(
          {
            id: request.subject,
            type,
          },
          request.objectType,
          request.object,
          request.action,
        )
        return new PermissionCheckResponse({ allowed: true })
      } catch (_) {
        return new PermissionCheckResponse({ allowed: false })
      }
    },
    async addMetadatas(request, context) {
      const resourceIdMap: { [key: string]: boolean } = {}
      for (const metadata of request.metadatas) {
        let collection = RootCollectionId
        if (metadata.collection && metadata.collection.length > 0) {
          collection = metadata.collection
        }
        resourceIdMap[collection] = true
      }
      const subject = context.values.get(SubjectKey)
      const validIdsMap = await toValidIds(subject, permissions, resourceIdMap)
      const ids: IdResponsesId[] = []
      for (const addRequest of request.metadatas) {
        try {
          if (!addRequest.metadata) continue
          let collection = RootCollectionId
          if (addRequest.collection && addRequest.collection.length > 0) {
            collection = addRequest.collection
          }
          if (!validIdsMap[collection]) {
            // noinspection ExceptionCaughtLocallyJS
            throw new ConnectError('permission check failed', Code.PermissionDenied)
          }
          ids.push(
            await addMetadata(dataSource, permissions, serviceAccountId, subject, collection, addRequest.metadata),
          )
        } catch (error: any) {
          logger.error({ error }, 'error adding metadata')
          ids.push(new IdResponsesId({ error: error.message }))
        }
      }
      return new IdResponses({ id: ids })
    },
    async addMetadataAttributes(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.id,
        PermissionAction.edit,
      )
      await dataSource.addMetadataAttributes(request.id, request.attributes)
      const metadata = await dataSource.getMetadata(request.id)
      if (!metadata) {
        throw new ConnectError('missing metadata', Code.NotFound)
      }
      return metadata
    },
    async addMetadataTrait(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.metadataId,
        PermissionAction.manage,
      )
      const metadata = await dataSource.getMetadata(request.metadataId)
      if (!metadata) {
        throw new ConnectError('missing metadata', Code.NotFound)
      }
      await dataSource.addMetadataTrait(request.metadataId, request.traitId)
      metadata.traitIds.push(request.traitId)
      await useServiceAccountClient(WorkflowService).beginTransitionWorkflow(
        new BeginTransitionWorkflowRequest({
          metadataId: request.metadataId,
          status: 'adding trait: ' + request.traitId,
          stateId: StateProcessing,
        }),
      )
      return metadata
    },
    async setMetadataTraits(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.metadataId,
        PermissionAction.manage,
      )
      const metadata = await dataSource.getMetadata(request.metadataId)
      if (!metadata) {
        throw new ConnectError('missing metadata', Code.NotFound)
      }

      // TODO: use transaction

      let removeIds = new Set(metadata.traitIds)
      const traitIds: string[] = []
      for (const id of request.traitId) {
        if (!metadata.traitIds.includes(id)) {
          await dataSource.addMetadataTrait(request.metadataId, id)
        }
        removeIds.delete(id)
        traitIds.push(id)
      }
      metadata.traitIds = traitIds

      for (const id of removeIds) {
        await dataSource.deleteMetadataTrait(request.metadataId, id)
      }

      await useServiceAccountClient(WorkflowService).beginTransitionWorkflow(
        new BeginTransitionWorkflowRequest({
          metadataId: request.metadataId,
          status: 'setting traits: ' + request.traitId.join(','),
          stateId: StateProcessing,
        }),
      )
      return metadata
    },
    async addMetadata(request, context) {
      if (!request.metadata) throw new ConnectError('missing metadata', Code.InvalidArgument)
      let collection = RootCollectionId
      if (request.collection && request.collection.length > 0) {
        collection = request.collection
      }
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.collection_type, collection, PermissionAction.edit)
      const id = await addMetadata(dataSource, permissions, serviceAccountId, subject, collection, request.metadata)
      return new IdResponse({ id: id.id })
    },
    async deleteMetadata(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, request.id, PermissionAction.delete)
      const metadata = await dataSource.getMetadata(request.id)
      if (!metadata) {
        throw new ConnectError('missing metadata', Code.NotFound)
      }
      await objectStore.delete(metadata)
      if (metadata.sourceIdentifier) {
        await objectStore.delete(metadata.sourceIdentifier.split('+')[0])
      }
      await dataSource.deleteMetadata(request.id)
      return new Empty()
    },
    async getMetadata(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, request.id, PermissionAction.view)
      const metadata = await dataSource.getMetadata(request.id)
      if (!metadata) throw new ConnectError('missing metadata', Code.NotFound)
      return metadata
    },
    async getMetadataCollections(request, context) {
      const subject = context.values.get(SubjectKey)
      const ids = await dataSource.getMetadataCollectionIds(request.id)
      const viewableIds = await permissions.bulkCheck(subject, PermissionObjectType.collection_type, ids, PermissionAction.view)
      const collections = viewableIds.map((id) => dataSource.getCollection(id))
      return new Collections({
        collections: (await Promise.all(collections)).filter((c) => c != null),
      })
    },
    async findMetadata(request, context) {
      const metadata = await dataSource.findMetadata(request.attributes)
      const metadataIds = metadata.map((m) => m.id)
      const subject = context.values.get(SubjectKey)
      const validIds = await permissions.bulkCheck(
        subject,
        PermissionObjectType.metadata_type,
        metadataIds,
        PermissionAction.view,
      )
      const validIdsMap: { [key: string]: boolean } = {}
      for (const id of validIds) {
        validIdsMap[id] = true
      }
      return new Metadatas({
        metadata: metadata.filter((m) => validIdsMap[m.id]),
      })
    },
    async getMetadataUploadUrl(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, request.id, PermissionAction.edit)
      const metadata = await dataSource.getMetadata(request.id)
      if (!metadata) {
        throw new ConnectError('missing metadata', Code.NotFound)
      }
      return await objectStore.createUploadUrl(metadata)
    },
    async getMetadataDownloadUrl(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, request.id, PermissionAction.view)
      const metadata = await dataSource.getMetadata(request.id)
      if (!metadata) {
        throw new ConnectError('missing metadata', Code.NotFound)
      }
      if (metadata.sourceIdentifier) {
        // TODO: enable the source to specify an object store
        if (metadata.sourceIdentifier.startsWith('http://') || metadata.sourceIdentifier.startsWith('https://')) {
          return new SignedUrl({
            url: metadata.sourceIdentifier,
          })
        }
      }
      return await objectStore.createDownloadUrl(metadata)
    },
    async addMetadataSupplementary(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.metadataId,
        PermissionAction.service,
      )
      await dataSource.addMetadataSupplementary(
        request.metadataId,
        request.key,
        request.name,
        request.contentType,
        Number(request.contentLength),
        request.traitIds,
        request.sourceId || null,
        request.sourceIdentifier || null,
      )
      return new MetadataSupplementary({
        metadataId: request.metadataId,
        key: request.key,
        name: request.name,
        contentType: request.contentType,
        contentLength: request.contentLength,
        traitIds: request.traitIds,
        sourceId: request.sourceId,
        sourceIdentifier: request.sourceIdentifier,
      })
    },
    async setMetadataSupplementaryReady(request, context) {
      const subject = context.values.get(SubjectKey)
      await setMetadataSupplementaryReady(dataSource, permissions, subject, request.id, request.key)
      return new Empty()
    },
    async getMetadataSupplementaryUploadUrl(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.id,
        PermissionAction.service,
      )
      const supplementary = await dataSource.getMetadataSupplementary(request.id, request.key)
      if (!supplementary) {
        throw new ConnectError('missing supplementary', Code.NotFound)
      }
      return objectStore.createUploadUrl(supplementary)
    },
    async getMetadataSupplementaryDownloadUrl(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, request.id, PermissionAction.view)
      const supplementary = await dataSource.getMetadataSupplementary(request.id, request.key)
      if (!supplementary) {
        throw new ConnectError('missing supplementary', Code.NotFound)
      }
      return objectStore.createDownloadUrl(supplementary)
    },
    async deleteMetadataSupplementary(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.id,
        PermissionAction.service,
      )
      const supplementary = await dataSource.getMetadataSupplementary(request.id, request.key)
      if (!supplementary) {
        throw new ConnectError('missing supplementary', Code.NotFound)
      }
      await objectStore.delete(supplementary)
      await dataSource.deleteMetadataSupplementary(request.id, request.key)
    },
    async getMetadataSupplementaries(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.id,
        PermissionAction.view,
      )
      return new MetadataSupplementaries({ supplementaries: await dataSource.getMetadataSupplementaries(request.id) })
    },
    async getMetadataSupplementary(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.id,
        PermissionAction.view,
      )
      const supplementary = await dataSource.getMetadataSupplementary(request.id, request.key)
      if (!supplementary) {
        throw new ConnectError('missing metadata supplementary', Code.NotFound)
      }
      return supplementary
    },
    async setMetadataReady(request, context) {
      const subject = context.values.get(SubjectKey)
      await setMetadataReady(dataSource, permissions, subject, request.id, request.sourceIdentifier)
      return new Empty()
    },
    async getMetadataPermissions(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, request.id, PermissionAction.manage)
      return await permissions.getPermissions(PermissionObjectType.metadata_type, request.id)
    },
    async addMetadataPermissions(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, request.id, PermissionAction.manage)
      for (const permission of request.permissions) {
        await permissions.createRelationship(PermissionObjectType.collection_type, permission)
      }
      return new Empty()
    },
    async addMetadataPermission(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, request.id, PermissionAction.manage)
      await permissions.createRelationship(
        PermissionObjectType.collection_type,
        new Permission({
          id: request.id,
          subject: request.subject,
          relation: request.relation,
          subjectType: request.subjectType,
        }),
      )
      return new Empty()
    },
    async setWorkflowState(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.metadataId,
        PermissionAction.service,
      )
      const metadata = await dataSource.getMetadata(request.metadataId)
      if (!metadata) {
        throw new ConnectError('missing metadata', Code.NotFound)
      }
      await dataSource.setWorkflowState(
        subject,
        request.metadataId,
        metadata.workflowStateId!,
        request.stateId,
        request.status,
        true,
        request.immediate,
      )
    },
    async setWorkflowStateComplete(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.metadataId,
        PermissionAction.service,
      )
      await setWorkflowStateComplete(subject, dataSource, request.metadataId, request.status)
    },
    async addMetadataRelationship(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.metadataId1,
        PermissionAction.manage,
      )
      await permissions.checkWithError(
        subject,
        PermissionObjectType.metadata_type,
        request.metadataId2,
        PermissionAction.manage,
      )
      await dataSource.addMetadataRelationship(request.metadataId1, request.metadataId2, request.relationship)
      return new Empty()
    },
    async getMetadataRelationships(request, context) {
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, request.id, PermissionAction.view)
      return new MetadataRelationships({
        relationships: await dataSource.getMetadataRelationships(request.id, request.relationship),
      })
    },
  })
}
