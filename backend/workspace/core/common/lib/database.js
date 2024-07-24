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
}
exports.DataSource = DataSource;
//# sourceMappingURL=database.js.map