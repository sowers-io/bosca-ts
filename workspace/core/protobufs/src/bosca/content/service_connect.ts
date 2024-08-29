//
// Copyright 2024 Sowers, LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// @generated by protoc-gen-connect-es v1.4.0 with parameter "target=ts,import_extension=none"
// @generated from file bosca/content/service.proto (package bosca.content, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import { Empty } from "../empty_pb";
import { Source, Sources } from "./sources_pb";
import { MethodKind } from "@bufbuild/protobuf";
import { IdAndVersionRequest, IdRequest, IdResponse, IdResponses, IdsRequest, SupplementaryIdRequest } from "../requests_pb";
import { Trait, Traits } from "./traits_pb";
import { AddCollectionItemRequest, AddCollectionRequest, AddCollectionsRequest, Collection, CollectionItems, Collections, FindCollectionRequest } from "./collections_pb";
import { Permission, PermissionCheckRequest, PermissionCheckResponse, Permissions } from "./permissions_pb";
import { AddMetadataAttributesRequest, AddMetadataRequest, AddMetadatasRequest, AddMetadataTraitRequest, AddSupplementaryRequest, FindMetadataRequest, Metadata, MetadataReadyRequest, MetadataRelationship, MetadataRelationshipIdRequest, MetadataRelationships, Metadatas, MetadataSupplementaries, MetadataSupplementary, SetMetadataTraitsRequest } from "./metadata_pb";
import { SignedUrl } from "./url_pb";
import { SetWorkflowStateCompleteRequest, SetWorkflowStateRequest } from "./workflows_pb";

/**
 * @generated from service bosca.content.ContentService
 */
export const ContentService = {
  typeName: "bosca.content.ContentService",
  methods: {
    /**
     * @generated from rpc bosca.content.ContentService.GetSources
     */
    getSources: {
      name: "GetSources",
      I: Empty,
      O: Sources,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetSource
     */
    getSource: {
      name: "GetSource",
      I: IdRequest,
      O: Source,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetTraits
     */
    getTraits: {
      name: "GetTraits",
      I: Empty,
      O: Traits,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetTrait
     */
    getTrait: {
      name: "GetTrait",
      I: IdRequest,
      O: Trait,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetRootCollectionItems
     */
    getRootCollectionItems: {
      name: "GetRootCollectionItems",
      I: Empty,
      O: CollectionItems,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetCollectionItems
     */
    getCollectionItems: {
      name: "GetCollectionItems",
      I: IdRequest,
      O: CollectionItems,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddCollection
     */
    addCollection: {
      name: "AddCollection",
      I: AddCollectionRequest,
      O: IdResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.SetCollectionReady
     */
    setCollectionReady: {
      name: "SetCollectionReady",
      I: IdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddCollections
     */
    addCollections: {
      name: "AddCollections",
      I: AddCollectionsRequest,
      O: IdResponses,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetCollection
     */
    getCollection: {
      name: "GetCollection",
      I: IdRequest,
      O: Collection,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.DeleteCollection
     */
    deleteCollection: {
      name: "DeleteCollection",
      I: IdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetCollectionPermissions
     */
    getCollectionPermissions: {
      name: "GetCollectionPermissions",
      I: IdRequest,
      O: Permissions,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddCollectionPermission
     */
    addCollectionPermission: {
      name: "AddCollectionPermission",
      I: Permission,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddCollectionItem
     */
    addCollectionItem: {
      name: "AddCollectionItem",
      I: AddCollectionItemRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.CheckPermission
     */
    checkPermission: {
      name: "CheckPermission",
      I: PermissionCheckRequest,
      O: PermissionCheckResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.FindCollection
     */
    findCollection: {
      name: "FindCollection",
      I: FindCollectionRequest,
      O: Collections,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.FindMetadata
     */
    findMetadata: {
      name: "FindMetadata",
      I: FindMetadataRequest,
      O: Metadatas,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadata
     */
    getMetadata: {
      name: "GetMetadata",
      I: IdRequest,
      O: Metadata,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadataVersion
     */
    getMetadataVersion: {
      name: "GetMetadataVersion",
      I: IdAndVersionRequest,
      O: Metadata,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadataCollections
     */
    getMetadataCollections: {
      name: "GetMetadataCollections",
      I: IdRequest,
      O: Collections,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadatas
     */
    getMetadatas: {
      name: "GetMetadatas",
      I: IdsRequest,
      O: Metadatas,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.SetMetadataActiveVersion
     */
    setMetadataActiveVersion: {
      name: "SetMetadataActiveVersion",
      I: IdAndVersionRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddMetadata
     */
    addMetadata: {
      name: "AddMetadata",
      I: AddMetadataRequest,
      O: IdResponse,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddMetadataAttributes
     */
    addMetadataAttributes: {
      name: "AddMetadataAttributes",
      I: AddMetadataAttributesRequest,
      O: Metadata,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddMetadatas
     */
    addMetadatas: {
      name: "AddMetadatas",
      I: AddMetadatasRequest,
      O: IdResponses,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddMetadataTrait
     */
    addMetadataTrait: {
      name: "AddMetadataTrait",
      I: AddMetadataTraitRequest,
      O: Metadata,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.SetMetadataTraits
     */
    setMetadataTraits: {
      name: "SetMetadataTraits",
      I: SetMetadataTraitsRequest,
      O: Metadata,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.DeleteMetadata
     */
    deleteMetadata: {
      name: "DeleteMetadata",
      I: IdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadataUploadUrl
     */
    getMetadataUploadUrl: {
      name: "GetMetadataUploadUrl",
      I: IdRequest,
      O: SignedUrl,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadataDownloadUrl
     */
    getMetadataDownloadUrl: {
      name: "GetMetadataDownloadUrl",
      I: IdRequest,
      O: SignedUrl,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddMetadataSupplementary
     */
    addMetadataSupplementary: {
      name: "AddMetadataSupplementary",
      I: AddSupplementaryRequest,
      O: MetadataSupplementary,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.SetMetadataSupplementaryReady
     */
    setMetadataSupplementaryReady: {
      name: "SetMetadataSupplementaryReady",
      I: SupplementaryIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadataSupplementaryUploadUrl
     */
    getMetadataSupplementaryUploadUrl: {
      name: "GetMetadataSupplementaryUploadUrl",
      I: SupplementaryIdRequest,
      O: SignedUrl,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadataSupplementaryDownloadUrl
     */
    getMetadataSupplementaryDownloadUrl: {
      name: "GetMetadataSupplementaryDownloadUrl",
      I: SupplementaryIdRequest,
      O: SignedUrl,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.DeleteMetadataSupplementary
     */
    deleteMetadataSupplementary: {
      name: "DeleteMetadataSupplementary",
      I: SupplementaryIdRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadataSupplementaries
     */
    getMetadataSupplementaries: {
      name: "GetMetadataSupplementaries",
      I: IdRequest,
      O: MetadataSupplementaries,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadataSupplementary
     */
    getMetadataSupplementary: {
      name: "GetMetadataSupplementary",
      I: SupplementaryIdRequest,
      O: MetadataSupplementary,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.SetMetadataReady
     */
    setMetadataReady: {
      name: "SetMetadataReady",
      I: MetadataReadyRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadataPermissions
     */
    getMetadataPermissions: {
      name: "GetMetadataPermissions",
      I: IdRequest,
      O: Permissions,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddMetadataPermissions
     */
    addMetadataPermissions: {
      name: "AddMetadataPermissions",
      I: Permissions,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddMetadataPermission
     */
    addMetadataPermission: {
      name: "AddMetadataPermission",
      I: Permission,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.SetWorkflowState
     */
    setWorkflowState: {
      name: "SetWorkflowState",
      I: SetWorkflowStateRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.SetWorkflowStateComplete
     */
    setWorkflowStateComplete: {
      name: "SetWorkflowStateComplete",
      I: SetWorkflowStateCompleteRequest,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.AddMetadataRelationship
     */
    addMetadataRelationship: {
      name: "AddMetadataRelationship",
      I: MetadataRelationship,
      O: Empty,
      kind: MethodKind.Unary,
    },
    /**
     * @generated from rpc bosca.content.ContentService.GetMetadataRelationships
     */
    getMetadataRelationships: {
      name: "GetMetadataRelationships",
      I: MetadataRelationshipIdRequest,
      O: MetadataRelationships,
      kind: MethodKind.Unary,
    },
  }
} as const;

