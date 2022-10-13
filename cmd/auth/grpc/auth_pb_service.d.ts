// package: grpc
// file: github.com/elojah/game_03/cmd/auth/grpc/auth.proto

import * as github_com_elojah_game_03_cmd_auth_grpc_auth_pb from "../../../../../../github.com/elojah/game_03/cmd/auth/grpc/auth_pb";
import * as google_protobuf_empty_pb from "google-protobuf/google/protobuf/empty_pb";
import * as google_protobuf_wrappers_pb from "google-protobuf/google/protobuf/wrappers_pb";
import {grpc} from "@improbable-eng/grpc-web";

type AuthSigninTwitch = {
  readonly methodName: string;
  readonly service: typeof Auth;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_wrappers_pb.StringValue;
  readonly responseType: typeof google_protobuf_wrappers_pb.StringValue;
};

type AuthSigninGoogle = {
  readonly methodName: string;
  readonly service: typeof Auth;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_wrappers_pb.StringValue;
  readonly responseType: typeof google_protobuf_wrappers_pb.StringValue;
};

type AuthPing = {
  readonly methodName: string;
  readonly service: typeof Auth;
  readonly requestStream: false;
  readonly responseStream: false;
  readonly requestType: typeof google_protobuf_empty_pb.Empty;
  readonly responseType: typeof google_protobuf_empty_pb.Empty;
};

export class Auth {
  static readonly serviceName: string;
  static readonly SigninTwitch: AuthSigninTwitch;
  static readonly SigninGoogle: AuthSigninGoogle;
  static readonly Ping: AuthPing;
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

export class AuthClient {
  readonly serviceHost: string;

  constructor(serviceHost: string, options?: grpc.RpcOptions);
  signinTwitch(
    requestMessage: google_protobuf_wrappers_pb.StringValue,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_wrappers_pb.StringValue|null) => void
  ): UnaryResponse;
  signinTwitch(
    requestMessage: google_protobuf_wrappers_pb.StringValue,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_wrappers_pb.StringValue|null) => void
  ): UnaryResponse;
  signinGoogle(
    requestMessage: google_protobuf_wrappers_pb.StringValue,
    metadata: grpc.Metadata,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_wrappers_pb.StringValue|null) => void
  ): UnaryResponse;
  signinGoogle(
    requestMessage: google_protobuf_wrappers_pb.StringValue,
    callback: (error: ServiceError|null, responseMessage: google_protobuf_wrappers_pb.StringValue|null) => void
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

