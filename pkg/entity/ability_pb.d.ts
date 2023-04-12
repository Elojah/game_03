// package: entity
// file: github.com/elojah/game_03/pkg/entity/ability.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class Ability extends jspb.Message {
  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): void;

  getAbilityid(): Uint8Array | string;
  getAbilityid_asU8(): Uint8Array;
  getAbilityid_asB64(): string;
  setAbilityid(value: Uint8Array | string): void;

  getLastcast(): number;
  setLastcast(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Ability.AsObject;
  static toObject(includeInstance: boolean, msg: Ability): Ability.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Ability, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Ability;
  static deserializeBinaryFromReader(message: Ability, reader: jspb.BinaryReader): Ability;
}

export namespace Ability {
  export type AsObject = {
    entityid: Uint8Array | string,
    abilityid: Uint8Array | string,
    lastcast: number,
  }
}

