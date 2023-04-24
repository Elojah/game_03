// package: dto
// file: github.com/elojah/game_03/pkg/rtc/dto/rtc.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class SDP extends jspb.Message {
  getEncoded(): string;
  setEncoded(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): SDP.AsObject;
  static toObject(includeInstance: boolean, msg: SDP): SDP.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: SDP, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): SDP;
  static deserializeBinaryFromReader(message: SDP, reader: jspb.BinaryReader): SDP;
}

export namespace SDP {
  export type AsObject = {
    encoded: string,
  }
}

export class ICECandidate extends jspb.Message {
  getEncoded(): string;
  setEncoded(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ICECandidate.AsObject;
  static toObject(includeInstance: boolean, msg: ICECandidate): ICECandidate.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ICECandidate, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ICECandidate;
  static deserializeBinaryFromReader(message: ICECandidate, reader: jspb.BinaryReader): ICECandidate;
}

export namespace ICECandidate {
  export type AsObject = {
    encoded: string,
  }
}

