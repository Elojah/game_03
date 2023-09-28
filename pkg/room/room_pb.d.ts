import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class R extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): R;

  getOwnerid(): Uint8Array | string;
  getOwnerid_asU8(): Uint8Array;
  getOwnerid_asB64(): string;
  setOwnerid(value: Uint8Array | string): R;

  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): R;

  getName(): string;
  setName(value: string): R;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): R.AsObject;
  static toObject(includeInstance: boolean, msg: R): R.AsObject;
  static serializeBinaryToWriter(message: R, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): R;
  static deserializeBinaryFromReader(message: R, reader: jspb.BinaryReader): R;
}

export namespace R {
  export type AsObject = {
    id: Uint8Array | string,
    ownerid: Uint8Array | string,
    worldid: Uint8Array | string,
    name: string,
  }
}

export class Public extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): Public;

  getRoomid(): Uint8Array | string;
  getRoomid_asU8(): Uint8Array;
  getRoomid_asB64(): string;
  setRoomid(value: Uint8Array | string): Public;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Public.AsObject;
  static toObject(includeInstance: boolean, msg: Public): Public.AsObject;
  static serializeBinaryToWriter(message: Public, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Public;
  static deserializeBinaryFromReader(message: Public, reader: jspb.BinaryReader): Public;
}

export namespace Public {
  export type AsObject = {
    id: Uint8Array | string,
    roomid: Uint8Array | string,
  }
}

