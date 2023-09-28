import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';
import * as github_com_elojah_game_03_pkg_geometry_geometry_pb from '../../../../../github.com/elojah/game_03/pkg/geometry/geometry_pb';
import * as github_com_elojah_game_03_pkg_entity_animation_pb from '../../../../../github.com/elojah/game_03/pkg/entity/animation_pb';


export class Stats extends jspb.Message {
  getDamage(): number;
  setDamage(value: number): Stats;

  getDefense(): number;
  setDefense(value: number): Stats;

  getMovespeed(): number;
  setMovespeed(value: number): Stats;

  getCastspeed(): number;
  setCastspeed(value: number): Stats;

  getCooldownreduction(): number;
  setCooldownreduction(value: number): Stats;

  getHp(): number;
  setHp(value: number): Stats;

  getMp(): number;
  setMp(value: number): Stats;

  getMaxhp(): number;
  setMaxhp(value: number): Stats;

  getMaxmp(): number;
  setMaxmp(value: number): Stats;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Stats.AsObject;
  static toObject(includeInstance: boolean, msg: Stats): Stats.AsObject;
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
  setId(value: Uint8Array | string): E;

  getUserid(): Uint8Array | string;
  getUserid_asU8(): Uint8Array;
  getUserid_asB64(): string;
  setUserid(value: Uint8Array | string): E;

  getCellid(): Uint8Array | string;
  getCellid_asU8(): Uint8Array;
  getCellid_asB64(): string;
  setCellid(value: Uint8Array | string): E;

  getFactionid(): Uint8Array | string;
  getFactionid_asU8(): Uint8Array;
  getFactionid_asB64(): string;
  setFactionid(value: Uint8Array | string): E;

  getName(): string;
  setName(value: string): E;

  getX(): number;
  setX(value: number): E;

  getY(): number;
  setY(value: number): E;

  getRot(): number;
  setRot(value: number): E;

  getRadius(): number;
  setRadius(value: number): E;

  getAt(): number;
  setAt(value: number): E;

  getAnimationid(): Uint8Array | string;
  getAnimationid_asU8(): Uint8Array;
  getAnimationid_asB64(): string;
  setAnimationid(value: Uint8Array | string): E;

  getAnimationat(): number;
  setAnimationat(value: number): E;

  getObjectsList(): Array<github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect>;
  setObjectsList(value: Array<github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect>): E;
  clearObjectsList(): E;
  addObjects(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect, index?: number): github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect;

  getStats(): Stats | undefined;
  setStats(value?: Stats): E;
  hasStats(): boolean;
  clearStats(): E;

  getEffectsMap(): jspb.Map<string, number>;
  clearEffectsMap(): E;

  getAbilitiesMap(): jspb.Map<string, github_com_elojah_game_03_pkg_entity_animation_pb.AnimationAbility>;
  clearAbilitiesMap(): E;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): E.AsObject;
  static toObject(includeInstance: boolean, msg: E): E.AsObject;
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

export enum Stat { 
  NONESTAT = 0,
  DAMAGE = 1,
  DEFENSE = 2,
  MOVESPEED = 4,
  CASTSPEED = 5,
  COOLDOWNREDUCTION = 6,
  HP = 7,
  MP = 8,
  MAXHP = 9,
  MAXMP = 10,
}
