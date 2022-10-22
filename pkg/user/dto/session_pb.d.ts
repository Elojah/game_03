// package: dto
// file: github.com/elojah/game_03/pkg/user/dto/session.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class CreateSessionReq extends jspb.Message {
  getPcid(): Uint8Array | string;
  getPcid_asU8(): Uint8Array;
  getPcid_asB64(): string;
  setPcid(value: Uint8Array | string): void;

  getRoomid(): Uint8Array | string;
  getRoomid_asU8(): Uint8Array;
  getRoomid_asB64(): string;
  setRoomid(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateSessionReq.AsObject;
  static toObject(includeInstance: boolean, msg: CreateSessionReq): CreateSessionReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateSessionReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateSessionReq;
  static deserializeBinaryFromReader(message: CreateSessionReq, reader: jspb.BinaryReader): CreateSessionReq;
}

export namespace CreateSessionReq {
  export type AsObject = {
    pcid: Uint8Array | string,
    roomid: Uint8Array | string,
  }
}

