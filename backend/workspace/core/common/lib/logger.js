"use strict";
var __importDefault = (this && this.__importDefault) || function (mod) {
    return (mod && mod.__esModule) ? mod : { "default": mod };
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.logger = void 0;
exports.newLoggingInterceptor = newLoggingInterceptor;
const pino_1 = __importDefault(require("pino"));
exports.logger = (0, pino_1.default)({
    level: 'debug',
    serializers: {
        err: pino_1.default.stdSerializers.err,
        error: pino_1.default.stdSerializers.err,
    },
});
function newLoggingInterceptor() {
    return (next) => async (req) => {
        try {
            return await next(req);
        }
        catch (e) {
            exports.logger.error({ error: e }, 'uncaught error');
            throw e;
        }
    };
}
//# sourceMappingURL=logger.js.map