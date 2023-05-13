// package: faction
// file: github.com/elojah/game_03/pkg/faction/faction.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class F extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): void;

  getName(): string;
  setName(value: string): void;

  getIcon(): string;
  setIcon(value: string): void;

  getMax(): number;
  setMax(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): F.AsObject;
  static toObject(includeInstance: boolean, msg: F): F.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: F, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): F;
  static deserializeBinaryFromReader(message: F, reader: jspb.BinaryReader): F;
}

export namespace F {
  export type AsObject = {
    id: Uint8Array | string,
    worldid: Uint8Array | string,
    name: string,
    icon: string,
    max: number,
  }
}

export class PC extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getFactionid(): Uint8Array | string;
  getFactionid_asU8(): Uint8Array;
  getFactionid_asB64(): string;
  setFactionid(value: Uint8Array | string): void;

  getPermission(): number;
  setPermission(value: number): void;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PC.AsObject;
  static toObject(includeInstance: boolean, msg: PC): PC.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
  static serializeBinaryToWriter(message: PC, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PC;
  static deserializeBinaryFromReader(message: PC, reader: jspb.BinaryReader): PC;
}

export namespace PC {
  export type AsObject = {
    id: Uint8Array | string,
    factionid: Uint8Array | string,
    permission: number,
  }
}

