/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.6.1
 * source: github.com/elojah/game_03/pkg/room/dto/user.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../../../../../gogo/protobuf/gogoproto/gogo";
import * as pb_1 from "google-protobuf";
export namespace dto {
    export class CreateRoomUserReq extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            RoomID?: Uint8Array;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("RoomID" in data && data.RoomID != undefined) {
                    this.RoomID = data.RoomID;
                }
            }
        }
        get RoomID() {
            return pb_1.Message.getFieldWithDefault(this, 1, new Uint8Array(0)) as Uint8Array;
        }
        set RoomID(value: Uint8Array) {
            pb_1.Message.setField(this, 1, value);
        }
        static fromObject(data: {
            RoomID?: Uint8Array;
        }): CreateRoomUserReq {
            const message = new CreateRoomUserReq({});
            if (data.RoomID != null) {
                message.RoomID = data.RoomID;
            }
            return message;
        }
        toObject() {
            const data: {
                RoomID?: Uint8Array;
            } = {};
            if (this.RoomID != null) {
                data.RoomID = this.RoomID;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.RoomID.length)
                writer.writeBytes(1, this.RoomID);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): CreateRoomUserReq {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new CreateRoomUserReq();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.RoomID = reader.readBytes();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): CreateRoomUserReq {
            return CreateRoomUserReq.deserialize(bytes);
        }
    }
}
