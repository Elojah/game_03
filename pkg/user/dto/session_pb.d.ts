// package: dto
// file: github.com/elojah/game_03/pkg/user/dto/session.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class CreateSessionReq extends jspb.Message {
  getPcid(): Uint8Array | string;
  getPcid_asU8(): Uint8Array;
  getPcid_asB64(): string;
  setPcid(value: Uint8Array | string): void;

  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSessionReq.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSessionReq): CreateSessionReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateSessionReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSessionReq;
  static deserializeBinaryFromReader(message: CreateSessionReq, reader: jspb.BinaryReader): CreateSessionReq;
}

export namespace CreateSessionReq {
  export type AsObject = {
    pcid: Uint8Array | string,
    worldid: Uint8Array | string,
  }
}

export class CreateSessionResp extends jspb.Message {
  getToken(): Uint8Array | string;
  getToken_asU8(): Uint8Array;
  getToken_asB64(): string;
  setToken(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSessionResp.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSessionResp): CreateSessionResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateSessionResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSessionResp;
  static deserializeBinaryFromReader(message: CreateSessionResp, reader: jspb.BinaryReader): CreateSessionResp;
}

export namespace CreateSessionResp {
  export type AsObject = {
    token: Uint8Array | string,
  }
}

export class SDP extends jspb.Message {
  getEncoded(): string;
  setEncoded(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SDP.AsObject;
  static toObject(includeInstance: boolean, msg: SDP): SDP.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SDP, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SDP;
  static deserializeBinaryFromReader(message: SDP, reader: jspb.BinaryReader): SDP;
}

export namespace SDP {
  export type AsObject = {
    encoded: string,
  }
}

