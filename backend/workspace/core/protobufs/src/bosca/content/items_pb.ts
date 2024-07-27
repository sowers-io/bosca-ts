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
// @generated from file bosca/content/items.proto (package bosca.content, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, Timestamp } from "@bufbuild/protobuf";

/**
 * @generated from enum bosca.content.ItemType
 */
export enum ItemType {
  /**
   * @generated from enum value: unknown = 0;
   */
  unknown = 0,

  /**
   * @generated from enum value: collection = 1;
   */
  collection = 1,

  /**
   * @generated from enum value: metadata = 2;
   */
  metadata = 2,

  /**
   * @generated from enum value: metadata_variant = 3;
   */
  metadata_variant = 3,
}
// Retrieve enum metadata with: proto3.getEnumType(ItemType)
proto3.util.setEnumType(ItemType, "bosca.content.ItemType", [
  { no: 0, name: "unknown" },
  { no: 1, name: "collection" },
  { no: 2, name: "metadata" },
  { no: 3, name: "metadata_variant" },
]);

/**
 * @generated from message bosca.content.Item
 */
export class Item extends Message<Item> {
  /**
   * @generated from field: string id = 2;
   */
  id = "";

  /**
   * @generated from field: string name = 3;
   */
  name = "";

  /**
   * @generated from field: repeated string category_ids = 12;
   */
  categoryIds: string[] = [];

  /**
   * @generated from field: repeated string tags = 13;
   */
  tags: string[] = [];

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

  constructor(data?: PartialMessage<Item>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.Item";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 2, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 12, name: "category_ids", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 13, name: "tags", kind: "scalar", T: 9 /* ScalarType.STRING */, repeated: true },
    { no: 14, name: "attributes", kind: "map", K: 9 /* ScalarType.STRING */, V: {kind: "scalar", T: 9 /* ScalarType.STRING */} },
    { no: 20, name: "created", kind: "message", T: Timestamp },
    { no: 21, name: "modified", kind: "message", T: Timestamp },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Item {
    return new Item().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Item {
    return new Item().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Item {
    return new Item().fromJsonString(jsonString, options);
  }

  static equals(a: Item | PlainMessage<Item> | undefined, b: Item | PlainMessage<Item> | undefined): boolean {
    return proto3.util.equals(Item, a, b);
  }
}

