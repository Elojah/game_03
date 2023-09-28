import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class U extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): U;

  getTwitchid(): string;
  setTwitchid(value: string): U;

  getGoogleid(): string;
  setGoogleid(value: string): U;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): U.AsObject;
  static toObject(includeInstance: boolean, msg: U): U.AsObject;
  static serializeBinaryToWriter(message: U, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): U;
  static deserializeBinaryFromReader(message: U, reader: jspb.BinaryReader): U;
}

export namespace U {
  export type AsObject = {
    id: Uint8Array | string,
    twitchid: string,
    googleid: string,
  }
}

