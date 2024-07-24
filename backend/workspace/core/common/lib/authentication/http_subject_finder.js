"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.HttpSessionInterceptor = exports.HttpSubjectFinder = void 0;
const permissions_1 = require("../permissions/permissions");
class HttpSubjectFinder {
    constructor(endpoint, serviceAccountId, serviceAccountToken, interceptor) {
        this.endpoint = endpoint;
        this.serviceAccountId = serviceAccountId;
        this.serviceAccountTokenHeader = 'Token ' + serviceAccountToken;
        this.interceptor = interceptor;
    }
    async findSubject(fromCookie, authorization) {
        if (authorization == this.serviceAccountTokenHeader) {
            return { id: this.serviceAccountId, type: permissions_1.SubjectType.serviceaccount };
        }
        const response = await fetch(this.endpoint, {
            method: 'GET',
            headers: {
                ...(fromCookie ? { Cookie: authorization } : { Authorization: authorization }),
            },
        });
        if (!response.ok) {
            throw new permissions_1.PermissionError('failed to get session: ' + response.status);
        }
        const subjectId = await this.interceptor.getSubjectId(response);
        if (!subjectId) {
            return { id: 'anonymous', type: permissions_1.SubjectType.user };
        }
        return { id: subjectId, type: permissions_1.SubjectType.user };
    }
}
exports.HttpSubjectFinder = HttpSubjectFinder;
class HttpSessionInterceptor {
    async getSubjectId(response) {
        const session = await response.json();
        if (!session || !session.identity || !session.identity.id) {
            return null;
        }
        return session.identity.id;
    }
}
exports.HttpSessionInterceptor = HttpSessionInterceptor;
//# sourceMappingURL=http_subject_finder.js.map