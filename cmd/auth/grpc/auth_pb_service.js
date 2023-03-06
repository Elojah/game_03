// package: grpc
// file: github.com/elojah/game_03/cmd/auth/grpc/auth.proto

var github_com_elojah_game_03_cmd_auth_grpc_auth_pb = require("../../../../../../github.com/elojah/game_03/cmd/auth/grpc/auth_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var google_protobuf_wrappers_pb = require("google-protobuf/google/protobuf/wrappers_pb");
var github_com_elojah_game_03_pkg_user_dto_user_pb = require("../../../../../../github.com/elojah/game_03/pkg/user/dto/user_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Auth = (function () {
  function Auth() {}
  Auth.serviceName = "grpc.Auth";
  return Auth;
}());

Auth.SigninTwitch = {
  methodName: "SigninTwitch",
  service: Auth,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_wrappers_pb.StringValue,
  responseType: github_com_elojah_game_03_pkg_user_dto_user_pb.SigninResp
};

Auth.SigninGoogle = {
  methodName: "SigninGoogle",
  service: Auth,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_wrappers_pb.StringValue,
  responseType: github_com_elojah_game_03_pkg_user_dto_user_pb.SigninResp
};

Auth.RefreshToken = {
  methodName: "RefreshToken",
  service: Auth,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_wrappers_pb.StringValue,
  responseType: github_com_elojah_game_03_pkg_user_dto_user_pb.SigninResp
};

Auth.Ping = {
  methodName: "Ping",
  service: Auth,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: google_protobuf_empty_pb.Empty
};

exports.Auth = Auth;

function AuthClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

AuthClient.prototype.signinTwitch = function signinTwitch(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Auth.SigninTwitch, {
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

AuthClient.prototype.signinGoogle = function signinGoogle(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Auth.SigninGoogle, {
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

AuthClient.prototype.refreshToken = function refreshToken(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Auth.RefreshToken, {
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

AuthClient.prototype.ping = function ping(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Auth.Ping, {
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

exports.AuthClient = AuthClient;

