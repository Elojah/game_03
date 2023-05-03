// package: event
// file: github.com/elojah/game_03/pkg/event/event.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";
import * as github_com_elojah_game_03_pkg_entity_entity_pb from "../../../../../github.com/elojah/game_03/pkg/entity/entity_pb";

export class E extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): void;

  hasSource(): boolean;
  clearSource(): void;
  getSource(): github_com_elojah_game_03_pkg_entity_entity_pb.E | undefined;
  setSource(value?: github_com_elojah_game_03_pkg_entity_entity_pb.E): void;

  getAt(): number;
  setAt(value: number): void;

  hasTrigger(): boolean;
  clearTrigger(): void;
  getTrigger(): E | undefined;
  setTrigger(value?: E): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): E.AsObject;
  static toObject(includeInstance: boolean, msg: E): E.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: E, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): E;
  static deserializeBinaryFromReader(message: E, reader: jspb.BinaryReader): E;
}

export namespace E {
  export type AsObject = {
    id: Uint8Array | string,
    entityid: Uint8Array | string,
    source?: github_com_elojah_game_03_pkg_entity_entity_pb.E.AsObject,
    at: number,
    trigger?: E.AsObject,
  }
}

