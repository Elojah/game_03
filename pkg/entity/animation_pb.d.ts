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

  clearSequenceList(): void;
  getSequenceList(): Array<number>;
  setSequenceList(value: Array<number>): void;
  addSequence(value: number, index?: number): number;

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
    sequenceList: Array<number>,
    rate: number,
    framewidth: number,
    frameheight: number,
    framestart: number,
    frameend: number,
    framemargin: number,
    framespacing: number,
  }
}

export class AnimationAbility extends jspb.Message {
  getAnimationid(): Uint8Array | string;
  getAnimationid_asU8(): Uint8Array;
  getAnimationid_asB64(): string;
  setAnimationid(value: Uint8Array | string): void;

  getCellid(): Uint8Array | string;
  getCellid_asU8(): Uint8Array;
  getCellid_asB64(): string;
  setCellid(value: Uint8Array | string): void;

  getX(): number;
  setX(value: number): void;

  getY(): number;
  setY(value: number): void;

  getRot(): number;
  setRot(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AnimationAbility.AsObject;
  static toObject(includeInstance: boolean, msg: AnimationAbility): AnimationAbility.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AnimationAbility, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AnimationAbility;
  static deserializeBinaryFromReader(message: AnimationAbility, reader: jspb.BinaryReader): AnimationAbility;
}

export namespace AnimationAbility {
  export type AsObject = {
    animationid: Uint8Array | string,
    cellid: Uint8Array | string,
    x: number,
    y: number,
    rot: number,
  }
}

