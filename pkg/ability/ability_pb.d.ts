// package: ability
// file: github.com/elojah/game_03/pkg/ability/ability.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_geometry_geometry_pb from "../../../../../github.com/elojah/game_03/pkg/geometry/geometry_pb";
import * as github_com_elojah_game_03_pkg_entity_entity_pb from "../../../../../github.com/elojah/game_03/pkg/entity/entity_pb";

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
    drain?: Amount.AsObject,
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

  getTarget(): TargetMap[keyof TargetMap];
  setTarget(value: TargetMap[keyof TargetMap]): void;

  getStat(): github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap];
  setStat(value: github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap]): void;

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
    id: Uint8Array | string,
    direct: number,
    target: TargetMap[keyof TargetMap],
    stat: github_com_elojah_game_03_pkg_entity_entity_pb.StatMap[keyof github_com_elojah_game_03_pkg_entity_entity_pb.StatMap],
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

  getDuration(): number;
  setDuration(value: number): void;

  getIcon(): Uint8Array | string;
  getIcon_asU8(): Uint8Array;
  getIcon_asB64(): string;
  setIcon(value: Uint8Array | string): void;

  hasStackrules(): boolean;
  clearStackrules(): void;
  getStackrules(): StackRules | undefined;
  setStackrules(value?: StackRules): void;

  getDelay(): number;
  setDelay(value: number): void;

  getRepeat(): number;
  setRepeat(value: number): void;

  getTargetsMap(): jspb.Map<string, TargetMap[keyof TargetMap]>;
  clearTargetsMap(): void;
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
    position?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.AsObject,
    force?: github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.AsObject,
    amount?: Amount.AsObject,
    duration: number,
    icon: Uint8Array | string,
    stackrules?: StackRules.AsObject,
    delay: number,
    repeat: number,
    targetsMap: Array<[string, TargetMap[keyof TargetMap]]>,
  }
}

export class Outcome extends jspb.Message {
  hasStats(): boolean;
  clearStats(): void;
  getStats(): github_com_elojah_game_03_pkg_entity_entity_pb.Stats | undefined;
  setStats(value?: github_com_elojah_game_03_pkg_entity_entity_pb.Stats): void;

  hasStatsmodified(): boolean;
  clearStatsmodified(): void;
  getStatsmodified(): github_com_elojah_game_03_pkg_entity_entity_pb.Stats | undefined;
  setStatsmodified(value?: github_com_elojah_game_03_pkg_entity_entity_pb.Stats): void;

  hasStacks(): boolean;
  clearStacks(): void;
  getStacks(): StackRules | undefined;
  setStacks(value?: StackRules): void;

  hasStacksmodified(): boolean;
  clearStacksmodified(): void;
  getStacksmodified(): StackRules | undefined;
  setStacksmodified(value?: StackRules): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Outcome.AsObject;
  static toObject(includeInstance: boolean, msg: Outcome): Outcome.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: Outcome, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Outcome;
  static deserializeBinaryFromReader(message: Outcome, reader: jspb.BinaryReader): Outcome;
}

export namespace Outcome {
  export type AsObject = {
    stats?: github_com_elojah_game_03_pkg_entity_entity_pb.Stats.AsObject,
    statsmodified?: github_com_elojah_game_03_pkg_entity_entity_pb.Stats.AsObject,
    stacks?: StackRules.AsObject,
    stacksmodified?: StackRules.AsObject,
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

  getEffectsMap(): jspb.Map<string, Effect>;
  clearEffectsMap(): void;
  getTriggersMap(): jspb.Map<string, Trigger>;
  clearTriggersMap(): void;
  getTriggersoutcomeMap(): jspb.Map<string, Trigger>;
  clearTriggersoutcomeMap(): void;
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
    effectsMap: Array<[string, Effect.AsObject]>,
    triggersMap: Array<[string, Trigger.AsObject]>,
    triggersoutcomeMap: Array<[string, Trigger.AsObject]>,
  }
}

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

