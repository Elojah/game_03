import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class PC extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): PC;

  getUserid(): Uint8Array | string;
  getUserid_asU8(): Uint8Array;
  getUserid_asB64(): string;
  setUserid(value: Uint8Array | string): PC;

  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): PC;

  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): PC;

  getAt(): number;
  setAt(value: number): PC;

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
    userid: Uint8Array | string,
    worldid: Uint8Array | string,
    entityid: Uint8Array | string,
    at: number,
  }
}

