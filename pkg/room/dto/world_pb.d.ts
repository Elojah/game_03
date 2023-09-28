import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';
import * as github_com_elojah_game_03_pkg_room_world_pb from '../../../../../../github.com/elojah/game_03/pkg/room/world_pb';


export class ListWorldReq extends jspb.Message {
  getIdsList(): Array<Uint8Array | string>;
  setIdsList(value: Array<Uint8Array | string>): ListWorldReq;
  clearIdsList(): ListWorldReq;
  addIds(value: Uint8Array | string, index?: number): ListWorldReq;

  getAll(): boolean;
  setAll(value: boolean): ListWorldReq;

  getSize(): number;
  setSize(value: number): ListWorldReq;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): ListWorldReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListWorldReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListWorldReq): ListWorldReq.AsObject;
  static serializeBinaryToWriter(message: ListWorldReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListWorldReq;
  static deserializeBinaryFromReader(message: ListWorldReq, reader: jspb.BinaryReader): ListWorldReq;
}

export namespace ListWorldReq {
  export type AsObject = {
    idsList: Array<Uint8Array | string>,
    all: boolean,
    size: number,
    state: Uint8Array | string,
  }
}

export class ListWorldResp extends jspb.Message {
  getWorldsList(): Array<github_com_elojah_game_03_pkg_room_world_pb.World>;
  setWorldsList(value: Array<github_com_elojah_game_03_pkg_room_world_pb.World>): ListWorldResp;
  clearWorldsList(): ListWorldResp;
  addWorlds(value?: github_com_elojah_game_03_pkg_room_world_pb.World, index?: number): github_com_elojah_game_03_pkg_room_world_pb.World;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): ListWorldResp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListWorldResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListWorldResp): ListWorldResp.AsObject;
  static serializeBinaryToWriter(message: ListWorldResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListWorldResp;
  static deserializeBinaryFromReader(message: ListWorldResp, reader: jspb.BinaryReader): ListWorldResp;
}

export namespace ListWorldResp {
  export type AsObject = {
    worldsList: Array<github_com_elojah_game_03_pkg_room_world_pb.World.AsObject>,
    state: Uint8Array | string,
  }
}

