import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';
import * as github_com_elojah_game_03_pkg_entity_animation_pb from '../../../../../../github.com/elojah/game_03/pkg/entity/animation_pb';


export class ListAnimationReq extends jspb.Message {
  getIdsList(): Array<Uint8Array | string>;
  setIdsList(value: Array<Uint8Array | string>): ListAnimationReq;
  clearIdsList(): ListAnimationReq;
  addIds(value: Uint8Array | string, index?: number): ListAnimationReq;

  getEntityidsList(): Array<Uint8Array | string>;
  setEntityidsList(value: Array<Uint8Array | string>): ListAnimationReq;
  clearEntityidsList(): ListAnimationReq;
  addEntityids(value: Uint8Array | string, index?: number): ListAnimationReq;

  getSize(): number;
  setSize(value: number): ListAnimationReq;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): ListAnimationReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAnimationReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListAnimationReq): ListAnimationReq.AsObject;
  static serializeBinaryToWriter(message: ListAnimationReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAnimationReq;
  static deserializeBinaryFromReader(message: ListAnimationReq, reader: jspb.BinaryReader): ListAnimationReq;
}

export namespace ListAnimationReq {
  export type AsObject = {
    idsList: Array<Uint8Array | string>,
    entityidsList: Array<Uint8Array | string>,
    size: number,
    state: Uint8Array | string,
  }
}

export class ListAnimationResp extends jspb.Message {
  getAnimationsList(): Array<github_com_elojah_game_03_pkg_entity_animation_pb.Animation>;
  setAnimationsList(value: Array<github_com_elojah_game_03_pkg_entity_animation_pb.Animation>): ListAnimationResp;
  clearAnimationsList(): ListAnimationResp;
  addAnimations(value?: github_com_elojah_game_03_pkg_entity_animation_pb.Animation, index?: number): github_com_elojah_game_03_pkg_entity_animation_pb.Animation;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): ListAnimationResp;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAnimationResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListAnimationResp): ListAnimationResp.AsObject;
  static serializeBinaryToWriter(message: ListAnimationResp, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): ListAnimationResp;
  static deserializeBinaryFromReader(message: ListAnimationResp, reader: jspb.BinaryReader): ListAnimationResp;
}

export namespace ListAnimationResp {
  export type AsObject = {
    animationsList: Array<github_com_elojah_game_03_pkg_entity_animation_pb.Animation.AsObject>,
    state: Uint8Array | string,
  }
}

export class CreateAnimationReq extends jspb.Message {
  getEntitytemplate(): string;
  setEntitytemplate(value: string): CreateAnimationReq;

  getAnimation(): github_com_elojah_game_03_pkg_entity_animation_pb.Animation | undefined;
  setAnimation(value?: github_com_elojah_game_03_pkg_entity_animation_pb.Animation): CreateAnimationReq;
  hasAnimation(): boolean;
  clearAnimation(): CreateAnimationReq;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAnimationReq.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAnimationReq): CreateAnimationReq.AsObject;
  static serializeBinaryToWriter(message: CreateAnimationReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAnimationReq;
  static deserializeBinaryFromReader(message: CreateAnimationReq, reader: jspb.BinaryReader): CreateAnimationReq;
}

export namespace CreateAnimationReq {
  export type AsObject = {
    entitytemplate: string,
    animation?: github_com_elojah_game_03_pkg_entity_animation_pb.Animation.AsObject,
  }
}

