// package: event
// file: github.com/elojah/game_03/pkg/event/event.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_ability_cast_pb from "../../../../../github.com/elojah/game_03/pkg/ability/cast_pb";

export class E extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): void;

  getSourceid(): Uint8Array | string;
  getSourceid_asU8(): Uint8Array;
  getSourceid_asB64(): string;
  setSourceid(value: Uint8Array | string): void;

  getAt(): number;
  setAt(value: number): void;

  hasEffect(): boolean;
  clearEffect(): void;
  getEffect(): github_com_elojah_game_03_pkg_ability_cast_pb.CastEffect | undefined;
  setEffect(value?: github_com_elojah_game_03_pkg_ability_cast_pb.CastEffect): void;

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
    entityid: Uint8Array | string,
    sourceid: Uint8Array | string,
    at: number,
    effect?: github_com_elojah_game_03_pkg_ability_cast_pb.CastEffect.AsObject,
  }
}
