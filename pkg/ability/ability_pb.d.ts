// package: ability
// file: github.com/elojah/game_03/pkg/ability/ability.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_geometry_geometry_pb from "../../../../../github.com/elojah/game_03/pkg/geometry/geometry_pb";

export class Target extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  hasRect(): boolean;
  clearRect(): void;
  getRect(): github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect | undefined;
  setRect(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect): void;

  hasCircle(): boolean;
  clearCircle(): void;
  getCircle(): github_com_elojah_game_03_pkg_geometry_geometry_pb.Circle | undefined;
  setCircle(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Circle): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Target.AsObject;
  static toObject(includeInstance: boolean, msg: Target): Target.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Target, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Target;
  static deserializeBinaryFromReader(message: Target, reader: jspb.BinaryReader): Target;
}

export namespace Target {
  export type AsObject = {
    id: Uint8Array | string,
    rect?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect.AsObject,
    circle?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Circle.AsObject,
  }
}

export class Effect extends jspb.Message {
  getStat(): string;
  setStat(value: string): void;

  getValue(): number;
  setValue(value: number): void;

  getElement(): ElementMap[keyof ElementMap];
  setElement(value: ElementMap[keyof ElementMap]): void;

  getTickrate(): number;
  setTickrate(value: number): void;

  getDuration(): number;
  setDuration(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Effect.AsObject;
  static toObject(includeInstance: boolean, msg: Effect): Effect.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Effect, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Effect;
  static deserializeBinaryFromReader(message: Effect, reader: jspb.BinaryReader): Effect;
}

export namespace Effect {
  export type AsObject = {
    stat: string,
    value: number,
    element: ElementMap[keyof ElementMap],
    tickrate: number,
    duration: number,
  }
}

export class Component extends jspb.Message {
  getCasttime(): number;
  setCasttime(value: number): void;

  getManacost(): number;
  setManacost(value: number): void;

  getCooldown(): number;
  setCooldown(value: number): void;

  clearTargetsList(): void;
  getTargetsList(): Array<Target>;
  setTargetsList(value: Array<Target>): void;
  addTargets(value?: Target, index?: number): Target;

  clearEffectsList(): void;
  getEffectsList(): Array<Effect>;
  setEffectsList(value: Array<Effect>): void;
  addEffects(value?: Effect, index?: number): Effect;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Component.AsObject;
  static toObject(includeInstance: boolean, msg: Component): Component.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Component, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Component;
  static deserializeBinaryFromReader(message: Component, reader: jspb.BinaryReader): Component;
}

export namespace Component {
  export type AsObject = {
    casttime: number,
    manacost: number,
    cooldown: number,
    targetsList: Array<Target.AsObject>,
    effectsList: Array<Effect.AsObject>,
  }
}

export class Ability extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getName(): string;
  setName(value: string): void;

  getIcon(): Uint8Array | string;
  getIcon_asU8(): Uint8Array;
  getIcon_asB64(): string;
  setIcon(value: Uint8Array | string): void;

  getAnimation(): Uint8Array | string;
  getAnimation_asU8(): Uint8Array;
  getAnimation_asB64(): string;
  setAnimation(value: Uint8Array | string): void;

  clearComponentsList(): void;
  getComponentsList(): Array<Component>;
  setComponentsList(value: Array<Component>): void;
  addComponents(value?: Component, index?: number): Component;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Ability.AsObject;
  static toObject(includeInstance: boolean, msg: Ability): Ability.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Ability, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Ability;
  static deserializeBinaryFromReader(message: Ability, reader: jspb.BinaryReader): Ability;
}

export namespace Ability {
  export type AsObject = {
    id: Uint8Array | string,
    name: string,
    icon: Uint8Array | string,
    animation: Uint8Array | string,
    componentsList: Array<Component.AsObject>,
  }
}

export interface ElementMap {
  FIRE: 0;
  WATER: 1;
  AIR: 2;
  EARTH: 3;
  LIGHT: 4;
  DARK: 5;
  TIME: 6;
}

export const Element: ElementMap;

