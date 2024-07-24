import { Permission, PermissionAction, PermissionObjectType, Permissions, PermissionSubjectType } from '@bosca/protobufs';
import { PermissionManager } from './permissions';
import { Subject } from '../authentication/subject_finder';
export declare class SpiceDBPermissionManager implements PermissionManager {
    private readonly client;
    constructor(endpoint: string, token: string);
    bulkCheck(subject: Subject, objectType: PermissionObjectType, resourceId: string[], action: PermissionAction): Promise<string[]>;
    checkWithError(subject: Subject, objectType: PermissionObjectType, resourceId: string, action: PermissionAction): Promise<void>;
    checkWithSubjectIdError(subjectType: PermissionSubjectType, subjectId: string, objectType: PermissionObjectType, resourceId: string, action: PermissionAction): Promise<void>;
    createRelationships(objectType: PermissionObjectType, permissions: Permission[]): Promise<void>;
    createRelationship(objectType: PermissionObjectType, permission: Permission): Promise<void>;
    getPermissions(objectType: PermissionObjectType, resourceId: string): Promise<Permissions>;
    private getObjectType;
    private getSubjectType;
    private getRelation;
    private getAction;
}
//# sourceMappingURL=spicedb.d.ts.map