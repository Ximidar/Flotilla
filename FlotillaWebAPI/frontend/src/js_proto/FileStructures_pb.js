/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.FileStructures = (function() {

    /**
     * Namespace FileStructures.
     * @exports FileStructures
     * @namespace
     */
    var FileStructures = {};

    FileStructures.FileAction = (function() {

        /**
         * Properties of a FileAction.
         * @memberof FileStructures
         * @interface IFileAction
         * @property {FileStructures.FileAction.Option|null} [Action] FileAction Action
         * @property {string|null} [Path] FileAction Path
         */

        /**
         * Constructs a new FileAction.
         * @memberof FileStructures
         * @classdesc Represents a FileAction.
         * @implements IFileAction
         * @constructor
         * @param {FileStructures.IFileAction=} [properties] Properties to set
         */
        function FileAction(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * FileAction Action.
         * @member {FileStructures.FileAction.Option} Action
         * @memberof FileStructures.FileAction
         * @instance
         */
        FileAction.prototype.Action = 0;

        /**
         * FileAction Path.
         * @member {string} Path
         * @memberof FileStructures.FileAction
         * @instance
         */
        FileAction.prototype.Path = "";

        /**
         * Creates a new FileAction instance using the specified properties.
         * @function create
         * @memberof FileStructures.FileAction
         * @static
         * @param {FileStructures.IFileAction=} [properties] Properties to set
         * @returns {FileStructures.FileAction} FileAction instance
         */
        FileAction.create = function create(properties) {
            return new FileAction(properties);
        };

        /**
         * Encodes the specified FileAction message. Does not implicitly {@link FileStructures.FileAction.verify|verify} messages.
         * @function encode
         * @memberof FileStructures.FileAction
         * @static
         * @param {FileStructures.IFileAction} message FileAction message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        FileAction.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Action != null && Object.hasOwnProperty.call(message, "Action"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.Action);
            if (message.Path != null && Object.hasOwnProperty.call(message, "Path"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.Path);
            return writer;
        };

        /**
         * Encodes the specified FileAction message, length delimited. Does not implicitly {@link FileStructures.FileAction.verify|verify} messages.
         * @function encodeDelimited
         * @memberof FileStructures.FileAction
         * @static
         * @param {FileStructures.IFileAction} message FileAction message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        FileAction.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a FileAction message from the specified reader or buffer.
         * @function decode
         * @memberof FileStructures.FileAction
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {FileStructures.FileAction} FileAction
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        FileAction.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.FileStructures.FileAction();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.Action = reader.int32();
                    break;
                case 2:
                    message.Path = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a FileAction message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof FileStructures.FileAction
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {FileStructures.FileAction} FileAction
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        FileAction.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a FileAction message.
         * @function verify
         * @memberof FileStructures.FileAction
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        FileAction.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Action != null && message.hasOwnProperty("Action"))
                switch (message.Action) {
                default:
                    return "Action: enum value expected";
                case 0:
                case 1:
                case 2:
                case 3:
                case 4:
                    break;
                }
            if (message.Path != null && message.hasOwnProperty("Path"))
                if (!$util.isString(message.Path))
                    return "Path: string expected";
            return null;
        };

        /**
         * Creates a FileAction message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof FileStructures.FileAction
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {FileStructures.FileAction} FileAction
         */
        FileAction.fromObject = function fromObject(object) {
            if (object instanceof $root.FileStructures.FileAction)
                return object;
            var message = new $root.FileStructures.FileAction();
            switch (object.Action) {
            case "SelectFile":
            case 0:
                message.Action = 0;
                break;
            case "GetFileStructure":
            case 1:
                message.Action = 1;
                break;
            case "AddFile":
            case 2:
                message.Action = 2;
                break;
            case "MoveFile":
            case 3:
                message.Action = 3;
                break;
            case "DeleteFile":
            case 4:
                message.Action = 4;
                break;
            }
            if (object.Path != null)
                message.Path = String(object.Path);
            return message;
        };

        /**
         * Creates a plain object from a FileAction message. Also converts values to other types if specified.
         * @function toObject
         * @memberof FileStructures.FileAction
         * @static
         * @param {FileStructures.FileAction} message FileAction
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        FileAction.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.Action = options.enums === String ? "SelectFile" : 0;
                object.Path = "";
            }
            if (message.Action != null && message.hasOwnProperty("Action"))
                object.Action = options.enums === String ? $root.FileStructures.FileAction.Option[message.Action] : message.Action;
            if (message.Path != null && message.hasOwnProperty("Path"))
                object.Path = message.Path;
            return object;
        };

        /**
         * Converts this FileAction to JSON.
         * @function toJSON
         * @memberof FileStructures.FileAction
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        FileAction.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        /**
         * Option enum.
         * @name FileStructures.FileAction.Option
         * @enum {number}
         * @property {number} SelectFile=0 SelectFile value
         * @property {number} GetFileStructure=1 GetFileStructure value
         * @property {number} AddFile=2 AddFile value
         * @property {number} MoveFile=3 MoveFile value
         * @property {number} DeleteFile=4 DeleteFile value
         */
        FileAction.Option = (function() {
            var valuesById = {}, values = Object.create(valuesById);
            values[valuesById[0] = "SelectFile"] = 0;
            values[valuesById[1] = "GetFileStructure"] = 1;
            values[valuesById[2] = "AddFile"] = 2;
            values[valuesById[3] = "MoveFile"] = 3;
            values[valuesById[4] = "DeleteFile"] = 4;
            return values;
        })();

        return FileAction;
    })();

    FileStructures.FileProg = (function() {

        /**
         * Properties of a FileProg.
         * @memberof FileStructures
         * @interface IFileProg
         * @property {string|null} [FileName] FileProg FileName
         * @property {number|Long|null} [Size] FileProg Size
         * @property {number|Long|null} [BytesRead] FileProg BytesRead
         * @property {number|Long|null} [CurrentLine] FileProg CurrentLine
         * @property {number|null} [Progress] FileProg Progress
         */

        /**
         * Constructs a new FileProg.
         * @memberof FileStructures
         * @classdesc Represents a FileProg.
         * @implements IFileProg
         * @constructor
         * @param {FileStructures.IFileProg=} [properties] Properties to set
         */
        function FileProg(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * FileProg FileName.
         * @member {string} FileName
         * @memberof FileStructures.FileProg
         * @instance
         */
        FileProg.prototype.FileName = "";

        /**
         * FileProg Size.
         * @member {number|Long} Size
         * @memberof FileStructures.FileProg
         * @instance
         */
        FileProg.prototype.Size = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * FileProg BytesRead.
         * @member {number|Long} BytesRead
         * @memberof FileStructures.FileProg
         * @instance
         */
        FileProg.prototype.BytesRead = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * FileProg CurrentLine.
         * @member {number|Long} CurrentLine
         * @memberof FileStructures.FileProg
         * @instance
         */
        FileProg.prototype.CurrentLine = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * FileProg Progress.
         * @member {number} Progress
         * @memberof FileStructures.FileProg
         * @instance
         */
        FileProg.prototype.Progress = 0;

        /**
         * Creates a new FileProg instance using the specified properties.
         * @function create
         * @memberof FileStructures.FileProg
         * @static
         * @param {FileStructures.IFileProg=} [properties] Properties to set
         * @returns {FileStructures.FileProg} FileProg instance
         */
        FileProg.create = function create(properties) {
            return new FileProg(properties);
        };

        /**
         * Encodes the specified FileProg message. Does not implicitly {@link FileStructures.FileProg.verify|verify} messages.
         * @function encode
         * @memberof FileStructures.FileProg
         * @static
         * @param {FileStructures.IFileProg} message FileProg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        FileProg.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.FileName != null && Object.hasOwnProperty.call(message, "FileName"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.FileName);
            if (message.Size != null && Object.hasOwnProperty.call(message, "Size"))
                writer.uint32(/* id 2, wireType 0 =*/16).uint64(message.Size);
            if (message.BytesRead != null && Object.hasOwnProperty.call(message, "BytesRead"))
                writer.uint32(/* id 3, wireType 0 =*/24).uint64(message.BytesRead);
            if (message.CurrentLine != null && Object.hasOwnProperty.call(message, "CurrentLine"))
                writer.uint32(/* id 4, wireType 0 =*/32).uint64(message.CurrentLine);
            if (message.Progress != null && Object.hasOwnProperty.call(message, "Progress"))
                writer.uint32(/* id 5, wireType 5 =*/45).float(message.Progress);
            return writer;
        };

        /**
         * Encodes the specified FileProg message, length delimited. Does not implicitly {@link FileStructures.FileProg.verify|verify} messages.
         * @function encodeDelimited
         * @memberof FileStructures.FileProg
         * @static
         * @param {FileStructures.IFileProg} message FileProg message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        FileProg.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a FileProg message from the specified reader or buffer.
         * @function decode
         * @memberof FileStructures.FileProg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {FileStructures.FileProg} FileProg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        FileProg.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.FileStructures.FileProg();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.FileName = reader.string();
                    break;
                case 2:
                    message.Size = reader.uint64();
                    break;
                case 3:
                    message.BytesRead = reader.uint64();
                    break;
                case 4:
                    message.CurrentLine = reader.uint64();
                    break;
                case 5:
                    message.Progress = reader.float();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a FileProg message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof FileStructures.FileProg
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {FileStructures.FileProg} FileProg
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        FileProg.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a FileProg message.
         * @function verify
         * @memberof FileStructures.FileProg
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        FileProg.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.FileName != null && message.hasOwnProperty("FileName"))
                if (!$util.isString(message.FileName))
                    return "FileName: string expected";
            if (message.Size != null && message.hasOwnProperty("Size"))
                if (!$util.isInteger(message.Size) && !(message.Size && $util.isInteger(message.Size.low) && $util.isInteger(message.Size.high)))
                    return "Size: integer|Long expected";
            if (message.BytesRead != null && message.hasOwnProperty("BytesRead"))
                if (!$util.isInteger(message.BytesRead) && !(message.BytesRead && $util.isInteger(message.BytesRead.low) && $util.isInteger(message.BytesRead.high)))
                    return "BytesRead: integer|Long expected";
            if (message.CurrentLine != null && message.hasOwnProperty("CurrentLine"))
                if (!$util.isInteger(message.CurrentLine) && !(message.CurrentLine && $util.isInteger(message.CurrentLine.low) && $util.isInteger(message.CurrentLine.high)))
                    return "CurrentLine: integer|Long expected";
            if (message.Progress != null && message.hasOwnProperty("Progress"))
                if (typeof message.Progress !== "number")
                    return "Progress: number expected";
            return null;
        };

        /**
         * Creates a FileProg message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof FileStructures.FileProg
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {FileStructures.FileProg} FileProg
         */
        FileProg.fromObject = function fromObject(object) {
            if (object instanceof $root.FileStructures.FileProg)
                return object;
            var message = new $root.FileStructures.FileProg();
            if (object.FileName != null)
                message.FileName = String(object.FileName);
            if (object.Size != null)
                if ($util.Long)
                    (message.Size = $util.Long.fromValue(object.Size)).unsigned = true;
                else if (typeof object.Size === "string")
                    message.Size = parseInt(object.Size, 10);
                else if (typeof object.Size === "number")
                    message.Size = object.Size;
                else if (typeof object.Size === "object")
                    message.Size = new $util.LongBits(object.Size.low >>> 0, object.Size.high >>> 0).toNumber(true);
            if (object.BytesRead != null)
                if ($util.Long)
                    (message.BytesRead = $util.Long.fromValue(object.BytesRead)).unsigned = true;
                else if (typeof object.BytesRead === "string")
                    message.BytesRead = parseInt(object.BytesRead, 10);
                else if (typeof object.BytesRead === "number")
                    message.BytesRead = object.BytesRead;
                else if (typeof object.BytesRead === "object")
                    message.BytesRead = new $util.LongBits(object.BytesRead.low >>> 0, object.BytesRead.high >>> 0).toNumber(true);
            if (object.CurrentLine != null)
                if ($util.Long)
                    (message.CurrentLine = $util.Long.fromValue(object.CurrentLine)).unsigned = true;
                else if (typeof object.CurrentLine === "string")
                    message.CurrentLine = parseInt(object.CurrentLine, 10);
                else if (typeof object.CurrentLine === "number")
                    message.CurrentLine = object.CurrentLine;
                else if (typeof object.CurrentLine === "object")
                    message.CurrentLine = new $util.LongBits(object.CurrentLine.low >>> 0, object.CurrentLine.high >>> 0).toNumber(true);
            if (object.Progress != null)
                message.Progress = Number(object.Progress);
            return message;
        };

        /**
         * Creates a plain object from a FileProg message. Also converts values to other types if specified.
         * @function toObject
         * @memberof FileStructures.FileProg
         * @static
         * @param {FileStructures.FileProg} message FileProg
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        FileProg.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.FileName = "";
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.Size = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.Size = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.BytesRead = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.BytesRead = options.longs === String ? "0" : 0;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.CurrentLine = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.CurrentLine = options.longs === String ? "0" : 0;
                object.Progress = 0;
            }
            if (message.FileName != null && message.hasOwnProperty("FileName"))
                object.FileName = message.FileName;
            if (message.Size != null && message.hasOwnProperty("Size"))
                if (typeof message.Size === "number")
                    object.Size = options.longs === String ? String(message.Size) : message.Size;
                else
                    object.Size = options.longs === String ? $util.Long.prototype.toString.call(message.Size) : options.longs === Number ? new $util.LongBits(message.Size.low >>> 0, message.Size.high >>> 0).toNumber(true) : message.Size;
            if (message.BytesRead != null && message.hasOwnProperty("BytesRead"))
                if (typeof message.BytesRead === "number")
                    object.BytesRead = options.longs === String ? String(message.BytesRead) : message.BytesRead;
                else
                    object.BytesRead = options.longs === String ? $util.Long.prototype.toString.call(message.BytesRead) : options.longs === Number ? new $util.LongBits(message.BytesRead.low >>> 0, message.BytesRead.high >>> 0).toNumber(true) : message.BytesRead;
            if (message.CurrentLine != null && message.hasOwnProperty("CurrentLine"))
                if (typeof message.CurrentLine === "number")
                    object.CurrentLine = options.longs === String ? String(message.CurrentLine) : message.CurrentLine;
                else
                    object.CurrentLine = options.longs === String ? $util.Long.prototype.toString.call(message.CurrentLine) : options.longs === Number ? new $util.LongBits(message.CurrentLine.low >>> 0, message.CurrentLine.high >>> 0).toNumber(true) : message.CurrentLine;
            if (message.Progress != null && message.hasOwnProperty("Progress"))
                object.Progress = options.json && !isFinite(message.Progress) ? String(message.Progress) : message.Progress;
            return object;
        };

        /**
         * Converts this FileProg to JSON.
         * @function toJSON
         * @memberof FileStructures.FileProg
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        FileProg.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return FileProg;
    })();

    FileStructures.File = (function() {

        /**
         * Properties of a File.
         * @memberof FileStructures
         * @interface IFile
         * @property {string|null} [PreviousPath] File PreviousPath
         * @property {string|null} [Name] File Name
         * @property {string|null} [Path] File Path
         * @property {string|null} [FileType] File FileType
         * @property {number|Long|null} [Size] File Size
         * @property {boolean|null} [IsDir] File IsDir
         * @property {number|Long|null} [UnixTime] File UnixTime
         * @property {Array.<FileStructures.IFile>|null} [Contents] File Contents
         */

        /**
         * Constructs a new File.
         * @memberof FileStructures
         * @classdesc Represents a File.
         * @implements IFile
         * @constructor
         * @param {FileStructures.IFile=} [properties] Properties to set
         */
        function File(properties) {
            this.Contents = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * File PreviousPath.
         * @member {string} PreviousPath
         * @memberof FileStructures.File
         * @instance
         */
        File.prototype.PreviousPath = "";

        /**
         * File Name.
         * @member {string} Name
         * @memberof FileStructures.File
         * @instance
         */
        File.prototype.Name = "";

        /**
         * File Path.
         * @member {string} Path
         * @memberof FileStructures.File
         * @instance
         */
        File.prototype.Path = "";

        /**
         * File FileType.
         * @member {string} FileType
         * @memberof FileStructures.File
         * @instance
         */
        File.prototype.FileType = "";

        /**
         * File Size.
         * @member {number|Long} Size
         * @memberof FileStructures.File
         * @instance
         */
        File.prototype.Size = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * File IsDir.
         * @member {boolean} IsDir
         * @memberof FileStructures.File
         * @instance
         */
        File.prototype.IsDir = false;

        /**
         * File UnixTime.
         * @member {number|Long} UnixTime
         * @memberof FileStructures.File
         * @instance
         */
        File.prototype.UnixTime = $util.Long ? $util.Long.fromBits(0,0,false) : 0;

        /**
         * File Contents.
         * @member {Array.<FileStructures.IFile>} Contents
         * @memberof FileStructures.File
         * @instance
         */
        File.prototype.Contents = $util.emptyArray;

        /**
         * Creates a new File instance using the specified properties.
         * @function create
         * @memberof FileStructures.File
         * @static
         * @param {FileStructures.IFile=} [properties] Properties to set
         * @returns {FileStructures.File} File instance
         */
        File.create = function create(properties) {
            return new File(properties);
        };

        /**
         * Encodes the specified File message. Does not implicitly {@link FileStructures.File.verify|verify} messages.
         * @function encode
         * @memberof FileStructures.File
         * @static
         * @param {FileStructures.IFile} message File message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        File.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.PreviousPath != null && Object.hasOwnProperty.call(message, "PreviousPath"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.PreviousPath);
            if (message.Name != null && Object.hasOwnProperty.call(message, "Name"))
                writer.uint32(/* id 2, wireType 2 =*/18).string(message.Name);
            if (message.Path != null && Object.hasOwnProperty.call(message, "Path"))
                writer.uint32(/* id 3, wireType 2 =*/26).string(message.Path);
            if (message.FileType != null && Object.hasOwnProperty.call(message, "FileType"))
                writer.uint32(/* id 4, wireType 2 =*/34).string(message.FileType);
            if (message.Size != null && Object.hasOwnProperty.call(message, "Size"))
                writer.uint32(/* id 5, wireType 0 =*/40).uint64(message.Size);
            if (message.IsDir != null && Object.hasOwnProperty.call(message, "IsDir"))
                writer.uint32(/* id 6, wireType 0 =*/48).bool(message.IsDir);
            if (message.UnixTime != null && Object.hasOwnProperty.call(message, "UnixTime"))
                writer.uint32(/* id 7, wireType 0 =*/56).int64(message.UnixTime);
            if (message.Contents != null && message.Contents.length)
                for (var i = 0; i < message.Contents.length; ++i)
                    $root.FileStructures.File.encode(message.Contents[i], writer.uint32(/* id 8, wireType 2 =*/66).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified File message, length delimited. Does not implicitly {@link FileStructures.File.verify|verify} messages.
         * @function encodeDelimited
         * @memberof FileStructures.File
         * @static
         * @param {FileStructures.IFile} message File message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        File.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a File message from the specified reader or buffer.
         * @function decode
         * @memberof FileStructures.File
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {FileStructures.File} File
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        File.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.FileStructures.File();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.PreviousPath = reader.string();
                    break;
                case 2:
                    message.Name = reader.string();
                    break;
                case 3:
                    message.Path = reader.string();
                    break;
                case 4:
                    message.FileType = reader.string();
                    break;
                case 5:
                    message.Size = reader.uint64();
                    break;
                case 6:
                    message.IsDir = reader.bool();
                    break;
                case 7:
                    message.UnixTime = reader.int64();
                    break;
                case 8:
                    if (!(message.Contents && message.Contents.length))
                        message.Contents = [];
                    message.Contents.push($root.FileStructures.File.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a File message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof FileStructures.File
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {FileStructures.File} File
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        File.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a File message.
         * @function verify
         * @memberof FileStructures.File
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        File.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.PreviousPath != null && message.hasOwnProperty("PreviousPath"))
                if (!$util.isString(message.PreviousPath))
                    return "PreviousPath: string expected";
            if (message.Name != null && message.hasOwnProperty("Name"))
                if (!$util.isString(message.Name))
                    return "Name: string expected";
            if (message.Path != null && message.hasOwnProperty("Path"))
                if (!$util.isString(message.Path))
                    return "Path: string expected";
            if (message.FileType != null && message.hasOwnProperty("FileType"))
                if (!$util.isString(message.FileType))
                    return "FileType: string expected";
            if (message.Size != null && message.hasOwnProperty("Size"))
                if (!$util.isInteger(message.Size) && !(message.Size && $util.isInteger(message.Size.low) && $util.isInteger(message.Size.high)))
                    return "Size: integer|Long expected";
            if (message.IsDir != null && message.hasOwnProperty("IsDir"))
                if (typeof message.IsDir !== "boolean")
                    return "IsDir: boolean expected";
            if (message.UnixTime != null && message.hasOwnProperty("UnixTime"))
                if (!$util.isInteger(message.UnixTime) && !(message.UnixTime && $util.isInteger(message.UnixTime.low) && $util.isInteger(message.UnixTime.high)))
                    return "UnixTime: integer|Long expected";
            if (message.Contents != null && message.hasOwnProperty("Contents")) {
                if (!Array.isArray(message.Contents))
                    return "Contents: array expected";
                for (var i = 0; i < message.Contents.length; ++i) {
                    var error = $root.FileStructures.File.verify(message.Contents[i]);
                    if (error)
                        return "Contents." + error;
                }
            }
            return null;
        };

        /**
         * Creates a File message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof FileStructures.File
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {FileStructures.File} File
         */
        File.fromObject = function fromObject(object) {
            if (object instanceof $root.FileStructures.File)
                return object;
            var message = new $root.FileStructures.File();
            if (object.PreviousPath != null)
                message.PreviousPath = String(object.PreviousPath);
            if (object.Name != null)
                message.Name = String(object.Name);
            if (object.Path != null)
                message.Path = String(object.Path);
            if (object.FileType != null)
                message.FileType = String(object.FileType);
            if (object.Size != null)
                if ($util.Long)
                    (message.Size = $util.Long.fromValue(object.Size)).unsigned = true;
                else if (typeof object.Size === "string")
                    message.Size = parseInt(object.Size, 10);
                else if (typeof object.Size === "number")
                    message.Size = object.Size;
                else if (typeof object.Size === "object")
                    message.Size = new $util.LongBits(object.Size.low >>> 0, object.Size.high >>> 0).toNumber(true);
            if (object.IsDir != null)
                message.IsDir = Boolean(object.IsDir);
            if (object.UnixTime != null)
                if ($util.Long)
                    (message.UnixTime = $util.Long.fromValue(object.UnixTime)).unsigned = false;
                else if (typeof object.UnixTime === "string")
                    message.UnixTime = parseInt(object.UnixTime, 10);
                else if (typeof object.UnixTime === "number")
                    message.UnixTime = object.UnixTime;
                else if (typeof object.UnixTime === "object")
                    message.UnixTime = new $util.LongBits(object.UnixTime.low >>> 0, object.UnixTime.high >>> 0).toNumber();
            if (object.Contents) {
                if (!Array.isArray(object.Contents))
                    throw TypeError(".FileStructures.File.Contents: array expected");
                message.Contents = [];
                for (var i = 0; i < object.Contents.length; ++i) {
                    if (typeof object.Contents[i] !== "object")
                        throw TypeError(".FileStructures.File.Contents: object expected");
                    message.Contents[i] = $root.FileStructures.File.fromObject(object.Contents[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a File message. Also converts values to other types if specified.
         * @function toObject
         * @memberof FileStructures.File
         * @static
         * @param {FileStructures.File} message File
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        File.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.Contents = [];
            if (options.defaults) {
                object.PreviousPath = "";
                object.Name = "";
                object.Path = "";
                object.FileType = "";
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.Size = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.Size = options.longs === String ? "0" : 0;
                object.IsDir = false;
                if ($util.Long) {
                    var long = new $util.Long(0, 0, false);
                    object.UnixTime = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.UnixTime = options.longs === String ? "0" : 0;
            }
            if (message.PreviousPath != null && message.hasOwnProperty("PreviousPath"))
                object.PreviousPath = message.PreviousPath;
            if (message.Name != null && message.hasOwnProperty("Name"))
                object.Name = message.Name;
            if (message.Path != null && message.hasOwnProperty("Path"))
                object.Path = message.Path;
            if (message.FileType != null && message.hasOwnProperty("FileType"))
                object.FileType = message.FileType;
            if (message.Size != null && message.hasOwnProperty("Size"))
                if (typeof message.Size === "number")
                    object.Size = options.longs === String ? String(message.Size) : message.Size;
                else
                    object.Size = options.longs === String ? $util.Long.prototype.toString.call(message.Size) : options.longs === Number ? new $util.LongBits(message.Size.low >>> 0, message.Size.high >>> 0).toNumber(true) : message.Size;
            if (message.IsDir != null && message.hasOwnProperty("IsDir"))
                object.IsDir = message.IsDir;
            if (message.UnixTime != null && message.hasOwnProperty("UnixTime"))
                if (typeof message.UnixTime === "number")
                    object.UnixTime = options.longs === String ? String(message.UnixTime) : message.UnixTime;
                else
                    object.UnixTime = options.longs === String ? $util.Long.prototype.toString.call(message.UnixTime) : options.longs === Number ? new $util.LongBits(message.UnixTime.low >>> 0, message.UnixTime.high >>> 0).toNumber() : message.UnixTime;
            if (message.Contents && message.Contents.length) {
                object.Contents = [];
                for (var j = 0; j < message.Contents.length; ++j)
                    object.Contents[j] = $root.FileStructures.File.toObject(message.Contents[j], options);
            }
            return object;
        };

        /**
         * Converts this File to JSON.
         * @function toJSON
         * @memberof FileStructures.File
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        File.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return File;
    })();

    return FileStructures;
})();

module.exports = $root;
