/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";

export const protobufPackage = "ntnft.ntnft";

export interface MsgMint {
  creator: string;
  class_id: string;
}

export interface MsgMintResponse {
  class_id: string;
  token_id: string;
  owner: string;
}

export interface MsgCreateClass {
  creator: string;
  name: string;
  uri: string;
  uriHash: string;
  data: string;
  price: string;
}

export interface MsgCreateClassResponse {
  creator: string;
  class_id: string;
}

export interface MsgCreateModuleAccountClass {
  creator: string;
  name: string;
  price: string;
  moduleName: string;
}

export interface MsgCreateModuleAccountClassResponse {
  class_id: string;
  owner: string;
}

export interface MsgRemoveToken {
  creator: string;
  tokenId: string;
}

export interface MsgRemoveTokenResponse {}

export interface MsgEditToken {
  creator: string;
  tokenId: string;
  uri: string;
  uriHash: string;
  data: string;
}

export interface MsgEditTokenResponse {
  class_id: string;
  token_id: string;
  owner: string;
  uri: string;
  uriHash: string;
  data: string;
}

export interface MsgEditClass {
  creator: string;
  classId: string;
  uri: string;
  uriHash: string;
  data: string;
}

export interface MsgEditClassResponse {}

const baseMsgMint: object = { creator: "", class_id: "" };

export const MsgMint = {
  encode(message: MsgMint, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.class_id !== "") {
      writer.uint32(18).string(message.class_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMint {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMint } as MsgMint;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.class_id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMint {
    const message = { ...baseMsgMint } as MsgMint;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = String(object.class_id);
    } else {
      message.class_id = "";
    }
    return message;
  },

  toJSON(message: MsgMint): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.class_id !== undefined && (obj.class_id = message.class_id);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMint>): MsgMint {
    const message = { ...baseMsgMint } as MsgMint;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = object.class_id;
    } else {
      message.class_id = "";
    }
    return message;
  },
};

const baseMsgMintResponse: object = { class_id: "", token_id: "", owner: "" };

export const MsgMintResponse = {
  encode(message: MsgMintResponse, writer: Writer = Writer.create()): Writer {
    if (message.class_id !== "") {
      writer.uint32(10).string(message.class_id);
    }
    if (message.token_id !== "") {
      writer.uint32(18).string(message.token_id);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgMintResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgMintResponse } as MsgMintResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.class_id = reader.string();
          break;
        case 2:
          message.token_id = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgMintResponse {
    const message = { ...baseMsgMintResponse } as MsgMintResponse;
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = String(object.class_id);
    } else {
      message.class_id = "";
    }
    if (object.token_id !== undefined && object.token_id !== null) {
      message.token_id = String(object.token_id);
    } else {
      message.token_id = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgMintResponse): unknown {
    const obj: any = {};
    message.class_id !== undefined && (obj.class_id = message.class_id);
    message.token_id !== undefined && (obj.token_id = message.token_id);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgMintResponse>): MsgMintResponse {
    const message = { ...baseMsgMintResponse } as MsgMintResponse;
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = object.class_id;
    } else {
      message.class_id = "";
    }
    if (object.token_id !== undefined && object.token_id !== null) {
      message.token_id = object.token_id;
    } else {
      message.token_id = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgCreateClass: object = {
  creator: "",
  name: "",
  uri: "",
  uriHash: "",
  data: "",
  price: "",
};

export const MsgCreateClass = {
  encode(message: MsgCreateClass, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.uri !== "") {
      writer.uint32(26).string(message.uri);
    }
    if (message.uriHash !== "") {
      writer.uint32(34).string(message.uriHash);
    }
    if (message.data !== "") {
      writer.uint32(42).string(message.data);
    }
    if (message.price !== "") {
      writer.uint32(50).string(message.price);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateClass {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateClass } as MsgCreateClass;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.uri = reader.string();
          break;
        case 4:
          message.uriHash = reader.string();
          break;
        case 5:
          message.data = reader.string();
          break;
        case 6:
          message.price = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateClass {
    const message = { ...baseMsgCreateClass } as MsgCreateClass;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = String(object.uri);
    } else {
      message.uri = "";
    }
    if (object.uriHash !== undefined && object.uriHash !== null) {
      message.uriHash = String(object.uriHash);
    } else {
      message.uriHash = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    return message;
  },

  toJSON(message: MsgCreateClass): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.uri !== undefined && (obj.uri = message.uri);
    message.uriHash !== undefined && (obj.uriHash = message.uriHash);
    message.data !== undefined && (obj.data = message.data);
    message.price !== undefined && (obj.price = message.price);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgCreateClass>): MsgCreateClass {
    const message = { ...baseMsgCreateClass } as MsgCreateClass;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = object.uri;
    } else {
      message.uri = "";
    }
    if (object.uriHash !== undefined && object.uriHash !== null) {
      message.uriHash = object.uriHash;
    } else {
      message.uriHash = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    return message;
  },
};

const baseMsgCreateClassResponse: object = { creator: "", class_id: "" };

export const MsgCreateClassResponse = {
  encode(
    message: MsgCreateClassResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.class_id !== "") {
      writer.uint32(18).string(message.class_id);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgCreateClassResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgCreateClassResponse } as MsgCreateClassResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.class_id = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateClassResponse {
    const message = { ...baseMsgCreateClassResponse } as MsgCreateClassResponse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = String(object.class_id);
    } else {
      message.class_id = "";
    }
    return message;
  },

  toJSON(message: MsgCreateClassResponse): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.class_id !== undefined && (obj.class_id = message.class_id);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateClassResponse>
  ): MsgCreateClassResponse {
    const message = { ...baseMsgCreateClassResponse } as MsgCreateClassResponse;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = object.class_id;
    } else {
      message.class_id = "";
    }
    return message;
  },
};

const baseMsgCreateModuleAccountClass: object = {
  creator: "",
  name: "",
  price: "",
  moduleName: "",
};

export const MsgCreateModuleAccountClass = {
  encode(
    message: MsgCreateModuleAccountClass,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.price !== "") {
      writer.uint32(26).string(message.price);
    }
    if (message.moduleName !== "") {
      writer.uint32(34).string(message.moduleName);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateModuleAccountClass {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateModuleAccountClass,
    } as MsgCreateModuleAccountClass;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.price = reader.string();
          break;
        case 4:
          message.moduleName = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateModuleAccountClass {
    const message = {
      ...baseMsgCreateModuleAccountClass,
    } as MsgCreateModuleAccountClass;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = String(object.price);
    } else {
      message.price = "";
    }
    if (object.moduleName !== undefined && object.moduleName !== null) {
      message.moduleName = String(object.moduleName);
    } else {
      message.moduleName = "";
    }
    return message;
  },

  toJSON(message: MsgCreateModuleAccountClass): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.name !== undefined && (obj.name = message.name);
    message.price !== undefined && (obj.price = message.price);
    message.moduleName !== undefined && (obj.moduleName = message.moduleName);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateModuleAccountClass>
  ): MsgCreateModuleAccountClass {
    const message = {
      ...baseMsgCreateModuleAccountClass,
    } as MsgCreateModuleAccountClass;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.price !== undefined && object.price !== null) {
      message.price = object.price;
    } else {
      message.price = "";
    }
    if (object.moduleName !== undefined && object.moduleName !== null) {
      message.moduleName = object.moduleName;
    } else {
      message.moduleName = "";
    }
    return message;
  },
};

const baseMsgCreateModuleAccountClassResponse: object = {
  class_id: "",
  owner: "",
};

export const MsgCreateModuleAccountClassResponse = {
  encode(
    message: MsgCreateModuleAccountClassResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.class_id !== "") {
      writer.uint32(10).string(message.class_id);
    }
    if (message.owner !== "") {
      writer.uint32(18).string(message.owner);
    }
    return writer;
  },

  decode(
    input: Reader | Uint8Array,
    length?: number
  ): MsgCreateModuleAccountClassResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = {
      ...baseMsgCreateModuleAccountClassResponse,
    } as MsgCreateModuleAccountClassResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.class_id = reader.string();
          break;
        case 2:
          message.owner = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgCreateModuleAccountClassResponse {
    const message = {
      ...baseMsgCreateModuleAccountClassResponse,
    } as MsgCreateModuleAccountClassResponse;
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = String(object.class_id);
    } else {
      message.class_id = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    return message;
  },

  toJSON(message: MsgCreateModuleAccountClassResponse): unknown {
    const obj: any = {};
    message.class_id !== undefined && (obj.class_id = message.class_id);
    message.owner !== undefined && (obj.owner = message.owner);
    return obj;
  },

  fromPartial(
    object: DeepPartial<MsgCreateModuleAccountClassResponse>
  ): MsgCreateModuleAccountClassResponse {
    const message = {
      ...baseMsgCreateModuleAccountClassResponse,
    } as MsgCreateModuleAccountClassResponse;
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = object.class_id;
    } else {
      message.class_id = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    return message;
  },
};

const baseMsgRemoveToken: object = { creator: "", tokenId: "" };

export const MsgRemoveToken = {
  encode(message: MsgRemoveToken, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenId !== "") {
      writer.uint32(18).string(message.tokenId);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRemoveToken {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRemoveToken } as MsgRemoveToken;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tokenId = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgRemoveToken {
    const message = { ...baseMsgRemoveToken } as MsgRemoveToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = String(object.tokenId);
    } else {
      message.tokenId = "";
    }
    return message;
  },

  toJSON(message: MsgRemoveToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgRemoveToken>): MsgRemoveToken {
    const message = { ...baseMsgRemoveToken } as MsgRemoveToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = object.tokenId;
    } else {
      message.tokenId = "";
    }
    return message;
  },
};

const baseMsgRemoveTokenResponse: object = {};

export const MsgRemoveTokenResponse = {
  encode(_: MsgRemoveTokenResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgRemoveTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgRemoveTokenResponse } as MsgRemoveTokenResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgRemoveTokenResponse {
    const message = { ...baseMsgRemoveTokenResponse } as MsgRemoveTokenResponse;
    return message;
  },

  toJSON(_: MsgRemoveTokenResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgRemoveTokenResponse>): MsgRemoveTokenResponse {
    const message = { ...baseMsgRemoveTokenResponse } as MsgRemoveTokenResponse;
    return message;
  },
};

const baseMsgEditToken: object = {
  creator: "",
  tokenId: "",
  uri: "",
  uriHash: "",
  data: "",
};

export const MsgEditToken = {
  encode(message: MsgEditToken, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.tokenId !== "") {
      writer.uint32(18).string(message.tokenId);
    }
    if (message.uri !== "") {
      writer.uint32(26).string(message.uri);
    }
    if (message.uriHash !== "") {
      writer.uint32(34).string(message.uriHash);
    }
    if (message.data !== "") {
      writer.uint32(42).string(message.data);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEditToken {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEditToken } as MsgEditToken;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.tokenId = reader.string();
          break;
        case 3:
          message.uri = reader.string();
          break;
        case 4:
          message.uriHash = reader.string();
          break;
        case 5:
          message.data = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEditToken {
    const message = { ...baseMsgEditToken } as MsgEditToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = String(object.tokenId);
    } else {
      message.tokenId = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = String(object.uri);
    } else {
      message.uri = "";
    }
    if (object.uriHash !== undefined && object.uriHash !== null) {
      message.uriHash = String(object.uriHash);
    } else {
      message.uriHash = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    return message;
  },

  toJSON(message: MsgEditToken): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.tokenId !== undefined && (obj.tokenId = message.tokenId);
    message.uri !== undefined && (obj.uri = message.uri);
    message.uriHash !== undefined && (obj.uriHash = message.uriHash);
    message.data !== undefined && (obj.data = message.data);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgEditToken>): MsgEditToken {
    const message = { ...baseMsgEditToken } as MsgEditToken;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.tokenId !== undefined && object.tokenId !== null) {
      message.tokenId = object.tokenId;
    } else {
      message.tokenId = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = object.uri;
    } else {
      message.uri = "";
    }
    if (object.uriHash !== undefined && object.uriHash !== null) {
      message.uriHash = object.uriHash;
    } else {
      message.uriHash = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    return message;
  },
};

const baseMsgEditTokenResponse: object = {
  class_id: "",
  token_id: "",
  owner: "",
  uri: "",
  uriHash: "",
  data: "",
};

export const MsgEditTokenResponse = {
  encode(
    message: MsgEditTokenResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.class_id !== "") {
      writer.uint32(10).string(message.class_id);
    }
    if (message.token_id !== "") {
      writer.uint32(18).string(message.token_id);
    }
    if (message.owner !== "") {
      writer.uint32(26).string(message.owner);
    }
    if (message.uri !== "") {
      writer.uint32(34).string(message.uri);
    }
    if (message.uriHash !== "") {
      writer.uint32(42).string(message.uriHash);
    }
    if (message.data !== "") {
      writer.uint32(50).string(message.data);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEditTokenResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEditTokenResponse } as MsgEditTokenResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.class_id = reader.string();
          break;
        case 2:
          message.token_id = reader.string();
          break;
        case 3:
          message.owner = reader.string();
          break;
        case 4:
          message.uri = reader.string();
          break;
        case 5:
          message.uriHash = reader.string();
          break;
        case 6:
          message.data = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEditTokenResponse {
    const message = { ...baseMsgEditTokenResponse } as MsgEditTokenResponse;
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = String(object.class_id);
    } else {
      message.class_id = "";
    }
    if (object.token_id !== undefined && object.token_id !== null) {
      message.token_id = String(object.token_id);
    } else {
      message.token_id = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = String(object.owner);
    } else {
      message.owner = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = String(object.uri);
    } else {
      message.uri = "";
    }
    if (object.uriHash !== undefined && object.uriHash !== null) {
      message.uriHash = String(object.uriHash);
    } else {
      message.uriHash = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    return message;
  },

  toJSON(message: MsgEditTokenResponse): unknown {
    const obj: any = {};
    message.class_id !== undefined && (obj.class_id = message.class_id);
    message.token_id !== undefined && (obj.token_id = message.token_id);
    message.owner !== undefined && (obj.owner = message.owner);
    message.uri !== undefined && (obj.uri = message.uri);
    message.uriHash !== undefined && (obj.uriHash = message.uriHash);
    message.data !== undefined && (obj.data = message.data);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgEditTokenResponse>): MsgEditTokenResponse {
    const message = { ...baseMsgEditTokenResponse } as MsgEditTokenResponse;
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = object.class_id;
    } else {
      message.class_id = "";
    }
    if (object.token_id !== undefined && object.token_id !== null) {
      message.token_id = object.token_id;
    } else {
      message.token_id = "";
    }
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = object.owner;
    } else {
      message.owner = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = object.uri;
    } else {
      message.uri = "";
    }
    if (object.uriHash !== undefined && object.uriHash !== null) {
      message.uriHash = object.uriHash;
    } else {
      message.uriHash = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    return message;
  },
};

const baseMsgEditClass: object = {
  creator: "",
  classId: "",
  uri: "",
  uriHash: "",
  data: "",
};

export const MsgEditClass = {
  encode(message: MsgEditClass, writer: Writer = Writer.create()): Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.classId !== "") {
      writer.uint32(18).string(message.classId);
    }
    if (message.uri !== "") {
      writer.uint32(26).string(message.uri);
    }
    if (message.uriHash !== "") {
      writer.uint32(34).string(message.uriHash);
    }
    if (message.data !== "") {
      writer.uint32(42).string(message.data);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEditClass {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEditClass } as MsgEditClass;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.classId = reader.string();
          break;
        case 3:
          message.uri = reader.string();
          break;
        case 4:
          message.uriHash = reader.string();
          break;
        case 5:
          message.data = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgEditClass {
    const message = { ...baseMsgEditClass } as MsgEditClass;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
    }
    if (object.classId !== undefined && object.classId !== null) {
      message.classId = String(object.classId);
    } else {
      message.classId = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = String(object.uri);
    } else {
      message.uri = "";
    }
    if (object.uriHash !== undefined && object.uriHash !== null) {
      message.uriHash = String(object.uriHash);
    } else {
      message.uriHash = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = String(object.data);
    } else {
      message.data = "";
    }
    return message;
  },

  toJSON(message: MsgEditClass): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.classId !== undefined && (obj.classId = message.classId);
    message.uri !== undefined && (obj.uri = message.uri);
    message.uriHash !== undefined && (obj.uriHash = message.uriHash);
    message.data !== undefined && (obj.data = message.data);
    return obj;
  },

  fromPartial(object: DeepPartial<MsgEditClass>): MsgEditClass {
    const message = { ...baseMsgEditClass } as MsgEditClass;
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
    }
    if (object.classId !== undefined && object.classId !== null) {
      message.classId = object.classId;
    } else {
      message.classId = "";
    }
    if (object.uri !== undefined && object.uri !== null) {
      message.uri = object.uri;
    } else {
      message.uri = "";
    }
    if (object.uriHash !== undefined && object.uriHash !== null) {
      message.uriHash = object.uriHash;
    } else {
      message.uriHash = "";
    }
    if (object.data !== undefined && object.data !== null) {
      message.data = object.data;
    } else {
      message.data = "";
    }
    return message;
  },
};

const baseMsgEditClassResponse: object = {};

export const MsgEditClassResponse = {
  encode(_: MsgEditClassResponse, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): MsgEditClassResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseMsgEditClassResponse } as MsgEditClassResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(_: any): MsgEditClassResponse {
    const message = { ...baseMsgEditClassResponse } as MsgEditClassResponse;
    return message;
  },

  toJSON(_: MsgEditClassResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<MsgEditClassResponse>): MsgEditClassResponse {
    const message = { ...baseMsgEditClassResponse } as MsgEditClassResponse;
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  Mint(request: MsgMint): Promise<MsgMintResponse>;
  CreateClass(request: MsgCreateClass): Promise<MsgCreateClassResponse>;
  CreateModuleAccountClass(
    request: MsgCreateModuleAccountClass
  ): Promise<MsgCreateModuleAccountClassResponse>;
  RemoveToken(request: MsgRemoveToken): Promise<MsgRemoveTokenResponse>;
  EditToken(request: MsgEditToken): Promise<MsgEditTokenResponse>;
  /** this line is used by starport scaffolding # proto/tx/rpc */
  EditClass(request: MsgEditClass): Promise<MsgEditClassResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Mint(request: MsgMint): Promise<MsgMintResponse> {
    const data = MsgMint.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Msg", "Mint", data);
    return promise.then((data) => MsgMintResponse.decode(new Reader(data)));
  }

  CreateClass(request: MsgCreateClass): Promise<MsgCreateClassResponse> {
    const data = MsgCreateClass.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Msg", "CreateClass", data);
    return promise.then((data) =>
      MsgCreateClassResponse.decode(new Reader(data))
    );
  }

  CreateModuleAccountClass(
    request: MsgCreateModuleAccountClass
  ): Promise<MsgCreateModuleAccountClassResponse> {
    const data = MsgCreateModuleAccountClass.encode(request).finish();
    const promise = this.rpc.request(
      "ntnft.ntnft.Msg",
      "CreateModuleAccountClass",
      data
    );
    return promise.then((data) =>
      MsgCreateModuleAccountClassResponse.decode(new Reader(data))
    );
  }

  RemoveToken(request: MsgRemoveToken): Promise<MsgRemoveTokenResponse> {
    const data = MsgRemoveToken.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Msg", "RemoveToken", data);
    return promise.then((data) =>
      MsgRemoveTokenResponse.decode(new Reader(data))
    );
  }

  EditToken(request: MsgEditToken): Promise<MsgEditTokenResponse> {
    const data = MsgEditToken.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Msg", "EditToken", data);
    return promise.then((data) =>
      MsgEditTokenResponse.decode(new Reader(data))
    );
  }

  EditClass(request: MsgEditClass): Promise<MsgEditClassResponse> {
    const data = MsgEditClass.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Msg", "EditClass", data);
    return promise.then((data) =>
      MsgEditClassResponse.decode(new Reader(data))
    );
  }
}

interface Rpc {
  request(
    service: string,
    method: string,
    data: Uint8Array
  ): Promise<Uint8Array>;
}

type Builtin = Date | Function | Uint8Array | string | number | undefined;
export type DeepPartial<T> = T extends Builtin
  ? T
  : T extends Array<infer U>
  ? Array<DeepPartial<U>>
  : T extends ReadonlyArray<infer U>
  ? ReadonlyArray<DeepPartial<U>>
  : T extends {}
  ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;
