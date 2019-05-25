/*
* @Author: Ximidar
* @Date:   2018-10-10 06:36:00
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-03-01 22:10:20
 */

package FileStructures

import (
	"time"

	"github.com/golang/protobuf/proto"
	"github.com/nats-io/go-nats"
)

const (
	// Name is the name of the program to call
	Name = "FILE_MANAGER."
	// Publishers

	// RunFileAction is the address to run a file action to
	RunFileAction = Name + "RunFileAction"

	// RequestLines is the address to get new lines from
	RequestLines = Name + "RequestLines"

	// UpdateFS will be called any time there is an update to the file system
	UpdateFS = Name + "UPDATE_FS"
	// FileProgress will be called every time the file progresses (0 - 100%)
	FileProgress = Name + "FILE_PROGRESS"
)

// NewFileAction will construct a FileAction object
func NewFileAction(action FileAction_Option, path string) (*FileAction, error) {
	fa := new(FileAction)
	fa.Action = action
	fa.Path = path
	return fa, nil
}

// NewFileActionFromMSG will construct a File action from a nats msg
func NewFileActionFromMSG(msg *nats.Msg) (*FileAction, error) {
	fa := new(FileAction)

	err := proto.Unmarshal(msg.Data, fa)

	if err != nil {
		return nil, err
	}

	return fa, nil
}

// SendAction will use the supplied nats conn to send the File Action
func SendAction(nc *nats.Conn, timeout time.Duration, fa *FileAction) (reply *nats.Msg, err error) {

	data, err := proto.Marshal(fa)

	if err != nil {
		return nil, err
	}

	reply, err = nc.Request(RunFileAction, data, timeout)

	if err != nil {
		return nil, err
	}

	return reply, nil

}
