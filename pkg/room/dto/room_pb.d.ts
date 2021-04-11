// package: dto
// file: github.com/elojah/game_03/pkg/room/dto/room.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_room_room_pb from "../../../../../../github.com/elojah/game_03/pkg/room/room_pb";

export class ListRoomReq extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoomReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoomReq): ListRoomReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListRoomReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoomReq;
  static deserializeBinaryFromReader(message: ListRoomReq, reader: jspb.BinaryReader): ListRoomReq;
}

export namespace ListRoomReq {
  export type AsObject = {
  }
}

export class ListRoomResp extends jspb.Message {
  clearRoomsList(): void;
  getRoomsList(): Array<github_com_elojah_game_03_pkg_room_room_pb.R>;
  setRoomsList(value: Array<github_com_elojah_game_03_pkg_room_room_pb.R>): void;
  addRooms(value?: github_com_elojah_game_03_pkg_room_room_pb.R, index?: number): github_com_elojah_game_03_pkg_room_room_pb.R;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListRoomResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListRoomResp): ListRoomResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListRoomResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListRoomResp;
  static deserializeBinaryFromReader(message: ListRoomResp, reader: jspb.BinaryReader): ListRoomResp;
}

export namespace ListRoomResp {
  export type AsObject = {
    roomsList: Array<github_com_elojah_game_03_pkg_room_room_pb.R.AsObject>,
  }
}

