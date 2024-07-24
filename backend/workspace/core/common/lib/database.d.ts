import { Pool, QueryResult } from 'pg';
export declare function createPool(connectionString?: string): Pool;
export declare class DataSource {
    private readonly pool;
    constructor(pool: Pool);
    query(sql: string, values?: any[]): Promise<QueryResult>;
}
//# sourceMappingURL=database.d.ts.map