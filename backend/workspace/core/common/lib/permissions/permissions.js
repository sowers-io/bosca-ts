"use strict";
Object.defineProperty(exports, "__esModule", { value: true });
exports.PermissionError = exports.SubjectType = void 0;
const connect_1 = require("@connectrpc/connect");
var SubjectType;
(function (SubjectType) {
    SubjectType["user"] = "user";
    SubjectType["group"] = "group";
    SubjectType["serviceaccount"] = "serviceaccount";
})(SubjectType || (exports.SubjectType = SubjectType = {}));
class PermissionError extends connect_1.ConnectError {
    constructor(message) {
        super(message, connect_1.Code.PermissionDenied);
    }
}
exports.PermissionError = PermissionError;
//# sourceMappingURL=permissions.js.map