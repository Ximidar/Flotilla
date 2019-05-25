/*
* @Author: Ximidar
* @Date:   2018-09-02 01:37:29
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-01-14 21:36:04
 */

package DataStructures

import (
	"encoding/json"

	nats "github.com/nats-io/go-nats"
)

// ReplyString will return a reply status with a success boolean and a message attached
type ReplyString struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// ReplyJSON is the same as ReplyString except the message will be a JSON byte array
type ReplyJSON struct {
	Success bool   `json:"success"`
	Message []byte `json:"message"`
}

// ConstructNegativeReplyString will create a replyString object with a failed status
func ConstructNegativeReplyString(mess string) ReplyString {
	rs := ReplyString{Success: false, Message: mess}
	return rs
}

// NewReplyStringFromMSG will take a nats msg and attempt to turn it into a ReplyString
func NewReplyStringFromMSG(msg *nats.Msg) (ReplyString, error) {
	rs := ReplyString{}
	err := json.Unmarshal(msg.Data, &rs)
	if err != nil {
		panic(err)
	}
	return rs, err
}

// PackageAndSendReplyJSON will package the reply and send it over nats
func PackageAndSendReplyJSON(repjson ReplyJSON, NC *nats.Conn, subj string) error {
	repBytes, err := json.Marshal(repjson)
	if err != nil {
		return err
	}

	err = NC.Publish(subj, repBytes)
	return err
}

// PackageAndSendReplyString will package a ReplyString and send it over nats
func PackageAndSendReplyString(repstring ReplyString, NC *nats.Conn, subj string) error {
	repBytes, err := json.Marshal(repstring)
	if err != nil {
		return err
	}

	err = NC.Publish(subj, repBytes)
	return err
}
