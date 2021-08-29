/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.CommRelayStructures = (function() {

    /**
     * Namespace CommRelayStructures.
     * @exports CommRelayStructures
     * @namespace
     */
    var CommRelayStructures = {};

    CommRelayStructures.Line = (function() {

        /**
         * Properties of a Line.
         * @memberof CommRelayStructures
         * @interface ILine
         * @property {string|null} [Line] Line Line
         * @property {number|Long|null} [LineNumber] Line LineNumber
         * @property {boolean|null} [KnownNumber] Line KnownNumber
         */

        /**
         * Constructs a new Line.
         * @memberof CommRelayStructures
         * @classdesc Represents a Line.
         * @implements ILine
         * @constructor
         * @param {CommRelayStructures.ILine=} [properties] Properties to set
         */
        function Line(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Line Line.
         * @member {string} Line
         * @memberof CommRelayStructures.Line
         * @instance
         */
        Line.prototype.Line = "";

        /**
         * Line LineNumber.
         * @member {number|Long} LineNumber
         * @memberof CommRelayStructures.Line
         * @instance
         */
        Line.prototype.LineNumber = $util.Long ? $util.Long.fromBits(0,0,true) : 0;

        /**
         * Line KnownNumber.
         * @member {boolean} KnownNumber
         * @memberof CommRelayStructures.Line
         * @instance
         */
        Line.prototype.KnownNumber = false;

        /**
         * Creates a new Line instance using the specified properties.
         * @function create
         * @memberof CommRelayStructures.Line
         * @static
         * @param {CommRelayStructures.ILine=} [properties] Properties to set
         * @returns {CommRelayStructures.Line} Line instance
         */
        Line.create = function create(properties) {
            return new Line(properties);
        };

        /**
         * Encodes the specified Line message. Does not implicitly {@link CommRelayStructures.Line.verify|verify} messages.
         * @function encode
         * @memberof CommRelayStructures.Line
         * @static
         * @param {CommRelayStructures.ILine} message Line message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Line.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Line != null && Object.hasOwnProperty.call(message, "Line"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.Line);
            if (message.LineNumber != null && Object.hasOwnProperty.call(message, "LineNumber"))
                writer.uint32(/* id 2, wireType 0 =*/16).uint64(message.LineNumber);
            if (message.KnownNumber != null && Object.hasOwnProperty.call(message, "KnownNumber"))
                writer.uint32(/* id 3, wireType 0 =*/24).bool(message.KnownNumber);
            return writer;
        };

        /**
         * Encodes the specified Line message, length delimited. Does not implicitly {@link CommRelayStructures.Line.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommRelayStructures.Line
         * @static
         * @param {CommRelayStructures.ILine} message Line message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Line.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Line message from the specified reader or buffer.
         * @function decode
         * @memberof CommRelayStructures.Line
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommRelayStructures.Line} Line
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Line.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommRelayStructures.Line();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.Line = reader.string();
                    break;
                case 2:
                    message.LineNumber = reader.uint64();
                    break;
                case 3:
                    message.KnownNumber = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Line message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommRelayStructures.Line
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommRelayStructures.Line} Line
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Line.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Line message.
         * @function verify
         * @memberof CommRelayStructures.Line
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Line.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Line != null && message.hasOwnProperty("Line"))
                if (!$util.isString(message.Line))
                    return "Line: string expected";
            if (message.LineNumber != null && message.hasOwnProperty("LineNumber"))
                if (!$util.isInteger(message.LineNumber) && !(message.LineNumber && $util.isInteger(message.LineNumber.low) && $util.isInteger(message.LineNumber.high)))
                    return "LineNumber: integer|Long expected";
            if (message.KnownNumber != null && message.hasOwnProperty("KnownNumber"))
                if (typeof message.KnownNumber !== "boolean")
                    return "KnownNumber: boolean expected";
            return null;
        };

        /**
         * Creates a Line message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommRelayStructures.Line
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommRelayStructures.Line} Line
         */
        Line.fromObject = function fromObject(object) {
            if (object instanceof $root.CommRelayStructures.Line)
                return object;
            var message = new $root.CommRelayStructures.Line();
            if (object.Line != null)
                message.Line = String(object.Line);
            if (object.LineNumber != null)
                if ($util.Long)
                    (message.LineNumber = $util.Long.fromValue(object.LineNumber)).unsigned = true;
                else if (typeof object.LineNumber === "string")
                    message.LineNumber = parseInt(object.LineNumber, 10);
                else if (typeof object.LineNumber === "number")
                    message.LineNumber = object.LineNumber;
                else if (typeof object.LineNumber === "object")
                    message.LineNumber = new $util.LongBits(object.LineNumber.low >>> 0, object.LineNumber.high >>> 0).toNumber(true);
            if (object.KnownNumber != null)
                message.KnownNumber = Boolean(object.KnownNumber);
            return message;
        };

        /**
         * Creates a plain object from a Line message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommRelayStructures.Line
         * @static
         * @param {CommRelayStructures.Line} message Line
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Line.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.Line = "";
                if ($util.Long) {
                    var long = new $util.Long(0, 0, true);
                    object.LineNumber = options.longs === String ? long.toString() : options.longs === Number ? long.toNumber() : long;
                } else
                    object.LineNumber = options.longs === String ? "0" : 0;
                object.KnownNumber = false;
            }
            if (message.Line != null && message.hasOwnProperty("Line"))
                object.Line = message.Line;
            if (message.LineNumber != null && message.hasOwnProperty("LineNumber"))
                if (typeof message.LineNumber === "number")
                    object.LineNumber = options.longs === String ? String(message.LineNumber) : message.LineNumber;
                else
                    object.LineNumber = options.longs === String ? $util.Long.prototype.toString.call(message.LineNumber) : options.longs === Number ? new $util.LongBits(message.LineNumber.low >>> 0, message.LineNumber.high >>> 0).toNumber(true) : message.LineNumber;
            if (message.KnownNumber != null && message.hasOwnProperty("KnownNumber"))
                object.KnownNumber = message.KnownNumber;
            return object;
        };

        /**
         * Converts this Line to JSON.
         * @function toJSON
         * @memberof CommRelayStructures.Line
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Line.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Line;
    })();

    CommRelayStructures.RequestLines = (function() {

        /**
         * Properties of a RequestLines.
         * @memberof CommRelayStructures
         * @interface IRequestLines
         * @property {number|null} [Amount] RequestLines Amount
         */

        /**
         * Constructs a new RequestLines.
         * @memberof CommRelayStructures
         * @classdesc Represents a RequestLines.
         * @implements IRequestLines
         * @constructor
         * @param {CommRelayStructures.IRequestLines=} [properties] Properties to set
         */
        function RequestLines(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * RequestLines Amount.
         * @member {number} Amount
         * @memberof CommRelayStructures.RequestLines
         * @instance
         */
        RequestLines.prototype.Amount = 0;

        /**
         * Creates a new RequestLines instance using the specified properties.
         * @function create
         * @memberof CommRelayStructures.RequestLines
         * @static
         * @param {CommRelayStructures.IRequestLines=} [properties] Properties to set
         * @returns {CommRelayStructures.RequestLines} RequestLines instance
         */
        RequestLines.create = function create(properties) {
            return new RequestLines(properties);
        };

        /**
         * Encodes the specified RequestLines message. Does not implicitly {@link CommRelayStructures.RequestLines.verify|verify} messages.
         * @function encode
         * @memberof CommRelayStructures.RequestLines
         * @static
         * @param {CommRelayStructures.IRequestLines} message RequestLines message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RequestLines.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Amount != null && Object.hasOwnProperty.call(message, "Amount"))
                writer.uint32(/* id 1, wireType 0 =*/8).int32(message.Amount);
            return writer;
        };

        /**
         * Encodes the specified RequestLines message, length delimited. Does not implicitly {@link CommRelayStructures.RequestLines.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommRelayStructures.RequestLines
         * @static
         * @param {CommRelayStructures.IRequestLines} message RequestLines message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        RequestLines.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a RequestLines message from the specified reader or buffer.
         * @function decode
         * @memberof CommRelayStructures.RequestLines
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommRelayStructures.RequestLines} RequestLines
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RequestLines.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommRelayStructures.RequestLines();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.Amount = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a RequestLines message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommRelayStructures.RequestLines
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommRelayStructures.RequestLines} RequestLines
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        RequestLines.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a RequestLines message.
         * @function verify
         * @memberof CommRelayStructures.RequestLines
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        RequestLines.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Amount != null && message.hasOwnProperty("Amount"))
                if (!$util.isInteger(message.Amount))
                    return "Amount: integer expected";
            return null;
        };

        /**
         * Creates a RequestLines message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommRelayStructures.RequestLines
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommRelayStructures.RequestLines} RequestLines
         */
        RequestLines.fromObject = function fromObject(object) {
            if (object instanceof $root.CommRelayStructures.RequestLines)
                return object;
            var message = new $root.CommRelayStructures.RequestLines();
            if (object.Amount != null)
                message.Amount = object.Amount | 0;
            return message;
        };

        /**
         * Creates a plain object from a RequestLines message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommRelayStructures.RequestLines
         * @static
         * @param {CommRelayStructures.RequestLines} message RequestLines
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        RequestLines.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.Amount = 0;
            if (message.Amount != null && message.hasOwnProperty("Amount"))
                object.Amount = message.Amount;
            return object;
        };

        /**
         * Converts this RequestLines to JSON.
         * @function toJSON
         * @memberof CommRelayStructures.RequestLines
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        RequestLines.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return RequestLines;
    })();

    CommRelayStructures.ReturnLines = (function() {

        /**
         * Properties of a ReturnLines.
         * @memberof CommRelayStructures
         * @interface IReturnLines
         * @property {Array.<CommRelayStructures.ILine>|null} [Lines] ReturnLines Lines
         * @property {boolean|null} [EOF] ReturnLines EOF
         */

        /**
         * Constructs a new ReturnLines.
         * @memberof CommRelayStructures
         * @classdesc Represents a ReturnLines.
         * @implements IReturnLines
         * @constructor
         * @param {CommRelayStructures.IReturnLines=} [properties] Properties to set
         */
        function ReturnLines(properties) {
            this.Lines = [];
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * ReturnLines Lines.
         * @member {Array.<CommRelayStructures.ILine>} Lines
         * @memberof CommRelayStructures.ReturnLines
         * @instance
         */
        ReturnLines.prototype.Lines = $util.emptyArray;

        /**
         * ReturnLines EOF.
         * @member {boolean} EOF
         * @memberof CommRelayStructures.ReturnLines
         * @instance
         */
        ReturnLines.prototype.EOF = false;

        /**
         * Creates a new ReturnLines instance using the specified properties.
         * @function create
         * @memberof CommRelayStructures.ReturnLines
         * @static
         * @param {CommRelayStructures.IReturnLines=} [properties] Properties to set
         * @returns {CommRelayStructures.ReturnLines} ReturnLines instance
         */
        ReturnLines.create = function create(properties) {
            return new ReturnLines(properties);
        };

        /**
         * Encodes the specified ReturnLines message. Does not implicitly {@link CommRelayStructures.ReturnLines.verify|verify} messages.
         * @function encode
         * @memberof CommRelayStructures.ReturnLines
         * @static
         * @param {CommRelayStructures.IReturnLines} message ReturnLines message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ReturnLines.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Lines != null && message.Lines.length)
                for (var i = 0; i < message.Lines.length; ++i)
                    $root.CommRelayStructures.Line.encode(message.Lines[i], writer.uint32(/* id 1, wireType 2 =*/10).fork()).ldelim();
            if (message.EOF != null && Object.hasOwnProperty.call(message, "EOF"))
                writer.uint32(/* id 2, wireType 0 =*/16).bool(message.EOF);
            return writer;
        };

        /**
         * Encodes the specified ReturnLines message, length delimited. Does not implicitly {@link CommRelayStructures.ReturnLines.verify|verify} messages.
         * @function encodeDelimited
         * @memberof CommRelayStructures.ReturnLines
         * @static
         * @param {CommRelayStructures.IReturnLines} message ReturnLines message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        ReturnLines.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a ReturnLines message from the specified reader or buffer.
         * @function decode
         * @memberof CommRelayStructures.ReturnLines
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {CommRelayStructures.ReturnLines} ReturnLines
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ReturnLines.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.CommRelayStructures.ReturnLines();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    if (!(message.Lines && message.Lines.length))
                        message.Lines = [];
                    message.Lines.push($root.CommRelayStructures.Line.decode(reader, reader.uint32()));
                    break;
                case 2:
                    message.EOF = reader.bool();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a ReturnLines message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof CommRelayStructures.ReturnLines
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {CommRelayStructures.ReturnLines} ReturnLines
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        ReturnLines.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a ReturnLines message.
         * @function verify
         * @memberof CommRelayStructures.ReturnLines
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        ReturnLines.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Lines != null && message.hasOwnProperty("Lines")) {
                if (!Array.isArray(message.Lines))
                    return "Lines: array expected";
                for (var i = 0; i < message.Lines.length; ++i) {
                    var error = $root.CommRelayStructures.Line.verify(message.Lines[i]);
                    if (error)
                        return "Lines." + error;
                }
            }
            if (message.EOF != null && message.hasOwnProperty("EOF"))
                if (typeof message.EOF !== "boolean")
                    return "EOF: boolean expected";
            return null;
        };

        /**
         * Creates a ReturnLines message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof CommRelayStructures.ReturnLines
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {CommRelayStructures.ReturnLines} ReturnLines
         */
        ReturnLines.fromObject = function fromObject(object) {
            if (object instanceof $root.CommRelayStructures.ReturnLines)
                return object;
            var message = new $root.CommRelayStructures.ReturnLines();
            if (object.Lines) {
                if (!Array.isArray(object.Lines))
                    throw TypeError(".CommRelayStructures.ReturnLines.Lines: array expected");
                message.Lines = [];
                for (var i = 0; i < object.Lines.length; ++i) {
                    if (typeof object.Lines[i] !== "object")
                        throw TypeError(".CommRelayStructures.ReturnLines.Lines: object expected");
                    message.Lines[i] = $root.CommRelayStructures.Line.fromObject(object.Lines[i]);
                }
            }
            if (object.EOF != null)
                message.EOF = Boolean(object.EOF);
            return message;
        };

        /**
         * Creates a plain object from a ReturnLines message. Also converts values to other types if specified.
         * @function toObject
         * @memberof CommRelayStructures.ReturnLines
         * @static
         * @param {CommRelayStructures.ReturnLines} message ReturnLines
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        ReturnLines.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.arrays || options.defaults)
                object.Lines = [];
            if (options.defaults)
                object.EOF = false;
            if (message.Lines && message.Lines.length) {
                object.Lines = [];
                for (var j = 0; j < message.Lines.length; ++j)
                    object.Lines[j] = $root.CommRelayStructures.Line.toObject(message.Lines[j], options);
            }
            if (message.EOF != null && message.hasOwnProperty("EOF"))
                object.EOF = message.EOF;
            return object;
        };

        /**
         * Converts this ReturnLines to JSON.
         * @function toJSON
         * @memberof CommRelayStructures.ReturnLines
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        ReturnLines.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return ReturnLines;
    })();

    return CommRelayStructures;
})();

module.exports = $root;
