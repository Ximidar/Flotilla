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

goog.exportSymbol('proto.FileStructures.File', null, global);
goog.exportSymbol('proto.FileStructures.FileAction', null, global);
goog.exportSymbol('proto.FileStructures.FileAction.Option', null, global);
goog.exportSymbol('proto.FileStructures.FileProg', null, global);

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
proto.FileStructures.FileAction = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.FileStructures.FileAction, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.FileStructures.FileAction.displayName = 'proto.FileStructures.FileAction';
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
proto.FileStructures.FileAction.prototype.toObject = function(opt_includeInstance) {
  return proto.FileStructures.FileAction.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.FileStructures.FileAction} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.FileStructures.FileAction.toObject = function(includeInstance, msg) {
  var f, obj = {
    action: jspb.Message.getFieldWithDefault(msg, 1, 0),
    path: jspb.Message.getFieldWithDefault(msg, 2, "")
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
 * @return {!proto.FileStructures.FileAction}
 */
proto.FileStructures.FileAction.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.FileStructures.FileAction;
  return proto.FileStructures.FileAction.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.FileStructures.FileAction} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.FileStructures.FileAction}
 */
proto.FileStructures.FileAction.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {!proto.FileStructures.FileAction.Option} */ (reader.readEnum());
      msg.setAction(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setPath(value);
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
proto.FileStructures.FileAction.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.FileStructures.FileAction.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.FileStructures.FileAction} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.FileStructures.FileAction.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getAction();
  if (f !== 0.0) {
    writer.writeEnum(
      1,
      f
    );
  }
  f = message.getPath();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
};


/**
 * @enum {number}
 */
proto.FileStructures.FileAction.Option = {
  SELECTFILE: 0,
  GETFILESTRUCTURE: 1,
  ADDFILE: 2,
  MOVEFILE: 3,
  DELETEFILE: 4
};

/**
 * optional Option Action = 1;
 * @return {!proto.FileStructures.FileAction.Option}
 */
proto.FileStructures.FileAction.prototype.getAction = function() {
  return /** @type {!proto.FileStructures.FileAction.Option} */ (jspb.Message.getFieldWithDefault(this, 1, 0));
};


/** @param {!proto.FileStructures.FileAction.Option} value */
proto.FileStructures.FileAction.prototype.setAction = function(value) {
  jspb.Message.setProto3EnumField(this, 1, value);
};


/**
 * optional string Path = 2;
 * @return {string}
 */
proto.FileStructures.FileAction.prototype.getPath = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.FileStructures.FileAction.prototype.setPath = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
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
proto.FileStructures.FileProg = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, null, null);
};
goog.inherits(proto.FileStructures.FileProg, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.FileStructures.FileProg.displayName = 'proto.FileStructures.FileProg';
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
proto.FileStructures.FileProg.prototype.toObject = function(opt_includeInstance) {
  return proto.FileStructures.FileProg.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.FileStructures.FileProg} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.FileStructures.FileProg.toObject = function(includeInstance, msg) {
  var f, obj = {
    filename: jspb.Message.getFieldWithDefault(msg, 1, ""),
    size: jspb.Message.getFieldWithDefault(msg, 2, 0),
    bytesread: jspb.Message.getFieldWithDefault(msg, 3, 0),
    currentline: jspb.Message.getFieldWithDefault(msg, 4, 0),
    progress: +jspb.Message.getFieldWithDefault(msg, 5, 0.0)
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
 * @return {!proto.FileStructures.FileProg}
 */
proto.FileStructures.FileProg.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.FileStructures.FileProg;
  return proto.FileStructures.FileProg.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.FileStructures.FileProg} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.FileStructures.FileProg}
 */
proto.FileStructures.FileProg.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setFilename(value);
      break;
    case 2:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setSize(value);
      break;
    case 3:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setBytesread(value);
      break;
    case 4:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setCurrentline(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readFloat());
      msg.setProgress(value);
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
proto.FileStructures.FileProg.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.FileStructures.FileProg.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.FileStructures.FileProg} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.FileStructures.FileProg.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getFilename();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getSize();
  if (f !== 0) {
    writer.writeUint64(
      2,
      f
    );
  }
  f = message.getBytesread();
  if (f !== 0) {
    writer.writeUint64(
      3,
      f
    );
  }
  f = message.getCurrentline();
  if (f !== 0) {
    writer.writeUint64(
      4,
      f
    );
  }
  f = message.getProgress();
  if (f !== 0.0) {
    writer.writeFloat(
      5,
      f
    );
  }
};


/**
 * optional string FileName = 1;
 * @return {string}
 */
proto.FileStructures.FileProg.prototype.getFilename = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.FileStructures.FileProg.prototype.setFilename = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional uint64 Size = 2;
 * @return {number}
 */
proto.FileStructures.FileProg.prototype.getSize = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 2, 0));
};


/** @param {number} value */
proto.FileStructures.FileProg.prototype.setSize = function(value) {
  jspb.Message.setProto3IntField(this, 2, value);
};


/**
 * optional uint64 BytesRead = 3;
 * @return {number}
 */
proto.FileStructures.FileProg.prototype.getBytesread = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 3, 0));
};


/** @param {number} value */
proto.FileStructures.FileProg.prototype.setBytesread = function(value) {
  jspb.Message.setProto3IntField(this, 3, value);
};


/**
 * optional uint64 CurrentLine = 4;
 * @return {number}
 */
proto.FileStructures.FileProg.prototype.getCurrentline = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 4, 0));
};


/** @param {number} value */
proto.FileStructures.FileProg.prototype.setCurrentline = function(value) {
  jspb.Message.setProto3IntField(this, 4, value);
};


/**
 * optional float Progress = 5;
 * @return {number}
 */
proto.FileStructures.FileProg.prototype.getProgress = function() {
  return /** @type {number} */ (+jspb.Message.getFieldWithDefault(this, 5, 0.0));
};


/** @param {number} value */
proto.FileStructures.FileProg.prototype.setProgress = function(value) {
  jspb.Message.setProto3FloatField(this, 5, value);
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
proto.FileStructures.File = function(opt_data) {
  jspb.Message.initialize(this, opt_data, 0, -1, proto.FileStructures.File.repeatedFields_, null);
};
goog.inherits(proto.FileStructures.File, jspb.Message);
if (goog.DEBUG && !COMPILED) {
  proto.FileStructures.File.displayName = 'proto.FileStructures.File';
}
/**
 * List of repeated fields within this message type.
 * @private {!Array<number>}
 * @const
 */
proto.FileStructures.File.repeatedFields_ = [8];



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
proto.FileStructures.File.prototype.toObject = function(opt_includeInstance) {
  return proto.FileStructures.File.toObject(opt_includeInstance, this);
};


/**
 * Static version of the {@see toObject} method.
 * @param {boolean|undefined} includeInstance Whether to include the JSPB
 *     instance for transitional soy proto support:
 *     http://goto/soy-param-migration
 * @param {!proto.FileStructures.File} msg The msg instance to transform.
 * @return {!Object}
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.FileStructures.File.toObject = function(includeInstance, msg) {
  var f, obj = {
    previouspath: jspb.Message.getFieldWithDefault(msg, 1, ""),
    name: jspb.Message.getFieldWithDefault(msg, 2, ""),
    path: jspb.Message.getFieldWithDefault(msg, 3, ""),
    filetype: jspb.Message.getFieldWithDefault(msg, 4, ""),
    size: jspb.Message.getFieldWithDefault(msg, 5, 0),
    isdir: jspb.Message.getFieldWithDefault(msg, 6, false),
    unixtime: jspb.Message.getFieldWithDefault(msg, 7, 0),
    contentsList: jspb.Message.toObjectList(msg.getContentsList(),
    proto.FileStructures.File.toObject, includeInstance)
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
 * @return {!proto.FileStructures.File}
 */
proto.FileStructures.File.deserializeBinary = function(bytes) {
  var reader = new jspb.BinaryReader(bytes);
  var msg = new proto.FileStructures.File;
  return proto.FileStructures.File.deserializeBinaryFromReader(msg, reader);
};


/**
 * Deserializes binary data (in protobuf wire format) from the
 * given reader into the given message object.
 * @param {!proto.FileStructures.File} msg The message object to deserialize into.
 * @param {!jspb.BinaryReader} reader The BinaryReader to use.
 * @return {!proto.FileStructures.File}
 */
proto.FileStructures.File.deserializeBinaryFromReader = function(msg, reader) {
  while (reader.nextField()) {
    if (reader.isEndGroup()) {
      break;
    }
    var field = reader.getFieldNumber();
    switch (field) {
    case 1:
      var value = /** @type {string} */ (reader.readString());
      msg.setPreviouspath(value);
      break;
    case 2:
      var value = /** @type {string} */ (reader.readString());
      msg.setName(value);
      break;
    case 3:
      var value = /** @type {string} */ (reader.readString());
      msg.setPath(value);
      break;
    case 4:
      var value = /** @type {string} */ (reader.readString());
      msg.setFiletype(value);
      break;
    case 5:
      var value = /** @type {number} */ (reader.readUint64());
      msg.setSize(value);
      break;
    case 6:
      var value = /** @type {boolean} */ (reader.readBool());
      msg.setIsdir(value);
      break;
    case 7:
      var value = /** @type {number} */ (reader.readInt64());
      msg.setUnixtime(value);
      break;
    case 8:
      var value = new proto.FileStructures.File;
      reader.readMessage(value,proto.FileStructures.File.deserializeBinaryFromReader);
      msg.addContents(value);
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
proto.FileStructures.File.prototype.serializeBinary = function() {
  var writer = new jspb.BinaryWriter();
  proto.FileStructures.File.serializeBinaryToWriter(this, writer);
  return writer.getResultBuffer();
};


/**
 * Serializes the given message to binary data (in protobuf wire
 * format), writing to the given BinaryWriter.
 * @param {!proto.FileStructures.File} message
 * @param {!jspb.BinaryWriter} writer
 * @suppress {unusedLocalVariables} f is only used for nested messages
 */
proto.FileStructures.File.serializeBinaryToWriter = function(message, writer) {
  var f = undefined;
  f = message.getPreviouspath();
  if (f.length > 0) {
    writer.writeString(
      1,
      f
    );
  }
  f = message.getName();
  if (f.length > 0) {
    writer.writeString(
      2,
      f
    );
  }
  f = message.getPath();
  if (f.length > 0) {
    writer.writeString(
      3,
      f
    );
  }
  f = message.getFiletype();
  if (f.length > 0) {
    writer.writeString(
      4,
      f
    );
  }
  f = message.getSize();
  if (f !== 0) {
    writer.writeUint64(
      5,
      f
    );
  }
  f = message.getIsdir();
  if (f) {
    writer.writeBool(
      6,
      f
    );
  }
  f = message.getUnixtime();
  if (f !== 0) {
    writer.writeInt64(
      7,
      f
    );
  }
  f = message.getContentsList();
  if (f.length > 0) {
    writer.writeRepeatedMessage(
      8,
      f,
      proto.FileStructures.File.serializeBinaryToWriter
    );
  }
};


/**
 * optional string PreviousPath = 1;
 * @return {string}
 */
proto.FileStructures.File.prototype.getPreviouspath = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 1, ""));
};


/** @param {string} value */
proto.FileStructures.File.prototype.setPreviouspath = function(value) {
  jspb.Message.setProto3StringField(this, 1, value);
};


/**
 * optional string Name = 2;
 * @return {string}
 */
proto.FileStructures.File.prototype.getName = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 2, ""));
};


/** @param {string} value */
proto.FileStructures.File.prototype.setName = function(value) {
  jspb.Message.setProto3StringField(this, 2, value);
};


/**
 * optional string Path = 3;
 * @return {string}
 */
proto.FileStructures.File.prototype.getPath = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 3, ""));
};


/** @param {string} value */
proto.FileStructures.File.prototype.setPath = function(value) {
  jspb.Message.setProto3StringField(this, 3, value);
};


/**
 * optional string FileType = 4;
 * @return {string}
 */
proto.FileStructures.File.prototype.getFiletype = function() {
  return /** @type {string} */ (jspb.Message.getFieldWithDefault(this, 4, ""));
};


/** @param {string} value */
proto.FileStructures.File.prototype.setFiletype = function(value) {
  jspb.Message.setProto3StringField(this, 4, value);
};


/**
 * optional uint64 Size = 5;
 * @return {number}
 */
proto.FileStructures.File.prototype.getSize = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 5, 0));
};


/** @param {number} value */
proto.FileStructures.File.prototype.setSize = function(value) {
  jspb.Message.setProto3IntField(this, 5, value);
};


/**
 * optional bool IsDir = 6;
 * Note that Boolean fields may be set to 0/1 when serialized from a Java server.
 * You should avoid comparisons like {@code val === true/false} in those cases.
 * @return {boolean}
 */
proto.FileStructures.File.prototype.getIsdir = function() {
  return /** @type {boolean} */ (jspb.Message.getFieldWithDefault(this, 6, false));
};


/** @param {boolean} value */
proto.FileStructures.File.prototype.setIsdir = function(value) {
  jspb.Message.setProto3BooleanField(this, 6, value);
};


/**
 * optional int64 UnixTime = 7;
 * @return {number}
 */
proto.FileStructures.File.prototype.getUnixtime = function() {
  return /** @type {number} */ (jspb.Message.getFieldWithDefault(this, 7, 0));
};


/** @param {number} value */
proto.FileStructures.File.prototype.setUnixtime = function(value) {
  jspb.Message.setProto3IntField(this, 7, value);
};


/**
 * repeated File Contents = 8;
 * @return {!Array<!proto.FileStructures.File>}
 */
proto.FileStructures.File.prototype.getContentsList = function() {
  return /** @type{!Array<!proto.FileStructures.File>} */ (
    jspb.Message.getRepeatedWrapperField(this, proto.FileStructures.File, 8));
};


/** @param {!Array<!proto.FileStructures.File>} value */
proto.FileStructures.File.prototype.setContentsList = function(value) {
  jspb.Message.setRepeatedWrapperField(this, 8, value);
};


/**
 * @param {!proto.FileStructures.File=} opt_value
 * @param {number=} opt_index
 * @return {!proto.FileStructures.File}
 */
proto.FileStructures.File.prototype.addContents = function(opt_value, opt_index) {
  return jspb.Message.addToRepeatedWrapperField(this, 8, opt_value, proto.FileStructures.File, opt_index);
};


proto.FileStructures.File.prototype.clearContentsList = function() {
  this.setContentsList([]);
};


goog.object.extend(exports, proto.FileStructures);