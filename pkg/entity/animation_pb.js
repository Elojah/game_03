// source: github.com/elojah/game_03/pkg/entity/animation.proto
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
goog.object.extend(proto, github_com_gogo_protobuf_gogoproto_gogo_pb);
goog.exportSymbol('proto.entity.Animation', null, global);
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
proto.entity.Animation = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.entity.Animation, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.entity.Animation.displayName = 'proto.entity.Animation';
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
proto.entity.Animation.prototype.toObject = function(opt_includeInstance) {
  return proto.entity.Animation.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.entity.Animation} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.entity.Animation.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: msg.getId_asB64(),
    entityid: msg.getEntityid_asB64(),
    sheetid: msg.getSheetid_asB64(),
    name: jspb.Message.getFieldWithDefault(msg, 4, ""),
    start: jspb.Message.getFieldWithDefault(msg, 5, 0),
    end: jspb.Message.getFieldWithDefault(msg, 6, 0),
    rate: jspb.Message.getFieldWithDefault(msg, 7, 0)
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
 * @return {!proto.entity.Animation}
 */
proto.entity.Animation.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.entity.Animation;
  return proto.entity.Animation.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.entity.Animation} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.entity.Animation}
 */
proto.entity.Animation.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.setEntityid(value);
      break;
    case 3:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setSheetid(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setStart(value);
      break;
    case 6:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setEnd(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt32());
      msg.setRate(value);
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
proto.entity.Animation.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.entity.Animation.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.entity.Animation} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.entity.Animation.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getEntityid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      2,
      f
    );
  }
  f = message.getSheetid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      3,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getStart();
  if (f !== 0) {
    writer.writeInt64(
      5,
      f
    );
  }
  f = message.getEnd();
  if (f !== 0) {
    writer.writeInt64(
      6,
      f
    );
  }
  f = message.getRate();
  if (f !== 0) {
    writer.writeInt32(
      7,
      f
    );
  }
};


/**
 * optional bytes ID = 1;
 * @return {!(string|Uint8Array)}
 */
proto.entity.Animation.prototype.getId = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes ID = 1;
 * This is a type-conversion wrapper around `getId()`
 * @return {string}
 */
proto.entity.Animation.prototype.getId_asB64 = function() {
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
proto.entity.Animation.prototype.getId_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getId()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.entity.Animation} returns this
 */
proto.entity.Animation.prototype.setId = function(value) {
  return jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional bytes EntityID = 2;
 * @return {!(string|Uint8Array)}
 */
proto.entity.Animation.prototype.getEntityid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes EntityID = 2;
 * This is a type-conversion wrapper around `getEntityid()`
 * @return {string}
 */
proto.entity.Animation.prototype.getEntityid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getEntityid()));
};


/**
 * optional bytes EntityID = 2;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getEntityid()`
 * @return {!Uint8Array}
 */
proto.entity.Animation.prototype.getEntityid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getEntityid()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.entity.Animation} returns this
 */
proto.entity.Animation.prototype.setEntityid = function(value) {
  return jspb.Message.setProto3BytesField(this, 2, value);
};


/**
 * optional bytes SheetID = 3;
 * @return {!(string|Uint8Array)}
 */
proto.entity.Animation.prototype.getSheetid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * optional bytes SheetID = 3;
 * This is a type-conversion wrapper around `getSheetid()`
 * @return {string}
 */
proto.entity.Animation.prototype.getSheetid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getSheetid()));
};


/**
 * optional bytes SheetID = 3;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getSheetid()`
 * @return {!Uint8Array}
 */
proto.entity.Animation.prototype.getSheetid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getSheetid()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.entity.Animation} returns this
 */
proto.entity.Animation.prototype.setSheetid = function(value) {
  return jspb.Message.setProto3BytesField(this, 3, value);
};


/**
 * optional string Name = 4;
 * @return {string}
 */
proto.entity.Animation.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * @param {string} value
 * @return {!proto.entity.Animation} returns this
 */
proto.entity.Animation.prototype.setName = function(value) {
  return jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional int64 Start = 5;
 * @return {number}
 */
proto.entity.Animation.prototype.getStart = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/**
 * @param {number} value
 * @return {!proto.entity.Animation} returns this
 */
proto.entity.Animation.prototype.setStart = function(value) {
  return jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional int64 End = 6;
 * @return {number}
 */
proto.entity.Animation.prototype.getEnd = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 6, 0));
};


/**
 * @param {number} value
 * @return {!proto.entity.Animation} returns this
 */
proto.entity.Animation.prototype.setEnd = function(value) {
  return jspb.Message.setProto3IntField(this, 6, value);
};


/**
 * optional int32 Rate = 7;
 * @return {number}
 */
proto.entity.Animation.prototype.getRate = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/**
 * @param {number} value
 * @return {!proto.entity.Animation} returns this
 */
proto.entity.Animation.prototype.setRate = function(value) {
  return jspb.Message.setProto3IntField(this, 7, value);
};


goog.object.extend(exports, proto.entity);