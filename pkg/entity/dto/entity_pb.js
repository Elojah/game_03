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

var github_com_elojah_game_03_pkg_gogoproto_gogo_pb = require('../../../../../../github.com/elojah/game_03/pkg/gogoproto/gogo_pb.js');
var github_com_elojah_game_03_pkg_ability_ability_pb = require('../../../../../../github.com/elojah/game_03/pkg/ability/ability_pb.js');
var github_com_elojah_game_03_pkg_entity_entity_pb = require('../../../../../../github.com/elojah/game_03/pkg/entity/entity_pb.js');
goog.exportSymbol('proto.dto.CreateEntityAbilityReq', null, global);
goog.exportSymbol('proto.dto.CreateEntityAbilityResp', null, global);
goog.exportSymbol('proto.dto.ListEntityReq', null, global);
goog.exportSymbol('proto.dto.ListEntityResp', null, global);

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
proto.dto.ListEntityReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.dto.ListEntityReq.repeatedFields_, null);
};
goog.inherits(proto.dto.ListEntityReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.dto.ListEntityReq.displayName = 'proto.dto.ListEntityReq';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.dto.ListEntityReq.repeatedFields_ = [1,2];



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
proto.dto.ListEntityReq.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.ListEntityReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.ListEntityReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListEntityReq.toObject = function(includeInstance, msg) {
  var f, obj = {
    idsList: msg.getIdsList_asB64(),
    cellidsList: msg.getCellidsList_asB64(),
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
 * @return {!proto.dto.ListEntityReq}
 */
proto.dto.ListEntityReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.ListEntityReq;
  return proto.dto.ListEntityReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.ListEntityReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.ListEntityReq}
 */
proto.dto.ListEntityReq.deserializeBinaryFromReader = function(msg, reader) {
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
      msg.addCellids(value);
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
proto.dto.ListEntityReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.ListEntityReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.ListEntityReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListEntityReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getIdsList_asU8();
  if (f.length > 0) {
    writer.writeRepeatedBytes(
      1,
      f
    );
  }
  f = message.getCellidsList_asU8();
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
proto.dto.ListEntityReq.prototype.getIdsList = function() {
  return /** @type {!(Array<!Uint8Array>|Array<string>)} */ (jspb.Message.getRepeatedField(this, 1));
};


/**
 * repeated bytes IDs = 1;
 * This is a type-conversion wrapper around `getIdsList()`
 * @return {!Array<string>}
 */
proto.dto.ListEntityReq.prototype.getIdsList_asB64 = function() {
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
proto.dto.ListEntityReq.prototype.getIdsList_asU8 = function() {
  return /** @type {!Array<!Uint8Array>} */ (jspb.Message.bytesListAsU8(
      this.getIdsList()));
};


/** @param {!(Array<!Uint8Array>|Array<string>)} value */
proto.dto.ListEntityReq.prototype.setIdsList = function(value) {
  jspb.Message.setField(this, 1, value || []);
};


/**
 * @param {!(string|Uint8Array)} value
 * @param {number=} opt_index
 */
proto.dto.ListEntityReq.prototype.addIds = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 1, value, opt_index);
};


proto.dto.ListEntityReq.prototype.clearIdsList = function() {
  this.setIdsList([]);
};


/**
 * repeated bytes CellIDs = 2;
 * @return {!(Array<!Uint8Array>|Array<string>)}
 */
proto.dto.ListEntityReq.prototype.getCellidsList = function() {
  return /** @type {!(Array<!Uint8Array>|Array<string>)} */ (jspb.Message.getRepeatedField(this, 2));
};


/**
 * repeated bytes CellIDs = 2;
 * This is a type-conversion wrapper around `getCellidsList()`
 * @return {!Array<string>}
 */
proto.dto.ListEntityReq.prototype.getCellidsList_asB64 = function() {
  return /** @type {!Array<string>} */ (jspb.Message.bytesListAsB64(
      this.getCellidsList()));
};


/**
 * repeated bytes CellIDs = 2;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getCellidsList()`
 * @return {!Array<!Uint8Array>}
 */
proto.dto.ListEntityReq.prototype.getCellidsList_asU8 = function() {
  return /** @type {!Array<!Uint8Array>} */ (jspb.Message.bytesListAsU8(
      this.getCellidsList()));
};


/** @param {!(Array<!Uint8Array>|Array<string>)} value */
proto.dto.ListEntityReq.prototype.setCellidsList = function(value) {
  jspb.Message.setField(this, 2, value || []);
};


/**
 * @param {!(string|Uint8Array)} value
 * @param {number=} opt_index
 */
proto.dto.ListEntityReq.prototype.addCellids = function(value, opt_index) {
  jspb.Message.addToRepeatedField(this, 2, value, opt_index);
};


proto.dto.ListEntityReq.prototype.clearCellidsList = function() {
  this.setCellidsList([]);
};


/**
 * optional int64 Size = 3;
 * @return {number}
 */
proto.dto.ListEntityReq.prototype.getSize = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.dto.ListEntityReq.prototype.setSize = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional bytes State = 4;
 * @return {!(string|Uint8Array)}
 */
proto.dto.ListEntityReq.prototype.getState = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/**
 * optional bytes State = 4;
 * This is a type-conversion wrapper around `getState()`
 * @return {string}
 */
proto.dto.ListEntityReq.prototype.getState_asB64 = function() {
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
proto.dto.ListEntityReq.prototype.getState_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getState()));
};


/** @param {!(string|Uint8Array)} value */
proto.dto.ListEntityReq.prototype.setState = function(value) {
  jspb.Message.setProto3BytesField(this, 4, value);
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
proto.dto.ListEntityResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.dto.ListEntityResp.repeatedFields_, null);
};
goog.inherits(proto.dto.ListEntityResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.dto.ListEntityResp.displayName = 'proto.dto.ListEntityResp';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.dto.ListEntityResp.repeatedFields_ = [1];



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
proto.dto.ListEntityResp.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.ListEntityResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.ListEntityResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListEntityResp.toObject = function(includeInstance, msg) {
  var f, obj = {
    entitiesList: jspb.Message.toObjectList(msg.getEntitiesList(),
    github_com_elojah_game_03_pkg_entity_entity_pb.E.toObject, includeInstance),
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
 * @return {!proto.dto.ListEntityResp}
 */
proto.dto.ListEntityResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.ListEntityResp;
  return proto.dto.ListEntityResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.ListEntityResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.ListEntityResp}
 */
proto.dto.ListEntityResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new github_com_elojah_game_03_pkg_entity_entity_pb.E;
      reader.readMessage(value,github_com_elojah_game_03_pkg_entity_entity_pb.E.deserializeBinaryFromReader);
      msg.addEntities(value);
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
proto.dto.ListEntityResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.ListEntityResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.ListEntityResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.ListEntityResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEntitiesList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      github_com_elojah_game_03_pkg_entity_entity_pb.E.serializeBinaryToWriter
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
 * repeated entity.E Entities = 1;
 * @return {!Array<!proto.entity.E>}
 */
proto.dto.ListEntityResp.prototype.getEntitiesList = function() {
  return /** @type{!Array<!proto.entity.E>} */ (
    jspb.Message.getRepeatedWrapperField(this, github_com_elojah_game_03_pkg_entity_entity_pb.E, 1));
};


/** @param {!Array<!proto.entity.E>} value */
proto.dto.ListEntityResp.prototype.setEntitiesList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.entity.E=} opt_value
 * @param {number=} opt_index
 * @return {!proto.entity.E}
 */
proto.dto.ListEntityResp.prototype.addEntities = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.entity.E, opt_index);
};


proto.dto.ListEntityResp.prototype.clearEntitiesList = function() {
  this.setEntitiesList([]);
};


/**
 * optional bytes State = 2;
 * @return {!(string|Uint8Array)}
 */
proto.dto.ListEntityResp.prototype.getState = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes State = 2;
 * This is a type-conversion wrapper around `getState()`
 * @return {string}
 */
proto.dto.ListEntityResp.prototype.getState_asB64 = function() {
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
proto.dto.ListEntityResp.prototype.getState_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getState()));
};


/** @param {!(string|Uint8Array)} value */
proto.dto.ListEntityResp.prototype.setState = function(value) {
  jspb.Message.setProto3BytesField(this, 2, value);
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
proto.dto.CreateEntityAbilityReq = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dto.CreateEntityAbilityReq, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.dto.CreateEntityAbilityReq.displayName = 'proto.dto.CreateEntityAbilityReq';
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
proto.dto.CreateEntityAbilityReq.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.CreateEntityAbilityReq.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.CreateEntityAbilityReq} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.CreateEntityAbilityReq.toObject = function(includeInstance, msg) {
  var f, obj = {
    entityid: msg.getEntityid_asB64(),
    ability: (f = msg.getAbility()) && github_com_elojah_game_03_pkg_ability_ability_pb.A.toObject(includeInstance, f)
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
 * @return {!proto.dto.CreateEntityAbilityReq}
 */
proto.dto.CreateEntityAbilityReq.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.CreateEntityAbilityReq;
  return proto.dto.CreateEntityAbilityReq.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.CreateEntityAbilityReq} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.CreateEntityAbilityReq}
 */
proto.dto.CreateEntityAbilityReq.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setEntityid(value);
      break;
    case 2:
      var value = new github_com_elojah_game_03_pkg_ability_ability_pb.A;
      reader.readMessage(value,github_com_elojah_game_03_pkg_ability_ability_pb.A.deserializeBinaryFromReader);
      msg.setAbility(value);
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
proto.dto.CreateEntityAbilityReq.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.CreateEntityAbilityReq.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.CreateEntityAbilityReq} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.CreateEntityAbilityReq.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEntityid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getAbility();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_elojah_game_03_pkg_ability_ability_pb.A.serializeBinaryToWriter
    );
  }
};


/**
 * optional bytes EntityID = 1;
 * @return {!(string|Uint8Array)}
 */
proto.dto.CreateEntityAbilityReq.prototype.getEntityid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes EntityID = 1;
 * This is a type-conversion wrapper around `getEntityid()`
 * @return {string}
 */
proto.dto.CreateEntityAbilityReq.prototype.getEntityid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getEntityid()));
};


/**
 * optional bytes EntityID = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getEntityid()`
 * @return {!Uint8Array}
 */
proto.dto.CreateEntityAbilityReq.prototype.getEntityid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getEntityid()));
};


/** @param {!(string|Uint8Array)} value */
proto.dto.CreateEntityAbilityReq.prototype.setEntityid = function(value) {
  jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional ability.A Ability = 2;
 * @return {?proto.ability.A}
 */
proto.dto.CreateEntityAbilityReq.prototype.getAbility = function() {
  return /** @type{?proto.ability.A} */ (
    jspb.Message.getWrapperField(this, github_com_elojah_game_03_pkg_ability_ability_pb.A, 2));
};


/** @param {?proto.ability.A|undefined} value */
proto.dto.CreateEntityAbilityReq.prototype.setAbility = function(value) {
  jspb.Message.setWrapperField(this, 2, value);
};


proto.dto.CreateEntityAbilityReq.prototype.clearAbility = function() {
  this.setAbility(undefined);
};


/**
 * Returns whether this field is set.
 * @return {!boolean}
 */
proto.dto.CreateEntityAbilityReq.prototype.hasAbility = function() {
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
proto.dto.CreateEntityAbilityResp = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.dto.CreateEntityAbilityResp, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.dto.CreateEntityAbilityResp.displayName = 'proto.dto.CreateEntityAbilityResp';
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
proto.dto.CreateEntityAbilityResp.prototype.toObject = function(opt_includeInstance) {
  return proto.dto.CreateEntityAbilityResp.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.dto.CreateEntityAbilityResp} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.CreateEntityAbilityResp.toObject = function(includeInstance, msg) {
  var f, obj = {
    entityid: msg.getEntityid_asB64(),
    abilityid: msg.getAbilityid_asB64()
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
 * @return {!proto.dto.CreateEntityAbilityResp}
 */
proto.dto.CreateEntityAbilityResp.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.dto.CreateEntityAbilityResp;
  return proto.dto.CreateEntityAbilityResp.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.dto.CreateEntityAbilityResp} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.dto.CreateEntityAbilityResp}
 */
proto.dto.CreateEntityAbilityResp.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setEntityid(value);
      break;
    case 2:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setAbilityid(value);
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
proto.dto.CreateEntityAbilityResp.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.dto.CreateEntityAbilityResp.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.dto.CreateEntityAbilityResp} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.dto.CreateEntityAbilityResp.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getEntityid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getAbilityid_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      2,
      f
    );
  }
};


/**
 * optional bytes EntityID = 1;
 * @return {!(string|Uint8Array)}
 */
proto.dto.CreateEntityAbilityResp.prototype.getEntityid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes EntityID = 1;
 * This is a type-conversion wrapper around `getEntityid()`
 * @return {string}
 */
proto.dto.CreateEntityAbilityResp.prototype.getEntityid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getEntityid()));
};


/**
 * optional bytes EntityID = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getEntityid()`
 * @return {!Uint8Array}
 */
proto.dto.CreateEntityAbilityResp.prototype.getEntityid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getEntityid()));
};


/** @param {!(string|Uint8Array)} value */
proto.dto.CreateEntityAbilityResp.prototype.setEntityid = function(value) {
  jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional bytes AbilityID = 2;
 * @return {!(string|Uint8Array)}
 */
proto.dto.CreateEntityAbilityResp.prototype.getAbilityid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes AbilityID = 2;
 * This is a type-conversion wrapper around `getAbilityid()`
 * @return {string}
 */
proto.dto.CreateEntityAbilityResp.prototype.getAbilityid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getAbilityid()));
};


/**
 * optional bytes AbilityID = 2;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getAbilityid()`
 * @return {!Uint8Array}
 */
proto.dto.CreateEntityAbilityResp.prototype.getAbilityid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getAbilityid()));
};


/** @param {!(string|Uint8Array)} value */
proto.dto.CreateEntityAbilityResp.prototype.setAbilityid = function(value) {
  jspb.Message.setProto3BytesField(this, 2, value);
};


goog.object.extend(exports, proto.dto);
