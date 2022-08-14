/* eslint-disable */
import { Params } from "../ntnft/params";
import { Owner } from "../ntnft/owner";
import { Class } from "../ntnft/class";
import { NtNft } from "../ntnft/nt_nft";
import { Writer, Reader } from "protobufjs/minimal";

export const protobufPackage = "ntnft.ntnft";

/** GenesisState defines the ntnft module's genesis state. */
export interface GenesisState {
  params: Params | undefined;
  ownerList: Owner[];
  classList: Class[];
  /** this line is used by starport scaffolding # genesis/proto/state */
  ntNftList: NtNft[];
}

const baseGenesisState: object = {};

export const GenesisState = {
  encode(message: GenesisState, writer: Writer = Writer.create()): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    for (const v of message.ownerList) {
      Owner.encode(v!, writer.uint32(18).fork()).ldelim();
    }
    for (const v of message.classList) {
      Class.encode(v!, writer.uint32(26).fork()).ldelim();
    }
    for (const v of message.ntNftList) {
      NtNft.encode(v!, writer.uint32(34).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): GenesisState {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseGenesisState } as GenesisState;
    message.ownerList = [];
    message.classList = [];
    message.ntNftList = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        case 2:
          message.ownerList.push(Owner.decode(reader, reader.uint32()));
          break;
        case 3:
          message.classList.push(Class.decode(reader, reader.uint32()));
          break;
        case 4:
          message.ntNftList.push(NtNft.decode(reader, reader.uint32()));
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.ownerList = [];
    message.classList = [];
    message.ntNftList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    if (object.ownerList !== undefined && object.ownerList !== null) {
      for (const e of object.ownerList) {
        message.ownerList.push(Owner.fromJSON(e));
      }
    }
    if (object.classList !== undefined && object.classList !== null) {
      for (const e of object.classList) {
        message.classList.push(Class.fromJSON(e));
      }
    }
    if (object.ntNftList !== undefined && object.ntNftList !== null) {
      for (const e of object.ntNftList) {
        message.ntNftList.push(NtNft.fromJSON(e));
      }
    }
    return message;
  },

  toJSON(message: GenesisState): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    if (message.ownerList) {
      obj.ownerList = message.ownerList.map((e) =>
        e ? Owner.toJSON(e) : undefined
      );
    } else {
      obj.ownerList = [];
    }
    if (message.classList) {
      obj.classList = message.classList.map((e) =>
        e ? Class.toJSON(e) : undefined
      );
    } else {
      obj.classList = [];
    }
    if (message.ntNftList) {
      obj.ntNftList = message.ntNftList.map((e) =>
        e ? NtNft.toJSON(e) : undefined
      );
    } else {
      obj.ntNftList = [];
    }
    return obj;
  },

  fromPartial(object: DeepPartial<GenesisState>): GenesisState {
    const message = { ...baseGenesisState } as GenesisState;
    message.ownerList = [];
    message.classList = [];
    message.ntNftList = [];
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    if (object.ownerList !== undefined && object.ownerList !== null) {
      for (const e of object.ownerList) {
        message.ownerList.push(Owner.fromPartial(e));
      }
    }
    if (object.classList !== undefined && object.classList !== null) {
      for (const e of object.classList) {
        message.classList.push(Class.fromPartial(e));
      }
    }
    if (object.ntNftList !== undefined && object.ntNftList !== null) {
      for (const e of object.ntNftList) {
        message.ntNftList.push(NtNft.fromPartial(e));
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
