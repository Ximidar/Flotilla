/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.CommStructures = (function() {

    /**
     * Namespace CommStructures.
     * @exports CommStructures
     * @namespace
     */
    var CommStructures = {};

    CommStructures.CommStatus = (function() {

        /**
         * Properties of a CommStatus.
         * @memberof CommStructures
         * @interface ICommStatus
         * @property {string|null} [port] CommStatus port
         * @property {number|null} [baud] CommStatus baud
         * @property {boolean|null} [connected] CommStatus connected
         */

        /**
         * Constructs a new CommStatus.
         * @memberof CommStructures
         * @classdesc Represents a CommStatus.
         * @implements ICommStatus
         * @constructor
         * @param {CommStructures.ICommStatus=} [properties] Properties to set
         */
        function CommStatus(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * CommStatus port.
         * @member {string} port
         * @memberof CommStructures.CommStatus
         * @instance
         */
        CommStatus.prototype.port = "";

        /**
         * CommStatus baud.
         * @member {number} baud
         * @memberof CommStructures.CommStatus
         * @instance
         */
        CommStatus.prototype.baud = 0;

        /**
         * CommStatus connected.
         * @member {boolean} connected
         * @memberof CommStructures.CommStatus
         * @instance
         */
        CommStatus.prototype.connected = false;

        /**
         * Creates a new CommStatus instance using the specified properties.
         * @function create
         * @memberof CommStructures.CommStatus
         * @static
         * @param {CommStructures.ICommStatus=} [properties] Properties to set
         * @returns {CommStructures.CommStatus} CommStatus instance
         */
        CommStatus.create = function create(properties) {
            return new CommStatus(properties);
        };

        /**
         * Encodes the specified CommStatus message. Does not implicitly {@link CommStructures.CommStatus.verify|verify} messages.
         * @function encode
         * @memberof CommStructures.CommStatus
         * @static
         * @param {CommStructures.ICommStatus} message CommStatus message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CommStatus.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.port != null && Object.hasOwnProperty.call(message, "port"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.port);
            if (message.baud != null && Object.hasOwnProperty.call(message, "baud"))
                writer.uint32(/* id 2, wireType 0 =*/16).int32(message.baud);
            if (message.connected != null && Object.hasOwnProperty.call(message, "connected"))
                writer.uint32(/* id 3, wireType 0 =*/24).bool(message.connected);
            return writer;
        };

        /**
         * Encodes the specified CommStatus message, length delimited. Does not implicitly {@link CommStructures.CommStatus.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommStructures.CommStatus
         * @static
         * @param {CommStructures.ICommStatus} message CommStatus message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CommStatus.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a CommStatus message from the specified reader or buffer.
         * @function decode
         * @memberof CommStructures.CommStatus
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommStructures.CommStatus} CommStatus
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        CommStatus.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommStructures.CommStatus();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.port = reader.string();
                    break;
                case 2:
                    message.baud = reader.int32();
                    break;
                case 3:
                    message.connected = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a CommStatus message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommStructures.CommStatus
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommStructures.CommStatus} CommStatus
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        CommStatus.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a CommStatus message.
         * @function verify
         * @memberof CommStructures.CommStatus
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        CommStatus.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.port != null && message.hasOwnProperty("port"))
                if (!$util.isString(message.port))
                    return "port: string expected";
            if (message.baud != null && message.hasOwnProperty("baud"))
                if (!$util.isInteger(message.baud))
                    return "baud: integer expected";
            if (message.connected != null && message.hasOwnProperty("connected"))
                if (typeof message.connected !== "boolean")
                    return "connected: boolean expected";
            return null;
        };

        /**
         * Creates a CommStatus message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommStructures.CommStatus
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommStructures.CommStatus} CommStatus
         */
        CommStatus.fromObject = function fromObject(object) {
            if (object instanceof $root.CommStructures.CommStatus)
                return object;
            var message = new $root.CommStructures.CommStatus();
            if (object.port != null)
                message.port = String(object.port);
            if (object.baud != null)
                message.baud = object.baud | 0;
            if (object.connected != null)
                message.connected = Boolean(object.connected);
            return message;
        };

        /**
         * Creates a plain object from a CommStatus message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommStructures.CommStatus
         * @static
         * @param {CommStructures.CommStatus} message CommStatus
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CommStatus.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.port = "";
                object.baud = 0;
                object.connected = false;
            }
            if (message.port != null && message.hasOwnProperty("port"))
                object.port = message.port;
            if (message.baud != null && message.hasOwnProperty("baud"))
                object.baud = message.baud;
            if (message.connected != null && message.hasOwnProperty("connected"))
                object.connected = message.connected;
            return object;
        };

        /**
         * Converts this CommStatus to JSON.
         * @function toJSON
         * @memberof CommStructures.CommStatus
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        CommStatus.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return CommStatus;
    })();

    CommStructures.InitComm = (function() {

        /**
         * Properties of an InitComm.
         * @memberof CommStructures
         * @interface IInitComm
         * @property {string|null} [port] InitComm port
         * @property {number|null} [baud] InitComm baud
         */

        /**
         * Constructs a new InitComm.
         * @memberof CommStructures
         * @classdesc Represents an InitComm.
         * @implements IInitComm
         * @constructor
         * @param {CommStructures.IInitComm=} [properties] Properties to set
         */
        function InitComm(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * InitComm port.
         * @member {string} port
         * @memberof CommStructures.InitComm
         * @instance
         */
        InitComm.prototype.port = "";

        /**
         * InitComm baud.
         * @member {number} baud
         * @memberof CommStructures.InitComm
         * @instance
         */
        InitComm.prototype.baud = 0;

        /**
         * Creates a new InitComm instance using the specified properties.
         * @function create
         * @memberof CommStructures.InitComm
         * @static
         * @param {CommStructures.IInitComm=} [properties] Properties to set
         * @returns {CommStructures.InitComm} InitComm instance
         */
        InitComm.create = function create(properties) {
            return new InitComm(properties);
        };

        /**
         * Encodes the specified InitComm message. Does not implicitly {@link CommStructures.InitComm.verify|verify} messages.
         * @function encode
         * @memberof CommStructures.InitComm
         * @static
         * @param {CommStructures.IInitComm} message InitComm message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InitComm.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.port != null && Object.hasOwnProperty.call(message, "port"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.port);
            if (message.baud != null && Object.hasOwnProperty.call(message, "baud"))
                writer.uint32(/* id 2, wireType 0 =*/16).int32(message.baud);
            return writer;
        };

        /**
         * Encodes the specified InitComm message, length delimited. Does not implicitly {@link CommStructures.InitComm.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommStructures.InitComm
         * @static
         * @param {CommStructures.IInitComm} message InitComm message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        InitComm.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an InitComm message from the specified reader or buffer.
         * @function decode
         * @memberof CommStructures.InitComm
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommStructures.InitComm} InitComm
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        InitComm.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommStructures.InitComm();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.port = reader.string();
                    break;
                case 2:
                    message.baud = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an InitComm message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommStructures.InitComm
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommStructures.InitComm} InitComm
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        InitComm.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an InitComm message.
         * @function verify
         * @memberof CommStructures.InitComm
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        InitComm.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.port != null && message.hasOwnProperty("port"))
                if (!$util.isString(message.port))
                    return "port: string expected";
            if (message.baud != null && message.hasOwnProperty("baud"))
                if (!$util.isInteger(message.baud))
                    return "baud: integer expected";
            return null;
        };

        /**
         * Creates an InitComm message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommStructures.InitComm
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommStructures.InitComm} InitComm
         */
        InitComm.fromObject = function fromObject(object) {
            if (object instanceof $root.CommStructures.InitComm)
                return object;
            var message = new $root.CommStructures.InitComm();
            if (object.port != null)
                message.port = String(object.port);
            if (object.baud != null)
                message.baud = object.baud | 0;
            return message;
        };

        /**
         * Creates a plain object from an InitComm message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommStructures.InitComm
         * @static
         * @param {CommStructures.InitComm} message InitComm
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        InitComm.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.port = "";
                object.baud = 0;
            }
            if (message.port != null && message.hasOwnProperty("port"))
                object.port = message.port;
            if (message.baud != null && message.hasOwnProperty("baud"))
                object.baud = message.baud;
            return object;
        };

        /**
         * Converts this InitComm to JSON.
         * @function toJSON
         * @memberof CommStructures.InitComm
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        InitComm.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return InitComm;
    })();

    CommStructures.CommOptions = (function() {

        /**
         * Properties of a CommOptions.
         * @memberof CommStructures
         * @interface ICommOptions
         * @property {CommStructures.IPorts|null} [ports] CommOptions ports
         * @property {CommStructures.IBauds|null} [bauds] CommOptions bauds
         */

        /**
         * Constructs a new CommOptions.
         * @memberof CommStructures
         * @classdesc Represents a CommOptions.
         * @implements ICommOptions
         * @constructor
         * @param {CommStructures.ICommOptions=} [properties] Properties to set
         */
        function CommOptions(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * CommOptions ports.
         * @member {CommStructures.IPorts|null|undefined} ports
         * @memberof CommStructures.CommOptions
         * @instance
         */
        CommOptions.prototype.ports = null;

        /**
         * CommOptions bauds.
         * @member {CommStructures.IBauds|null|undefined} bauds
         * @memberof CommStructures.CommOptions
         * @instance
         */
        CommOptions.prototype.bauds = null;

        /**
         * Creates a new CommOptions instance using the specified properties.
         * @function create
         * @memberof CommStructures.CommOptions
         * @static
         * @param {CommStructures.ICommOptions=} [properties] Properties to set
         * @returns {CommStructures.CommOptions} CommOptions instance
         */
        CommOptions.create = function create(properties) {
            return new CommOptions(properties);
        };

        /**
         * Encodes the specified CommOptions message. Does not implicitly {@link CommStructures.CommOptions.verify|verify} messages.
         * @function encode
         * @memberof CommStructures.CommOptions
         * @static
         * @param {CommStructures.ICommOptions} message CommOptions message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CommOptions.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ports != null && Object.hasOwnProperty.call(message, "ports"))
                $root.CommStructures.Ports.encode(message.ports, writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.bauds != null && Object.hasOwnProperty.call(message, "bauds"))
                $root.CommStructures.Bauds.encode(message.bauds, writer.uint32(/* id 2, wireType 2 =*/18).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified CommOptions message, length delimited. Does not implicitly {@link CommStructures.CommOptions.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommStructures.CommOptions
         * @static
         * @param {CommStructures.ICommOptions} message CommOptions message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CommOptions.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a CommOptions message from the specified reader or buffer.
         * @function decode
         * @memberof CommStructures.CommOptions
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommStructures.CommOptions} CommOptions
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        CommOptions.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommStructures.CommOptions();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.ports = $root.CommStructures.Ports.decode(reader, reader.uint32());
                    break;
                case 2:
                    message.bauds = $root.CommStructures.Bauds.decode(reader, reader.uint32());
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a CommOptions message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommStructures.CommOptions
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommStructures.CommOptions} CommOptions
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        CommOptions.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a CommOptions message.
         * @function verify
         * @memberof CommStructures.CommOptions
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        CommOptions.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ports != null && message.hasOwnProperty("ports")) {
                var error = $root.CommStructures.Ports.verify(message.ports);
                if (error)
                    return "ports." + error;
            }
            if (message.bauds != null && message.hasOwnProperty("bauds")) {
                var error = $root.CommStructures.Bauds.verify(message.bauds);
                if (error)
                    return "bauds." + error;
            }
            return null;
        };

        /**
         * Creates a CommOptions message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommStructures.CommOptions
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommStructures.CommOptions} CommOptions
         */
        CommOptions.fromObject = function fromObject(object) {
            if (object instanceof $root.CommStructures.CommOptions)
                return object;
            var message = new $root.CommStructures.CommOptions();
            if (object.ports != null) {
                if (typeof object.ports !== "object")
                    throw TypeError(".CommStructures.CommOptions.ports: object expected");
                message.ports = $root.CommStructures.Ports.fromObject(object.ports);
            }
            if (object.bauds != null) {
                if (typeof object.bauds !== "object")
                    throw TypeError(".CommStructures.CommOptions.bauds: object expected");
                message.bauds = $root.CommStructures.Bauds.fromObject(object.bauds);
            }
            return message;
        };

        /**
         * Creates a plain object from a CommOptions message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommStructures.CommOptions
         * @static
         * @param {CommStructures.CommOptions} message CommOptions
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CommOptions.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.ports = null;
                object.bauds = null;
            }
            if (message.ports != null && message.hasOwnProperty("ports"))
                object.ports = $root.CommStructures.Ports.toObject(message.ports, options);
            if (message.bauds != null && message.hasOwnProperty("bauds"))
                object.bauds = $root.CommStructures.Bauds.toObject(message.bauds, options);
            return object;
        };

        /**
         * Converts this CommOptions to JSON.
         * @function toJSON
         * @memberof CommStructures.CommOptions
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        CommOptions.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return CommOptions;
    })();

    CommStructures.CommMessage = (function() {

        /**
         * Properties of a CommMessage.
         * @memberof CommStructures
         * @interface ICommMessage
         * @property {string|null} [message] CommMessage message
         */

        /**
         * Constructs a new CommMessage.
         * @memberof CommStructures
         * @classdesc Represents a CommMessage.
         * @implements ICommMessage
         * @constructor
         * @param {CommStructures.ICommMessage=} [properties] Properties to set
         */
        function CommMessage(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * CommMessage message.
         * @member {string} message
         * @memberof CommStructures.CommMessage
         * @instance
         */
        CommMessage.prototype.message = "";

        /**
         * Creates a new CommMessage instance using the specified properties.
         * @function create
         * @memberof CommStructures.CommMessage
         * @static
         * @param {CommStructures.ICommMessage=} [properties] Properties to set
         * @returns {CommStructures.CommMessage} CommMessage instance
         */
        CommMessage.create = function create(properties) {
            return new CommMessage(properties);
        };

        /**
         * Encodes the specified CommMessage message. Does not implicitly {@link CommStructures.CommMessage.verify|verify} messages.
         * @function encode
         * @memberof CommStructures.CommMessage
         * @static
         * @param {CommStructures.ICommMessage} message CommMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CommMessage.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.message != null && Object.hasOwnProperty.call(message, "message"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.message);
            return writer;
        };

        /**
         * Encodes the specified CommMessage message, length delimited. Does not implicitly {@link CommStructures.CommMessage.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommStructures.CommMessage
         * @static
         * @param {CommStructures.ICommMessage} message CommMessage message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        CommMessage.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a CommMessage message from the specified reader or buffer.
         * @function decode
         * @memberof CommStructures.CommMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommStructures.CommMessage} CommMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        CommMessage.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommStructures.CommMessage();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.message = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a CommMessage message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommStructures.CommMessage
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommStructures.CommMessage} CommMessage
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        CommMessage.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a CommMessage message.
         * @function verify
         * @memberof CommStructures.CommMessage
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        CommMessage.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.message != null && message.hasOwnProperty("message"))
                if (!$util.isString(message.message))
                    return "message: string expected";
            return null;
        };

        /**
         * Creates a CommMessage message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommStructures.CommMessage
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommStructures.CommMessage} CommMessage
         */
        CommMessage.fromObject = function fromObject(object) {
            if (object instanceof $root.CommStructures.CommMessage)
                return object;
            var message = new $root.CommStructures.CommMessage();
            if (object.message != null)
                message.message = String(object.message);
            return message;
        };

        /**
         * Creates a plain object from a CommMessage message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommStructures.CommMessage
         * @static
         * @param {CommStructures.CommMessage} message CommMessage
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        CommMessage.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.message = "";
            if (message.message != null && message.hasOwnProperty("message"))
                object.message = message.message;
            return object;
        };

        /**
         * Converts this CommMessage to JSON.
         * @function toJSON
         * @memberof CommStructures.CommMessage
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        CommMessage.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return CommMessage;
    })();

    CommStructures.WrittenBytes = (function() {

        /**
         * Properties of a WrittenBytes.
         * @memberof CommStructures
         * @interface IWrittenBytes
         * @property {number|null} [bytes] WrittenBytes bytes
         */

        /**
         * Constructs a new WrittenBytes.
         * @memberof CommStructures
         * @classdesc Represents a WrittenBytes.
         * @implements IWrittenBytes
         * @constructor
         * @param {CommStructures.IWrittenBytes=} [properties] Properties to set
         */
        function WrittenBytes(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * WrittenBytes bytes.
         * @member {number} bytes
         * @memberof CommStructures.WrittenBytes
         * @instance
         */
        WrittenBytes.prototype.bytes = 0;

        /**
         * Creates a new WrittenBytes instance using the specified properties.
         * @function create
         * @memberof CommStructures.WrittenBytes
         * @static
         * @param {CommStructures.IWrittenBytes=} [properties] Properties to set
         * @returns {CommStructures.WrittenBytes} WrittenBytes instance
         */
        WrittenBytes.create = function create(properties) {
            return new WrittenBytes(properties);
        };

        /**
         * Encodes the specified WrittenBytes message. Does not implicitly {@link CommStructures.WrittenBytes.verify|verify} messages.
         * @function encode
         * @memberof CommStructures.WrittenBytes
         * @static
         * @param {CommStructures.IWrittenBytes} message WrittenBytes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        WrittenBytes.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.bytes != null && Object.hasOwnProperty.call(message, "bytes"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.bytes);
            return writer;
        };

        /**
         * Encodes the specified WrittenBytes message, length delimited. Does not implicitly {@link CommStructures.WrittenBytes.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommStructures.WrittenBytes
         * @static
         * @param {CommStructures.IWrittenBytes} message WrittenBytes message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        WrittenBytes.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a WrittenBytes message from the specified reader or buffer.
         * @function decode
         * @memberof CommStructures.WrittenBytes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommStructures.WrittenBytes} WrittenBytes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        WrittenBytes.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommStructures.WrittenBytes();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.bytes = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a WrittenBytes message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommStructures.WrittenBytes
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommStructures.WrittenBytes} WrittenBytes
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        WrittenBytes.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a WrittenBytes message.
         * @function verify
         * @memberof CommStructures.WrittenBytes
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        WrittenBytes.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.bytes != null && message.hasOwnProperty("bytes"))
                if (!$util.isInteger(message.bytes))
                    return "bytes: integer expected";
            return null;
        };

        /**
         * Creates a WrittenBytes message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommStructures.WrittenBytes
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommStructures.WrittenBytes} WrittenBytes
         */
        WrittenBytes.fromObject = function fromObject(object) {
            if (object instanceof $root.CommStructures.WrittenBytes)
                return object;
            var message = new $root.CommStructures.WrittenBytes();
            if (object.bytes != null)
                message.bytes = object.bytes | 0;
            return message;
        };

        /**
         * Creates a plain object from a WrittenBytes message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommStructures.WrittenBytes
         * @static
         * @param {CommStructures.WrittenBytes} message WrittenBytes
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        WrittenBytes.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.bytes = 0;
            if (message.bytes != null && message.hasOwnProperty("bytes"))
                object.bytes = message.bytes;
            return object;
        };

        /**
         * Converts this WrittenBytes to JSON.
         * @function toJSON
         * @memberof CommStructures.WrittenBytes
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        WrittenBytes.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return WrittenBytes;
    })();

    CommStructures.Port = (function() {

        /**
         * Properties of a Port.
         * @memberof CommStructures
         * @interface IPort
         * @property {string|null} [address] Port address
         */

        /**
         * Constructs a new Port.
         * @memberof CommStructures
         * @classdesc Represents a Port.
         * @implements IPort
         * @constructor
         * @param {CommStructures.IPort=} [properties] Properties to set
         */
        function Port(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Port address.
         * @member {string} address
         * @memberof CommStructures.Port
         * @instance
         */
        Port.prototype.address = "";

        /**
         * Creates a new Port instance using the specified properties.
         * @function create
         * @memberof CommStructures.Port
         * @static
         * @param {CommStructures.IPort=} [properties] Properties to set
         * @returns {CommStructures.Port} Port instance
         */
        Port.create = function create(properties) {
            return new Port(properties);
        };

        /**
         * Encodes the specified Port message. Does not implicitly {@link CommStructures.Port.verify|verify} messages.
         * @function encode
         * @memberof CommStructures.Port
         * @static
         * @param {CommStructures.IPort} message Port message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Port.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.address != null && Object.hasOwnProperty.call(message, "address"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.address);
            return writer;
        };

        /**
         * Encodes the specified Port message, length delimited. Does not implicitly {@link CommStructures.Port.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommStructures.Port
         * @static
         * @param {CommStructures.IPort} message Port message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Port.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Port message from the specified reader or buffer.
         * @function decode
         * @memberof CommStructures.Port
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommStructures.Port} Port
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Port.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommStructures.Port();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.address = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Port message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommStructures.Port
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommStructures.Port} Port
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Port.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Port message.
         * @function verify
         * @memberof CommStructures.Port
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Port.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.address != null && message.hasOwnProperty("address"))
                if (!$util.isString(message.address))
                    return "address: string expected";
            return null;
        };

        /**
         * Creates a Port message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommStructures.Port
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommStructures.Port} Port
         */
        Port.fromObject = function fromObject(object) {
            if (object instanceof $root.CommStructures.Port)
                return object;
            var message = new $root.CommStructures.Port();
            if (object.address != null)
                message.address = String(object.address);
            return message;
        };

        /**
         * Creates a plain object from a Port message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommStructures.Port
         * @static
         * @param {CommStructures.Port} message Port
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Port.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.address = "";
            if (message.address != null && message.hasOwnProperty("address"))
                object.address = message.address;
            return object;
        };

        /**
         * Converts this Port to JSON.
         * @function toJSON
         * @memberof CommStructures.Port
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Port.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Port;
    })();

    CommStructures.Ports = (function() {

        /**
         * Properties of a Ports.
         * @memberof CommStructures
         * @interface IPorts
         * @property {Array.<CommStructures.IPort>|null} [ports] Ports ports
         */

        /**
         * Constructs a new Ports.
         * @memberof CommStructures
         * @classdesc Represents a Ports.
         * @implements IPorts
         * @constructor
         * @param {CommStructures.IPorts=} [properties] Properties to set
         */
        function Ports(properties) {
            this.ports = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Ports ports.
         * @member {Array.<CommStructures.IPort>} ports
         * @memberof CommStructures.Ports
         * @instance
         */
        Ports.prototype.ports = $util.emptyArray;

        /**
         * Creates a new Ports instance using the specified properties.
         * @function create
         * @memberof CommStructures.Ports
         * @static
         * @param {CommStructures.IPorts=} [properties] Properties to set
         * @returns {CommStructures.Ports} Ports instance
         */
        Ports.create = function create(properties) {
            return new Ports(properties);
        };

        /**
         * Encodes the specified Ports message. Does not implicitly {@link CommStructures.Ports.verify|verify} messages.
         * @function encode
         * @memberof CommStructures.Ports
         * @static
         * @param {CommStructures.IPorts} message Ports message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Ports.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.ports != null && message.ports.length)
                for (var i = 0; i < message.ports.length; ++i)
                    $root.CommStructures.Port.encode(message.ports[i], writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified Ports message, length delimited. Does not implicitly {@link CommStructures.Ports.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommStructures.Ports
         * @static
         * @param {CommStructures.IPorts} message Ports message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Ports.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Ports message from the specified reader or buffer.
         * @function decode
         * @memberof CommStructures.Ports
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommStructures.Ports} Ports
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Ports.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommStructures.Ports();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    if (!(message.ports && message.ports.length))
                        message.ports = [];
                    message.ports.push($root.CommStructures.Port.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Ports message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommStructures.Ports
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommStructures.Ports} Ports
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Ports.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Ports message.
         * @function verify
         * @memberof CommStructures.Ports
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Ports.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.ports != null && message.hasOwnProperty("ports")) {
                if (!Array.isArray(message.ports))
                    return "ports: array expected";
                for (var i = 0; i < message.ports.length; ++i) {
                    var error = $root.CommStructures.Port.verify(message.ports[i]);
                    if (error)
                        return "ports." + error;
                }
            }
            return null;
        };

        /**
         * Creates a Ports message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommStructures.Ports
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommStructures.Ports} Ports
         */
        Ports.fromObject = function fromObject(object) {
            if (object instanceof $root.CommStructures.Ports)
                return object;
            var message = new $root.CommStructures.Ports();
            if (object.ports) {
                if (!Array.isArray(object.ports))
                    throw TypeError(".CommStructures.Ports.ports: array expected");
                message.ports = [];
                for (var i = 0; i < object.ports.length; ++i) {
                    if (typeof object.ports[i] !== "object")
                        throw TypeError(".CommStructures.Ports.ports: object expected");
                    message.ports[i] = $root.CommStructures.Port.fromObject(object.ports[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a Ports message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommStructures.Ports
         * @static
         * @param {CommStructures.Ports} message Ports
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Ports.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.ports = [];
            if (message.ports && message.ports.length) {
                object.ports = [];
                for (var j = 0; j < message.ports.length; ++j)
                    object.ports[j] = $root.CommStructures.Port.toObject(message.ports[j], options);
            }
            return object;
        };

        /**
         * Converts this Ports to JSON.
         * @function toJSON
         * @memberof CommStructures.Ports
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Ports.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Ports;
    })();

    CommStructures.Baud = (function() {

        /**
         * Properties of a Baud.
         * @memberof CommStructures
         * @interface IBaud
         * @property {number|null} [speed] Baud speed
         */

        /**
         * Constructs a new Baud.
         * @memberof CommStructures
         * @classdesc Represents a Baud.
         * @implements IBaud
         * @constructor
         * @param {CommStructures.IBaud=} [properties] Properties to set
         */
        function Baud(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Baud speed.
         * @member {number} speed
         * @memberof CommStructures.Baud
         * @instance
         */
        Baud.prototype.speed = 0;

        /**
         * Creates a new Baud instance using the specified properties.
         * @function create
         * @memberof CommStructures.Baud
         * @static
         * @param {CommStructures.IBaud=} [properties] Properties to set
         * @returns {CommStructures.Baud} Baud instance
         */
        Baud.create = function create(properties) {
            return new Baud(properties);
        };

        /**
         * Encodes the specified Baud message. Does not implicitly {@link CommStructures.Baud.verify|verify} messages.
         * @function encode
         * @memberof CommStructures.Baud
         * @static
         * @param {CommStructures.IBaud} message Baud message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Baud.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.speed != null && Object.hasOwnProperty.call(message, "speed"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.speed);
            return writer;
        };

        /**
         * Encodes the specified Baud message, length delimited. Does not implicitly {@link CommStructures.Baud.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommStructures.Baud
         * @static
         * @param {CommStructures.IBaud} message Baud message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Baud.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Baud message from the specified reader or buffer.
         * @function decode
         * @memberof CommStructures.Baud
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommStructures.Baud} Baud
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Baud.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommStructures.Baud();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.speed = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Baud message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommStructures.Baud
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommStructures.Baud} Baud
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Baud.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Baud message.
         * @function verify
         * @memberof CommStructures.Baud
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Baud.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.speed != null && message.hasOwnProperty("speed"))
                if (!$util.isInteger(message.speed))
                    return "speed: integer expected";
            return null;
        };

        /**
         * Creates a Baud message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommStructures.Baud
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommStructures.Baud} Baud
         */
        Baud.fromObject = function fromObject(object) {
            if (object instanceof $root.CommStructures.Baud)
                return object;
            var message = new $root.CommStructures.Baud();
            if (object.speed != null)
                message.speed = object.speed | 0;
            return message;
        };

        /**
         * Creates a plain object from a Baud message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommStructures.Baud
         * @static
         * @param {CommStructures.Baud} message Baud
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Baud.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.speed = 0;
            if (message.speed != null && message.hasOwnProperty("speed"))
                object.speed = message.speed;
            return object;
        };

        /**
         * Converts this Baud to JSON.
         * @function toJSON
         * @memberof CommStructures.Baud
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Baud.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Baud;
    })();

    CommStructures.Bauds = (function() {

        /**
         * Properties of a Bauds.
         * @memberof CommStructures
         * @interface IBauds
         * @property {Array.<CommStructures.IBaud>|null} [bauds] Bauds bauds
         */

        /**
         * Constructs a new Bauds.
         * @memberof CommStructures
         * @classdesc Represents a Bauds.
         * @implements IBauds
         * @constructor
         * @param {CommStructures.IBauds=} [properties] Properties to set
         */
        function Bauds(properties) {
            this.bauds = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Bauds bauds.
         * @member {Array.<CommStructures.IBaud>} bauds
         * @memberof CommStructures.Bauds
         * @instance
         */
        Bauds.prototype.bauds = $util.emptyArray;

        /**
         * Creates a new Bauds instance using the specified properties.
         * @function create
         * @memberof CommStructures.Bauds
         * @static
         * @param {CommStructures.IBauds=} [properties] Properties to set
         * @returns {CommStructures.Bauds} Bauds instance
         */
        Bauds.create = function create(properties) {
            return new Bauds(properties);
        };

        /**
         * Encodes the specified Bauds message. Does not implicitly {@link CommStructures.Bauds.verify|verify} messages.
         * @function encode
         * @memberof CommStructures.Bauds
         * @static
         * @param {CommStructures.IBauds} message Bauds message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Bauds.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.bauds != null && message.bauds.length)
                for (var i = 0; i < message.bauds.length; ++i)
                    $root.CommStructures.Baud.encode(message.bauds[i], writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            return writer;
        };

        /**
         * Encodes the specified Bauds message, length delimited. Does not implicitly {@link CommStructures.Bauds.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommStructures.Bauds
         * @static
         * @param {CommStructures.IBauds} message Bauds message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Bauds.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Bauds message from the specified reader or buffer.
         * @function decode
         * @memberof CommStructures.Bauds
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommStructures.Bauds} Bauds
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Bauds.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommStructures.Bauds();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    if (!(message.bauds && message.bauds.length))
                        message.bauds = [];
                    message.bauds.push($root.CommStructures.Baud.decode(reader, reader.uint32()));
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Bauds message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommStructures.Bauds
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommStructures.Bauds} Bauds
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Bauds.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Bauds message.
         * @function verify
         * @memberof CommStructures.Bauds
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Bauds.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.bauds != null && message.hasOwnProperty("bauds")) {
                if (!Array.isArray(message.bauds))
                    return "bauds: array expected";
                for (var i = 0; i < message.bauds.length; ++i) {
                    var error = $root.CommStructures.Baud.verify(message.bauds[i]);
                    if (error)
                        return "bauds." + error;
                }
            }
            return null;
        };

        /**
         * Creates a Bauds message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommStructures.Bauds
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommStructures.Bauds} Bauds
         */
        Bauds.fromObject = function fromObject(object) {
            if (object instanceof $root.CommStructures.Bauds)
                return object;
            var message = new $root.CommStructures.Bauds();
            if (object.bauds) {
                if (!Array.isArray(object.bauds))
                    throw TypeError(".CommStructures.Bauds.bauds: array expected");
                message.bauds = [];
                for (var i = 0; i < object.bauds.length; ++i) {
                    if (typeof object.bauds[i] !== "object")
                        throw TypeError(".CommStructures.Bauds.bauds: object expected");
                    message.bauds[i] = $root.CommStructures.Baud.fromObject(object.bauds[i]);
                }
            }
            return message;
        };

        /**
         * Creates a plain object from a Bauds message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommStructures.Bauds
         * @static
         * @param {CommStructures.Bauds} message Bauds
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Bauds.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.bauds = [];
            if (message.bauds && message.bauds.length) {
                object.bauds = [];
                for (var j = 0; j < message.bauds.length; ++j)
                    object.bauds[j] = $root.CommStructures.Baud.toObject(message.bauds[j], options);
            }
            return object;
        };

        /**
         * Converts this Bauds to JSON.
         * @function toJSON
         * @memberof CommStructures.Bauds
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Bauds.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Bauds;
    })();

    return CommStructures;
})();

module.exports = $root;
