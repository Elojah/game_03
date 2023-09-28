import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class Animation extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): Animation;

  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): Animation;

  getSheetid(): Uint8Array | string;
  getSheetid_asU8(): Uint8Array;
  getSheetid_asB64(): string;
  setSheetid(value: Uint8Array | string): Animation;

  getDuplicateid(): Uint8Array | string;
  getDuplicateid_asU8(): Uint8Array;
  getDuplicateid_asB64(): string;
  setDuplicateid(value: Uint8Array | string): Animation;

  getName(): string;
  setName(value: string): Animation;

  getStart(): number;
  setStart(value: number): Animation;

  getEnd(): number;
  setEnd(value: number): Animation;

  getSequenceList(): Array<number>;
  setSequenceList(value: Array<number>): Animation;
  clearSequenceList(): Animation;
  addSequence(value: number, index?: number): Animation;

  getRate(): number;
  setRate(value: number): Animation;

  getFramewidth(): number;
  setFramewidth(value: number): Animation;

  getFrameheight(): number;
  setFrameheight(value: number): Animation;

  getFramestart(): number;
  setFramestart(value: number): Animation;

  getFrameend(): number;
  setFrameend(value: number): Animation;

  getFramemargin(): number;
  setFramemargin(value: number): Animation;

  getFramespacing(): number;
  setFramespacing(value: number): Animation;

  getRepeat(): number;
  setRepeat(value: number): Animation;

  getDelay(): number;
  setDelay(value: number): Animation;

  getDuration(): number;
  setDuration(value: number): Animation;

  getShowandhide(): boolean;
  setShowandhide(value: boolean): Animation;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Animation.AsObject;
  static toObject(includeInstance: boolean, msg: Animation): Animation.AsObject;
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
    repeat: number,
    delay: number,
    duration: number,
    showandhide: boolean,
  }
}

export class AnimationAbility extends jspb.Message {
  getAnimationid(): Uint8Array | string;
  getAnimationid_asU8(): Uint8Array;
  getAnimationid_asB64(): string;
  setAnimationid(value: Uint8Array | string): AnimationAbility;

  getCellid(): Uint8Array | string;
  getCellid_asU8(): Uint8Array;
  getCellid_asB64(): string;
  setCellid(value: Uint8Array | string): AnimationAbility;

  getX(): number;
  setX(value: number): AnimationAbility;

  getY(): number;
  setY(value: number): AnimationAbility;

  getRot(): number;
  setRot(value: number): AnimationAbility;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AnimationAbility.AsObject;
  static toObject(includeInstance: boolean, msg: AnimationAbility): AnimationAbility.AsObject;
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

