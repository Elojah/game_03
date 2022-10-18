// package: grpc
// file: github.com/elojah/game_03/cmd/api/grpc/api.proto

var github_com_elojah_game_03_cmd_api_grpc_api_pb = require("../../../../../../github.com/elojah/game_03/cmd/api/grpc/api_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var github_com_elojah_game_03_pkg_entity_entity_pb = require("../../../../../../github.com/elojah/game_03/pkg/entity/entity_pb");
var github_com_elojah_game_03_pkg_entity_pc_pb = require("../../../../../../github.com/elojah/game_03/pkg/entity/pc_pb");
var github_com_elojah_game_03_pkg_entity_dto_entity_pb = require("../../../../../../github.com/elojah/game_03/pkg/entity/dto/entity_pb");
var github_com_elojah_game_03_pkg_entity_dto_animation_pb = require("../../../../../../github.com/elojah/game_03/pkg/entity/dto/animation_pb");
var github_com_elojah_game_03_pkg_entity_dto_pc_pb = require("../../../../../../github.com/elojah/game_03/pkg/entity/dto/pc_pb");
var github_com_elojah_game_03_pkg_entity_dto_template_pb = require("../../../../../../github.com/elojah/game_03/pkg/entity/dto/template_pb");
var github_com_elojah_game_03_pkg_room_room_pb = require("../../../../../../github.com/elojah/game_03/pkg/room/room_pb");
var github_com_elojah_game_03_pkg_room_dto_cell_pb = require("../../../../../../github.com/elojah/game_03/pkg/room/dto/cell_pb");
var github_com_elojah_game_03_pkg_room_dto_room_pb = require("../../../../../../github.com/elojah/game_03/pkg/room/dto/room_pb");
var github_com_elojah_game_03_pkg_room_dto_world_pb = require("../../../../../../github.com/elojah/game_03/pkg/room/dto/world_pb");
var github_com_elojah_game_03_pkg_twitch_dto_follow_pb = require("../../../../../../github.com/elojah/game_03/pkg/twitch/dto/follow_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var API = (function () {
  function API() {}
  API.serviceName = "grpc.API";
  return API;
}());

API.ConnectPC = {
  methodName: "ConnectPC",
  service: API,
  requestStream: false,
  responseStream: true,
  requestType: github_com_elojah_game_03_pkg_entity_pc_pb.PC,
  responseType: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityResp
};

API.UpdateEntity = {
  methodName: "UpdateEntity",
  service: API,
  requestStream: true,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_entity_entity_pb.E,
  responseType: google_protobuf_empty_pb.Empty
};

API.ListEntity = {
  methodName: "ListEntity",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityReq,
  responseType: github_com_elojah_game_03_pkg_entity_dto_entity_pb.ListEntityResp
};

API.ListAnimation = {
  methodName: "ListAnimation",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationReq,
  responseType: github_com_elojah_game_03_pkg_entity_dto_animation_pb.ListAnimationResp
};

API.CreatePC = {
  methodName: "CreatePC",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_entity_dto_pc_pb.CreatePCReq,
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

API.ListTemplate = {
  methodName: "ListTemplate",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateReq,
  responseType: github_com_elojah_game_03_pkg_entity_dto_template_pb.ListTemplateResp
};

API.CreateEntity = {
  methodName: "CreateEntity",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_entity_entity_pb.E,
  responseType: github_com_elojah_game_03_pkg_entity_entity_pb.E
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

API.ListCell = {
  methodName: "ListCell",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellReq,
  responseType: github_com_elojah_game_03_pkg_room_dto_cell_pb.ListCellResp
};

API.ListWorld = {
  methodName: "ListWorld",
  service: API,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldReq,
  responseType: github_com_elojah_game_03_pkg_room_dto_world_pb.ListWorldResp
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

APIClient.prototype.connectPC = function connectPC(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(API.ConnectPC, {
    request: requestMessage,
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport,
    debug: this.options.debug,
    onMessage: function (responseMessage) {
      listeners.data.forEach(function (handler) {
        handler(responseMessage);
      });
    },
    onEnd: function (status, statusMessage, trailers) {
      listeners.status.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners.end.forEach(function (handler) {
        handler({ code: status, details: statusMessage, metadata: trailers });
      });
      listeners = null;
    }
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

APIClient.prototype.updateEntity = function updateEntity(metadata) {
  var listeners = {
    end: [],
    status: []
  };
  var client = grpc.client(API.UpdateEntity, {
    host: this.serviceHost,
    metadata: metadata,
    transport: this.options.transport
  });
  client.onEnd(function (status, statusMessage, trailers) {
    listeners.status.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners.end.forEach(function (handler) {
      handler({ code: status, details: statusMessage, metadata: trailers });
    });
    listeners = null;
  });
  return {
    on: function (type, handler) {
      listeners[type].push(handler);
      return this;
    },
    write: function (requestMessage) {
      if (!client.started) {
        client.start(metadata);
      }
      client.send(requestMessage);
      return this;
    },
    end: function () {
      client.finishSend();
    },
    cancel: function () {
      listeners = null;
      client.close();
    }
  };
};

APIClient.prototype.listEntity = function listEntity(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.ListEntity, {
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

APIClient.prototype.listAnimation = function listAnimation(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.ListAnimation, {
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

APIClient.prototype.listTemplate = function listTemplate(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.ListTemplate, {
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

APIClient.prototype.createEntity = function createEntity(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.CreateEntity, {
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

APIClient.prototype.listCell = function listCell(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.ListCell, {
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

APIClient.prototype.listWorld = function listWorld(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(API.ListWorld, {
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

