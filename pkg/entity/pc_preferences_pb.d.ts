import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';


export class PCPreferences extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): PCPreferences;

  getPcid(): Uint8Array | string;
  getPcid_asU8(): Uint8Array;
  getPcid_asB64(): string;
  setPcid(value: Uint8Array | string): PCPreferences;

  getAbilityhotbarsMap(): jspb.Map<string, Uint8Array | string>;
  clearAbilityhotbarsMap(): PCPreferences;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PCPreferences.AsObject;
  static toObject(includeInstance: boolean, msg: PCPreferences): PCPreferences.AsObject;
  static serializeBinaryToWriter(message: PCPreferences, writer: jspb.BinaryWriter): void;
  static deserializeBinary(bytes: Uint8Array): PCPreferences;
  static deserializeBinaryFromReader(message: PCPreferences, reader: jspb.BinaryReader): PCPreferences;
}

export namespace PCPreferences {
  export type AsObject = {
    id: Uint8Array | string,
    pcid: Uint8Array | string,
    abilityhotbarsMap: Array<[string, Uint8Array | string]>,
  }
}

