"use strict";
var __createBinding = (this && this.__createBinding) || (Object.create ? (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    var desc = Object.getOwnPropertyDescriptor(m, k);
    if (!desc || ("get" in desc ? !m.__esModule : desc.writable || desc.configurable)) {
      desc = { enumerable: true, get: function() { return m[k]; } };
    }
    Object.defineProperty(o, k2, desc);
}) : (function(o, m, k, k2) {
    if (k2 === undefined) k2 = k;
    o[k2] = m[k];
}));
var __setModuleDefault = (this && this.__setModuleDefault) || (Object.create ? (function(o, v) {
    Object.defineProperty(o, "default", { enumerable: true, value: v });
}) : function(o, v) {
    o["default"] = v;
});
var __importStar = (this && this.__importStar) || function (mod) {
    if (mod && mod.__esModule) return mod;
    var result = {};
    if (mod != null) for (var k in mod) if (k !== "default" && Object.prototype.hasOwnProperty.call(mod, k)) __createBinding(result, mod, k);
    __setModuleDefault(result, mod);
    return result;
};
Object.defineProperty(exports, "__esModule", { value: true });
exports.SpiceDBPermissionManager = void 0;
const protobufs_1 = require("@bosca/protobufs");
const permissions_1 = require("./permissions");
const authzed_node_1 = require("@authzed/authzed-node");
const grpc = __importStar(require("@grpc/grpc-js"));
class SpiceDBPermissionManager {
    constructor(endpoint, token) {
        const client = authzed_node_1.v1.NewClient(token, endpoint, authzed_node_1.v1.ClientSecurity.INSECURE_PLAINTEXT_CREDENTIALS, 1, grpc.ServerCredentials.createInsecure());
        this.client = client.promises;
    }
    async bulkCheck(subject, objectType, resourceId, action) {
        const subjectId = subject.id;
        let subjectType = protobufs_1.PermissionSubjectType.user;
        if (subject.type == permissions_1.SubjectType.serviceaccount) {
            subjectType = protobufs_1.PermissionSubjectType.service_account;
        }
        const items = [];
        for (const id of resourceId) {
            items.push(authzed_node_1.v1.CheckBulkPermissionsRequestItem.create({
                subject: authzed_node_1.v1.SubjectReference.create({
                    object: authzed_node_1.v1.ObjectReference.create({
                        objectType: this.getSubjectType(subjectType),
                        objectId: subjectId,
                    }),
                }),
                resource: authzed_node_1.v1.ObjectReference.create({
                    objectType: this.getObjectType(objectType),
                    objectId: id,
                }),
                permission: this.getAction(action),
            }));
        }
        const check = authzed_node_1.v1.CheckBulkPermissionsRequest.create({ items: items });
        if (subjectType == protobufs_1.PermissionSubjectType.service_account) {
            check.consistency = authzed_node_1.v1.Consistency.create({
                requirement: {
                    oneofKind: 'fullyConsistent',
                    fullyConsistent: true,
                },
            });
        }
        const responses = await this.client.checkBulkPermissions(check);
        if (!responses)
            return [];
        const ids = [];
        for (let i = 0; i < responses.pairs.length; i++) {
            const pair = responses.pairs[i];
            const response = pair.response;
            switch (response.oneofKind) {
                case 'item':
                    // @ts-ignore
                    const item = response.item;
                    if (item.permissionship == authzed_node_1.v1.CheckPermissionResponse_Permissionship.HAS_PERMISSION ||
                        item.permissionship == authzed_node_1.v1.CheckPermissionResponse_Permissionship.CONDITIONAL_PERMISSION) {
                        ids.push(resourceId[i]);
                    }
                    else if (item.permissionship == authzed_node_1.v1.CheckPermissionResponse_Permissionship.NO_PERMISSION) {
                        console.log('permission check failed', resourceId[i], subjectId, this.getAction(action));
                    }
                    break;
                case 'error':
                    throw new permissions_1.PermissionError('check failed');
            }
        }
        return ids;
    }
    async checkWithError(subject, objectType, resourceId, action) {
        const subjectId = subject.id;
        let subjectType = protobufs_1.PermissionSubjectType.user;
        if (subject.type === permissions_1.SubjectType.serviceaccount) {
            subjectType = protobufs_1.PermissionSubjectType.service_account;
        }
        return this.checkWithSubjectIdError(subjectType, subjectId, objectType, resourceId, action);
    }
    async checkWithSubjectIdError(subjectType, subjectId, objectType, resourceId, action) {
        const request = authzed_node_1.v1.CheckPermissionRequest.create({
            resource: authzed_node_1.v1.ObjectReference.create({
                objectType: this.getObjectType(objectType),
                objectId: resourceId,
            }),
            subject: authzed_node_1.v1.SubjectReference.create({
                object: authzed_node_1.v1.ObjectReference.create({
                    objectType: this.getSubjectType(subjectType),
                    objectId: subjectId,
                }),
            }),
            permission: this.getAction(action),
        });
        if (subjectType === protobufs_1.PermissionSubjectType.service_account) {
            request.consistency = authzed_node_1.v1.Consistency.create({
                requirement: {
                    oneofKind: 'fullyConsistent',
                    fullyConsistent: true,
                },
            });
        }
        try {
            const response = await this.client.checkPermission(request);
            if (!response)
                throw new permissions_1.PermissionError('no response');
            if (response.permissionship === authzed_node_1.v1.CheckPermissionResponse_Permissionship.NO_PERMISSION ||
                response.permissionship === authzed_node_1.v1.CheckPermissionResponse_Permissionship.UNSPECIFIED) {
                console.error('permission check failed', resourceId, subjectId, this.getAction(action));
                throw new permissions_1.PermissionError('permission check failed');
            }
        }
        catch (e) {
            console.error('permission check failed', resourceId, subjectId, this.getAction(action));
            throw new permissions_1.PermissionError('permission check failed');
        }
    }
    async createRelationships(objectType, permissions) {
        const updates = [];
        for (const permission of permissions) {
            updates.push(authzed_node_1.v1.RelationshipUpdate.create({
                operation: authzed_node_1.v1.RelationshipUpdate_Operation.CREATE,
                relationship: authzed_node_1.v1.Relationship.create({
                    resource: authzed_node_1.v1.ObjectReference.create({
                        objectType: this.getObjectType(objectType),
                        objectId: permission.id,
                    }),
                    relation: this.getRelation(permission.relation),
                    subject: authzed_node_1.v1.SubjectReference.create({
                        object: authzed_node_1.v1.ObjectReference.create({
                            objectType: this.getSubjectType(permission.subjectType),
                            objectId: permission.subject,
                        }),
                    }),
                }),
            }));
        }
        await this.client.writeRelationships(authzed_node_1.v1.WriteRelationshipsRequest.create({ updates: updates }));
    }
    async createRelationship(objectType, permission) {
        await this.createRelationships(objectType, [permission]);
    }
    async getPermissions(objectType, resourceId) {
        const response = await this.client.readRelationships(authzed_node_1.v1.ReadRelationshipsRequest.create({
            relationshipFilter: {
                resourceType: this.getObjectType(objectType),
                optionalResourceId: resourceId,
            },
            optionalLimit: 0,
            consistency: authzed_node_1.v1.Consistency.create({ requirement: { oneofKind: 'fullyConsistent', fullyConsistent: true } }),
        }));
        const permissions = [];
        for (const r of response) {
            const relationship = r.relationship;
            if (!relationship || !relationship.subject || !relationship.subject.object)
                continue;
            let action = protobufs_1.PermissionRelation.viewers;
            switch (relationship.relation) {
                case 'viewers':
                    action = protobufs_1.PermissionRelation.viewers;
                    break;
                case 'discoverers':
                    action = protobufs_1.PermissionRelation.discoverers;
                    break;
                case 'editors':
                    action = protobufs_1.PermissionRelation.editors;
                    break;
                case 'managers':
                    action = protobufs_1.PermissionRelation.managers;
                    break;
                case 'servicers':
                    action = protobufs_1.PermissionRelation.serviceaccounts;
                    break;
                case 'owners':
                    action = protobufs_1.PermissionRelation.owners;
                    break;
            }
            const permission = new protobufs_1.Permission({
                id: relationship.subject.object.objectId,
                relation: action,
                subject: relationship.subject.object.objectId,
                subjectType: protobufs_1.PermissionSubjectType.user,
            });
            switch (relationship.subject.object.objectType) {
                case permissions_1.SubjectType.user:
                    permission.subjectType = protobufs_1.PermissionSubjectType.user;
                    break;
                case permissions_1.SubjectType.group:
                    permission.subjectType = protobufs_1.PermissionSubjectType.group;
                    break;
                case permissions_1.SubjectType.serviceaccount:
                    permission.subjectType = protobufs_1.PermissionSubjectType.service_account;
                    break;
            }
            permissions.push(permission);
        }
        return new protobufs_1.Permissions({ permissions: permissions });
    }
    getObjectType(objectType) {
        switch (objectType) {
            case protobufs_1.PermissionObjectType.metadata_type:
                return 'metadata';
            case protobufs_1.PermissionObjectType.collection_type:
                return 'collection';
            case protobufs_1.PermissionObjectType.system_resource_type:
                return 'systemresource';
            case protobufs_1.PermissionObjectType.workflow_type:
                return 'workflow';
            case protobufs_1.PermissionObjectType.workflow_state_type:
                return 'workflowstate';
        }
        return '';
    }
    getSubjectType(objectType) {
        switch (objectType) {
            case protobufs_1.PermissionSubjectType.user:
                return 'user';
            case protobufs_1.PermissionSubjectType.group:
                return 'group';
            case protobufs_1.PermissionSubjectType.service_account:
                return 'serviceaccount';
        }
        return '';
    }
    getRelation(relation) {
        switch (relation) {
            case protobufs_1.PermissionRelation.viewers:
                return 'viewers';
            case protobufs_1.PermissionRelation.discoverers:
                return 'discoverers';
            case protobufs_1.PermissionRelation.editors:
                return 'editors';
            case protobufs_1.PermissionRelation.managers:
                return 'managers';
            case protobufs_1.PermissionRelation.serviceaccounts:
                return 'servicers';
            case protobufs_1.PermissionRelation.owners:
                return 'owners';
        }
    }
    getAction(relation) {
        switch (relation) {
            case protobufs_1.PermissionAction.view:
                return 'view';
            case protobufs_1.PermissionAction.list:
                return 'list';
            case protobufs_1.PermissionAction.edit:
                return 'edit';
            case protobufs_1.PermissionAction.manage:
                return 'manage';
            case protobufs_1.PermissionAction.service:
                return 'service';
            case protobufs_1.PermissionAction.delete:
                return 'delete';
        }
    }
}
exports.SpiceDBPermissionManager = SpiceDBPermissionManager;
//# sourceMappingURL=spicedb.js.map