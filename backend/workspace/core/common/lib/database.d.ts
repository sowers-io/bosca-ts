import { Pool, PoolClient, QueryResult } from 'pg';
import { Message } from '@bufbuild/protobuf';
export declare function createPool(connectionString?: string): Pool;
export declare class DataSource {
    private readonly pool;
    constructor(pool: Pool);
    transaction<T>(txn: (client: PoolClient) => Promise<T>): Promise<T>;
    query(sql: string, values?: any[]): Promise<QueryResult>;
    queryAndMap<T extends Message>(factory: () => T, sql: string, values?: any[], mapper?: ((row: any) => void) | undefined): Promise<T[]>;
    queryAndMapFirst<T extends Message>(factory: () => T, sql: string, values?: any[], mapper?: ((row: any) => void) | undefined): Promise<T | null>;
}
//# sourceMappingURL=database.d.ts.map