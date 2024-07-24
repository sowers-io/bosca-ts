import {
  AddSupplementaryRequest,
  ContentService,
  Empty,
  IdRequest,
  IdResponse,
  MetadataRelationship,
  MetadataRelationshipIdRequest,
  MetadataRelationships,
  MetadataSupplementaries,
  MetadataSupplementary,
  Permission,
  PermissionAction,
  PermissionObjectType,
  Permissions,
  SetWorkflowStateCompleteRequest,
  SetWorkflowStateRequest,
  SignedUrl,
  SupplementaryIdRequest,
} from '@bosca/protobufs'
import { Code, ConnectError, type ConnectRouter } from '@connectrpc/connect'
import { ContentDataSource } from './datasource'
import {
  createPool,
  DataSource,
  newMetadataPermissions,
  PermissionManager,
  SpiceDBPermissionManager,
  SubjectKey,
} from '@bosca/common'
import { MinioObjectStore } from './objectstores/minio'
import { ObjectStore } from './objectstores/objectstore'

const RootCollectionId = '00000000-0000-0000-0000-000000000000'

export default (router: ConnectRouter) => {
  const pool = createPool()
  const objectStore: ObjectStore = new MinioObjectStore()
  const permissions: PermissionManager = new SpiceDBPermissionManager(
    process.env.BOSCA_PERMISSIONS_ENDPOINT!,
    process.env.BOSCA_PERMISSIONS_SHARED_TOKEN!
  )
  const dataSource = new ContentDataSource(pool)
  const serviceAccountId = process.env.BOSCA_SERVICE_ACCOUNT_ID!
  return router.service(ContentService, {
    async getSources(request, context) {
      throw new Error('unimplemented')
    },
    async getSource(request, context) {
      throw new Error('unimplemented')
    },
    async getTraits(request, context) {
      throw new Error('unimplemented')
    },
    async getTrait(request, context) {
      throw new Error('unimplemented')
    },
    async getRootCollectionItems(request, context) {
      throw new Error('unimplemented')
    },
    async getCollectionItems(request, context) {
      throw new Error('unimplemented')
    },
    async addCollections(request, context) {
      throw new Error('unimplemented')
    },
    async addCollection(request, context) {
      throw new Error('unimplemented')
    },
    async deleteCollection(request, context) {
      throw new Error('unimplemented')
    },
    async getCollectionPermissions(request, context) {
      throw new Error('unimplemented')
    },
    async addCollectionPermission(request, context) {
      throw new Error('unimplemented')
    },
    async addCollectionItem(request, context) {
      throw new Error('unimplemented')
    },
    async setCollectionReady(request, context) {
      throw new Error('unimplemented')
    },
    async getCollection(request, context) {
      throw new Error('unimplemented')
    },
    async findCollection(request, context) {
      throw new Error('unimplemented')
    },
    async checkPermission(request, context) {
      throw new Error('unimplemented')
    },
    async addMetadatas(request, context) {
      throw new Error('unimplemented')
    },
    async addMetadataTrait(request, context) {
      throw new Error('unimplemented')
    },
    async addMetadata(request, context) {
      if (!request.metadata) throw new ConnectError('missing metadata', Code.InvalidArgument)
      let collection = RootCollectionId
      if (request.collection && request.collection.length > 0) {
        collection = request.collection
      }
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.collection_type, collection, PermissionAction.edit)
      const id = await dataSource.addMetadata(request.metadata)
      const metadataPermissions = newMetadataPermissions(serviceAccountId, subject.id, id)
      await permissions.createRelationships(PermissionObjectType.metadata_type, metadataPermissions)
      return new IdResponse({ id: id })
    },
    async deleteMetadata(request, context) {
      const metadata = await dataSource.getMetadata(request.id)
      if (!metadata) {
        throw new ConnectError('missing metadata', Code.NotFound)
      }
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, metadata.id, PermissionAction.delete)
      await objectStore.delete(metadata.id)
      if (metadata.sourceIdentifier) {
        await objectStore.delete(metadata.sourceIdentifier)
      }
      await dataSource.deleteMetadata(request.id)
      return new IdResponse({ id: request.id })
    },
    async getMetadata(request, context) {
      const metadata = await dataSource.getMetadata(request.id)
      if (!metadata) throw new ConnectError('missing metadata', Code.NotFound)
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, metadata.id, PermissionAction.view)
      return metadata
    },
    async findMetadata(request, context) {
      throw new Error('unimplemented')
    },
    async getMetadataUploadUrl(request, context) {
      const metadata = await dataSource.getMetadata(request.id)
      if (!metadata) {
        throw new ConnectError('missing metadata', Code.NotFound)
      }
      const subject = context.values.get(SubjectKey)
      await permissions.checkWithError(subject, PermissionObjectType.metadata_type, metadata.id, PermissionAction.view)
      return objectStore.createUploadUrl(metadata)
    },
    async getMetadataDownloadUrl(request) {
      const metadata = await dataSource.getMetadata(request.id)
      if (!metadata) {
        throw new ConnectError('missing metadata', Code.NotFound)
      }
      return objectStore.createDownloadUrl(metadata.id)
    },
    async addMetadataSupplementary(request, context) {
      throw new Error('unimplemented')
    },
    async setMetadataSupplementaryReady(request, context) {
      throw new Error('unimplemented')
    },
    async getMetadataSupplementaryUploadUrl(request, context) {
      throw new Error('unimplemented')
    },
    async getMetadataSupplementaryDownloadUrl(request, context) {
      throw new Error('unimplemented')
    },
    async deleteMetadataSupplementary(request, context) {
      throw new Error('unimplemented')
    },
    async getMetadataSupplementaries(request, context) {
      throw new Error('unimplemented')
    },
    async getMetadataSupplementary(request, context) {
      throw new Error('unimplemented')
    },
    async setMetadataReady(request, context) {
      throw new Error('unimplemented')
    },
    async getMetadataPermissions(request, context) {
      throw new Error('unimplemented')
    },
    async addMetadataPermissions(request, context) {
      throw new Error('unimplemented')
    },
    async addMetadataPermission(request, context) {
      throw new Error('unimplemented')
    },
    async setWorkflowState(request, context) {
      throw new Error('unimplemented')
    },
    async setWorkflowStateComplete(request, context) {
      throw new Error('unimplemented')
    },
    async addMetadataRelationship(request, context) {
      throw new Error('unimplemented')
    },
    async getMetadataRelationships(request, context) {
      throw new Error('unimplemented')
    },
  })
}
