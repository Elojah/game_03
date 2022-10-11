// package: dto
// file: github.com/elojah/game_03/pkg/entity/dto/animation.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_entity_animation_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/animation_pb";

export class ListAnimationReq extends jspb.Message {
  clearIdsList(): void;
  getIdsList(): Array<Uint8Array | string>;
  getIdsList_asU8(): Array<Uint8Array>;
  getIdsList_asB64(): Array<string>;
  setIdsList(value: Array<Uint8Array | string>): void;
  addIds(value: Uint8Array | string, index?: number): Uint8Array | string;

  clearEntityidsList(): void;
  getEntityidsList(): Array<Uint8Array | string>;
  getEntityidsList_asU8(): Array<Uint8Array>;
  getEntityidsList_asB64(): Array<string>;
  setEntityidsList(value: Array<Uint8Array | string>): void;
  addEntityids(value: Uint8Array | string, index?: number): Uint8Array | string;

  getSize(): number;
  setSize(value: number): void;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAnimationReq.AsObject;
  static toObject(includeInstance: boolean, msg: ListAnimationReq): ListAnimationReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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
  clearAnimationsList(): void;
  getAnimationsList(): Array<github_com_elojah_game_03_pkg_entity_animation_pb.Animation>;
  setAnimationsList(value: Array<github_com_elojah_game_03_pkg_entity_animation_pb.Animation>): void;
  addAnimations(value?: github_com_elojah_game_03_pkg_entity_animation_pb.Animation, index?: number): github_com_elojah_game_03_pkg_entity_animation_pb.Animation;

  getState(): Uint8Array | string;
  getState_asU8(): Uint8Array;
  getState_asB64(): string;
  setState(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): ListAnimationResp.AsObject;
  static toObject(includeInstance: boolean, msg: ListAnimationResp): ListAnimationResp.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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
  hasAnimation(): boolean;
  clearAnimation(): void;
  getAnimation(): github_com_elojah_game_03_pkg_entity_animation_pb.Animation | undefined;
  setAnimation(value?: github_com_elojah_game_03_pkg_entity_animation_pb.Animation): void;

  getEntitytemplate(): string;
  setEntitytemplate(value: string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): CreateAnimationReq.AsObject;
  static toObject(includeInstance: boolean, msg: CreateAnimationReq): CreateAnimationReq.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: CreateAnimationReq, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): CreateAnimationReq;
  static deserializeBinaryFromReader(message: CreateAnimationReq, reader: jspb.BinaryReader): CreateAnimationReq;
}

export namespace CreateAnimationReq {
  export type AsObject = {
    animation?: github_com_elojah_game_03_pkg_entity_animation_pb.Animation.AsObject,
    entitytemplate: string,
  }
}

