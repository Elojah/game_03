// package: entity
// file: github.com/elojah/game_03/pkg/entity/pc.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_entity_entity_pb from "../../../../../github.com/elojah/game_03/pkg/entity/entity_pb";

export class PC extends jspb.Message {
  getUserid(): Uint8Array | string;
  getUserid_asU8(): Uint8Array;
  getUserid_asB64(): string;
  setUserid(value: Uint8Array | string): void;

  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): void;

  hasEntity(): boolean;
  clearEntity(): void;
  getEntity(): github_com_elojah_game_03_pkg_entity_entity_pb.E | undefined;
  setEntity(value?: github_com_elojah_game_03_pkg_entity_entity_pb.E): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PC.AsObject;
  static toObject(includeInstance: boolean, msg: PC): PC.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PC, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PC;
  static deserializeBinaryFromReader(message: PC, reader: jspb.BinaryReader): PC;
}

export namespace PC {
  export type AsObject = {
    userid: Uint8Array | string,
    worldid: Uint8Array | string,
    entity?: github_com_elojah_game_03_pkg_entity_entity_pb.E.AsObject,
  }
}

