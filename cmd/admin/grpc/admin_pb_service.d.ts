// package: grpc
// file: github.com/elojah/game_03/cmd/admin/grpc/admin.proto

import * as github_com_elojah_game_03_cmd_admin_grpc_admin_pb from "../../../../../../github.com/elojah/game_03/cmd/admin/grpc/admin_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import * as github_com_elojah_game_03_pkg_tile_dto_set_pb from "../../../../../../github.com/elojah/game_03/pkg/tile/dto/set_pb";
import * as github_com_elojah_game_03_pkg_entity_template_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/template_pb";
import * as github_com_elojah_game_03_pkg_entity_dto_animation_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/dto/animation_pb";
import * as github_com_elojah_game_03_pkg_entity_dto_template_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/dto/template_pb";
import {grpc} from "@improbable-eng/grpc-web";

type AdminMigrateUp = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_wrappers_pb.StringValue;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type AdminCreateTileset = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_tile_dto_set_pb.CreateTilesetReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_tile_dto_set_pb.CreateTilesetResp;
};

type AdminCreateWorld = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof google_protobuf_wrappers_pb.StringValue;
};

type AdminCreateTemplate = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_entity_dto_template_pb.CreateTemplateReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_entity_template_pb.Template;
};

type AdminCreateAnimation = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_entity_dto_animation_pb.CreateAnimationReq;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type AdminPing = {
  readonly methodName: string;
  readonly service: typeof Admin;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class Admin {
  static readonly serviceName: string;
  static readonly MigrateUp: AdminMigrateUp;
  static readonly CreateTileset: AdminCreateTileset;
  static readonly CreateWorld: AdminCreateWorld;
  static readonly CreateTemplate: AdminCreateTemplate;
  static readonly CreateAnimation: AdminCreateAnimation;
  static readonly Ping: AdminPing;
}

export type ServiceError = { message: string, code: number; metadata: grpc.Metadata }
export type Status = { details: string, code: number; metadata: grpc.Metadata }

interface UnaryResponse {
  cancel(): void;
}
interface ResponseStream<T> {
  cancel(): void;
  on(type: 'data', handler: (message: T) => void): ResponseStream<T>;
  on(type: 'end', handler: (status?: Status) => void): ResponseStream<T>;
  on(type: 'status', handler: (status: Status) => void): ResponseStream<T>;
}
interface RequestStream<T> {
  write(message: T): RequestStream<T>;
  end(): void;
  cancel(): void;
  on(type: 'end', handler: (status?: Status) => void): RequestStream<T>;
  on(type: 'status', handler: (status: Status) => void): RequestStream<T>;
}
interface BidirectionalStream<ReqT, ResT> {
  write(message: ReqT): BidirectionalStream<ReqT, ResT>;
  end(): void;
  cancel(): void;
  on(type: 'data', handler: (message: ResT) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'end', handler: (status?: Status) => void): BidirectionalStream<ReqT, ResT>;
  on(type: 'status', handler: (status: Status) => void): BidirectionalStream<ReqT, ResT>;
}

export class AdminClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  migrateUp(
    requestMessage: google_protobuf_wrappers_pb.StringValue,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  migrateUp(
    requestMessage: google_protobuf_wrappers_pb.StringValue,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  createTileset(
    requestMessage: github_com_elojah_game_03_pkg_tile_dto_set_pb.CreateTilesetReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_tile_dto_set_pb.CreateTilesetResp|null) => void
  ): UnaryResponse;
  createTileset(
    requestMessage: github_com_elojah_game_03_pkg_tile_dto_set_pb.CreateTilesetReq,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_tile_dto_set_pb.CreateTilesetResp|null) => void
  ): UnaryResponse;
  createWorld(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_wrappers_pb.StringValue|null) => void
  ): UnaryResponse;
  createWorld(
    requestMessage: google_protobuf_empty_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_wrappers_pb.StringValue|null) => void
  ): UnaryResponse;
  createTemplate(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_template_pb.CreateTemplateReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_entity_template_pb.Template|null) => void
  ): UnaryResponse;
  createTemplate(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_template_pb.CreateTemplateReq,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_entity_template_pb.Template|null) => void
  ): UnaryResponse;
  createAnimation(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_animation_pb.CreateAnimationReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  createAnimation(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_animation_pb.CreateAnimationReq,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  ping(
    requestMessage: google_protobuf_empty_pb.Empty,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
  ping(
    requestMessage: google_protobuf_empty_pb.Empty,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_empty_pb.Empty|null) => void
  ): UnaryResponse;
}

