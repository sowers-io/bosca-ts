import { type ConnectRouter } from '@connectrpc/connect'
import { createPool, logger, PermissionManager, SpiceDBPermissionManager } from '@bosca/common'
import { WorkflowDataSource } from '../datasources/workflow'
import { workflow } from './workflow'

export default (router: ConnectRouter) => {
  const pool = createPool(process.env.BOSCA_WORKFLOW_CONNECTION_STRING!)
  const permissions: PermissionManager = new SpiceDBPermissionManager(
    process.env.BOSCA_PERMISSIONS_ENDPOINT!,
    process.env.BOSCA_PERMISSIONS_SHARED_TOKEN!
  )
  const workflowDataSource = new WorkflowDataSource(pool)
  workflowDataSource.initialize().catch((error) => {
    logger.error({ error }, 'Failed to initialize workflow data source')
    process.exit(1)
  })
  return workflow(router, permissions, workflowDataSource)
}
