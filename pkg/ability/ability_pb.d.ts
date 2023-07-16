// package: ability
// file: github.com/elojah/game_03/pkg/ability/ability.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_entity_entity_pb from "../../../../../github.com/elojah/game_03/pkg/entity/entity_pb";

export class MoveTarget extends jspb.Message {
  getMove(): MoveMap[keyof MoveMap];
  setMove(value: MoveMap[keyof MoveMap]): void;

  getTargettype(): TargetTypeMap[keyof TargetTypeMap];
  setTargettype(value: TargetTypeMap[keyof TargetTypeMap]): void;

  getTargetid(): Uint8Array | string;
  getTargetid_asU8(): Uint8Array;
  getTargetid_asB64(): string;
  setTargetid(value: Uint8Array | string): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoveTarget.AsObject;
  static toObject(includeInstance: boolean, msg: MoveTarget): MoveTarget.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: MoveTarget, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoveTarget;
  static deserializeBinaryFromReader(message: MoveTarget, reader: jspb.BinaryReader): MoveTarget;
}

export namespace MoveTarget {
  export type AsObject = {
    move: MoveMap[keyof MoveMap],
    targettype: TargetTypeMap[keyof TargetTypeMap],
    targetid: Uint8Array | string,
  }
}

export class Target extends jspb.Message {
  getType(): TargetTypeMap[keyof TargetTypeMap];
  setType(value: TargetTypeMap[keyof TargetTypeMap]): void;

  getGroupid(): Uint8Array | string;
  getGroupid_asU8(): Uint8Array;
  getGroupid_asB64(): string;
  setGroupid(value: Uint8Array | string): void;

  getClosest(): boolean;
  setClosest(value: boolean): void;

  getRange(): number;
  setRange(value: number): void;

  getRadius(): number;
  setRadius(value: number): void;

  getWidth(): number;
  setWidth(value: number): void;

  getHeight(): number;
  setHeight(value: number): void;

  hasMove(): boolean;
  clearMove(): void;
  getMove(): MoveTarget | undefined;
  setMove(value?: MoveTarget): void;

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
    type: TargetTypeMap[keyof TargetTypeMap],
    groupid: Uint8Array | string,
    closest: boolean,
    range: number,
    radius: number,
    width: number,
    height: number,
    move?: MoveTarget.AsObject,
  }
}

export class AbilityModifier extends jspb.Message {
  getCancel(): boolean;
  setCancel(value: boolean): void;

  getCasttime(): number;
  setCasttime(value: number): void;

  getManacost(): number;
  setManacost(value: number): void;

  getCooldown(): number;
  setCooldown(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AbilityModifier.AsObject;
  static toObject(includeInstance: boolean, msg: AbilityModifier): AbilityModifier.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: AbilityModifier, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): AbilityModifier;
  static deserializeBinaryFromReader(message: AbilityModifier, reader: jspb.BinaryReader): AbilityModifier;
}

export namespace AbilityModifier {
  export type AsObject = {
    cancel: boolean,
    casttime: number,
    manacost: number,
    cooldown: number,
  }
}

export class EffectModifier extends jspb.Message {
  getEffectid(): Uint8Array | string;
  getEffectid_asU8(): Uint8Array;
  getEffectid_asB64(): string;
  setEffectid(value: Uint8Array | string): void;

  getCancel(): boolean;
  setCancel(value: boolean): void;

  getStat(): github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap];
  setStat(value: github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap]): void;

  hasAmount(): boolean;
  clearAmount(): void;
  getAmount(): Amount | undefined;
  setAmount(value?: Amount): void;

  getDuration(): number;
  setDuration(value: number): void;

  getDelay(): number;
  setDelay(value: number): void;

  getRepeat(): number;
  setRepeat(value: number): void;

  hasStackrules(): boolean;
  clearStackrules(): void;
  getStackrules(): StackRules | undefined;
  setStackrules(value?: StackRules): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EffectModifier.AsObject;
  static toObject(includeInstance: boolean, msg: EffectModifier): EffectModifier.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: EffectModifier, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EffectModifier;
  static deserializeBinaryFromReader(message: EffectModifier, reader: jspb.BinaryReader): EffectModifier;
}

export namespace EffectModifier {
  export type AsObject = {
    effectid: Uint8Array | string,
    cancel: boolean,
    stat: github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap],
    amount?: Amount.AsObject,
    duration: number,
    delay: number,
    repeat: number,
    stackrules?: StackRules.AsObject,
  }
}

export class Amount extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getDirect(): number;
  setDirect(value: number): void;

  hasTarget(): boolean;
  clearTarget(): void;
  getTarget(): Target | undefined;
  setTarget(value?: Target): void;

  getStat(): github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap];
  setStat(value: github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap]): void;

  getPercentage(): number;
  setPercentage(value: number): void;

  getEffectid(): Uint8Array | string;
  getEffectid_asU8(): Uint8Array;
  getEffectid_asB64(): string;
  setEffectid(value: Uint8Array | string): void;

  getStatoutcome(): boolean;
  setStatoutcome(value: boolean): void;

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
    id: Uint8Array | string,
    direct: number,
    target?: Target.AsObject,
    stat: github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap],
    percentage: number,
    effectid: Uint8Array | string,
    statoutcome: boolean,
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

  getAbilitymodifiersMap(): jspb.Map<string, AbilityModifier>;
  clearAbilitymodifiersMap(): void;
  getEffectmodifiersMap(): jspb.Map<string, EffectModifier>;
  clearEffectmodifiersMap(): void;
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
    abilitymodifiersMap: Array<[string, AbilityModifier.AsObject]>,
    effectmodifiersMap: Array<[string, EffectModifier.AsObject]>,
  }
}

export class StackRules extends jspb.Message {
  getStacks(): number;
  setStacks(value: number): void;

  getMaxstacks(): number;
  setMaxstacks(value: number): void;

  getMaxduration(): number;
  setMaxduration(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StackRules.AsObject;
  static toObject(includeInstance: boolean, msg: StackRules): StackRules.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: StackRules, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): StackRules;
  static deserializeBinaryFromReader(message: StackRules, reader: jspb.BinaryReader): StackRules;
}

export namespace StackRules {
  export type AsObject = {
    stacks: number,
    maxstacks: number,
    maxduration: number,
  }
}

export class Effect extends jspb.Message {
  getStat(): github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap];
  setStat(value: github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap]): void;

  hasAmount(): boolean;
  clearAmount(): void;
  getAmount(): Amount | undefined;
  setAmount(value?: Amount): void;

  getDuration(): number;
  setDuration(value: number): void;

  getIcon(): Uint8Array | string;
  getIcon_asU8(): Uint8Array;
  getIcon_asB64(): string;
  setIcon(value: Uint8Array | string): void;

  getDelay(): number;
  setDelay(value: number): void;

  getRepeat(): number;
  setRepeat(value: number): void;

  hasStackrules(): boolean;
  clearStackrules(): void;
  getStackrules(): StackRules | undefined;
  setStackrules(value?: StackRules): void;

  getTargetsMap(): jspb.Map<string, Target>;
  clearTargetsMap(): void;
  getTriggersMap(): jspb.Map<string, Trigger>;
  clearTriggersMap(): void;
  getEffectsMap(): jspb.Map<string, Effect>;
  clearEffectsMap(): void;
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
    stat: github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap],
    amount?: Amount.AsObject,
    duration: number,
    icon: Uint8Array | string,
    delay: number,
    repeat: number,
    stackrules?: StackRules.AsObject,
    targetsMap: Array<[string, Target.AsObject]>,
    triggersMap: Array<[string, Trigger.AsObject]>,
    effectsMap: Array<[string, Effect.AsObject]>,
  }
}

export class A extends jspb.Message {
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

  getEffectsMap(): jspb.Map<string, Effect>;
  clearEffectsMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): A.AsObject;
  static toObject(includeInstance: boolean, msg: A): A.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: A, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): A;
  static deserializeBinaryFromReader(message: A, reader: jspb.BinaryReader): A;
}

export namespace A {
  export type AsObject = {
    id: Uint8Array | string,
    name: string,
    icon: Uint8Array | string,
    animation: Uint8Array | string,
    casttime: number,
    manacost: number,
    cooldown: number,
    effectsMap: Array<[string, Effect.AsObject]>,
  }
}

export interface TargetTypeMap {
  NONETARGET: 0;
  SELF: 1;
  FOE: 2;
  ALLY: 3;
  RECT: 4;
  CIRCLE: 5;
}

export const TargetType: TargetTypeMap;

export interface MoveMap {
  NONEMOVE: 0;
  WALK: 1;
  TELEPORT: 2;
  PUSH: 3;
}

export const Move: MoveMap;

export interface OperatorMap {
  NONEOPERATOR: 0;
  EQUAL: 1;
  NOTEQUAL: 2;
  GREATER: 3;
  LESSER: 4;
}

export const Operator: OperatorMap;

