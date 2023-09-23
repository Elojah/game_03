// package: entity
// file: github.com/elojah/game_03/pkg/entity/pc_preferences.proto

import * as jspb from "google-protobuf";
import * as github_com_gogo_protobuf_gogoproto_gogo_pb from "../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb";

export class PCPreferences extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): void;

  getPcid(): Uint8Array | string;
  getPcid_asU8(): Uint8Array;
  getPcid_asB64(): string;
  setPcid(value: Uint8Array | string): void;

  getAbilityhotbarsMap(): jspb.Map<string, Uint8Array | string>;
  clearAbilityhotbarsMap(): void;
  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): PCPreferences.AsObject;
  static toObject(includeInstance: boolean, msg: PCPreferences): PCPreferences.AsObject;
  static extensions: {[key: number]: jspb.ExtensionFieldInfo<jspb.Message>};
  static extensionsBinary: {[key: number]: jspb.ExtensionFieldBinaryInfo<jspb.Message>};
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

