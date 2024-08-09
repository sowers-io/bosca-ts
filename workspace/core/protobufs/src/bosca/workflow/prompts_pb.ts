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
// @generated from file bosca/workflow/prompts.proto (package bosca.workflow, syntax proto3)
/* eslint-disable */
// @ts-nocheck

import type { BinaryReadOptions, FieldList, JsonReadOptions, JsonValue, PartialMessage, PlainMessage } from "@bufbuild/protobuf";
import { Message, proto3 } from "@bufbuild/protobuf";

/**
 * @generated from message bosca.workflow.Prompt
 */
export class Prompt extends Message<Prompt> {
  /**
   * @generated from field: string id = 1;
   */
  id = "";

  /**
   * @generated from field: string name = 2;
   */
  name = "";

  /**
   * @generated from field: string description = 3;
   */
  description = "";

  /**
   * @generated from field: string system_prompt = 4;
   */
  systemPrompt = "";

  /**
   * @generated from field: string user_prompt = 5;
   */
  userPrompt = "";

  /**
   * @generated from field: string input_type = 6;
   */
  inputType = "";

  /**
   * @generated from field: string output_type = 7;
   */
  outputType = "";

  constructor(data?: PartialMessage<Prompt>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.workflow.Prompt";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "id", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 2, name: "name", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 3, name: "description", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 4, name: "system_prompt", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 5, name: "user_prompt", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 6, name: "input_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
    { no: 7, name: "output_type", kind: "scalar", T: 9 /* ScalarType.STRING */ },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Prompt {
    return new Prompt().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Prompt {
    return new Prompt().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Prompt {
    return new Prompt().fromJsonString(jsonString, options);
  }

  static equals(a: Prompt | PlainMessage<Prompt> | undefined, b: Prompt | PlainMessage<Prompt> | undefined): boolean {
    return proto3.util.equals(Prompt, a, b);
  }
}

/**
 * @generated from message bosca.workflow.Prompts
 */
export class Prompts extends Message<Prompts> {
  /**
   * @generated from field: repeated bosca.workflow.Prompt prompts = 1;
   */
  prompts: Prompt[] = [];

  constructor(data?: PartialMessage<Prompts>) {
    super();
    proto3.util.initPartial(data, this);
  }

  static readonly runtime: typeof proto3 = proto3;
  static readonly typeName = "bosca.workflow.Prompts";
  static readonly fields: FieldList = proto3.util.newFieldList(() => [
    { no: 1, name: "prompts", kind: "message", T: Prompt, repeated: true },
  ]);

  static fromBinary(bytes: Uint8Array, options?: Partial<BinaryReadOptions>): Prompts {
    return new Prompts().fromBinary(bytes, options);
  }

  static fromJson(jsonValue: JsonValue, options?: Partial<JsonReadOptions>): Prompts {
    return new Prompts().fromJson(jsonValue, options);
  }

  static fromJsonString(jsonString: string, options?: Partial<JsonReadOptions>): Prompts {
    return new Prompts().fromJsonString(jsonString, options);
  }

  static equals(a: Prompts | PlainMessage<Prompts> | undefined, b: Prompts | PlainMessage<Prompts> | undefined): boolean {
    return proto3.util.equals(Prompts, a, b);
  }
}

