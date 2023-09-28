import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class World extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): World;

  getName(): string;
  setName(value: string): World;

  getHeight(): number;
  setHeight(value: number): World;

  getWidth(): number;
  setWidth(value: number): World;

  getCellheight(): number;
  setCellheight(value: number): World;

  getCellwidth(): number;
  setCellwidth(value: number): World;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): World.AsObject;
  static toObject(includeInstance: boolean, msg: World): World.AsObject;
  static serializeBinaryToWriter(message: World, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): World;
  static deserializeBinaryFromReader(message: World, reader: jspb.BinaryReader): World;
}

export namespace World {
  export type AsObject = {
    id: Uint8Array | string,
    name: string,
    height: number,
    width: number,
    cellheight: number,
    cellwidth: number,
  }
}

