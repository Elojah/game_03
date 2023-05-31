// package: entity
// file: github.com/elojah/game_03/pkg/entity/entity.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_geometry_geometry_pb from "../../../../../github.com/elojah/game_03/pkg/geometry/geometry_pb";
import * as github_com_elojah_game_03_pkg_entity_animation_pb from "../../../../../github.com/elojah/game_03/pkg/entity/animation_pb";

export class Stats extends jspb.Message {
  getDamage(): number;
  setDamage(value: number): void;

  getDefense(): number;
  setDefense(value: number): void;

  getMovespeed(): number;
  setMovespeed(value: number): void;

  getCastspeed(): number;
  setCastspeed(value: number): void;

  getCooldownreduction(): number;
  setCooldownreduction(value: number): void;

  getHp(): number;
  setHp(value: number): void;

  getMp(): number;
  setMp(value: number): void;

  getMaxhp(): number;
  setMaxhp(value: number): void;

  getMaxmp(): number;
  setMaxmp(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Stats.AsObject;
  static toObject(includeInstance: boolean, msg: Stats): Stats.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Stats, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Stats;
  static deserializeBinaryFromReader(message: Stats, reader: jspb.BinaryReader): Stats;
}

export namespace Stats {
  export type AsObject = {
    damage: number,
    defense: number,
    movespeed: number,
    castspeed: number,
    cooldownreduction: number,
    hp: number,
    mp: number,
    maxhp: number,
    maxmp: number,
  }
}

export class E extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getUserid(): Uint8Array | string;
  getUserid_asU8(): Uint8Array;
  getUserid_asB64(): string;
  setUserid(value: Uint8Array | string): void;

  getCellid(): Uint8Array | string;
  getCellid_asU8(): Uint8Array;
  getCellid_asB64(): string;
  setCellid(value: Uint8Array | string): void;

  getFactionid(): Uint8Array | string;
  getFactionid_asU8(): Uint8Array;
  getFactionid_asB64(): string;
  setFactionid(value: Uint8Array | string): void;

  getName(): string;
  setName(value: string): void;

  getX(): number;
  setX(value: number): void;

  getY(): number;
  setY(value: number): void;

  getRot(): number;
  setRot(value: number): void;

  getRadius(): number;
  setRadius(value: number): void;

  getAt(): number;
  setAt(value: number): void;

  getAnimationid(): Uint8Array | string;
  getAnimationid_asU8(): Uint8Array;
  getAnimationid_asB64(): string;
  setAnimationid(value: Uint8Array | string): void;

  getAnimationat(): number;
  setAnimationat(value: number): void;

  clearObjectsList(): void;
  getObjectsList(): Array<github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect>;
  setObjectsList(value: Array<github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect>): void;
  addObjects(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect, index?: number): github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect;

  hasStats(): boolean;
  clearStats(): void;
  getStats(): Stats | undefined;
  setStats(value?: Stats): void;

  getEffectsMap(): jspb.Map<string, number>;
  clearEffectsMap(): void;
  getAbilitiesMap(): jspb.Map<string, github_com_elojah_game_03_pkg_entity_animation_pb.AnimationAbility>;
  clearAbilitiesMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): E.AsObject;
  static toObject(includeInstance: boolean, msg: E): E.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: E, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): E;
  static deserializeBinaryFromReader(message: E, reader: jspb.BinaryReader): E;
}

export namespace E {
  export type AsObject = {
    id: Uint8Array | string,
    userid: Uint8Array | string,
    cellid: Uint8Array | string,
    factionid: Uint8Array | string,
    name: string,
    x: number,
    y: number,
    rot: number,
    radius: number,
    at: number,
    animationid: Uint8Array | string,
    animationat: number,
    objectsList: Array<github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect.AsObject>,
    stats?: Stats.AsObject,
    effectsMap: Array<[string, number]>,
    abilitiesMap: Array<[string, github_com_elojah_game_03_pkg_entity_animation_pb.AnimationAbility.AsObject]>,
  }
}

export interface StatMap {
  NONESTAT: 0;
  DAMAGE: 1;
  DEFENSE: 2;
  MOVESPEED: 4;
  CASTSPEED: 5;
  COOLDOWNREDUCTION: 6;
  HP: 7;
  MP: 8;
  MAXHP: 9;
  MAXMP: 10;
}

export const Stat: StatMap;

