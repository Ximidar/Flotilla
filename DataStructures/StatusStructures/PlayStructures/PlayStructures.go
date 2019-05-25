/*
* @Author: Ximidar
* @Date:   2018-12-21 17:47:08
* @Last Modified by:   Ximidar
* @Last Modified time: 2019-02-27 19:23:52
 */

package PlayStructures

import (
	"errors"
	"fmt"
	"time"

	nats "github.com/nats-io/go-nats"
	"github.com/ximidar/Flotilla/DataStructures"
)

const (

	// Actions

	// PLAY Use this action to play the selected file
	PLAY = "Play"
	//CANCEL Use this action to cancel the currently playing file
	CANCEL = "Cancel"
	//PAUSE use this action to pause the playing file
	PAUSE = "Pause"
	//RESUME use this action to resume the playing file
	RESUME = "Resume"
	// DONE Use this action to return the printer to idle
	DONE = "Done"
	// IDLE use this as a status for being ide
	IDLE = "Idle"
	// RESET will reset all flags and clear any errors that may have occured
	// If asked to reset during a print it will cancel the current playing file and
	// Reset the entire server
	RESET = "Reset"

	//////////////
	// Requests //
	//////////////

	// Name is the name of these requests
	Name = "PlayAction."

	// ControlAction will request a specific action to happen
	ControlAction = Name + "ControlAction"

	// RegisterNode will register a name with a node
	RegisterNode = Name + "RegisterNode"

	// RogerUp will change a registered Node to a status
	RogerUp = Name + "RogerUp"

	//GetStatus will get the status of of the nodes
	GetStatus = Name + "GetStatus"

	// SetError will set the error for the rest of the machine
	SetError = Name + "SetError"

	////////////////
	// Publishers //
	////////////////

	// PublishStatus will report a new status to nats
	PublishStatus = Name + "PublishAction"
)

// ProposeAction will ask the play status system to set a status
func ProposeAction(action string, nc *nats.Conn) error {
	proposedAction, err := NewPlayAction(action)
	if err != nil {
		return err
	}
	err = proposedAction.Send(nc)
	if err != nil {
		return err
	}
	return nil
}

// SetAnError will register an error with the PlayStatus system
func SetAnError(mess string, nc *nats.Conn) error {

	err := nc.Publish(SetError, []byte(mess))

	return err
}

// PlayAction will set an action on the server
type PlayAction struct {
	action string
}

// NewPlayAction will create an action to perform
func NewPlayAction(action string) (*PlayAction, error) {
	pa := new(PlayAction)
	err := pa.SetAction(action)
	if err != nil {
		return nil, err
	}
	return pa, nil

}

// SetAction protects the PlayAction.action variable
func (pa *PlayAction) SetAction(action string) error {

	err := CheckSubj(action)
	if err != nil {
		return fmt.Errorf("action: %v does not exist", action)
	}
	pa.action = action
	return nil

}

// Send will send the play action
func (pa *PlayAction) Send(nc *nats.Conn) error {
	if !pa.isActionSet() {
		return errors.New("action is not set")
	}

	resp, err := nc.Request(ControlAction, []byte(pa.action), 2*time.Second)
	if err != nil {
		return err
	}

	rs, err := DataStructures.NewReplyStringFromMSG(resp)
	if err != nil {
		return err
	}
	if !rs.Success {
		fmt.Printf("Success: %v\nMessage: %v\n", rs.Success, rs.Message)
		return errors.New(rs.Message)
	}
	return nil
}

func (pa *PlayAction) isActionSet() bool {
	if pa.action == "" {
		return false
	}
	return true
}

// Helpers

// CheckSubj will check if the subject is valid or not
func CheckSubj(subj string) error {
	ok := true
	switch subj {
	case PLAY:
	case CANCEL:
	case PAUSE:
	case RESUME:
	case DONE:
	case IDLE:
	case RESET:
	default:
		ok = false
	}
	if !ok {
		return fmt.Errorf("subj: %v does not exist", subj)
	}
	return nil
}
