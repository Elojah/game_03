/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.6.1
 * source: github.com/elojah/game_03/pkg/ability/cast.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../../../../gogo/protobuf/gogoproto/gogo";
import * as dependency_2 from "./../geometry/geometry";
import * as dependency_3 from "./ability";
import * as pb_1 from "google-protobuf";
export namespace ability {
    export class CastTarget extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            ID?: Uint8Array;
            CellID?: Uint8Array;
            Rect?: dependency_2.geometry.Rect;
            Circle?: dependency_2.geometry.Circle;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("ID" in data && data.ID != undefined) {
                    this.ID = data.ID;
                }
                if ("CellID" in data && data.CellID != undefined) {
                    this.CellID = data.CellID;
                }
                if ("Rect" in data && data.Rect != undefined) {
                    this.Rect = data.Rect;
                }
                if ("Circle" in data && data.Circle != undefined) {
                    this.Circle = data.Circle;
                }
            }
        }
        get ID() {
            return pb_1.Message.getFieldWithDefault(this, 1, new Uint8Array(0)) as Uint8Array;
        }
        set ID(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        get CellID() {
            return pb_1.Message.getFieldWithDefault(this, 2, new Uint8Array(0)) as Uint8Array;
        }
        set CellID(value: Uint8Array) {
            pb_1.Message.setField(this, 2, value);
        }
        get Rect() {
            return pb_1.Message.getWrapperField(this, dependency_2.geometry.Rect, 3) as dependency_2.geometry.Rect;
        }
        set Rect(value: dependency_2.geometry.Rect) {
            pb_1.Message.setWrapperField(this, 3, value);
        }
        get has_Rect() {
            return pb_1.Message.getField(this, 3) != null;
        }
        get Circle() {
            return pb_1.Message.getWrapperField(this, dependency_2.geometry.Circle, 4) as dependency_2.geometry.Circle;
        }
        set Circle(value: dependency_2.geometry.Circle) {
            pb_1.Message.setWrapperField(this, 4, value);
        }
        get has_Circle() {
            return pb_1.Message.getField(this, 4) != null;
        }
        static fromObject(data: {
            ID?: Uint8Array;
            CellID?: Uint8Array;
            Rect?: ReturnType<typeof dependency_2.geometry.Rect.prototype.toObject>;
            Circle?: ReturnType<typeof dependency_2.geometry.Circle.prototype.toObject>;
        }): CastTarget {
            const message = new CastTarget({});
            if (data.ID != null) {
                message.ID = data.ID;
            }
            if (data.CellID != null) {
                message.CellID = data.CellID;
            }
            if (data.Rect != null) {
                message.Rect = dependency_2.geometry.Rect.fromObject(data.Rect);
            }
            if (data.Circle != null) {
                message.Circle = dependency_2.geometry.Circle.fromObject(data.Circle);
            }
            return message;
        }
        toObject() {
            const data: {
                ID?: Uint8Array;
                CellID?: Uint8Array;
                Rect?: ReturnType<typeof dependency_2.geometry.Rect.prototype.toObject>;
                Circle?: ReturnType<typeof dependency_2.geometry.Circle.prototype.toObject>;
            } = {};
            if (this.ID != null) {
                data.ID = this.ID;
            }
            if (this.CellID != null) {
                data.CellID = this.CellID;
            }
            if (this.Rect != null) {
                data.Rect = this.Rect.toObject();
            }
            if (this.Circle != null) {
                data.Circle = this.Circle.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.ID.length)
                writer.writeBytes(1, this.ID);
            if (this.CellID.length)
                writer.writeBytes(2, this.CellID);
            if (this.has_Rect)
                writer.writeMessage(3, this.Rect, () => this.Rect.serialize(writer));
            if (this.has_Circle)
                writer.writeMessage(4, this.Circle, () => this.Circle.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): CastTarget {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new CastTarget();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.ID = reader.readBytes();
                        break;
                    case 2:
                        message.CellID = reader.readBytes();
                        break;
                    case 3:
                        reader.readMessage(message.Rect, () => message.Rect = dependency_2.geometry.Rect.deserialize(reader));
                        break;
                    case 4:
                        reader.readMessage(message.Circle, () => message.Circle = dependency_2.geometry.Circle.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): CastTarget {
            return CastTarget.deserialize(bytes);
        }
    }
    export class Cast extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            ID?: Uint8Array;
            AbilityID?: Uint8Array;
            Targets?: Map<string, CastTarget>;
            At?: number;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("ID" in data && data.ID != undefined) {
                    this.ID = data.ID;
                }
                if ("AbilityID" in data && data.AbilityID != undefined) {
                    this.AbilityID = data.AbilityID;
                }
                if ("Targets" in data && data.Targets != undefined) {
                    this.Targets = data.Targets;
                }
                if ("At" in data && data.At != undefined) {
                    this.At = data.At;
                }
            }
            if (!this.Targets)
                this.Targets = new Map();
        }
        get ID() {
            return pb_1.Message.getFieldWithDefault(this, 1, new Uint8Array(0)) as Uint8Array;
        }
        set ID(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        get AbilityID() {
            return pb_1.Message.getFieldWithDefault(this, 2, new Uint8Array(0)) as Uint8Array;
        }
        set AbilityID(value: Uint8Array) {
            pb_1.Message.setField(this, 2, value);
        }
        get Targets() {
            return pb_1.Message.getField(this, 3) as any as Map<string, CastTarget>;
        }
        set Targets(value: Map<string, CastTarget>) {
            pb_1.Message.setField(this, 3, value as any);
        }
        get At() {
            return pb_1.Message.getFieldWithDefault(this, 4, 0) as number;
        }
        set At(value: number) {
            pb_1.Message.setField(this, 4, value);
        }
        static fromObject(data: {
            ID?: Uint8Array;
            AbilityID?: Uint8Array;
            Targets?: {
                [key: string]: ReturnType<typeof CastTarget.prototype.toObject>;
            };
            At?: number;
        }): Cast {
            const message = new Cast({});
            if (data.ID != null) {
                message.ID = data.ID;
            }
            if (data.AbilityID != null) {
                message.AbilityID = data.AbilityID;
            }
            if (typeof data.Targets == "object") {
                message.Targets = new Map(Object.entries(data.Targets).map(([key, value]) => [key, CastTarget.fromObject(value)]));
            }
            if (data.At != null) {
                message.At = data.At;
            }
            return message;
        }
        toObject() {
            const data: {
                ID?: Uint8Array;
                AbilityID?: Uint8Array;
                Targets?: {
                    [key: string]: ReturnType<typeof CastTarget.prototype.toObject>;
                };
                At?: number;
            } = {};
            if (this.ID != null) {
                data.ID = this.ID;
            }
            if (this.AbilityID != null) {
                data.AbilityID = this.AbilityID;
            }
            if (this.Targets != null) {
                data.Targets = (Object.fromEntries)((Array.from)(this.Targets).map(([key, value]) => [key, value.toObject()]));
            }
            if (this.At != null) {
                data.At = this.At;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.ID.length)
                writer.writeBytes(1, this.ID);
            if (this.AbilityID.length)
                writer.writeBytes(2, this.AbilityID);
            for (const [key, value] of this.Targets) {
                writer.writeMessage(3, this.Targets, () => {
                    writer.writeString(1, key);
                    writer.writeMessage(2, value, () => value.serialize(writer));
                });
            }
            if (this.At != 0)
                writer.writeInt64(4, this.At);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): Cast {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new Cast();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.ID = reader.readBytes();
                        break;
                    case 2:
                        message.AbilityID = reader.readBytes();
                        break;
                    case 3:
                        reader.readMessage(message, () => pb_1.Map.deserializeBinary(message.Targets as any, reader, reader.readString, () => {
                            let value;
                            reader.readMessage(message, () => value = CastTarget.deserialize(reader));
                            return value;
                        }));
                        break;
                    case 4:
                        message.At = reader.readInt64();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): Cast {
            return Cast.deserialize(bytes);
        }
    }
    export class CastEffect extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            ID?: Uint8Array;
            Self?: boolean;
            AbilityID?: Uint8Array;
            EffectID?: Uint8Array;
            CurrentID?: Uint8Array;
            Effect?: dependency_3.ability.Effect;
            Targets?: Map<string, CastTarget>;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("ID" in data && data.ID != undefined) {
                    this.ID = data.ID;
                }
                if ("Self" in data && data.Self != undefined) {
                    this.Self = data.Self;
                }
                if ("AbilityID" in data && data.AbilityID != undefined) {
                    this.AbilityID = data.AbilityID;
                }
                if ("EffectID" in data && data.EffectID != undefined) {
                    this.EffectID = data.EffectID;
                }
                if ("CurrentID" in data && data.CurrentID != undefined) {
                    this.CurrentID = data.CurrentID;
                }
                if ("Effect" in data && data.Effect != undefined) {
                    this.Effect = data.Effect;
                }
                if ("Targets" in data && data.Targets != undefined) {
                    this.Targets = data.Targets;
                }
            }
            if (!this.Targets)
                this.Targets = new Map();
        }
        get ID() {
            return pb_1.Message.getFieldWithDefault(this, 1, new Uint8Array(0)) as Uint8Array;
        }
        set ID(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        get Self() {
            return pb_1.Message.getFieldWithDefault(this, 2, false) as boolean;
        }
        set Self(value: boolean) {
            pb_1.Message.setField(this, 2, value);
        }
        get AbilityID() {
            return pb_1.Message.getFieldWithDefault(this, 3, new Uint8Array(0)) as Uint8Array;
        }
        set AbilityID(value: Uint8Array) {
            pb_1.Message.setField(this, 3, value);
        }
        get EffectID() {
            return pb_1.Message.getFieldWithDefault(this, 4, new Uint8Array(0)) as Uint8Array;
        }
        set EffectID(value: Uint8Array) {
            pb_1.Message.setField(this, 4, value);
        }
        get CurrentID() {
            return pb_1.Message.getFieldWithDefault(this, 5, new Uint8Array(0)) as Uint8Array;
        }
        set CurrentID(value: Uint8Array) {
            pb_1.Message.setField(this, 5, value);
        }
        get Effect() {
            return pb_1.Message.getWrapperField(this, dependency_3.ability.Effect, 6) as dependency_3.ability.Effect;
        }
        set Effect(value: dependency_3.ability.Effect) {
            pb_1.Message.setWrapperField(this, 6, value);
        }
        get has_Effect() {
            return pb_1.Message.getField(this, 6) != null;
        }
        get Targets() {
            return pb_1.Message.getField(this, 7) as any as Map<string, CastTarget>;
        }
        set Targets(value: Map<string, CastTarget>) {
            pb_1.Message.setField(this, 7, value as any);
        }
        static fromObject(data: {
            ID?: Uint8Array;
            Self?: boolean;
            AbilityID?: Uint8Array;
            EffectID?: Uint8Array;
            CurrentID?: Uint8Array;
            Effect?: ReturnType<typeof dependency_3.ability.Effect.prototype.toObject>;
            Targets?: {
                [key: string]: ReturnType<typeof CastTarget.prototype.toObject>;
            };
        }): CastEffect {
            const message = new CastEffect({});
            if (data.ID != null) {
                message.ID = data.ID;
            }
            if (data.Self != null) {
                message.Self = data.Self;
            }
            if (data.AbilityID != null) {
                message.AbilityID = data.AbilityID;
            }
            if (data.EffectID != null) {
                message.EffectID = data.EffectID;
            }
            if (data.CurrentID != null) {
                message.CurrentID = data.CurrentID;
            }
            if (data.Effect != null) {
                message.Effect = dependency_3.ability.Effect.fromObject(data.Effect);
            }
            if (typeof data.Targets == "object") {
                message.Targets = new Map(Object.entries(data.Targets).map(([key, value]) => [key, CastTarget.fromObject(value)]));
            }
            return message;
        }
        toObject() {
            const data: {
                ID?: Uint8Array;
                Self?: boolean;
                AbilityID?: Uint8Array;
                EffectID?: Uint8Array;
                CurrentID?: Uint8Array;
                Effect?: ReturnType<typeof dependency_3.ability.Effect.prototype.toObject>;
                Targets?: {
                    [key: string]: ReturnType<typeof CastTarget.prototype.toObject>;
                };
            } = {};
            if (this.ID != null) {
                data.ID = this.ID;
            }
            if (this.Self != null) {
                data.Self = this.Self;
            }
            if (this.AbilityID != null) {
                data.AbilityID = this.AbilityID;
            }
            if (this.EffectID != null) {
                data.EffectID = this.EffectID;
            }
            if (this.CurrentID != null) {
                data.CurrentID = this.CurrentID;
            }
            if (this.Effect != null) {
                data.Effect = this.Effect.toObject();
            }
            if (this.Targets != null) {
                data.Targets = (Object.fromEntries)((Array.from)(this.Targets).map(([key, value]) => [key, value.toObject()]));
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.ID.length)
                writer.writeBytes(1, this.ID);
            if (this.Self != false)
                writer.writeBool(2, this.Self);
            if (this.AbilityID.length)
                writer.writeBytes(3, this.AbilityID);
            if (this.EffectID.length)
                writer.writeBytes(4, this.EffectID);
            if (this.CurrentID.length)
                writer.writeBytes(5, this.CurrentID);
            if (this.has_Effect)
                writer.writeMessage(6, this.Effect, () => this.Effect.serialize(writer));
            for (const [key, value] of this.Targets) {
                writer.writeMessage(7, this.Targets, () => {
                    writer.writeString(1, key);
                    writer.writeMessage(2, value, () => value.serialize(writer));
                });
            }
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): CastEffect {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new CastEffect();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.ID = reader.readBytes();
                        break;
                    case 2:
                        message.Self = reader.readBool();
                        break;
                    case 3:
                        message.AbilityID = reader.readBytes();
                        break;
                    case 4:
                        message.EffectID = reader.readBytes();
                        break;
                    case 5:
                        message.CurrentID = reader.readBytes();
                        break;
                    case 6:
                        reader.readMessage(message.Effect, () => message.Effect = dependency_3.ability.Effect.deserialize(reader));
                        break;
                    case 7:
                        reader.readMessage(message, () => pb_1.Map.deserializeBinary(message.Targets as any, reader, reader.readString, () => {
                            let value;
                            reader.readMessage(message, () => value = CastTarget.deserialize(reader));
                            return value;
                        }));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): CastEffect {
            return CastEffect.deserialize(bytes);
        }
    }
}
