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
// @generated from file bosca/search/search.proto (package bosca.search, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3, protoInt64 } from "@bufbuild/protobuf";
import { Metadata } from "../content/metadata_pb";

/**
 * @generated from message bosca.search.SearchRequest
 */
export class SearchRequest extends Message<SearchRequest> {
  /**
   * @generated from field: string query = 1;
   */
  query = "";

  /**
   * @generated from field: uint64 offset = 2;
   */
  offset = protoInt64.zero;

  /**
   * @generated from field: uint64 limit = 3;
   */
  limit = protoInt64.zero;

  constructor(data?: PartialMessage<SearchRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.search.SearchRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "query", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "offset", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
    { no: 3, name: "limit", kind: "scalar", T: 4 /* ScalarType.UINT64 */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SearchRequest {
    return new SearchRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SearchRequest {
    return new SearchRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SearchRequest {
    return new SearchRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SearchRequest | PlainMessage<SearchRequest> | undefined, b: SearchRequest | PlainMessage<SearchRequest> | undefined): boolean {
    return proto3.util.equals(SearchRequest, a, b);
  }
}

/**
 * @generated from message bosca.search.SearchResponse
 */
export class SearchResponse extends Message<SearchResponse> {
  /**
   * @generated from field: repeated bosca.content.Metadata metadata = 1;
   */
  metadata: Metadata[] = [];

  constructor(data?: PartialMessage<SearchResponse>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.search.SearchResponse";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "metadata", kind: "message", T: Metadata, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SearchResponse {
    return new SearchResponse().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SearchResponse {
    return new SearchResponse().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SearchResponse {
    return new SearchResponse().fromJsonString(jsonString, options);
  }

  static equals(a: SearchResponse | PlainMessage<SearchResponse> | undefined, b: SearchResponse | PlainMessage<SearchResponse> | undefined): boolean {
    return proto3.util.equals(SearchResponse, a, b);
  }
}

