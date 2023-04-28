// package: grpc
// file: github.com/elojah/game_03/cmd/core/grpc/core.proto

var github_com_elojah_game_03_cmd_core_grpc_core_pb = require("../../../../../../github.com/elojah/game_03/cmd/core/grpc/core_pb");
var google_protobuf_empty_pb = require("google-protobuf/google/protobuf/empty_pb");
var github_com_elojah_game_03_pkg_rtc_dto_rtc_pb = require("../../../../../../github.com/elojah/game_03/pkg/rtc/dto/rtc_pb");
var grpc = require("@improbable-eng/grpc-web").grpc;

var Core = (function () {
  function Core() {}
  Core.serviceName = "grpc.Core";
  return Core;
}());

Core.Connect = {
  methodName: "Connect",
  service: Core,
  requestStream: false,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ConnectReq,
  responseType: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.SDP
};

Core.SendICE = {
  methodName: "SendICE",
  service: Core,
  requestStream: true,
  responseStream: false,
  requestType: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate,
  responseType: google_protobuf_empty_pb.Empty
};

Core.ReceiveICE = {
  methodName: "ReceiveICE",
  service: Core,
  requestStream: false,
  responseStream: true,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: github_com_elojah_game_03_pkg_rtc_dto_rtc_pb.ICECandidate
};

Core.Ping = {
  methodName: "Ping",
  service: Core,
  requestStream: false,
  responseStream: false,
  requestType: google_protobuf_empty_pb.Empty,
  responseType: google_protobuf_empty_pb.Empty
};

exports.Core = Core;

function CoreClient(serviceHost, options) {
  this.serviceHost = serviceHost;
  this.options = options || {};
}

CoreClient.prototype.connect = function connect(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Core.Connect, {
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

CoreClient.prototype.sendICE = function sendICE(metadata) {
  var listeners = {
    end: [],
    status: []
  };
  var client = grpc.client(Core.SendICE, {
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

CoreClient.prototype.receiveICE = function receiveICE(requestMessage, metadata) {
  var listeners = {
    data: [],
    end: [],
    status: []
  };
  var client = grpc.invoke(Core.ReceiveICE, {
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

CoreClient.prototype.ping = function ping(requestMessage, metadata, callback) {
  if (arguments.length === 2) {
    callback = arguments[1];
  }
  var client = grpc.unary(Core.Ping, {
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

exports.CoreClient = CoreClient;

