import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class Cell extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): Cell;

  getContiguousMap(): jspb.Map<number, Uint8Array | string>;
  clearContiguousMap(): Cell;

  getTilemap(): Uint8Array | string;
  getTilemap_asU8(): Uint8Array;
  getTilemap_asB64(): string;
  setTilemap(value: Uint8Array | string): Cell;

  getX(): number;
  setX(value: number): Cell;

  getY(): number;
  setY(value: number): Cell;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Cell.AsObject;
  static toObject(includeInstance: boolean, msg: Cell): Cell.AsObject;
  static serializeBinaryToWriter(message: Cell, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Cell;
  static deserializeBinaryFromReader(message: Cell, reader: jspb.BinaryReader): Cell;
}

export namespace Cell {
  export type AsObject = {
    id: Uint8Array | string,
    contiguousMap: Array<[number, Uint8Array | string]>,
    tilemap: Uint8Array | string,
    x: number,
    y: number,
  }
}

export class WorldCell extends jspb.Message {
  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): WorldCell;

  getCellid(): Uint8Array | string;
  getCellid_asU8(): Uint8Array;
  getCellid_asB64(): string;
  setCellid(value: Uint8Array | string): WorldCell;

  getX(): number;
  setX(value: number): WorldCell;

  getY(): number;
  setY(value: number): WorldCell;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WorldCell.AsObject;
  static toObject(includeInstance: boolean, msg: WorldCell): WorldCell.AsObject;
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

export enum Orientation { 
  NONE = 0,
  UP = 1,
  UPRIGHT = 2,
  RIGHT = 3,
  DOWNRIGHT = 4,
  DOWN = 5,
  DOWNLEFT = 6,
  LEFT = 7,
  UPLEFT = 8,
}
