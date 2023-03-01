// package: dto
// file: github.com/elojah/game_03/pkg/room/dto/user.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class CreateRoomUserReq extends jspb.Message {
  getRoomid(): Uint8Array | string;
  getRoomid_asU8(): Uint8Array;
  getRoomid_asB64(): string;
  setRoomid(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateRoomUserReq.AsObject;
  static toObject(includeInstance: boolean, msg: CreateRoomUserReq): CreateRoomUserReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateRoomUserReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateRoomUserReq;
  static deserializeBinaryFromReader(message: CreateRoomUserReq, reader: jspb.BinaryReader): CreateRoomUserReq;
}

export namespace CreateRoomUserReq {
  export type AsObject = {
    roomid: Uint8Array | string,
  }
}

