// package: twitch
// file: github.com/elojah/game_03/pkg/twitch/token.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class Token extends jspb.Message {
  getAccesstoken(): string;
  setAccesstoken(value: string): void;

  getRefreshtoken(): string;
  setRefreshtoken(value: string): void;

  getExpiresin(): number;
  setExpiresin(value: number): void;

  clearScopeList(): void;
  getScopeList(): Array<string>;
  setScopeList(value: Array<string>): void;
  addScope(value: string, index?: number): string;

  getTokentype(): string;
  setTokentype(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Token.AsObject;
  static toObject(includeInstance: boolean, msg: Token): Token.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Token, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Token;
  static deserializeBinaryFromReader(message: Token, reader: jspb.BinaryReader): Token;
}

export namespace Token {
  export type AsObject = {
    accesstoken: string,
    refreshtoken: string,
    expiresin: number,
    scopeList: Array<string>,
    tokentype: string,
  }
}

