import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class CreateTilesetReq extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): CreateTilesetReq;

  getSet(): Uint8Array | string;
  getSet_asU8(): Uint8Array;
  getSet_asB64(): string;
  setSet(value: Uint8Array | string): CreateTilesetReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTilesetReq.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTilesetReq): CreateTilesetReq.AsObject;
  static serializeBinaryToWriter(message: CreateTilesetReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTilesetReq;
  static deserializeBinaryFromReader(message: CreateTilesetReq, reader: jspb.BinaryReader): CreateTilesetReq;
}

export namespace CreateTilesetReq {
  export type AsObject = {
    id: Uint8Array | string,
    set: Uint8Array | string,
  }
}

export class CreateTilesetResp extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): CreateTilesetResp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTilesetResp.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTilesetResp): CreateTilesetResp.AsObject;
  static serializeBinaryToWriter(message: CreateTilesetResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTilesetResp;
  static deserializeBinaryFromReader(message: CreateTilesetResp, reader: jspb.BinaryReader): CreateTilesetResp;
}

export namespace CreateTilesetResp {
  export type AsObject = {
    id: Uint8Array | string,
  }
}

