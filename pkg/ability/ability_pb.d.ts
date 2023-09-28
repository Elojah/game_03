import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';
import * as github_com_elojah_game_03_pkg_entity_entity_pb from '../../../../../github.com/elojah/game_03/pkg/entity/entity_pb';


export class MoveTarget extends jspb.Message {
  getMove(): Move;
  setMove(value: Move): MoveTarget;

  getTargettype(): TargetType;
  setTargettype(value: TargetType): MoveTarget;

  getTargetid(): Uint8Array | string;
  getTargetid_asU8(): Uint8Array;
  getTargetid_asB64(): string;
  setTargetid(value: Uint8Array | string): MoveTarget;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): MoveTarget.AsObject;
  static toObject(includeInstance: boolean, msg: MoveTarget): MoveTarget.AsObject;
  static serializeBinaryToWriter(message: MoveTarget, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): MoveTarget;
  static deserializeBinaryFromReader(message: MoveTarget, reader: jspb.BinaryReader): MoveTarget;
}

export namespace MoveTarget {
  export type AsObject = {
    move: Move,
    targettype: TargetType,
    targetid: Uint8Array | string,
  }
}

export class Target extends jspb.Message {
  getType(): TargetType;
  setType(value: TargetType): Target;

  getGroupid(): Uint8Array | string;
  getGroupid_asU8(): Uint8Array;
  getGroupid_asB64(): string;
  setGroupid(value: Uint8Array | string): Target;

  getClosest(): boolean;
  setClosest(value: boolean): Target;

  getRange(): number;
  setRange(value: number): Target;

  getRadius(): number;
  setRadius(value: number): Target;

  getWidth(): number;
  setWidth(value: number): Target;

  getHeight(): number;
  setHeight(value: number): Target;

  getMovetarget(): MoveTarget | undefined;
  setMovetarget(value?: MoveTarget): Target;
  hasMovetarget(): boolean;
  clearMovetarget(): Target;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Target.AsObject;
  static toObject(includeInstance: boolean, msg: Target): Target.AsObject;
  static serializeBinaryToWriter(message: Target, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Target;
  static deserializeBinaryFromReader(message: Target, reader: jspb.BinaryReader): Target;
}

export namespace Target {
  export type AsObject = {
    type: TargetType,
    groupid: Uint8Array | string,
    closest: boolean,
    range: number,
    radius: number,
    width: number,
    height: number,
    movetarget?: MoveTarget.AsObject,
  }
}

export class AbilityModifier extends jspb.Message {
  getCancel(): boolean;
  setCancel(value: boolean): AbilityModifier;

  getCasttime(): number;
  setCasttime(value: number): AbilityModifier;

  getManacost(): number;
  setManacost(value: number): AbilityModifier;

  getCooldown(): number;
  setCooldown(value: number): AbilityModifier;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): AbilityModifier.AsObject;
  static toObject(includeInstance: boolean, msg: AbilityModifier): AbilityModifier.AsObject;
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
  setEffectid(value: Uint8Array | string): EffectModifier;

  getCancel(): boolean;
  setCancel(value: boolean): EffectModifier;

  getStat(): github_com_elojah_game_03_pkg_entity_entity_pb.Stat;
  setStat(value: github_com_elojah_game_03_pkg_entity_entity_pb.Stat): EffectModifier;

  getAmount(): Amount | undefined;
  setAmount(value?: Amount): EffectModifier;
  hasAmount(): boolean;
  clearAmount(): EffectModifier;

  getDuration(): number;
  setDuration(value: number): EffectModifier;

  getDelay(): number;
  setDelay(value: number): EffectModifier;

  getRepeat(): number;
  setRepeat(value: number): EffectModifier;

  getStackrules(): StackRules | undefined;
  setStackrules(value?: StackRules): EffectModifier;
  hasStackrules(): boolean;
  clearStackrules(): EffectModifier;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): EffectModifier.AsObject;
  static toObject(includeInstance: boolean, msg: EffectModifier): EffectModifier.AsObject;
  static serializeBinaryToWriter(message: EffectModifier, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): EffectModifier;
  static deserializeBinaryFromReader(message: EffectModifier, reader: jspb.BinaryReader): EffectModifier;
}

export namespace EffectModifier {
  export type AsObject = {
    effectid: Uint8Array | string,
    cancel: boolean,
    stat: github_com_elojah_game_03_pkg_entity_entity_pb.Stat,
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
  setId(value: Uint8Array | string): Amount;

  getDirect(): number;
  setDirect(value: number): Amount;

  getTarget(): Target | undefined;
  setTarget(value?: Target): Amount;
  hasTarget(): boolean;
  clearTarget(): Amount;

  getStat(): github_com_elojah_game_03_pkg_entity_entity_pb.Stat;
  setStat(value: github_com_elojah_game_03_pkg_entity_entity_pb.Stat): Amount;

  getPercentage(): number;
  setPercentage(value: number): Amount;

  getEffectid(): Uint8Array | string;
  getEffectid_asU8(): Uint8Array;
  getEffectid_asB64(): string;
  setEffectid(value: Uint8Array | string): Amount;

  getStatoutcome(): boolean;
  setStatoutcome(value: boolean): Amount;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Amount.AsObject;
  static toObject(includeInstance: boolean, msg: Amount): Amount.AsObject;
  static serializeBinaryToWriter(message: Amount, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Amount;
  static deserializeBinaryFromReader(message: Amount, reader: jspb.BinaryReader): Amount;
}

export namespace Amount {
  export type AsObject = {
    id: Uint8Array | string,
    direct: number,
    target?: Target.AsObject,
    stat: github_com_elojah_game_03_pkg_entity_entity_pb.Stat,
    percentage: number,
    effectid: Uint8Array | string,
    statoutcome: boolean,
  }
}

export class Trigger extends jspb.Message {
  getOperator(): Operator;
  setOperator(value: Operator): Trigger;

  getAmount(): Amount | undefined;
  setAmount(value?: Amount): Trigger;
  hasAmount(): boolean;
  clearAmount(): Trigger;

  getTreshold(): Amount | undefined;
  setTreshold(value?: Amount): Trigger;
  hasTreshold(): boolean;
  clearTreshold(): Trigger;

  getAbilitymodifiersMap(): jspb.Map<string, AbilityModifier>;
  clearAbilitymodifiersMap(): Trigger;

  getEffectmodifiersMap(): jspb.Map<string, EffectModifier>;
  clearEffectmodifiersMap(): Trigger;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Trigger.AsObject;
  static toObject(includeInstance: boolean, msg: Trigger): Trigger.AsObject;
  static serializeBinaryToWriter(message: Trigger, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Trigger;
  static deserializeBinaryFromReader(message: Trigger, reader: jspb.BinaryReader): Trigger;
}

export namespace Trigger {
  export type AsObject = {
    operator: Operator,
    amount?: Amount.AsObject,
    treshold?: Amount.AsObject,
    abilitymodifiersMap: Array<[string, AbilityModifier.AsObject]>,
    effectmodifiersMap: Array<[string, EffectModifier.AsObject]>,
  }
}

export class StackRules extends jspb.Message {
  getStacks(): number;
  setStacks(value: number): StackRules;

  getMaxstacks(): number;
  setMaxstacks(value: number): StackRules;

  getMaxduration(): number;
  setMaxduration(value: number): StackRules;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): StackRules.AsObject;
  static toObject(includeInstance: boolean, msg: StackRules): StackRules.AsObject;
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
  getStat(): github_com_elojah_game_03_pkg_entity_entity_pb.Stat;
  setStat(value: github_com_elojah_game_03_pkg_entity_entity_pb.Stat): Effect;

  getAmount(): Amount | undefined;
  setAmount(value?: Amount): Effect;
  hasAmount(): boolean;
  clearAmount(): Effect;

  getDuration(): number;
  setDuration(value: number): Effect;

  getIcon(): Uint8Array | string;
  getIcon_asU8(): Uint8Array;
  getIcon_asB64(): string;
  setIcon(value: Uint8Array | string): Effect;

  getDelay(): number;
  setDelay(value: number): Effect;

  getRepeat(): number;
  setRepeat(value: number): Effect;

  getStackrules(): StackRules | undefined;
  setStackrules(value?: StackRules): Effect;
  hasStackrules(): boolean;
  clearStackrules(): Effect;

  getTargetsMap(): jspb.Map<string, Target>;
  clearTargetsMap(): Effect;

  getTriggersMap(): jspb.Map<string, Trigger>;
  clearTriggersMap(): Effect;

  getEffectsMap(): jspb.Map<string, Effect>;
  clearEffectsMap(): Effect;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Effect.AsObject;
  static toObject(includeInstance: boolean, msg: Effect): Effect.AsObject;
  static serializeBinaryToWriter(message: Effect, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Effect;
  static deserializeBinaryFromReader(message: Effect, reader: jspb.BinaryReader): Effect;
}

export namespace Effect {
  export type AsObject = {
    stat: github_com_elojah_game_03_pkg_entity_entity_pb.Stat,
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
  setId(value: Uint8Array | string): A;

  getName(): string;
  setName(value: string): A;

  getIcon(): Uint8Array | string;
  getIcon_asU8(): Uint8Array;
  getIcon_asB64(): string;
  setIcon(value: Uint8Array | string): A;

  getAnimation(): Uint8Array | string;
  getAnimation_asU8(): Uint8Array;
  getAnimation_asB64(): string;
  setAnimation(value: Uint8Array | string): A;

  getCasttime(): number;
  setCasttime(value: number): A;

  getManacost(): number;
  setManacost(value: number): A;

  getCooldown(): number;
  setCooldown(value: number): A;

  getEffectsMap(): jspb.Map<string, Effect>;
  clearEffectsMap(): A;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): A.AsObject;
  static toObject(includeInstance: boolean, msg: A): A.AsObject;
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

export enum TargetType { 
  NONETARGET = 0,
  SELF = 1,
  FOE = 2,
  ALLY = 3,
  SPAWN = 4,
  RECT = 5,
  CIRCLE = 6,
}
export enum Move { 
  NONEMOVE = 0,
  WALK = 1,
  TELEPORT = 2,
  PUSH = 3,
}
export enum Operator { 
  NONEOPERATOR = 0,
  EQUAL = 1,
  NOTEQUAL = 2,
  GREATER = 3,
  LESSER = 4,
}
