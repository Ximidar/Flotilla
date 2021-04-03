/*
* @Author: Ximidar
* @Date:   2019-02-16 23:05:57
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-21 02:47:58
 */

package ConfigStructures

import (
	"bytes"
	"encoding/gob"
	"encoding/json"
	"errors"
	"time"

	"github.com/nats-io/nats.go"
)

// ConfigReply will be sent to the config to retrieve certain values
type ConfigReply struct {
	// VarNames is the collection of vars that will be returned
	VarNames []string `json:"varnames"`

	// ToJSON will return a json object instead of the map
	ToJSON bool `json:"tojson"`
}

// RequestConfigJSON will return a JSON map of the keys asked for
func RequestConfigJSON(nc *nats.Conn, keys ...string) ([]byte, error) {
	cr := ConfigReply{
		VarNames: keys,
		ToJSON:   true,
	}

	// encode options to JSON
	mOptions, err := json.Marshal(cr)
	if err != nil {
		return []byte(nil), err
	}

	// Request those values
	reply, err := nc.Request(GetVar, mOptions, 5*time.Second)

	if err != nil {
		return []byte(nil), err
	}

	if len(reply.Data) == 0 {
		return []byte(nil), errors.New("no data returned")
	}
	return reply.Data, err

}

// EncodeBytes will encode an interface to bytes
func EncodeBytes(key interface{}) ([]byte, error) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(key)
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

// DecodeBytes will take a byteslice and decode it into an interface
func DecodeBytes(data []byte, out interface{}) error {

	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	err := dec.Decode(out)
	if err != nil {
		return err
	}
	return nil
}
