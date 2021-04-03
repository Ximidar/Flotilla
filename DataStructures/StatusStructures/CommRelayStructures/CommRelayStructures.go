/*
* @Author: Ximidar
* @Date:   2018-12-17 10:48:25
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-03-01 22:06:30
 */

package CommRelayStructures

import (
	"github.com/golang/protobuf/proto"
	"github.com/nats-io/nats.go"
)

const (
	// Name is the common name for this Nats interface
	Name = "CommRelay."

	// WriteLine is the subject name for sending raw message lines to the comm
	WriteLine = Name + "WriteLine"
)

// NewLine will construct a line from input
func NewLine(line string, lineNumber uint64, knownNumber bool) (*Line, error) {
	Nline := new(Line)
	Nline.Line = line
	Nline.LineNumber = lineNumber
	Nline.KnownNumber = knownNumber

	return Nline, nil
}

// SendLine will send the line over the Nats Conn
func SendLine(NC *nats.Conn, line *Line) error {
	lineM, err := proto.Marshal(line)
	if err != nil {
		return err
	}

	err = NC.Publish(WriteLine, lineM)
	if err != nil {
		return err
	}
	return nil
}

func NewLineFromMsg(bl []byte) (*Line, error) {
	l := new(Line)
	err := proto.Unmarshal(bl, l)
	return l, err

}
