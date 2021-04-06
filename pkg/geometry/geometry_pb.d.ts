// package: geometry
// file: github.com/elojah/game_03/pkg/geometry/geometry.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class Vec2 extends jspb.Message {
  getX(): number;
  setX(value: number): void;

  getY(): number;
  setY(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Vec2.AsObject;
  static toObject(includeInstance: boolean, msg: Vec2): Vec2.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Vec2, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Vec2;
  static deserializeBinaryFromReader(message: Vec2, reader: jspb.BinaryReader): Vec2;
}

export namespace Vec2 {
  export type AsObject = {
    x: number,
    y: number,
  }
}

