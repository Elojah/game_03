// package: dto
// file: github.com/elojah/game_03/pkg/twitch/dto/follow.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_twitch_follow_pb from "../../../../../../github.com/elojah/game_03/pkg/twitch/follow_pb";

export class ListFollowReq extends jspb.Message {
  getAfter(): string;
  setAfter(value: string): void;

  getFirst(): string;
  setFirst(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFollowReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListFollowReq): ListFollowReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListFollowReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFollowReq;
  static deserializeBinaryFromReader(message: ListFollowReq, reader: jspb.BinaryReader): ListFollowReq;
}

export namespace ListFollowReq {
  export type AsObject = {
    after: string,
    first: string,
  }
}

export class ListFollowResp extends jspb.Message {
  clearFollowsList(): void;
  getFollowsList(): Array<github_com_elojah_game_03_pkg_twitch_follow_pb.Follow>;
  setFollowsList(value: Array<github_com_elojah_game_03_pkg_twitch_follow_pb.Follow>): void;
  addFollows(value?: github_com_elojah_game_03_pkg_twitch_follow_pb.Follow, index?: number): github_com_elojah_game_03_pkg_twitch_follow_pb.Follow;

  getTotal(): number;
  setTotal(value: number): void;

  getCursor(): string;
  setCursor(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListFollowResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListFollowResp): ListFollowResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: ListFollowResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListFollowResp;
  static deserializeBinaryFromReader(message: ListFollowResp, reader: jspb.BinaryReader): ListFollowResp;
}

export namespace ListFollowResp {
  export type AsObject = {
    followsList: Array<github_com_elojah_game_03_pkg_twitch_follow_pb.Follow.AsObject>,
    total: number,
    cursor: string,
  }
}

