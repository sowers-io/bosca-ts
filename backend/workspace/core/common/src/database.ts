import { Pool, QueryResult } from 'pg'

export function createPool(connectionString?: string): Pool {
  return new Pool({
    connectionString: connectionString || process.env.BOSCA_CONTENT_CONNECTION_STRING,
  })
}

export class DataSource {
  private readonly pool: Pool

  constructor(pool: Pool) {
    this.pool = pool
  }

  async query(sql: string, values: any[] = []): Promise<QueryResult> {
    const client = await this.pool.connect()
    try {
      return await client.query(sql, values)
    } finally {
      client.release()
    }
  }
}
