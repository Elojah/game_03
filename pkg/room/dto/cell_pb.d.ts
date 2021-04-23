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
  hasCell(): boolean;
  clearCell(): void;
  getCell(): github_com_elojah_game_03_pkg_room_cell_pb.Cell | undefined;
  setCell(value?: github_com_elojah_game_03_pkg_room_cell_pb.Cell): void;

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
    cell?: github_com_elojah_game_03_pkg_room_cell_pb.Cell.AsObject,
    entitiesList: Array<Entity.AsObject>,
    ts: number,
  }
}

