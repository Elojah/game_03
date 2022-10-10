// package: dto
// file: github.com/elojah/game_03/pkg/entity/dto/template.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class CreateTemplateReq extends jspb.Message {
  getName(): string;
  setName(value: string): void;

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
  }
}

