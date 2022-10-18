// package: room
// file: github.com/elojah/game_03/pkg/room/world.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class World extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getName(): string;
  setName(value: string): void;

  getHeight(): number;
  setHeight(value: number): void;

  getWidth(): number;
  setWidth(value: number): void;

  getCellheight(): number;
  setCellheight(value: number): void;

  getCellwidth(): number;
  setCellwidth(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): World.AsObject;
  static toObject(includeInstance: boolean, msg: World): World.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: World, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): World;
  static deserializeBinaryFromReader(message: World, reader: jspb.BinaryReader): World;
}

export namespace World {
  export type AsObject = {
    id: Uint8Array | string,
    name: string,
    height: number,
    width: number,
    cellheight: number,
    cellwidth: number,
  }
}

