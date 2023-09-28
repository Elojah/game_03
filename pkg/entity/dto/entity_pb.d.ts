import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';
import * as github_com_elojah_game_03_pkg_ability_ability_pb from '../../../../../../github.com/elojah/game_03/pkg/ability/ability_pb';
import * as github_com_elojah_game_03_pkg_entity_entity_pb from '../../../../../../github.com/elojah/game_03/pkg/entity/entity_pb';


export class ListEntityReq extends jspb.Message {
  getIdsList(): Array<Uint8Array | string>;
  setIdsList(value: Array<Uint8Array | string>): ListEntityReq;
  clearIdsList(): ListEntityReq;
  addIds(value: Uint8Array | string, index?: number): ListEntityReq;

  getCellidsList(): Array<Uint8Array | string>;
  setCellidsList(value: Array<Uint8Array | string>): ListEntityReq;
  clearCellidsList(): ListEntityReq;
  addCellids(value: Uint8Array | string, index?: number): ListEntityReq;

  getSize(): number;
  setSize(value: number): ListEntityReq;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): ListEntityReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListEntityReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListEntityReq): ListEntityReq.AsObject;
  static serializeBinaryToWriter(message: ListEntityReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListEntityReq;
  static deserializeBinaryFromReader(message: ListEntityReq, reader: jspb.BinaryReader): ListEntityReq;
}

export namespace ListEntityReq {
  export type AsObject = {
    idsList: Array<Uint8Array | string>,
    cellidsList: Array<Uint8Array | string>,
    size: number,
    state: Uint8Array | string,
  }
}

export class ListEntityResp extends jspb.Message {
  getEntitiesList(): Array<github_com_elojah_game_03_pkg_entity_entity_pb.E>;
  setEntitiesList(value: Array<github_com_elojah_game_03_pkg_entity_entity_pb.E>): ListEntityResp;
  clearEntitiesList(): ListEntityResp;
  addEntities(value?: github_com_elojah_game_03_pkg_entity_entity_pb.E, index?: number): github_com_elojah_game_03_pkg_entity_entity_pb.E;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): ListEntityResp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListEntityResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListEntityResp): ListEntityResp.AsObject;
  static serializeBinaryToWriter(message: ListEntityResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListEntityResp;
  static deserializeBinaryFromReader(message: ListEntityResp, reader: jspb.BinaryReader): ListEntityResp;
}

export namespace ListEntityResp {
  export type AsObject = {
    entitiesList: Array<github_com_elojah_game_03_pkg_entity_entity_pb.E.AsObject>,
    state: Uint8Array | string,
  }
}

export class CreateEntityAbilityReq extends jspb.Message {
  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): CreateEntityAbilityReq;

  getAbility(): github_com_elojah_game_03_pkg_ability_ability_pb.A | undefined;
  setAbility(value?: github_com_elojah_game_03_pkg_ability_ability_pb.A): CreateEntityAbilityReq;
  hasAbility(): boolean;
  clearAbility(): CreateEntityAbilityReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateEntityAbilityReq.AsObject;
  static toObject(includeInstance: boolean, msg: CreateEntityAbilityReq): CreateEntityAbilityReq.AsObject;
  static serializeBinaryToWriter(message: CreateEntityAbilityReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateEntityAbilityReq;
  static deserializeBinaryFromReader(message: CreateEntityAbilityReq, reader: jspb.BinaryReader): CreateEntityAbilityReq;
}

export namespace CreateEntityAbilityReq {
  export type AsObject = {
    entityid: Uint8Array | string,
    ability?: github_com_elojah_game_03_pkg_ability_ability_pb.A.AsObject,
  }
}

export class CreateEntityAbilityResp extends jspb.Message {
  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): CreateEntityAbilityResp;

  getAbilityid(): Uint8Array | string;
  getAbilityid_asU8(): Uint8Array;
  getAbilityid_asB64(): string;
  setAbilityid(value: Uint8Array | string): CreateEntityAbilityResp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateEntityAbilityResp.AsObject;
  static toObject(includeInstance: boolean, msg: CreateEntityAbilityResp): CreateEntityAbilityResp.AsObject;
  static serializeBinaryToWriter(message: CreateEntityAbilityResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateEntityAbilityResp;
  static deserializeBinaryFromReader(message: CreateEntityAbilityResp, reader: jspb.BinaryReader): CreateEntityAbilityResp;
}

export namespace CreateEntityAbilityResp {
  export type AsObject = {
    entityid: Uint8Array | string,
    abilityid: Uint8Array | string,
  }
}

