import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class F extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): F;

  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): F;

  getName(): string;
  setName(value: string): F;

  getIcon(): string;
  setIcon(value: string): F;

  getMax(): number;
  setMax(value: number): F;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): F.AsObject;
  static toObject(includeInstance: boolean, msg: F): F.AsObject;
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
  setId(value: Uint8Array | string): PC;

  getFactionid(): Uint8Array | string;
  getFactionid_asU8(): Uint8Array;
  getFactionid_asB64(): string;
  setFactionid(value: Uint8Array | string): PC;

  getPermission(): number;
  setPermission(value: number): PC;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PC.AsObject;
  static toObject(includeInstance: boolean, msg: PC): PC.AsObject;
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

