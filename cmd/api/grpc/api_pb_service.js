// package: grpc
// file: github.com/elojah/game_03/cmd/api/grpc/api.proto

var github_com_elojah_game_03_cmd_api_grpc_api_pb = require("../../../../../../github.com/elojah/game_03/cmd/api/grpc/api_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var github_com_elojah_game_03_pkg_entity_pc_pb = require("../../../../../../github.com/elojah/game_03/pkg/entity/pc_pb");
var github_com_elojah_game_03_pkg_entity_dto_pc_pb = require("../../../../../../github.com/elojah/game_03/pkg/entity/dto/pc_pb");
var github_com_elojah_game_03_pkg_room_room_pb = require("../../../../../../github.com/elojah/game_03/pkg/room/room_pb");
var github_com_elojah_game_03_pkg_room_dto_room_pb = require("../../../../../../github.com/elojah/game_03/pkg/room/dto/room_pb");
var github_com_elojah_game_03_pkg_twitch_dto_follow_pb = require("../../../../../../github.com/elojah/game_03/pkg/twitch/dto/follow_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var API = (function () {
  function API() {}
  API.serviceName = "grpc.API";
  return API;
}());

API.CreatePC = {
  methodName: "CreatePC",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_entity_pc_pb.PC,
  responseType: github_com_elojah_game_03_pkg_entity_pc_pb.PC
};

API.ListPC = {
  methodName: "ListPC",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCReq,
  responseType: github_com_elojah_game_03_pkg_entity_dto_pc_pb.ListPCResp
};

API.CreateRoom = {
  methodName: "CreateRoom",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_room_room_pb.R,
  responseType: github_com_elojah_game_03_pkg_room_room_pb.R
};

API.ListRoom = {
  methodName: "ListRoom",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomReq,
  responseType: github_com_elojah_game_03_pkg_room_dto_room_pb.ListRoomResp
};

API.ListFollow = {
  methodName: "ListFollow",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowReq,
  responseType: github_com_elojah_game_03_pkg_twitch_dto_follow_pb.ListFollowResp
};

API.Ping = {
  methodName: "Ping",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: google_protobuf_empty_pb.Empty
};

exports.API = API;

function APIClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

APIClient.prototype.createPC = function createPC(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.CreatePC, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

APIClient.prototype.listPC = function listPC(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.ListPC, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

APIClient.prototype.createRoom = function createRoom(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.CreateRoom, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

APIClient.prototype.listRoom = function listRoom(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.ListRoom, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

APIClient.prototype.listFollow = function listFollow(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.ListFollow, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

APIClient.prototype.ping = function ping(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.Ping, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onEnd: function (response) {
      if (callback) {
        if (response.status !== grpc.Code.OK) {
          var err = new Error(response.statusMessage);
          err.code = response.status;
          err.metadata = response.trailers;
          callback(err, null);
        } else {
          callback(null, response.message);
        }
      }
    }
  });
  return {
    cancel: function () {
      callback = null;
      client.close();
    }
  };
};

exports.APIClient = APIClient;

