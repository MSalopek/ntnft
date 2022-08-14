/* eslint-disable */
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "ntnft.ntnft";

export interface Collection {
  token_ids: string[];
}

export interface Class {
  index: string;
  name: string;
  creator: string;
  uri: string;
  uriHash: string;
  data: string;
  token_ids: Collection | undefined;
}

const baseCollection: object = { token_ids: "" };

export const Collection = {
  encode(message: Collection, writer: Writer = Writer.create()): Writer {
    for (const v of message.token_ids) {
      writer.uint32(10).string(v!);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Collection {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseCollection } as Collection;
    message.token_ids = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.token_ids.push(reader.string());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Collection {
    const message = { ...baseCollection } as Collection;
    message.token_ids = [];
    if (object.token_ids !== undefined && object.token_ids !== null) {
      for (const e of object.token_ids) {
        message.token_ids.push(String(e));
      }
    }
    return message;
  },

  toJSON(message: Collection): unknown {
    const obj: any = {};
    if (message.token_ids) {
      obj.token_ids = message.token_ids.map((e) => e);
    } else {
      obj.token_ids = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Collection>): Collection {
    const message = { ...baseCollection } as Collection;
    message.token_ids = [];
    if (object.token_ids !== undefined && object.token_ids !== null) {
      for (const e of object.token_ids) {
        message.token_ids.push(e);
      }
    }
    return message;
  },
};

const baseClass: object = {
  index: "",
  name: "",
  creator: "",
  uri: "",
  uriHash: "",
  data: "",
};

export const Class = {
  encode(message: Class, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.creator !== "") {
      writer.uint32(26).string(message.creator);
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
    if (message.token_ids !== undefined) {
      Collection.encode(message.token_ids, writer.uint32(58).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Class {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseClass } as Class;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.name = reader.string();
          break;
        case 3:
          message.creator = reader.string();
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
        case 7:
          message.token_ids = Collection.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Class {
    const message = { ...baseClass } as Class;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = String(object.name);
    } else {
      message.name = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = String(object.creator);
    } else {
      message.creator = "";
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
    if (object.token_ids !== undefined && object.token_ids !== null) {
      message.token_ids = Collection.fromJSON(object.token_ids);
    } else {
      message.token_ids = undefined;
    }
    return message;
  },

  toJSON(message: Class): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.name !== undefined && (obj.name = message.name);
    message.creator !== undefined && (obj.creator = message.creator);
    message.uri !== undefined && (obj.uri = message.uri);
    message.uriHash !== undefined && (obj.uriHash = message.uriHash);
    message.data !== undefined && (obj.data = message.data);
    message.token_ids !== undefined &&
      (obj.token_ids = message.token_ids
        ? Collection.toJSON(message.token_ids)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<Class>): Class {
    const message = { ...baseClass } as Class;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.name !== undefined && object.name !== null) {
      message.name = object.name;
    } else {
      message.name = "";
    }
    if (object.creator !== undefined && object.creator !== null) {
      message.creator = object.creator;
    } else {
      message.creator = "";
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
    if (object.token_ids !== undefined && object.token_ids !== null) {
      message.token_ids = Collection.fromPartial(object.token_ids);
    } else {
      message.token_ids = undefined;
    }
    return message;
  },
};

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
