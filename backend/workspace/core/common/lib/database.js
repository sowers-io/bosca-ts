"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.DataSource = void 0;
exports.createPool = createPool;
const pg_1 = require("pg");
function createPool(connectionString) {
    return new pg_1.Pool({
        connectionString: connectionString || process.env.BOSCA_CONTENT_CONNECTION_STRING,
    });
}
class DataSource {
    constructor(pool) {
        this.pool = pool;
    }
    async query(sql, values = []) {
        const client = await this.pool.connect();
        try {
            return await client.query(sql, values);
        }
        finally {
            client.release();
        }
    }
    async queryAndMap(factory, sql, values = [], mapper = undefined) {
        const client = await this.pool.connect();
        try {
            const records = await client.query(sql, values);
            return records.rows.map((r) => {
                const instance = factory();
                if (mapper) {
                    mapper(r);
                }
                return instance.fromJson(r, { ignoreUnknownFields: true });
            });
        }
        finally {
            client.release();
        }
    }
    async queryAndMapFirst(factory, sql, values = [], mapper = undefined) {
        const client = await this.pool.connect();
        try {
            const records = await client.query(sql, values);
            if (records && records.rows && records.rows.length > 0) {
                if (mapper) {
                    mapper(records.rows[0]);
                }
                return factory().fromJson(records.rows[0], { ignoreUnknownFields: true });
            }
            return null;
        }
        finally {
            client.release();
        }
    }
}
exports.DataSource = DataSource;
//# sourceMappingURL=database.js.map