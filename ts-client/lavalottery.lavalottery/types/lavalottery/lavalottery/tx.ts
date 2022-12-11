/* eslint-disable */
import _m0 from "protobufjs/minimal";

export const protobufPackage = "lavalottery.lavalottery";

export interface MsgSendTicket {
  creator: string;
  fee: string;
  bet: string;
}

export interface MsgSendTicketResponse {
}

function createBaseMsgSendTicket(): MsgSendTicket {
  return { creator: "", fee: "", bet: "" };
}

export const MsgSendTicket = {
  encode(message: MsgSendTicket, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    if (message.creator !== "") {
      writer.uint32(10).string(message.creator);
    }
    if (message.fee !== "") {
      writer.uint32(18).string(message.fee);
    }
    if (message.bet !== "") {
      writer.uint32(26).string(message.bet);
    }
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSendTicket {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSendTicket();
    while (reader.pos < end) {
      const tag = reader.uint32();
      switch (tag >>> 3) {
        case 1:
          message.creator = reader.string();
          break;
        case 2:
          message.fee = reader.string();
          break;
        case 3:
          message.bet = reader.string();
          break;
        default:
          reader.skipType(tag & 7);
          break;
      }
    }
    return message;
  },

  fromJSON(object: any): MsgSendTicket {
    return {
      creator: isSet(object.creator) ? String(object.creator) : "",
      fee: isSet(object.fee) ? String(object.fee) : "",
      bet: isSet(object.bet) ? String(object.bet) : "",
    };
  },

  toJSON(message: MsgSendTicket): unknown {
    const obj: any = {};
    message.creator !== undefined && (obj.creator = message.creator);
    message.fee !== undefined && (obj.fee = message.fee);
    message.bet !== undefined && (obj.bet = message.bet);
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSendTicket>, I>>(object: I): MsgSendTicket {
    const message = createBaseMsgSendTicket();
    message.creator = object.creator ?? "";
    message.fee = object.fee ?? "";
    message.bet = object.bet ?? "";
    return message;
  },
};

function createBaseMsgSendTicketResponse(): MsgSendTicketResponse {
  return {};
}

export const MsgSendTicketResponse = {
  encode(_: MsgSendTicketResponse, writer: _m0.Writer = _m0.Writer.create()): _m0.Writer {
    return writer;
  },

  decode(input: _m0.Reader | Uint8Array, length?: number): MsgSendTicketResponse {
    const reader = input instanceof _m0.Reader ? input : new _m0.Reader(input);
    let end = length === undefined ? reader.len : reader.pos + length;
    const message = createBaseMsgSendTicketResponse();
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

  fromJSON(_: any): MsgSendTicketResponse {
    return {};
  },

  toJSON(_: MsgSendTicketResponse): unknown {
    const obj: any = {};
    return obj;
  },

  fromPartial<I extends Exact<DeepPartial<MsgSendTicketResponse>, I>>(_: I): MsgSendTicketResponse {
    const message = createBaseMsgSendTicketResponse();
    return message;
  },
};

/** Msg defines the Msg service. */
export interface Msg {
  /** this line is used by starport scaffolding # proto/tx/rpc */
  SendTicket(request: MsgSendTicket): Promise<MsgSendTicketResponse>;
}

export class MsgClientImpl implements Msg {
  private readonly rpc: Rpc;
  constructor(rpc: Rpc) {
    this.rpc = rpc;
    this.SendTicket = this.SendTicket.bind(this);
  }
  SendTicket(request: MsgSendTicket): Promise<MsgSendTicketResponse> {
    const data = MsgSendTicket.encode(request).finish();
    const promise = this.rpc.request("lavalottery.lavalottery.Msg", "SendTicket", data);
    return promise.then((data) => MsgSendTicketResponse.decode(new _m0.Reader(data)));
  }
}

interface Rpc {
  request(service: string, method: string, data: Uint8Array): Promise<Uint8Array>;
}

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
