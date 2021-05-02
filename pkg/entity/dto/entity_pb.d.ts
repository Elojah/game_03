// package: dto
// file: github.com/elojah/game_03/pkg/entity/dto/entity.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_entity_entity_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/entity_pb";

export class ListEntityReq extends jspb.Message {
  clearIdsList(): void;
  getIdsList(): Array<Uint8Array | string>;
  getIdsList_asU8(): Array<Uint8Array>;
  getIdsList_asB64(): Array<string>;
  setIdsList(value: Array<Uint8Array | string>): void;
  addIds(value: Uint8Array | string, index?: number): Uint8Array | string;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListEntityReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListEntityReq): ListEntityReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListEntityReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListEntityReq;
  static deserializeBinaryFromReader(message: ListEntityReq, reader: jspb.BinaryReader): ListEntityReq;
}

export namespace ListEntityReq {
  export type AsObject = {
    idsList: Array<Uint8Array | string>,
  }
}

export class ListEntityResp extends jspb.Message {
  clearEntitiesList(): void;
  getEntitiesList(): Array<github_com_elojah_game_03_pkg_entity_entity_pb.E>;
  setEntitiesList(value: Array<github_com_elojah_game_03_pkg_entity_entity_pb.E>): void;
  addEntities(value?: github_com_elojah_game_03_pkg_entity_entity_pb.E, index?: number): github_com_elojah_game_03_pkg_entity_entity_pb.E;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListEntityResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListEntityResp): ListEntityResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListEntityResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListEntityResp;
  static deserializeBinaryFromReader(message: ListEntityResp, reader: jspb.BinaryReader): ListEntityResp;
}

export namespace ListEntityResp {
  export type AsObject = {
    entitiesList: Array<github_com_elojah_game_03_pkg_entity_entity_pb.E.AsObject>,
  }
}

