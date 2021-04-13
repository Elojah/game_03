// source: github.com/elojah/game_03/pkg/entity/dto/pc.proto
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

var github_com_gogo_protobuf_gogoproto_gogo_pb = require('../../../../../../github.com/gogo/protobuf/gogoproto/gogo_pb.js');
goog.object.extend(proto, github_com_gogo_protobuf_gogoproto_gogo_pb);
var github_com_elojah_game_03_pkg_entity_pc_pb = require('../../../../../../github.com/elojah/game_03/pkg/entity/pc_pb.js');
goog.object.extend(proto, github_com_elojah_game_03_pkg_entity_pc_pb);
goog.exportSymbol('proto.dto.ListPCReq', null, global);
goog.exportSymbol('proto.dto.ListPCResp', null, global);
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
proto.dto.ListPCReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dto.ListPCReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dto.ListPCReq.displayName = 'proto.dto.ListPCReq';
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
proto.dto.ListPCResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.dto.ListPCResp.repeatedFields_, null);
};
goog.inherits(proto.dto.ListPCResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dto.ListPCResp.displayName = 'proto.dto.ListPCResp';
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
proto.dto.ListPCReq.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.ListPCReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.ListPCReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListPCReq.toObject = function(includeInstance, msg) {
  var f, obj = {

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
 * @return {!proto.dto.ListPCReq}
 */
proto.dto.ListPCReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.ListPCReq;
  return proto.dto.ListPCReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.ListPCReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.ListPCReq}
 */
proto.dto.ListPCReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
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
proto.dto.ListPCReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.ListPCReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.ListPCReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListPCReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.dto.ListPCResp.repeatedFields_ = [1];



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
proto.dto.ListPCResp.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.ListPCResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.ListPCResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListPCResp.toObject = function(includeInstance, msg) {
  var f, obj = {
    pcsList: jspb.Message.toObjectList(msg.getPcsList(),
    github_com_elojah_game_03_pkg_entity_pc_pb.PC.toObject, includeInstance)
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
 * @return {!proto.dto.ListPCResp}
 */
proto.dto.ListPCResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.ListPCResp;
  return proto.dto.ListPCResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.ListPCResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.ListPCResp}
 */
proto.dto.ListPCResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_elojah_game_03_pkg_entity_pc_pb.PC;
      reader.readMessage(value,github_com_elojah_game_03_pkg_entity_pc_pb.PC.deserializeBinaryFromReader);
      msg.addPcs(value);
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
proto.dto.ListPCResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.ListPCResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.ListPCResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListPCResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPcsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      github_com_elojah_game_03_pkg_entity_pc_pb.PC.serializeBinaryToWriter
    );
  }
};


/**
 * repeated entity.PC PCs = 1;
 * @return {!Array<!proto.entity.PC>}
 */
proto.dto.ListPCResp.prototype.getPcsList = function() {
  return /** @type{!Array<!proto.entity.PC>} */ (
    jspb.Message.getRepeatedWrapperField(this, github_com_elojah_game_03_pkg_entity_pc_pb.PC, 1));
};


/**
 * @param {!Array<!proto.entity.PC>} value
 * @return {!proto.dto.ListPCResp} returns this
*/
proto.dto.ListPCResp.prototype.setPcsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.entity.PC=} opt_value
 * @param {number=} opt_index
 * @return {!proto.entity.PC}
 */
proto.dto.ListPCResp.prototype.addPcs = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.entity.PC, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.dto.ListPCResp} returns this
 */
proto.dto.ListPCResp.prototype.clearPcsList = function() {
  return this.setPcsList([]);
};


goog.object.extend(exports, proto.dto);
