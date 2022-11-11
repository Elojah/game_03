// package: dto
// file: github.com/elojah/game_03/pkg/tile/dto/set.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class CreateTilesetReq extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getSet(): Uint8Array | string;
  getSet_asU8(): Uint8Array;
  getSet_asB64(): string;
  setSet(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTilesetReq.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTilesetReq): CreateTilesetReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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
  setId(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTilesetResp.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTilesetResp): CreateTilesetResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateTilesetResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTilesetResp;
  static deserializeBinaryFromReader(message: CreateTilesetResp, reader: jspb.BinaryReader): CreateTilesetResp;
}

export namespace CreateTilesetResp {
  export type AsObject = {
    id: Uint8Array | string,
  }
}

