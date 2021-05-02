// package: entity
// file: github.com/elojah/game_03/pkg/entity/entity.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class E extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getUserid(): Uint8Array | string;
  getUserid_asU8(): Uint8Array;
  getUserid_asB64(): string;
  setUserid(value: Uint8Array | string): void;

  getCellid(): Uint8Array | string;
  getCellid_asU8(): Uint8Array;
  getCellid_asB64(): string;
  setCellid(value: Uint8Array | string): void;

  getX(): number;
  setX(value: number): void;

  getY(): number;
  setY(value: number): void;

  getRot(): number;
  setRot(value: number): void;

  getRadius(): number;
  setRadius(value: number): void;

  getTilemap(): Uint8Array | string;
  getTilemap_asU8(): Uint8Array;
  getTilemap_asB64(): string;
  setTilemap(value: Uint8Array | string): void;

  getTileset(): Uint8Array | string;
  getTileset_asU8(): Uint8Array;
  getTileset_asB64(): string;
  setTileset(value: Uint8Array | string): void;

  getAt(): number;
  setAt(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): E.AsObject;
  static toObject(includeInstance: boolean, msg: E): E.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: E, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): E;
  static deserializeBinaryFromReader(message: E, reader: jspb.BinaryReader): E;
}

export namespace E {
  export type AsObject = {
    id: Uint8Array | string,
    userid: Uint8Array | string,
    cellid: Uint8Array | string,
    x: number,
    y: number,
    rot: number,
    radius: number,
    tilemap: Uint8Array | string,
    tileset: Uint8Array | string,
    at: number,
  }
}

