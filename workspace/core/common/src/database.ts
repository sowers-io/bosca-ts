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

import { Pool, PoolClient, QueryResult } from 'pg'
import { Message } from '@bufbuild/protobuf'

export function createPool(connectionString: string): Pool {
  return new Pool({
    connectionString,
  })
}

export class DataSource {
  private readonly pool: Pool

  constructor(pool: Pool) {
    this.pool = pool
  }

  async transaction<T>(txn: (client: PoolClient) => Promise<T>): Promise<T> {
    const client = await this.pool.connect()
    try {
      await client.query('BEGIN')
      try {
        const result = await txn(client)
        await client.query('COMMIT')
        return result
      } catch (e) {
        await client.query('ROLLBACK')
        throw e
      }
    } finally {
      client.release()
    }
  }

  async query(sql: string, values: any[] = []): Promise<QueryResult> {
    const client = await this.pool.connect()
    try {
      return await client.query(sql, values)
    } finally {
      client.release()
    }
  }

  async queryAndMap<T extends Message>(
    factory: () => T,
    sql: string,
    values: any[] = [],
    mapper: ((row: any) => void) | undefined = undefined,
  ): Promise<T[]> {
    const client = await this.pool.connect()
    try {
      const records = await client.query(sql, values)
      return records.rows.map((r) => {
        const instance = factory()
        if (mapper) {
          mapper(r)
        }
        return instance.fromJson(r, { ignoreUnknownFields: true })
      })
    } finally {
      client.release()
    }
  }

  async queryAndMapFirst<T extends Message>(
    factory: () => T,
    sql: string,
    values: any[] = [],
    mapper: ((row: any) => void) | undefined = undefined,
  ): Promise<T | null> {
    const client = await this.pool.connect()
    try {
      const records = await client.query(sql, values)
      if (records && records.rows && records.rows.length > 0) {
        if (mapper) {
          mapper(records.rows[0])
        }
        return factory().fromJson(records.rows[0], { ignoreUnknownFields: true })
      }
      return null
    } finally {
      client.release()
    }
  }
}
