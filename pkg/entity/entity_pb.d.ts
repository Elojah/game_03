// package: entity
// file: github.com/elojah/game_03/pkg/entity/entity.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_geometry_geometry_pb from "../../../../../github.com/elojah/game_03/pkg/geometry/geometry_pb";
import * as github_com_elojah_game_03_pkg_space_world_pb from "../../../../../github.com/elojah/game_03/pkg/space/world_pb";

export class E extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  hasSection(): boolean;
  clearSection(): void;
  getSection(): github_com_elojah_game_03_pkg_space_world_pb.Section | undefined;
  setSection(value?: github_com_elojah_game_03_pkg_space_world_pb.Section): void;

  hasCoord(): boolean;
  clearCoord(): void;
  getCoord(): github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2 | undefined;
  setCoord(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2): void;

  getMaxhp(): number;
  setMaxhp(value: number): void;

  getHp(): number;
  setHp(value: number): void;

  getMaxmp(): number;
  setMaxmp(value: number): void;

  getMp(): number;
  setMp(value: number): void;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): E.AsObject;
  static toObject(includeInstance: boolean, msg: E): E.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: E, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): E;
  static deserializeBinaryFromReader(message: E, reader: jspb.BinaryReader): E;
}

export namespace E {
  export type AsObject = {
    id: Uint8Array | string,
    section?: github_com_elojah_game_03_pkg_space_world_pb.Section.AsObject,
    coord?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.AsObject,
    maxhp: number,
    hp: number,
    maxmp: number,
    mp: number,
    state: Uint8Array | string,
  }
}

