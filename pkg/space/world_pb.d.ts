// package: space
// file: github.com/elojah/game_03/pkg/space/world.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_geometry_geometry_pb from "../../../../../github.com/elojah/game_03/pkg/geometry/geometry_pb";

export class World extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getOwnerid(): Uint8Array | string;
  getOwnerid_asU8(): Uint8Array;
  getOwnerid_asB64(): string;
  setOwnerid(value: Uint8Array | string): void;

  hasNsection(): boolean;
  clearNsection(): void;
  getNsection(): github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2 | undefined;
  setNsection(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2): void;

  hasDimsection(): boolean;
  clearDimsection(): void;
  getDimsection(): github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2 | undefined;
  setDimsection(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): World.AsObject;
  static toObject(includeInstance: boolean, msg: World): World.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: World, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): World;
  static deserializeBinaryFromReader(message: World, reader: jspb.BinaryReader): World;
}

export namespace World {
  export type AsObject = {
    id: Uint8Array | string,
    ownerid: Uint8Array | string,
    nsection?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.AsObject,
    dimsection?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.AsObject,
  }
}

export class Section extends jspb.Message {
  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): void;

  hasCoord(): boolean;
  clearCoord(): void;
  getCoord(): github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2 | undefined;
  setCoord(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Section.AsObject;
  static toObject(includeInstance: boolean, msg: Section): Section.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Section, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Section;
  static deserializeBinaryFromReader(message: Section, reader: jspb.BinaryReader): Section;
}

export namespace Section {
  export type AsObject = {
    worldid: Uint8Array | string,
    coord?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.AsObject,
  }
}

