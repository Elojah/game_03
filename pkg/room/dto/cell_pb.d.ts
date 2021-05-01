// package: dto
// file: github.com/elojah/game_03/pkg/room/dto/cell.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_entity_entity_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/entity_pb";
import * as github_com_elojah_game_03_pkg_room_cell_pb from "../../../../../../github.com/elojah/game_03/pkg/room/cell_pb";

export class Entity extends jspb.Message {
  hasEntity(): boolean;
  clearEntity(): void;
  getEntity(): github_com_elojah_game_03_pkg_entity_entity_pb.E | undefined;
  setEntity(value?: github_com_elojah_game_03_pkg_entity_entity_pb.E): void;

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
    entity?: github_com_elojah_game_03_pkg_entity_entity_pb.E.AsObject,
  }
}

export class Cell extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  clearEntitiesList(): void;
  getEntitiesList(): Array<Entity>;
  setEntitiesList(value: Array<Entity>): void;
  addEntities(value?: Entity, index?: number): Entity;

  getTs(): number;
  setTs(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Cell.AsObject;
  static toObject(includeInstance: boolean, msg: Cell): Cell.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Cell, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Cell;
  static deserializeBinaryFromReader(message: Cell, reader: jspb.BinaryReader): Cell;
}

export namespace Cell {
  export type AsObject = {
    id: Uint8Array | string,
    entitiesList: Array<Entity.AsObject>,
    ts: number,
  }
}

export class ListCellReq extends jspb.Message {
  clearIdsList(): void;
  getIdsList(): Array<Uint8Array | string>;
  getIdsList_asU8(): Array<Uint8Array>;
  getIdsList_asB64(): Array<string>;
  setIdsList(value: Array<Uint8Array | string>): void;
  addIds(value: Uint8Array | string, index?: number): Uint8Array | string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCellReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListCellReq): ListCellReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListCellReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCellReq;
  static deserializeBinaryFromReader(message: ListCellReq, reader: jspb.BinaryReader): ListCellReq;
}

export namespace ListCellReq {
  export type AsObject = {
    idsList: Array<Uint8Array | string>,
  }
}

export class ListCellResp extends jspb.Message {
  clearCellsList(): void;
  getCellsList(): Array<github_com_elojah_game_03_pkg_room_cell_pb.Cell>;
  setCellsList(value: Array<github_com_elojah_game_03_pkg_room_cell_pb.Cell>): void;
  addCells(value?: github_com_elojah_game_03_pkg_room_cell_pb.Cell, index?: number): github_com_elojah_game_03_pkg_room_cell_pb.Cell;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListCellResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListCellResp): ListCellResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListCellResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListCellResp;
  static deserializeBinaryFromReader(message: ListCellResp, reader: jspb.BinaryReader): ListCellResp;
}

export namespace ListCellResp {
  export type AsObject = {
    cellsList: Array<github_com_elojah_game_03_pkg_room_cell_pb.Cell.AsObject>,
  }
}

