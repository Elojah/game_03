// package: dto
// file: github.com/elojah/game_03/pkg/entity/dto/template.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_entity_template_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/template_pb";
import * as github_com_elojah_game_03_pkg_entity_entity_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/entity_pb";

export class CreateTemplateReq extends jspb.Message {
  getName(): string;
  setName(value: string): void;

  hasEntity(): boolean;
  clearEntity(): void;
  getEntity(): github_com_elojah_game_03_pkg_entity_entity_pb.E | undefined;
  setEntity(value?: github_com_elojah_game_03_pkg_entity_entity_pb.E): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateTemplateReq.AsObject;
  static toObject(includeInstance: boolean, msg: CreateTemplateReq): CreateTemplateReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateTemplateReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateTemplateReq;
  static deserializeBinaryFromReader(message: CreateTemplateReq, reader: jspb.BinaryReader): CreateTemplateReq;
}

export namespace CreateTemplateReq {
  export type AsObject = {
    name: string,
    entity?: github_com_elojah_game_03_pkg_entity_entity_pb.E.AsObject,
  }
}

export class ListTemplateReq extends jspb.Message {
  getAll(): boolean;
  setAll(value: boolean): void;

  getSize(): number;
  setSize(value: number): void;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTemplateReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListTemplateReq): ListTemplateReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListTemplateReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTemplateReq;
  static deserializeBinaryFromReader(message: ListTemplateReq, reader: jspb.BinaryReader): ListTemplateReq;
}

export namespace ListTemplateReq {
  export type AsObject = {
    all: boolean,
    size: number,
    state: Uint8Array | string,
  }
}

export class ListTemplateResp extends jspb.Message {
  clearTemplatesList(): void;
  getTemplatesList(): Array<github_com_elojah_game_03_pkg_entity_template_pb.Template>;
  setTemplatesList(value: Array<github_com_elojah_game_03_pkg_entity_template_pb.Template>): void;
  addTemplates(value?: github_com_elojah_game_03_pkg_entity_template_pb.Template, index?: number): github_com_elojah_game_03_pkg_entity_template_pb.Template;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListTemplateResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListTemplateResp): ListTemplateResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListTemplateResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListTemplateResp;
  static deserializeBinaryFromReader(message: ListTemplateResp, reader: jspb.BinaryReader): ListTemplateResp;
}

export namespace ListTemplateResp {
  export type AsObject = {
    templatesList: Array<github_com_elojah_game_03_pkg_entity_template_pb.Template.AsObject>,
    state: Uint8Array | string,
  }
}

