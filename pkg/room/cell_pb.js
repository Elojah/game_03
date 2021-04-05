// source: github.com/elojah/game_03/pkg/room/cell.proto
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
goog.exportSymbol('proto.room.Cell', null, global);
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
proto.room.Cell = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.room.Cell, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.room.Cell.displayName = 'proto.room.Cell';
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
proto.room.Cell.prototype.toObject = function(opt_includeInstance) {
  return proto.room.Cell.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.room.Cell} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.room.Cell.toObject = function(includeInstance, msg) {
  var f, obj = {
    worldid: msg.getWorldid_asB64(),
    x: jspb.Message.getFieldWithDefault(msg, 2, 0),
    y: jspb.Message.getFieldWithDefault(msg, 3, 0),
    field: msg.getField_asB64()
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
 * @return {!proto.room.Cell}
 */
proto.room.Cell.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.room.Cell;
  return proto.room.Cell.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.room.Cell} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.room.Cell}
 */
proto.room.Cell.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = /** @type {number} */ (reader.readInt64());
      msg.setX(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setY(value);
      break;
    case 4:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setField(value);
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
proto.room.Cell.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.room.Cell.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.room.Cell} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.room.Cell.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getWorldid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getX();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getY();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getField_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      4,
      f
    );
  }
};


/**
 * optional bytes WorldID = 1;
 * @return {!(string|Uint8Array)}
 */
proto.room.Cell.prototype.getWorldid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes WorldID = 1;
 * This is a type-conversion wrapper around `getWorldid()`
 * @return {string}
 */
proto.room.Cell.prototype.getWorldid_asB64 = function() {
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
proto.room.Cell.prototype.getWorldid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getWorldid()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.room.Cell} returns this
 */
proto.room.Cell.prototype.setWorldid = function(value) {
  return jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional int64 X = 2;
 * @return {number}
 */
proto.room.Cell.prototype.getX = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/**
 * @param {number} value
 * @return {!proto.room.Cell} returns this
 */
proto.room.Cell.prototype.setX = function(value) {
  return jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional int64 Y = 3;
 * @return {number}
 */
proto.room.Cell.prototype.getY = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.room.Cell} returns this
 */
proto.room.Cell.prototype.setY = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional bytes Field = 4;
 * @return {!(string|Uint8Array)}
 */
proto.room.Cell.prototype.getField = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * optional bytes Field = 4;
 * This is a type-conversion wrapper around `getField()`
 * @return {string}
 */
proto.room.Cell.prototype.getField_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getField()));
};


/**
 * optional bytes Field = 4;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getField()`
 * @return {!Uint8Array}
 */
proto.room.Cell.prototype.getField_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getField()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.room.Cell} returns this
 */
proto.room.Cell.prototype.setField = function(value) {
  return jspb.Message.setProto3BytesField(this, 4, value);
};


goog.object.extend(exports, proto.room);