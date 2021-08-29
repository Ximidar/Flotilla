/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.TemperatureStructures = (function() {

    /**
     * Namespace TemperatureStructures.
     * @exports TemperatureStructures
     * @namespace
     */
    var TemperatureStructures = {};

    TemperatureStructures.Temperature = (function() {

        /**
         * Properties of a Temperature.
         * @memberof TemperatureStructures
         * @interface ITemperature
         * @property {string|null} [Tool] Temperature Tool
         * @property {number|null} [Temp] Temperature Temp
         * @property {number|null} [Target] Temperature Target
         */

        /**
         * Constructs a new Temperature.
         * @memberof TemperatureStructures
         * @classdesc Represents a Temperature.
         * @implements ITemperature
         * @constructor
         * @param {TemperatureStructures.ITemperature=} [properties] Properties to set
         */
        function Temperature(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Temperature Tool.
         * @member {string} Tool
         * @memberof TemperatureStructures.Temperature
         * @instance
         */
        Temperature.prototype.Tool = "";

        /**
         * Temperature Temp.
         * @member {number} Temp
         * @memberof TemperatureStructures.Temperature
         * @instance
         */
        Temperature.prototype.Temp = 0;

        /**
         * Temperature Target.
         * @member {number} Target
         * @memberof TemperatureStructures.Temperature
         * @instance
         */
        Temperature.prototype.Target = 0;

        /**
         * Creates a new Temperature instance using the specified properties.
         * @function create
         * @memberof TemperatureStructures.Temperature
         * @static
         * @param {TemperatureStructures.ITemperature=} [properties] Properties to set
         * @returns {TemperatureStructures.Temperature} Temperature instance
         */
        Temperature.create = function create(properties) {
            return new Temperature(properties);
        };

        /**
         * Encodes the specified Temperature message. Does not implicitly {@link TemperatureStructures.Temperature.verify|verify} messages.
         * @function encode
         * @memberof TemperatureStructures.Temperature
         * @static
         * @param {TemperatureStructures.ITemperature} message Temperature message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Temperature.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Tool != null && Object.hasOwnProperty.call(message, "Tool"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.Tool);
            if (message.Temp != null && Object.hasOwnProperty.call(message, "Temp"))
                writer.uint32(/* id 2, wireType 0 =*/16).int32(message.Temp);
            if (message.Target != null && Object.hasOwnProperty.call(message, "Target"))
                writer.uint32(/* id 3, wireType 0 =*/24).int32(message.Target);
            return writer;
        };

        /**
         * Encodes the specified Temperature message, length delimited. Does not implicitly {@link TemperatureStructures.Temperature.verify|verify} messages.
         * @function encodeDelimited
         * @memberof TemperatureStructures.Temperature
         * @static
         * @param {TemperatureStructures.ITemperature} message Temperature message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Temperature.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes a Temperature message from the specified reader or buffer.
         * @function decode
         * @memberof TemperatureStructures.Temperature
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {TemperatureStructures.Temperature} Temperature
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Temperature.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.TemperatureStructures.Temperature();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.Tool = reader.string();
                    break;
                case 2:
                    message.Temp = reader.int32();
                    break;
                case 3:
                    message.Target = reader.int32();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes a Temperature message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof TemperatureStructures.Temperature
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {TemperatureStructures.Temperature} Temperature
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Temperature.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies a Temperature message.
         * @function verify
         * @memberof TemperatureStructures.Temperature
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Temperature.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Tool != null && message.hasOwnProperty("Tool"))
                if (!$util.isString(message.Tool))
                    return "Tool: string expected";
            if (message.Temp != null && message.hasOwnProperty("Temp"))
                if (!$util.isInteger(message.Temp))
                    return "Temp: integer expected";
            if (message.Target != null && message.hasOwnProperty("Target"))
                if (!$util.isInteger(message.Target))
                    return "Target: integer expected";
            return null;
        };

        /**
         * Creates a Temperature message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof TemperatureStructures.Temperature
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {TemperatureStructures.Temperature} Temperature
         */
        Temperature.fromObject = function fromObject(object) {
            if (object instanceof $root.TemperatureStructures.Temperature)
                return object;
            var message = new $root.TemperatureStructures.Temperature();
            if (object.Tool != null)
                message.Tool = String(object.Tool);
            if (object.Temp != null)
                message.Temp = object.Temp | 0;
            if (object.Target != null)
                message.Target = object.Target | 0;
            return message;
        };

        /**
         * Creates a plain object from a Temperature message. Also converts values to other types if specified.
         * @function toObject
         * @memberof TemperatureStructures.Temperature
         * @static
         * @param {TemperatureStructures.Temperature} message Temperature
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Temperature.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults) {
                object.Tool = "";
                object.Temp = 0;
                object.Target = 0;
            }
            if (message.Tool != null && message.hasOwnProperty("Tool"))
                object.Tool = message.Tool;
            if (message.Temp != null && message.hasOwnProperty("Temp"))
                object.Temp = message.Temp;
            if (message.Target != null && message.hasOwnProperty("Target"))
                object.Target = message.Target;
            return object;
        };

        /**
         * Converts this Temperature to JSON.
         * @function toJSON
         * @memberof TemperatureStructures.Temperature
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Temperature.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Temperature;
    })();

    return TemperatureStructures;
})();

module.exports = $root;
