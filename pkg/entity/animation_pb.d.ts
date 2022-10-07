// package: entity
// file: github.com/elojah/game_03/pkg/entity/animation.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class Animation extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): void;

  getSheetid(): Uint8Array | string;
  getSheetid_asU8(): Uint8Array;
  getSheetid_asB64(): string;
  setSheetid(value: Uint8Array | string): void;

  getDuplicateid(): Uint8Array | string;
  getDuplicateid_asU8(): Uint8Array;
  getDuplicateid_asB64(): string;
  setDuplicateid(value: Uint8Array | string): void;

  getName(): string;
  setName(value: string): void;

  getStart(): number;
  setStart(value: number): void;

  getEnd(): number;
  setEnd(value: number): void;

  getRate(): number;
  setRate(value: number): void;

  getFramewidth(): number;
  setFramewidth(value: number): void;

  getFrameheight(): number;
  setFrameheight(value: number): void;

  getFramestart(): number;
  setFramestart(value: number): void;

  getFrameend(): number;
  setFrameend(value: number): void;

  getFramemargin(): number;
  setFramemargin(value: number): void;

  getFramespacing(): number;
  setFramespacing(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Animation.AsObject;
  static toObject(includeInstance: boolean, msg: Animation): Animation.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Animation, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Animation;
  static deserializeBinaryFromReader(message: Animation, reader: jspb.BinaryReader): Animation;
}

export namespace Animation {
  export type AsObject = {
    id: Uint8Array | string,
    entityid: Uint8Array | string,
    sheetid: Uint8Array | string,
    duplicateid: Uint8Array | string,
    name: string,
    start: number,
    end: number,
    rate: number,
    framewidth: number,
    frameheight: number,
    framestart: number,
    frameend: number,
    framemargin: number,
    framespacing: number,
  }
}

