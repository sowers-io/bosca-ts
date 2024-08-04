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

import { type ConnectRouter } from '@connectrpc/connect'
import { health, createPool, logger, PermissionManager, SpiceDBPermissionManager } from '@bosca/common'
import { WorkflowDataSource } from '../datasources/workflow'
import { workflow } from './workflow'

export default (router: ConnectRouter) => {
  const pool = createPool(process.env.BOSCA_WORKFLOW_CONNECTION_STRING!)
  const permissions: PermissionManager = new SpiceDBPermissionManager(
    process.env.BOSCA_PERMISSIONS_ENDPOINT!,
    process.env.BOSCA_PERMISSIONS_SHARED_TOKEN!,
  )
  const workflowDataSource = new WorkflowDataSource(pool)
  workflowDataSource.initialize().catch((error) => {
    logger.error({ error }, 'Failed to initialize workflow data source')
    process.exit(1)
  })
  return workflow(health(router), permissions, workflowDataSource)
}
