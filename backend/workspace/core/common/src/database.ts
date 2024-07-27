import { Pool, PoolClient, QueryResult } from 'pg'
import { Message } from '@bufbuild/protobuf'

export function createPool(connectionString: string): Pool {
  return new Pool({
    connectionString: connectionString,
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
    mapper: ((row: any) => void) | undefined = undefined
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
    mapper: ((row: any) => void) | undefined = undefined
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
