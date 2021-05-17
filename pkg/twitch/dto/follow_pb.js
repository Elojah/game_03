// source: github.com/elojah/game_03/pkg/twitch/dto/follow.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var github_com_gogo_protobuf_gogoproto_gogo_pb = require('../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb.js');
goog.object.extend(proto, github_com_gogo_protobuf_gogoproto_gogo_pb);
var github_com_elojah_game_03_pkg_twitch_follow_pb = require('../../../../../../github.com/elojah/game_03/pkg/twitch/follow_pb.js');
goog.object.extend(proto, github_com_elojah_game_03_pkg_twitch_follow_pb);
goog.exportSymbol('proto.dto.ListFollowReq', null, global);
goog.exportSymbol('proto.dto.ListFollowResp', null, global);
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.dto.ListFollowReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dto.ListFollowReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dto.ListFollowReq.displayName = 'proto.dto.ListFollowReq';
}
/**
 * Generated by JsPbCodeGenerator.
 * @param {Array=} opt_data Optional initial data array, typically from a
 * server response, or constructed directly in Javascript. The array is used
 * in place and becomes part of the constructed object. It is not cloned.
 * If no data is provided, the constructed object will be empty, but still
 * valid.
 * @extends {jspb.Message}
 * @constructor
 */
proto.dto.ListFollowResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.dto.ListFollowResp.repeatedFields_, null);
};
goog.inherits(proto.dto.ListFollowResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dto.ListFollowResp.displayName = 'proto.dto.ListFollowResp';
}



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.dto.ListFollowReq.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.ListFollowReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.ListFollowReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListFollowReq.toObject = function(includeInstance, msg) {
  var f, obj = {
    after: jspb.Message.getFieldWithDefault(msg, 1, ""),
    first: jspb.Message.getFieldWithDefault(msg, 2, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.dto.ListFollowReq}
 */
proto.dto.ListFollowReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.ListFollowReq;
  return proto.dto.ListFollowReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.ListFollowReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.ListFollowReq}
 */
proto.dto.ListFollowReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setAfter(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setFirst(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.dto.ListFollowReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.ListFollowReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.ListFollowReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListFollowReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAfter();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getFirst();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * optional string After = 1;
 * @return {string}
 */
proto.dto.ListFollowReq.prototype.getAfter = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.dto.ListFollowReq} returns this
 */
proto.dto.ListFollowReq.prototype.setAfter = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string First = 2;
 * @return {string}
 */
proto.dto.ListFollowReq.prototype.getFirst = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * @param {string} value
 * @return {!proto.dto.ListFollowReq} returns this
 */
proto.dto.ListFollowReq.prototype.setFirst = function(value) {
  return jspb.Message.setProto3StringField(this, 2, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.dto.ListFollowResp.repeatedFields_ = [1];



if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * Optional fields that are not set will be set to undefined.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     net/proto2/compiler/js/internal/generator.cc#kKeyword.
 * @param {boolean=} opt_includeInstance Deprecated. whether to include the
 *     JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @return {!Object}
 */
proto.dto.ListFollowResp.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.ListFollowResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.ListFollowResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListFollowResp.toObject = function(includeInstance, msg) {
  var f, obj = {
    followsList: jspb.Message.toObjectList(msg.getFollowsList(),
    github_com_elojah_game_03_pkg_twitch_follow_pb.Follow.toObject, includeInstance),
    total: jspb.Message.getFieldWithDefault(msg, 2, 0),
    cursor: jspb.Message.getFieldWithDefault(msg, 3, "")
  };

  if (includeInstance) {
    obj.$jspbMessageInstance = msg;
  }
  return obj;
};
}


/**
 * Deserializes binary data (in protobuf wire format).
 * @param {jspb.ByteSource} bytes The bytes to deserialize.
 * @return {!proto.dto.ListFollowResp}
 */
proto.dto.ListFollowResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.ListFollowResp;
  return proto.dto.ListFollowResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.ListFollowResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.ListFollowResp}
 */
proto.dto.ListFollowResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_elojah_game_03_pkg_twitch_follow_pb.Follow;
      reader.readMessage(value,github_com_elojah_game_03_pkg_twitch_follow_pb.Follow.deserializeBinaryFromReader);
      msg.addFollows(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setTotal(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setCursor(value);
      break;
    default:
      reader.skipField();
      break;
    }
  }
  return msg;
};


/**
 * Serializes the message to binary data (in protobuf wire format).
 * @return {!Uint8Array}
 */
proto.dto.ListFollowResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.ListFollowResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.ListFollowResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListFollowResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFollowsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      github_com_elojah_game_03_pkg_twitch_follow_pb.Follow.serializeBinaryToWriter
    );
  }
  f = message.getTotal();
  if (f !== 0) {
    writer.writeUint64(
      2,
      f
    );
  }
  f = message.getCursor();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
};


/**
 * repeated twitch.Follow Follows = 1;
 * @return {!Array<!proto.twitch.Follow>}
 */
proto.dto.ListFollowResp.prototype.getFollowsList = function() {
  return /** @type{!Array<!proto.twitch.Follow>} */ (
    jspb.Message.getRepeatedWrapperField(this, github_com_elojah_game_03_pkg_twitch_follow_pb.Follow, 1));
};


/**
 * @param {!Array<!proto.twitch.Follow>} value
 * @return {!proto.dto.ListFollowResp} returns this
*/
proto.dto.ListFollowResp.prototype.setFollowsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.twitch.Follow=} opt_value
 * @param {number=} opt_index
 * @return {!proto.twitch.Follow}
 */
proto.dto.ListFollowResp.prototype.addFollows = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.twitch.Follow, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.dto.ListFollowResp} returns this
 */
proto.dto.ListFollowResp.prototype.clearFollowsList = function() {
  return this.setFollowsList([]);
};


/**
 * optional uint64 Total = 2;
 * @return {number}
 */
proto.dto.ListFollowResp.prototype.getTotal = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.dto.ListFollowResp} returns this
 */
proto.dto.ListFollowResp.prototype.setTotal = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional string Cursor = 3;
 * @return {string}
 */
proto.dto.ListFollowResp.prototype.getCursor = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * @param {string} value
 * @return {!proto.dto.ListFollowResp} returns this
 */
proto.dto.ListFollowResp.prototype.setCursor = function(value) {
  return jspb.Message.setProto3StringField(this, 3, value);
};


goog.object.extend(exports, proto.dto);
