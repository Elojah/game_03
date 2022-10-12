// source: github.com/elojah/game_03/pkg/entity/dto/animation.proto
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
var github_com_elojah_game_03_pkg_entity_animation_pb = require('../../../../../../github.com/elojah/game_03/pkg/entity/animation_pb.js');
goog.object.extend(proto, github_com_elojah_game_03_pkg_entity_animation_pb);
goog.exportSymbol('proto.dto.CreateAnimationReq', null, global);
goog.exportSymbol('proto.dto.ListAnimationReq', null, global);
goog.exportSymbol('proto.dto.ListAnimationResp', null, global);
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
proto.dto.ListAnimationReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.dto.ListAnimationReq.repeatedFields_, null);
};
goog.inherits(proto.dto.ListAnimationReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dto.ListAnimationReq.displayName = 'proto.dto.ListAnimationReq';
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
proto.dto.ListAnimationResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.dto.ListAnimationResp.repeatedFields_, null);
};
goog.inherits(proto.dto.ListAnimationResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dto.ListAnimationResp.displayName = 'proto.dto.ListAnimationResp';
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
proto.dto.CreateAnimationReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dto.CreateAnimationReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.dto.CreateAnimationReq.displayName = 'proto.dto.CreateAnimationReq';
}

/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.dto.ListAnimationReq.repeatedFields_ = [1,2];



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
proto.dto.ListAnimationReq.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.ListAnimationReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.ListAnimationReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListAnimationReq.toObject = function(includeInstance, msg) {
  var f, obj = {
    idsList: msg.getIdsList_asB64(),
    entityidsList: msg.getEntityidsList_asB64(),
    size: jspb.Message.getFieldWithDefault(msg, 3, 0),
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
 * @return {!proto.dto.ListAnimationReq}
 */
proto.dto.ListAnimationReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.ListAnimationReq;
  return proto.dto.ListAnimationReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.ListAnimationReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.ListAnimationReq}
 */
proto.dto.ListAnimationReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.addIds(value);
      break;
    case 2:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.addEntityids(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setSize(value);
      break;
    case 4:
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
proto.dto.ListAnimationReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.ListAnimationReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.ListAnimationReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListAnimationReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getIdsList_asU8();
  if (f.length > 0) {
    writer.writeRepeatedBytes(
      1,
      f
    );
  }
  f = message.getEntityidsList_asU8();
  if (f.length > 0) {
    writer.writeRepeatedBytes(
      2,
      f
    );
  }
  f = message.getSize();
  if (f !== 0) {
    writer.writeInt64(
      3,
      f
    );
  }
  f = message.getState_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      4,
      f
    );
  }
};


/**
 * repeated bytes IDs = 1;
 * @return {!(Array<!Uint8Array>|Array<string>)}
 */
proto.dto.ListAnimationReq.prototype.getIdsList = function() {
  return /** @type {!(Array<!Uint8Array>|Array<string>)} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * repeated bytes IDs = 1;
 * This is a type-conversion wrapper around `getIdsList()`
 * @return {!Array<string>}
 */
proto.dto.ListAnimationReq.prototype.getIdsList_asB64 = function() {
  return /** @type {!Array<string>} */ (jspb.Message.bytesListAsB64(
      this.getIdsList()));
};


/**
 * repeated bytes IDs = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getIdsList()`
 * @return {!Array<!Uint8Array>}
 */
proto.dto.ListAnimationReq.prototype.getIdsList_asU8 = function() {
  return /** @type {!Array<!Uint8Array>} */ (jspb.Message.bytesListAsU8(
      this.getIdsList()));
};


/**
 * @param {!(Array<!Uint8Array>|Array<string>)} value
 * @return {!proto.dto.ListAnimationReq} returns this
 */
proto.dto.ListAnimationReq.prototype.setIdsList = function(value) {
  return jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {!(string|Uint8Array)} value
 * @param {number=} opt_index
 * @return {!proto.dto.ListAnimationReq} returns this
 */
proto.dto.ListAnimationReq.prototype.addIds = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.dto.ListAnimationReq} returns this
 */
proto.dto.ListAnimationReq.prototype.clearIdsList = function() {
  return this.setIdsList([]);
};


/**
 * repeated bytes EntityIDs = 2;
 * @return {!(Array<!Uint8Array>|Array<string>)}
 */
proto.dto.ListAnimationReq.prototype.getEntityidsList = function() {
  return /** @type {!(Array<!Uint8Array>|Array<string>)} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * repeated bytes EntityIDs = 2;
 * This is a type-conversion wrapper around `getEntityidsList()`
 * @return {!Array<string>}
 */
proto.dto.ListAnimationReq.prototype.getEntityidsList_asB64 = function() {
  return /** @type {!Array<string>} */ (jspb.Message.bytesListAsB64(
      this.getEntityidsList()));
};


/**
 * repeated bytes EntityIDs = 2;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getEntityidsList()`
 * @return {!Array<!Uint8Array>}
 */
proto.dto.ListAnimationReq.prototype.getEntityidsList_asU8 = function() {
  return /** @type {!Array<!Uint8Array>} */ (jspb.Message.bytesListAsU8(
      this.getEntityidsList()));
};


/**
 * @param {!(Array<!Uint8Array>|Array<string>)} value
 * @return {!proto.dto.ListAnimationReq} returns this
 */
proto.dto.ListAnimationReq.prototype.setEntityidsList = function(value) {
  return jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {!(string|Uint8Array)} value
 * @param {number=} opt_index
 * @return {!proto.dto.ListAnimationReq} returns this
 */
proto.dto.ListAnimationReq.prototype.addEntityids = function(value, opt_index) {
  return jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.dto.ListAnimationReq} returns this
 */
proto.dto.ListAnimationReq.prototype.clearEntityidsList = function() {
  return this.setEntityidsList([]);
};


/**
 * optional int64 Size = 3;
 * @return {number}
 */
proto.dto.ListAnimationReq.prototype.getSize = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/**
 * @param {number} value
 * @return {!proto.dto.ListAnimationReq} returns this
 */
proto.dto.ListAnimationReq.prototype.setSize = function(value) {
  return jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional bytes State = 4;
 * @return {!(string|Uint8Array)}
 */
proto.dto.ListAnimationReq.prototype.getState = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * optional bytes State = 4;
 * This is a type-conversion wrapper around `getState()`
 * @return {string}
 */
proto.dto.ListAnimationReq.prototype.getState_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getState()));
};


/**
 * optional bytes State = 4;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getState()`
 * @return {!Uint8Array}
 */
proto.dto.ListAnimationReq.prototype.getState_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getState()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.dto.ListAnimationReq} returns this
 */
proto.dto.ListAnimationReq.prototype.setState = function(value) {
  return jspb.Message.setProto3BytesField(this, 4, value);
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.dto.ListAnimationResp.repeatedFields_ = [1];



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
proto.dto.ListAnimationResp.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.ListAnimationResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.ListAnimationResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListAnimationResp.toObject = function(includeInstance, msg) {
  var f, obj = {
    animationsList: jspb.Message.toObjectList(msg.getAnimationsList(),
    github_com_elojah_game_03_pkg_entity_animation_pb.Animation.toObject, includeInstance),
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
 * @return {!proto.dto.ListAnimationResp}
 */
proto.dto.ListAnimationResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.ListAnimationResp;
  return proto.dto.ListAnimationResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.ListAnimationResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.ListAnimationResp}
 */
proto.dto.ListAnimationResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_elojah_game_03_pkg_entity_animation_pb.Animation;
      reader.readMessage(value,github_com_elojah_game_03_pkg_entity_animation_pb.Animation.deserializeBinaryFromReader);
      msg.addAnimations(value);
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
proto.dto.ListAnimationResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.ListAnimationResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.ListAnimationResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListAnimationResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAnimationsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      github_com_elojah_game_03_pkg_entity_animation_pb.Animation.serializeBinaryToWriter
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
 * repeated entity.Animation Animations = 1;
 * @return {!Array<!proto.entity.Animation>}
 */
proto.dto.ListAnimationResp.prototype.getAnimationsList = function() {
  return /** @type{!Array<!proto.entity.Animation>} */ (
    jspb.Message.getRepeatedWrapperField(this, github_com_elojah_game_03_pkg_entity_animation_pb.Animation, 1));
};


/**
 * @param {!Array<!proto.entity.Animation>} value
 * @return {!proto.dto.ListAnimationResp} returns this
*/
proto.dto.ListAnimationResp.prototype.setAnimationsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.entity.Animation=} opt_value
 * @param {number=} opt_index
 * @return {!proto.entity.Animation}
 */
proto.dto.ListAnimationResp.prototype.addAnimations = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.entity.Animation, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.dto.ListAnimationResp} returns this
 */
proto.dto.ListAnimationResp.prototype.clearAnimationsList = function() {
  return this.setAnimationsList([]);
};


/**
 * optional bytes State = 2;
 * @return {!(string|Uint8Array)}
 */
proto.dto.ListAnimationResp.prototype.getState = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes State = 2;
 * This is a type-conversion wrapper around `getState()`
 * @return {string}
 */
proto.dto.ListAnimationResp.prototype.getState_asB64 = function() {
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
proto.dto.ListAnimationResp.prototype.getState_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getState()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.dto.ListAnimationResp} returns this
 */
proto.dto.ListAnimationResp.prototype.setState = function(value) {
  return jspb.Message.setProto3BytesField(this, 2, value);
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
proto.dto.CreateAnimationReq.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.CreateAnimationReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.CreateAnimationReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.CreateAnimationReq.toObject = function(includeInstance, msg) {
  var f, obj = {
    entitytemplate: jspb.Message.getFieldWithDefault(msg, 1, ""),
    animation: (f = msg.getAnimation()) && github_com_elojah_game_03_pkg_entity_animation_pb.Animation.toObject(includeInstance, f)
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
 * @return {!proto.dto.CreateAnimationReq}
 */
proto.dto.CreateAnimationReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.CreateAnimationReq;
  return proto.dto.CreateAnimationReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.CreateAnimationReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.CreateAnimationReq}
 */
proto.dto.CreateAnimationReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setEntitytemplate(value);
      break;
    case 2:
      var value = new github_com_elojah_game_03_pkg_entity_animation_pb.Animation;
      reader.readMessage(value,github_com_elojah_game_03_pkg_entity_animation_pb.Animation.deserializeBinaryFromReader);
      msg.setAnimation(value);
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
proto.dto.CreateAnimationReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.CreateAnimationReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.CreateAnimationReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.CreateAnimationReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEntitytemplate();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getAnimation();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_elojah_game_03_pkg_entity_animation_pb.Animation.serializeBinaryToWriter
    );
  }
};


/**
 * optional string EntityTemplate = 1;
 * @return {string}
 */
proto.dto.CreateAnimationReq.prototype.getEntitytemplate = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * @param {string} value
 * @return {!proto.dto.CreateAnimationReq} returns this
 */
proto.dto.CreateAnimationReq.prototype.setEntitytemplate = function(value) {
  return jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional entity.Animation Animation = 2;
 * @return {?proto.entity.Animation}
 */
proto.dto.CreateAnimationReq.prototype.getAnimation = function() {
  return /** @type{?proto.entity.Animation} */ (
    jspb.Message.getWrapperField(this, github_com_elojah_game_03_pkg_entity_animation_pb.Animation, 2));
};


/**
 * @param {?proto.entity.Animation|undefined} value
 * @return {!proto.dto.CreateAnimationReq} returns this
*/
proto.dto.CreateAnimationReq.prototype.setAnimation = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.dto.CreateAnimationReq} returns this
 */
proto.dto.CreateAnimationReq.prototype.clearAnimation = function() {
  return this.setAnimation(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.dto.CreateAnimationReq.prototype.hasAnimation = function() {
  return jspb.Message.getField(this, 2) != null;
};


goog.object.extend(exports, proto.dto);
