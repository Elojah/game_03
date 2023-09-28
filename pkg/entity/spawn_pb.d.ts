import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class Spawn extends jspb.Message {
  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): Spawn;

  getSpawnid(): Uint8Array | string;
  getSpawnid_asU8(): Uint8Array;
  getSpawnid_asB64(): string;
  setSpawnid(value: Uint8Array | string): Spawn;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): Spawn.AsObject;
  static toObject(includeInstance: boolean, msg: Spawn): Spawn.AsObject;
  static serializeBinaryToWriter(message: Spawn, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): Spawn;
  static deserializeBinaryFromReader(message: Spawn, reader: jspb.BinaryReader): Spawn;
}

export namespace Spawn {
  export type AsObject = {
    entityid: Uint8Array | string,
    spawnid: Uint8Array | string,
  }
}

