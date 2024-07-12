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
// @generated from file bosca/content/workflows.proto (package bosca.content, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message bosca.content.SetWorkflowStateRequest
 */
export class SetWorkflowStateRequest extends Message<SetWorkflowStateRequest> {
  /**
   * @generated from field: string metadata_id = 1;
   */
  metadataId = "";

  /**
   * @generated from field: string state_id = 2;
   */
  stateId = "";

  /**
   * @generated from field: string status = 3;
   */
  status = "";

  /**
   * @generated from field: bool immediate = 4;
   */
  immediate = false;

  constructor(data?: PartialMessage<SetWorkflowStateRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.SetWorkflowStateRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "metadata_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "state_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "status", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "immediate", kind: "scalar", T: 8 /* ScalarType.BOOL */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetWorkflowStateRequest {
    return new SetWorkflowStateRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetWorkflowStateRequest {
    return new SetWorkflowStateRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetWorkflowStateRequest {
    return new SetWorkflowStateRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SetWorkflowStateRequest | PlainMessage<SetWorkflowStateRequest> | undefined, b: SetWorkflowStateRequest | PlainMessage<SetWorkflowStateRequest> | undefined): boolean {
    return proto3.util.equals(SetWorkflowStateRequest, a, b);
  }
}

/**
 * @generated from message bosca.content.SetWorkflowStateCompleteRequest
 */
export class SetWorkflowStateCompleteRequest extends Message<SetWorkflowStateCompleteRequest> {
  /**
   * @generated from field: string metadata_id = 1;
   */
  metadataId = "";

  /**
   * @generated from field: string status = 2;
   */
  status = "";

  constructor(data?: PartialMessage<SetWorkflowStateCompleteRequest>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.content.SetWorkflowStateCompleteRequest";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "metadata_id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "status", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): SetWorkflowStateCompleteRequest {
    return new SetWorkflowStateCompleteRequest().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): SetWorkflowStateCompleteRequest {
    return new SetWorkflowStateCompleteRequest().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): SetWorkflowStateCompleteRequest {
    return new SetWorkflowStateCompleteRequest().fromJsonString(jsonString, options);
  }

  static equals(a: SetWorkflowStateCompleteRequest | PlainMessage<SetWorkflowStateCompleteRequest> | undefined, b: SetWorkflowStateCompleteRequest | PlainMessage<SetWorkflowStateCompleteRequest> | undefined): boolean {
    return proto3.util.equals(SetWorkflowStateCompleteRequest, a, b);
  }
}

