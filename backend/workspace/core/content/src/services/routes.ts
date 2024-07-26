import { type ConnectRouter } from '@connectrpc/connect'
import { ContentDataSource } from '../datasources/content'
import { createPool, PermissionManager, SpiceDBPermissionManager } from '@bosca/common'
import { MinioObjectStore } from '../objectstores/minio'
import { ObjectStore } from '../objectstores/objectstore'
import { WorkflowDataSource } from '../datasources/workflow'
import { content } from './content'
import { workflow } from './workflow'

export default (router: ConnectRouter) => {
  const pool = createPool()
  const objectStore: ObjectStore = new MinioObjectStore()
  const permissions: PermissionManager = new SpiceDBPermissionManager(
    process.env.BOSCA_PERMISSIONS_ENDPOINT!,
    process.env.BOSCA_PERMISSIONS_SHARED_TOKEN!
  )
  const workflowDataSource = new WorkflowDataSource(pool)
  const dataSource = new ContentDataSource(pool, workflowDataSource)

  dataSource.addRootCollection().catch((e: any) => {
    console.error('failed to create root collection', e)
  })

  const serviceAccountId = process.env.BOSCA_SERVICE_ACCOUNT_ID!

  router = content(router, serviceAccountId, permissions, dataSource, objectStore)
  router = workflow(router, permissions, workflowDataSource, dataSource)

  return router
}
