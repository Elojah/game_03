// package: grpc
// file: github.com/elojah/game_03/cmd/core/grpc/core.proto

import * as github_com_elojah_game_03_cmd_core_grpc_core_pb from "../../../../../../github.com/elojah/game_03/cmd/core/grpc/core_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as github_com_elojah_game_03_pkg_rtc_dto_rtc_pb from "../../../../../../github.com/elojah/game_03/pkg/rtc/dto/rtc_pb";
import {grpc} from "@improbable-eng/grpc-web";

type CoreConnect = {
  readonly methodName: string;
  readonly service: typeof Core;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ConnectReq;
  readonly responseType: typeof github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.SDP;
};

type CoreSendICE = {
  readonly methodName: string;
  readonly service: typeof Core;
  readonly requestStream: true;
  readonly responseStream: false;
  readonly requestType: typeof github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

type CoreReceiveICE = {
  readonly methodName: string;
  readonly service: typeof Core;
  readonly requestStream: false;
  readonly responseStream: true;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate;
};

type CorePing = {
  readonly methodName: string;
  readonly service: typeof Core;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class Core {
  static readonly serviceName: string;
  static readonly Connect: CoreConnect;
  static readonly SendICE: CoreSendICE;
  static readonly ReceiveICE: CoreReceiveICE;
  static readonly Ping: CorePing;
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

export class CoreClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  connect(
    requestMessage: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ConnectReq,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.SDP|null) => void
  ): UnaryResponse;
  connect(
    requestMessage: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ConnectReq,
    callback: (error: ServiceError|null, responseMessage: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.SDP|null) => void
  ): UnaryResponse;
  sendICE(metadata?: grpc.Metadata): RequestStream<github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate>;
  receiveICE(requestMessage: google_protobuf_empty_pb.Empty, metadata?: grpc.Metadata): ResponseStream<github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate>;
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

