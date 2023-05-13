// package: dto
// file: github.com/elojah/game_03/pkg/ability/dto/ability.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_ability_ability_pb from "../../../../../../github.com/elojah/game_03/pkg/ability/ability_pb";

export class ListAbilityReq extends jspb.Message {
  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): void;

  getSize(): number;
  setSize(value: number): void;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAbilityReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListAbilityReq): ListAbilityReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListAbilityReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAbilityReq;
  static deserializeBinaryFromReader(message: ListAbilityReq, reader: jspb.BinaryReader): ListAbilityReq;
}

export namespace ListAbilityReq {
  export type AsObject = {
    entityid: Uint8Array | string,
    size: number,
    state: Uint8Array | string,
  }
}

export class ListAbilityResp extends jspb.Message {
  clearAbilitiesList(): void;
  getAbilitiesList(): Array<github_com_elojah_game_03_pkg_ability_ability_pb.A>;
  setAbilitiesList(value: Array<github_com_elojah_game_03_pkg_ability_ability_pb.A>): void;
  addAbilities(value?: github_com_elojah_game_03_pkg_ability_ability_pb.A, index?: number): github_com_elojah_game_03_pkg_ability_ability_pb.A;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAbilityResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListAbilityResp): ListAbilityResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListAbilityResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAbilityResp;
  static deserializeBinaryFromReader(message: ListAbilityResp, reader: jspb.BinaryReader): ListAbilityResp;
}

export namespace ListAbilityResp {
  export type AsObject = {
    abilitiesList: Array<github_com_elojah_game_03_pkg_ability_ability_pb.A.AsObject>,
    state: Uint8Array | string,
  }
}

