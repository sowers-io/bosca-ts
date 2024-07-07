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

// @generated by protoc-gen-es v1.10.0 with parameter "target=ts,import_extension=none"
// @generated from file bosca/content/metadata.proto (package bosca.content, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64, Struct, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from message bosca.content.Metadata
 */
export class Metadata extends Message<Metadata> {
  /**
   * @generated from field: string default_id = 1;
   */
  defaultId = "";

  /**
   * @generated from field: string id = 2;
   */
  id = "";

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: string content_type = 4;
   */
  contentType = "";

  /**
   * @generated from field: optional string source_id = 5;
   */
  sourceId?: string;

  /**
   * @generated from field: optional string source_identifier = 6;
   */
  sourceIdentifier?: string;

  /**
   * @generated from field: string language_tag = 7;
   */
  languageTag = "";

  /**
   * @generated from field: optional int64 content_length = 8;
   */
  contentLength?: bigint;

  /**
   * @generated from field: repeated string trait_ids = 11;
   */
  traitIds: string[] = [];

  /**
   * @generated from field: repeated string category_ids = 12;
   */
  categoryIds: string[] = [];

  /**
   * @generated from field: repeated string labels = 13;
   */
  labels: string[] = [];

  /**
   * @generated from field: map<string, string> attributes = 14;
   */
  attributes: { [key: string]: string } = {};

  /**
   * @generated from field: google.protobuf.Timestamp created = 20;
   */
  created?: Timestamp;

  /**
   * @generated from field: google.protobuf.Timestamp modified = 21;
   */
  modified?: Timestamp;

  /**
   * @generated from field: string workflow_state_id = 31;
   */
  workflowStateId = "";

  /**
   * @generated from field: optional string workflow_state_pending_id = 32;
   */
  workflowStatePendingId?: string;

  /**
   * @generated from field: google.protobuf.Struct metadata = 33;
   */
  metadata?: Struct;

  constructor(data?: PartialMessage<Metadata>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.Metadata";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "default_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "content_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "source_id", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 6, name: "source_identifier", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 7, name: "language_tag", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 8, name: "content_length", kind: "scalar", T: 3 /* ScalarType.INT64 */, opt: true },
    { no: 11, name: "trait_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 12, name: "category_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 13, name: "labels", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 14, name: "attributes", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 20, name: "created", kind: "message", T: Timestamp },
    { no: 21, name: "modified", kind: "message", T: Timestamp },
    { no: 31, name: "workflow_state_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 32, name: "workflow_state_pending_id", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 33, name: "metadata", kind: "message", T: Struct },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Metadata {
    return new Metadata().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Metadata {
    return new Metadata().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Metadata {
    return new Metadata().fromJsonString(jsonString, options);
  }

  static equals(a: Metadata | PlainMessage<Metadata> | undefined, b: Metadata | PlainMessage<Metadata> | undefined): boolean {
    return proto3.util.equals(Metadata, a, b);
  }
}

/**
 * @generated from message bosca.content.AddMetadataRequest
 */
export class AddMetadataRequest extends Message<AddMetadataRequest> {
  /**
   * @generated from field: optional string collection = 1;
   */
  collection?: string;

  /**
   * @generated from field: bosca.content.Metadata metadata = 2;
   */
  metadata?: Metadata;

  constructor(data?: PartialMessage<AddMetadataRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.AddMetadataRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "collection", kind: "scalar", T: 9 /* ScalarType.STRING */, opt: true },
    { no: 2, name: "metadata", kind: "message", T: Metadata },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddMetadataRequest {
    return new AddMetadataRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddMetadataRequest {
    return new AddMetadataRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddMetadataRequest {
    return new AddMetadataRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AddMetadataRequest | PlainMessage<AddMetadataRequest> | undefined, b: AddMetadataRequest | PlainMessage<AddMetadataRequest> | undefined): boolean {
    return proto3.util.equals(AddMetadataRequest, a, b);
  }
}

/**
 * @generated from message bosca.content.AddMetadatasRequest
 */
export class AddMetadatasRequest extends Message<AddMetadatasRequest> {
  /**
   * @generated from field: repeated bosca.content.AddMetadataRequest metadatas = 1;
   */
  metadatas: AddMetadataRequest[] = [];

  constructor(data?: PartialMessage<AddMetadatasRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.AddMetadatasRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "metadatas", kind: "message", T: AddMetadataRequest, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddMetadatasRequest {
    return new AddMetadatasRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddMetadatasRequest {
    return new AddMetadatasRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddMetadatasRequest {
    return new AddMetadatasRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AddMetadatasRequest | PlainMessage<AddMetadatasRequest> | undefined, b: AddMetadatasRequest | PlainMessage<AddMetadatasRequest> | undefined): boolean {
    return proto3.util.equals(AddMetadatasRequest, a, b);
  }
}

/**
 * @generated from message bosca.content.Metadatas
 */
export class Metadatas extends Message<Metadatas> {
  /**
   * @generated from field: repeated bosca.content.Metadata metadata = 1;
   */
  metadata: Metadata[] = [];

  constructor(data?: PartialMessage<Metadatas>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.Metadatas";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "metadata", kind: "message", T: Metadata, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Metadatas {
    return new Metadatas().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Metadatas {
    return new Metadatas().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Metadatas {
    return new Metadatas().fromJsonString(jsonString, options);
  }

  static equals(a: Metadatas | PlainMessage<Metadatas> | undefined, b: Metadatas | PlainMessage<Metadatas> | undefined): boolean {
    return proto3.util.equals(Metadatas, a, b);
  }
}

/**
 * @generated from message bosca.content.MetadataRelationshipIdRequest
 */
export class MetadataRelationshipIdRequest extends Message<MetadataRelationshipIdRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string relationship = 2;
   */
  relationship = "";

  constructor(data?: PartialMessage<MetadataRelationshipIdRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.MetadataRelationshipIdRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "relationship", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MetadataRelationshipIdRequest {
    return new MetadataRelationshipIdRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MetadataRelationshipIdRequest {
    return new MetadataRelationshipIdRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MetadataRelationshipIdRequest {
    return new MetadataRelationshipIdRequest().fromJsonString(jsonString, options);
  }

  static equals(a: MetadataRelationshipIdRequest | PlainMessage<MetadataRelationshipIdRequest> | undefined, b: MetadataRelationshipIdRequest | PlainMessage<MetadataRelationshipIdRequest> | undefined): boolean {
    return proto3.util.equals(MetadataRelationshipIdRequest, a, b);
  }
}

/**
 * @generated from message bosca.content.MetadataRelationship
 */
export class MetadataRelationship extends Message<MetadataRelationship> {
  /**
   * @generated from field: string metadata_id1 = 1;
   */
  metadataId1 = "";

  /**
   * @generated from field: string metadata_id2 = 2;
   */
  metadataId2 = "";

  /**
   * @generated from field: string relationship = 3;
   */
  relationship = "";

  /**
   * @generated from field: map<string, string> attributes = 4;
   */
  attributes: { [key: string]: string } = {};

  constructor(data?: PartialMessage<MetadataRelationship>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.MetadataRelationship";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "metadata_id1", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "metadata_id2", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "relationship", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "attributes", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MetadataRelationship {
    return new MetadataRelationship().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MetadataRelationship {
    return new MetadataRelationship().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MetadataRelationship {
    return new MetadataRelationship().fromJsonString(jsonString, options);
  }

  static equals(a: MetadataRelationship | PlainMessage<MetadataRelationship> | undefined, b: MetadataRelationship | PlainMessage<MetadataRelationship> | undefined): boolean {
    return proto3.util.equals(MetadataRelationship, a, b);
  }
}

/**
 * @generated from message bosca.content.MetadataRelationships
 */
export class MetadataRelationships extends Message<MetadataRelationships> {
  /**
   * @generated from field: repeated bosca.content.MetadataRelationship relationships = 1;
   */
  relationships: MetadataRelationship[] = [];

  constructor(data?: PartialMessage<MetadataRelationships>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.MetadataRelationships";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "relationships", kind: "message", T: MetadataRelationship, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): MetadataRelationships {
    return new MetadataRelationships().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): MetadataRelationships {
    return new MetadataRelationships().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): MetadataRelationships {
    return new MetadataRelationships().fromJsonString(jsonString, options);
  }

  static equals(a: MetadataRelationships | PlainMessage<MetadataRelationships> | undefined, b: MetadataRelationships | PlainMessage<MetadataRelationships> | undefined): boolean {
    return proto3.util.equals(MetadataRelationships, a, b);
  }
}

/**
 * @generated from message bosca.content.AddMetadataTraitRequest
 */
export class AddMetadataTraitRequest extends Message<AddMetadataTraitRequest> {
  /**
   * @generated from field: string metadata_id = 1;
   */
  metadataId = "";

  /**
   * @generated from field: string trait_id = 2;
   */
  traitId = "";

  constructor(data?: PartialMessage<AddMetadataTraitRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.AddMetadataTraitRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "metadata_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "trait_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddMetadataTraitRequest {
    return new AddMetadataTraitRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddMetadataTraitRequest {
    return new AddMetadataTraitRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddMetadataTraitRequest {
    return new AddMetadataTraitRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AddMetadataTraitRequest | PlainMessage<AddMetadataTraitRequest> | undefined, b: AddMetadataTraitRequest | PlainMessage<AddMetadataTraitRequest> | undefined): boolean {
    return proto3.util.equals(AddMetadataTraitRequest, a, b);
  }
}

/**
 * @generated from message bosca.content.AddSupplementaryRequest
 */
export class AddSupplementaryRequest extends Message<AddSupplementaryRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string type = 2;
   */
  type = "";

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: string content_type = 4;
   */
  contentType = "";

  /**
   * @generated from field: int64 content_length = 5;
   */
  contentLength = protoInt64.zero;

  constructor(data?: PartialMessage<AddSupplementaryRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.AddSupplementaryRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "content_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "content_length", kind: "scalar", T: 3 /* ScalarType.INT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): AddSupplementaryRequest {
    return new AddSupplementaryRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): AddSupplementaryRequest {
    return new AddSupplementaryRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): AddSupplementaryRequest {
    return new AddSupplementaryRequest().fromJsonString(jsonString, options);
  }

  static equals(a: AddSupplementaryRequest | PlainMessage<AddSupplementaryRequest> | undefined, b: AddSupplementaryRequest | PlainMessage<AddSupplementaryRequest> | undefined): boolean {
    return proto3.util.equals(AddSupplementaryRequest, a, b);
  }
}

/**
 * @generated from message bosca.content.SupplementaryIdRequest
 */
export class SupplementaryIdRequest extends Message<SupplementaryIdRequest> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string type = 2;
   */
  type = "";

  constructor(data?: PartialMessage<SupplementaryIdRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.SupplementaryIdRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SupplementaryIdRequest {
    return new SupplementaryIdRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SupplementaryIdRequest {
    return new SupplementaryIdRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SupplementaryIdRequest {
    return new SupplementaryIdRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SupplementaryIdRequest | PlainMessage<SupplementaryIdRequest> | undefined, b: SupplementaryIdRequest | PlainMessage<SupplementaryIdRequest> | undefined): boolean {
    return proto3.util.equals(SupplementaryIdRequest, a, b);
  }
}

/**
 * @generated from message bosca.content.FindMetadataRequest
 */
export class FindMetadataRequest extends Message<FindMetadataRequest> {
  /**
   * @generated from field: map<string, string> attributes = 1;
   */
  attributes: { [key: string]: string } = {};

  constructor(data?: PartialMessage<FindMetadataRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.FindMetadataRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "attributes", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): FindMetadataRequest {
    return new FindMetadataRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): FindMetadataRequest {
    return new FindMetadataRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): FindMetadataRequest {
    return new FindMetadataRequest().fromJsonString(jsonString, options);
  }

  static equals(a: FindMetadataRequest | PlainMessage<FindMetadataRequest> | undefined, b: FindMetadataRequest | PlainMessage<FindMetadataRequest> | undefined): boolean {
    return proto3.util.equals(FindMetadataRequest, a, b);
  }
}

