/* eslint-disable */
import { NtNft } from "../ntnft/nt_nft";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "ntnft.ntnft";

export interface OwnerCollection {
  class_id: string;
  token: NtNft | undefined;
}

export interface Owner {
  index: string;
  address: string;
  collection: OwnerCollection[];
}

const baseOwnerCollection: object = { class_id: "" };

export const OwnerCollection = {
  encode(message: OwnerCollection, writer: Writer = Writer.create()): Writer {
    if (message.class_id !== "") {
      writer.uint32(10).string(message.class_id);
    }
    if (message.token !== undefined) {
      NtNft.encode(message.token, writer.uint32(18).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): OwnerCollection {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOwnerCollection } as OwnerCollection;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.class_id = reader.string();
          break;
        case 2:
          message.token = NtNft.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): OwnerCollection {
    const message = { ...baseOwnerCollection } as OwnerCollection;
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = String(object.class_id);
    } else {
      message.class_id = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = NtNft.fromJSON(object.token);
    } else {
      message.token = undefined;
    }
    return message;
  },

  toJSON(message: OwnerCollection): unknown {
    const obj: any = {};
    message.class_id !== undefined && (obj.class_id = message.class_id);
    message.token !== undefined &&
      (obj.token = message.token ? NtNft.toJSON(message.token) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<OwnerCollection>): OwnerCollection {
    const message = { ...baseOwnerCollection } as OwnerCollection;
    if (object.class_id !== undefined && object.class_id !== null) {
      message.class_id = object.class_id;
    } else {
      message.class_id = "";
    }
    if (object.token !== undefined && object.token !== null) {
      message.token = NtNft.fromPartial(object.token);
    } else {
      message.token = undefined;
    }
    return message;
  },
};

const baseOwner: object = { index: "", address: "" };

export const Owner = {
  encode(message: Owner, writer: Writer = Writer.create()): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.address !== "") {
      writer.uint32(18).string(message.address);
    }
    for (const v of message.collection) {
      OwnerCollection.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): Owner {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseOwner } as Owner;
    message.collection = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        case 2:
          message.address = reader.string();
          break;
        case 3:
          message.collection.push(
            OwnerCollection.decode(reader, reader.uint32())
          );
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Owner {
    const message = { ...baseOwner } as Owner;
    message.collection = [];
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = String(object.address);
    } else {
      message.address = "";
    }
    if (object.collection !== undefined && object.collection !== null) {
      for (const e of object.collection) {
        message.collection.push(OwnerCollection.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: Owner): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.address !== undefined && (obj.address = message.address);
    if (message.collection) {
      obj.collection = message.collection.map((e) =>
        e ? OwnerCollection.toJSON(e) : undefined
      );
    } else {
      obj.collection = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<Owner>): Owner {
    const message = { ...baseOwner } as Owner;
    message.collection = [];
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    if (object.address !== undefined && object.address !== null) {
      message.address = object.address;
    } else {
      message.address = "";
    }
    if (object.collection !== undefined && object.collection !== null) {
      for (const e of object.collection) {
        message.collection.push(OwnerCollection.fromPartial(e));
      }
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
