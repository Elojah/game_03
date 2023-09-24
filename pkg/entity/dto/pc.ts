/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.6.1
 * source: github.com/elojah/game_03/pkg/entity/dto/pc.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../../../../../gogo/protobuf/gogoproto/gogo";
import * as dependency_2 from "./../pc";
import * as dependency_3 from "./../entity";
import * as pb_1 from "google-protobuf";
export namespace dto {
    export class PC extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            PC?: dependency_2.entity.PC;
            Entity?: dependency_3.entity.E;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("PC" in data && data.PC != undefined) {
                    this.PC = data.PC;
                }
                if ("Entity" in data && data.Entity != undefined) {
                    this.Entity = data.Entity;
                }
            }
        }
        get PC() {
            return pb_1.Message.getWrapperField(this, dependency_2.entity.PC, 1) as dependency_2.entity.PC;
        }
        set PC(value: dependency_2.entity.PC) {
            pb_1.Message.setWrapperField(this, 1, value);
        }
        get has_PC() {
            return pb_1.Message.getField(this, 1) != null;
        }
        get Entity() {
            return pb_1.Message.getWrapperField(this, dependency_3.entity.E, 2) as dependency_3.entity.E;
        }
        set Entity(value: dependency_3.entity.E) {
            pb_1.Message.setWrapperField(this, 2, value);
        }
        get has_Entity() {
            return pb_1.Message.getField(this, 2) != null;
        }
        static fromObject(data: {
            PC?: ReturnType<typeof dependency_2.entity.PC.prototype.toObject>;
            Entity?: ReturnType<typeof dependency_3.entity.E.prototype.toObject>;
        }): PC {
            const message = new PC({});
            if (data.PC != null) {
                message.PC = dependency_2.entity.PC.fromObject(data.PC);
            }
            if (data.Entity != null) {
                message.Entity = dependency_3.entity.E.fromObject(data.Entity);
            }
            return message;
        }
        toObject() {
            const data: {
                PC?: ReturnType<typeof dependency_2.entity.PC.prototype.toObject>;
                Entity?: ReturnType<typeof dependency_3.entity.E.prototype.toObject>;
            } = {};
            if (this.PC != null) {
                data.PC = this.PC.toObject();
            }
            if (this.Entity != null) {
                data.Entity = this.Entity.toObject();
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.has_PC)
                writer.writeMessage(1, this.PC, () => this.PC.serialize(writer));
            if (this.has_Entity)
                writer.writeMessage(2, this.Entity, () => this.Entity.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): PC {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new PC();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.PC, () => message.PC = dependency_2.entity.PC.deserialize(reader));
                        break;
                    case 2:
                        reader.readMessage(message.Entity, () => message.Entity = dependency_3.entity.E.deserialize(reader));
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): PC {
            return PC.deserialize(bytes);
        }
    }
    export class CreatePCReq extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            RoomID?: Uint8Array;
            EntityTemplate?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("RoomID" in data && data.RoomID != undefined) {
                    this.RoomID = data.RoomID;
                }
                if ("EntityTemplate" in data && data.EntityTemplate != undefined) {
                    this.EntityTemplate = data.EntityTemplate;
                }
            }
        }
        get RoomID() {
            return pb_1.Message.getFieldWithDefault(this, 1, new Uint8Array(0)) as Uint8Array;
        }
        set RoomID(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        get EntityTemplate() {
            return pb_1.Message.getFieldWithDefault(this, 2, "") as string;
        }
        set EntityTemplate(value: string) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            RoomID?: Uint8Array;
            EntityTemplate?: string;
        }): CreatePCReq {
            const message = new CreatePCReq({});
            if (data.RoomID != null) {
                message.RoomID = data.RoomID;
            }
            if (data.EntityTemplate != null) {
                message.EntityTemplate = data.EntityTemplate;
            }
            return message;
        }
        toObject() {
            const data: {
                RoomID?: Uint8Array;
                EntityTemplate?: string;
            } = {};
            if (this.RoomID != null) {
                data.RoomID = this.RoomID;
            }
            if (this.EntityTemplate != null) {
                data.EntityTemplate = this.EntityTemplate;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.RoomID.length)
                writer.writeBytes(1, this.RoomID);
            if (this.EntityTemplate.length)
                writer.writeString(2, this.EntityTemplate);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): CreatePCReq {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new CreatePCReq();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.RoomID = reader.readBytes();
                        break;
                    case 2:
                        message.EntityTemplate = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): CreatePCReq {
            return CreatePCReq.deserialize(bytes);
        }
    }
    export class ListPCReq extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            RoomID?: Uint8Array;
            Size?: number;
            State?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("RoomID" in data && data.RoomID != undefined) {
                    this.RoomID = data.RoomID;
                }
                if ("Size" in data && data.Size != undefined) {
                    this.Size = data.Size;
                }
                if ("State" in data && data.State != undefined) {
                    this.State = data.State;
                }
            }
        }
        get RoomID() {
            return pb_1.Message.getFieldWithDefault(this, 1, new Uint8Array(0)) as Uint8Array;
        }
        set RoomID(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        get Size() {
            return pb_1.Message.getFieldWithDefault(this, 2, 0) as number;
        }
        set Size(value: number) {
            pb_1.Message.setField(this, 2, value);
        }
        get State() {
            return pb_1.Message.getFieldWithDefault(this, 3, new Uint8Array(0)) as Uint8Array;
        }
        set State(value: Uint8Array) {
            pb_1.Message.setField(this, 3, value);
        }
        static fromObject(data: {
            RoomID?: Uint8Array;
            Size?: number;
            State?: Uint8Array;
        }): ListPCReq {
            const message = new ListPCReq({});
            if (data.RoomID != null) {
                message.RoomID = data.RoomID;
            }
            if (data.Size != null) {
                message.Size = data.Size;
            }
            if (data.State != null) {
                message.State = data.State;
            }
            return message;
        }
        toObject() {
            const data: {
                RoomID?: Uint8Array;
                Size?: number;
                State?: Uint8Array;
            } = {};
            if (this.RoomID != null) {
                data.RoomID = this.RoomID;
            }
            if (this.Size != null) {
                data.Size = this.Size;
            }
            if (this.State != null) {
                data.State = this.State;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.RoomID.length)
                writer.writeBytes(1, this.RoomID);
            if (this.Size != 0)
                writer.writeInt64(2, this.Size);
            if (this.State.length)
                writer.writeBytes(3, this.State);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): ListPCReq {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new ListPCReq();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.RoomID = reader.readBytes();
                        break;
                    case 2:
                        message.Size = reader.readInt64();
                        break;
                    case 3:
                        message.State = reader.readBytes();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): ListPCReq {
            return ListPCReq.deserialize(bytes);
        }
    }
    export class ListPCResp extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            PCs?: PC[];
            State?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [1], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("PCs" in data && data.PCs != undefined) {
                    this.PCs = data.PCs;
                }
                if ("State" in data && data.State != undefined) {
                    this.State = data.State;
                }
            }
        }
        get PCs() {
            return pb_1.Message.getRepeatedWrapperField(this, PC, 1) as PC[];
        }
        set PCs(value: PC[]) {
            pb_1.Message.setRepeatedWrapperField(this, 1, value);
        }
        get State() {
            return pb_1.Message.getFieldWithDefault(this, 2, new Uint8Array(0)) as Uint8Array;
        }
        set State(value: Uint8Array) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            PCs?: ReturnType<typeof PC.prototype.toObject>[];
            State?: Uint8Array;
        }): ListPCResp {
            const message = new ListPCResp({});
            if (data.PCs != null) {
                message.PCs = data.PCs.map(item => PC.fromObject(item));
            }
            if (data.State != null) {
                message.State = data.State;
            }
            return message;
        }
        toObject() {
            const data: {
                PCs?: ReturnType<typeof PC.prototype.toObject>[];
                State?: Uint8Array;
            } = {};
            if (this.PCs != null) {
                data.PCs = this.PCs.map((item: PC) => item.toObject());
            }
            if (this.State != null) {
                data.State = this.State;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.PCs.length)
                writer.writeRepeatedMessage(1, this.PCs, (item: PC) => item.serialize(writer));
            if (this.State.length)
                writer.writeBytes(2, this.State);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): ListPCResp {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new ListPCResp();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.PCs, () => pb_1.Message.addToRepeatedWrapperField(message, 1, PC.deserialize(reader), PC));
                        break;
                    case 2:
                        message.State = reader.readBytes();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): ListPCResp {
            return ListPCResp.deserialize(bytes);
        }
    }
    export class GetPCReq extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            WorldID?: Uint8Array;
            ID?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("WorldID" in data && data.WorldID != undefined) {
                    this.WorldID = data.WorldID;
                }
                if ("ID" in data && data.ID != undefined) {
                    this.ID = data.ID;
                }
            }
        }
        get WorldID() {
            return pb_1.Message.getFieldWithDefault(this, 1, new Uint8Array(0)) as Uint8Array;
        }
        set WorldID(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        get ID() {
            return pb_1.Message.getFieldWithDefault(this, 2, new Uint8Array(0)) as Uint8Array;
        }
        set ID(value: Uint8Array) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            WorldID?: Uint8Array;
            ID?: Uint8Array;
        }): GetPCReq {
            const message = new GetPCReq({});
            if (data.WorldID != null) {
                message.WorldID = data.WorldID;
            }
            if (data.ID != null) {
                message.ID = data.ID;
            }
            return message;
        }
        toObject() {
            const data: {
                WorldID?: Uint8Array;
                ID?: Uint8Array;
            } = {};
            if (this.WorldID != null) {
                data.WorldID = this.WorldID;
            }
            if (this.ID != null) {
                data.ID = this.ID;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.WorldID.length)
                writer.writeBytes(1, this.WorldID);
            if (this.ID.length)
                writer.writeBytes(2, this.ID);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): GetPCReq {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new GetPCReq();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.WorldID = reader.readBytes();
                        break;
                    case 2:
                        message.ID = reader.readBytes();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): GetPCReq {
            return GetPCReq.deserialize(bytes);
        }
    }
}