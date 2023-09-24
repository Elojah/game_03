/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.6.1
 * source: github.com/elojah/game_03/pkg/entity/dto/template.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../../../../../gogo/protobuf/gogoproto/gogo";
import * as dependency_2 from "./../template";
import * as dependency_3 from "./../entity";
import * as pb_1 from "google-protobuf";
export namespace dto {
    export class CreateTemplateReq extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            Name?: string;
            Entity?: dependency_3.entity.E;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("Name" in data && data.Name != undefined) {
                    this.Name = data.Name;
                }
                if ("Entity" in data && data.Entity != undefined) {
                    this.Entity = data.Entity;
                }
            }
        }
        get Name() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set Name(value: string) {
            pb_1.Message.setField(this, 1, value);
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
            Name?: string;
            Entity?: ReturnType<typeof dependency_3.entity.E.prototype.toObject>;
        }): CreateTemplateReq {
            const message = new CreateTemplateReq({});
            if (data.Name != null) {
                message.Name = data.Name;
            }
            if (data.Entity != null) {
                message.Entity = dependency_3.entity.E.fromObject(data.Entity);
            }
            return message;
        }
        toObject() {
            const data: {
                Name?: string;
                Entity?: ReturnType<typeof dependency_3.entity.E.prototype.toObject>;
            } = {};
            if (this.Name != null) {
                data.Name = this.Name;
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
            if (this.Name.length)
                writer.writeString(1, this.Name);
            if (this.has_Entity)
                writer.writeMessage(2, this.Entity, () => this.Entity.serialize(writer));
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): CreateTemplateReq {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new CreateTemplateReq();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.Name = reader.readString();
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
        static deserializeBinary(bytes: Uint8Array): CreateTemplateReq {
            return CreateTemplateReq.deserialize(bytes);
        }
    }
    export class ListTemplateReq extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            All?: boolean;
            Size?: number;
            State?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("All" in data && data.All != undefined) {
                    this.All = data.All;
                }
                if ("Size" in data && data.Size != undefined) {
                    this.Size = data.Size;
                }
                if ("State" in data && data.State != undefined) {
                    this.State = data.State;
                }
            }
        }
        get All() {
            return pb_1.Message.getFieldWithDefault(this, 1, false) as boolean;
        }
        set All(value: boolean) {
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
            All?: boolean;
            Size?: number;
            State?: Uint8Array;
        }): ListTemplateReq {
            const message = new ListTemplateReq({});
            if (data.All != null) {
                message.All = data.All;
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
                All?: boolean;
                Size?: number;
                State?: Uint8Array;
            } = {};
            if (this.All != null) {
                data.All = this.All;
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
            if (this.All != false)
                writer.writeBool(1, this.All);
            if (this.Size != 0)
                writer.writeInt64(2, this.Size);
            if (this.State.length)
                writer.writeBytes(3, this.State);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): ListTemplateReq {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new ListTemplateReq();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.All = reader.readBool();
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
        static deserializeBinary(bytes: Uint8Array): ListTemplateReq {
            return ListTemplateReq.deserialize(bytes);
        }
    }
    export class ListTemplateResp extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            Templates?: dependency_2.entity.Template[];
            State?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [1], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("Templates" in data && data.Templates != undefined) {
                    this.Templates = data.Templates;
                }
                if ("State" in data && data.State != undefined) {
                    this.State = data.State;
                }
            }
        }
        get Templates() {
            return pb_1.Message.getRepeatedWrapperField(this, dependency_2.entity.Template, 1) as dependency_2.entity.Template[];
        }
        set Templates(value: dependency_2.entity.Template[]) {
            pb_1.Message.setRepeatedWrapperField(this, 1, value);
        }
        get State() {
            return pb_1.Message.getFieldWithDefault(this, 2, new Uint8Array(0)) as Uint8Array;
        }
        set State(value: Uint8Array) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            Templates?: ReturnType<typeof dependency_2.entity.Template.prototype.toObject>[];
            State?: Uint8Array;
        }): ListTemplateResp {
            const message = new ListTemplateResp({});
            if (data.Templates != null) {
                message.Templates = data.Templates.map(item => dependency_2.entity.Template.fromObject(item));
            }
            if (data.State != null) {
                message.State = data.State;
            }
            return message;
        }
        toObject() {
            const data: {
                Templates?: ReturnType<typeof dependency_2.entity.Template.prototype.toObject>[];
                State?: Uint8Array;
            } = {};
            if (this.Templates != null) {
                data.Templates = this.Templates.map((item: dependency_2.entity.Template) => item.toObject());
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
            if (this.Templates.length)
                writer.writeRepeatedMessage(1, this.Templates, (item: dependency_2.entity.Template) => item.serialize(writer));
            if (this.State.length)
                writer.writeBytes(2, this.State);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): ListTemplateResp {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new ListTemplateResp();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.Templates, () => pb_1.Message.addToRepeatedWrapperField(message, 1, dependency_2.entity.Template.deserialize(reader), dependency_2.entity.Template));
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
        static deserializeBinary(bytes: Uint8Array): ListTemplateResp {
            return ListTemplateResp.deserialize(bytes);
        }
    }
}
