import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class WorldWaypoint extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): WorldWaypoint;

  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): WorldWaypoint;

  getCellid(): Uint8Array | string;
  getCellid_asU8(): Uint8Array;
  getCellid_asB64(): string;
  setCellid(value: Uint8Array | string): WorldWaypoint;

  getX(): number;
  setX(value: number): WorldWaypoint;

  getY(): number;
  setY(value: number): WorldWaypoint;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WorldWaypoint.AsObject;
  static toObject(includeInstance: boolean, msg: WorldWaypoint): WorldWaypoint.AsObject;
  static serializeBinaryToWriter(message: WorldWaypoint, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WorldWaypoint;
  static deserializeBinaryFromReader(message: WorldWaypoint, reader: jspb.BinaryReader): WorldWaypoint;
}

export namespace WorldWaypoint {
  export type AsObject = {
    id: Uint8Array | string,
    worldid: Uint8Array | string,
    cellid: Uint8Array | string,
    x: number,
    y: number,
  }
}

