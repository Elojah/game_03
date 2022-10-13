// package: user
// file: github.com/elojah/game_03/pkg/user/user.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class U extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getTwitchid(): string;
  setTwitchid(value: string): void;

  getGoogleid(): string;
  setGoogleid(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): U.AsObject;
  static toObject(includeInstance: boolean, msg: U): U.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: U, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): U;
  static deserializeBinaryFromReader(message: U, reader: jspb.BinaryReader): U;
}

export namespace U {
  export type AsObject = {
    id: Uint8Array | string,
    twitchid: string,
    googleid: string,
  }
}

