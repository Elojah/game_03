/**
 * @fileoverview gRPC-Web generated client stub for grpc
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.6.1
// source: github.com/elojah/game_03/cmd/core/grpc/core.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as github_com_elojah_game_03_pkg_rtc_dto_rtc_pb from '../../../../../../github.com/elojah/game_03/pkg/rtc/dto/rtc_pb';
import * as github_com_elojah_game_03_pkg_pbtypes_empty_pb from '../../../../../../github.com/elojah/game_03/pkg/pbtypes/empty_pb';


export class CoreClient {
  client_: grpcWeb.AbstractClientBase;
  hostname_: string;
  credentials_: null | { [index: string]: string; };
  options_: null | { [index: string]: any; };

  constructor (hostname: string,
               credentials?: null | { [index: string]: string; },
               options?: null | { [index: string]: any; }) {
    if (!options) options = {};
    if (!credentials) credentials = {};
    options['format'] = 'binary';

    this.client_ = new grpcWeb.GrpcWebClientBase(options);
    this.hostname_ = hostname.replace(/\/+$/, '');
    this.credentials_ = credentials;
    this.options_ = options;
  }

  methodDescriptorConnect = new grpcWeb.MethodDescriptor(
    '/grpc.Core/Connect',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ConnectReq,
    github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.SDP,
    (request: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ConnectReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.SDP.deserializeBinary
  );

  connect(
    request: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ConnectReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.SDP>;

  connect(
    request: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ConnectReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.SDP) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.SDP>;

  connect(
    request: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ConnectReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.SDP) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.Core/Connect',
        request,
        metadata || {},
        this.methodDescriptorConnect,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.Core/Connect',
    request,
    metadata || {},
    this.methodDescriptorConnect);
  }

  methodDescriptorSendICE = new grpcWeb.MethodDescriptor(
    '/grpc.Core/SendICE',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate,
    github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty,
    (request: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty.deserializeBinary
  );

  sendICE(
    request: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty>;

  sendICE(
    request: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty>;

  sendICE(
    request: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.Core/SendICE',
        request,
        metadata || {},
        this.methodDescriptorSendICE,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.Core/SendICE',
    request,
    metadata || {},
    this.methodDescriptorSendICE);
  }

  methodDescriptorReceiveICE = new grpcWeb.MethodDescriptor(
    '/grpc.Core/ReceiveICE',
    grpcWeb.MethodType.SERVER_STREAMING,
    github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty,
    github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate,
    (request: github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate.deserializeBinary
  );

  receiveICE(
    request: github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty,
    metadata?: grpcWeb.Metadata): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate> {
    return this.client_.serverStreaming(
      this.hostname_ +
        '/grpc.Core/ReceiveICE',
      request,
      metadata || {},
      this.methodDescriptorReceiveICE);
  }

  methodDescriptorPing = new grpcWeb.MethodDescriptor(
    '/grpc.Core/Ping',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty,
    github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty,
    (request: github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty.deserializeBinary
  );

  ping(
    request: github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty>;

  ping(
    request: github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty>;

  ping(
    request: github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_pbtypes_empty_pb.Empty) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.Core/Ping',
        request,
        metadata || {},
        this.methodDescriptorPing,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.Core/Ping',
    request,
    metadata || {},
    this.methodDescriptorPing);
  }

}

