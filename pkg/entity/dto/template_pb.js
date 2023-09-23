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
var github_com_elojah_game_03_pkg_entity_template_pb = require('../../../../../../github.com/elojah/game_03/pkg/entity/template_pb.js');
var github_com_elojah_game_03_pkg_entity_entity_pb = require('../../../../../../github.com/elojah/game_03/pkg/entity/entity_pb.js');
goog.exportSymbol('proto.dto.CreateTemplateReq', null, global);
goog.exportSymbol('proto.dto.ListTemplateReq', null, global);
goog.exportSymbol('proto.dto.ListTemplateResp', null, global);

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
proto.dto.CreateTemplateReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dto.CreateTemplateReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.dto.CreateTemplateReq.displayName = 'proto.dto.CreateTemplateReq';
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
proto.dto.CreateTemplateReq.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.CreateTemplateReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.CreateTemplateReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.CreateTemplateReq.toObject = function(includeInstance, msg) {
  var f, obj = {
    name: jspb.Message.getFieldWithDefault(msg, 1, ""),
    entity: (f = msg.getEntity()) && github_com_elojah_game_03_pkg_entity_entity_pb.E.toObject(includeInstance, f)
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
 * @return {!proto.dto.CreateTemplateReq}
 */
proto.dto.CreateTemplateReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.CreateTemplateReq;
  return proto.dto.CreateTemplateReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.CreateTemplateReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.CreateTemplateReq}
 */
proto.dto.CreateTemplateReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 2:
      var value = new github_com_elojah_game_03_pkg_entity_entity_pb.E;
      reader.readMessage(value,github_com_elojah_game_03_pkg_entity_entity_pb.E.deserializeBinaryFromReader);
      msg.setEntity(value);
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
proto.dto.CreateTemplateReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.CreateTemplateReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.CreateTemplateReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.CreateTemplateReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getEntity();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_elojah_game_03_pkg_entity_entity_pb.E.serializeBinaryToWriter
    );
  }
};


/**
 * optional string Name = 1;
 * @return {string}
 */
proto.dto.CreateTemplateReq.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.dto.CreateTemplateReq.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional entity.E Entity = 2;
 * @return {?proto.entity.E}
 */
proto.dto.CreateTemplateReq.prototype.getEntity = function() {
  return /** @type{?proto.entity.E} */ (
    jspb.Message.getWrapperField(this, github_com_elojah_game_03_pkg_entity_entity_pb.E, 2));
};


/** @param {?proto.entity.E|undefined} value */
proto.dto.CreateTemplateReq.prototype.setEntity = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.dto.CreateTemplateReq.prototype.clearEntity = function() {
  this.setEntity(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.dto.CreateTemplateReq.prototype.hasEntity = function() {
  return jspb.Message.getField(this, 2) != null;
};



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
proto.dto.ListTemplateReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dto.ListTemplateReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.dto.ListTemplateReq.displayName = 'proto.dto.ListTemplateReq';
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
proto.dto.ListTemplateReq.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.ListTemplateReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.ListTemplateReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListTemplateReq.toObject = function(includeInstance, msg) {
  var f, obj = {
    all: jspb.Message.getFieldWithDefault(msg, 1, false),
    size: jspb.Message.getFieldWithDefault(msg, 2, 0),
    state: msg.getState_asB64()
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
 * @return {!proto.dto.ListTemplateReq}
 */
proto.dto.ListTemplateReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.ListTemplateReq;
  return proto.dto.ListTemplateReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.ListTemplateReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.ListTemplateReq}
 */
proto.dto.ListTemplateReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setAll(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSize(value);
      break;
    case 3:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setState(value);
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
proto.dto.ListTemplateReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.ListTemplateReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.ListTemplateReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListTemplateReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAll();
  if (f) {
    writer.writeBool(
      1,
      f
    );
  }
  f = message.getSize();
  if (f !== 0) {
    writer.writeInt64(
      2,
      f
    );
  }
  f = message.getState_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      3,
      f
    );
  }
};


/**
 * optional bool All = 1;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.dto.ListTemplateReq.prototype.getAll = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 1, false));
};


/** @param {boolean} value */
proto.dto.ListTemplateReq.prototype.setAll = function(value) {
  jspb.Message.setProto3BooleanField(this, 1, value);
};


/**
 * optional int64 Size = 2;
 * @return {number}
 */
proto.dto.ListTemplateReq.prototype.getSize = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.dto.ListTemplateReq.prototype.setSize = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional bytes State = 3;
 * @return {!(string|Uint8Array)}
 */
proto.dto.ListTemplateReq.prototype.getState = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/**
 * optional bytes State = 3;
 * This is a type-conversion wrapper around `getState()`
 * @return {string}
 */
proto.dto.ListTemplateReq.prototype.getState_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getState()));
};


/**
 * optional bytes State = 3;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getState()`
 * @return {!Uint8Array}
 */
proto.dto.ListTemplateReq.prototype.getState_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getState()));
};


/** @param {!(string|Uint8Array)} value */
proto.dto.ListTemplateReq.prototype.setState = function(value) {
  jspb.Message.setProto3BytesField(this, 3, value);
};



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
proto.dto.ListTemplateResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.dto.ListTemplateResp.repeatedFields_, null);
};
goog.inherits(proto.dto.ListTemplateResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.dto.ListTemplateResp.displayName = 'proto.dto.ListTemplateResp';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.dto.ListTemplateResp.repeatedFields_ = [1];



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
proto.dto.ListTemplateResp.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.ListTemplateResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.ListTemplateResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListTemplateResp.toObject = function(includeInstance, msg) {
  var f, obj = {
    templatesList: jspb.Message.toObjectList(msg.getTemplatesList(),
    github_com_elojah_game_03_pkg_entity_template_pb.Template.toObject, includeInstance),
    state: msg.getState_asB64()
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
 * @return {!proto.dto.ListTemplateResp}
 */
proto.dto.ListTemplateResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.ListTemplateResp;
  return proto.dto.ListTemplateResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.ListTemplateResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.ListTemplateResp}
 */
proto.dto.ListTemplateResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_elojah_game_03_pkg_entity_template_pb.Template;
      reader.readMessage(value,github_com_elojah_game_03_pkg_entity_template_pb.Template.deserializeBinaryFromReader);
      msg.addTemplates(value);
      break;
    case 2:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setState(value);
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
proto.dto.ListTemplateResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.ListTemplateResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.ListTemplateResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListTemplateResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTemplatesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      github_com_elojah_game_03_pkg_entity_template_pb.Template.serializeBinaryToWriter
    );
  }
  f = message.getState_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      2,
      f
    );
  }
};


/**
 * repeated entity.Template Templates = 1;
 * @return {!Array<!proto.entity.Template>}
 */
proto.dto.ListTemplateResp.prototype.getTemplatesList = function() {
  return /** @type{!Array<!proto.entity.Template>} */ (
    jspb.Message.getRepeatedWrapperField(this, github_com_elojah_game_03_pkg_entity_template_pb.Template, 1));
};


/** @param {!Array<!proto.entity.Template>} value */
proto.dto.ListTemplateResp.prototype.setTemplatesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.entity.Template=} opt_value
 * @param {number=} opt_index
 * @return {!proto.entity.Template}
 */
proto.dto.ListTemplateResp.prototype.addTemplates = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.entity.Template, opt_index);
};


proto.dto.ListTemplateResp.prototype.clearTemplatesList = function() {
  this.setTemplatesList([]);
};


/**
 * optional bytes State = 2;
 * @return {!(string|Uint8Array)}
 */
proto.dto.ListTemplateResp.prototype.getState = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes State = 2;
 * This is a type-conversion wrapper around `getState()`
 * @return {string}
 */
proto.dto.ListTemplateResp.prototype.getState_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getState()));
};


/**
 * optional bytes State = 2;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getState()`
 * @return {!Uint8Array}
 */
proto.dto.ListTemplateResp.prototype.getState_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getState()));
};


/** @param {!(string|Uint8Array)} value */
proto.dto.ListTemplateResp.prototype.setState = function(value) {
  jspb.Message.setProto3BytesField(this, 2, value);
};


goog.object.extend(exports, proto.dto);
