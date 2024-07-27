import { type ConnectRouter } from '@connectrpc/connect'
import { ContentDataSource } from '../datasources/content'
import { createPool, logger, PermissionManager, SpiceDBPermissionManager } from '@bosca/common'
import { S3ObjectStore } from '../objectstores/s3'
import { ObjectStore } from '../objectstores/objectstore'
import { content } from './content'

export default (router: ConnectRouter) => {
  const pool = createPool(process.env.BOSCA_CONTENT_CONNECTION_STRING!)
  const objectStore: ObjectStore = new S3ObjectStore()
  const permissions: PermissionManager = new SpiceDBPermissionManager(
    process.env.BOSCA_PERMISSIONS_ENDPOINT!,
    process.env.BOSCA_PERMISSIONS_SHARED_TOKEN!
  )
  const dataSource = new ContentDataSource(pool)
  dataSource.addRootCollection().catch((e: any) => {
    logger.error({ error: e }, 'failed to create root collection')
    process.exit(1)
  })
  const serviceAccountId = process.env.BOSCA_SERVICE_ACCOUNT_ID!
  return content(router, serviceAccountId, permissions, dataSource, objectStore)
}
