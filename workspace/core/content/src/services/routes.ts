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
import { ContentDataSource } from '../datasources/content'
import { health, createPool, logger, PermissionManager, SpiceDBPermissionManager } from '@bosca/common'
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
  return content(health(router), serviceAccountId, permissions, dataSource, objectStore)
}
