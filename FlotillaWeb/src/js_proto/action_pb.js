/*eslint-disable block-scoped-var, id-length, no-control-regex, no-magic-numbers, no-prototype-builtins, no-redeclare, no-shadow, no-var, sort-vars*/
"use strict";

var $protobuf = require("protobufjs/minimal");

// Common aliases
var $Reader = $protobuf.Reader, $Writer = $protobuf.Writer, $util = $protobuf.util;

// Exported root namespace
var $root = $protobuf.roots["default"] || ($protobuf.roots["default"] = {});

$root.PlayStructures = (function() {

    /**
     * Namespace PlayStructures.
     * @exports PlayStructures
     * @namespace
     */
    var PlayStructures = {};

    PlayStructures.Action = (function() {

        /**
         * Properties of an Action.
         * @memberof PlayStructures
         * @interface IAction
         * @property {string|null} [Action] Action Action
         */

        /**
         * Constructs a new Action.
         * @memberof PlayStructures
         * @classdesc Represents an Action.
         * @implements IAction
         * @constructor
         * @param {PlayStructures.IAction=} [properties] Properties to set
         */
        function Action(properties) {
            if (properties)
                for (var keys = Object.keys(properties), i = 0; i < keys.length; ++i)
                    if (properties[keys[i]] != null)
                        this[keys[i]] = properties[keys[i]];
        }

        /**
         * Action Action.
         * @member {string} Action
         * @memberof PlayStructures.Action
         * @instance
         */
        Action.prototype.Action = "";

        /**
         * Creates a new Action instance using the specified properties.
         * @function create
         * @memberof PlayStructures.Action
         * @static
         * @param {PlayStructures.IAction=} [properties] Properties to set
         * @returns {PlayStructures.Action} Action instance
         */
        Action.create = function create(properties) {
            return new Action(properties);
        };

        /**
         * Encodes the specified Action message. Does not implicitly {@link PlayStructures.Action.verify|verify} messages.
         * @function encode
         * @memberof PlayStructures.Action
         * @static
         * @param {PlayStructures.IAction} message Action message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Action.encode = function encode(message, writer) {
            if (!writer)
                writer = $Writer.create();
            if (message.Action != null && message.hasOwnProperty("Action"))
                writer.uint32(/* id 1, wireType 2 =*/10).string(message.Action);
            return writer;
        };

        /**
         * Encodes the specified Action message, length delimited. Does not implicitly {@link PlayStructures.Action.verify|verify} messages.
         * @function encodeDelimited
         * @memberof PlayStructures.Action
         * @static
         * @param {PlayStructures.IAction} message Action message or plain object to encode
         * @param {$protobuf.Writer} [writer] Writer to encode to
         * @returns {$protobuf.Writer} Writer
         */
        Action.encodeDelimited = function encodeDelimited(message, writer) {
            return this.encode(message, writer).ldelim();
        };

        /**
         * Decodes an Action message from the specified reader or buffer.
         * @function decode
         * @memberof PlayStructures.Action
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @param {number} [length] Message length if known beforehand
         * @returns {PlayStructures.Action} Action
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Action.decode = function decode(reader, length) {
            if (!(reader instanceof $Reader))
                reader = $Reader.create(reader);
            var end = length === undefined ? reader.len : reader.pos + length, message = new $root.PlayStructures.Action();
            while (reader.pos < end) {
                var tag = reader.uint32();
                switch (tag >>> 3) {
                case 1:
                    message.Action = reader.string();
                    break;
                default:
                    reader.skipType(tag & 7);
                    break;
                }
            }
            return message;
        };

        /**
         * Decodes an Action message from the specified reader or buffer, length delimited.
         * @function decodeDelimited
         * @memberof PlayStructures.Action
         * @static
         * @param {$protobuf.Reader|Uint8Array} reader Reader or buffer to decode from
         * @returns {PlayStructures.Action} Action
         * @throws {Error} If the payload is not a reader or valid buffer
         * @throws {$protobuf.util.ProtocolError} If required fields are missing
         */
        Action.decodeDelimited = function decodeDelimited(reader) {
            if (!(reader instanceof $Reader))
                reader = new $Reader(reader);
            return this.decode(reader, reader.uint32());
        };

        /**
         * Verifies an Action message.
         * @function verify
         * @memberof PlayStructures.Action
         * @static
         * @param {Object.<string,*>} message Plain object to verify
         * @returns {string|null} `null` if valid, otherwise the reason why it is not
         */
        Action.verify = function verify(message) {
            if (typeof message !== "object" || message === null)
                return "object expected";
            if (message.Action != null && message.hasOwnProperty("Action"))
                if (!$util.isString(message.Action))
                    return "Action: string expected";
            return null;
        };

        /**
         * Creates an Action message from a plain object. Also converts values to their respective internal types.
         * @function fromObject
         * @memberof PlayStructures.Action
         * @static
         * @param {Object.<string,*>} object Plain object
         * @returns {PlayStructures.Action} Action
         */
        Action.fromObject = function fromObject(object) {
            if (object instanceof $root.PlayStructures.Action)
                return object;
            var message = new $root.PlayStructures.Action();
            if (object.Action != null)
                message.Action = String(object.Action);
            return message;
        };

        /**
         * Creates a plain object from an Action message. Also converts values to other types if specified.
         * @function toObject
         * @memberof PlayStructures.Action
         * @static
         * @param {PlayStructures.Action} message Action
         * @param {$protobuf.IConversionOptions} [options] Conversion options
         * @returns {Object.<string,*>} Plain object
         */
        Action.toObject = function toObject(message, options) {
            if (!options)
                options = {};
            var object = {};
            if (options.defaults)
                object.Action = "";
            if (message.Action != null && message.hasOwnProperty("Action"))
                object.Action = message.Action;
            return object;
        };

        /**
         * Converts this Action to JSON.
         * @function toJSON
         * @memberof PlayStructures.Action
         * @instance
         * @returns {Object.<string,*>} JSON object
         */
        Action.prototype.toJSON = function toJSON() {
            return this.constructor.toObject(this, $protobuf.util.toJSONOptions);
        };

        return Action;
    })();

    return PlayStructures;
})();

module.exports = $root;
