"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.SubjectKey = void 0;
exports.newAuthenticationInterceptor = newAuthenticationInterceptor;
const connect_1 = require("@connectrpc/connect");
const permissions_1 = require("../permissions/permissions");
exports.SubjectKey = (0, connect_1.createContextKey)({
    id: 'anonymous',
    type: permissions_1.SubjectType.user,
} // Default value
);
function newAuthenticationInterceptor(subjectFinder) {
    async function authenticate(req, fromCookie, authorization) {
        if (authorization && authorization.length > 0) {
            const subject = await subjectFinder.findSubject(fromCookie, authorization);
            req.contextValues.set(exports.SubjectKey, subject);
            return true;
        }
        return false;
    }
    return (next) => async (req) => {
        if (await authenticate(req, false, req.header.get('Authorization'))) {
            return await next(req);
        }
        if (await authenticate(req, false, req.header.get('X-Service-Authorization'))) {
            return await next(req);
        }
        if (await authenticate(req, true, req.header.get('Cookie'))) {
            return await next(req);
        }
        return await next(req);
    };
}
//# sourceMappingURL=interceptor.js.map