/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.6.1
 * source: github.com/elojah/game_03/pkg/tile/dto/set.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../../../../../gogo/protobuf/gogoproto/gogo";
import * as pb_1 from "google-protobuf";
export namespace dto {
    export class CreateTilesetReq extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            ID?: Uint8Array;
            Set?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("ID" in data && data.ID != undefined) {
                    this.ID = data.ID;
                }
                if ("Set" in data && data.Set != undefined) {
                    this.Set = data.Set;
                }
            }
        }
        get ID() {
            return pb_1.Message.getFieldWithDefault(this, 1, new Uint8Array(0)) as Uint8Array;
        }
        set ID(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        get Set() {
            return pb_1.Message.getFieldWithDefault(this, 2, new Uint8Array(0)) as Uint8Array;
        }
        set Set(value: Uint8Array) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            ID?: Uint8Array;
            Set?: Uint8Array;
        }): CreateTilesetReq {
            const message = new CreateTilesetReq({});
            if (data.ID != null) {
                message.ID = data.ID;
            }
            if (data.Set != null) {
                message.Set = data.Set;
            }
            return message;
        }
        toObject() {
            const data: {
                ID?: Uint8Array;
                Set?: Uint8Array;
            } = {};
            if (this.ID != null) {
                data.ID = this.ID;
            }
            if (this.Set != null) {
                data.Set = this.Set;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.ID.length)
                writer.writeBytes(1, this.ID);
            if (this.Set.length)
                writer.writeBytes(2, this.Set);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): CreateTilesetReq {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new CreateTilesetReq();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.ID = reader.readBytes();
                        break;
                    case 2:
                        message.Set = reader.readBytes();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): CreateTilesetReq {
            return CreateTilesetReq.deserialize(bytes);
        }
    }
    export class CreateTilesetResp extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            ID?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("ID" in data && data.ID != undefined) {
                    this.ID = data.ID;
                }
            }
        }
        get ID() {
            return pb_1.Message.getFieldWithDefault(this, 1, new Uint8Array(0)) as Uint8Array;
        }
        set ID(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        static fromObject(data: {
            ID?: Uint8Array;
        }): CreateTilesetResp {
            const message = new CreateTilesetResp({});
            if (data.ID != null) {
                message.ID = data.ID;
            }
            return message;
        }
        toObject() {
            const data: {
                ID?: Uint8Array;
            } = {};
            if (this.ID != null) {
                data.ID = this.ID;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.ID.length)
                writer.writeBytes(1, this.ID);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): CreateTilesetResp {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new CreateTilesetResp();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
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
        static deserializeBinary(bytes: Uint8Array): CreateTilesetResp {
            return CreateTilesetResp.deserialize(bytes);
        }
    }
}
