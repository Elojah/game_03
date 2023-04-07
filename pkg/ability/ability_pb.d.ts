// package: ability
// file: github.com/elojah/game_03/pkg/ability/ability.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_geometry_geometry_pb from "../../../../../github.com/elojah/game_03/pkg/geometry/geometry_pb";

export class Modifier extends jspb.Message {
  getCasttime(): number;
  setCasttime(value: number): void;

  getManacost(): number;
  setManacost(value: number): void;

  getCooldown(): number;
  setCooldown(value: number): void;

  getStat(): StatMap[keyof StatMap];
  setStat(value: StatMap[keyof StatMap]): void;

  hasAmount(): boolean;
  clearAmount(): void;
  getAmount(): Amount | undefined;
  setAmount(value?: Amount): void;

  hasDrain(): boolean;
  clearDrain(): void;
  getDrain(): Amount | undefined;
  setDrain(value?: Amount): void;

  getDuration(): number;
  setDuration(value: number): void;

  getDelay(): number;
  setDelay(value: number): void;

  getRepeat(): number;
  setRepeat(value: number): void;

  getStacks(): number;
  setStacks(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Modifier.AsObject;
  static toObject(includeInstance: boolean, msg: Modifier): Modifier.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Modifier, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Modifier;
  static deserializeBinaryFromReader(message: Modifier, reader: jspb.BinaryReader): Modifier;
}

export namespace Modifier {
  export type AsObject = {
    casttime: number,
    manacost: number,
    cooldown: number,
    stat: StatMap[keyof StatMap],
    amount?: Amount.AsObject,
    drain?: Amount.AsObject,
    duration: number,
    delay: number,
    repeat: number,
    stacks: number,
  }
}

export class Amount extends jspb.Message {
  getTarget(): TargetMap[keyof TargetMap];
  setTarget(value: TargetMap[keyof TargetMap]): void;

  getDirect(): number;
  setDirect(value: number): void;

  getStat(): number;
  setStat(value: number): void;

  getPercentage(): number;
  setPercentage(value: number): void;

  getEffectid(): Uint8Array | string;
  getEffectid_asU8(): Uint8Array;
  getEffectid_asB64(): string;
  setEffectid(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Amount.AsObject;
  static toObject(includeInstance: boolean, msg: Amount): Amount.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Amount, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Amount;
  static deserializeBinaryFromReader(message: Amount, reader: jspb.BinaryReader): Amount;
}

export namespace Amount {
  export type AsObject = {
    target: TargetMap[keyof TargetMap],
    direct: number,
    stat: number,
    percentage: number,
    effectid: Uint8Array | string,
  }
}

export class Trigger extends jspb.Message {
  getOperator(): OperatorMap[keyof OperatorMap];
  setOperator(value: OperatorMap[keyof OperatorMap]): void;

  hasAmount(): boolean;
  clearAmount(): void;
  getAmount(): Amount | undefined;
  setAmount(value?: Amount): void;

  hasTreshold(): boolean;
  clearTreshold(): void;
  getTreshold(): Amount | undefined;
  setTreshold(value?: Amount): void;

  getConsumetreshold(): boolean;
  setConsumetreshold(value: boolean): void;

  clearModifiersList(): void;
  getModifiersList(): Array<Modifier>;
  setModifiersList(value: Array<Modifier>): void;
  addModifiers(value?: Modifier, index?: number): Modifier;

  hasTrigger(): boolean;
  clearTrigger(): void;
  getTrigger(): Trigger | undefined;
  setTrigger(value?: Trigger): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Trigger.AsObject;
  static toObject(includeInstance: boolean, msg: Trigger): Trigger.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Trigger, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Trigger;
  static deserializeBinaryFromReader(message: Trigger, reader: jspb.BinaryReader): Trigger;
}

export namespace Trigger {
  export type AsObject = {
    operator: OperatorMap[keyof OperatorMap],
    amount?: Amount.AsObject,
    treshold?: Amount.AsObject,
    consumetreshold: boolean,
    modifiersList: Array<Modifier.AsObject>,
    trigger?: Trigger.AsObject,
  }
}

export class Effect extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getStat(): StatMap[keyof StatMap];
  setStat(value: StatMap[keyof StatMap]): void;

  hasPosition(): boolean;
  clearPosition(): void;
  getPosition(): github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2 | undefined;
  setPosition(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2): void;

  hasForce(): boolean;
  clearForce(): void;
  getForce(): github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2 | undefined;
  setForce(value?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2): void;

  hasAmount(): boolean;
  clearAmount(): void;
  getAmount(): Amount | undefined;
  setAmount(value?: Amount): void;

  hasDrain(): boolean;
  clearDrain(): void;
  getDrain(): Amount | undefined;
  setDrain(value?: Amount): void;

  clearDraintargetsList(): void;
  getDraintargetsList(): Array<TargetMap[keyof TargetMap]>;
  setDraintargetsList(value: Array<TargetMap[keyof TargetMap]>): void;
  addDraintargets(value: TargetMap[keyof TargetMap], index?: number): TargetMap[keyof TargetMap];

  getDuration(): number;
  setDuration(value: number): void;

  getIcon(): Uint8Array | string;
  getIcon_asU8(): Uint8Array;
  getIcon_asB64(): string;
  setIcon(value: Uint8Array | string): void;

  getStacks(): number;
  setStacks(value: number): void;

  getMaxstack(): number;
  setMaxstack(value: number): void;

  getDelay(): number;
  setDelay(value: number): void;

  getRepeat(): number;
  setRepeat(value: number): void;

  clearTriggersList(): void;
  getTriggersList(): Array<Trigger>;
  setTriggersList(value: Array<Trigger>): void;
  addTriggers(value?: Trigger, index?: number): Trigger;

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
    id: Uint8Array | string,
    stat: StatMap[keyof StatMap],
    position?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.AsObject,
    force?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.AsObject,
    amount?: Amount.AsObject,
    drain?: Amount.AsObject,
    draintargetsList: Array<TargetMap[keyof TargetMap]>,
    duration: number,
    icon: Uint8Array | string,
    stacks: number,
    maxstack: number,
    delay: number,
    repeat: number,
    triggersList: Array<Trigger.AsObject>,
  }
}

export class Component extends jspb.Message {
  clearTargetsList(): void;
  getTargetsList(): Array<TargetMap[keyof TargetMap]>;
  setTargetsList(value: Array<TargetMap[keyof TargetMap]>): void;
  addTargets(value: TargetMap[keyof TargetMap], index?: number): TargetMap[keyof TargetMap];

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
    targetsList: Array<TargetMap[keyof TargetMap]>,
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

  getCasttime(): number;
  setCasttime(value: number): void;

  getManacost(): number;
  setManacost(value: number): void;

  getCooldown(): number;
  setCooldown(value: number): void;

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
    casttime: number,
    manacost: number,
    cooldown: number,
    componentsList: Array<Component.AsObject>,
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
  DAMAGERECEIVED: 11;
}

export const Stat: StatMap;

export interface TargetMap {
  NONETARGET: 0;
  SELF: 1;
  FOE: 3;
  CLOSESTSELF: 4;
  CLOSESTFOE: 5;
  RECT: 6;
  CIRCLE: 7;
}

export const Target: TargetMap;

export interface OperatorMap {
  NONEOPERATOR: 0;
  EQUAL: 1;
  NOTEQUAL: 2;
  GREATER: 3;
  LESSER: 4;
}

export const Operator: OperatorMap;

