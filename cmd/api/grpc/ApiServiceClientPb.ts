/**
 * @fileoverview gRPC-Web generated client stub for grpc
 * @enhanceable
 * @public
 */

// Code generated by protoc-gen-grpc-web. DO NOT EDIT.
// versions:
// 	protoc-gen-grpc-web v1.4.2
// 	protoc              v3.6.1
// source: github.com/elojah/game_03/cmd/api/grpc/api.proto


/* eslint-disable */
// @ts-nocheck


import * as grpcWeb from 'grpc-web';

import * as github_com_elojah_game_03_pkg_entity_dto_entity_pb from '../../../../../../github.com/elojah/game_03/pkg/entity/dto/entity_pb';
import * as github_com_elojah_game_03_pkg_entity_dto_pc_pb from '../../../../../../github.com/elojah/game_03/pkg/entity/dto/pc_pb';
import * as github_com_elojah_game_03_pkg_room_dto_user_pb from '../../../../../../github.com/elojah/game_03/pkg/room/dto/user_pb';
import * as github_com_elojah_game_03_pkg_user_dto_session_pb from '../../../../../../github.com/elojah/game_03/pkg/user/dto/session_pb';
import * as github_com_elojah_game_03_pkg_ability_dto_ability_pb from '../../../../../../github.com/elojah/game_03/pkg/ability/dto/ability_pb';
import * as github_com_elojah_game_03_pkg_entity_dto_animation_pb from '../../../../../../github.com/elojah/game_03/pkg/entity/dto/animation_pb';
import * as github_com_elojah_game_03_pkg_room_dto_cell_pb from '../../../../../../github.com/elojah/game_03/pkg/room/dto/cell_pb';
import * as github_com_elojah_game_03_pkg_twitch_dto_follow_pb from '../../../../../../github.com/elojah/game_03/pkg/twitch/dto/follow_pb';
import * as github_com_elojah_game_03_pkg_room_dto_room_pb from '../../../../../../github.com/elojah/game_03/pkg/room/dto/room_pb';
import * as github_com_elojah_game_03_pkg_entity_dto_template_pb from '../../../../../../github.com/elojah/game_03/pkg/entity/dto/template_pb';
import * as github_com_elojah_game_03_pkg_room_dto_world_pb from '../../../../../../github.com/elojah/game_03/pkg/room/dto/world_pb';
import * as github_com_elojah_game_03_pkg_entity_entity_pb from '../../../../../../github.com/elojah/game_03/pkg/entity/entity_pb';
import * as github_com_elojah_game_03_pkg_entity_pc_pb from '../../../../../../github.com/elojah/game_03/pkg/entity/pc_pb';
import * as github_com_elojah_game_03_pkg_entity_pc_preferences_pb from '../../../../../../github.com/elojah/game_03/pkg/entity/pc_preferences_pb';
import * as github_com_elojah_game_03_pkg_pbtypes_empty_pb from '../../../../../../github.com/elojah/game_03/pkg/pbtypes/empty_pb';
import * as github_com_elojah_game_03_pkg_room_room_pb from '../../../../../../github.com/elojah/game_03/pkg/room/room_pb';
import * as github_com_elojah_game_03_pkg_room_user_pb from '../../../../../../github.com/elojah/game_03/pkg/room/user_pb';


export class APIClient {
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

  methodDescriptorCreateSession = new grpcWeb.MethodDescriptor(
    '/grpc.API/CreateSession',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_user_dto_session_pb.CreateSessionReq,
    github_com_elojah_game_03_pkg_user_dto_session_pb.CreateSessionResp,
    (request: github_com_elojah_game_03_pkg_user_dto_session_pb.CreateSessionReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_user_dto_session_pb.CreateSessionResp.deserializeBinary
  );

  createSession(
    request: github_com_elojah_game_03_pkg_user_dto_session_pb.CreateSessionReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_user_dto_session_pb.CreateSessionResp>;

  createSession(
    request: github_com_elojah_game_03_pkg_user_dto_session_pb.CreateSessionReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_user_dto_session_pb.CreateSessionResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_user_dto_session_pb.CreateSessionResp>;

  createSession(
    request: github_com_elojah_game_03_pkg_user_dto_session_pb.CreateSessionReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_user_dto_session_pb.CreateSessionResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/CreateSession',
        request,
        metadata || {},
        this.methodDescriptorCreateSession,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/CreateSession',
    request,
    metadata || {},
    this.methodDescriptorCreateSession);
  }

  methodDescriptorListEntity = new grpcWeb.MethodDescriptor(
    '/grpc.API/ListEntity',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityReq,
    github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityResp,
    (request: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityResp.deserializeBinary
  );

  listEntity(
    request: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityResp>;

  listEntity(
    request: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityResp>;

  listEntity(
    request: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/ListEntity',
        request,
        metadata || {},
        this.methodDescriptorListEntity,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/ListEntity',
    request,
    metadata || {},
    this.methodDescriptorListEntity);
  }

  methodDescriptorCreateEntity = new grpcWeb.MethodDescriptor(
    '/grpc.API/CreateEntity',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_entity_entity_pb.E,
    github_com_elojah_game_03_pkg_entity_entity_pb.E,
    (request: github_com_elojah_game_03_pkg_entity_entity_pb.E) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_entity_entity_pb.E.deserializeBinary
  );

  createEntity(
    request: github_com_elojah_game_03_pkg_entity_entity_pb.E,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_entity_entity_pb.E>;

  createEntity(
    request: github_com_elojah_game_03_pkg_entity_entity_pb.E,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_entity_pb.E) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_entity_entity_pb.E>;

  createEntity(
    request: github_com_elojah_game_03_pkg_entity_entity_pb.E,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_entity_pb.E) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/CreateEntity',
        request,
        metadata || {},
        this.methodDescriptorCreateEntity,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/CreateEntity',
    request,
    metadata || {},
    this.methodDescriptorCreateEntity);
  }

  methodDescriptorCreateEntityAbility = new grpcWeb.MethodDescriptor(
    '/grpc.API/CreateEntityAbility',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_entity_dto_entity_pb.CreateEntityAbilityReq,
    github_com_elojah_game_03_pkg_entity_dto_entity_pb.CreateEntityAbilityResp,
    (request: github_com_elojah_game_03_pkg_entity_dto_entity_pb.CreateEntityAbilityReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_entity_dto_entity_pb.CreateEntityAbilityResp.deserializeBinary
  );

  createEntityAbility(
    request: github_com_elojah_game_03_pkg_entity_dto_entity_pb.CreateEntityAbilityReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_entity_dto_entity_pb.CreateEntityAbilityResp>;

  createEntityAbility(
    request: github_com_elojah_game_03_pkg_entity_dto_entity_pb.CreateEntityAbilityReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_entity_pb.CreateEntityAbilityResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_entity_dto_entity_pb.CreateEntityAbilityResp>;

  createEntityAbility(
    request: github_com_elojah_game_03_pkg_entity_dto_entity_pb.CreateEntityAbilityReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_entity_pb.CreateEntityAbilityResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/CreateEntityAbility',
        request,
        metadata || {},
        this.methodDescriptorCreateEntityAbility,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/CreateEntityAbility',
    request,
    metadata || {},
    this.methodDescriptorCreateEntityAbility);
  }

  methodDescriptorListAbility = new grpcWeb.MethodDescriptor(
    '/grpc.API/ListAbility',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_ability_dto_ability_pb.ListAbilityReq,
    github_com_elojah_game_03_pkg_ability_dto_ability_pb.ListAbilityResp,
    (request: github_com_elojah_game_03_pkg_ability_dto_ability_pb.ListAbilityReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_ability_dto_ability_pb.ListAbilityResp.deserializeBinary
  );

  listAbility(
    request: github_com_elojah_game_03_pkg_ability_dto_ability_pb.ListAbilityReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_ability_dto_ability_pb.ListAbilityResp>;

  listAbility(
    request: github_com_elojah_game_03_pkg_ability_dto_ability_pb.ListAbilityReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_ability_dto_ability_pb.ListAbilityResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_ability_dto_ability_pb.ListAbilityResp>;

  listAbility(
    request: github_com_elojah_game_03_pkg_ability_dto_ability_pb.ListAbilityReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_ability_dto_ability_pb.ListAbilityResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/ListAbility',
        request,
        metadata || {},
        this.methodDescriptorListAbility,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/ListAbility',
    request,
    metadata || {},
    this.methodDescriptorListAbility);
  }

  methodDescriptorListAnimation = new grpcWeb.MethodDescriptor(
    '/grpc.API/ListAnimation',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationReq,
    github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationResp,
    (request: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationResp.deserializeBinary
  );

  listAnimation(
    request: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationResp>;

  listAnimation(
    request: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationResp>;

  listAnimation(
    request: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/ListAnimation',
        request,
        metadata || {},
        this.methodDescriptorListAnimation,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/ListAnimation',
    request,
    metadata || {},
    this.methodDescriptorListAnimation);
  }

  methodDescriptorCreatePC = new grpcWeb.MethodDescriptor(
    '/grpc.API/CreatePC',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_entity_dto_pc_pb.CreatePCReq,
    github_com_elojah_game_03_pkg_entity_pc_pb.PC,
    (request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.CreatePCReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_entity_pc_pb.PC.deserializeBinary
  );

  createPC(
    request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.CreatePCReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_entity_pc_pb.PC>;

  createPC(
    request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.CreatePCReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_pc_pb.PC) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_entity_pc_pb.PC>;

  createPC(
    request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.CreatePCReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_pc_pb.PC) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/CreatePC',
        request,
        metadata || {},
        this.methodDescriptorCreatePC,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/CreatePC',
    request,
    metadata || {},
    this.methodDescriptorCreatePC);
  }

  methodDescriptorListPC = new grpcWeb.MethodDescriptor(
    '/grpc.API/ListPC',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCReq,
    github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCResp,
    (request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCResp.deserializeBinary
  );

  listPC(
    request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCResp>;

  listPC(
    request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCResp>;

  listPC(
    request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/ListPC',
        request,
        metadata || {},
        this.methodDescriptorListPC,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/ListPC',
    request,
    metadata || {},
    this.methodDescriptorListPC);
  }

  methodDescriptorGetPC = new grpcWeb.MethodDescriptor(
    '/grpc.API/GetPC',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_entity_dto_pc_pb.GetPCReq,
    github_com_elojah_game_03_pkg_entity_dto_pc_pb.PC,
    (request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.GetPCReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_entity_dto_pc_pb.PC.deserializeBinary
  );

  getPC(
    request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.GetPCReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_entity_dto_pc_pb.PC>;

  getPC(
    request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.GetPCReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_pc_pb.PC) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_entity_dto_pc_pb.PC>;

  getPC(
    request: github_com_elojah_game_03_pkg_entity_dto_pc_pb.GetPCReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_pc_pb.PC) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/GetPC',
        request,
        metadata || {},
        this.methodDescriptorGetPC,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/GetPC',
    request,
    metadata || {},
    this.methodDescriptorGetPC);
  }

  methodDescriptorGetPCPreferences = new grpcWeb.MethodDescriptor(
    '/grpc.API/GetPCPreferences',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_entity_pc_pb.PC,
    github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences,
    (request: github_com_elojah_game_03_pkg_entity_pc_pb.PC) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences.deserializeBinary
  );

  getPCPreferences(
    request: github_com_elojah_game_03_pkg_entity_pc_pb.PC,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences>;

  getPCPreferences(
    request: github_com_elojah_game_03_pkg_entity_pc_pb.PC,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences>;

  getPCPreferences(
    request: github_com_elojah_game_03_pkg_entity_pc_pb.PC,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/GetPCPreferences',
        request,
        metadata || {},
        this.methodDescriptorGetPCPreferences,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/GetPCPreferences',
    request,
    metadata || {},
    this.methodDescriptorGetPCPreferences);
  }

  methodDescriptorUpdatePCPreferences = new grpcWeb.MethodDescriptor(
    '/grpc.API/UpdatePCPreferences',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences,
    github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences,
    (request: github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences.deserializeBinary
  );

  updatePCPreferences(
    request: github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences>;

  updatePCPreferences(
    request: github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences>;

  updatePCPreferences(
    request: github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_pc_preferences_pb.PCPreferences) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/UpdatePCPreferences',
        request,
        metadata || {},
        this.methodDescriptorUpdatePCPreferences,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/UpdatePCPreferences',
    request,
    metadata || {},
    this.methodDescriptorUpdatePCPreferences);
  }

  methodDescriptorListTemplate = new grpcWeb.MethodDescriptor(
    '/grpc.API/ListTemplate',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateReq,
    github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateResp,
    (request: github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateResp.deserializeBinary
  );

  listTemplate(
    request: github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateResp>;

  listTemplate(
    request: github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateResp>;

  listTemplate(
    request: github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/ListTemplate',
        request,
        metadata || {},
        this.methodDescriptorListTemplate,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/ListTemplate',
    request,
    metadata || {},
    this.methodDescriptorListTemplate);
  }

  methodDescriptorCreateRoom = new grpcWeb.MethodDescriptor(
    '/grpc.API/CreateRoom',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_room_room_pb.R,
    github_com_elojah_game_03_pkg_room_room_pb.R,
    (request: github_com_elojah_game_03_pkg_room_room_pb.R) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_room_room_pb.R.deserializeBinary
  );

  createRoom(
    request: github_com_elojah_game_03_pkg_room_room_pb.R,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_room_room_pb.R>;

  createRoom(
    request: github_com_elojah_game_03_pkg_room_room_pb.R,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_room_pb.R) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_room_room_pb.R>;

  createRoom(
    request: github_com_elojah_game_03_pkg_room_room_pb.R,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_room_pb.R) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/CreateRoom',
        request,
        metadata || {},
        this.methodDescriptorCreateRoom,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/CreateRoom',
    request,
    metadata || {},
    this.methodDescriptorCreateRoom);
  }

  methodDescriptorListRoom = new grpcWeb.MethodDescriptor(
    '/grpc.API/ListRoom',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq,
    github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp,
    (request: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp.deserializeBinary
  );

  listRoom(
    request: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp>;

  listRoom(
    request: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp>;

  listRoom(
    request: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/ListRoom',
        request,
        metadata || {},
        this.methodDescriptorListRoom,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/ListRoom',
    request,
    metadata || {},
    this.methodDescriptorListRoom);
  }

  methodDescriptorListRoomPublic = new grpcWeb.MethodDescriptor(
    '/grpc.API/ListRoomPublic',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq,
    github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp,
    (request: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp.deserializeBinary
  );

  listRoomPublic(
    request: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp>;

  listRoomPublic(
    request: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp>;

  listRoomPublic(
    request: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/ListRoomPublic',
        request,
        metadata || {},
        this.methodDescriptorListRoomPublic,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/ListRoomPublic',
    request,
    metadata || {},
    this.methodDescriptorListRoomPublic);
  }

  methodDescriptorCreateRoomUser = new grpcWeb.MethodDescriptor(
    '/grpc.API/CreateRoomUser',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_room_dto_user_pb.CreateRoomUserReq,
    github_com_elojah_game_03_pkg_room_user_pb.User,
    (request: github_com_elojah_game_03_pkg_room_dto_user_pb.CreateRoomUserReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_room_user_pb.User.deserializeBinary
  );

  createRoomUser(
    request: github_com_elojah_game_03_pkg_room_dto_user_pb.CreateRoomUserReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_room_user_pb.User>;

  createRoomUser(
    request: github_com_elojah_game_03_pkg_room_dto_user_pb.CreateRoomUserReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_user_pb.User) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_room_user_pb.User>;

  createRoomUser(
    request: github_com_elojah_game_03_pkg_room_dto_user_pb.CreateRoomUserReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_user_pb.User) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/CreateRoomUser',
        request,
        metadata || {},
        this.methodDescriptorCreateRoomUser,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/CreateRoomUser',
    request,
    metadata || {},
    this.methodDescriptorCreateRoomUser);
  }

  methodDescriptorListCell = new grpcWeb.MethodDescriptor(
    '/grpc.API/ListCell',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellReq,
    github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellResp,
    (request: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellResp.deserializeBinary
  );

  listCell(
    request: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellResp>;

  listCell(
    request: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellResp>;

  listCell(
    request: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/ListCell',
        request,
        metadata || {},
        this.methodDescriptorListCell,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/ListCell',
    request,
    metadata || {},
    this.methodDescriptorListCell);
  }

  methodDescriptorListWorld = new grpcWeb.MethodDescriptor(
    '/grpc.API/ListWorld',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldReq,
    github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldResp,
    (request: github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldResp.deserializeBinary
  );

  listWorld(
    request: github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldResp>;

  listWorld(
    request: github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldResp>;

  listWorld(
    request: github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/ListWorld',
        request,
        metadata || {},
        this.methodDescriptorListWorld,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/ListWorld',
    request,
    metadata || {},
    this.methodDescriptorListWorld);
  }

  methodDescriptorListFollow = new grpcWeb.MethodDescriptor(
    '/grpc.API/ListFollow',
    grpcWeb.MethodType.UNARY,
    github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq,
    github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp,
    (request: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq) => {
      return request.serializeBinary();
    },
    github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp.deserializeBinary
  );

  listFollow(
    request: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq,
    metadata: grpcWeb.Metadata | null): Promise<github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp>;

  listFollow(
    request: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq,
    metadata: grpcWeb.Metadata | null,
    callback: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp) => void): grpcWeb.ClientReadableStream<github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp>;

  listFollow(
    request: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq,
    metadata: grpcWeb.Metadata | null,
    callback?: (err: grpcWeb.RpcError,
               response: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp) => void) {
    if (callback !== undefined) {
      return this.client_.rpcCall(
        this.hostname_ +
          '/grpc.API/ListFollow',
        request,
        metadata || {},
        this.methodDescriptorListFollow,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/ListFollow',
    request,
    metadata || {},
    this.methodDescriptorListFollow);
  }

  methodDescriptorPing = new grpcWeb.MethodDescriptor(
    '/grpc.API/Ping',
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
          '/grpc.API/Ping',
        request,
        metadata || {},
        this.methodDescriptorPing,
        callback);
    }
    return this.client_.unaryCall(
    this.hostname_ +
      '/grpc.API/Ping',
    request,
    metadata || {},
    this.methodDescriptorPing);
  }

}
