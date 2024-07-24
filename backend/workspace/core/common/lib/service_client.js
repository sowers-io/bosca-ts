"use strict";
/*
 * Copyright 2024 Sowers, LLC
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
Object.defineProperty(exports, "__esModule", { value: true });
exports.useServiceAccountClient = useServiceAccountClient;
exports.useClient = useClient;
const connect_1 = require("@connectrpc/connect");
const connect_node_1 = require("@connectrpc/connect-node");
const serviceAccountClients = {};
const clients = {};
function newTransport(typeName, serviceAuthorization = undefined) {
    const interceptors = [];
    if (serviceAuthorization) {
        interceptors.push(serviceAuthorization);
    }
    switch (typeName) {
        case 'bosca.content.ContentService':
            return (0, connect_node_1.createGrpcTransport)({
                baseUrl: 'http://' + process.env.BOSCA_CONTENT_API_ADDRESS,
                httpVersion: '2',
                interceptors: interceptors,
            });
        case 'bosca.workflow.WorkflowService':
            return (0, connect_node_1.createGrpcTransport)({
                baseUrl: 'http://' + process.env.BOSCA_WORKFLOW_API_ADDRESS,
                httpVersion: '2',
                interceptors: interceptors,
            });
        default:
            throw new Error('unsupported transport: ' + typeName);
    }
}
function useServiceAccountClient(service) {
    let client = serviceAccountClients[service.typeName];
    if (client != null) {
        return client;
    }
    const serviceAuthorization = (next) => async (req) => {
        req.header.set('authorization', 'Token ' + process.env.BOSCA_SERVICE_ACCOUNT_TOKEN);
        return await next(req);
    };
    const transport = newTransport(service.typeName, serviceAuthorization);
    client = (0, connect_1.createPromiseClient)(service, transport);
    serviceAccountClients[service.typeName] = client;
    return client;
}
function useClient(service) {
    let client = clients[service.typeName];
    if (client != null) {
        return client;
    }
    const transport = newTransport(service.typeName);
    client = (0, connect_1.createPromiseClient)(service, transport);
    clients[service.typeName] = client;
    return client;
}
//# sourceMappingURL=service_client.js.map