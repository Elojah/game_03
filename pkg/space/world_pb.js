// source: github.com/elojah/game_03/pkg/space/world.proto
/**
 * @fileoverview
 * @enhanceable
 * @suppress {messageConventions} JS Compiler reports an error if a variable or
 *     field starts with 'MSG_' and isn't a translatable message.
 * @public
 */
// GENERATED CODE -- DO NOT EDIT!
/* eslint-disable */
// @ts-nocheck

var jspb = require('google-protobuf');
var goog = jspb;
var global = Function('return this')();

var github_com_gogo_protobuf_gogoproto_gogo_pb = require('../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb.js');
goog.object.extend(proto, github_com_gogo_protobuf_gogoproto_gogo_pb);
var github_com_elojah_game_03_pkg_geometry_geometry_pb = require('../../../../../github.com/elojah/game_03/pkg/geometry/geometry_pb.js');
goog.object.extend(proto, github_com_elojah_game_03_pkg_geometry_geometry_pb);
goog.exportSymbol('proto.space.Section', null, global);
goog.exportSymbol('proto.space.World', null, global);
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
proto.space.World = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.space.World, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.space.World.displayName = 'proto.space.World';
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
proto.space.Section = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.space.Section, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.space.Section.displayName = 'proto.space.Section';
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
proto.space.World.prototype.toObject = function(opt_includeInstance) {
  return proto.space.World.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.space.World} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.space.World.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: msg.getId_asB64(),
    ownerid: msg.getOwnerid_asB64(),
    nsection: (f = msg.getNsection()) && github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.toObject(includeInstance, f),
    dimsection: (f = msg.getDimsection()) && github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.toObject(includeInstance, f)
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
 * @return {!proto.space.World}
 */
proto.space.World.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.space.World;
  return proto.space.World.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.space.World} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.space.World}
 */
proto.space.World.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setId(value);
      break;
    case 2:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setOwnerid(value);
      break;
    case 3:
      var value = new github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2;
      reader.readMessage(value,github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.deserializeBinaryFromReader);
      msg.setNsection(value);
      break;
    case 4:
      var value = new github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2;
      reader.readMessage(value,github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.deserializeBinaryFromReader);
      msg.setDimsection(value);
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
proto.space.World.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.space.World.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.space.World} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.space.World.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getOwnerid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      2,
      f
    );
  }
  f = message.getNsection();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.serializeBinaryToWriter
    );
  }
  f = message.getDimsection();
  if (f != null) {
    writer.writeMessage(
      4,
      f,
      github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.serializeBinaryToWriter
    );
  }
};


/**
 * optional bytes ID = 1;
 * @return {!(string|Uint8Array)}
 */
proto.space.World.prototype.getId = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes ID = 1;
 * This is a type-conversion wrapper around `getId()`
 * @return {string}
 */
proto.space.World.prototype.getId_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getId()));
};


/**
 * optional bytes ID = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getId()`
 * @return {!Uint8Array}
 */
proto.space.World.prototype.getId_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getId()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.space.World} returns this
 */
proto.space.World.prototype.setId = function(value) {
  return jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional bytes OwnerID = 2;
 * @return {!(string|Uint8Array)}
 */
proto.space.World.prototype.getOwnerid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes OwnerID = 2;
 * This is a type-conversion wrapper around `getOwnerid()`
 * @return {string}
 */
proto.space.World.prototype.getOwnerid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getOwnerid()));
};


/**
 * optional bytes OwnerID = 2;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getOwnerid()`
 * @return {!Uint8Array}
 */
proto.space.World.prototype.getOwnerid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getOwnerid()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.space.World} returns this
 */
proto.space.World.prototype.setOwnerid = function(value) {
  return jspb.Message.setProto3BytesField(this, 2, value);
};


/**
 * optional geometry.Vec2 NSection = 3;
 * @return {?proto.geometry.Vec2}
 */
proto.space.World.prototype.getNsection = function() {
  return /** @type{?proto.geometry.Vec2} */ (
    jspb.Message.getWrapperField(this, github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2, 3));
};


/**
 * @param {?proto.geometry.Vec2|undefined} value
 * @return {!proto.space.World} returns this
*/
proto.space.World.prototype.setNsection = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.space.World} returns this
 */
proto.space.World.prototype.clearNsection = function() {
  return this.setNsection(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.space.World.prototype.hasNsection = function() {
  return jspb.Message.getField(this, 3) != null;
};


/**
 * optional geometry.Vec2 DimSection = 4;
 * @return {?proto.geometry.Vec2}
 */
proto.space.World.prototype.getDimsection = function() {
  return /** @type{?proto.geometry.Vec2} */ (
    jspb.Message.getWrapperField(this, github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2, 4));
};


/**
 * @param {?proto.geometry.Vec2|undefined} value
 * @return {!proto.space.World} returns this
*/
proto.space.World.prototype.setDimsection = function(value) {
  return jspb.Message.setWrapperField(this, 4, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.space.World} returns this
 */
proto.space.World.prototype.clearDimsection = function() {
  return this.setDimsection(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.space.World.prototype.hasDimsection = function() {
  return jspb.Message.getField(this, 4) != null;
};





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
proto.space.Section.prototype.toObject = function(opt_includeInstance) {
  return proto.space.Section.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.space.Section} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.space.Section.toObject = function(includeInstance, msg) {
  var f, obj = {
    worldid: msg.getWorldid_asB64(),
    coord: (f = msg.getCoord()) && github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.toObject(includeInstance, f)
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
 * @return {!proto.space.Section}
 */
proto.space.Section.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.space.Section;
  return proto.space.Section.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.space.Section} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.space.Section}
 */
proto.space.Section.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setWorldid(value);
      break;
    case 2:
      var value = new github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2;
      reader.readMessage(value,github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.deserializeBinaryFromReader);
      msg.setCoord(value);
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
proto.space.Section.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.space.Section.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.space.Section} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.space.Section.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getWorldid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getCoord();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2.serializeBinaryToWriter
    );
  }
};


/**
 * optional bytes WorldID = 1;
 * @return {!(string|Uint8Array)}
 */
proto.space.Section.prototype.getWorldid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes WorldID = 1;
 * This is a type-conversion wrapper around `getWorldid()`
 * @return {string}
 */
proto.space.Section.prototype.getWorldid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getWorldid()));
};


/**
 * optional bytes WorldID = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getWorldid()`
 * @return {!Uint8Array}
 */
proto.space.Section.prototype.getWorldid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getWorldid()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.space.Section} returns this
 */
proto.space.Section.prototype.setWorldid = function(value) {
  return jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional geometry.Vec2 Coord = 2;
 * @return {?proto.geometry.Vec2}
 */
proto.space.Section.prototype.getCoord = function() {
  return /** @type{?proto.geometry.Vec2} */ (
    jspb.Message.getWrapperField(this, github_com_elojah_game_03_pkg_geometry_geometry_pb.Vec2, 2));
};


/**
 * @param {?proto.geometry.Vec2|undefined} value
 * @return {!proto.space.Section} returns this
*/
proto.space.Section.prototype.setCoord = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.space.Section} returns this
 */
proto.space.Section.prototype.clearCoord = function() {
  return this.setCoord(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.space.Section.prototype.hasCoord = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.space);
