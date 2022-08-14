/* eslint-disable */
import { Reader, Writer } from "protobufjs/minimal";
import { Params } from "../ntnft/params";
import { Owner } from "../ntnft/owner";
import {
  PageRequest,
  PageResponse,
} from "../cosmos/base/query/v1beta1/pagination";
import { Class } from "../ntnft/class";
import { NtNft } from "../ntnft/nt_nft";

export const protobufPackage = "ntnft.ntnft";

/** QueryParamsRequest is request type for the Query/Params RPC method. */
export interface QueryParamsRequest {}

/** QueryParamsResponse is response type for the Query/Params RPC method. */
export interface QueryParamsResponse {
  /** params holds all the parameters of this module. */
  params: Params | undefined;
}

export interface QueryGetOwnerRequest {
  index: string;
}

export interface QueryGetOwnerResponse {
  owner: Owner | undefined;
}

export interface QueryAllOwnerRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllOwnerResponse {
  owner: Owner[];
  pagination: PageResponse | undefined;
}

export interface QueryGetClassRequest {
  index: string;
}

export interface QueryGetClassResponse {
  class: Class | undefined;
}

export interface QueryAllClassRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllClassResponse {
  class: Class[];
  pagination: PageResponse | undefined;
}

export interface QueryGetNtNftRequest {
  index: string;
}

export interface QueryGetNtNftResponse {
  ntNft: NtNft | undefined;
}

export interface QueryAllNtNftRequest {
  pagination: PageRequest | undefined;
}

export interface QueryAllNtNftResponse {
  ntNft: NtNft[];
  pagination: PageResponse | undefined;
}

const baseQueryParamsRequest: object = {};

export const QueryParamsRequest = {
  encode(_: QueryParamsRequest, writer: Writer = Writer.create()): Writer {
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
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

  fromJSON(_: any): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },

  toJSON(_: QueryParamsRequest): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial(_: DeepPartial<QueryParamsRequest>): QueryParamsRequest {
    const message = { ...baseQueryParamsRequest } as QueryParamsRequest;
    return message;
  },
};

const baseQueryParamsResponse: object = {};

export const QueryParamsResponse = {
  encode(
    message: QueryParamsResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.params !== undefined) {
      Params.encode(message.params, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryParamsResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.params = Params.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromJSON(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },

  toJSON(message: QueryParamsResponse): unknown {
    const obj: any = {};
    message.params !== undefined &&
      (obj.params = message.params ? Params.toJSON(message.params) : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryParamsResponse>): QueryParamsResponse {
    const message = { ...baseQueryParamsResponse } as QueryParamsResponse;
    if (object.params !== undefined && object.params !== null) {
      message.params = Params.fromPartial(object.params);
    } else {
      message.params = undefined;
    }
    return message;
  },
};

const baseQueryGetOwnerRequest: object = { index: "" };

export const QueryGetOwnerRequest = {
  encode(
    message: QueryGetOwnerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOwnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetOwnerRequest } as QueryGetOwnerRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOwnerRequest {
    const message = { ...baseQueryGetOwnerRequest } as QueryGetOwnerRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetOwnerRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetOwnerRequest>): QueryGetOwnerRequest {
    const message = { ...baseQueryGetOwnerRequest } as QueryGetOwnerRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetOwnerResponse: object = {};

export const QueryGetOwnerResponse = {
  encode(
    message: QueryGetOwnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.owner !== undefined) {
      Owner.encode(message.owner, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetOwnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetOwnerResponse } as QueryGetOwnerResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner = Owner.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetOwnerResponse {
    const message = { ...baseQueryGetOwnerResponse } as QueryGetOwnerResponse;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = Owner.fromJSON(object.owner);
    } else {
      message.owner = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetOwnerResponse): unknown {
    const obj: any = {};
    message.owner !== undefined &&
      (obj.owner = message.owner ? Owner.toJSON(message.owner) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetOwnerResponse>
  ): QueryGetOwnerResponse {
    const message = { ...baseQueryGetOwnerResponse } as QueryGetOwnerResponse;
    if (object.owner !== undefined && object.owner !== null) {
      message.owner = Owner.fromPartial(object.owner);
    } else {
      message.owner = undefined;
    }
    return message;
  },
};

const baseQueryAllOwnerRequest: object = {};

export const QueryAllOwnerRequest = {
  encode(
    message: QueryAllOwnerRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllOwnerRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllOwnerRequest } as QueryAllOwnerRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOwnerRequest {
    const message = { ...baseQueryAllOwnerRequest } as QueryAllOwnerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOwnerRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllOwnerRequest>): QueryAllOwnerRequest {
    const message = { ...baseQueryAllOwnerRequest } as QueryAllOwnerRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllOwnerResponse: object = {};

export const QueryAllOwnerResponse = {
  encode(
    message: QueryAllOwnerResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.owner) {
      Owner.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllOwnerResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllOwnerResponse } as QueryAllOwnerResponse;
    message.owner = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.owner.push(Owner.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllOwnerResponse {
    const message = { ...baseQueryAllOwnerResponse } as QueryAllOwnerResponse;
    message.owner = [];
    if (object.owner !== undefined && object.owner !== null) {
      for (const e of object.owner) {
        message.owner.push(Owner.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllOwnerResponse): unknown {
    const obj: any = {};
    if (message.owner) {
      obj.owner = message.owner.map((e) => (e ? Owner.toJSON(e) : undefined));
    } else {
      obj.owner = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllOwnerResponse>
  ): QueryAllOwnerResponse {
    const message = { ...baseQueryAllOwnerResponse } as QueryAllOwnerResponse;
    message.owner = [];
    if (object.owner !== undefined && object.owner !== null) {
      for (const e of object.owner) {
        message.owner.push(Owner.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetClassRequest: object = { index: "" };

export const QueryGetClassRequest = {
  encode(
    message: QueryGetClassRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetClassRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetClassRequest } as QueryGetClassRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClassRequest {
    const message = { ...baseQueryGetClassRequest } as QueryGetClassRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetClassRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetClassRequest>): QueryGetClassRequest {
    const message = { ...baseQueryGetClassRequest } as QueryGetClassRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetClassResponse: object = {};

export const QueryGetClassResponse = {
  encode(
    message: QueryGetClassResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.class !== undefined) {
      Class.encode(message.class, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetClassResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetClassResponse } as QueryGetClassResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.class = Class.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetClassResponse {
    const message = { ...baseQueryGetClassResponse } as QueryGetClassResponse;
    if (object.class !== undefined && object.class !== null) {
      message.class = Class.fromJSON(object.class);
    } else {
      message.class = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetClassResponse): unknown {
    const obj: any = {};
    message.class !== undefined &&
      (obj.class = message.class ? Class.toJSON(message.class) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetClassResponse>
  ): QueryGetClassResponse {
    const message = { ...baseQueryGetClassResponse } as QueryGetClassResponse;
    if (object.class !== undefined && object.class !== null) {
      message.class = Class.fromPartial(object.class);
    } else {
      message.class = undefined;
    }
    return message;
  },
};

const baseQueryAllClassRequest: object = {};

export const QueryAllClassRequest = {
  encode(
    message: QueryAllClassRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllClassRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllClassRequest } as QueryAllClassRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllClassRequest {
    const message = { ...baseQueryAllClassRequest } as QueryAllClassRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClassRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllClassRequest>): QueryAllClassRequest {
    const message = { ...baseQueryAllClassRequest } as QueryAllClassRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllClassResponse: object = {};

export const QueryAllClassResponse = {
  encode(
    message: QueryAllClassResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.class) {
      Class.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllClassResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllClassResponse } as QueryAllClassResponse;
    message.class = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.class.push(Class.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllClassResponse {
    const message = { ...baseQueryAllClassResponse } as QueryAllClassResponse;
    message.class = [];
    if (object.class !== undefined && object.class !== null) {
      for (const e of object.class) {
        message.class.push(Class.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllClassResponse): unknown {
    const obj: any = {};
    if (message.class) {
      obj.class = message.class.map((e) => (e ? Class.toJSON(e) : undefined));
    } else {
      obj.class = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllClassResponse>
  ): QueryAllClassResponse {
    const message = { ...baseQueryAllClassResponse } as QueryAllClassResponse;
    message.class = [];
    if (object.class !== undefined && object.class !== null) {
      for (const e of object.class) {
        message.class.push(Class.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryGetNtNftRequest: object = { index: "" };

export const QueryGetNtNftRequest = {
  encode(
    message: QueryGetNtNftRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNtNftRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetNtNftRequest } as QueryGetNtNftRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.index = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNtNftRequest {
    const message = { ...baseQueryGetNtNftRequest } as QueryGetNtNftRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = String(object.index);
    } else {
      message.index = "";
    }
    return message;
  },

  toJSON(message: QueryGetNtNftRequest): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryGetNtNftRequest>): QueryGetNtNftRequest {
    const message = { ...baseQueryGetNtNftRequest } as QueryGetNtNftRequest;
    if (object.index !== undefined && object.index !== null) {
      message.index = object.index;
    } else {
      message.index = "";
    }
    return message;
  },
};

const baseQueryGetNtNftResponse: object = {};

export const QueryGetNtNftResponse = {
  encode(
    message: QueryGetNtNftResponse,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.ntNft !== undefined) {
      NtNft.encode(message.ntNft, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryGetNtNftResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryGetNtNftResponse } as QueryGetNtNftResponse;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ntNft = NtNft.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryGetNtNftResponse {
    const message = { ...baseQueryGetNtNftResponse } as QueryGetNtNftResponse;
    if (object.ntNft !== undefined && object.ntNft !== null) {
      message.ntNft = NtNft.fromJSON(object.ntNft);
    } else {
      message.ntNft = undefined;
    }
    return message;
  },

  toJSON(message: QueryGetNtNftResponse): unknown {
    const obj: any = {};
    message.ntNft !== undefined &&
      (obj.ntNft = message.ntNft ? NtNft.toJSON(message.ntNft) : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryGetNtNftResponse>
  ): QueryGetNtNftResponse {
    const message = { ...baseQueryGetNtNftResponse } as QueryGetNtNftResponse;
    if (object.ntNft !== undefined && object.ntNft !== null) {
      message.ntNft = NtNft.fromPartial(object.ntNft);
    } else {
      message.ntNft = undefined;
    }
    return message;
  },
};

const baseQueryAllNtNftRequest: object = {};

export const QueryAllNtNftRequest = {
  encode(
    message: QueryAllNtNftRequest,
    writer: Writer = Writer.create()
  ): Writer {
    if (message.pagination !== undefined) {
      PageRequest.encode(message.pagination, writer.uint32(10).fork()).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNtNftRequest {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllNtNftRequest } as QueryAllNtNftRequest;
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.pagination = PageRequest.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNtNftRequest {
    const message = { ...baseQueryAllNtNftRequest } as QueryAllNtNftRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNtNftRequest): unknown {
    const obj: any = {};
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageRequest.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(object: DeepPartial<QueryAllNtNftRequest>): QueryAllNtNftRequest {
    const message = { ...baseQueryAllNtNftRequest } as QueryAllNtNftRequest;
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageRequest.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

const baseQueryAllNtNftResponse: object = {};

export const QueryAllNtNftResponse = {
  encode(
    message: QueryAllNtNftResponse,
    writer: Writer = Writer.create()
  ): Writer {
    for (const v of message.ntNft) {
      NtNft.encode(v!, writer.uint32(10).fork()).ldelim();
    }
    if (message.pagination !== undefined) {
      PageResponse.encode(
        message.pagination,
        writer.uint32(18).fork()
      ).ldelim();
    }
    return writer;
  },

  decode(input: Reader | Uint8Array, length?: number): QueryAllNtNftResponse {
    const reader = input instanceof Uint8Array ? new Reader(input) : input;
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = { ...baseQueryAllNtNftResponse } as QueryAllNtNftResponse;
    message.ntNft = [];
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.ntNft.push(NtNft.decode(reader, reader.uint32()));
          break;
        case 2:
          message.pagination = PageResponse.decode(reader, reader.uint32());
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): QueryAllNtNftResponse {
    const message = { ...baseQueryAllNtNftResponse } as QueryAllNtNftResponse;
    message.ntNft = [];
    if (object.ntNft !== undefined && object.ntNft !== null) {
      for (const e of object.ntNft) {
        message.ntNft.push(NtNft.fromJSON(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromJSON(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },

  toJSON(message: QueryAllNtNftResponse): unknown {
    const obj: any = {};
    if (message.ntNft) {
      obj.ntNft = message.ntNft.map((e) => (e ? NtNft.toJSON(e) : undefined));
    } else {
      obj.ntNft = [];
    }
    message.pagination !== undefined &&
      (obj.pagination = message.pagination
        ? PageResponse.toJSON(message.pagination)
        : undefined);
    return obj;
  },

  fromPartial(
    object: DeepPartial<QueryAllNtNftResponse>
  ): QueryAllNtNftResponse {
    const message = { ...baseQueryAllNtNftResponse } as QueryAllNtNftResponse;
    message.ntNft = [];
    if (object.ntNft !== undefined && object.ntNft !== null) {
      for (const e of object.ntNft) {
        message.ntNft.push(NtNft.fromPartial(e));
      }
    }
    if (object.pagination !== undefined && object.pagination !== null) {
      message.pagination = PageResponse.fromPartial(object.pagination);
    } else {
      message.pagination = undefined;
    }
    return message;
  },
};

/** Query defines the gRPC querier service. */
export interface Query {
  /** Parameters queries the parameters of the module. */
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse>;
  /** Queries a Owner by index. */
  Owner(request: QueryGetOwnerRequest): Promise<QueryGetOwnerResponse>;
  /** Queries a list of Owner items. */
  OwnerAll(request: QueryAllOwnerRequest): Promise<QueryAllOwnerResponse>;
  /** Queries a Class by index. */
  Class(request: QueryGetClassRequest): Promise<QueryGetClassResponse>;
  /** Queries a list of Class items. */
  ClassAll(request: QueryAllClassRequest): Promise<QueryAllClassResponse>;
  /** Queries a NtNft by index. */
  NtNft(request: QueryGetNtNftRequest): Promise<QueryGetNtNftResponse>;
  /** Queries a list of NtNft items. */
  NtNftAll(request: QueryAllNtNftRequest): Promise<QueryAllNtNftResponse>;
}

export class QueryClientImpl implements Query {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
  }
  Params(request: QueryParamsRequest): Promise<QueryParamsResponse> {
    const data = QueryParamsRequest.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Query", "Params", data);
    return promise.then((data) => QueryParamsResponse.decode(new Reader(data)));
  }

  Owner(request: QueryGetOwnerRequest): Promise<QueryGetOwnerResponse> {
    const data = QueryGetOwnerRequest.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Query", "Owner", data);
    return promise.then((data) =>
      QueryGetOwnerResponse.decode(new Reader(data))
    );
  }

  OwnerAll(request: QueryAllOwnerRequest): Promise<QueryAllOwnerResponse> {
    const data = QueryAllOwnerRequest.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Query", "OwnerAll", data);
    return promise.then((data) =>
      QueryAllOwnerResponse.decode(new Reader(data))
    );
  }

  Class(request: QueryGetClassRequest): Promise<QueryGetClassResponse> {
    const data = QueryGetClassRequest.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Query", "Class", data);
    return promise.then((data) =>
      QueryGetClassResponse.decode(new Reader(data))
    );
  }

  ClassAll(request: QueryAllClassRequest): Promise<QueryAllClassResponse> {
    const data = QueryAllClassRequest.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Query", "ClassAll", data);
    return promise.then((data) =>
      QueryAllClassResponse.decode(new Reader(data))
    );
  }

  NtNft(request: QueryGetNtNftRequest): Promise<QueryGetNtNftResponse> {
    const data = QueryGetNtNftRequest.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Query", "NtNft", data);
    return promise.then((data) =>
      QueryGetNtNftResponse.decode(new Reader(data))
    );
  }

  NtNftAll(request: QueryAllNtNftRequest): Promise<QueryAllNtNftResponse> {
    const data = QueryAllNtNftRequest.encode(request).finish();
    const promise = this.rpc.request("ntnft.ntnft.Query", "NtNftAll", data);
    return promise.then((data) =>
      QueryAllNtNftResponse.decode(new Reader(data))
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
