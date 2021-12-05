// package: dto
// file: github.com/elojah/game_03/pkg/room/dto/world.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_room_world_pb from "../../../../../../github.com/elojah/game_03/pkg/room/world_pb";

export class ListWorldReq extends jspb.Message {
  clearIdsList(): void;
  getIdsList(): Array<Uint8Array | string>;
  getIdsList_asU8(): Array<Uint8Array>;
  getIdsList_asB64(): Array<string>;
  setIdsList(value: Array<Uint8Array | string>): void;
  addIds(value: Uint8Array | string, index?: number): Uint8Array | string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListWorldReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListWorldReq): ListWorldReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListWorldReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListWorldReq;
  static deserializeBinaryFromReader(message: ListWorldReq, reader: jspb.BinaryReader): ListWorldReq;
}

export namespace ListWorldReq {
  export type AsObject = {
    idsList: Array<Uint8Array | string>,
  }
}

export class ListWorldResp extends jspb.Message {
  clearWorldsList(): void;
  getWorldsList(): Array<github_com_elojah_game_03_pkg_room_world_pb.World>;
  setWorldsList(value: Array<github_com_elojah_game_03_pkg_room_world_pb.World>): void;
  addWorlds(value?: github_com_elojah_game_03_pkg_room_world_pb.World, index?: number): github_com_elojah_game_03_pkg_room_world_pb.World;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListWorldResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListWorldResp): ListWorldResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListWorldResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListWorldResp;
  static deserializeBinaryFromReader(message: ListWorldResp, reader: jspb.BinaryReader): ListWorldResp;
}

export namespace ListWorldResp {
  export type AsObject = {
    worldsList: Array<github_com_elojah_game_03_pkg_room_world_pb.World.AsObject>,
  }
}

