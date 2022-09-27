// package: grpc
// file: github.com/elojah/game_03/cmd/admin/grpc/admin.proto

var github_com_elojah_game_03_cmd_admin_grpc_admin_pb = require("../../../../../../github.com/elojah/game_03/cmd/admin/grpc/admin_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var google_protobuf_wrappers_pb = require("google-protobuf/google/protobuf/wrappers_pb");
var github_com_elojah_game_03_pkg_tile_dto_set_pb = require("../../../../../../github.com/elojah/game_03/pkg/tile/dto/set_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Admin = (function () {
  function Admin() {}
  Admin.serviceName = "grpc.Admin";
  return Admin;
}());

Admin.MigrateUp = {
  methodName: "MigrateUp",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_wrappers_pb.StringValue,
  responseType: google_protobuf_empty_pb.Empty
};

Admin.CreateTilemap = {
  methodName: "CreateTilemap",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: google_protobuf_wrappers_pb.StringValue
};

Admin.CreateTileset = {
  methodName: "CreateTileset",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_tile_dto_set_pb.CreateTilesetReq,
  responseType: github_com_elojah_game_03_pkg_tile_dto_set_pb.CreateTilesetResp
};

Admin.CreateWorld = {
  methodName: "CreateWorld",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: google_protobuf_wrappers_pb.StringValue
};

Admin.Ping = {
  methodName: "Ping",
  service: Admin,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: google_protobuf_empty_pb.Empty
};

exports.Admin = Admin;

function AdminClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

AdminClient.prototype.migrateUp = function migrateUp(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.MigrateUp, {
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

AdminClient.prototype.createTilemap = function createTilemap(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.CreateTilemap, {
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

AdminClient.prototype.createTileset = function createTileset(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.CreateTileset, {
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

AdminClient.prototype.createWorld = function createWorld(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.CreateWorld, {
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

AdminClient.prototype.ping = function ping(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Admin.Ping, {
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

exports.AdminClient = AdminClient;

