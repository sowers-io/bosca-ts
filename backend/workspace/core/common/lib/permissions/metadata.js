"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.AdministratorGroup = void 0;
exports.newMetadataPermissions = newMetadataPermissions;
const protobufs_1 = require("@bosca/protobufs");
exports.AdministratorGroup = 'administrators';
function newMetadataPermissions(serviceAccountId, userId, metadataId) {
    return [
        new protobufs_1.Permission({
            id: metadataId,
            subject: exports.AdministratorGroup,
            subjectType: protobufs_1.PermissionSubjectType.group,
            relation: protobufs_1.PermissionRelation.owners,
        }),
        new protobufs_1.Permission({
            id: metadataId,
            subject: serviceAccountId,
            subjectType: protobufs_1.PermissionSubjectType.service_account,
            relation: protobufs_1.PermissionRelation.owners,
        }),
        new protobufs_1.Permission({
            id: metadataId,
            subject: userId,
            subjectType: protobufs_1.PermissionSubjectType.user,
            relation: protobufs_1.PermissionRelation.owners,
        }),
    ];
}
//# sourceMappingURL=metadata.js.map