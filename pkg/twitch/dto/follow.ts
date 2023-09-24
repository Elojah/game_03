/**
 * Generated by the protoc-gen-ts.  DO NOT EDIT!
 * compiler version: 3.6.1
 * source: github.com/elojah/game_03/pkg/twitch/dto/follow.proto
 * git: https://github.com/thesayyn/protoc-gen-ts */
import * as dependency_1 from "./../../../../../gogo/protobuf/gogoproto/gogo";
import * as dependency_2 from "./../follow";
import * as pb_1 from "google-protobuf";
export namespace dto {
    export class ListFollowReq extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            After?: string;
            First?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("After" in data && data.After != undefined) {
                    this.After = data.After;
                }
                if ("First" in data && data.First != undefined) {
                    this.First = data.First;
                }
            }
        }
        get After() {
            return pb_1.Message.getFieldWithDefault(this, 1, "") as string;
        }
        set After(value: string) {
            pb_1.Message.setField(this, 1, value);
        }
        get First() {
            return pb_1.Message.getFieldWithDefault(this, 2, "") as string;
        }
        set First(value: string) {
            pb_1.Message.setField(this, 2, value);
        }
        static fromObject(data: {
            After?: string;
            First?: string;
        }): ListFollowReq {
            const message = new ListFollowReq({});
            if (data.After != null) {
                message.After = data.After;
            }
            if (data.First != null) {
                message.First = data.First;
            }
            return message;
        }
        toObject() {
            const data: {
                After?: string;
                First?: string;
            } = {};
            if (this.After != null) {
                data.After = this.After;
            }
            if (this.First != null) {
                data.First = this.First;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.After.length)
                writer.writeString(1, this.After);
            if (this.First.length)
                writer.writeString(2, this.First);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): ListFollowReq {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new ListFollowReq();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        message.After = reader.readString();
                        break;
                    case 2:
                        message.First = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): ListFollowReq {
            return ListFollowReq.deserialize(bytes);
        }
    }
    export class ListFollowResp extends pb_1.Message {
        #one_of_decls: number[][] = [];
        constructor(data?: any[] | {
            Follows?: dependency_2.twitch.Follow[];
            Total?: number;
            Cursor?: string;
        }) {
            super();
            pb_1.Message.initialize(this, Array.isArray(data) ? data : [], 0, -1, [1], this.#one_of_decls);
            if (!Array.isArray(data) && typeof data == "object") {
                if ("Follows" in data && data.Follows != undefined) {
                    this.Follows = data.Follows;
                }
                if ("Total" in data && data.Total != undefined) {
                    this.Total = data.Total;
                }
                if ("Cursor" in data && data.Cursor != undefined) {
                    this.Cursor = data.Cursor;
                }
            }
        }
        get Follows() {
            return pb_1.Message.getRepeatedWrapperField(this, dependency_2.twitch.Follow, 1) as dependency_2.twitch.Follow[];
        }
        set Follows(value: dependency_2.twitch.Follow[]) {
            pb_1.Message.setRepeatedWrapperField(this, 1, value);
        }
        get Total() {
            return pb_1.Message.getFieldWithDefault(this, 2, 0) as number;
        }
        set Total(value: number) {
            pb_1.Message.setField(this, 2, value);
        }
        get Cursor() {
            return pb_1.Message.getFieldWithDefault(this, 3, "") as string;
        }
        set Cursor(value: string) {
            pb_1.Message.setField(this, 3, value);
        }
        static fromObject(data: {
            Follows?: ReturnType<typeof dependency_2.twitch.Follow.prototype.toObject>[];
            Total?: number;
            Cursor?: string;
        }): ListFollowResp {
            const message = new ListFollowResp({});
            if (data.Follows != null) {
                message.Follows = data.Follows.map(item => dependency_2.twitch.Follow.fromObject(item));
            }
            if (data.Total != null) {
                message.Total = data.Total;
            }
            if (data.Cursor != null) {
                message.Cursor = data.Cursor;
            }
            return message;
        }
        toObject() {
            const data: {
                Follows?: ReturnType<typeof dependency_2.twitch.Follow.prototype.toObject>[];
                Total?: number;
                Cursor?: string;
            } = {};
            if (this.Follows != null) {
                data.Follows = this.Follows.map((item: dependency_2.twitch.Follow) => item.toObject());
            }
            if (this.Total != null) {
                data.Total = this.Total;
            }
            if (this.Cursor != null) {
                data.Cursor = this.Cursor;
            }
            return data;
        }
        serialize(): Uint8Array;
        serialize(w: pb_1.BinaryWriter): void;
        serialize(w?: pb_1.BinaryWriter): Uint8Array | void {
            const writer = w || new pb_1.BinaryWriter();
            if (this.Follows.length)
                writer.writeRepeatedMessage(1, this.Follows, (item: dependency_2.twitch.Follow) => item.serialize(writer));
            if (this.Total != 0)
                writer.writeUint64(2, this.Total);
            if (this.Cursor.length)
                writer.writeString(3, this.Cursor);
            if (!w)
                return writer.getResultBuffer();
        }
        static deserialize(bytes: Uint8Array | pb_1.BinaryReader): ListFollowResp {
            const reader = bytes instanceof pb_1.BinaryReader ? bytes : new pb_1.BinaryReader(bytes), message = new ListFollowResp();
            while (reader.nextField()) {
                if (reader.isEndGroup())
                    break;
                switch (reader.getFieldNumber()) {
                    case 1:
                        reader.readMessage(message.Follows, () => pb_1.Message.addToRepeatedWrapperField(message, 1, dependency_2.twitch.Follow.deserialize(reader), dependency_2.twitch.Follow));
                        break;
                    case 2:
                        message.Total = reader.readUint64();
                        break;
                    case 3:
                        message.Cursor = reader.readString();
                        break;
                    default: reader.skipField();
                }
            }
            return message;
        }
        serializeBinary(): Uint8Array {
            return this.serialize();
        }
        static deserializeBinary(bytes: Uint8Array): ListFollowResp {
            return ListFollowResp.deserialize(bytes);
        }
    }
}
