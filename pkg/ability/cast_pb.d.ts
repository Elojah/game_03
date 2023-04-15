// package: ability
// file: github.com/elojah/game_03/pkg/ability/cast.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_geometry_geometry_pb from "../../../../../github.com/elojah/game_03/pkg/geometry/geometry_pb";

export class CastTarget extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  hasRect(): boolean;
  clearRect(): void;
  getRect(): github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect | undefined;
  setRect(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect): void;

  hasCircle(): boolean;
  clearCircle(): void;
  getCircle(): github_com_elojah_game_03_pkg_geometry_geometry_pb.Circle | undefined;
  setCircle(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Circle): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CastTarget.AsObject;
  static toObject(includeInstance: boolean, msg: CastTarget): CastTarget.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CastTarget, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CastTarget;
  static deserializeBinaryFromReader(message: CastTarget, reader: jspb.BinaryReader): CastTarget;
}

export namespace CastTarget {
  export type AsObject = {
    id: Uint8Array | string,
    rect?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect.AsObject,
    circle?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Circle.AsObject,
  }
}

export class Cast extends jspb.Message {
  getSourceid(): Uint8Array | string;
  getSourceid_asU8(): Uint8Array;
  getSourceid_asB64(): string;
  setSourceid(value: Uint8Array | string): void;

  getAbilityid(): Uint8Array | string;
  getAbilityid_asU8(): Uint8Array;
  getAbilityid_asB64(): string;
  setAbilityid(value: Uint8Array | string): void;

  getTargetsMap(): jspb.Map<string, CastTarget>;
  clearTargetsMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Cast.AsObject;
  static toObject(includeInstance: boolean, msg: Cast): Cast.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Cast, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Cast;
  static deserializeBinaryFromReader(message: Cast, reader: jspb.BinaryReader): Cast;
}

export namespace Cast {
  export type AsObject = {
    sourceid: Uint8Array | string,
    abilityid: Uint8Array | string,
    targetsMap: Array<[string, CastTarget.AsObject]>,
  }
}

