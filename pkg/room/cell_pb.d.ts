// package: room
// file: github.com/elojah/game_03/pkg/room/cell.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class WorldCell extends jspb.Message {
  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): void;

  getCellid(): Uint8Array | string;
  getCellid_asU8(): Uint8Array;
  getCellid_asB64(): string;
  setCellid(value: Uint8Array | string): void;

  getX(): number;
  setX(value: number): void;

  getY(): number;
  setY(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WorldCell.AsObject;
  static toObject(includeInstance: boolean, msg: WorldCell): WorldCell.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: WorldCell, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WorldCell;
  static deserializeBinaryFromReader(message: WorldCell, reader: jspb.BinaryReader): WorldCell;
}

export namespace WorldCell {
  export type AsObject = {
    worldid: Uint8Array | string,
    cellid: Uint8Array | string,
    x: number,
    y: number,
  }
}

export class Cell extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getContiguousMap(): jspb.Map<number, Uint8Array | string>;
  clearContiguousMap(): void;
  getTilemap(): Uint8Array | string;
  getTilemap_asU8(): Uint8Array;
  getTilemap_asB64(): string;
  setTilemap(value: Uint8Array | string): void;

  getTileset(): Uint8Array | string;
  getTileset_asU8(): Uint8Array;
  getTileset_asB64(): string;
  setTileset(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Cell.AsObject;
  static toObject(includeInstance: boolean, msg: Cell): Cell.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Cell, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Cell;
  static deserializeBinaryFromReader(message: Cell, reader: jspb.BinaryReader): Cell;
}

export namespace Cell {
  export type AsObject = {
    id: Uint8Array | string,
    contiguousMap: Array<[number, Uint8Array | string]>,
    tilemap: Uint8Array | string,
    tileset: Uint8Array | string,
  }
}

