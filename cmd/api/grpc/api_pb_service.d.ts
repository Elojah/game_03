// package: grpc
// file: github.com/elojah/game_03/cmd/api/grpc/api.proto

import * as github_com_elojah_game_03_cmd_api_grpc_api_pb from "../../../../../../github.com/elojah/game_03/cmd/api/grpc/api_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as github_com_elojah_game_03_pkg_twitch_dto_follow_pb from "../../../../../../github.com/elojah/game_03/pkg/twitch/dto/follow_pb";
import * as github_com_elojah_game_03_pkg_room_dto_room_pb from "../../../../../../github.com/elojah/game_03/pkg/room/dto/room_pb";
import {grpc} from "@improbable-eng/grpc-web";

type APIListFollow = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp;
};

type APIListRoom = {
  readonly methodName: string;
  readonly service: typeof API;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp;
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
  static readonly ListFollow: APIListFollow;
  static readonly ListRoom: APIListRoom;
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
  listFollow(
    requestMessage: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp|null) => void
  ): UnaryResponse;
  listFollow(
    requestMessage: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp|null) => void
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

