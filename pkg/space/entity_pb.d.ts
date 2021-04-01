// package: space
// file: github.com/elojah/game_03/pkg/space/entity.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_space_world_pb from "../../../../../github.com/elojah/game_03/pkg/space/world_pb";

export class Entity extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  hasSection(): boolean;
  clearSection(): void;
  getSection(): github_com_elojah_game_03_pkg_space_world_pb.Section | undefined;
  setSection(value?: github_com_elojah_game_03_pkg_space_world_pb.Section): void;

  getCross(): CrossMap[keyof CrossMap];
  setCross(value: CrossMap[keyof CrossMap]): void;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Entity.AsObject;
  static toObject(includeInstance: boolean, msg: Entity): Entity.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Entity, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Entity;
  static deserializeBinaryFromReader(message: Entity, reader: jspb.BinaryReader): Entity;
}

export namespace Entity {
  export type AsObject = {
    id: Uint8Array | string,
    section?: github_com_elojah_game_03_pkg_space_world_pb.Section.AsObject,
    cross: CrossMap[keyof CrossMap],
    state: Uint8Array | string,
  }
}

export interface CrossMap {
  IN: 0;
  OUT: 1;
}

export const Cross: CrossMap;

