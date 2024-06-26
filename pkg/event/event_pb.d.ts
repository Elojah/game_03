import * as jspb from 'google-protobuf'

import * as github_com_elojah_game_03_pkg_gogoproto_gogo_pb from '../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb';
import * as github_com_elojah_game_03_pkg_entity_entity_pb from '../../../../../github.com/elojah/game_03/pkg/entity/entity_pb';
import * as github_com_elojah_game_03_pkg_ability_cast_pb from '../../../../../github.com/elojah/game_03/pkg/ability/cast_pb';


export class E extends jspb.Message {
  getId(): Uint8Array | string;
  getId_asU8(): Uint8Array;
  getId_asB64(): string;
  setId(value: Uint8Array | string): E;

  getEntityid(): Uint8Array | string;
  getEntityid_asU8(): Uint8Array;
  getEntityid_asB64(): string;
  setEntityid(value: Uint8Array | string): E;

  getSource(): github_com_elojah_game_03_pkg_entity_entity_pb.E | undefined;
  setSource(value?: github_com_elojah_game_03_pkg_entity_entity_pb.E): E;
  hasSource(): boolean;
  clearSource(): E;

  getAt(): number;
  setAt(value: number): E;

  getEffect(): github_com_elojah_game_03_pkg_ability_cast_pb.CastEffect | undefined;
  setEffect(value?: github_com_elojah_game_03_pkg_ability_cast_pb.CastEffect): E;
  hasEffect(): boolean;
  clearEffect(): E;

  serializeBinary(): Uint8Array;
  toObject(includeInstance?: boolean): E.AsObject;
  static toObject(includeInstance: boolean, msg: E): E.AsObject;
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
    effect?: github_com_elojah_game_03_pkg_ability_cast_pb.CastEffect.AsObject,
  }
}

