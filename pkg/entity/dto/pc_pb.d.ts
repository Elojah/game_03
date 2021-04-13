// package: dto
// file: github.com/elojah/game_03/pkg/entity/dto/pc.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_entity_pc_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/pc_pb";

export class ListPCReq extends jspb.Message {
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPCReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListPCReq): ListPCReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPCReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPCReq;
  static deserializeBinaryFromReader(message: ListPCReq, reader: jspb.BinaryReader): ListPCReq;
}

export namespace ListPCReq {
  export type AsObject = {
  }
}

export class ListPCResp extends jspb.Message {
  clearPcsList(): void;
  getPcsList(): Array<github_com_elojah_game_03_pkg_entity_pc_pb.PC>;
  setPcsList(value: Array<github_com_elojah_game_03_pkg_entity_pc_pb.PC>): void;
  addPcs(value?: github_com_elojah_game_03_pkg_entity_pc_pb.PC, index?: number): github_com_elojah_game_03_pkg_entity_pc_pb.PC;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListPCResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListPCResp): ListPCResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListPCResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListPCResp;
  static deserializeBinaryFromReader(message: ListPCResp, reader: jspb.BinaryReader): ListPCResp;
}

export namespace ListPCResp {
  export type AsObject = {
    pcsList: Array<github_com_elojah_game_03_pkg_entity_pc_pb.PC.AsObject>,
  }
}

