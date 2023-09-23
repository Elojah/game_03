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

var github_com_gogo_protobuf_gogoproto_gogo_pb = require('../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb.js');
goog.exportSymbol('proto.room.WorldSpawn', null, global);

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
proto.room.WorldSpawn = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.room.WorldSpawn, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.room.WorldSpawn.displayName = 'proto.room.WorldSpawn';
}


if (jspb.Message.GENERATE_TO_OBJECT) {
/**
 * Creates an object representation of this proto suitable for use in Soy templates.
 * Field names that are reserved in JavaScript and will be renamed to pb_name.
 * To access a reserved field use, foo.pb_<name>, eg, foo.pb_default.
 * For the list of reserved names please see:
 *     com.google.apps.jspb.JsClassTemplate.JS_RESERVED_WORDS.
 * @param {boolean=} opt_includeInstance Whether to include the JSPB instance
 *     for transitional soy proto support: http://goto/soy-param-migration
 * @return {!Object}
 */
proto.room.WorldSpawn.prototype.toObject = function(opt_includeInstance) {
  return proto.room.WorldSpawn.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.room.WorldSpawn} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.room.WorldSpawn.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: msg.getId_asB64(),
    worldid: msg.getWorldid_asB64(),
    entityid: msg.getEntityid_asB64(),
    ownerid: msg.getOwnerid_asB64()
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
 * @return {!proto.room.WorldSpawn}
 */
proto.room.WorldSpawn.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.room.WorldSpawn;
  return proto.room.WorldSpawn.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.room.WorldSpawn} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.room.WorldSpawn}
 */
proto.room.WorldSpawn.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setWorldid(value);
      break;
    case 3:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setEntityid(value);
      break;
    case 4:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setOwnerid(value);
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
proto.room.WorldSpawn.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.room.WorldSpawn.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.room.WorldSpawn} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.room.WorldSpawn.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getWorldid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      2,
      f
    );
  }
  f = message.getEntityid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      3,
      f
    );
  }
  f = message.getOwnerid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      4,
      f
    );
  }
};


/**
 * optional bytes ID = 1;
 * @return {!(string|Uint8Array)}
 */
proto.room.WorldSpawn.prototype.getId = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes ID = 1;
 * This is a type-conversion wrapper around `getId()`
 * @return {string}
 */
proto.room.WorldSpawn.prototype.getId_asB64 = function() {
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
proto.room.WorldSpawn.prototype.getId_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getId()));
};


/** @param {!(string|Uint8Array)} value */
proto.room.WorldSpawn.prototype.setId = function(value) {
  jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional bytes WorldID = 2;
 * @return {!(string|Uint8Array)}
 */
proto.room.WorldSpawn.prototype.getWorldid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes WorldID = 2;
 * This is a type-conversion wrapper around `getWorldid()`
 * @return {string}
 */
proto.room.WorldSpawn.prototype.getWorldid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getWorldid()));
};


/**
 * optional bytes WorldID = 2;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getWorldid()`
 * @return {!Uint8Array}
 */
proto.room.WorldSpawn.prototype.getWorldid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getWorldid()));
};


/** @param {!(string|Uint8Array)} value */
proto.room.WorldSpawn.prototype.setWorldid = function(value) {
  jspb.Message.setProto3BytesField(this, 2, value);
};


/**
 * optional bytes EntityID = 3;
 * @return {!(string|Uint8Array)}
 */
proto.room.WorldSpawn.prototype.getEntityid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * optional bytes EntityID = 3;
 * This is a type-conversion wrapper around `getEntityid()`
 * @return {string}
 */
proto.room.WorldSpawn.prototype.getEntityid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getEntityid()));
};


/**
 * optional bytes EntityID = 3;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getEntityid()`
 * @return {!Uint8Array}
 */
proto.room.WorldSpawn.prototype.getEntityid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getEntityid()));
};


/** @param {!(string|Uint8Array)} value */
proto.room.WorldSpawn.prototype.setEntityid = function(value) {
  jspb.Message.setProto3BytesField(this, 3, value);
};


/**
 * optional bytes OwnerID = 4;
 * @return {!(string|Uint8Array)}
 */
proto.room.WorldSpawn.prototype.getOwnerid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * optional bytes OwnerID = 4;
 * This is a type-conversion wrapper around `getOwnerid()`
 * @return {string}
 */
proto.room.WorldSpawn.prototype.getOwnerid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getOwnerid()));
};


/**
 * optional bytes OwnerID = 4;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getOwnerid()`
 * @return {!Uint8Array}
 */
proto.room.WorldSpawn.prototype.getOwnerid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getOwnerid()));
};


/** @param {!(string|Uint8Array)} value */
proto.room.WorldSpawn.prototype.setOwnerid = function(value) {
  jspb.Message.setProto3BytesField(this, 4, value);
};


goog.object.extend(exports, proto.room);
