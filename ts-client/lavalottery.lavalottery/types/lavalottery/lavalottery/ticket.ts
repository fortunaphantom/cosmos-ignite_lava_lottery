/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "lavalottery.lavalottery";

export interface Ticket {
  index: string;
  name: string;
  fee: string;
  bet: string;
}

function createBaseTicket(): Ticket {
  return { index: "", name: "", fee: "", bet: "" };
}

export const Ticket = {
  encode(message: Ticket, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.index !== "") {
      writer.uint32(10).string(message.index);
    }
    if (message.name !== "") {
      writer.uint32(18).string(message.name);
    }
    if (message.fee !== "") {
      writer.uint32(26).string(message.fee);
    }
    if (message.bet !== "") {
      writer.uint32(34).string(message.bet);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): Ticket {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseTicket();
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
          message.fee = reader.string();
          break;
        case 4:
          message.bet = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): Ticket {
    return {
      index: isSet(object.index) ? String(object.index) : "",
      name: isSet(object.name) ? String(object.name) : "",
      fee: isSet(object.fee) ? String(object.fee) : "",
      bet: isSet(object.bet) ? String(object.bet) : "",
    };
  },

  toJSON(message: Ticket): unknown {
    const obj: any = {};
    message.index !== undefined && (obj.index = message.index);
    message.name !== undefined && (obj.name = message.name);
    message.fee !== undefined && (obj.fee = message.fee);
    message.bet !== undefined && (obj.bet = message.bet);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<Ticket>, I>>(object: I): Ticket {
    const message = createBaseTicket();
    message.index = object.index ?? "";
    message.name = object.name ?? "";
    message.fee = object.fee ?? "";
    message.bet = object.bet ?? "";
    return message;
  },
};

type Builtin = Date | Function | Uint8Array | string | number | boolean | undefined;

export type DeepPartial<T> = T extends Builtin ? T
  : T extends Array<infer U> ? Array<DeepPartial<U>> : T extends ReadonlyArray<infer U> ? ReadonlyArray<DeepPartial<U>>
  : T extends {} ? { [K in keyof T]?: DeepPartial<T[K]> }
  : Partial<T>;

type KeysOfUnion<T> = T extends T ? keyof T : never;
export type Exact<P, I extends P> = P extends Builtin ? P
  : P & { [K in keyof P]: Exact<P[K], I[K]> } & { [K in Exclude<keyof I, KeysOfUnion<P>>]: never };

function isSet(value: any): boolean {
  return value !== null && value !== undefined;
}
