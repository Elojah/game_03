// source: github.com/elojah/game_03/pkg/ability/cast.proto
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
var github_com_elojah_game_03_pkg_geometry_geometry_pb = require('../../../../../github.com/elojah/game_03/pkg/geometry/geometry_pb.js');
goog.object.extend(proto, github_com_elojah_game_03_pkg_geometry_geometry_pb);
goog.exportSymbol('proto.ability.Cast', null, global);
goog.exportSymbol('proto.ability.CastTarget', null, global);
goog.exportSymbol('proto.ability.CastTargets', null, global);
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
proto.ability.CastTarget = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ability.CastTarget, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ability.CastTarget.displayName = 'proto.ability.CastTarget';
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
proto.ability.CastTargets = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.ability.CastTargets.repeatedFields_, null);
};
goog.inherits(proto.ability.CastTargets, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ability.CastTargets.displayName = 'proto.ability.CastTargets';
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
proto.ability.Cast = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.ability.Cast, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  /**
   * @public
   * @override
   */
  proto.ability.Cast.displayName = 'proto.ability.Cast';
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
proto.ability.CastTarget.prototype.toObject = function(opt_includeInstance) {
  return proto.ability.CastTarget.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ability.CastTarget} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ability.CastTarget.toObject = function(includeInstance, msg) {
  var f, obj = {
    id: msg.getId_asB64(),
    rect: (f = msg.getRect()) && github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect.toObject(includeInstance, f),
    circle: (f = msg.getCircle()) && github_com_elojah_game_03_pkg_geometry_geometry_pb.Circle.toObject(includeInstance, f)
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
 * @return {!proto.ability.CastTarget}
 */
proto.ability.CastTarget.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ability.CastTarget;
  return proto.ability.CastTarget.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ability.CastTarget} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ability.CastTarget}
 */
proto.ability.CastTarget.deserializeBinaryFromReader = function(msg, reader) {
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
      var value = new github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect;
      reader.readMessage(value,github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect.deserializeBinaryFromReader);
      msg.setRect(value);
      break;
    case 3:
      var value = new github_com_elojah_game_03_pkg_geometry_geometry_pb.Circle;
      reader.readMessage(value,github_com_elojah_game_03_pkg_geometry_geometry_pb.Circle.deserializeBinaryFromReader);
      msg.setCircle(value);
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
proto.ability.CastTarget.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ability.CastTarget.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ability.CastTarget} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ability.CastTarget.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getId_asU8();
  if (f.length > 0) {
    writer.writeBytes(
      1,
      f
    );
  }
  f = message.getRect();
  if (f != null) {
    writer.writeMessage(
      2,
      f,
      github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect.serializeBinaryToWriter
    );
  }
  f = message.getCircle();
  if (f != null) {
    writer.writeMessage(
      3,
      f,
      github_com_elojah_game_03_pkg_geometry_geometry_pb.Circle.serializeBinaryToWriter
    );
  }
};


/**
 * optional bytes ID = 1;
 * @return {!(string|Uint8Array)}
 */
proto.ability.CastTarget.prototype.getId = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes ID = 1;
 * This is a type-conversion wrapper around `getId()`
 * @return {string}
 */
proto.ability.CastTarget.prototype.getId_asB64 = function() {
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
proto.ability.CastTarget.prototype.getId_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getId()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.ability.CastTarget} returns this
 */
proto.ability.CastTarget.prototype.setId = function(value) {
  return jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional geometry.Rect Rect = 2;
 * @return {?proto.geometry.Rect}
 */
proto.ability.CastTarget.prototype.getRect = function() {
  return /** @type{?proto.geometry.Rect} */ (
    jspb.Message.getWrapperField(this, github_com_elojah_game_03_pkg_geometry_geometry_pb.Rect, 2));
};


/**
 * @param {?proto.geometry.Rect|undefined} value
 * @return {!proto.ability.CastTarget} returns this
*/
proto.ability.CastTarget.prototype.setRect = function(value) {
  return jspb.Message.setWrapperField(this, 2, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.ability.CastTarget} returns this
 */
proto.ability.CastTarget.prototype.clearRect = function() {
  return this.setRect(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.ability.CastTarget.prototype.hasRect = function() {
  return jspb.Message.getField(this, 2) != null;
};


/**
 * optional geometry.Circle Circle = 3;
 * @return {?proto.geometry.Circle}
 */
proto.ability.CastTarget.prototype.getCircle = function() {
  return /** @type{?proto.geometry.Circle} */ (
    jspb.Message.getWrapperField(this, github_com_elojah_game_03_pkg_geometry_geometry_pb.Circle, 3));
};


/**
 * @param {?proto.geometry.Circle|undefined} value
 * @return {!proto.ability.CastTarget} returns this
*/
proto.ability.CastTarget.prototype.setCircle = function(value) {
  return jspb.Message.setWrapperField(this, 3, value);
};


/**
 * Clears the message field making it undefined.
 * @return {!proto.ability.CastTarget} returns this
 */
proto.ability.CastTarget.prototype.clearCircle = function() {
  return this.setCircle(undefined);
};


/**
 * Returns whether this field is set.
 * @return {boolean}
 */
proto.ability.CastTarget.prototype.hasCircle = function() {
  return jspb.Message.getField(this, 3) != null;
};



/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.ability.CastTargets.repeatedFields_ = [1];



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
proto.ability.CastTargets.prototype.toObject = function(opt_includeInstance) {
  return proto.ability.CastTargets.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ability.CastTargets} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ability.CastTargets.toObject = function(includeInstance, msg) {
  var f, obj = {
    targetsList: jspb.Message.toObjectList(msg.getTargetsList(),
    proto.ability.CastTarget.toObject, includeInstance)
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
 * @return {!proto.ability.CastTargets}
 */
proto.ability.CastTargets.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ability.CastTargets;
  return proto.ability.CastTargets.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ability.CastTargets} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ability.CastTargets}
 */
proto.ability.CastTargets.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = new proto.ability.CastTarget;
      reader.readMessage(value,proto.ability.CastTarget.deserializeBinaryFromReader);
      msg.addTargets(value);
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
proto.ability.CastTargets.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ability.CastTargets.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ability.CastTargets} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ability.CastTargets.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getTargetsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      1,
      f,
      proto.ability.CastTarget.serializeBinaryToWriter
    );
  }
};


/**
 * repeated CastTarget Targets = 1;
 * @return {!Array<!proto.ability.CastTarget>}
 */
proto.ability.CastTargets.prototype.getTargetsList = function() {
  return /** @type{!Array<!proto.ability.CastTarget>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.ability.CastTarget, 1));
};


/**
 * @param {!Array<!proto.ability.CastTarget>} value
 * @return {!proto.ability.CastTargets} returns this
*/
proto.ability.CastTargets.prototype.setTargetsList = function(value) {
  return jspb.Message.setRepeatedWrapperField(this, 1, value);
};


/**
 * @param {!proto.ability.CastTarget=} opt_value
 * @param {number=} opt_index
 * @return {!proto.ability.CastTarget}
 */
proto.ability.CastTargets.prototype.addTargets = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 1, opt_value, proto.ability.CastTarget, opt_index);
};


/**
 * Clears the list making it empty but non-null.
 * @return {!proto.ability.CastTargets} returns this
 */
proto.ability.CastTargets.prototype.clearTargetsList = function() {
  return this.setTargetsList([]);
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
proto.ability.Cast.prototype.toObject = function(opt_includeInstance) {
  return proto.ability.Cast.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Deprecated. Whether to include
 *     the JSPB instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.ability.Cast} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ability.Cast.toObject = function(includeInstance, msg) {
  var f, obj = {
    sourceid: msg.getSourceid_asB64(),
    abilityid: msg.getAbilityid_asB64(),
    targetsMap: (f = msg.getTargetsMap()) ? f.toObject(includeInstance, proto.ability.CastTargets.toObject) : []
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
 * @return {!proto.ability.Cast}
 */
proto.ability.Cast.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.ability.Cast;
  return proto.ability.Cast.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.ability.Cast} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.ability.Cast}
 */
proto.ability.Cast.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setSourceid(value);
      break;
    case 2:
      var value = /** @type {!Uint8Array} */ (reader.readBytes());
      msg.setAbilityid(value);
      break;
    case 3:
      var value = msg.getTargetsMap();
      reader.readMessage(value, function(message, reader) {
        jspb.Map.deserializeBinary(message, reader, jspb.BinaryReader.prototype.readString, jspb.BinaryReader.prototype.readMessage, proto.ability.CastTargets.deserializeBinaryFromReader, "", new proto.ability.CastTargets());
         });
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
proto.ability.Cast.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.ability.Cast.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.ability.Cast} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.ability.Cast.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getSourceid_asU8();
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
  f = message.getTargetsMap(true);
  if (f && f.getLength() > 0) {
    f.serializeBinary(3, writer, jspb.BinaryWriter.prototype.writeString, jspb.BinaryWriter.prototype.writeMessage, proto.ability.CastTargets.serializeBinaryToWriter);
  }
};


/**
 * optional bytes SourceID = 1;
 * @return {!(string|Uint8Array)}
 */
proto.ability.Cast.prototype.getSourceid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/**
 * optional bytes SourceID = 1;
 * This is a type-conversion wrapper around `getSourceid()`
 * @return {string}
 */
proto.ability.Cast.prototype.getSourceid_asB64 = function() {
  return /** @type {string} */ (jspb.Message.bytesAsB64(
      this.getSourceid()));
};


/**
 * optional bytes SourceID = 1;
 * Note that Uint8Array is not supported on all browsers.
 * @see http://caniuse.com/Uint8Array
 * This is a type-conversion wrapper around `getSourceid()`
 * @return {!Uint8Array}
 */
proto.ability.Cast.prototype.getSourceid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getSourceid()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.ability.Cast} returns this
 */
proto.ability.Cast.prototype.setSourceid = function(value) {
  return jspb.Message.setProto3BytesField(this, 1, value);
};


/**
 * optional bytes AbilityID = 2;
 * @return {!(string|Uint8Array)}
 */
proto.ability.Cast.prototype.getAbilityid = function() {
  return /** @type {!(string|Uint8Array)} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/**
 * optional bytes AbilityID = 2;
 * This is a type-conversion wrapper around `getAbilityid()`
 * @return {string}
 */
proto.ability.Cast.prototype.getAbilityid_asB64 = function() {
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
proto.ability.Cast.prototype.getAbilityid_asU8 = function() {
  return /** @type {!Uint8Array} */ (jspb.Message.bytesAsU8(
      this.getAbilityid()));
};


/**
 * @param {!(string|Uint8Array)} value
 * @return {!proto.ability.Cast} returns this
 */
proto.ability.Cast.prototype.setAbilityid = function(value) {
  return jspb.Message.setProto3BytesField(this, 2, value);
};


/**
 * map<string, CastTargets> Targets = 3;
 * @param {boolean=} opt_noLazyCreate Do not create the map if
 * empty, instead returning `undefined`
 * @return {!jspb.Map<string,!proto.ability.CastTargets>}
 */
proto.ability.Cast.prototype.getTargetsMap = function(opt_noLazyCreate) {
  return /** @type {!jspb.Map<string,!proto.ability.CastTargets>} */ (
      jspb.Message.getMapField(this, 3, opt_noLazyCreate,
      proto.ability.CastTargets));
};


/**
 * Clears values from the map. The map will be non-null.
 * @return {!proto.ability.Cast} returns this
 */
proto.ability.Cast.prototype.clearTargetsMap = function() {
  this.getTargetsMap().clear();
  return this;};


goog.object.extend(exports, proto.ability);
