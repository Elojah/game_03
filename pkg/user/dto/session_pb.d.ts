import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class CreateSessionReq extends jspb.Message {
  getPcid(): Uint8Array | string;
  getPcid_asU8(): Uint8Array;
  getPcid_asB64(): string;
  setPcid(value: Uint8Array | string): CreateSessionReq;

  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): CreateSessionReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSessionReq.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSessionReq): CreateSessionReq.AsObject;
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
  setToken(value: Uint8Array | string): CreateSessionResp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSessionResp.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSessionResp): CreateSessionResp.AsObject;
  static serializeBinaryToWriter(message: CreateSessionResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSessionResp;
  static deserializeBinaryFromReader(message: CreateSessionResp, reader: jspb.BinaryReader): CreateSessionResp;
}

export namespace CreateSessionResp {
  export type AsObject = {
    token: Uint8Array | string,
  }
}

