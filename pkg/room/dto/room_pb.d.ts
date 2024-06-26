import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';
import * as github_com_elojah_game_03_pkg_room_room_pb from '../../../../../../github.com/elojah/game_03/pkg/room/room_pb';
import * as github_com_elojah_game_03_pkg_user_user_pb from '../../../../../../github.com/elojah/game_03/pkg/user/user_pb';


export class Room extends jspb.Message {
  getRoom(): github_com_elojah_game_03_pkg_room_room_pb.R | undefined;
  setRoom(value?: github_com_elojah_game_03_pkg_room_room_pb.R): Room;
  hasRoom(): boolean;
  clearRoom(): Room;

  getOwner(): github_com_elojah_game_03_pkg_user_user_pb.U | undefined;
  setOwner(value?: github_com_elojah_game_03_pkg_user_user_pb.U): Room;
  hasOwner(): boolean;
  clearOwner(): Room;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Room.AsObject;
  static toObject(includeInstance: boolean, msg: Room): Room.AsObject;
  static serializeBinaryToWriter(message: Room, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Room;
  static deserializeBinaryFromReader(message: Room, reader: jspb.BinaryReader): Room;
}

export namespace Room {
  export type AsObject = {
    room?: github_com_elojah_game_03_pkg_room_room_pb.R.AsObject,
    owner?: github_com_elojah_game_03_pkg_user_user_pb.U.AsObject,
  }
}

export class ListRoomReq extends jspb.Message {
  getSize(): number;
  setSize(value: number): ListRoomReq;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): ListRoomReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoomReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoomReq): ListRoomReq.AsObject;
  static serializeBinaryToWriter(message: ListRoomReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoomReq;
  static deserializeBinaryFromReader(message: ListRoomReq, reader: jspb.BinaryReader): ListRoomReq;
}

export namespace ListRoomReq {
  export type AsObject = {
    size: number,
    state: Uint8Array | string,
  }
}

export class ListRoomResp extends jspb.Message {
  getRoomsList(): Array<Room>;
  setRoomsList(value: Array<Room>): ListRoomResp;
  clearRoomsList(): ListRoomResp;
  addRooms(value?: Room, index?: number): Room;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): ListRoomResp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoomResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoomResp): ListRoomResp.AsObject;
  static serializeBinaryToWriter(message: ListRoomResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoomResp;
  static deserializeBinaryFromReader(message: ListRoomResp, reader: jspb.BinaryReader): ListRoomResp;
}

export namespace ListRoomResp {
  export type AsObject = {
    roomsList: Array<Room.AsObject>,
    state: Uint8Array | string,
  }
}

