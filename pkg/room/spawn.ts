/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.6.1
 * source: github.com/elojah/game_03/pkg/room/spawn.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../../../../gogo/protobuf/gogoproto/gogo";
import * as pb_1 from "google-protobuf";
export namespace room {
    export class WorldSpawn extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            ID?: Uint8Array;
            WorldID?: Uint8Array;
            EntityID?: Uint8Array;
            OwnerID?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("ID" in data && data.ID != undefined) {
                    this.ID = data.ID;
                }
                if ("WorldID" in data && data.WorldID != undefined) {
                    this.WorldID = data.WorldID;
                }
                if ("EntityID" in data && data.EntityID != undefined) {
                    this.EntityID = data.EntityID;
                }
                if ("OwnerID" in data && data.OwnerID != undefined) {
                    this.OwnerID = data.OwnerID;
                }
            }
        }
        get ID() {
            return pb_1.Message.getFieldWithDefault(this, 1, new Uint8Array(0)) as Uint8Array;
        }
        set ID(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        get WorldID() {
            return pb_1.Message.getFieldWithDefault(this, 2, new Uint8Array(0)) as Uint8Array;
        }
        set WorldID(value: Uint8Array) {
            pb_1.Message.setField(this, 2, value);
        }
        get EntityID() {
            return pb_1.Message.getFieldWithDefault(this, 3, new Uint8Array(0)) as Uint8Array;
        }
        set EntityID(value: Uint8Array) {
            pb_1.Message.setField(this, 3, value);
        }
        get OwnerID() {
            return pb_1.Message.getFieldWithDefault(this, 4, new Uint8Array(0)) as Uint8Array;
        }
        set OwnerID(value: Uint8Array) {
            pb_1.Message.setField(this, 4, value);
        }
        static fromObject(data: {
            ID?: Uint8Array;
            WorldID?: Uint8Array;
            EntityID?: Uint8Array;
            OwnerID?: Uint8Array;
        }): WorldSpawn {
            const message = new WorldSpawn({});
            if (data.ID != null) {
                message.ID = data.ID;
            }
            if (data.WorldID != null) {
                message.WorldID = data.WorldID;
            }
            if (data.EntityID != null) {
                message.EntityID = data.EntityID;
            }
            if (data.OwnerID != null) {
                message.OwnerID = data.OwnerID;
            }
            return message;
        }
        toObject() {
            const data: {
                ID?: Uint8Array;
                WorldID?: Uint8Array;
                EntityID?: Uint8Array;
                OwnerID?: Uint8Array;
            } = {};
            if (this.ID != null) {
                data.ID = this.ID;
            }
            if (this.WorldID != null) {
                data.WorldID = this.WorldID;
            }
            if (this.EntityID != null) {
                data.EntityID = this.EntityID;
            }
            if (this.OwnerID != null) {
                data.OwnerID = this.OwnerID;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.ID.length)
                writer.writeBytes(1, this.ID);
            if (this.WorldID.length)
                writer.writeBytes(2, this.WorldID);
            if (this.EntityID.length)
                writer.writeBytes(3, this.EntityID);
            if (this.OwnerID.length)
                writer.writeBytes(4, this.OwnerID);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): WorldSpawn {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new WorldSpawn();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.ID = reader.readBytes();
                        break;
                    case 2:
                        message.WorldID = reader.readBytes();
                        break;
                    case 3:
                        message.EntityID = reader.readBytes();
                        break;
                    case 4:
                        message.OwnerID = reader.readBytes();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): WorldSpawn {
            return WorldSpawn.deserialize(bytes);
        }
    }
}