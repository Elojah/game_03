import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';
import * as github_com_elojah_game_03_pkg_entity_pc_pb from '../../../../../../github.com/elojah/game_03/pkg/entity/pc_pb';
import * as github_com_elojah_game_03_pkg_entity_entity_pb from '../../../../../../github.com/elojah/game_03/pkg/entity/entity_pb';


export class PC extends jspb.Message {
  getPc(): github_com_elojah_game_03_pkg_entity_pc_pb.PC | undefined;
  setPc(value?: github_com_elojah_game_03_pkg_entity_pc_pb.PC): PC;
  hasPc(): boolean;
  clearPc(): PC;

  getEntity(): github_com_elojah_game_03_pkg_entity_entity_pb.E | undefined;
  setEntity(value?: github_com_elojah_game_03_pkg_entity_entity_pb.E): PC;
  hasEntity(): boolean;
  clearEntity(): PC;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PC.AsObject;
  static toObject(includeInstance: boolean, msg: PC): PC.AsObject;
  static serializeBinaryToWriter(message: PC, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PC;
  static deserializeBinaryFromReader(message: PC, reader: jspb.BinaryReader): PC;
}

export namespace PC {
  export type AsObject = {
    pc?: github_com_elojah_game_03_pkg_entity_pc_pb.PC.AsObject,
    entity?: github_com_elojah_game_03_pkg_entity_entity_pb.E.AsObject,
  }
}

export class CreatePCReq extends jspb.Message {
  getRoomid(): Uint8Array | string;
  getRoomid_asU8(): Uint8Array;
  getRoomid_asB64(): string;
  setRoomid(value: Uint8Array | string): CreatePCReq;

  getEntitytemplate(): string;
  setEntitytemplate(value: string): CreatePCReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreatePCReq.AsObject;
  static toObject(includeInstance: boolean, msg: CreatePCReq): CreatePCReq.AsObject;
  static serializeBinaryToWriter(message: CreatePCReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreatePCReq;
  static deserializeBinaryFromReader(message: CreatePCReq, reader: jspb.BinaryReader): CreatePCReq;
}

export namespace CreatePCReq {
  export type AsObject = {
    roomid: Uint8Array | string,
    entitytemplate: string,
  }
}

export class ListPCReq extends jspb.Message {
  getRoomid(): Uint8Array | string;
  getRoomid_asU8(): Uint8Array;
  getRoomid_asB64(): string;
  setRoomid(value: Uint8Array | string): ListPCReq;

  getSize(): number;
  setSize(value: number): ListPCReq;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): ListPCReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPCReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListPCReq): ListPCReq.AsObject;
  static serializeBinaryToWriter(message: ListPCReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPCReq;
  static deserializeBinaryFromReader(message: ListPCReq, reader: jspb.BinaryReader): ListPCReq;
}

export namespace ListPCReq {
  export type AsObject = {
    roomid: Uint8Array | string,
    size: number,
    state: Uint8Array | string,
  }
}

export class ListPCResp extends jspb.Message {
  getPcsList(): Array<PC>;
  setPcsList(value: Array<PC>): ListPCResp;
  clearPcsList(): ListPCResp;
  addPcs(value?: PC, index?: number): PC;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): ListPCResp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPCResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListPCResp): ListPCResp.AsObject;
  static serializeBinaryToWriter(message: ListPCResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPCResp;
  static deserializeBinaryFromReader(message: ListPCResp, reader: jspb.BinaryReader): ListPCResp;
}

export namespace ListPCResp {
  export type AsObject = {
    pcsList: Array<PC.AsObject>,
    state: Uint8Array | string,
  }
}

export class GetPCReq extends jspb.Message {
  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): GetPCReq;

  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): GetPCReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): GetPCReq.AsObject;
  static toObject(includeInstance: boolean, msg: GetPCReq): GetPCReq.AsObject;
  static serializeBinaryToWriter(message: GetPCReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): GetPCReq;
  static deserializeBinaryFromReader(message: GetPCReq, reader: jspb.BinaryReader): GetPCReq;
}

export namespace GetPCReq {
  export type AsObject = {
    worldid: Uint8Array | string,
    id: Uint8Array | string,
  }
}

