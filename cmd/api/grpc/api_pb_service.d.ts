// package: grpc
// file: github.com/elojah/game_03/cmd/api/grpc/api.proto

import * as github_com_elojah_game_03_cmd_api_grpc_api_pb from "../../../../../../github.com/elojah/game_03/cmd/api/grpc/api_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as github_com_elojah_game_03_pkg_entity_pc_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/pc_pb";
import * as github_com_elojah_game_03_pkg_entity_dto_entity_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/dto/entity_pb";
import * as github_com_elojah_game_03_pkg_entity_dto_animation_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/dto/animation_pb";
import * as github_com_elojah_game_03_pkg_entity_dto_pc_pb from "../../../../../../github.com/elojah/game_03/pkg/entity/dto/pc_pb";
import * as github_com_elojah_game_03_pkg_room_room_pb from "../../../../../../github.com/elojah/game_03/pkg/room/room_pb";
import * as github_com_elojah_game_03_pkg_room_dto_cell_pb from "../../../../../../github.com/elojah/game_03/pkg/room/dto/cell_pb";
import * as github_com_elojah_game_03_pkg_room_dto_room_pb from "../../../../../../github.com/elojah/game_03/pkg/room/dto/room_pb";
import * as github_com_elojah_game_03_pkg_twitch_dto_follow_pb from "../../../../../../github.com/elojah/game_03/pkg/twitch/dto/follow_pb";
import {grpc} from "@improbable-eng/grpc-web";

type APIListEntity = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityResp;
};

type APIListAnimation = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationResp;
};

type APICreatePC = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_entity_dto_pc_pb.CreatePCReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_entity_pc_pb.PC;
};

type APIListPC = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCResp;
};

type APIConnectPC = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof github_com_elojah_game_03_pkg_entity_pc_pb.PC;
  readonly responseType: typeof github_com_elojah_game_03_pkg_room_dto_cell_pb.Cell;
};

type APICreateRoom = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_room_room_pb.R;
  readonly responseType: typeof github_com_elojah_game_03_pkg_room_room_pb.R;
};

type APIListRoom = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp;
};

type APIListCell = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellResp;
};

type APIListFollow = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp;
};

type APIPing = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class API {
  static readonly serviceName: string;
  static readonly ListEntity: APIListEntity;
  static readonly ListAnimation: APIListAnimation;
  static readonly CreatePC: APICreatePC;
  static readonly ListPC: APIListPC;
  static readonly ConnectPC: APIConnectPC;
  static readonly CreateRoom: APICreateRoom;
  static readonly ListRoom: APIListRoom;
  static readonly ListCell: APIListCell;
  static readonly ListFollow: APIListFollow;
  static readonly Ping: APIPing;
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

export class APIClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  listEntity(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityResp|null) => void
  ): UnaryResponse;
  listEntity(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityReq,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityResp|null) => void
  ): UnaryResponse;
  listAnimation(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationResp|null) => void
  ): UnaryResponse;
  listAnimation(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationReq,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationResp|null) => void
  ): UnaryResponse;
  createPC(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_pc_pb.CreatePCReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_entity_pc_pb.PC|null) => void
  ): UnaryResponse;
  createPC(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_pc_pb.CreatePCReq,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_entity_pc_pb.PC|null) => void
  ): UnaryResponse;
  listPC(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCResp|null) => void
  ): UnaryResponse;
  listPC(
    requestMessage: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCReq,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCResp|null) => void
  ): UnaryResponse;
  connectPC(requestMessage: github_com_elojah_game_03_pkg_entity_pc_pb.PC, metadata?: grpc.Metadata): ResponseStream<github_com_elojah_game_03_pkg_room_dto_cell_pb.Cell>;
  createRoom(
    requestMessage: github_com_elojah_game_03_pkg_room_room_pb.R,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_room_room_pb.R|null) => void
  ): UnaryResponse;
  createRoom(
    requestMessage: github_com_elojah_game_03_pkg_room_room_pb.R,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_room_room_pb.R|null) => void
  ): UnaryResponse;
  listRoom(
    requestMessage: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp|null) => void
  ): UnaryResponse;
  listRoom(
    requestMessage: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp|null) => void
  ): UnaryResponse;
  listCell(
    requestMessage: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellResp|null) => void
  ): UnaryResponse;
  listCell(
    requestMessage: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellReq,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellResp|null) => void
  ): UnaryResponse;
  listFollow(
    requestMessage: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp|null) => void
  ): UnaryResponse;
  listFollow(
    requestMessage: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp|null) => void
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

