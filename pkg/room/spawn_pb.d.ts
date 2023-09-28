import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class WorldSpawn extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): WorldSpawn;

  getWorldid(): Uint8Array | string;
  getWorldid_asU8(): Uint8Array;
  getWorldid_asB64(): string;
  setWorldid(value: Uint8Array | string): WorldSpawn;

  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): WorldSpawn;

  getOwnerid(): Uint8Array | string;
  getOwnerid_asU8(): Uint8Array;
  getOwnerid_asB64(): string;
  setOwnerid(value: Uint8Array | string): WorldSpawn;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): WorldSpawn.AsObject;
  static toObject(includeInstance: boolean, msg: WorldSpawn): WorldSpawn.AsObject;
  static serializeBinaryToWriter(message: WorldSpawn, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): WorldSpawn;
  static deserializeBinaryFromReader(message: WorldSpawn, reader: jspb.BinaryReader): WorldSpawn;
}

export namespace WorldSpawn {
  export type AsObject = {
    id: Uint8Array | string,
    worldid: Uint8Array | string,
    entityid: Uint8Array | string,
    ownerid: Uint8Array | string,
  }
}

